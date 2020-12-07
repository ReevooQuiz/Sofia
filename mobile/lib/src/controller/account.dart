import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:mobile/src/controller.dart' show ControllerMVC, hostUrl;
import 'package:mobile/src/model/form.dart';
import 'package:mobile/src/model/user.dart';
import 'package:http/http.dart' as http;

class AccountCon extends ControllerMVC {
  static AccountCon _this;
  final _formKey = GlobalKey<FormState>();
  static bool loginState = false;
  String _name;
  String _password;
  User _user;

  factory AccountCon() {
    _this ??= AccountCon._();
    return _this;
  }

  AccountCon._();

  User get user => _user;

  GlobalKey<FormState> get formKey => _formKey;

  Future<User> get future => Future.value(_user);

  static AccountCon get con => _this;

  set name(String name) {
    this._name = name;
  }

  set password(String pass) {
    this._password = pass;
  }

  Widget get login => RaisedButton(
        color: Colors.cyan[600],
        textColor: Colors.white,
        onPressed: () {
          // Validate will return true if the form is valid, or false if
          // the form is invalid.
          if (_formKey.currentState.validate()) {
            _formKey.currentState.save();
            fetchAccount(LoginForm(_name, _password), http.Client())
                .then((value) {
              switch (value.type) {
                case "mismatch":
                  {
                    Scaffold.of(stateMVC.context).showSnackBar(SnackBar(
                      content: Text("用户名或密码错误"),
                    ));
                  }
                  break;
                case "banned":
                  {
                    Scaffold.of(stateMVC.context).showSnackBar(SnackBar(
                      content: Text("你的账号已被禁用"),
                    ));
                  }
                  break;
                default:
                  {
                    Scaffold.of(stateMVC.context).showSnackBar(SnackBar(
                      content: Text("登录成功"),
                    ));
                    loginState = true;
                    Navigator.pop(stateMVC.context);
                  }
              }
            });
          }
        },
        child: Text('登录'),
      );
  Widget get signIn => OutlineButton(
        onPressed: null,
        child: Text('注册'),
      );
  Future<User> fetchAccount(LoginForm form, http.Client client) async {
    final response = await client.post(hostUrl + 'login', body: form.toJson());
    final resJson = jsonDecode(response.body);
    if (resJson['code'] == 0) {
      return User.fromJson(resJson);
    } else
      return User.fromError(resJson['type']);
  }
}
