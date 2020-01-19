import 'package:boncuisine_app/models/cuisine.dart';
import 'package:boncuisine_app/models/recipe.dart';
import 'package:boncuisine_app/models/user.dart';
import 'package:flutter/widgets.dart';

enum Status { INITIAL, SUCCESS, ERROR }

class AppState {
  final List<Recipe> allRecipes;
  final List<Recipe> cuisineRecipes;
  final List<Recipe> userRecipes;
  final User currentUser;
  final List<Cuisine> cuisines;
  final Recipe selectedRecipe;
  final Status loginStatus;

  const AppState({
    @required this.allRecipes,
    @required this.cuisineRecipes,
    @required this.userRecipes,
    @required this.currentUser,
    @required this.cuisines,
    @required this.selectedRecipe,
    @required this.loginStatus,
  });

  AppState.initialState()
      : allRecipes = List.unmodifiable(<Recipe>[]),
        cuisineRecipes = List.unmodifiable(<Recipe>[]),
        userRecipes = List.unmodifiable(<Recipe>[]),
        currentUser = null,
        cuisines = List.unmodifiable(<Cuisine>[]),
        selectedRecipe = null,
        loginStatus = Status.INITIAL;

  AppState copyWith({
    List<Recipe> allRecipes,
    List<Recipe> cuisineRecipes,
    List<Recipe> userRecipes,
    User currentUser,
    List<Cuisine> cuisines,
    Recipe selectedRecipe,
    Status loginStatus,
  }) {
    return AppState(
        allRecipes: allRecipes ?? this.allRecipes,
        cuisineRecipes: cuisineRecipes ?? this.cuisineRecipes,
        userRecipes: userRecipes ?? this.userRecipes,
        currentUser: currentUser ?? this.currentUser,
        cuisines: cuisines ?? this.cuisines,
        selectedRecipe: selectedRecipe ?? this.selectedRecipe,
        loginStatus: loginStatus ?? this.loginStatus);
  }
}
