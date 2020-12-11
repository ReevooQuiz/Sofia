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
  String _title = "验证邮箱";
  _ForgetPasswordState() : super(AccountCon()) {
    _accountCon = AccountCon.con;
    _accountCon.changePwTrigger = () => _changePassword();
  }

  void _changePassword() {
    setState(() {
      _change = true;
      _title = "重置密码";
    });
  }

  @override
  Widget build(BuildContext context) {
    if (_change)
      return Scaffold(
          appBar: AppBar(
            leading: IconButton(
                icon: Icon(Icons.arrow_back_outlined, color: Color(0xFF5F6772)),
                onPressed: () {
                  Navigator.pop(context);
                  setState(() {
                    _change = false;
                    _title = "验证邮箱";
                  });
                }),
            title: Text(
              _title,
              style: TextStyle(color: Color(0xFF5F6772)),
            ),
            backgroundColor: Colors.white,
          ),
          body: SingleChildScrollView(child: _changePasswordForm()));
    return Scaffold(
        appBar: AppBar(
          leading: IconButton(
              icon: Icon(Icons.arrow_back_outlined, color: Color(0xFF5F6772)),
              onPressed: () {
                Navigator.pop(context);
                setState(() {
                  _change = false;
                  _title = "验证邮箱";
                });
              }),
          title: Text(
            _title,
            style: TextStyle(color: Color(0xFF5F6772)),
          ),
          backgroundColor: Colors.white,
        ),
        body: SingleChildScrollView(child: _validateEmail()));
  }

  Widget _validateEmail() {
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

  Widget _changePasswordForm() {
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
                  decoration:
                      const InputDecoration(hintText: '密码', labelText: '密码'),
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
                  decoration: const InputDecoration(
                      hintText: '确认密码', labelText: '确认密码'),
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
  }
}
