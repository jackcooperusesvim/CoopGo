package model

import (
	"context"
	"database/sql"
	_ "embed"
	"github.com/jackcooperusesvim/coopGo/model/sqlgen"
	"log"
)

//go:embed schema.sql
var ddl string

const DATABASE_LOCATION = "data.sqlite3"

func CreateTables() error {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", DATABASE_LOCATION)
	if err != nil {
		return err
	}
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}
	return nil
}

func BuildTables() error {
	q, ctx, err := DbInfo()
	names := []string{"CompSci", "JavaScript", "Mathematics", "History", "Science"}
	desc := []string{"Interesting, but sometimes boring", "Horrible", "Ok", "Nonononono", "Boring"}
	start_date := []string{"2000-12-30", "2000-12-30", "2000-12-30", "2000-12-30", "2000-12-30"}
	end_date := []string{"1999-01-01", "1999-01-01", "1999-01-01", "1999-01-01", "1999-01-01"}

	if err != nil {
		return err
	}
	for i := range names {
		_, err := q.CreateCourse(ctx, sqlgen.CreateCourseParams{
			Name:      names[i],
			Desc:      desc[i],
			StartDate: start_date[i],
			EndDate:   end_date[i],
		})

		if err != nil {
			return err
		}
	}
	return nil
}

func DbInfo() (*sqlgen.Queries, context.Context, error) {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", DATABASE_LOCATION)

	if err != nil {
		return nil, nil, err
	}

	q := sqlgen.New(db)
	// courses, err := q.ListCourse(ctx)
	// if err != nil {
	// 	print("error")
	// 	return nil, nil, err
	// }
	return q, ctx, nil
}
func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return
	}

	queries := sqlgen.New(db)

	// list all authors
	authors, err := queries.ListCourse(ctx)
	if err != nil {
		return
	}
	log.Println(authors)
}
