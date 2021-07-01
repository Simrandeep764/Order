package Controllers

import (
	"Order/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateOrder(c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)
	fmt.Println(order)
	err := Models.GenerateOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func GetOrderDetailsById(c *gin.Context) {
	id := c.Params.ByName("oid")
	var order []Models.Order
	err := Models.GetOrderDetailsById(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
