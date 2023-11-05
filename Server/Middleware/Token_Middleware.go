package Middleware

import (
	struct_def "Maria_Demo/Structs"
	"Maria_Demo/Token_Handler"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Validate_Token_Middleware() gin.HandlerFunc {
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
		fmt.Println(token)
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

func Parse_Token_Middleware() gin.HandlerFunc {
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

		if token_payload.Expires <= time.Now().UTC().UnixNano() {
			c.JSON(401, gin.H{"error": "Session Expired"})
		}

		// Add the token payload to the session
		c.Set("token_payload", token_payload)
	}
}
