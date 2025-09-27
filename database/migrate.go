package database

import (
	"fmt"

	"github.com/ADMex1/goweb/model"
)

func Migration() {
	err := DB.AutoMigrate(&model.Book{})
	if err != nil {
		panic("Migration Failed" + err.Error())
	}
	fmt.Println("Migration success")
}
