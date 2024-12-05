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

### Using the API

One the application is running, you can use the following endpoints to interact with the API. By default, the rate limit is set to 10 per hour. You can change this in the `.env` file.

#### Endpoints

- `GET /users`: Get all users
- `GET /users/:id`: Get a user by ID
- `POST /users`: Create a new user
- `PUT /users/:id`: Update a user by ID
- `DELETE /users/:id`: Delete a user by ID
- `GET /health`: Get the health status of the application

POST and PUT requests require a JSON body with the following structure:

```json
{
  "name": "John Doe",
  "email": "test@test.com"
}
```

## License

This project is licensed under the [GNU General Public License v3.0](LICENSE).

## Security

If you discover a security vulnerability within this project, please refer to the [security policy](SECURITY.md) for more information.
