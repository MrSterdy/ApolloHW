package timetable

import (
	"context"
	"database/sql"
	"time"

	"github.com/MrSterdy/ApolloHW/util"

	"github.com/uptrace/bun"
)

var db *bun.DB
var ctx context.Context

func Init(database *bun.DB, databaseCtx context.Context) {
	db = database
	ctx = databaseCtx

	_, err := db.NewCreateTable().IfNotExists().Model((*Subject)(nil)).Exec(ctx)
	if err != nil {
		panic(err)
	}

	_, err = db.NewCreateTable().IfNotExists().Model((*Weekday)(nil)).Exec(ctx)
	if err != nil {
		panic(err)
	}

	_, err = db.NewCreateTable().IfNotExists().Model((*Date)(nil)).Exec(ctx)
	if err != nil {
		panic(err)
	}
}

func FindByDate(date *util.Date) *Date {
	dt := new(Date)

	err := db.NewSelect().Model(dt).Where("date = ?", date).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		panic(err)
	}

	return dt
}

func FindByWeekday(weekday time.Weekday) *Weekday {
	wt := new(Weekday)

	err := db.NewSelect().Model(wt).Where("day = ?", weekday).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		panic(err)
	}

	return wt
}

func FindNextTimetable(subject string, startDate *util.Date) (Timetable, *util.Date) {
	day := startDate.Weekday()

	wt := findNearestWeekday(day, subject)
	if wt == nil {
		return nil, nil
	}

	nearestDay := int(wt.Day) - int(day)
	if nearestDay <= 0 {
		nearestDay += 7
	}

	nearestDate := &util.Date{Time: startDate.AddDate(0, 0, nearestDay)}

	dt := findInDateRange(startDate, nearestDate, subject)
	if dt == nil {
		dt = FindByDate(nearestDate)
		if dt == nil {
			return wt, nearestDate
		}

		return FindNextTimetable(subject, nearestDate)
	}

	return dt, dt.Date
}

func findInDateRange(startDate, endDate *util.Date, subject string) *Date {
	dt := new(Date)

	err := db.NewSelect().
		Model(dt).
		Where("date > ?", startDate).
		Where("date <= ?", endDate).
		Where("JSON_CONTAINS(subjects, '{\"name\":\"?\"}')", bun.Safe(subject)).
		Order("date ASC").
		Limit(1).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		panic(err)
	}

	return dt
}

func findNearestWeekday(day time.Weekday, subject string) *Weekday {
	wt := new(Weekday)

	err := db.NewSelect().
		Model(wt).
		Where("day > ?", day).
		Where("JSON_CONTAINS(subjects, '{\"name\":\"?\"}')", bun.Safe(subject)).
		Order("day ASC").
		Limit(1).
		Scan(ctx)
	if err == sql.ErrNoRows {
		err = db.NewSelect().
			Model(wt).
			Where("day <= ?", day).
			Where("JSON_CONTAINS(subjects, '{\"name\":\"?\"}')", bun.Safe(subject)).
			Order("day ASC").
			Limit(1).
			Scan(ctx)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}

			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	return wt
}

func GetSubjects() Subjects {
	var subjects Subjects

	err := db.NewSelect().Model(&subjects).Scan(ctx)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	return subjects
}

func SaveSubjects(subjects Subjects) {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		panic(err)
	}

	_, err = tx.NewTruncateTable().Model((*Subject)(nil)).Exec(ctx)
	if err != nil {
		panic(err)
	}
	_, err = tx.NewInsert().Model(&subjects).Exec(ctx)
	if err != nil {
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}

func DeleteSubjects() {
	if _, err := db.NewTruncateTable().Model((*Subject)(nil)).Exec(ctx); err != nil {
		panic(err)
	}
}

func SaveByDate(dt *Date) {
	_, err := db.NewInsert().Model(dt).On("DUPLICATE KEY UPDATE").Exec(ctx)
	if err != nil {
		panic(err)
	}
}

func SaveByWeekday(wt *Weekday) {
	_, err := db.NewInsert().Model(wt).On("DUPLICATE KEY UPDATE").Exec(ctx)
	if err != nil {
		panic(err)
	}
}
