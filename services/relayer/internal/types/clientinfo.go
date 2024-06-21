package types

type ClientInfo struct {
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Host        string `json:"host" bson:"host"`
	ID          string `json:"id" bson:"id"`
	Icon        string `json:"icon,omitempty" bson:"icon,omitempty"`
	Name        string `json:"name" bson:"name"`
}
