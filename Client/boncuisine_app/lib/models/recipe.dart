import 'package:flutter/cupertino.dart';

class Recipe {
  int recipeID;
  String recipeName;
  String recipeTime;
  String servings;
  String difficulty;
  Image image;
  String cuisineName;

  static Recipe fromJSON(Map json) {
    Recipe recipe = Recipe();
    recipe.recipeID = json["id"];
    recipe.recipeName = json["name"];
    recipe.recipeTime = json["time"];
    recipe.servings = json["servings"];
    recipe.difficulty = json["difficulty"];
    recipe.image = json["image"];
    recipe.cuisineName = json["cuisineName"];
    return recipe;
  }
}
