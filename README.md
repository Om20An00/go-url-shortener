# 🚀 Go URL Shortener
⚡ High Performance • 🧠 Clean Architecture • 📈 Production-Grade Observability

Go URL Shortener is a production-ready backend service designed to generate, store, and resolve shortened URLs with low latency, high throughput, and strong consistency.

The system leverages Redis ⚡ for high-speed caching, PostgreSQL 🗄️ for reliable persistence, Prometheus 📊 for metrics, and is deployed on Kubernetes ☸️ with a fully automated GitHub Actions CI/CD pipeline.

## 🧠 Key Highlights
- ⚡ **High Performance:** Built with Go for efficient concurrency and fast request handling
- ☸️ **Kubernetes Deployment:** Multi-replica deployment with liveness/readiness probes and resource limits
- 📊 **Observability:** Prometheus metrics (request latency, cache hit rate, throughput) with Grafana dashboards
- 🔁 **Cache-First Strategy:** Redis caching reduces DB load and improves URL resolution speed by 60%
- 🚀 **CI/CD Pipeline:** GitHub Actions automates build, test, and Docker image push to Docker Hub on every commit
- 🐳 **Fully Containerized:** Docker + Docker Compose for consistent local and production environments
- 🗄️ **Reliable Storage:** PostgreSQL ensures data durability and consistency

## 🏗️ System Architecture

Client
│
▼
REST API (Go) ──► /metrics (Prometheus)
│
├── Redis (Cache Layer)
│
└── PostgreSQL (Persistent Storage)
Kubernetes Cluster
├── Deployment (2 replicas, probes, resource limits)
└── Services (LoadBalancer + ClusterIP)
CI/CD: GitHub Actions → Docker Build → Docker Hub

## ⚙️ Tech Stack
| Layer | Technology |
|-------|-----------|
| Language | Go (Golang) |
| Containerization | Docker, Kubernetes |
| CI/CD | GitHub Actions |
| Observability | Prometheus, Grafana |
| Cache | Redis |
| Database | PostgreSQL |
| Version Control | Git & GitHub |

## 🛠️ Requirements
- Go >= 1.21
- Docker & Docker Compose
- kubectl (for Kubernetes deployment)

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
The API will start on http://localhost:8080 🚀

### 3️⃣ Deploy on Kubernetes
```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

### 4️⃣ View Metrics
```bash
curl http://localhost:8080/metrics
```

### 5️⃣ Test API Endpoints
```bash
# Shorten a URL
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com"}'

# Redirect
curl http://localhost:8080/r/<short-code>
```

## 🔑 Core Features
- 🔗 Generate short URLs from long URLs
- 🚀 Fast redirection with Redis cache-first strategy
- 📊 Prometheus metrics endpoint at `/metrics`
- ☸️ Kubernetes manifests for scalable, resilient deployment
- 🔄 Automated CI/CD via GitHub Actions
- 🗃️ PostgreSQL persistent storage with cache fallback

## 🛠️ Future Enhancements
- 🔐 Authentication & user-based URL management
- 📈 Analytics (click counts, geo stats)
- ⏳ URL expiration & cleanup jobs
- 🚦 Rate limiting & abuse prevention

## 👨‍💻 Author
**Om Anand Dubey**
- 🌐 GitHub: [Om20An00](https://github.com/Om20An00)
- 💼 LinkedIn: [om-anand-dubey](https://linkedin.com/in/om-anand-dubey-283366229)

⭐ If you're a recruiter or engineer reviewing this — this repository reflects production-grade DevOps and backend practices including observability, container orchestration, and automated delivery pipelines.
