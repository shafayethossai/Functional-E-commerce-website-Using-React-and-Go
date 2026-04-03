# Go Backend E-commerce Project Documentation

## 📚 Overview

This is a complete **backend API** for an e-commerce application built with **Go (Golang)**. It demonstrates professional backend architecture using:
- **Domain-Driven Design (DDD)** for clean code organization
- **Interface-based programming** to avoid tight coupling
- **Middleware pattern** for cross-cutting concerns
- **PostgreSQL database** with migrations
- **JWT Authentication** for secure API access
- **Dependency Injection** through constructor functions

---

## 🔧 Essential Files & Configuration

### 1. **go.mod** - Module Definition
```
Purpose: Dependency management file for Go projects
Location: Project root
Contains:
- Module name: "first-program"
- Go version: 1.25.5
- List of all external packages your project depends on
```

**Why needed?**
- Go needs to know what external libraries you're using
- Ensures version consistency across environments
- Similar to `package.json` (Node.js) or `requirements.txt` (Python)

**Example dependencies in your project:**
- `github.com/lib/pq` - PostgreSQL driver (for database)
- `github.com/jmoiron/sqlx` - SQL utilities (for database queries)
- `github.com/joho/godotenv` - Load .env files
- `golang-jwt/jwt/v5` - JWT authentication (implied)

---

### 2. **go.sum** - Dependency Lock File
```
Purpose: Ensures reproducible builds
Location: Project root
Contains: Hash values of each dependency version
```

**Why needed?**
- Locks specific versions so everyone uses the same code
- Prevents unexpected breaking changes from dependency updates
- Go runs: `go mod tidy` → automatically generates/updates this file

---

### 3. **.env** - Environment Configuration
```
Purpose: Store sensitive credentials and environment variables
Location: Project root (NOT in git repository)
Contains: Database credentials, JWT secret, port number, etc.

Example .env file:
```
VERSION=1.0.0
SERVICE_NAME=ecommerce-api
HTTP_PORT=8080
JWT_SECRET_KEY=your-secret-key-here

DB_HOST=localhost
DB_PORT=5432
DB_NAME=ecommerce
DB_USER=postgres
DB_PASSWORD=your-password
DB_ENABLE_SSL=false
```

**Why needed?**
- Separates secrets from code (security best practice)
- Allows different configs for dev/production without code changes
- `.env` file is loaded by `config.go` using `godotenv` package

---

### 4. **main.go** - Application Entry Point
```
Current Status: Mostly commented out (old code from learning)
Current Architecture: The active entry point is actually in cmd/serve.go
```

**What happens when you run the app:**

1. Go looks for the `main()` function in `main.go`
2. Executes the main logic (currently commented)
3. The real startup code is in `cmd/serve.go` (the `Serve()` function)

**Typical flow (if main.go was active):**
```go
func main() {
    // 1. Load configuration from .env
    // 2. Call Serve() from cmd package
    // 3. Start the HTTP server
}
```

---

## 🏗️ Project Architecture - Layer by Layer

### **Architecture Diagram - Data Flow:**

```
Client (Browser/Postman)
        ↓
   HTTP Request
        ↓
rest/server.go (HTTP Server)
        ↓
rest/middlewares/ (Logger, CORS, JWT Auth)
        ↓
rest/handlers/ (API Handlers)
        ↓
product/ or user/ (Business Logic/Services)
        ↓
repo/ (Database Access Layer)
        ↓
domain/ (Data Models)
        ↓
PostgreSQL Database
```

---

## 📂 Folder Structure & Purpose

### 1. **domain/** - Data Models & Entities
```
Purpose: Define the core business objects
Files: user.go, product.go

What is a Struct?
- A collection of related fields bundled together
- Like a blueprint or template for objects
- Example: User has id, email, password, etc.

Example from user.go:
```go
type User struct {
    ID          int    `json:"id" db:"id"`
    FirstName   string `json:"first_name" db:"first_name"`
    Email       string `json:"email" db:"email"`
    Password    string `json:"password" db:"password"`
    IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}
```

**Tags explained:**
- `json:"id"` - When converting to/from JSON, use this key
- `db:"id"` - When reading/writing to database, map to this column name

**Flow:** Database → Struct → JSON → Frontend

---

### 2. **config/** - Configuration Management
```
Purpose: Load settings from .env file into Go structs
File: config.go

Process:
1. Reads .env file using godotenv package
2. Extracts values into Config struct
3. Makes config available globally to other packages

Example structures:
```go
type Config struct {
    Version      string
    ServiceName  string
    HttpPort     int
    JwtSecretKey string
    DB           *DBConfig  // Nested struct
}

type DBConfig struct {
    Host          string
    Port          int
    Name          string
    User          string
    Password      string
    EnableSSLMODE bool
}
```

**Why separate struct?**
- Clean organization
- Keeps database config details together
- Easy to pass around as `config.DB`

---

### 3. **infra/** - Infrastructure/Database Connection
```
Purpose: Handle database connections and migrations
Files: connection.go, migrate.go

What happens:
1. connection.go: Creates connection string and connects to PostgreSQL
2. migrate.go: Runs SQL migration files to create/update tables

Flow:
```
.env variables (username, password, host, port)
        ↓
GetConnectionString() (formats into PostgreSQL connection string)
        ↓
sqlx.Connect() (actual database connection)
        ↓
Migration files (000001-create-user.up.sql, etc.)
        ↓
Tables created in database
```

**Example connection string:**
```
user=postgres password=secret host=localhost port=5432 dbname=ecommerce sslmode=disable
```

---

### 4. **migrations/** - Database Version Control
```
Purpose: Track database schema changes over time
Files: 
- 000001-create-user.up.sql      (creates user table)
- 000001-create-user.down.sql    (deletes user table - rollback)
- 000002-create-products.up.sql  (creates product table)
- 000002-create-products.down.sql (deletes product table)

Why migrations?
- Database changes are tracked like code in git
- Can rollback to previous schema if needed
- Team members have consistent database structure
- Automated deployment process

How they work:
```
Up files:   Applied when moving forward
             CREATE TABLE users (id INT, email VARCHAR...)

Down files: Applied when rolling back
             DROP TABLE users
```

---

### 5. **repo/** - Repository Pattern (Data Access Layer)
```
Purpose: Abstract database operations
Files: user.go, product.go

What is the Repository Pattern?
- Isolates database logic from business logic
- Makes testing easier (can mock the repository)
- If you switch databases (PostgreSQL → MongoDB), only change repo
- Loose coupling!

Example: ProductRepo interface in repo/product.go

Interface Definition (What methods must exist):
```go
type ProductRepo interface {
    Create(p domain.Product) (*domain.Product, error)
    Update(p domain.Product) error
    Delete(id int) error
    Get(id int) (*domain.Product, error)
    GetAll() ([]*domain.Product, error)
}

Implementation:
```go
type productRepo struct {
    db *sqlx.DB  // Injected database connection
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
    query := `INSERT INTO products...`
    row := r.db.QueryRow(query, p.Title, p.Description...)
    err := row.Scan(&p.ID)
    return &p, nil
}
```

**Key Concept: Constructor Function**
```go
func NewProductRepo(db *sqlx.DB) ProductRepo {
    return &productRepo{db: db}  // Dependency Injection!
}
```
- Creates a new ProductRepo instance
- Takes database connection as parameter
- Returns it as ProductRepo interface (not concrete struct)
- Why? Loose coupling - code depends on interface, not implementation

---

### 6. **product/** & **user/** - Business Logic Layer
```
Purpose: Contain service classes with business rules
Files: 
- product/service.go
- product/port.go
- user/service.go
- user/port.go
```

**What is a Service?**
- Contains business logic
- Orchestrates between handlers and repository
- Example: "Create product" service checks validation, calls repo, handles errors

**Example Service Structure:**
```go
type service struct {
    prdRepo ProductRepo  // Repository injected
}

func NewService(prdRepo ProductRepo) Service {
    return &service{prdRepo: prdRepo}
}

func (s *service) CreateProduct(p domain.Product) (*domain.Product, error) {
    // Validation logic
    // Call repository
    // Return result
    return s.prdRepo.Create(p)
}
```

**port.go - Interface Definition**
- Defines what methods the service must implement
- Other packages depend on this interface, not the implementation
- Enables testing and loose coupling

---

### 7. **rest/** - HTTP Server & Routing

#### **rest/server.go** - Main HTTP Server
```
Purpose: Start the HTTP server and register routes
Key components:

1. Struct for server:
```go
type Server struct {
    cnf            *config.Config
    productHandler *product.Handler
    userHandler    *user.Handler
}

2. Start() method:
   - Creates a multiplexer (router): http.NewServeMux()
   - Applies middlewares
   - Registers routes
   - Starts listening on HTTP port
```

#### **rest/handlers/** - Request Handlers

```
Purpose: Handle HTTP requests and return responses
Structure:
rest/handlers/
  ├── product/
  │   ├── handler.go        (Main Handler struct)
  │   ├── createProduct.go  (Handle POST /products)
  │   ├── getProducts.go    (Handle GET /products)
  │   ├── get_product.go    (Handle GET /products/{id})
  │   ├── update_product.go (Handle PUT /products/{id})
  │   ├── delete_product.go (Handle DELETE /products/{id})
  │   ├── routes.go         (Register all product routes)
  │   └── port.go           (Handler interface definition)
  │
  └── user/
      ├── handler.go        (Main Handler struct)
      ├── create_user.go    (Handle POST /users)
      ├── login.go          (Handle POST /login)
      ├── port.go
      └── routes.go
```

**Example Handler Flow:**

```
HTTP Request: POST /products
        ↓
rest/handlers/product/routes.go (routes request)
        ↓
rest/handlers/product/createProduct.go (handler function)
        ↓
Extract JSON body → Convert to domain.Product
        ↓
Call service.CreateProduct()
        ↓
Service calls repository.Create()
        ↓
Repository executes SQL query
        ↓
Convert domain.Product back to JSON
        ↓
Return HTTP response (200 OK + JSON body)
```

---

### 8. **rest/middlewares/** - Cross-Cutting Concerns
```
Purpose: Logic that applies to ALL requests
Files:
- middleware.go      (Base middleware structure)
- manager.go         (Middleware pipeline)
- authenticateJWT.go (JWT token validation)
- cors.go            (Cross-origin resource sharing)
- logger.go          (Log requests)
- preflight.go       (Handle OPTIONS requests)
```

**What is Middleware?**
- Functions that process requests before they reach handlers
- Execute in order: Preflight → CORS → Logger → authenticateJWT → Handler

**Example Flow:**
```
Request comes in
        ↓
Preflight middleware (handles OPTIONS)
        ↓
CORS middleware (sets headers)
        ↓
Logger middleware (logs request)
        ↓
JWT middleware (validates authentication token)
        ↓
Handler (if all middlewares passed)
        ↓
Response
```

---

### 9. **util/** - Utility Functions
```
Purpose: Reusable helper functions
Files:
- creat_jwt.go (Generate JWT tokens)
- sendData.go  (Send JSON responses)

These are used throughout the app for common tasks
```

---

### 10. **db_queries/** - SQL Queries Reference
```
Purpose: Example/reference SQL queries (not used by app)
Files:
- 00011-Insert.sql  (INSERT statements)
- 00012-Select.sql  (SELECT statements)
- 00013-Update.sql  (UPDATE statements)
- 00014-Delete.sql  (DELETE statements)

These are learning/reference files - migrations are the actual schema
```

---

### 11. **cmd/** - Command/Startup
```
Purpose: Contains startup/CLI commands
File: serve.go (The Serve() function)

What happens in Serve():
```go
func Serve() {
    // 1. Load configuration
    cnf := config.GetConfig()
    
    // 2. Connect to database
    dbCon, err := infra.NewConnection(cnf.DB)
    
    // 3. Run migrations
    infra.MigrateDB(dbCon, "./migrations")
    
    // 4. Create repositories (injecting database connection)
    userRepo := repo.NewUserRepo(dbCon)
    productRepo := repo.NewProductRepo(dbCon)
    
    // 5. Create services (injecting repositories)
    userSvc := user.NewService(userRepo)
    prdSvc := product.NewService(productRepo)
    
    // 6. Create middlewares
    middlewares := middleware.NewMiddlewares(cnf)
    
    // 7. Create handlers (injecting services)
    productHandler := prdHandler.NewHandler(middlewares, prdSvc)
    userHandler := usrHandler.NewHandler(cnf, userSvc)
    
    // 8. Create server and start
    server := rest.NewServer(cnf, productHandler, userHandler)
    server.Start()
}
```

This is **Dependency Injection** - each layer receives its dependencies as parameters!

---

## 🔄 Complete Request Flow Example

**Scenario: Create a new product**

```
1. Frontend sends:
   POST /products
   Content-Type: application/json
   {
     "title": "Laptop",
     "description": "Gaming laptop",
     "price": 1500,
     "imageUrl": "image.jpg"
   }

2. Server receives request at rest/server.go

3. Middlewares process (logger, CORS, etc.)

4. Router matches POST /products

5. Calls handler: rest/handlers/product/createProduct.go handler function

6. Handler:
   - Parses JSON into domain.Product struct
   - Calls productService.Create(product)

7. Service (product/service.go):
   - Validates data
   - Calls productRepo.Create(product)

8. Repository (repo/product.go):
   - Executes SQL: INSERT INTO products (title, description, price, img_url)
   - Scans returned ID back into product struct
   - Returns *domain.Product

9. Service returns result to handler

10. Handler converts domain.Product to JSON

11. Handler returns HTTP response:
    200 OK
    {
      "id": 1,
      "title": "Laptop",
      "description": "Gaming laptop",
      "price": 1500,
      "imageUrl": "image.jpg"
    }

12. Frontend receives and displays product
```

---

## 🎯 Key Concepts & Terms Explained

### **Interface**
```go
type ProductRepo interface {
    Create(p domain.Product) (*domain.Product, error)
    Get(id int) (*domain.Product, error)
}
```
- Defines a contract: "any type implementing ProductRepo must have these methods"
- Enables loose coupling
- Allows testing with mock implementations
- Why? Code depends on interface, not concrete types

### **Constructor Function**
```go
func NewProductRepo(db *sqlx.DB) ProductRepo {
    return &productRepo{db: db}
}
```
- Creates and returns new instances
- Injects dependencies
- Named convention: `New{TypeName}`
- Why? Centralized object creation

### **Dependency Injection (DI)**
```go
// Instead of creating dependencies inside:
func badWay() {
    db := sqlx.Connect(...)  // Creates its own database
    repo := productRepo{db}
}

// Do this - receive dependencies:
func goodWay(db *sqlx.DB) {
    repo := productRepo{db}  // DB passed in
}
```
- Receives required objects as parameters
- Loose coupling
- Testable
- Flexible

### **Receiver Function (Method)**
```go
func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
    // 'r' is the receiver - the object this method belongs to
    // Can access r.db
}
```
- Method on a type
- `(r *productRepo)` means this method belongs to productRepo
- `*productRepo` means pointer receiver (can modify the object)

### **Struct Tags**
```go
type Product struct {
    ID    int     `json:"id" db:"id"`
    Title string  `json:"title" db:"title"`
}
```
- `json:"id"` - JSON encoder/decoder uses "id" as key
- `db:"id"` - sqlx uses "id" as database column name
- Metadata for packages to interpret

### **Domain-Driven Design (DDD)**
- **domain/** - The core business entities (User, Product)
- **repo/** - Technical details for persistence
- **service/** - Business logic and rules
- Separates business logic from technical concerns

---

## 🚀 How to Run the Project

1. **Install PostgreSQL** and create database "ecommerce"

2. **Create .env file:**
```
VERSION=1.0.0
SERVICE_NAME=ecommerce-api
HTTP_PORT=8080
JWT_SECRET_KEY=your-secret-key

DB_HOST=localhost
DB_PORT=5432
DB_NAME=ecommerce
DB_USER=postgres
DB_PASSWORD=your-password
DB_ENABLE_SSL=false
```

3. **Update main.go to call Serve():**
```go
package main

import "first-program/cmd"

func main() {
    cmd.Serve()
}
```

4. **Run the project:**
```bash
go run main.go
```

5. **Server starts on:** http://localhost:8080

---

## 📊 Package Dependencies Summary

| Package | Purpose | Import Path |
|---------|---------|-------------|
| github.com/lib/pq | PostgreSQL driver | Used for database connection |
| github.com/jmoiron/sqlx | SQL utilities | Used in infra and repo layers |
| github.com/joho/godotenv | Load .env files | Used in config.go |
| golang-jwt/jwt | JWT tokens | Used for authentication |
| net/http | HTTP server | Used in rest/server.go |

---

## 🎓 Learning Path

To understand this project, study in this order:

1. **Basics:** main.go, go.mod, .env
2. **Configuration:** config/config.go
3. **Database:** infra/connection.go, migrations/
4. **Data Models:** domain/user.go, domain/product.go
5. **Data Access:** repo/product.go, repo/user.go
6. **Business Logic:** product/service.go, user/service.go
7. **HTTP Layer:** rest/server.go
8. **Routing:** rest/handlers/product/routes.go
9. **Request Handlers:** rest/handlers/product/createProduct.go
10. **Middleware:** rest/middlewares/

---

## ✨ Architecture Benefits

✅ **Loose Coupling** - Change database? Only modify repo/ layer  
✅ **Testable** - Mock interfaces for unit testing  
✅ **Scalable** - Add features without affecting existing code  
✅ **Maintainable** - Clear separation of concerns  
✅ **Professional** - Follows industry best practices (DDD, DI)  

---

# 🚀 SIMPLE STEP-BY-STEP EXPLANATION: How Backend Works

## When You Run `go run main.go`

### **STEP 1: main.go Executes**
```
$ go run main.go

↓

main.go (Entry Point)
└─→ func main() { 
       cmd.Serve()  // Calls the Serve() function from cmd/serve.go
    }
```

**What happens?**
- Go looks for `main()` function in main.go
- Finds `cmd.Serve()` call
- Jumps to cmd/serve.go to run Serve() function

---

## **STEP 2: cmd/serve.go Runs the Setup**

In cmd/serve.go, this is the actual startup sequence:

```
func Serve() {
    // What happens in which order:
}
```

### **Order 1️⃣ - Load Configuration**

```go
cnf := config.GetConfig()
```

**What it does:**
- Opens .env file
- Reads: DATABASE_HOST, DATABASE_USER, PASSWORD, PORT, etc.
- Creates a `Config` struct with all settings
- Stores it in `cnf` variable

**File:** config/config.go

**Why needed?**
- PostgreSQL needs host, username, password to connect
- Server needs to know which port (8080)
- JWT needs secret key

**Flow:**
```
.env file
   ↓ (read by godotenv)
Environment Variables
   ↓ (read by config.GetConfig())
Config struct (cnf)
   ↓ (used in all other steps)
Database Connection, Server, Handlers
```

---

### **Order 2️⃣ - Connect to Database**

```go
dbCon, err := infra.NewConnection(cnf.DB)
```

**What it does:**
- Takes database credentials from config
- Creates connection string
- Connects to PostgreSQL server
- Returns database connection object

**File:** infra/connection.go

**Connection string looks like:**
```
user=postgres password=secret host=localhost port=5432 dbname=ecommerce sslmode=disable
```

**Flow:**
```
cnf.DB (database credentials)
   ↓ (passed to NewConnection)
infra/connection.go
   ↓ (formats connection string)
PostgreSQL Server
   ↓ (connects)
dbCon (database connection object)
```

**Why needed?**
- Every repo needs to talk to database
- Must establish connection first before doing queries

---

### **Order 3️⃣ - Run Database Migrations**

```go
err = infra.MigrateDB(dbCon, "./migrations")
```

**What it does:**
- Reads migration files from migrations/ folder
- Checks which migrations have been run
- Runs new migrations (creates tables if needed)
- Example: Creates `users` table, `products` table

**File:** infra/migrate.go

**Migration files:**
```
000001-create-user.up.sql      (creates user table)
000001-create-user.down.sql    (deletes user table)
000002-create-products.up.sql  (creates product table)
```

**Flow:**
```
migrations/ folder
   ↓ (read migration files)
infra/migrate.go
   ↓ (check which ran, run new ones)
PostgreSQL Database
   ↓ (tables created)
Database ready to use
```

**Why needed?**
- SQL queries need tables to exist
- Without this, SELECT/INSERT will fail
- Ensures database structure is correct

---

### **Order 4️⃣ - Create Repository Layer**

```go
userRepo := repo.NewUserRepo(dbCon)
productRepo := repo.NewProductRepo(dbCon)
```

**What it does:**
- Takes the database connection
- Wraps it in a repository object
- Repository knows how to do database operations

**Files:** repo/user.go, repo/product.go

**What is a Repository?**
- A wrapper around database
- Contains methods like: Create(), Get(), Update(), Delete()
- Handles SQL queries

**Example from repo/product.go:**
```go
type productRepo struct {
    db *sqlx.DB  // Database connection
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
    return &productRepo{
        db: db,  // Receives database connection as parameter
    }
}

// Repository method to create product
func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
    query := `INSERT INTO products (title, description, price, img_url) 
              VALUES ($1, $2, $3, $4) RETURNING id`
    err := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl).Scan(&p.ID)
    return &p, nil
}
```

**Flow:**
```
dbCon (database connection)
   ↓ (passed to NewUserRepo/NewProductRepo)
repo/user.go & repo/product.go
   ↓ (wraps database connection)
userRepo & productRepo objects
   ↓ (ready to do database operations)
Used by services
```

**Why needed?**
- Separates database logic from business logic
- If switching databases (PostgreSQL → MongoDB), only change repo
- Services don't need to know about SQL details

---

### **Order 5️⃣ - Create Service Layer (Business Logic)**

```go
userSvc := user.NewService(userRepo)
prdSvc := product.NewService(productRepo)
```

**What it does:**
- Takes the repository
- Creates a service that wraps the repository
- Service handles business logic

**Files:** product/service.go, user/service.go

**What is a Service?**
- Contains business rules
- Calls repository methods
- Example: "Create product" validates data, then calls repo.Create()

**Example from product/service.go:**
```go
type service struct {
    prdRepo ProductRepo  // Repository injected
}

func NewService(prdRepo ProductRepo) Service {
    return &service{
        prdRepo: prdRepo,
    }
}

// Service method with business logic
func (s *service) CreateProduct(p domain.Product) (*domain.Product, error) {
    // Validation logic here
    if p.Title == "" {
        return nil, errors.New("Title required")
    }
    if p.Price <= 0 {
        return nil, errors.New("Price must be positive")
    }
    
    // If validation passes, call repository
    return s.prdRepo.Create(p)
}
```

**Flow:**
```
userRepo & productRepo
   ↓ (passed to NewService)
product/service.go & user/service.go
   ↓ (wraps repositories)
userSvc & prdSvc objects
   ↓ (contains business logic)
Used by handlers
```

**Why needed?**
- Business logic separated from API endpoints
- Validation, calculations happen here
- Reusable across different API versions

---

### **Order 6️⃣ - Create Middleware**

```go
middlewares := middleware.NewMiddlewares(cnf)
```

**What it does:**
- Creates middleware objects
- Middlewares handle: logging, CORS, JWT authentication, etc.

**Files:** rest/middlewares/

**What are Middlewares?**
- Functions that run BEFORE handlers
- Process every request
- Example: Check JWT token before allowing access

**Middleware order:**
```
Request comes → Preflight → CORS → Logger → JWT Check → Handler → Response
```

**Flow:**
```
cnf (configuration with JWT secret)
   ↓ (passed to NewMiddlewares)
rest/middlewares/
   ↓ (creates middleware objects)
middlewares object
   ↓ (used by handlers)
Every request passes through
```

---

### **Order 7️⃣ - Create Handlers**

```go
productHandler := prdHandler.NewHandler(middlewares, prdSvc)
userHandler := usrHandler.NewHandler(cnf, userSvc)
```

**What it does:**
- Takes service and middleware
- Creates handler objects
- Handlers handle HTTP requests

**Files:** rest/handlers/product/, rest/handlers/user/

**What is a Handler?**
- Receives HTTP request
- Calls service method
- Returns JSON response

**Example from rest/handlers/product/createProduct.go:**
```go
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
    // 1. Parse JSON from request body
    var product domain.Product
    json.NewDecoder(r.Body).Decode(&product)
    
    // 2. Call service
    result, err := h.svc.CreateProduct(product)
    
    // 3. Return JSON response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}
```

**Flow:**
```
middlewares & prdSvc/userSvc
   ↓ (passed to NewHandler)
rest/handlers/product/ & rest/handlers/user/
   ↓ (creates handler objects)
productHandler & userHandler
   ↓ (ready to handle requests)
Used by server
```

**Why needed?**
- Receives HTTP requests from clients
- Converts HTTP data to Go structs
- Calls business logic
- Returns JSON responses

---

### **Order 8️⃣ - Create Server & Start**

```go
server := rest.NewServer(cnf, productHandler, userHandler)
server.Start()
```

**What it does:**
- Creates HTTP server
- Registers all routes
- Starts listening for requests

**File:** rest/server.go

**What happens in server.Start():**
```go
func (server *Server) Start() {
    // 1. Create router
    mux := http.NewServeMux()
    
    // 2. Register routes
    server.productHandler.RegisterRoutes(mux, manager)
    server.userHandler.RegisterRoutes(mux, manager)
    
    // 3. Start server on port 8080
    addr := ":8080"
    http.ListenAndServe(addr, mux)
}
```

**Routes registered:**
```
POST   /products          → Create product
GET    /products          → Get all products
GET    /products/{id}     → Get one product
PUT    /products/{id}     → Update product
DELETE /products/{id}     → Delete product

POST   /signup            → Create user
POST   /login             → User login
```

**Flow:**
```
productHandler & userHandler
   ↓ (passed to NewServer)
rest/server.go
   ↓ (creates HTTP server)
server.Start()
   ↓ (registers routes, starts listening)
Server running on :8080
   ↓ (waiting for requests)
Ready to handle API calls
```

---

## **COMPLETE FLOW CHART: From main.go to API Response**

```
START: go run main.go
        ↓
[STEP 1] main.go
        ↓
        cmd.Serve()
        ↓
[STEP 2] config.GetConfig()
        │
        ├─ Read .env file
        ├─ Load DATABASE_URL, PORT, JWT_SECRET
        └─ Return Config struct
        ↓
[STEP 3] infra.NewConnection(cnf.DB)
        │
        ├─ Format connection string
        ├─ Connect to PostgreSQL
        └─ Return database connection (dbCon)
        ↓
[STEP 4] infra.MigrateDB(dbCon, "./migrations")
        │
        ├─ Read migration files
        ├─ Create tables: users, products
        └─ Database ready
        ↓
[STEP 5] repo.NewUserRepo(dbCon)
        │    repo.NewProductRepo(dbCon)
        ├─ userRepo: knows how to query users table
        └─ productRepo: knows how to query products table
        ↓
[STEP 6] user.NewService(userRepo)
        │    product.NewService(productRepo)
        ├─ userSvc: business logic for users
        └─ prdSvc: business logic for products
        ↓
[STEP 7] middleware.NewMiddlewares(cnf)
        │
        ├─ Logger middleware: log requests
        ├─ CORS middleware: allow cross-origin
        ├─ Preflight middleware: handle OPTIONS
        └─ JWT middleware: check authentication
        ↓
[STEP 8] prdHandler.NewHandler(middlewares, prdSvc)
        │    usrHandler.NewHandler(cnf, userSvc)
        ├─ productHandler: handles product API calls
        └─ userHandler: handles user API calls
        ↓
[STEP 9] rest.NewServer(cnf, productHandler, userHandler)
        │
        ├─ Create HTTP server
        ├─ Register all routes
        └─ server.Start()
        ↓
[STEP 10] Server listening on localhost:8080
        ↓
        ✅ READY TO HANDLE API REQUESTS
```

---

## **When Client Sends API Request**

### **Example: POST /products (Create Product)**

```
Frontend sends:
┌─────────────────────────────────┐
│ POST /products                  │
│ Content-Type: application/json  │
│                                 │
│ {                               │
│   "title": "Laptop",            │
│   "description": "Gaming laptop"│
│   "price": 1500,                │
│   "imageUrl": "image.jpg"       │
│ }                               │
└─────────────────────────────────┘
        ↓
[STEP 1] Server receives request at localhost:8080
        ↓
[STEP 2] Middlewares process:
        │
        ├─ Preflight: Check if OPTIONS request? (No, continue)
        ├─ CORS: Add CORS headers
        ├─ Logger: Log this request
        └─ JWT: Check authentication token (if needed)
        ↓
[STEP 3] Router matches: POST /products
        │
        └─ Calls: productHandler.CreateProduct()
        ↓
[STEP 4] Handler (rest/handlers/product/createProduct.go):
        │
        ├─ Parse JSON from request body
        ├─ Extract: title, description, price, imageUrl
        ├─ Create domain.Product struct
        └─ Call: prdSvc.CreateProduct(product)
        ↓
[STEP 5] Service (product/service.go):
        │
        ├─ Validate: Title not empty? Price > 0?
        ├─ if validation passes:
        └─ Call: prdRepo.Create(product)
        ↓
[STEP 6] Repository (repo/product.go):
        │
        ├─ Create SQL query: INSERT INTO products...
        ├─ Execute on database: prdRepo.db.QueryRow(query)
        ├─ Get returned product ID from database
        ├─ Put ID into product struct
        └─ Return: (*domain.Product, error)
        ↓
[STEP 7] Database (PostgreSQL):
        │
        ├─ INSERT product into products table
        ├─ Auto-generate ID
        └─ Return: new product with ID
        ↓
[STEP 8] Back to Repository:
        │
        └─ Convert database result to domain.Product
        ↓
[STEP 9] Back to Service:
        │
        └─ Return result to handler
        ↓
[STEP 10] Handler converts to JSON:
        │
        └─ Struct → JSON format
        ↓
[STEP 11] Send Response:
        │
        ├─ HTTP Status: 200 OK
        ├─ Content-Type: application/json
        └─ Body:
            {
              "id": 1,
              "title": "Laptop",
              "description": "Gaming laptop",
              "price": 1500,
              "imageUrl": "image.jpg"
            }
        ↓
[STEP 12] Frontend receives response
        │
        └─ Displays product in browser
```

---

## **WHICH WORKS AFTER WHICH ONE - Summary**

| Order | Component | Input | Output | File |
|-------|-----------|-------|--------|------|
| 1 | main.go | - | Calls Serve() | main.go |
| 2 | config | .env file | Config struct | config/config.go |
| 3 | database connection | Config struct | dbCon object | infra/connection.go |
| 4 | migrations | dbCon + migrations/ | Tables created | infra/migrate.go |
| 5 | repositories | dbCon | productRepo, userRepo | repo/ |
| 6 | services | repos | productSvc, userSvc | product/, user/ |
| 7 | middlewares | Config | middleware objects | rest/middlewares/ |
| 8 | handlers | services + middlewares | productHandler, userHandler | rest/handlers/ |
| 9 | server | handlers + config | HTTP server | rest/server.go |
| 10 | server.Start() | - | Listening on port 8080 | rest/server.go |

---

## **SIMPLE RULE TO REMEMBER**

Each layer depends on the previous layer:

```
Config (settings)
   ↓
Database Connection (access to database)
   ↓
Migrations (create tables)
   ↓
Repositories (query database)
   ↓
Services (business logic)
   ↓
Middlewares (process requests)
   ↓
Handlers (receive requests)
   ↓
Server (start HTTP server)
   ↓
Accept API requests from clients
```

**If any layer is missing or wrong → entire chain breaks!**

---

## **Example: Adding New API Endpoint**

Want to add new endpoint: `POST /products/{id}/reviews` (add review to product)?

**What you need to implement:**

1. **Step 1:** Add Review struct in `domain/review.go`
   ```go
   type Review struct {
       ID        int
       ProductID int
       Rating    int
       Comment   string
   }
   ```

2. **Step 2:** Add methods in `repo/review.go`
   ```go
   func (r *reviewRepo) Create(rev domain.Review) error {
       // SQL: INSERT INTO reviews...
   }
   ```

3. **Step 3:** Add service in `review/service.go`
   ```go
   func (s *service) CreateReview(rev domain.Review) error {
       // Validate: Rating 1-5?
       // Call repo.Create()
   }
   ```

4. **Step 4:** Add handler in `rest/handlers/review/createReview.go`
   ```go
   func (h *Handler) CreateReview(w http.ResponseWriter, r *http.Request) {
       // Parse JSON
       // Call service
       // Return JSON
   }
   ```

5. **Step 5:** Add route in `rest/handlers/review/routes.go`
   ```go
   mux.HandleFunc("POST /products/{id}/reviews", h.CreateReview)
   ```

6. **Step 6:** Register in `cmd/serve.go`
   ```go
   reviewHandler := revHandler.NewHandler(middlewares, revSvc)
   ```

**Pattern is always same!** Domain → Repo → Service → Handler → Route

---

FIRST-PROGRAM/
├── cmd/
│   └── serve.go
├── config/
│   └── config.go
├── db_queries/
│   ├── 001-Create-UserTable.sql
│   ├── 002-Create-ProductTable.sql
│   ├── 002-product-Table.sql
│   ├── 00011-Insert.sql
│   ├── 00012-Select.sql
│   ├── 00013-Update.sql
│   └── 00014-Delete.sql
├── domain/
│   ├── product.go
│   └── user.go
├── infra/
│   ├── connection.go
│   └── migrate.go
├── migrations/
│   ├── 000001-create-user.down.sql
│   ├── 000001-create-user.up.sql
│   ├── 000002-create-products.down.sql
│   └── 000002-create-products.up.sql
├── product/
│   ├── port.go
│   └── service.go
├── repo/
│   ├── product.go
│   └── user.go
├── rest/
│   └── handlers/
│       ├── product/
│       │   ├── createProduct.go
│       │   ├── delete_product.go
│       │   ├── get_product.go
│       │   ├── getProducts.go
│       │   ├── handler.go
│       │   ├── port.go
│       │   ├── routes.go
│       │   └── update_product.go
│       └── user/
│           ├── create_user.go
│           ├── handler.go
│           ├── login.go
│           ├── port.go
│           └── routes.go
├── middlewares/
│   ├── authenticate_JWT.go
│   ├── cors.go
│   ├── logger.go
│   ├── manager.go
│   ├── middleware.go
│   ├── preflight.go
│   └── server.go
├── user/
│   ├── port.go
│   └── service.go
├── util/
│   ├── creat_jwt.go
│   └── sendData.go
├── .env
├── go.mod
├── go.sum
├── main.go
└── readme.md

**Happy Learning! 🎉**