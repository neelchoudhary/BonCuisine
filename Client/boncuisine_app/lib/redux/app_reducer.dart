import 'app_state.dart';
import 'package:redux/redux.dart';

import 'cuisine/cuisine_reducer.dart';

final appReducer = combineReducers<AppState>([
  ...cuisineReducers,
]);
