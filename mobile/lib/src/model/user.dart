class User {
  final String id;
  final String type;
  String _name;
  String _nickName;
  String _password;
  String _email;
  int _gender;

  set name(String name) {
    this._name = name;
  }

  String get name => _name;
  String get password => _password;
  String get nickName => _nickName;
  String get email => _email;
  int get gender => _gender;
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

  User(this.id, this.type);
  factory User.fromJson(Map<String, dynamic> json) {
    return User(json['ID'].toString(), json['type']);
  }
  factory User.fromError(int code) {
    if (code == 1) {
      return User("<<<invalid>>>", "mismatch");
    } else
      return User("<<<invalid>>>", "banned");
  }
}
