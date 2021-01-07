import 'package:flutter/cupertino.dart';
import 'package:mobile/src/controller.dart';

class Category extends StatefulWidget {
  @override
  CategoryState createState() => CategoryState();
}

class CategoryState extends StateMVC<Category> {
  @override
  Widget build(BuildContext context) {
    return Text("分类");
  }
}
