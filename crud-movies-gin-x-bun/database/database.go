package database

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

var Connect = func() *bun.DB {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	uriDb := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@/" + os.Getenv("DB_NAME")

	DB, err := sql.Open("mysql", uriDb)

	if err != nil {
		panic(err)
	}

	return bun.NewDB(DB, mysqldialect.New())

}()

func MigrateModel(model interface{}) {
	bunDb := bun.NewDB(Connect.DB, mysqldialect.New())

	if err := bunDb.ResetModel(context.TODO(), model); err != nil {
		panic(err)
	}
}
