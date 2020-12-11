import 'package:flutter/cupertino.dart';
import 'package:mobile/src/controller.dart';

class Me extends StatefulWidget {
  @override
  MeState createState() => MeState();
}

class MeState extends StateMVC<Me> {
  @override
  Widget build(BuildContext context) {
    return Text("我的");
  }
}
