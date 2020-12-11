import 'package:flutter/material.dart';
import 'package:mobile/src/controller.dart';
import 'package:mobile/src/model/question.dart';

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
    return null;
  }

  Widget plainQuestionPreview(BuildContext context, int index) {
    return null;
  }
}
