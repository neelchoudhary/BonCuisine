package models

// Ingredient ...
type Ingredient struct {
	ID       int64  `json:"ingredient_id"`
	IngredientName string `json:"ingredient_name"`
	ImageID int64 `json:"image_id"`
}
