package database

import (
	"os"

	"github.com/go-rel/mysql"
	"github.com/go-rel/rel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Database() rel.Repository {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	uriDb := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@/" + os.Getenv("DB_NAME")

	adapter, err := mysql.Open(uriDb)

	if err != nil {
		panic(err)
	}

	return rel.New(adapter)

}
