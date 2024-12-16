services:
  postgres:
    container_name: bob-pg
    image: bitnami/postgresql:latest
    ports:
      - "5432:5432"
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_DB=bob-db
    volumes:
      - ./.docker/postgres:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres"]
      interval: 5s
      retries: 5
      timeout: 5s

  app:
    container_name: bob-api
    image: bob
    depends_on:
      postgres:
        condition: service_healthy
    build:
      context: .
      dockerfile: Dockerfile
    # depends on the api port
    ports:
      - "3000:3000"
    env_file:
      - .env