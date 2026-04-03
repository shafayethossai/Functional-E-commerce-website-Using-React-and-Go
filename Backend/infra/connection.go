package infra

import (
	"first-program/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DBConfig) string {
	// user -> postgresql
	// password -> shafayet@@@01851287806
	// host -> localhost
	// port -> 5432
	// db name -> ecommerce

	connString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
	)
	if !cnf.EnableSSLMODE {
		connString += " sslmode=disable"
	}
	return connString
}

func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)             // database connection string
	dbCon, err := sqlx.Connect("postgres", dbSource) // connect to the postgres database using sqlx
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbCon, nil
}
