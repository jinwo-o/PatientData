package models

// Post type details
type Post struct {
	ID      int64  `json:"ID"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Disease string `json:"disease"`
}
