import '../model.dart';

class Answer {
  final int answerId;
  final User answerer;
  final int questionId;
  String content;
  int commentCount;
  int criticismCount;
  int likeCount;
  int approvalCount;
  bool liked;
  bool approved;
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
      this.liked,
      this.approved,
      this.time);

  factory Answer.fromJson(Map<String, dynamic> json) {
    return Answer(
        json['aid'],
        User.fromUid(json['answerer']),
        json['qid'],
        json['head'] ?? json['content'],
        json['approval_count'],
        json['criticism_count'],
        json['comment_count'],
        json['like_count'],
        json['liked'],
        json['approved'],
        DateTime.tryParse(json['time']));
  }
  Map<String, dynamic> toJson() =>
      {'qid': this.questionId, 'content': this.content};
}
