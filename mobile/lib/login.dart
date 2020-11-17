import 'dart:convert';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

class Login extends StatefulWidget{
  @override
  State<StatefulWidget> createState()=>LoginState();
}
class LoginState extends State<Login>{

  final _formKey = GlobalKey<FormState>();
  String _name;
  String _password;
  @override
  Widget build(BuildContext context) {
    return Form(
        key: _formKey,
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.center,
          children: <Widget>[
            Padding(
                padding: const EdgeInsets.symmetric(vertical: 15.0),
                child:Icon(
                  Icons.filter_vintage,
                  size: 50,
                )),
            Text(
              '登录',
              style: TextStyle(
                  fontSize: 30,
                  fontWeight: FontWeight.bold
              ),
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
                    http.post('http://192.168.3.40:8070/' + 'Login', body: {
                      'user_name': _name,
                      'password': _password
                    }).then((value) => {
                      if (jsonDecode(value.body)['LoginStatus'] == true)
                        {
                          Scaffold.of(context).showSnackBar(
                              SnackBar(content: Text('登录成功'))),
                          Navigator.pop(context),
                          Navigator.push(
                              context,
                              MaterialPageRoute(
                                  builder: (context) => StateBookTile(
                                      User.fromJson(
                                          jsonDecode(value.body)))))
                        }
                      else
                        {
                          Scaffold.of(context).showSnackBar(
                              SnackBar(content: Text('登录失败'))),
                        }
                    });
                  }
                },
                child: Text('登录'),
              ),
            ),
          ],
        ));
  }
}