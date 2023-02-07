package start

import (
	"context"
	"fmt"

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
