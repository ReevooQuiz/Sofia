class User {
  // final String name;
  final String id;
  // final String avatar;
  // final String joinDate;
  final String type;
  User(this.id, this.type);
  factory User.fromJson(Map<String, dynamic> json) {
    return User(json['ID'].toString(), json['type']);
  }
  factory User.fromError(int code) {
    if (code==1) {
      return User("<<<invalid>>>", "mismatch");
    } else
      return User("<<<invalid>>>", "banned");
  }
}
