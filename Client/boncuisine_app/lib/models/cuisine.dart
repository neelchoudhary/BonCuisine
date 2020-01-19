class Cuisine {
  int cuisineID;
  String cuisineName;

  static Cuisine fromJSON(Map json) {
    Cuisine cuisine = Cuisine();
    cuisine.cuisineID = json["id"];
    cuisine.cuisineID = json["name"];
    return cuisine;
  }

  Cuisine({this.cuisineID, this.cuisineName});
}
