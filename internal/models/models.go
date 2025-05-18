package models

import "time"

// type Events struct {
// 	Type string `json:"type"`
// 	Repo
// }

type API_Response []struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		Commits []struct {
			Author struct {
				Email string `json:"email"`
				Name  string `json:"name"`
			} `json:"author"`
			Message string `json:"message"`
		} `json:"commits"`
	} `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}
