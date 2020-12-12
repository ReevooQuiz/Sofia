class Question {
  final int questionId;
  final String raiser;
  String imageUrl;
  String title;
  String content;
  String category;
  int acceptedAnswer;
  int answerCount;
  int viewCount;
  int favoriteCount;
  bool isSelected = false;
  List<String> labels;
  DateTime time;

  Question(this.questionId, this.raiser, this.title, this.time,
      {this.imageUrl});

  factory Question.fromJson(Map<String, dynamic> json) {
    return Question(json['qid'], json['raiser'], json['title'],
        DateTime.tryParse(json['time']));
  }

  Map<String, dynamic> toJson() => {
        'title': this.title,
        'content': this.content,
        'category': this.category,
        'labels': this.labels.toString()
      };
}
