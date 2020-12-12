import 'package:flutter/material.dart';
import 'package:mobile/src/controller.dart';
import 'package:mobile/src/model.dart';
import 'package:mobile/src/view.dart';

class MainEntryCon extends ControllerMVC {
  static MainEntryCon _this;
  List<Question> questions;
  factory MainEntryCon() {
    _this ??= MainEntryCon._();
    return _this;
  }
  MainEntryCon._();
  static MainEntryCon get con => _this;
  List<Widget> get featuredQuestion {
    List<QuestionPreviewCard> list = List<QuestionPreviewCard>();
    questions.forEach((question) {
      int index = list.length;
      list.add(QuestionPreviewCard(
          title: question.title,
          description: question.content,
          isSelected: question.isSelected,
          onSelected: () => selectQuestion(index),
          onPressed: () => enterQuestion(index),
          imageUrl: question.imageUrl));
    });
    return list;
  }

  void enterQuestion(int index) {}

  void selectQuestion(int index) {}
  Widget plainQuestionPreview(BuildContext context, int index) {
    Question question = questions[index];
    //NOTE : should replace with avatar
    // return QuestionView(avatarUrl: question.imageUrl,onPressed: ()=>enterQuestion(index),);
    return null;
  }
}
