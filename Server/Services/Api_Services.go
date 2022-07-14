package Api_Services

import (
	Authentication_Handler "Maria_Demo/Authentication"
	struct_def "Maria_Demo/Structs"
	"Maria_Demo/convert_handler"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Add_App_Group(app *gin.Engine, db_connection *sql.DB) *gin.RouterGroup {
	// Create a new router group for the app to handle database requests
	api_group := app.Group("/api")
	{
		app.GET("/api/:truck_id", func(c *gin.Context) {
			var user struct_def.Jwt_Payload = c.MustGet("token_payload").(struct_def.Jwt_Payload)
			if !Authentication_Handler.Check_Roles(user.Roles, []string{"admin", "user"}) {
				c.JSON(401, gin.H{"error": "You are not authorized to access this resource."})
				return
			}
			truck_id := c.Param("truck_id")
			rows, err := db_connection.Query("SELECT * FROM trucks.trucks WHERE truck_id = ?", truck_id)
			if err != nil {
				c.JSON(400, err)
			}

			truck_data, err := convert_handler.Dataset_To_Truck_Data(rows)
			if err != nil {
				c.JSON(500, err)
			}

			c.JSON(200, truck_data)
		})
	}
	return api_group
}
