package models

// Recipe ...
type Recipe struct {
	ID       int64  `json:"recipe_id"`
	RecipeName string `json:"recipe_name"`
	Time string `json:"r_time"`
	Servings    string `json:"num_servings"`
	Difficulty string `json:"difficulty"`
	CuisineName string `json:"cuisine_name"`
	ImageData string `json:"image_data"`
	
}