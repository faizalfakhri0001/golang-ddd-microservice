<h1 align="center">Go Gin DDD Boilerplate</h1>

<p align="center">
  <!-- <a href="#"><img src="https://github.com/resotto/goilerplate/workflows/test/badge.svg" /></a>
  <a href="#"><img src="https://goreportcard.com/badge/github.com/resotto/goilerplate" /></a>
  <a href="#"><img src="https://pkg.go.dev/badge/github.com/resotto/goilerplate" /></a>
  <a href="#"><img src="https://img.shields.io/badge/chat-on%20issue-yellow"></a> -->
  <a href="#"><img src="https://img.shields.io/badge/license-GPL%20v3.0-brightgreen.svg" /></a>
</p>

<p align="center">
  Clean Boilerplate of Go, Domain-Driven Design, Clean Architecture, Gin and GORM.  
</p>

---

What is this?

- **Good example of Go with Clean Architecture.**
- **Rocket start guide of Go, Domain-Driven Design, [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html), [Gin](https://github.com/gin-gonic/gin), and [GORM](https://github.com/go-gorm/gorm)**.

Note

- Default application/test code is trivial because you will write cool logic.

---

## Table of Contents

- [Getting Started](#getting-started)
- [Endpoints](#endpoints)
- [Package Structure](#package-structure)
- [How to Cross the Border of Those Layers](#how-to-cross-the-border-of-those-layers)
- [Dependency Injection](#dependency-injection)
- [How to Start with Goilerplate](#how-to-start-with-goilerplate)
- [Testing](#testing)
- [Naming Convention](#naming-convention)
- [With Gochk](#with-gochk)
- [With PostgreSQL](#with-postgresql)
- [Feedbacks](#feedbacks)
- [License](#license)
- [Author](#author)

## Getting Started

```zsh
go mod init {repository}/{service}        # might take few minutes
go get -u github.com/gin-gonic/gin
go get -u github.com/spf13/viper
go get -u gorm.io/gorm
go get gorm.io/driver/mysql
go get -u github.com/sirupsen/logrus
go install github.com/google/wire/cmd/wire@latest
go get github.com/stretchr/testify/assert
go install github.com/cosmtrek/air@latest


`install another package`
```

## Endpoints

- With Template
  - `GET /`
    - NOTICE: Following path is from CURRENT directory, so please run Gin from root directory.
      ```go
      r.LoadHTMLGlob("internal/app/adapter/view/*")
      ```
- With Public API of bitbank
  - `GET /ticker`
  - `GET /candlestick`
    - NOTICE: This works from 0AM ~ 3PM (UTC) due to its API constraints.
- With PostgreSQL
  - [NOTICE: Please run postgres container first with this step.](#with-postgresql)
    - `GET /parameter`
    - `GET /order`

## Package Structure

```zsh
.
├── LICENSE
├── README.md
├── build                                         # Packaging and Continuous Integration
│   ├── Dockerfile
│   └── init.sql
├── cmd                                           # Main Application
│   └── app
│       └── main.go
├── internal                                      # Private Codes
│   └── app
│       ├── {feature}                             # Feature Module
│       │   ├── interfaces                        # Interfaces specific for the feature
│       │   │   ├── {feature}_controller.go       # Controller Implementation
│       │   │   └── {feature}_controller_impl.go  # Controller Contract Abstraction
│       │   ├── infrastructure                    # Infrastructure layer for the feature
│       │   │   ├── {feature}_repository_impl.go  # Repository Implementation
│       │   │   └── {feature}_service_impl.go     # Service Implementation
│       │   └── domain                            # Domain layer for the feature
│       │       ├── {feature}_repository.go       # Repository Contract Abstraction
│       │       └── {feature}_service.go          # Service Contract Abstraction
│       └── shared                                # Shared Component
│           ├── {db}                              # Database (Replace {db} with Database you will use)
│           │   ├── model                         # Model
│           │   └── conn.go                       # Database Connection
│           ├── util.go                           # Utility Function
│           └── constant.go                       # Shared Constant
└── mocks                                         # Test Data
    └── feature_mock.go
```

- **Domain Layer:**
  - Purpose: The domain layer contains the core business logic and domain models of the application. It encapsulates the concepts, rules, and behaviors that are central to the problem domain.
  - Responsibilities:
    - Define domain entities: Objects that represent the key concepts and entities within the problem domain.
    - Implement domain logic: Business rules, validations, and behaviors that govern how entities interact and evolve over time.
    - Define domain events: Events that capture significant changes or state transitions within the domain.
- **Application Layer:**
  - Purpose: The application layer orchestrates interactions between the domain layer and external systems or interfaces. It coordinates the execution of use cases and manages application-specific logic.
  - Responsibilities:
    - Define application services: High-level operations or use cases that drive the application's behavior.
    - Define application services: High-level operations or use cases that drive the application's behavior.
    - Define application services: High-level operations or use cases that drive the application's behavior.
- **Adapter (Infrastructure) Layer:**
  - Purpose: The adapter layer provides a bridge between the application's internal components and external systems or frameworks. It handles communication with external interfaces, such as databases, web services, or user interfaces.
  - Responsibilities:
    - Implement interfaces with external systems: Adapters translate domain concepts into formats understandable by external systems and vice versa.
    - Handle input/output: Convert external requests into internal commands and vice versa, ensuring proper data transformation and validation.
    - Manage infrastructure concerns: Deal with aspects like database access, HTTP routing, messaging, and external service integrations.
