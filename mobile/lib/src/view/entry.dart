import 'package:flutter/cupertino.dart';
import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:mobile/src/controller.dart';
import 'package:mobile/src/view/gadget/not_valid.dart';

class MainEntry extends StatefulWidget {
  final ScrollController controller;
  final String title;
  MainEntry({Key key, this.controller, this.title}) : super(key: key);
  @override
  MainEntryState createState() => MainEntryState(controller);
}

class MainEntryState extends StateMVC<MainEntry> {
  final ScrollController _scontroller;
  MainEntryCon _mainEntryCon;
  MainEntryState(this._scontroller) : super(MainEntryCon()) {
    this._mainEntryCon = MainEntryCon.con;
    _mainEntryCon.refreshFeature = (len) => refreshFeature(len);
    _mainEntryCon.refreshQuestion = (len) => refreshQuestions(len);
  }
  int _featureLength = 4;
  int _questionLength = 1;
  void refreshFeature(int len) {
    setState(() {
      _featureLength = len;
    });
  }

  void refreshQuestions(int len) {
    setState(() {
      _questionLength = len;
    });
  }

  void refresh() {}

  @override
  Widget build(BuildContext context) {
    return RefreshIndicator(
        onRefresh: _mainEntryCon.refreshAll,
        child: CustomScrollView(
          controller: _scontroller,
          slivers: <Widget>[
            SliverAppBar(
              floating: true,
              snap: true,
              title: Text(widget.title, style: TextStyle(color: Colors.white)),
            ),
            SliverToBoxAdapter(
                child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                  _featureLength>0?
                  Padding(
                    child: Text(
                      "热榜",
                      style: Theme.of(context)
                          .textTheme
                          .headline6
                          .copyWith(fontWeight: FontWeight.bold),
                      textAlign: TextAlign.left,
                    ),
                    padding: const EdgeInsets.fromLTRB(16, 0, 0, 0),
                  ):Container(
              padding: const EdgeInsets.all(15.0),
              alignment: Alignment.center,
              child: SizedBox(
                  width: 80.0,
                  height: 80.0,
                  child: CircularProgressIndicator(strokeWidth: 2.0)),
            ),
                  Divider(),
                  Container(
                      height: 250.0,
                      child: Center(
                          child: ListView.builder(
                        itemBuilder: _mainEntryCon.featuredQuestion,
                        scrollDirection: Axis.horizontal,
                        primary: true,
                        itemCount: _featureLength,
                      ))),
                  Divider(),
                ])),
            _questionLength > 0
                ? new SliverList(
                    delegate: new SliverChildBuilderDelegate(
                        _mainEntryCon.plainQuestion,
                        childCount: _questionLength),
                  )
                : new SliverToBoxAdapter(
                    child: NotValidIndicator(),
                  )
          ],
        ));
    // new SliverFixedExtentList(delegate: new SliverChildBuilderDelegate((BuildContext context,int index)=>_mainEntryCon.plainQuestionPreview(context, index)), itemExtent: 80.0)
    // ],
    // );
    // new SliverFixedExtentList(
    //     delegate: new SliverChildBuilderDelegate(
    //         (BuildContext context, int index) =>
    //             _mainEntryCon.plainQuestionPreview(context, index)),
    //     itemExtent: 80.0)
  }
}
