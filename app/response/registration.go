package response

import "time"

type ListRegistration struct {
	ID          string    `bun:"id" json:"id"`
	ActivityID  string    `bun:"activity_id" json:"activity_id"`
	Name        string    `bun:"name" json:"name"`
	Description string    `bun:"description" json:"description"`
	ReleaseDate time.Time `bun:"release_date" json:"release_date"`

	StudentID     string `bun:"student_id" json:"student_id"`
	FirstName     string `bun:"first_name" json:"first_name"`
	LastName      string `bun:"last_name" json:"last_name"`
	StudentNumber string `bun:"student_number" json:"student_number"`

	CreatedAt int64 `bun:"created_at" json:"created_at"`
	UpdatedAt int64 `bun:"updated_at" json:"updated_at"`
}
