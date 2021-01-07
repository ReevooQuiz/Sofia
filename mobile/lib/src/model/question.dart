import 'package:mobile/src/model.dart';

class Question {
  final String questionId;
  final User raiser;
  String imageUrl;
  String title;
  String content;
  String category;
  int acceptedAnswer;
  int answerCount;
  int viewCount;
  int favoriteCount;
  String description;
  bool isSelected = false;
  List<String> labels;
  DateTime time;
  Question(this.questionId, this.raiser, this.title,this.description, this.time,
      {this.imageUrl,this.labels});

  factory Question.fromJson(Map<String, dynamic> json) {
    return Question(json['qid'], User.fromUid(json['raiser']), json['title'],json['head']??json['content'],
        DateTime.tryParse(json['time']),
        imageUrl: json['picture_urls'][0],
        labels:json['labels'].cast<String>());
  } 
  Map<String, dynamic> toJson() => {
        'title': this.title,
        'content': this.content,
        'category': this.category,
        'labels': this.labels.toString()
      };
}
