import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class QuestionPreviewCard extends StatelessWidget {
  final String title;
  final String description;
  final String imageUrl;
  final VoidCallback onPressed;
  const QuestionPreviewCard(
      {Key key,
      @required this.title,
      @required this.description,
      @required this.isSelected,
      @required this.onSelected,
      @required this.onPressed,
      @required this.imageUrl})
      : super(key: key);
  final bool isSelected;
  final VoidCallback onSelected;
  @override
  Widget build(BuildContext context) {
    final ColorScheme colorScheme = Theme.of(context).colorScheme;
    const ShapeBorder border = ContinuousRectangleBorder(
        borderRadius: BorderRadius.all(Radius.circular(10)));
    return SafeArea(
      top: false,
        child: Padding(
      padding: const EdgeInsets.fromLTRB(8,0,8,8),
      child: Card(
        clipBehavior: Clip.antiAlias,
        shape: border,
        child: InkWell(
          onLongPress: () => onSelected(),
          onTap: ()=>onPressed(),
          splashColor: colorScheme.onSurface.withOpacity(0.12),
          highlightColor: Colors.transparent,
          child: Stack(
            alignment: Alignment.bottomLeft,
            children: [
              Container(
                  color: isSelected
                      ? colorScheme.primary.withOpacity(0.08)
                      : Colors.transparent),
              Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  imageUrl != null
                      ? SizedBox(
                          height: 150,
                          width: 300,
                          child: Stack(
                            children: [
                              Positioned.fill(
                                  child: Ink.image(
                                image: NetworkImage(imageUrl),
                                fit: BoxFit.cover,
                                child: Container(),
                              )),
                              Positioned(
                                  bottom: 16,
                                  left: 16,
                                  right: 16,
                                  child: FittedBox(
                                    fit: BoxFit.scaleDown,
                                    alignment: Alignment.centerLeft,
                                    child: Text(
                                      title,
                                      style: Theme.of(context)
                                          .textTheme
                                          .headline5
                                          .copyWith(color: Colors.white),
                                    ),
                                  )),
                              
                            ]
                          ),
                        )
                      : Text(
                          title,
                          style: Theme.of(context)
                              .textTheme
                              .headline5
                              .copyWith(color: Colors.white),
                        ),Padding(
                                padding:
                                    const EdgeInsets.fromLTRB(16, 16, 16, 0),
                                child: DefaultTextStyle(
                                  softWrap: false,
                                  overflow: TextOverflow.ellipsis,
                                  style: Theme.of(context).textTheme.subtitle1,
                                  child: Column(
                                    crossAxisAlignment:
                                        CrossAxisAlignment.start,
                                    children: [Text(description)],
                                  ),
                                ),
                              )
                ],
              ),
              Align(
                alignment: Alignment.topRight,
                child: Padding(
                  padding: const EdgeInsets.all(8),
                  child: Icon(
                    Icons.check_circle_outline,
                    color:
                        isSelected ? colorScheme.primary : Colors.transparent,
                  ),
                ),
              )
            ],
          ),
        ),
      ),
    ));
  }
}
