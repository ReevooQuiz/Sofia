import 'dart:io';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:mobile/src/controller.dart';
import 'package:mvc_pattern/mvc_pattern.dart';
import 'package:mobile/src/view.dart';
enum imageOpt { gallery, camera }

class SignIn extends StatefulWidget {
  @override
  State<StatefulWidget> createState() => SignInState();
}

class SignInState extends StateMVC<SignIn> {
  AccountCon _accountCon;
  File _imgPath;
  SignInState() : super(AccountCon()) {
    _accountCon = AccountCon.con;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: Text("注册"),
          leading: IconButton(
              icon: Icon(Icons.arrow_back_outlined),
              onPressed: () => Navigator.pop(context)),
          backgroundColor: Colors.white,
        ),
        body: SingleChildScrollView(
            child: Padding(
                padding: EdgeInsets.symmetric(vertical: 10),
                child: Column(children: [
                  Form(
                      key: _accountCon.signInFormKey,
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.center,
                        children: <Widget>[
                          Padding(
                            padding: const EdgeInsets.symmetric(
                                vertical: 16.0, horizontal: 50.0),
                            child: Stack(
                              alignment: Alignment.bottomRight,
                              children: <Widget>[
                                Padding(
                                    padding: const EdgeInsets.all(20.0),
                                    child: _imageView(_imgPath)),
                                PopupMenuButton<imageOpt>(
                                  icon: Icon(Icons.camera_alt_outlined),
                                  offset: Offset.fromDirection(1.07, 60),
                                  itemBuilder: (BuildContext context) =>
                                      <PopupMenuEntry<imageOpt>>[
                                    const PopupMenuItem<imageOpt>(
                                      value: imageOpt.camera,
                                      child: Text('拍照'),
                                    ),
                                    const PopupMenuItem<imageOpt>(
                                      value: imageOpt.gallery,
                                      child: Text('从相册选取'),
                                    )
                                  ],
                                  onSelected: (imageOpt result) {
                                    _accountCon
                                        .inputImage(result)
                                        .then((value) => setState(() {
                                              _imgPath = File(value.path);
                                            }));
                                  },
                                )
                              ],
                            ),
                          ),
                          Padding(
                              padding: const EdgeInsets.symmetric(
                                  vertical: 16.0, horizontal: 50.0),
                              child: TextFormField(
                                decoration:
                                    const InputDecoration(hintText: '用户名'),
                                validator: (value) {
                                  if (value.trim().isEmpty) {
                                    return '请输入用户名';
                                  }
                                  return null;
                                },
                                onSaved: (value) {
                                  _accountCon.name = value;
                                },
                                textInputAction: TextInputAction.next,
                              )),
                          Padding(
                              padding: const EdgeInsets.symmetric(
                                  vertical: 16.0, horizontal: 50.0),
                              child: TextFormField(
                                decoration:
                                    const InputDecoration(hintText: '昵称'),
                                validator: (value) {
                                  if (value.trim().isEmpty) {
                                    return '请输入昵称';
                                  }
                                  return null;
                                },
                                onSaved: (value) {
                                  _accountCon.nickName = value;
                                },
                                textInputAction: TextInputAction.next,
                              )),
                          Padding(
                              padding: const EdgeInsets.symmetric(
                                  vertical: 16.0, horizontal: 50.0),
                              child: DropdownButtonFormField(
                                hint: Text('性别'),
                                items: <String>[
                                  '男',
                                  '女',
                                  '其它'
                                ].map<DropdownMenuItem<String>>((String value) {
                                  return DropdownMenuItem<String>(
                                      value: value, child: Text(value));
                                }).toList(),
                                onChanged: (value) {
                                  _accountCon.gender = value;
                                },
                                validator: (value) =>
                                    value == null ? '请选择性别' : null,
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
                                if (value.trim().length < 6)
                                  return '请输入至少6位的密码';
                                return null;
                              },
                              obscureText: true,
                              onSaved: (value) {
                                _accountCon.password = value;
                              },
                              onChanged: (value) {
                                _accountCon.password = value;
                              },
                              textInputAction: TextInputAction.next,
                            ),
                          ),
                          Padding(
                            padding: const EdgeInsets.symmetric(
                                vertical: 16.0, horizontal: 50.0),
                            child: TextFormField(
                              decoration:
                                  const InputDecoration(hintText: '确认密码'),
                              validator: (value) {
                                if (value.trim().isEmpty) return '请再次输入密码';
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
                          EmailValidater(_accountCon.signInFormKey),
                          Padding(
                              padding: const EdgeInsets.symmetric(
                                  vertical: 20.0, horizontal: 10.0),
                              child: Row(
                                  mainAxisAlignment:
                                      MainAxisAlignment.spaceEvenly,
                                  children: [_accountCon.submit])),
                        ],
                      ))
                ]))));
  }

  Widget _imageView(File imgPath) {
    if (imgPath == null)
      return Center(
          child: Text(
        "上传头像",
        style: Theme.of(context).textTheme.bodyText2,
      ));
    else
      return CircleAvatar(
        radius: 80,
        backgroundImage: FileImage(imgPath),
      );
  }
}
