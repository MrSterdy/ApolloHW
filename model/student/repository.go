package student

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

	_, err := db.NewCreateTable().IfNotExists().Model((*Student)(nil)).Exec(ctx)
	if err != nil {
		panic(err)
	}
}

func IsEmpty() bool {
	exists, err := db.NewSelect().Model((*Student)(nil)).Exists(ctx)
	if err != nil {
		panic(err)
	}

	return !exists
}

func Find(id int64) *Student {
	student := new(Student)

	err := db.NewSelect().Model(student).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		panic(err)
	}

	return student
}

func GetAll() []*Student {
	var students []*Student

	err := db.NewSelect().Model(&students).Scan(ctx)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	return students
}

func Save(student *Student) {
	_, err := db.NewInsert().Model(student).On("DUPLICATE KEY UPDATE").Exec(ctx)
	if err != nil {
		panic(err)
	}
}

func SaveAll(students []*Student) {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		panic(err)
	}

	_, err = tx.NewTruncateTable().Model((*Student)(nil)).Exec(ctx)
	if err != nil {
		panic(err)
	}
	_, err = tx.NewInsert().Model(&students).Exec(ctx)
	if err != nil {
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}

func DeleteAll() {
	if _, err := db.NewTruncateTable().Model((*Student)(nil)).Exec(ctx); err != nil {
		panic(err)
	}
}
