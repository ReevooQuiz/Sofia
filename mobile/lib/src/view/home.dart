import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
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
  bool _isVisible;
  ScrollController _hideButtonController;
  _HomeState() : super(HomeCon()) {
    _index = 0;
    _homeCon = HomeCon.con;
    _isVisible = true;
  }
  @override
  void initState() {
    super.initState();
    _isVisible = true;
    _hideButtonController = new ScrollController();
    _hideButtonController.addListener(() {
      print("listener");
      if (_hideButtonController.position.userScrollDirection ==
          ScrollDirection.reverse) {
        setState(() {
          _isVisible = false;
          print("**** $_isVisible up");
        });
      }
      if (_hideButtonController.position.userScrollDirection ==
          ScrollDirection.forward) {
        setState(() {
          _isVisible = true;
          print("**** $_isVisible down");
        });
      }
    });
  }

  void _onItemTapped(int index) {
    setState(() {
      _index = index;
    });
  }

  @override
  void dispose() {
    _hideButtonController.dispose();
    _hideButtonController.removeListener(() {});
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      bottomNavigationBar: AnimatedContainer(
          duration: Duration(milliseconds: 200),
          height: _isVisible ? 60 : 0.0,
          child: BottomNavigationBar(
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
          )),
      body: SafeArea(
        child: Column(
          children: <Widget>[
            AnimatedContainer(
              height: _isVisible ? 56.0 : 0.0,
              duration: Duration(milliseconds: 200),
              child: AppBar(
                title: Text(
                  widget.title,
                  style: TextStyle(color: Colors.white),
                ),
              ),
            ),
            _show(),
          ],
        ),
      ),
    );
  }

  Widget _show() {
    Widget display;

    switch (_index) {
      case 0:
        display = MainEntry();
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
