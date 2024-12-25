services:
  postgres:
    image: bitnami/postgresql:latest
    ports:
      - "${DATABASE_PORT}:5432"
    environment:
      - POSTGRES_USER=${DATABASE_USER}
      - POSTGRES_PASSWORD=${DATABASE_PASSWORD}
      - POSTGRES_DB=${DATABASE_NAME}
    volumes:
      - ./.docker/postgres:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres"]
      interval: 5s
      retries: 5
      timeout: 5s

  app:
    image: bob-app
    depends_on:
      postgres:
        condition: service_healthy
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "${SERVER_PORT}:3000"
    env_file:
      - .env