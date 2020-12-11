import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';

class QuestionView extends StatelessWidget {
  final String title;
  final String description;
  final String raiser;
  final String avatarUrl;
  final bool favorite; // 是否被收藏
  final VoidCallback onFavorite;
  final VoidCallback onPressed;
  final VoidCallback onAnswer;
  QuestionView(
      {Key key,
      @required this.title,
      @required this.description,
      @required this.raiser,
      @required this.avatarUrl,
      @required this.onFavorite,
      @required this.onPressed,
      @required this.onAnswer,
      @required this.favorite})
      : super(key: key);

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
                  child: favorite
                      ? const FaIcon(FontAwesomeIcons.solidStar)
                      : const FaIcon(FontAwesomeIcons.star),
                  onPressed: () => onFavorite(),
                ),
                const SizedBox(width: 8),
                TextButton(
                  child: const FaIcon(FontAwesomeIcons.pen),
                  onPressed: () => onAnswer,
                ),
                const SizedBox(width: 8),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
