import 'package:flutter/cupertino.dart';
import 'package:mobile/src/controller.dart';

class HotList extends StatefulWidget {
  @override
  HotListState createState() => HotListState();
}

class HotListState extends StateMVC<HotList> {
  @override
  Widget build(BuildContext context) {
    return Text("热榜");
  }
}
