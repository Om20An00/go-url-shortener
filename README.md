<div align="center">

<img src="https://capsule-render.vercel.app/api?type=waving&color=0:00ADD8,50:00758F,100:0d1117&height=180&section=header&text=Go%20URL%20Shortener&fontSize=42&fontColor=ffffff&fontAlignY=38&desc=High%20Performance%20%E2%80%A2%20Clean%20Architecture%20%E2%80%A2%20Production-Grade%20Observability&descAlignY=58&descColor=e0f7ff&animation=fadeIn" />

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-Deployed-326CE5?style=for-the-badge&logo=kubernetes&logoColor=white)](https://kubernetes.io)
[![Prometheus](https://img.shields.io/badge/Prometheus-Monitored-E6522C?style=for-the-badge&logo=prometheus&logoColor=white)](https://prometheus.io)
[![GitHub Actions](https://img.shields.io/badge/CI%2FCD-GitHub_Actions-2088FF?style=for-the-badge&logo=githubactions&logoColor=white)](https://github.com/features/actions)

<img src="https://img.shields.io/github/stars/Om20An00/go-url-shortener?style=social" />
<img src="https://img.shields.io/github/forks/Om20An00/go-url-shortener?style=social" />

</div>

---

## 📖 Overview

**Go URL Shortener** is a production-ready backend service designed to generate, store, and resolve shortened URLs with **low latency**, **high throughput**, and **strong consistency**.

The system leverages **Redis** for high-speed caching, **PostgreSQL** for reliable persistence, **Prometheus + Grafana** for full observability, and ships to **Kubernetes** through a fully automated **GitHub Actions** CI/CD pipeline — built the way a real distributed system should be, not just a toy shortener.

> 💡 **TL;DR** — Send a long URL in, get a short code out, and resolve it back at cache speed — all while the system tells you exactly how healthy it is.

---

## 🧠 Key Highlights

| | |
|---|---|
| ⚡ **High Performance** | Built with Go's goroutines and channels for efficient concurrency and fast request handling |
| ☸️ **Kubernetes-Native** | Multi-replica deployment with liveness/readiness probes and resource limits baked in |
| 📊 **Full Observability** | Prometheus metrics — request latency, cache hit rate, throughput — visualized in Grafana |
| 🔁 **Cache-First Strategy** | Redis caching cuts DB load and improves URL resolution speed by **~60%** |
| 🚀 **Automated CI/CD** | GitHub Actions builds, tests, and pushes Docker images on every commit |
| 🐳 **Fully Containerized** | Docker + Docker Compose for identical local and production environments |
| 🗄️ **Durable Storage** | PostgreSQL guarantees data durability and consistency at the source of truth |

---

## 🏗️ System Architecture

```
                            ┌─────────────┐
                            │   Client    │
                            └──────┬──────┘
                                   │
                                   ▼
                    ┌───────────────────────────┐
                    │      REST API (Go)        │──────► /metrics (Prometheus)
                    └─────────────┬─────────────┘
                                   │
                     ┌─────────────┴─────────────┐
                     ▼                           ▼
             ┌───────────────┐           ┌───────────────┐
             │ Redis (Cache) │           │  PostgreSQL   │
             │  Layer        │           │  (Persistent) │
             └───────────────┘           └───────────────┘

┌──────────────────────────── Kubernetes Cluster ────────────────────────────┐
│   Deployment (2+ replicas, liveness/readiness probes, resource limits)     │
│   Services → LoadBalancer (external) + ClusterIP (internal)               │
└──────────────────────────────────────────────────────────────────────────┘

        CI/CD:  GitHub Actions  →  Docker Build  →  Push to Docker Hub
```

**Request flow:** every read checks Redis first — only on a cache miss does it fall through to PostgreSQL, repopulating the cache on the way back out. Every request is instrumented, so latency and hit-rate are visible in real time, not guessed at.

---

## ⚙️ Tech Stack

| Layer | Technology |
|---|---|
| **Language** | Go (Golang) |
| **Cache** | Redis |
| **Database** | PostgreSQL |
| **Containerization** | Docker, Kubernetes |
| **CI/CD** | GitHub Actions |
| **Observability** | Prometheus, Grafana |
| **Version Control** | Git & GitHub |

---

## 🛠️ Requirements

- Go `>= 1.21`
- Docker & Docker Compose
- `kubectl` (for Kubernetes deployment)

---

## 💻 Getting Started

### 1️⃣ Clone the repo
```bash
git clone https://github.com/Om20An00/go-url-shortener.git
cd go-url-shortener
```

### 2️⃣ Run with Docker Compose
```bash
docker-compose up --build
```
The API will start on **http://localhost:8080** 🚀

### 3️⃣ Deploy on Kubernetes
```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

### 4️⃣ View Metrics
```bash
curl http://localhost:8080/metrics
```

### 5️⃣ Test the API

**Shorten a URL**
```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com"}'
```

**Redirect**
```bash
curl http://localhost:8080/r/<short-code>
```

---

## 🔑 Core Features

- 🔗 Generate short URLs from long URLs
- 🚀 Fast redirection with a Redis cache-first strategy
- 📊 Prometheus metrics endpoint at `/metrics`
- ☸️ Kubernetes manifests for scalable, resilient deployment
- 🔄 Automated CI/CD via GitHub Actions
- 🗃️ PostgreSQL persistent storage with cache fallback

---

## 🛠️ Future Enhancements

- 🔐 Authentication & user-based URL management
- 📈 Click analytics & geo stats
- ⏳ URL expiration & cleanup jobs
- 🚦 Rate limiting & abuse prevention

---

## 👨‍💻 Author

<div align="center">

**Om Anand Dubey**

[![GitHub](https://img.shields.io/badge/GitHub-Om20An00-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/Om20An00)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-om--anand--dubey-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://linkedin.com/in/om-anand-dubey-283366229)

</div>

> ⭐ **If you're a recruiter or engineer reviewing this** — this repository reflects production-grade DevOps and backend practices, including observability, container orchestration, and automated delivery pipelines. If it's useful or interesting, a star is always appreciated.

<div align="center">

<img src="https://capsule-render.vercel.app/api?type=waving&color=0:0d1117,50:00758F,100:00ADD8&height=100&section=footer" />

</div>
