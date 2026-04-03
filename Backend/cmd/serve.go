package cmd

import (
	"first-program/config"
	"first-program/infra"
	"first-program/product"
	"first-program/repo"
	"first-program/rest"
	prdHandler "first-program/rest/handlers/product"
	usrHandler "first-program/rest/handlers/user"
	middleware "first-program/rest/middlewares"
	"first-program/user"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := infra.NewConnection(cnf.DB) // db connection

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = infra.MigrateDB(dbCon, "./migrations") // run the database migrations using the sql-migrate package
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// repos
	userRepo := repo.NewUserRepo(dbCon)       // create a new user repository using the sqlx database connection
	productRepo := repo.NewProductRepo(dbCon) // create a new product repository using the sqlx database connection

	// domains
	userSvc := user.NewService(userRepo)
	prdSvc := product.NewService(productRepo)
	middlewares := middleware.NewMiddlewares(cnf)

	// Handlers
	productHandler := prdHandler.NewHandler(middlewares, prdSvc) // create a new product handler using the middlewares and product service
	userHandler := usrHandler.NewHandler(cnf, userSvc)           // create a new user handler using the config and user service

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}

/*
"pipeline" :

globalrouter -> hudai middleware -> logger middleware -> arekta middleware -> Request handlers

*/
