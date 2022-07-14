package Api_Services

import (
	QRHandler "Maria_Demo/QR_Handler"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Add_QR_Group(app *gin.Engine, db_connection *sql.DB) *gin.RouterGroup {
	qr_group := app.Group("/qr")
	{
		qr_group.GET("/:truck_id", func(c *gin.Context) {
			qr_image, err := QRHandler.Generate_QR("http://localhost:3000/trucks/" + c.Param("truck_id"))
			if err != nil {
				c.JSON(500, err)
			}
			c.Data(200, "image/png", qr_image)
		})
	}
	return qr_group
}
