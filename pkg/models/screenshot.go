package models

const (
	SCREENSHOT_TYPE_OLD  = "old"
	SCREENSHOT_TYPE_NEW  = "new"
	SCREENSHOT_TYPE_DIFF = "diff"
)

type Screenshot struct {
	ID int

	Type     string
	FilePath string

	// Relations
	Url Url
}
