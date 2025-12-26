package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"go-url-shortener/internal/storage"

	"github.com/redis/go-redis/v9"
)

func generateShortCode(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	// Initialize PostgreSQL and Redis
	storage.InitPostgres()
	storage.ConnectRedis()

	rand.Seed(time.Now().UnixNano())

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Shorten URL endpoint
	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			URL string `json:"url"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		short := generateShortCode(6)
		if err := storage.InsertURL(short, req.URL); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Store in Redis cache as well
		storage.Rdb.Set(storage.RedisCtx, short, req.URL, 24*time.Hour)

		json.NewEncoder(w).Encode(map[string]string{"short_url": short})
	})

	// Redirect endpoint
	http.HandleFunc("/r/", func(w http.ResponseWriter, r *http.Request) {
		short := r.URL.Path[len("/r/"):]
		var original string

		// Try Redis cache first
		original, err := storage.Rdb.Get(storage.RedisCtx, short).Result()
		if err == redis.Nil { // not in cache
			original, err = storage.GetOriginalURL(short)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			// Save to Redis for future requests
			storage.Rdb.Set(storage.RedisCtx, short, original, 24*time.Hour)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, original, http.StatusFound)
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
