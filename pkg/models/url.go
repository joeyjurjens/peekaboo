package models

type Url struct {
	ID  int
	Url string

	// Relations
	Project     Project
	Screenshots []Screenshot
}
