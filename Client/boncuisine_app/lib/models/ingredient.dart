import 'package:flutter/cupertino.dart';

class Ingredient {
  int ingredientID;
  String name;
  String amount;
  Image image;

  static Ingredient fromJSON(Map json) {
    Ingredient ingredient = Ingredient();
    ingredient.ingredientID = json["id"];
    ingredient.name = json["name"];
    ingredient.amount = json["amount"];
    ingredient.image = json["image"];
    return ingredient;
  }
}
