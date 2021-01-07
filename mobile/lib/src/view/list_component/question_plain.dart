import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';

class QuestionView extends StatefulWidget {
  final String title;
  final String description;
  final String raiser;
  final String avatarUrl;
  final VoidCallback onFavorite;
  final VoidCallback onPressed;
  final VoidCallback onAnswer;
  final VoidCallback onLike;
  QuestionView(
      {Key key,
      @required this.title,
      @required this.description,
      @required this.raiser,
      @required this.avatarUrl,
      @required this.onFavorite,
      @required this.onLike,
      @required this.onPressed,
      @required this.onAnswer})
      : super(key: key);

  @override
  _QuestionViewState createState() => _QuestionViewState(
      title: title,
      description: description,
      raiser: raiser,
      avatarUrl: avatarUrl,
      onLike: onLike,
      onFavorite: onFavorite,
      onPressed: onPressed,
      onAnswer: onAnswer);
}

class _QuestionViewState extends State<QuestionView> {
  final String title;
  final String description;
  final String raiser;
  final String avatarUrl;
  final VoidCallback onLike;

  bool _favorite = false; // 是否被收藏
  final VoidCallback onFavorite;
  final VoidCallback onPressed;
  final VoidCallback onAnswer;
  _QuestionViewState(
      {@required this.title,
      @required this.description,
      @required this.raiser,
      @required this.avatarUrl,
      @required this.onFavorite,
      @required this.onPressed,
      @required this.onLike,
      @required this.onAnswer});

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Card(
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: <Widget>[
            ListTile(
              leading: CircleAvatar(
                backgroundImage: NetworkImage(avatarUrl),
                radius: 30,
              ),
              title: Text(title),
              subtitle: Text(description),
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.end,
              children: <Widget>[
                TextButton(
                  child: _favorite
                      ? const FaIcon(FontAwesomeIcons.solidStar)
                      : const FaIcon(FontAwesomeIcons.star),
                  onPressed: () {
                    setState(() {
                      _favorite = !_favorite;
                    });
                    onFavorite();
                  },
                  style: ButtonStyle(enableFeedback: true),
                ),
                const SizedBox(width: 8),
                TextButton(
                  child: const FaIcon(FontAwesomeIcons.pen),
                  onPressed: () => onAnswer,
                ),
                const SizedBox(width: 8),
                TextButton(
                  child: const FaIcon(FontAwesomeIcons.heart),
                  onPressed: () => onLike,
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
