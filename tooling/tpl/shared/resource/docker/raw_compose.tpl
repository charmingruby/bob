services:
  app:
    image: bob-app
    ports:
      - "${SERVER_PORT}:3000"
    env_file:
      - .env