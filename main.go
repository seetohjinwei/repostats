package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/seetohjinwei/repostats/cli"
	"github.com/seetohjinwei/repostats/messages"
	"github.com/seetohjinwei/repostats/web"
)

var modes = map[string]func(*pgxpool.Pool){
	"cli": cli.Start,
	"web": web.Start,
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	maybeMode := flag.String("mode", "cli", "either cli or web")
	flag.Parse()

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	found := false
	for mode, app := range modes {
		if *maybeMode == mode {
			app(pool)
			return
		}
	}
	if !found {
		fmt.Printf(messages.INVALID_MODE, *maybeMode)
	}
}
