import 'package:flutter/cupertino.dart';
import 'package:mobile/src/controller.dart';

class TimeLine extends StatefulWidget {
  @override
  TimeLineState createState() => TimeLineState();
}

class TimeLineState extends StateMVC<TimeLine> {
  @override
  Widget build(BuildContext context) {
    return SliverToBoxAdapter(child:Text("关注"));
  }
}
