package entity

import "time"

type User struct {
	ID       string
	Name     string
	Mail     string
	Password string
}

type TimeLineCalendar struct {
	ID            string
	Date          time.Time
	Name          string
	TimeLineStart time.Time
	TimeLIneEnd   time.Time
	Color         string
}

type Task struct {
	ID            string
	UserID        string
	Name          string
	DescriptionID []string
	SubTasksID    []string
	NotesID       []string
	TagsID        []string
	StatusID      string
	Severity      int
	DateUse       time.Time
}

type Note struct {
	ID       string
	UserID   string
	TypeNone string
}

type Tag struct {
	ID     string
	UserID string
	Name   string
	Color  string
}

type Status struct {
	ID     string
	UserID string
	Name   string
}
