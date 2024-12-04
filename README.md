## KuReDoPoGo

Kubernetes Redis Docker Postgres Golang—KuReDoPoGo is a backend API service built using Golang and Gin Framework. It uses Redis for rate limiting and Postgres for storing data. It is deployed on Kubernetes. For local development, it uses Docker Compose.

## Using and Developing

### Running Locally

#### Prerequisites
- Go
- Postgres (local, remote, or Docker instance)
- Redis (local, remote, or Docker instance)

1. Copy the `.env.example` file to `.env` and update the values like postgres connection string, redis connection string, and other values as needed. You can use the following command to copy the file:

```bash
cp .env.example .env
```

2. Install the Go dependencies and run the application:

```bash
go mod download
go run main.go
```

You can access the application at `http://localhost:8080`.

### Running with Docker Compose

#### Prerequisites

- Docker
- Docker Compose

1. Run the following command to start the application:

```bash
docker compose up
```
If you're using the latest version of Docker Compose, you can leverage the Compose Watch feature. This will automatically rebuild the application when code changes are made:

```bash
docker compose up --watch
docker compose up --build --watch
```

You can access the application at `http://localhost:8080` and Postgres at `http://localhost:5432`.

## License

This project is licensed under the [GNU General Public License v3.0](LICENSE).

## Security

If you discover a security vulnerability within this project, please refer to the [security policy](SECURITY.md) for more information.
