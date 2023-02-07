package main

import (
	"context"
	"fmt"
	"log"

	"github.com/GRTheory/my-model/ent"

	_ "github.com/go-sql-driver/mysql"
)

func GetClient(username, password, host, port, database string) (*ent.Client, error) {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", username, password, host, port, database)
	client, err := ent.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	err = client.Schema.Create(context.Background())
	return client, err
}

func main() {
	username := "root"
	password := "pass_123"
	host := "192.168.228.128"
	port := "3306"
	database := "base"
	client, err := GetClient(username, password, host, port, database)
	if err != nil {
		log.Fatalf("failed getting a new client: %v", err)
	}

	ctx := context.Background()
	teacher, err := client.Teacher.
		Create().
		SetTID(89).
		SetName("teacher1").
		SetAge(23).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed synchronizing entities to database: %v", err)
	}
	log.Println(teacher)
}
