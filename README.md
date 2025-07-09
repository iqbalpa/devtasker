# DevTasker

A simple task management backend application.

## Description

I hone my Go knowledge by implementing this system. The goal is to create a simple application but end-to-end. So basically, it will mimic a large scale system design, but with small application.

These are things that I want to implement here.

- [x] Backend Endpoints
- [x] HTTP Handler
- [x] API docs
- [x] Validation and Logging
- [x] Database
- [ ] Unit Testing
- [x] Monitoring (prometheus)
- [x] Docker
- [x] CI/CD
- [x] Reverse Proxy (nginx)
- [ ] GraphQL API
- [ ] gRPC Layer

## Features

- **RESTful API:** For creating, reading, updating, and deleting tasks.
- **Go Fiber Framework:** High-performance Go web framework.
- **PostgreSQL Database:** Robust and reliable database for storing task data.
- **GORM:** Developer-friendly ORM for Go.
- **Swagger Documentation:** Interactive API documentation.
- **Prometheus & Grafana:** For monitoring application metrics.
- **Nginx Reverse Proxy:** Manages and routes traffic to the application and monitoring services.
- **Docker-Compose:** For easy setup and deployment of the entire application stack.
- **Live Reloading:** With Air for a better development experience.

## Getting Started

To get the application running locally, follow these steps:

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/devtasker.git
   cd devtasker
   ```

2. **Set up environment variables:**
   Create `.devtasker.env`, `.grafana.env`, and `.postgres.env` from the provided `.env.example` files and customize them if needed.

3. **Run with Docker Compose:**
   ```bash
   docker-compose up -d
   ```

The application will be accessible at `http://localhost`.

## API Endpoints

All endpoints are prefixed with `/api`.

| Method | Endpoint      | Description          |
| ------ | ------------- | -------------------- |
| POST   | `/task`       | Create a new task    |
| GET    | `/task`       | Get all tasks        |
| GET    | `/task/:id`   | Get a task by ID     |
| PATCH  | `/task/:id`   | Update a task        |
| DELETE | `/task/:id`   | Delete a task        |

- **Swagger Docs:** `http://localhost/doc`
- **Health Check:** `http://localhost/health`

## Monitoring

- **Prometheus:** `http://localhost:9090`
- **Grafana:** `http://localhost:3001`

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature-name`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature/your-feature-name`).
6. Open a pull request.
