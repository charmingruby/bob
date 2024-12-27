# Bob

Bob is a Go tool designed to streamline the creation of modular and scalable applications using predefined templates.

## Description
bob is a tool that provides a set of predefined templates to facilitate the creation of projects in various contexts. It aims to accelerate the initial setup of projects, whether for enterprise applications or proof of concepts (PoC), ensuring a solid foundation for further development.

## Philosophy
The philosophy of bob is to be able to grow into different contexts in an opinionated manner, ensuring an evolutionary architecture and allowing the easy composition of new components in the architecture. And from just the beginning, we have the following design principles:

* Keep it simple
* Easy to extend
* Try best to be friendly to the business logic development, encapsulate the complexity
* Ability to create various templates for different use cases, such as REST, gRPC, and serverless
* Modular design for easy integration and maintenance


## Installation

### Get Started

1. Run the installation script:
   ```shell
   curl -sSL https://raw.githubusercontent.com/charmingruby/bob/main/install.sh | bash
   ```

2. Create a configuration file:
   ```shell
   bob init
   ```

3. Optionally, you can change the configs for your project;

4. Create a project:
   ```shell
   bob create template rest base -database dynamo
   ```

5. Run go mod tidy to ensure all dependencies are installed:
   ```shell
   go mod tidy
   ```

6. Set the environment variables, based on .env.example

7. Run docker compose:
   ```shell
   docker compose up -d
   ```

8. Run application:
   ```shell
   go run ./cmd/api/main.go
   ```

9. Call the api:
   ```shell
   curl -X POST http://localhost:3000/example/ping -H "Content-Type: application/json" -d '{"name": "bob"}'
   ```

## Give a Star! ⭐

If you like this project or are using it to learn or using for your own solution or purpose, give it a star. Your support matters!