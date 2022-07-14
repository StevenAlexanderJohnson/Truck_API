package Api_Services

import (
	Authentication_Handler "Maria_Demo/Authentication"
	struct_def "Maria_Demo/Structs"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Add_Root_Group(app *gin.Engine, db_connection *sql.DB) *gin.RouterGroup {
	// Create a new router group for the app to handle database requests
	api_group := app.Group("/")
	{
		app.GET("/", func(c *gin.Context) {
			var user struct_def.Jwt_Payload = c.MustGet("token_payload").(struct_def.Jwt_Payload)
			Authentication_Handler.Check_Roles(user.Roles, []string{"user"})
			c.JSON(200, "Welcom to the Demo: "+user.User+"!")
		})

		app.POST("/sign_in", func(c *gin.Context) {
			var login_data struct_def.Login_Data
			c.BindJSON(&login_data)
			token, err := Authentication_Handler.Sign_In(login_data.Email, login_data.Password, db_connection)
			if err != nil {
				fmt.Println(err)
				c.JSON(401, gin.H{"error": "Invalid Credentials"})
				return
			}
			c.SetCookie("Auth_Token", token, 3600, "/", "localhost", false, true)
			c.JSON(200, gin.H{"success": true})
			return
		})

		app.POST("/register", func(c *gin.Context) {
			var register_data struct_def.Register_Data
			c.BindJSON(&register_data)
			token, err := Authentication_Handler.Register(register_data.Email, register_data.Username, register_data.Password, db_connection)
			if err != nil {
				fmt.Println(err)
				c.JSON(401, gin.H{"error": "Credentials are already in use."})
				return
			}
			c.SetCookie("Auth_Token", token, 3600, "/", "localhost", false, true)
			c.JSON(200, gin.H{"success": true})
			return
		})

		app.POST("/logout", func(c *gin.Context) {
			c.SetCookie("Auth_Token", "", -1, "/", "localhost", false, true)
			c.JSON(200, gin.H{"success": true})
			return
		})

		app.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	return api_group
}
