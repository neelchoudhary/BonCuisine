import 'package:boncuisine_app/redux/cuisine/cuisine_actions.dart';
import 'package:redux/redux.dart';
import '../app_state.dart';

final cuisineReducers = <AppState Function(AppState, dynamic)>[
  TypedReducer<AppState, OnLoadedCuisinesAction>(_onLoadedCuisines),
];

AppState _onLoadedCuisines(AppState state, OnLoadedCuisinesAction action) {
  return state.copyWith(cuisines: action.cuisines);
}
