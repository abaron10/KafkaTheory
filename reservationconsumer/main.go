package main

import (
	"encoding/json"
	"fmt"
	"reservationconsumer/consumer"
	"reservationconsumer/database"
	"reservationconsumer/models"
)

func main() {

	connector := database.NewConnector()
	if err := connector.CreateTable(); err != nil {
		fmt.Println(err)
		return
	}

	defer connector.Close()

	cr := consumer.NewConsumer(consumer.ConsumerConfig{Host: "localhost:9092", Topics: []string{"message-log"}, GroupId: "my-reservationconsumer"})

	for {
		bytesEvent := <-cr.Poll()

		var user models.User
		json.Unmarshal(bytesEvent, &user)

		if err := connector.InsertTable(user.Name, user.Email); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(bytesEvent))
	}
}
