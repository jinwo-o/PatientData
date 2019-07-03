package models

// Post type details
type Post struct {
	patient_id int64  `json:"patient_id"`
	name       string `json:"name"`
	gender     string `json:"gender"`
	disease    string `json:"disease"`
	// created_at time.Time `json:"created_at"`
	// updated_at time.Time `json:"updated_at"`
}
