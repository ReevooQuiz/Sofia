import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:mobile/src/controller/account.dart';
import 'package:mobile/src/model/form.dart';
import 'package:mvc_application/view.dart';

class Login extends StatefulWidget {
  @override
  State<StatefulWidget> createState() => LoginState();
}

class LoginState extends StateMVC<Login> {
  final _formKey = GlobalKey<FormState>();
  String _name;
  String _password;
  AccountCon _accountCon;
  LoginState() : super(AccountCon()) {
    _accountCon = controller;
  }
  @override
  Widget build(BuildContext context) {
    return Form(
        key: _formKey,
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.center,
          children: <Widget>[
            Padding(
                padding: const EdgeInsets.symmetric(vertical: 15.0),
                child: Icon(
                  Icons.filter_vintage,
                  size: 50,
                )),
            Text(
              '登录',
              style: TextStyle(fontSize: 30, fontWeight: FontWeight.bold),
            ),
            Padding(
                padding: const EdgeInsets.symmetric(
                    vertical: 16.0, horizontal: 30.0),
                child: TextFormField(
                  decoration: const InputDecoration(
                      icon: Icon(Icons.account_circle), hintText: '用户名'),
                  validator: (value) {
                    if (value.isEmpty) {
                      return '请输入用户名';
                    }
                    return null;
                  },
                  onSaved: (value) {
                    _name = value;
                  },
                )),
            Padding(
              padding:
                  const EdgeInsets.symmetric(vertical: 16.0, horizontal: 30.0),
              child: TextFormField(
                decoration: const InputDecoration(
                    icon: Icon(Icons.lock), hintText: '密码'),
                validator: (value) {
                  if (value.isEmpty) {
                    return '请输入密码';
                  }
                  return null;
                },
                obscureText: true,
                onSaved: (value) {
                  _password = value;
                },
              ),
            ),
            Padding(
              padding:
                  const EdgeInsets.symmetric(vertical: 20.0, horizontal: 10.0),
              child: RaisedButton(
                color: Colors.cyan[600],
                textColor: Colors.white,
                onPressed: () {
                  // Validate will return true if the form is valid, or false if
                  // the form is invalid.

                  if (_formKey.currentState.validate()) {
                    _formKey.currentState.save();
                    _accountCon.fetchAccount(LoginForm(_name, _password)) ;
                  }
                },
                child: Text('登录'),
              ),
            ),
          ],
        ));
  }
}
