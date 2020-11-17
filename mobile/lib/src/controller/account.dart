import 'dart:convert';

import 'package:mobile/src/controller.dart' show ControllerMVC, hostUrl;
import 'package:mobile/src/model/form.dart';
import 'package:mobile/src/model/user.dart';
import 'package:http/http.dart' as http;

class AccountCon extends ControllerMVC {
  static AccountCon _this;
  factory AccountCon() {
    _this ??= AccountCon._();
    return _this;
  }
  AccountCon._();
  User get user => _user;
  User _user;
  Future<User> get future => Future.value(_user);

  Future<User> fetchAccount(LoginForm form) async {
    final response = await http.post(hostUrl + 'login', body: form.toJson());
    final resJson = jsonDecode(response.body);
    if (resJson['code'] == 0) {
      return User.fromJson(resJson);
    } else
      return User.fromError(resJson['type']);
  }
}
