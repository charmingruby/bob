#  Base modular structure panoramic view
```
.
├─── cmd // all executables
│   └── api
|       └── main.go
├─── config
│   └── config.go // environment configurations
├─── db // database files in case of using a SQL database
│   └── migration // migration files
│       ├── 00001_create_x_table.up.sql
│       └── 00001_create_x_table.down.sql
├─── internal // application logic
│   └── module_a
│       └── core
│           ├── model // domain entities
│           ├── repository // repository contracts
│           └── service // implementation of use cases
│       └── database // database stuff
│           └── database_choosen // database choosen module
│               └── module_a_repository.go // implementation of a repository contract
│       ├── transport // communication module
│           └── rest // communication protocol
│               └── dto
│                   └── response // response structs
│                   └── request // request structs with validation
│               └── endpoint // routes registrations and handlers
│           └── grpc
│       └── module_a.go // exposures all module actions to be used
│   ├── module_b
│       ├── ...
│       └── integration
│           └── provider
│               └── public_provider.go // implementation of integration contract
│   └── shared
│       ├── core // common core logic across modules
│       ├── database // common database logic across modules
│       ├── transport // common transport logic across modules
│       ├── helper // common helper logic across modules
│       └── integration // contracts for modules communication
│           └── module_a_b_integration.go // contract with minimum business logic
├─── pkg // external libraries without business logic
│   └── mysql // library logic and setup to be used
└── test // test helpers and more complex tests
│   └── module_a
│       ├── integration_test
│       └── factory
│   ├── module_b
│   └── shared
│       ├── helper
│       ├── integration // actions in case of test modules needs integrated actions
        └── container
```