package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/farrukhnaveed/go-ecom/cmd/api"
	"github.com/farrukhnaveed/go-ecom/db"
	"github.com/farrukhnaveed/go-ecom/config"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Net: "tcp",
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		AllowNativePasswords: true,
		ParseTime: true,
	})

	if (err != nil) {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(config.Envs.Port, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

function initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Successfully connected to database")
}