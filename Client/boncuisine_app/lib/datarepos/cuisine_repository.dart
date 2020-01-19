import 'package:boncuisine_app/models/cuisine.dart';
import 'package:boncuisine_app/utils/constants.dart';
import 'package:boncuisine_app/utils/logger.dart';
import 'package:http/http.dart' as http;
import 'dart:convert' as convert;

abstract class ICuisineRepository {
  Future<List<Cuisine>> getCuisines();
}

class CuisineAPIRepository implements ICuisineRepository {
  Future<List<Cuisine>> getCuisines() async {
    await Future.delayed(Duration(milliseconds: Constants.delay));
    final http.Response response = await http
        .get('${Constants.apiURI}/cuisines', headers: Constants.header);
    Logger.network(response.statusCode, response.body);
    Iterable decodedCuisinesJSON = convert.jsonDecode(response.body);
    List<Cuisine> cuisines =
        decodedCuisinesJSON.map((json) => Cuisine.fromJSON(json)).toList();
    return cuisines;
  }
}

class CuisineTestRepository implements ICuisineRepository {
  Future<List<Cuisine>> getCuisines() async {
    await Future.delayed(Duration(milliseconds: Constants.delay));
    List<Cuisine> cuisines = [];
    cuisines.add(Cuisine(cuisineID: 0, cuisineName: "Asian"));
    cuisines.add(Cuisine(cuisineID: 1, cuisineName: "Italian"));
    cuisines.add(Cuisine(cuisineID: 2, cuisineName: "Contemporary"));
    return cuisines;
  }
}
