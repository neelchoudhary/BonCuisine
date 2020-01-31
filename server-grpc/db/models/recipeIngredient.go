package models

// RecipeIngredient ...
type RecipeIngredient struct {
	IngredientID   int64  `json:"ingredient_id"`
	IngredientName string `json:"ingredient_name"`
	Amount         int64  `json:"amount"`
	ImageData      []byte `json:"image_data"`
}
