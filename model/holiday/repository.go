package holiday

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
)

var db *bun.DB
var ctx context.Context

func Init(database *bun.DB, databaseCtx context.Context) {
	db = database
	ctx = databaseCtx

	_, err := db.NewCreateTable().IfNotExists().Model((*Holiday)(nil)).Exec(ctx)
	if err != nil {
		panic(err)
	}
}

func GetAll() []Holiday {
	var holidays []Holiday

	err := db.NewSelect().Model(&holidays).Scan(ctx)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	return holidays
}

func SaveAll(holidays []Holiday) {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		panic(err)
	}

	_, err = tx.NewTruncateTable().Model((*Holiday)(nil)).Exec(ctx)
	if err != nil {
		panic(err)
	}
	_, err = tx.NewInsert().Model(&holidays).Exec(ctx)
	if err != nil {
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}

func DeleteAll() {
	if _, err := db.NewTruncateTable().Model((*Holiday)(nil)).Exec(ctx); err != nil {
		panic(err)
	}
}
