package start

import (
	"context"
	"log"
	"testing"
)

func TestGetClient(t *testing.T) {
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
