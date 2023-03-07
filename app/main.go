package main

import (
	"app/controllers"
	"app/models"
	"fmt"
	"log"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)
	//
	// log.Println("test")

	err := controllers.StartMainServer()
	if err != nil {
		log.Println(err)
		return
	}

	// CREATE
	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.Password = "testtest"
	fmt.Println(u)
	// u.CreateUser()

	// READ
	user, err := models.GetUser(1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(user)

	// UPDATE
	user.Name = "test2"
	user.Email = "test2@example.com"
	user.UpdateUser()
	fmt.Println(user)

	// DELETE
	user.DeleteUser()
	user, _ = models.GetUser(1)
	fmt.Println(user)

}
