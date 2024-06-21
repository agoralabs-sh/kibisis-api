package models

import "github.com/kamva/mgm/v3"

// DefaultModel struct contains a model's default fields: _id, created_at and updated_at.
type DefaultModel struct {
	mgm.IDField `bson:",inline"`
	DateFields  `bson:",inline"`
}

// Creating function calls the inner fields' defined hooks
func (model *DefaultModel) Creating() error {
	return model.DateFields.Creating()
}

// Saving function calls the inner fields' defined hooks
func (model *DefaultModel) Saving() error {
	return model.DateFields.Saving()
}
