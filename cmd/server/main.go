package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"go-url-shortener/internal/storage"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)
	redirectLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "redirect_latency_seconds",
			Help:    "Latency of redirect requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"cache_hit"},
	)
	cacheHitsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "cache_hits_total",
			Help: "Total number of Redis cache hits",
		},
	)
	cacheMissesTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "cache_misses_total",
			Help: "Total number of Redis cache misses",
		},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(redirectLatency)
	prometheus.MustRegister(cacheHitsTotal)
	prometheus.MustRegister(cacheMissesTotal)
}

func generateShortCode(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	storage.InitPostgres()
	storage.ConnectRedis()

	rand.Seed(time.Now().UnixNano())

	// Prometheus metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Shorten URL endpoint
	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			httpRequestsTotal.WithLabelValues(r.Method, "/shorten", "405").Inc()
			return
		}

		var req struct {
			URL string `json:"url"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			httpRequestsTotal.WithLabelValues(r.Method, "/shorten", "400").Inc()
			return
		}

		short := generateShortCode(6)
		if err := storage.InsertURL(short, req.URL); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			httpRequestsTotal.WithLabelValues(r.Method, "/shorten", "500").Inc()
			return
		}

		storage.Rdb.Set(storage.RedisCtx, short, req.URL, 24*time.Hour)
		httpRequestsTotal.WithLabelValues(r.Method, "/shorten", "200").Inc()
		json.NewEncoder(w).Encode(map[string]string{"short_url": short})
	})

	// Redirect endpoint
	http.HandleFunc("/r/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		short := r.URL.Path[len("/r/"):]
		var original string
		cacheHit := "true"

		original, err := storage.Rdb.Get(storage.RedisCtx, short).Result()
		if err == redis.Nil {
			cacheHit = "false"
			cacheMissesTotal.Inc()
			original, err = storage.GetOriginalURL(short)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				httpRequestsTotal.WithLabelValues(r.Method, "/r/", "404").Inc()
				return
			}
			storage.Rdb.Set(storage.RedisCtx, short, original, 24*time.Hour)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			httpRequestsTotal.WithLabelValues(r.Method, "/r/", "500").Inc()
			return
		} else {
			cacheHitsTotal.Inc()
		}

		redirectLatency.WithLabelValues(cacheHit).Observe(time.Since(start).Seconds())
		httpRequestsTotal.WithLabelValues(r.Method, "/r/", "302").Inc()
		http.Redirect(w, r, original, http.StatusFound)
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
