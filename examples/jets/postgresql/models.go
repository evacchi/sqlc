// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package jets

import ()

type Jet struct {
	ID      int32
	PilotID int32
	Age     int32
	Name    string
	Color   string
}

type Language struct {
	ID       int32
	Language string
}

type Pilot struct {
	ID   int32
	Name string
}

type PilotLanguage struct {
	PilotID    int32
	LanguageID int32
}
