import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:mobile/src/utils.dart';
import 'package:mobile/src/view.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
          primarySwatch: createMaterialColor(Color(0xFF88D5D1)),
          visualDensity: VisualDensity.adaptivePlatformDensity,
          buttonTheme: ButtonThemeData(
              shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.all(Radius.circular(30)))),
          inputDecorationTheme: InputDecorationTheme(
              filled: true,
              fillColor: Color(0x2088D5D2),
              focusedErrorBorder: OutlineInputBorder(
                  borderSide: BorderSide(color: Colors.redAccent),
                  borderRadius: BorderRadius.all(Radius.circular(30))),
              errorBorder: OutlineInputBorder(
                  borderSide: BorderSide(color: Colors.redAccent),
                  borderRadius: BorderRadius.all(Radius.circular(30))),
              enabledBorder: OutlineInputBorder(
                  borderSide: BorderSide(color: Color(0xFF67A29E)),
                  borderRadius: BorderRadius.all(Radius.circular(30))),
              focusedBorder: OutlineInputBorder(
                  borderSide: BorderSide(color: Color(0xFF88D5D1)),
                  borderRadius: BorderRadius.all(Radius.circular(30))))),
      home: Home(title: 'Sofia'),
    );
  }
}
