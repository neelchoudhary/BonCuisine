package models

// SavedRecipe ...
type SavedRecipe struct {
	UserID   string `json:"user_id"`
	RecipeID int64  `json:"recipe_id"`
}
