import 'package:flutter/cupertino.dart';
import 'package:mobile/src/controller.dart';

class MainEntry extends StatefulWidget {
  @override
  MainEntryState createState() => MainEntryState();
}

class MainEntryState extends StateMVC<MainEntry> {
  MainEntryCon _mainEntryCon;
  MainEntryState() : super(MainEntryCon()) {
    this._mainEntryCon = MainEntryCon.con;
  }
  @override
  Widget build(BuildContext context) {
    return CustomScrollView(
      slivers: <Widget>[
        SliverPadding(
            padding: const EdgeInsets.all(8.0),
            sliver: new ListView(scrollDirection: Axis.horizontal,children: _mainEntryCon.featuredQuestion,)),
        new SliverFixedExtentList(delegate: new SliverChildBuilderDelegate((BuildContext context,int index)=>_mainEntryCon.plainQuestionPreview(context, index)), itemExtent: 80.0)
      ],
    );
  }
}
