// ========================================
// COMPLETE PROJECT CONSOLIDATION
// All code organized by file
// ========================================

package main

// This file is a reference document containing all source code
// from the project organized by file location and section

// =====================================
// 1. MAIN.GO
// =====================================
// package main
//
// import (
// 	"first-program/cmd"
// )
//
// func main() {
// 	cmd.Serve()
// }

// =====================================
// 2. CMD/SERVE.GO
// =====================================
// package cmd
//
// import (
// 	"first-program/config"
// 	"first-program/infra"
// 	"first-program/product"
// 	"first-program/repo"
// 	"first-program/rest"
// 	prdHandler "first-program/rest/handlers/product"
// 	usrHandler "first-program/rest/handlers/user"
// 	middleware "first-program/rest/middlewares"
// 	"first-program/user"
// 	"fmt"
// 	"os"
// 	_ "github.com/lib/pq"
// )
//
// func Serve() {
// 	cnf := config.GetConfig()
// 	dbCon, err := infra.NewConnection(cnf.DB)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	err = infra.MigrateDB(dbCon, "./migrations")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	userRepo := repo.NewUserRepo(dbCon)
// 	productRepo := repo.NewProductRepo(dbCon)
// 	userSvc := user.NewService(userRepo)
// 	prdSvc := product.NewService(productRepo)
// 	middlewares := middleware.NewMiddlewares(cnf)
// 	productHandler := prdHandler.NewHandler(middlewares, prdSvc)
// 	userHandler := usrHandler.NewHandler(cnf, userSvc)
// 	server := rest.NewServer(cnf, productHandler, userHandler)
// 	server.Start()
// }

// =====================================
// 3. CONFIG/CONFIG.GO
// =====================================
// package config
//
// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"github.com/joho/godotenv"
// )
//
// var configurations *Config
//
// type DBConfig struct {
// 	Host          string
// 	Port          int
// 	Name          string
// 	User          string
// 	Password      string
// 	EnableSSLMODE bool
// }
//
// type Config struct {
// 	Version      string
// 	ServiceName  string
// 	HttpPort     int
// 	JwtSecretKey string
// 	DB           *DBConfig
// }
//
// func loadConfig() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		fmt.Println("Failed to load the env variables: ", err)
// 		os.Exit(1)
// 	}
//
// 	version := os.Getenv("VERSION")
// 	if version == "" {
// 		fmt.Println("Version is required")
// 		os.Exit(1)
// 	}
//
// 	serviceName := os.Getenv("SERVICE_NAME")
// 	if serviceName == "" {
// 		fmt.Println("Service Name is required")
// 		os.Exit(1)
// 	}
//
// 	httpPort := os.Getenv("HTTP_PORT")
// 	if httpPort == "" {
// 		fmt.Println("Http Port is required")
// 		os.Exit(1)
// 	}
//
// 	port, err := strconv.ParseInt(httpPort, 10, 64)
// 	if err != nil {
// 		fmt.Println("Port must be number")
// 		os.Exit(1)
// 	}
//
// 	jwtSecretKey := os.Getenv("JWT_SECRECT_KEY")
// 	if jwtSecretKey == "" {
// 		fmt.Println("Jwt secret key is required")
// 		os.Exit(1)
// 	}
//
// 	db_host := os.Getenv("DB_HOST")
// 	if db_host == "" {
// 		fmt.Println("DB Host is required")
// 		os.Exit(1)
// 	}
//
// 	db_port := os.Getenv("DB_PORT")
// 	if db_port == "" {
// 		fmt.Println("DB Port is required")
// 		os.Exit(1)
// 	}
//
// 	db_prt, err := strconv.ParseInt(db_port, 10, 64)
// 	if err != nil {
// 		fmt.Println("Port must be number")
// 		os.Exit(1)
// 	}
//
// 	db_name := os.Getenv("DB_NAME")
// 	if db_name == "" {
// 		fmt.Println("DB Name is required")
// 		os.Exit(1)
// 	}
//
// 	db_user := os.Getenv("DB_USER")
// 	if db_user == "" {
// 		fmt.Println("DB Name is required")
// 		os.Exit(1)
// 	}
//
// 	db_pass := os.Getenv("DB_PASSWORD")
// 	if db_pass == "" {
// 		fmt.Println("DB Password is required")
// 		os.Exit(1)
// 	}
//
// 	db_enableSslMode := os.Getenv("DB_ENABLE_SSL_MODE")
// 	db_enableSSlMode, err := strconv.ParseBool(db_enableSslMode)
// 	if err != nil {
// 		fmt.Println("DB enable ssl mode is required")
// 		os.Exit(1)
// 	}
//
// 	dbConfig := &DBConfig{
// 		Host:          db_host,
// 		Port:          int(db_prt),
// 		Name:          db_name,
// 		User:          db_user,
// 		Password:      db_pass,
// 		EnableSSLMODE: db_enableSSlMode,
// 	}
//
// 	configurations = &Config{
// 		Version:      version,
// 		ServiceName:  serviceName,
// 		HttpPort:     int(port),
// 		JwtSecretKey: jwtSecretKey,
// 		DB:           dbConfig,
// 	}
// }
//
// func GetConfig() *Config {
// 	if configurations == nil {
// 		loadConfig()
// 	}
// 	return configurations
// }

// =====================================
// 4. DOMAIN/PRODUCT.GO
// =====================================
// package domain
//
// type Product struct {
// 	ID          int     `json:"id" db:"id"`
// 	Title       string  `json:"title" db:"title"`
// 	Description string  `json:"description" db:"description"`
// 	Price       float64 `json:"price" db:"price"`
// 	ImgUrl      string  `json:"imageUrl" db:"img_url"`
// }

// =====================================
// 5. DOMAIN/USER.GO
// =====================================
// package domain
//
// type User struct {
// 	ID          int    `json:"id" db:"id"`
// 	FirstName   string `json:"first_name" db:"first_name"`
// 	LastName    string `json:"last_name" db:"last_name"`
// 	Email       string `json:"email" db:"email"`
// 	Password    string `json:"password" db:"password"`
// 	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
// }

// =====================================
// 6. INFRA/CONNECTION.GO
// =====================================
// package infra
//
// import (
// 	"first-program/config"
// 	"fmt"
// 	"github.com/jmoiron/sqlx"
// 	_ "github.com/lib/pq"
// )
//
// func GetConnectionString(cnf *config.DBConfig) string {
// 	connString := fmt.Sprintf(
// 		"user=%s password=%s host=%s port=%d dbname=%s",
// 		cnf.User,
// 		cnf.Password,
// 		cnf.Host,
// 		cnf.Port,
// 		cnf.Name,
// 	)
// 	if !cnf.EnableSSLMODE {
// 		connString += " sslmode=disable"
// 	}
// 	return connString
// }
//
// func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
// 	dbSource := GetConnectionString(cnf)
// 	dbCon, err := sqlx.Connect("postgres", dbSource)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}
// 	return dbCon, nil
// }

// =====================================
// 7. INFRA/MIGRATE.GO
// =====================================
// package infra
//
// import (
// 	"fmt"
// 	"github.com/jmoiron/sqlx"
// 	_ "github.com/lib/pq"
// 	migrate "github.com/rubenv/sql-migrate"
// )
//
// func MigrateDB(db *sqlx.DB, dir string) error {
// 	migrations := &migrate.FileMigrationSource{
// 		Dir: dir,
// 	}
//
// 	_, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("Successfully migrated database")
// 	return nil
// }

// =====================================
// 8. PRODUCT/PORT.GO
// =====================================
// package product
//
// import (
// 	"first-program/domain"
// 	prdHandler "first-program/rest/handlers/product"
// )
//
// type Service interface {
// 	prdHandler.Service
// }
//
// type ProductRepo interface {
// 	Create(p domain.Product) (*domain.Product, error)
// 	Get(productID int) (*domain.Product, error)
// 	List() ([]*domain.Product, error)
// 	Update(product domain.Product) (*domain.Product, error)
// 	Delete(productID int) error
// }

// =====================================
// 9. PRODUCT/SERVICE.GO
// =====================================
// package product
//
// import "first-program/domain"
//
// type service struct {
// 	prdRepo ProductRepo
// }
//
// func NewService(prdRepo ProductRepo) Service {
// 	return &service{
// 		prdRepo: prdRepo,
// 	}
// }
//
// func (s *service) Create(p domain.Product) (*domain.Product, error) {
// 	return s.prdRepo.Create(p)
// }
//
// func (s *service) Get(productID int) (*domain.Product, error) {
// 	return s.prdRepo.Get(productID)
// }
//
// func (s *service) List() ([]*domain.Product, error) {
// 	return s.prdRepo.List()
// }
//
// func (s *service) Update(p domain.Product) (*domain.Product, error) {
// 	return s.prdRepo.Update(p)
// }
//
// func (s *service) Delete(productID int) error {
// 	return s.prdRepo.Delete(productID)
// }

// =====================================
// 10. REPO/USER.GO
// =====================================
// package repo
//
// import (
// 	"database/sql"
// 	"first-program/domain"
// 	"first-program/user"
// 	"github.com/jmoiron/sqlx"
// )
//
// type UserRepo interface {
// 	user.UserRepo
// }
//
// type userRepo struct {
// 	db *sqlx.DB
// }
//
// func NewUserRepo(db *sqlx.DB) UserRepo {
// 	return &userRepo{
// 		db: db,
// 	}
// }
//
// func (r *userRepo) Create(user domain.User) (*domain.User, error) {
// 	query := `
// 		INSERT INTO users (
// 			first_name, last_name, email, password, is_shop_owner
// 		) VALUES (
// 			:first_name, :last_name, :email, :password, :is_shop_owner
// 		) RETURNING id
// 	`
// 	var userID int
// 	rows, err := r.db.NamedQuery(query, user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if rows.Next() {
// 		rows.Scan(&userID)
// 	}
// 	user.ID = userID
// 	return &user, nil
// }
//
// func (r *userRepo) Find(email, pass string) (*domain.User, error) {
// 	var user domain.User
// 	query := `SELECT id, first_name, last_name, email, password, is_shop_owner
// 		FROM users WHERE email = $1 and password = $2 LIMIT 1`
// 	err := r.db.Get(&user, query, email, pass)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &user, nil
// }

// =====================================
// 11. REPO/PRODUCT.GO
// =====================================
// package repo
//
// import (
// 	"database/sql"
// 	"first-program/domain"
// 	"first-program/product"
// 	"github.com/jmoiron/sqlx"
// )
//
// type ProductRepo interface {
// 	product.ProductRepo
// }
//
// type productRepo struct {
// 	db *sqlx.DB
// }
//
// func NewProductRepo(db *sqlx.DB) ProductRepo {
// 	repo := &productRepo{
// 		db: db,
// 	}
// 	return repo
// }
//
// func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
// 	query := `INSERT INTO products (title, description, price, img_url)
// 		VALUES ($1, $2, $3, $4) RETURNING id`
// 	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl)
// 	err := row.Scan(&p.ID)
// 	if err != nil {
// 		return nil, nil
// 	}
// 	return &p, nil
// }
//
// func (r *productRepo) Get(id int) (*domain.Product, error) {
// 	var prd domain.Product
// 	query := `SELECT id, title, description, price, img_url FROM products WHERE id = $1`
// 	err := r.db.Get(&prd, query, id)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &prd, nil
// }
//
// func (r *productRepo) List() ([]*domain.Product, error) {
// 	var ProductList []*domain.Product
// 	query := `SELECT id, title, description, price, img_url FROM products`
// 	err := r.db.Select(&ProductList, query)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return ProductList, nil
// }
//
// func (r *productRepo) Delete(productID int) error {
// 	query := `DELETE FROM products WHERE id = $1`
// 	_, err := r.db.Exec(query, productID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil
// 		}
// 		return err
// 	}
// 	return nil
// }
//
// func (r *productRepo) Update(p domain.Product) (*domain.Product, error) {
// 	query := `UPDATE products SET title=$1, description=$2, price=$3, img_url=$4
// 		WHERE id = $5`
// 	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl, p.ID)
// 	err := row.Err()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &p, nil
// }

// =====================================
// 12. REST/SERVER.GO
// =====================================
// package rest
//
// import (
// 	"first-program/config"
// 	"first-program/rest/handlers/product"
// 	"first-program/rest/handlers/user"
// 	middleware "first-program/rest/middlewares"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"strconv"
// )
//
// type Server struct {
// 	cnf            *config.Config
// 	productHandler *product.Handler
// 	userHandler    *user.Handler
// }
//
// func NewServer(cnf *config.Config, productHandler *product.Handler, userHandler *user.Handler) *Server {
// 	return &Server{
// 		cnf:            cnf,
// 		productHandler: productHandler,
// 		userHandler:    userHandler,
// 	}
// }
//
// func (server *Server) Start() {
// 	manager := middleware.NewManager()
// 	manager.Use(middleware.Preflight, middleware.Cors, middleware.Logger)
// 	mux := http.NewServeMux()
// 	WrappedMux := manager.WrapMux(mux)
// 	server.productHandler.RegisterRoutes(mux, manager)
// 	server.userHandler.RegisterRoutes(mux, manager)
// 	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
// 	fmt.Println("Server running on port", addr)
// 	err := http.ListenAndServe(addr, WrappedMux)
// 	if err != nil {
// 		fmt.Println("Error starting the server", err)
// 		os.Exit(1)
// 	}
// }

// =====================================
// 13. REST/HANDLERS/PRODUCT/HANDLER.GO
// =====================================
// package product
//
// import middleware "first-program/rest/middlewares"
//
// type Handler struct {
// 	middlewares *middleware.Middlewares
// 	svc         Service
// }
//
// func NewHandler(middlewares *middleware.Middlewares, svc Service) *Handler {
// 	return &Handler{
// 		middlewares: middlewares,
// 		svc:         svc,
// 	}
// }

// =====================================
// 14. REST/HANDLERS/PRODUCT/PORT.GO
// =====================================
// package product
//
// import "first-program/domain"
//
// type Service interface {
// 	Create(domain.Product) (*domain.Product, error)
// 	Get(productID int) (*domain.Product, error)
// 	List() ([]*domain.Product, error)
// 	Update(p domain.Product) (*domain.Product, error)
// 	Delete(productID int) error
// }

// =====================================
// 15. REST/HANDLERS/PRODUCT/ROUTES.GO
// =====================================
// package product
//
// import (
// 	middleware "first-program/rest/middlewares"
// 	"net/http"
// )
//
// func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
// 	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProducts)))
// 	mux.Handle("POST /products", manager.With(
// 		http.HandlerFunc(h.CreateProduct), h.middlewares.AuthenticateJWT))
// 	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(h.GetProduct)))
// 	mux.Handle("PUT /products/{id}", manager.With(
// 		http.HandlerFunc(h.UpdateProducts), h.middlewares.AuthenticateJWT))
// 	mux.Handle("DELETE /products/{id}", manager.With(
// 		http.HandlerFunc(h.DeleteProducts), h.middlewares.AuthenticateJWT))
// }

// =====================================
// 16. REST/HANDLERS/PRODUCT/CREATEPRODUCT.GO
// =====================================
// package product
//
// import (
// 	"encoding/json"
// 	"first-program/domain"
// 	"first-program/util"
// 	"fmt"
// 	"net/http"
// )
//
// type ReqCreateProduct struct {
// 	Title       string  `json:"title"`
// 	Description string  `json:"description"`
// 	Price       float64 `json:"price"`
// 	ImgUrl      string  `json:"imageUrl"`
// }
//
// func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
// 	var req ReqCreateProduct
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&req)
// 	if err != nil {
// 		fmt.Println(err)
// 		util.SendError(w, http.StatusBadRequest, "Invalid req body")
// 		return
// 	}
// 	createdProduct, err := h.svc.Create(domain.Product{
// 		Title: req.Title, Description: req.Description, Price: req.Price, ImgUrl: r.RequestURI,
// 	})
// 	if err != nil {
// 		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
// 		return
// 	}
// 	util.SendData(w, http.StatusCreated, createdProduct)
// }

// =====================================
// 17. REST/HANDLERS/PRODUCT/GETPRODUCTS.GO
// =====================================
// package product
//
// import (
// 	"first-program/util"
// 	"net/http"
// )
//
// func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
// 	productList, err := h.svc.List()
// 	if err != nil {
// 		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
// 	}
// 	util.SendData(w, http.StatusOK, productList)
// }

// =====================================
// 18. REST/HANDLERS/PRODUCT/GET_PRODUCT.GO
// =====================================
// package product
//
// import (
// 	"first-program/util"
// 	"net/http"
// 	"strconv"
// )
//
// func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
// 	productID := r.PathValue("id")
// 	pID, err := strconv.Atoi(productID)
// 	if err != nil {
// 		util.SendError(w, http.StatusBadRequest, "Invalid req body")
// 		return
// 	}
// 	product, err := h.svc.Get(pID)
// 	if err != nil {
// 		util.SendError(w, http.StatusInternalServerError, "Invalid req body")
// 		return
// 	}
// 	if product == nil {
// 		util.SendError(w, http.StatusNotFound, "Invalid req body")
// 		return
// 	}
// 	util.SendData(w, http.StatusOK, product)
// }

// =====================================
// 19. REST/HANDLERS/PRODUCT/UPDATE_PRODUCT.GO
// =====================================
// package product
//
// import (
// 	"encoding/json"
// 	"first-program/domain"
// 	"first-program/util"
// 	"fmt"
// 	"net/http"
// 	"strconv"
// )
//
// type ReqUpdateProducts struct {
// 	Title       string  `json:"title"`
// 	Description string  `json:"description"`
// 	Price       float64 `json:"price"`
// 	ImgUrl      string  `json:"imageUrl"`
// }
//
// func (h *Handler) UpdateProducts(w http.ResponseWriter, r *http.Request) {
// 	productID := r.PathValue("id")
// 	pID, err := strconv.Atoi(productID)
// 	if err != nil {
// 		util.SendError(w, http.StatusBadRequest, "Invalid product ID")
// 		return
// 	}
// 	var req ReqUpdateProducts
// 	decoder := json.NewDecoder(r.Body)
// 	err = decoder.Decode(&req)
// 	if err != nil {
// 		fmt.Println(err)
// 		http.Error(w, "Please give me valid json", 400)
// 		util.SendError(w, http.StatusBadRequest, "Invalid req body")
// 		return
// 	}
// 	_, err = h.svc.Update(domain.Product{
// 		ID: pID, Title: req.Title, Description: req.Description,
// 		Price: req.Price, ImgUrl: r.RequestURI,
// 	})
// 	if err != nil {
// 		util.SendError(w, http.StatusInternalServerError, "Internal Server error")
// 		return
// 	}
// 	util.SendData(w, http.StatusOK, "Successfully updated product")
// }

// =====================================
// 20. REST/HANDLERS/PRODUCT/DELETE_PRODUCT.GO
// =====================================
// package product
//
// import (
// 	"first-program/util"
// 	"fmt"
// 	"net/http"
// 	"strconv"
// )
//
// func (h *Handler) DeleteProducts(w http.ResponseWriter, r *http.Request) {
// 	productID := r.PathValue("id")
// 	pID, err := strconv.Atoi(productID)
// 	if err != nil {
// 		fmt.Println(err)
// 		util.SendError(w, http.StatusBadRequest, "Invalid product ID")
// 		return
// 	}
// 	err = h.svc.Delete(pID)
// 	if err != nil {
// 		fmt.Println(err)
// 		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
// 		return
// 	}
// 	util.SendData(w, http.StatusOK, "Successfully Deleted product")
// }

// =====================================
// 21. REST/HANDLERS/USER/HANDLER.GO
// =====================================
// package user
//
// import "first-program/config"
//
// type Handler struct {
// 	cnf *config.Config
// 	svc Service
// }
//
// func NewHandler(cnf *config.Config, svc Service) *Handler {
// 	return &Handler{
// 		cnf: cnf,
// 		svc: svc,
// 	}
// }

// =====================================
// 22. REST/HANDLERS/USER/PORT.GO
// =====================================
// package user
//
// import "first-program/domain"
//
// type Service interface {
// 	Create(user domain.User) (*domain.User, error)
// 	Find(email string, pass string) (*domain.User, error)
// }

// =====================================
// 23. REST/HANDLERS/USER/ROUTES.GO
// =====================================
// package user
//
// import (
// 	middleware "first-program/rest/middlewares"
// 	"net/http"
// )
//
// func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
// 	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.CreateUser)))
// 	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(h.Login)))
// }

// =====================================
// 24. REST/HANDLERS/USER/CREATE_USER.GO
// =====================================
// package user
//
// import (
// 	"encoding/json"
// 	"first-program/domain"
// 	"first-program/util"
// 	"fmt"
// 	"net/http"
// )
//
// type ReqCreateUser struct {
// 	FirstName   string `json:"first_name"`
// 	LastName    string `json:"last_name"`
// 	Email       string `json:"email"`
// 	Password    string `json:"password"`
// 	IsShopOwner bool   `json:"is_shop_owner"`
// }
//
// func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
// 	var req ReqCreateUser
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&req)
// 	if err != nil {
// 		fmt.Println(err)
// 		util.SendError(w, http.StatusBadRequest, "Invalid req body")
// 		return
// 	}
// 	usr, err := h.svc.Create(domain.User{
// 		FirstName: req.FirstName, LastName: req.LastName, Email: req.Email,
// 		Password: req.Password, IsShopOwner: req.IsShopOwner,
// 	})
// 	if err != nil {
// 		util.SendError(w, http.StatusInternalServerError, "Invalid server Error")
// 		return
// 	}
// 	util.SendData(w, http.StatusCreated, usr)
// }

// =====================================
// 25. REST/HANDLERS/USER/LOGIN.GO
// =====================================
// package user
//
// import (
// 	"encoding/json"
// 	"first-program/util"
// 	"fmt"
// 	"net/http"
// )
//
// type Reqlogin struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }
//
// func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
// 	var req Reqlogin
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&req)
// 	if err != nil {
// 		fmt.Println(err)
// 		util.SendError(w, http.StatusBadRequest, "Invalid red Body")
// 		return
// 	}
// 	usr, err := h.svc.Find(req.Email, req.Password)
// 	if err != nil {
// 		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
// 		return
// 	}
// 	if usr == nil {
// 		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
// 		return
// 	}
// 	accressToken, err := util.CreateJWT(h.cnf.JwtSecretKey, util.Payload{
// 		Sub: usr.ID, FirstName: usr.FirstName, LastName: usr.LastName, Email: usr.Email,
// 	})
// 	if err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		return
// 	}
// 	util.SendData(w, http.StatusCreated, accressToken)
// }

// =====================================
// 26. REST/MIDDLEWARES/MIDDLEWARE.GO
// =====================================
// package middleware
//
// import "first-program/config"
//
// type Middlewares struct {
// 	cnf *config.Config
// }
//
// func NewMiddlewares(cnf *config.Config) *Middlewares {
// 	return &Middlewares{
// 		cnf: cnf,
// 	}
// }

// =====================================
// 27. REST/MIDDLEWARES/MANAGER.GO
// =====================================
// package middleware
//
// import "net/http"
//
// type Middleware func(http.Handler) http.Handler
//
// type Manager struct {
// 	globalMiddlewares []Middleware
// }
//
// func NewManager() *Manager {
// 	return &Manager{
// 		globalMiddlewares: make([]Middleware, 0),
// 	}
// }
//
// func (mngr *Manager) Use(middlewares ...Middleware) {
// 	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
// }
//
// func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
// 	h := next
// 	for _, middleware := range middlewares {
// 		h = middleware(h)
// 	}
// 	return h
// }
//
// func (mngr *Manager) WrapMux(next http.Handler) http.Handler {
// 	h := next
// 	for _, gblMiddleware := range mngr.globalMiddlewares {
// 		h = gblMiddleware(h)
// 	}
// 	return h
// }

// =====================================
// 28. REST/MIDDLEWARES/AUTHENTICATEJWT.GO
// =====================================
// package middleware
//
// import (
// 	"crypto/hmac"
// 	"crypto/sha256"
// 	"encoding/base64"
// 	"net/http"
// 	"strings"
// )
//
// func base64URLEncode(data []byte) string {
// 	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
// }
//
// func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		header := r.Header.Get("Authorization")
// 		if header == "" {
// 			http.Error(w, "Unathorized", http.StatusUnauthorized)
// 			return
// 		}
// 		headerArr := strings.Split(header, " ")
// 		if len(headerArr) != 2 {
// 			http.Error(w, "Unathorized", http.StatusUnauthorized)
// 			return
// 		}
// 		accessToken := headerArr[1]
// 		tokenParts := strings.Split(accessToken, ".")
// 		if len(tokenParts) != 3 {
// 			http.Error(w, "Unathorized", http.StatusUnauthorized)
// 			return
// 		}
// 		jwtHeader := tokenParts[0]
// 		jwtPayload := tokenParts[1]
// 		signature := tokenParts[2]
// 		message := jwtHeader + "." + jwtPayload
// 		byteArrMessage := []byte(message)
// 		byteArrSecret := []byte(m.cnf.JwtSecretKey)
// 		h := hmac.New(sha256.New, byteArrSecret)
// 		h.Write(byteArrMessage)
// 		hash := h.Sum(nil)
// 		newsignature := base64URLEncode(hash)
// 		if newsignature != signature {
// 			http.Error(w, "hacker", http.StatusUnauthorized)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }

// =====================================
// 29. REST/MIDDLEWARES/CORS.GO
// =====================================
// package middleware
//
// import "net/http"
//
// func Cors(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Method", "GET, POST, PUT, PATCH, DELETE")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 		w.Header().Set("Content-Type", "application/json")
// 		next.ServeHTTP(w, r)
// 	})
// }

// =====================================
// 30. REST/MIDDLEWARES/LOGGER.GO
// =====================================
// package middleware
//
// import (
// 	"log"
// 	"net/http"
// 	"time"
// )
//
// func Logger(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		startTime := time.Now()
// 		next.ServeHTTP(w, r)
// 		log.Println(r.Method, r.URL.Path, time.Since(startTime))
// 	})
// }

// =====================================
// 31. REST/MIDDLEWARES/PREFLIGHT.GO
// =====================================
// package middleware
//
// import "net/http"
//
// func Preflight(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method == "OPTIONS" {
// 			w.WriteHeader(200)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }

// =====================================
// 32. USER/PORT.GO
// =====================================
// package user
//
// import (
// 	"first-program/domain"
// 	userHandler "first-program/rest/handlers/user"
// )
//
// type Service interface {
// 	userHandler.Service
// }
//
// type UserRepo interface {
// 	Create(user domain.User) (*domain.User, error)
// 	Find(email string, pass string) (*domain.User, error)
// }

// =====================================
// 33. USER/SERVICE.GO
// =====================================
// package user
//
// import "first-program/domain"
//
// type service struct {
// 	usrRepo UserRepo
// }
//
// func NewService(usrRepo UserRepo) Service {
// 	return &service{
// 		usrRepo: usrRepo,
// 	}
// }
//
// func (svc *service) Create(user domain.User) (*domain.User, error) {
// 	usr, err := svc.usrRepo.Create(user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if usr == nil {
// 		return nil, nil
// 	}
// 	return usr, nil
// }
//
// func (svc *service) Find(email string, pass string) (*domain.User, error) {
// 	usr, err := svc.usrRepo.Find(email, pass)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if usr == nil {
// 		return nil, nil
// 	}
// 	return usr, nil
// }

// =====================================
// 34. UTIL/CREAT_JWT.GO
// =====================================
// package util
//
// import (
// 	"crypto/hmac"
// 	"crypto/sha256"
// 	"encoding/base64"
// 	"encoding/json"
// )
//
// type Header struct {
// 	Alg string `json:"alg"`
// 	Typ string `json:"typ"`
// }
//
// type Payload struct {
// 	Sub         int    `json:"sub"`
// 	FirstName   string `json:"first_name"`
// 	LastName    string `json:"last_name"`
// 	Email       string `json:"email"`
// 	IsShopOwner bool   `json:"is_shop_owner"`
// }
//
// func base64URLEncode(data []byte) string {
// 	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
// }
//
// func CreateJWT(secret string, data Payload) (string, error) {
// 	header := Header{
// 		Alg: "HS256",
// 		Typ: "JWT",
// 	}
// 	bytArrHeader, err := json.Marshal(header)
// 	if err != nil {
// 		return "", err
// 	}
// 	headerB64 := base64URLEncode(bytArrHeader)
// 	byteArrData, err := json.Marshal(data)
// 	if err != nil {
// 		return "", err
// 	}
// 	payload64 := base64URLEncode(byteArrData)
// 	message := headerB64 + "." + payload64
// 	byteArrMessage := []byte(message)
// 	byteArrSecret := []byte(secret)
// 	h := hmac.New(sha256.New, byteArrSecret)
// 	h.Write(byteArrMessage)
// 	signature := h.Sum(nil)
// 	signatureB64 := base64URLEncode(signature)
// 	jwt := headerB64 + "." + payload64 + "." + signatureB64
// 	return jwt, nil
// }

// =====================================
// 35. UTIL/SENDDATA.GO
// =====================================
// package util
//
// import (
// 	"encoding/json"
// 	"net/http"
// )
//
// func SendData(w http.ResponseWriter, StatusCode int, data interface{}) {
// 	w.WriteHeader(StatusCode)
// 	encoder := json.NewEncoder(w)
// 	encoder.Encode(data)
// }
//
// func SendError(w http.ResponseWriter, statusCode int, msg string) {
// 	w.WriteHeader(statusCode)
// 	encoder := json.NewEncoder(w)
// 	encoder.Encode(msg)
// }

// =====================================
// DATABASE QUERIES
// =====================================

// 1. DB_QUERIES/00011-INSERT.SQL
// INSERT INTO users (
//     first_name, last_name, email, password
// ) VALUES (
//     'Habibur', 'Rahman', 'habibur@gmail.com', '12345'
// )

// 2. DB_QUERIES/00012-SELECT.SQL
// SELECT * FROM users;
// SELECT first_name, last_name FROM users;

// 3. DB_QUERIES/00013-UPDATE.SQL
// UPDATE users SET
//     first_name = 'Ismat', last_name = 'Rahman'
// WHERE id = 10

// 4. DB_QUERIES/00014-DELETE.SQL
// DELETE FROM users WHERE id = 10;

// 5. DB_QUERIES/001-CREATE-USERTABLE.SQL
// CREATE TABLE users (
// 	id SERIAL PRIMARY KEY,
// 	first_name VARCHAR(255) NOT NULL,
// 	last_name VARCHAR (255) NOT NULL,
// 	email VARCHAR(255) UNIQUE NOT NULL,
// 	password VARCHAR(255) NOT NULL,
// 	is_shop_owner BOOLEAN DEFAULT FALSE,
// 	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
// 	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
// );

// 6. DB_QUERIES/002-CREATE-PRODUCTTABLE.SQL
// CREATE TABLE products (
// 	id BIGSERIAL PRIMARY KEY,
// 	title VARCHAR(255) NOT NULL,
// 	description TEXT,
// 	price DOUBLE PRECISION NOT NULL,
// 	img_url TEXT,
// 	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
// 	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
// );

// 7. DB_QUERIES/002-PRODUCT-TABLE.SQL
// CREATE TABLE users (
// 	id BIGSERIAL PRIMARY KEY,
// 	title VARCHAR(255) NOT NULL,
// 	description TEXT,
// 	price DOUBLE PRECISION NOT NULL,
// 	img_url TEXT,
// 	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
// 	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
// );

// =====================================
// MIGRATIONS
// =====================================

// 1. MIGRATIONS/000001-CREATE-USER.UP.SQL
// -- +migrate Up
// CREATE TABLE IF NOT EXISTS users (
// 	id SERIAL PRIMARY KEY,
// 	first_name VARCHAR(255) NOT NULL,
// 	last_name VARCHAR (255) NOT NULL,
// 	email VARCHAR(255) UNIQUE NOT NULL,
// 	password VARCHAR(255) NOT NULL,
// 	is_shop_owner BOOLEAN DEFAULT FALSE,
// 	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
// 	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
// );

// 2. MIGRATIONS/000001-CREATE-USER.DOWN.SQL
// -- +migrate Down
// DROP TABLE IF EXISTS users;

// 3. MIGRATIONS/000002-CREATE-PRODUCTS.UP.SQL
// -- +migrate Up
// CREATE TABLE IF NOT EXISTS products (
// 	id BIGSERIAL PRIMARY KEY,
// 	title VARCHAR(255) NOT NULL,
// 	description TEXT,
// 	price DOUBLE PRECISION NOT NULL,
// 	img_url TEXT,
// 	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
// 	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
// );

// 4. MIGRATIONS/000002-CREATE-PRODUCTS.DOWN.SQL
// -- +migrate Down
// DROP TABLE IF EXISTS products;

// =====================================
// END OF PROJECT CONSOLIDATION
// This file consolidates all source code in the project
// =====================================
