package entity

import "time"

type Task struct {
	ID          string
	UserID      string
	Name        string
	Description []Note
	SubTask     []Task
	StatusID    string
	Severity    int
	DateUse     time.Time
}

type Note struct {
	ID       string
	UserID   string
	TypeNone string
}

type User struct {
	ID       string
	Name     string
	Mail     string
	Password string
}

// Tag TODO: Добавить цвет тегу
type Tag struct {
	ID     string
	UserID string
	Name   string
}

type Status struct {
	ID     string
	UserID string
	Name   string
}
