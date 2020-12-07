
class Answer {
  final int answerId;
  final int answerer;
  final int questionId;
  String content;
  int commentCount;
  int criticismCount;
  int likeCount;
  int approvalCount;
  DateTime time;

  Answer(
      this.answerId,
      this.answerer,
      this.questionId,
      this.content,
      this.approvalCount,
      this.criticismCount,
      this.commentCount,
      this.likeCount,
      this.time);

  factory Answer.fromJson(Map<String, dynamic> json) {
    return Answer(
        json['aid'],
        json['answerer'],
        json['qid'],
        null,
        json['approval_count'],
        json['criticism_count'],
        json['comment_count'],
        json['like_count'],
        DateTime.tryParse(json['time']));
  }
  Map<String, dynamic> toJson()=>{'qid':this.questionId,'content':this.content};
}
