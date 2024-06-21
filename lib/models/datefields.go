package models

import "time"

// DateFields struct contains the `created_at` and `updated_at`
// fields that autofill when inserting or updating a model.
type DateFields struct {
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updated_at"`
}

// Creating hook is used here to set the `created_at` field
// value when inserting a new model into the database.
func (f *DateFields) Creating() error {
	f.CreatedAt = time.Now().UTC()
	return nil
}

// Saving hook is used here to set the `updated_at` field
// value when creating or updating a model.
// TODO: get context as param the next version(4).
func (f *DateFields) Saving() error {
	f.UpdatedAt = time.Now().UTC()
	return nil
}
