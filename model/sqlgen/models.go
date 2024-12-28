// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlgen

import (
	"database/sql"
)

type Child struct {
	ID          int64
	FirstName   string
	LastName    sql.NullString
	Birthday    string
	GradeOffset int64
	FamilyID    sql.NullInt64
}

type Course struct {
	ID        int64
	Name      string
	Desc      string
	StartDate sql.NullString
	EndDate   sql.NullString
}

type Coursejoinchild struct {
	ID       int64
	CourseID int64
	ChildID  int64
}

type Family struct {
	ID         int64
	LastName   string
	MainParent string
	SecParent  string
	Phone1     string
	Phone1     string
	Phone2     string
	Phone3     sql.NullString
}