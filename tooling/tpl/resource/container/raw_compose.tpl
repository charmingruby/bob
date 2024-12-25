services:
  app:
    image: bob-app
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "${SERVER_PORT}:3000"
    env_file:
      - .env