package db

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	"github.com/MrSterdy/ApolloHW/model/holiday"
	"github.com/MrSterdy/ApolloHW/model/student"
	"github.com/MrSterdy/ApolloHW/model/timetable"
)

func InitMySQL() {
	sqldb, err := sql.Open("mysql", os.Getenv("DB_MYSQL"))
	if err != nil {
		panic(err)
	}

	mysql := bun.NewDB(sqldb, mysqldialect.New())
	ctx := context.Background()

	holiday.Init(mysql, ctx)
	student.Init(mysql, ctx)
	timetable.Init(mysql, ctx)
}
