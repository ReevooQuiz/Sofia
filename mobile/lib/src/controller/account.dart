import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';
import 'package:mobile/src/controller.dart' show ControllerMVC, hostUrl;
import 'package:mobile/src/model/form.dart';
import 'package:mobile/src/model/user.dart';
import 'package:http/http.dart' as http;
import 'package:mobile/src/view.dart';
import 'package:simple_auth/simple_auth.dart' as simpleAuth;
import 'package:font_awesome_flutter/font_awesome_flutter.dart';

class AccountCon extends ControllerMVC {
  static AccountCon _this;
  final _loginFormKey = GlobalKey<FormState>();
  final _signInFormKey = GlobalKey<FormState>();
  final _forgetFormKey = GlobalKey<FormState>();
  final _changeFormKey = GlobalKey<FormState>();
  static bool loginState = false;
  bool _codeSent = false;
  String _code;
  User _user;
  final String regexEmail =
      "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*\$";
  final simpleAuth.GithubApi githubApi = new simpleAuth.GithubApi(
      "github",
      "51f0dde36e2f4fcee97c",
      "04aee9d3c62d4ea10577113dedbf62b842f8a855",
      "http://localhost",
      scopes: [
        "user",
        "repo",
        "public_repo",
      ]);
  factory AccountCon() {
    _this ??= AccountCon._();
    return _this;
  }

  AccountCon._();

  User get user => _user;

  GlobalKey<FormState> get formKey => _loginFormKey;
  GlobalKey<FormState> get signInFormKey => _signInFormKey;
  GlobalKey<FormState> get forgetFormKey => _forgetFormKey;
  GlobalKey<FormState> get changeFormKey => _changeFormKey;
  Future<User> get future => Future.value(_user);

  static AccountCon get con => _this;

  set name(String name) {
    _user.name = name;
  }

  set password(String pass) {
    _user.password = pass;
  }

  set email(String email) {
    _user.email = email;
  }

  set code(String code) {
    _code = code;
  }

  set gender(String value) {
    switch (value) {
      case '男':
        _user.gender = 0;
        break;
      case '女':
        _user.gender = 1;
        break;
      default:
        _user.gender = 2;
    }
  }

  set nickName(String name) {
    _user.nickName = name;
  }

  Widget get login => RaisedButton(
        textColor: Colors.white,
        onPressed: () {
          // Validate will return true if the form is valid, or false if
          // the form is invalid.
          if (_loginFormKey.currentState.validate()) {
            _loginFormKey.currentState.save();
            fetchAccount(LoginForm(_user.name, _user.password), http.Client())
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
  Widget get signIn => OutlinedButton(
        onPressed: () {
          _user = User("-1", "<<invalid>>");
          Navigator.push(stateMVC.context,
              MaterialPageRoute(builder: (BuildContext context) => SignIn()));
        },
        child: Text('注册'),
      );

  Widget get submit => RaisedButton(
        textColor: Colors.white,
        onPressed: () {
          if (_signInFormKey.currentState.validate()) {
            _signInFormKey.currentState.save();
            loginState = true;
            Navigator.pop(stateMVC.context);
            Navigator.pop(stateMVC.context);
            Navigator.push(
              stateMVC.context,
              MaterialPageRoute(
                  builder: (BuildContext context) => Home(
                        title: 'Sofia',
                      ),
                  maintainState: false),
            );
          }
        },
        child: Text('提交'),
      );

  Function _changePassword;
  set changePwTrigger(Function trigger) {
    _changePassword = trigger;
  }

  Widget get forgetPassword => FlatButton(
      onPressed: () {
        _user = User("-1", "<<invalid>>");
        Navigator.push(
            stateMVC.context,
            MaterialPageRoute(
                builder: (BuildContext context) => ForgetPassword()));
      },
      child: Text('忘记密码？'));

  Widget get codeVerifier => OutlinedButton(
        onPressed: () {
          if (_forgetFormKey.currentState.validate()) {
            _changePassword();
          }
        },
        child: Text('确认'),
      );

  Widget get changePassword => RaisedButton(
      textColor: Colors.white,
      onPressed: () {
        if (_changeFormKey.currentState.validate()) {
          _changeFormKey.currentState.save();
          Navigator.pop(stateMVC.context);
        }
      },
      child: Text('确认修改'));

  Widget get loginWithGithub => Padding(
      padding: const EdgeInsets.symmetric(horizontal: 80.0),
      child: RaisedButton(
        color: Colors.black,
        onPressed: () async {
          try {
            var success = await githubApi.authenticate();
            Scaffold.of(stateMVC.context).showSnackBar(
                SnackBar(content: Text("Logged in success: $success")));
          } catch (e) {
            print("$e");
          }
        },
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceEvenly,
          children: [
            FaIcon(
              FontAwesomeIcons.github,
              color: Colors.white,
            ),
            Text("使用GitHub登录", style: TextStyle(color: Colors.white))
          ],
        ),
      ));

  Future<User> fetchAccount(LoginForm form, http.Client client) async {
    final response = await client.post(hostUrl + 'login', body: form.toJson());
    final resJson = jsonDecode(response.body);
    if (resJson['code'] == 0) {
      return User.fromJson(resJson);
    } else
      return User.fromError(resJson['type']);
  }

  bool validPassword(String confirmPassword) {
    return _user.password == confirmPassword;
  }

  bool isEmail(String mail) {
    return new RegExp(regexEmail).hasMatch(mail);
  }

  Future<PickedFile> inputImage(imageOpt opt) async {
    if (opt == imageOpt.camera)
      return await ImagePicker().getImage(source: ImageSource.camera);
    else
      return await ImagePicker().getImage(source: ImageSource.gallery);
  }

  Future<bool> verifyEmail(GlobalKey<FormState> formState) async {
    if (formState.currentState.validate()) {
      formState.currentState.save();
      _codeSent = true;
    } else {
      _codeSent = false;
      return Future.value(false);
    }
    return Future.value(true);
  }

  Future<bool> verifyCode(String value) async =>
      value != null ? value == _code : false;

  bool validateCode(String value) {
    if (_codeSent) return value.trim().length != 0;
    return true;
  }
}
