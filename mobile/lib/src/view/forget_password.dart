import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:mobile/src/view.dart';
import 'package:mvc_pattern/mvc_pattern.dart';
import 'package:mobile/src/controller.dart';

class ForgetPassword extends StatefulWidget {
  @override
  _ForgetPasswordState createState() => _ForgetPasswordState();
}

class _ForgetPasswordState extends StateMVC<ForgetPassword> {
  AccountCon _accountCon;
  bool _change = false;
  _ForgetPasswordState() : super(AccountCon()) {
    _accountCon = AccountCon.con;
    _accountCon.changePwTrigger = () => _changePassword();
  }

  void _changePassword() {
    setState(() {
      _change = true;
    });
  }

  @override
  Widget build(BuildContext context) {
    if (_change)
      return Padding(
          padding: const EdgeInsets.symmetric(vertical: 100),
          child: Form(
            key: _accountCon.changeFormKey,
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Padding(
                  padding: const EdgeInsets.symmetric(
                      vertical: 16.0, horizontal: 50.0),
                  child: TextFormField(
                    decoration: const InputDecoration(hintText: '密码'),
                    validator: (value) {
                      if (value.isEmpty) {
                        return '请输入密码';
                      }
                      if (value.trim().length < 6) return '请输入至少6位的密码';
                      return null;
                    },
                    obscureText: true,
                    onChanged: (value) {
                      _accountCon.password = value;
                    },
                    onSaved: (value) {
                      _accountCon.password = value;
                    },
                    textInputAction: TextInputAction.next,
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.symmetric(
                      vertical: 16.0, horizontal: 50.0),
                  child: TextFormField(
                    decoration: const InputDecoration(hintText: '确认密码'),
                    validator: (value) {
                      if (value.isEmpty) return '请再次输入密码';
                      if (_accountCon.validPassword(value) == false)
                        return '两次输入密码不一致';
                      return null;
                    },
                    obscureText: true,
                    onSaved: (value) {
                      _accountCon.password = value;
                    },
                    textInputAction: TextInputAction.next,
                  ),
                ),
                _accountCon.changePassword
              ],
            ),
          ));
    return Padding(
        padding: const EdgeInsets.symmetric(vertical: 100),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Form(
              key: _accountCon.forgetFormKey,
              child: Column(children: [
                EmailValidater(_accountCon.forgetFormKey),
                Padding(
                  padding: const EdgeInsets.symmetric(vertical: 50.0),
                  child: _accountCon.codeVerifier,
                )
              ]),
            )
          ],
        ));
  }
}
