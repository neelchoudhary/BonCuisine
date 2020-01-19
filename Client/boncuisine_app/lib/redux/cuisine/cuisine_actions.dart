import 'dart:async';

import 'package:boncuisine_app/models/cuisine.dart';

class LoadCuisinesAction {
  final Completer completer;
  LoadCuisinesAction({this.completer});
}

class OnLoadedCuisinesAction {
  final List<Cuisine> cuisines;

  OnLoadedCuisinesAction({this.cuisines});
}

class TapCuisineSelector {
  final int cuisineID;

  TapCuisineSelector({this.cuisineID});
}
