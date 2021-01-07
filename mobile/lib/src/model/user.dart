enum LoginType { Banned, Mismatch, Inactive, Unknown, Normal }

class User {
  final String id;
  final LoginType type;
  String _name;
  String _nickName;
  String _password;
  String _email;
  String _icon;
  int _gender;
  set name(String name) {
    this._name = name;
  }

  String get name => _name;
  String get password => _password;
  String get nickName => _nickName;
  String get email => _email;
  int get gender => _gender;
  String get icon => _icon;
  set password(String pass) {
    this._password = pass;
  }

  set email(String email) {
    this._email = email;
  }

  set gender(int value) {
    this._gender = value;
  }

  set nickName(String name) {
    this._nickName = name;
  }

  set icon(String icon) {
    this._icon = icon;
  }

  User(this.id, this.type);
  //used when log in
  factory User.fromJson(Map<String, dynamic> json) {
    User user = User(json['uid'].toString(), LoginType.values[json['role']]);
    user.icon = json['icon'];
    user.name = json['name'];
    user.nickName = json['nickname'];
    return user;
  }
  //used in question & answer displays, only contain avatar&name
  factory User.fromUid(Map<String, dynamic> json) {
    User user = User(json['uid'], LoginType.Unknown);
    user.icon = json['icon'];
    user.name = json['name'];
    user.nickName = json['nickname'];
    return user;
  }
  factory User.fromError(int code) {
    return User("<<<invalid>>>", LoginType.values[code]);
  }
}
