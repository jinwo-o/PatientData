package models

// Post type details
type Post struct {
	Patient_id int64  `json:"patient_id"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Disease    string `json:"disease"`
	// created_at time.Time `json:"created_at"`
	// updated_at time.Time `json:"updated_at"`
}
