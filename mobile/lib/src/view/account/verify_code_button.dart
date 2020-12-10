import 'dart:async';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:mobile/src/controller.dart';

class VerifyCodeButton extends StatefulWidget {
  @override
  State<StatefulWidget> createState() => _EmailVaildater();
}

class _EmailVaildater extends StateMVC<VerifyCodeButton> {
  AccountCon _accountCon;
  bool _sendState;
  int _countDown;
  _EmailVaildater() : super(AccountCon()) {
    _accountCon = AccountCon.con;
    _countDown = 60;
    _sendState = false;
  }
  @override
  Widget build(BuildContext context) {
    if (_sendState == false) {
      return OutlineButton(onPressed: _startTimer, child: Text('发送验证码'));
    } else {
      return OutlineButton(
          onPressed: null, child: Text(_countDown.toString() + '秒后重新发送'));
    }
  }

  void _startTimer() async {
    const oneSec = const Duration(seconds: 1);
    _sendState = true;
    if (await _accountCon.verifyEmail()) {
      Scaffold.of(context)
          .showSnackBar(SnackBar(content: Text('已向邮箱发送验证码，请注意查收。')));
      new Timer.periodic(
        oneSec,
        (Timer timer) {
          if (_countDown == 0) {
            setState(() {
              timer.cancel();
              _countDown = 60;
              _sendState = false;
            });
          } else {
            setState(() {
              _countDown--;
            });
          }
        },
      );
    }
  }
}
