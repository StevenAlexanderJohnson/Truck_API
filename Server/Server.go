package main

import (
	Database "Maria_Demo/Database"
	Middleware "Maria_Demo/Middleware"
	Api_Services "Maria_Demo/Services"
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Set a flag for initializing the database
	var database_init bool
	flag.BoolVar(&database_init, "init", false, "Initialize the database")
	flag.Parse()

	// Create a connection to the database
	db_connection, err := sql.Open("mysql", "root:demo@tcp(localhost:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	defer db_connection.Close()

	// Check if the user want to initialize a new database
	if database_init {
		fmt.Println("Initializing Database...")
		// Initialize the database and check for errors
		err := Database.Initialize_Database(db_connection)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Database Initialized")
	}

	// Create the app to handle the requests
	app := gin.Default()
	// Set the cors headers
	app.Use(Middleware.CORSMiddleware())
	// Add the middleware to validate the token
	app.Use(Middleware.Validate_Token_Middleware())
	app.Use(Middleware.Parse_Token_Middleware())

	// Setup services to handle different routes
	Api_Services.Add_Root_Group(app, db_connection)
	Api_Services.Add_App_Group(app, db_connection)
	Api_Services.Add_QR_Group(app, db_connection)

	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "NO_EXISTING_ROUTE", "message": "No existing route"})
	})

	app.Run()
}
