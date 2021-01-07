class Comment {
  final int commentId;
  final int userId;
  final int answerId;
  String content;
  DateTime time;
  Comment(this.commentId, this.userId, this.answerId, this.time);
  factory Comment.fromJson(Map<String, dynamic> json) {
    return Comment(json['cmid'], json['uid'], json['aid'],
        DateTime.tryParse(json['time']));
  }
  // Map<String, dynamic> toJson()=>{}
}
