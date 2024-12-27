#  Base modular structure panoramic view
```
.
├─── cmd -> e.g.: all executables
│    └── api
|       └── main.go
├─── config
│    └── config.go -> e.g.: environment configurations
├─── db -> e.g.: database files in case of using a SQL database
│    └── migration -> e.g.: migration files
│        ├── 00001_create_x_table.up.sql
│        └── 00001_create_x_table.down.sql
├─── internal -> e.g.: application logic
│    └── module_a
│        └── core
│            ├── model -> e.g.: domain entities
│            ├── repository -> e.g.: repository contracts
│            └── service -> e.g.: implementation of use cases
│        └── database -> e.g.: persistence stuff
│            └── database_choosen -> e.g.: database choosen module
│                └── module_a_repository.go -> e.g.: implementation of a repository contract
│       ├── transport -> e.g.: communication module
│           └── rest -> e.g.: communication protocol
│               └── dto
│                   ├── response -> e.g.: response structs
│                   └── request -> e.g.: request structs with validation
│               └── endpoint -> e.g.: routes registrations and handlers
│           └── grpc
│        └── module_a.go -> e.g.: exposures all module actions to be used
│    ├── module_b
│        ├── ...
│        └── integration
│            └── provider
│                └── public_provider.go -> e.g.: implementation of integration contract
│   └── shared
│       ├── core -> e.g.: common core logic across modules
│       ├── database -> e.g.: common database logic across modules
│       ├── transport -> e.g.: common transport logic across modules
│       ├── helper -> e.g.: common helper logic across modules
│       └── integration -> e.g.: contracts for modules communication
│           └── module_a_b_integration.go -> e.g.: contract with minimum business logic
├─── pkg -> e.g.: external libraries without business logic
│    └── mysql -> e.g.: library logic and setup to be used
└── test -> e.g.: test helpers and more complex tests
    └── module_a
        ├── integration_test
        └── factory
    ├── module_b
    └── shared
        ├── helper
        ├── integration -> e.g.: actions in case of test modules needs integrated actions
        └── docker
```