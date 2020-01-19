import 'package:boncuisine_app/datarepos/cuisine_repository.dart';
import 'package:boncuisine_app/models/cuisine.dart';
import 'package:boncuisine_app/redux/app_state.dart';
import 'package:boncuisine_app/redux/cuisine/cuisine_actions.dart';
import 'package:boncuisine_app/utils/logger.dart';
import 'package:redux/redux.dart';

class CuisineMiddleware extends MiddlewareClass<AppState> {
  final ICuisineRepository _cuisineRepository;

  CuisineMiddleware(this._cuisineRepository);

  @override
  void call(Store<AppState> store, dynamic action, NextDispatcher next) async {
    if (action is LoadCuisinesAction) {
      try {
        List<Cuisine> cuisines = await _cuisineRepository.getCuisines();
        store.dispatch(OnLoadedCuisinesAction(cuisines: cuisines));
        action.completer.complete();
      } catch (error) {
        Logger.e("Failed to get cuisines", e: error, s: StackTrace.current);
        action.completer.completeError(error);
      }
    }

    if (action is TapCuisineSelector) {}

    next(action);
  }
}
