import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:mobile/src/controller/account.dart';
import 'package:mvc_pattern/mvc_pattern.dart';

class Login extends StatefulWidget {
  @override
  State<StatefulWidget> createState() => LoginState();
}

class LoginState extends StateMVC<Login> {
  AccountCon _accountCon;

  LoginState() : super(AccountCon()) {
    _accountCon = AccountCon.con;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: Text("登录"),
          backgroundColor: Colors.white,
        ),
        body: SingleChildScrollView(
            child: Padding(
                padding: EdgeInsets.symmetric(vertical: 100),
                child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Form(
                          key: _accountCon.formKey,
                          child: Column(
                            crossAxisAlignment: CrossAxisAlignment.center,
                            children: <Widget>[
                              Padding(
                                  padding: const EdgeInsets.symmetric(
                                      vertical: 50.0),
                                  child: Text(
                                    'Reevoo σοφία',
                                    style: TextStyle(
                                        fontSize: 30,
                                        fontWeight: FontWeight.bold),
                                  )),
                              Padding(
                                  padding: const EdgeInsets.symmetric(
                                      vertical: 16.0, horizontal: 50.0),
                                  child: TextFormField(
                                    decoration:
                                        const InputDecoration(hintText: '用户名'),
                                    validator: (value) {
                                      if (value.isEmpty) {
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
                                  decoration:
                                      const InputDecoration(hintText: '密码'),
                                  validator: (value) {
                                    if (value.isEmpty) {
                                      return '请输入密码';
                                    }
                                    return null;
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
                                  child: Row(
                                      mainAxisAlignment:
                                          MainAxisAlignment.spaceEvenly,
                                      children: [
                                        _accountCon.signIn,
                                        _accountCon.login
                                      ])),
                            ],
                          ))
                    ]))));
  }
}
