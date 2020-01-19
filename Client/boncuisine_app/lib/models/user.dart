class User {
  int userID;
  String username;
  String password;
  String name;
  String email;

  static User fromJSON(Map json) {
    User user = User();
    user.userID = json["id"];
    user.username = json["username"];
    user.password = json["password"];
    user.name = json["name"];
    user.email = json["email"];
    return user;
  }
}
