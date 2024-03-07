package models

type Project struct {
	ID   int
	Name string

	Urls []Url
}
