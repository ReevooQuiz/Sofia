import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:mobile/src/controller.dart';
import 'package:mobile/src/view.dart';
class EmailValidater extends StatefulWidget {
  EmailValidater(this._formKey,{Key key}) : super(key: key);
  final GlobalKey<FormState> _formKey;

  @override
  _EmailValidaterState createState() => _EmailValidaterState(_formKey);
}

class _EmailValidaterState extends StateMVC<EmailValidater> {
  AccountCon _accountCon;
  final GlobalKey<FormState> _formKey;
  _EmailValidaterState(this._formKey) : super(AccountCon()) {
    this._accountCon = AccountCon.con;
  }
  @override
  Widget build(BuildContext context) {
    return Container(
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.baseline,
        textBaseline: TextBaseline.ideographic,
        children: <Widget>[
          Padding(
            padding:
                const EdgeInsets.symmetric(vertical: 16.0, horizontal: 50.0),
            child: TextFormField(
              decoration: const InputDecoration(hintText: '邮箱'),
              validator: (value) {
                if (value.trim().isEmpty) return '请输入邮箱';
                if (_accountCon.isEmail(value) == false)
                  return '请输入正确格式的电子邮箱地址';
                return null;
              },
              onSaved: (value) {
                _accountCon.email = value;
              },
              textInputAction: TextInputAction.next,
            ),
          ),
          Padding(
              padding:
                  const EdgeInsets.symmetric(vertical: 16.0, horizontal: 50.0),
              child: 
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    Expanded(
                        flex: 3,
                        child: TextFormField(
                          decoration: const InputDecoration(hintText: '验证码'),
                          onSaved: (value) {
                            _accountCon.code = value;
                          },
                          validator: (value) =>
                              _accountCon.validateCode(value) ? null : '输入验证码',
                          textInputAction: TextInputAction.done,
                        )),
                    Expanded(
                        flex: 5,
                        child: Padding(
                          child: VerifyCodeButton(_formKey),
                          padding: const EdgeInsets.symmetric(horizontal: 10.0),
                        ))
                  ],
                ),
              ),
        ],
      ),
    );
  }
}
