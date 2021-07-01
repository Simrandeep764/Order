package Routes

import (
	"Order/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter (for configuring ports and running on 8004)
func SetupRouter() {

	r4 := gin.Default()
	grp5 := r4.Group("v1/")
	{
		grp5.GET("order/:oid", Controllers.GetOrderDetailsById)
		grp5.POST("order", Controllers.GenerateOrder)
	}

	r4.Run(":7004")
}
