import 'package:flutter/material.dart';
import 'package:mobile/src/controller.dart';
import 'package:mobile/src/model.dart';
import 'package:mobile/src/view.dart';
import 'package:mobile/src/resources/api_provider.dart';
import 'package:http/http.dart' as http;

class MainEntryCon extends ControllerMVC {
  static MainEntryCon _this;
  List<Question> hotlist = List<Question>();
  List<Question> questions = List<Question>();
  int _currentPage = 0;
  int _questionLen = 0;
  int _pageQuestionLen = 5;
  int _featureLen = 0;
  int _maxFeatureLen = 5;
  factory MainEntryCon() {
    _this ??= MainEntryCon._();
    return _this;
  }
  MainEntryCon._();
  Function refreshFeature;
  Function refreshQuestion;
  static MainEntryCon get con => _this;
  Function get featuredQuestion => (BuildContext context, int index) {
        if (index >= _featureLen) {
          if (index < _maxFeatureLen) {
            _getFeatureQuestions();
            return Container(
              padding: const EdgeInsets.all(15.0),
              alignment: Alignment.center,
              child: SizedBox(
                  width: 80.0,
                  height: 80.0,
                  child: CircularProgressIndicator(strokeWidth: 2.0)),
            );
          }
          return Container(
              alignment: Alignment.center,
              padding: EdgeInsets.all(16.0),
              child: Text(
                "没有更多了",
                style: TextStyle(color: Colors.grey),
              ));
        } else
          return QuestionPreviewCard(
              title: hotlist[index].title,
              description: hotlist[index].description,
              isSelected: hotlist[index].isSelected,
              onSelected: () => selectQuestion(index),
              onPressed: () => enterQuestion(hotlist[index].questionId),
              imageUrl: hotlist[index].imageUrl);
      };

  Future<void> refreshAll() async {
    refreshFeature(0);
    refreshQuestion(1);
    questions.clear();
    hotlist.clear();
    _getAllQuestions(0);
    await _getFeatureQuestions();
  }

  Function get plainQuestion => (BuildContext context, int index) {
        if (index >= questions.length) {
          _getAllQuestions(++_currentPage);
          if (index >= questions.length) return null;
          return Container(
            padding: const EdgeInsets.all(15.0),
            alignment: Alignment.center,
            child: SizedBox(
                width: 80.0,
                height: 80.0,
                child: CircularProgressIndicator(strokeWidth: 2.0)),
          );
        } else
          return QuestionView(
            title: questions[index].title,
            description: questions[index].description,
            onPressed: () => enterQuestion(questions[index].questionId),
            avatarUrl: questions[index].raiser.icon,
            onFavorite: () => favoriteQuestion(questions[index].questionId),
            onAnswer: () => answerQuestion(questions[index].questionId),
            raiser: questions[index].raiser.id,
          );
      };

  Future<void> _getFeatureQuestions() async {
    final client = http.Client();
    ApiProvider(client).hotlist(stateMVC.context).then((value) {
      hotlist = value.toList();
      _featureLen = hotlist.length;
      _maxFeatureLen = _featureLen;
      refreshFeature(_featureLen + 1);
    });
  }

  Future<void> _getAllQuestions(int page) async {
    final client = http.Client();
    await ApiProvider(client).questions(stateMVC.context, page).then((value) {
      questions += value.toList();
      if (value.length == 0) _currentPage--;
      _questionLen = questions.length;
      refreshQuestion(_questionLen);
    });
  }

  void enterQuestion(String id) {}

  void favoriteQuestion(String id) {}

  void answerQuestion(String id) {}

  void selectQuestion(int index) {
    hotlist[index].isSelected = !hotlist[index].isSelected;
    refreshFeature(_featureLen + 1);
  }

  Widget plainQuestionPreview(BuildContext context, int index) {
    Question question = hotlist[index];
    //NOTE : should replace with avatar
    // return QuestionView(avatarUrl: question.imageUrl,onPressed: ()=>enterQuestion(index),);
    return null;
  }
}
