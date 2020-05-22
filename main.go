package main

import (
	"fmt"
	"go-gorm-crud-postgresql/driver"
	"go-gorm-crud-postgresql/model"
	"go-gorm-crud-postgresql/repository/client"
	"log"
)

func main() {

	repo := client.IClientRepository{}
	db := driver.GetPostgreSQLConnection()

	fmt.Println("######## Save #######")
	client := model.Client{
		Email: "gorm3@gorm3",
		Password: "gorm123",
	}
	err := repo.Save(&client,db)
	if err !=nil {
		log.Fatal(err)
	}

	//fmt.Println("######## FindById #######")
	//client , err := repo.FindById(2,db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(client)

	//fmt.Println("######## FindAll #######")
	//clients, err :=	repo.FindAll(db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%+v\n", clients)
	//
	//fmt.Println("######## Update #######")
	//client, err := repo.Update(&model.Client{Id: 2, Email: "updatedEmail", UpdateAt: time.Now()}, db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(client)

	//fmt.Println("######## Delete #######")
	//client := model.Client{Id: 2}
	//err := repo.Delete(&client, db)
	//if err != nil {
	//	log.Fatal(err)
	//}

}
