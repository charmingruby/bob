services:
  app:
    container_name: bob-api
    image: bob
    build:
      context: .
      dockerfile: Dockerfile
    # depends on the api port
    ports:
      - "3000:3000"
    env_file:
      - .env