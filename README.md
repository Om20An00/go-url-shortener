🚀 Go URL Shortener

⚡ High Performance • 🧠 Clean Architecture • 📈 Scalable Backend

Go URL Shortener is a production-ready backend service designed to generate, store, and resolve shortened URLs with low latency, high throughput, and strong consistency ⚙️🚀.

The system follows clean architecture principles 🧩, leverages Redis ⚡ for high-speed caching, PostgreSQL 🗄️ for reliable persistence, and exposes a RESTful API 🌐 built in Go for maximum performance and scalability

🧠 Key Highlights (Recruiter Focused)

⚡ High Performance: Built with Go for efficient concurrency and fast request handling

🧩 Scalable Design: Stateless REST APIs with Redis-based caching

🗄️ Reliable Storage: PostgreSQL ensures data durability and consistency

🔁 Cache-First Strategy: Reduces DB load and improves URL resolution speed

🛡️ Clean & Maintainable Code: Modular structure following best practices

📈 Production-Oriented: Designed with extensibility and real-world usage in mind

🏗️ System Architecture
```
Client
  │
  ▼
REST API (Go)
  │
  ├── Redis (Cache)
  │
  └── PostgreSQL (Persistent Storage)
```
⚙️ Tech Stack
Layer	Technology
Language	Go (Golang)
API	RESTful APIs
Cache	Redis
Database	PostgreSQL
Architecture	Clean / Layered Architecture
Version Control	Git & GitHub
🔑 Core Features

🔗 Generate short URLs from long URLs

🚀 Fast redirection using Redis caching

🗃️ Persistent storage with PostgreSQL

🔄 Cache fallback to database on cache miss

🧪 Easily testable REST endpoints

📦 Modular and extensible codebase



It demonstrates:

Real backend system design

Cache + database coordination

API design and data flow

Performance-conscious engineering

Code written with production standards

💡 The same architectural patterns used here are applicable to real systems at scale.

🛠️ Future Enhancements

🔐 Authentication & user-based URL management

📊 Analytics (click counts, geo stats)

⏳ URL expiration & cleanup jobs

🚦 Rate limiting & abuse prevention

🐳 Dockerized deployment

👨‍💻 Author

Om Anand Dubey
🎓 Computer Science Graduate
💻 Backend & Systems Enthusiast
🌐 GitHub: Om20An00

⭐ Final Note

If you’re a recruiter or engineer reviewing this project:

This repository reflects industry-grade backend development practices, not just academic code.

Feel free to ⭐ the repo or reach out for discussion!
