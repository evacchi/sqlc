// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package querytest

import (
	"time"

	"github.com/google/uuid"
)

type TddTest struct {
	TestID    uuid.UUID
	Title     string
	Descr     string
	TsCreated time.Time
	TsUpdated time.Time
}
