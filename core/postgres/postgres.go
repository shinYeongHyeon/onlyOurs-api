package postgres

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/shinYeongHyeon/onlyOurs-api/ent"
	"os"
)

func Connect() {
	url := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))

	client, err := ent.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
