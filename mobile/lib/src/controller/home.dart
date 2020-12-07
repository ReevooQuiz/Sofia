import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:mobile/src/controller.dart';

class HomeCon extends ControllerMVC {
  static HomeCon _this;

  factory HomeCon() {
    _this ??= HomeCon._();
    return _this;
  }

  HomeCon._();

  static HomeCon get con => _this;

  BottomNavigationBarItem get hot =>
      BottomNavigationBarItem(icon: Icon(Icons.home), label: "趋势");
  BottomNavigationBarItem get category =>
      BottomNavigationBarItem(icon: Icon(Icons.category), label: "分类");
  BottomNavigationBarItem get timeline =>
      BottomNavigationBarItem(icon: Icon(Icons.menu), label: "关注");
  BottomNavigationBarItem get me =>
      BottomNavigationBarItem(icon: Icon(Icons.person_rounded), label: "我的");
}
