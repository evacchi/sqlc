// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package querytest

import (
	"database/sql"
)

type User struct {
	FirstName sql.NullString `json:"firstName"`
	LastName  sql.NullString `json:"lastName"`
	Age       sql.NullInt16  `json:"age"`
}
