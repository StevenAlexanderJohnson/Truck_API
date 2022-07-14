package main

import (
	Api_Services "Maria_Demo/Services"
	struct_def "Maria_Demo/Structs"
	"Maria_Demo/Token_Handler"
	"bufio"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func validate_token_middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// If they are attempting to sign in or sign up, skip the validation
		if c.Request.URL.Path == "/sign_in" || c.Request.URL.Path == "/register" {
			c.Next()
			return
		}

		// Get the token from the request
		token, err := c.Cookie("Auth_Token")
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		// Validate the token
		validated, err := Token_Handler.Verify_Token(token)
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		// If the token is valid, return unauthorized
		if !validated {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
	}
}

func parse_token_middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// If they are attempting to login or register skip the validation
		if c.Request.URL.Path == "/sign_in" || c.Request.URL.Path == "/register" {
			c.Next()
			return
		}
		// Get the token from the request
		token, err := c.Cookie("Auth_Token")
		if err != nil {
			fmt.Println("ERR: ?", err)
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		// Validate the token
		token_data, err := Token_Handler.Read_Token_Data(token)
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		var token_payload struct_def.Jwt_Payload
		err = json.Unmarshal(token_data, &token_payload)
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Add the token payload to the session
		c.Set("token_payload", token_payload)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func initialize_database(db *sql.DB) error {
	fmt.Println("Executing Schema...")
	err := filepath.Walk("./Database/schema", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		fmt.Printf("Executing: %s\n", info.Name())
		script, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		_, err = db.Exec(string(script))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	fmt.Println("Schema Executed")

	fmt.Println("Initializing Data...")
	data_file, err := os.Open("./Database/load_data.sql")
	if err != nil {
		fmt.Println("ERR: ?", err)
		return err
	}
	defer data_file.Close()

	scanner := bufio.NewScanner(data_file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		_, err := db.Exec(scanner.Text())
		if err != nil {
			return err
		}
	}
	fmt.Println("Data Executed")
	return nil
}

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
		err := initialize_database(db_connection)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Database Initialized")
	}

	// Create the app to handle the requests
	app := gin.Default()
	// Set the cors headers
	app.Use(CORSMiddleware())
	// Add the middleware to validate the token
	app.Use(validate_token_middleware())
	app.Use(parse_token_middleware())

	// Setup services to handle different routes
	Api_Services.Add_Root_Group(app, db_connection)
	Api_Services.Add_App_Group(app, db_connection)
	Api_Services.Add_QR_Group(app, db_connection)

	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "NO_EXISTING_ROUTE", "message": "No existing route"})
	})

	app.Run()
}
