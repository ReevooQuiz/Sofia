import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:mobile/src/controller/account.dart';
import 'package:mobile/src/view.dart';
import 'package:mvc_pattern/mvc_pattern.dart';
class Login extends StatefulWidget {
  @override
  State<StatefulWidget> createState() => LoginState();
}

class LoginState extends StateMVC<Login> {
  AccountCon _accountCon;
  String _title;

  LoginState() : super(AccountCon()) {
    _accountCon = AccountCon.con;
    _accountCon.forgetTrigger = () => forgetPassword();
    _title = '登录';
  }
  void forgetPassword() {
    setState(() {
      if (_title == '登录')
        _title = '忘记密码';
      else
        _title = '登录';
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: Text(_title),
          leading: _title == '登录'
              ? null
              : IconButton(
                  icon: Icon(Icons.arrow_back),
                  onPressed: () => forgetPassword()),
          backgroundColor: Colors.white,
        ),
        body: SingleChildScrollView(child: _form()));
  }

  Widget _form() {
    if (_title == '登录') {
      return Padding(
          padding: EdgeInsets.symmetric(vertical: 100),
          child: Column(mainAxisAlignment: MainAxisAlignment.center, children: [
            Form(
                key: _accountCon.formKey,
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.center,
                  children: <Widget>[
                    Padding(
                        padding: const EdgeInsets.symmetric(vertical: 50.0),
                        child: Text(
                          'Reevoo σοφία',
                          style: TextStyle(
                              fontSize: 30, fontWeight: FontWeight.bold),
                        )),
                    Padding(
                        padding: const EdgeInsets.symmetric(
                            vertical: 16.0, horizontal: 50.0),
                        child: TextFormField(
                          decoration: const InputDecoration(hintText: '用户名'),
                          validator: (value) {
                            if (value.trim().isEmpty) {
                              return '请输入用户名';
                            }
                            return null;
                          },
                          onSaved: (value) {
                            _accountCon.name = value;
                          },
                          autofocus: true,
                        )),
                    Padding(
                      padding: const EdgeInsets.symmetric(
                          vertical: 16.0, horizontal: 50.0),
                      child: TextFormField(
                        decoration: const InputDecoration(hintText: '密码'),
                        validator: (value) {
                          if (value.trim().isEmpty) {
                            return '请输入密码';
                          }
                          return null;
                        },
                        onChanged: (value) {
                          _accountCon.password = value;
                        },
                        obscureText: true,
                        onSaved: (value) {
                          _accountCon.password = value;
                        },
                      ),
                    ),
                    Padding(
                        padding: const EdgeInsets.symmetric(
                            vertical: 20.0, horizontal: 10.0),
                        child: Column(
                            mainAxisAlignment: MainAxisAlignment.center,
                            children: [
                              _accountCon.loginWithGithub,
                              _accountCon.forgetPassword,
                              Row(
                                  mainAxisAlignment:
                                      MainAxisAlignment.spaceEvenly,
                                  children: [
                                    _accountCon.signIn,
                                    _accountCon.login
                                  ])
                            ])),
                  ],
                ))
          ]));
    } else
      return ForgetPassword();
  }
}
