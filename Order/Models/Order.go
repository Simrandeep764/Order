package Models

import (
	"Order/Config"
	_ "fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GenerateOrder(ord *Order) (err error) {
	if err = Config.DB.Create(ord).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderDetailsById(ord *[]Order, id string) (err error) {
	if err = Config.DB.Where("order_id = ?", id).Find(ord).Error; err != nil {
		return err
	}
	return nil
}
