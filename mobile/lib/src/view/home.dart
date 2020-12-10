import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:mobile/src/controller.dart';
import 'package:mobile/src/view.dart';
import 'package:mvc_pattern/mvc_pattern.dart';

class Home extends StatefulWidget {
  final String title;
  Home({Key key, this.title}) : super(key: key);

  @override
  State<StatefulWidget> createState() => _HomeState();
}

class _HomeState extends StateMVC<Home> {
  HomeCon _homeCon;
  int _index;
  _HomeState() : super(HomeCon()) {
    _index = 0;
    _homeCon = HomeCon.con;
  }

  void _onItemTapped(int index) {
    setState(() {
      _index = index;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title,style: TextStyle(color: Colors.white),),
      ),
      bottomNavigationBar: BottomNavigationBar(
        items: <BottomNavigationBarItem>[
          _homeCon.hot,
          _homeCon.category,
          _homeCon.timeline,
          _homeCon.me
        ],
        currentIndex: _index,
        selectedItemColor: Theme.of(context).primaryColor,
        unselectedItemColor: Colors.grey[400],
        onTap: _onItemTapped,
        showUnselectedLabels: false,
      ),
      body: _show(),
    );
  }

  Widget _show() {
    Widget display;

    switch (_index) {
      case 0:
        display = HotList();
        break;
      case 1:
        display = Category();
        break;
      case 2:
        display = TimeLine();
        break;
      case 3:
        display = Me();
        break;
    }
    return display;
  }
}
