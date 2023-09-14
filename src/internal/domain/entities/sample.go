package entities

type Sample struct {
	Name        string `json:"name" bson:"name"`
	Value       int    `json:"value" bson:"value"`
	Description string `json:"description" bson:"description"`
}
