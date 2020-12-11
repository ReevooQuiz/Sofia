import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:mobile/src/utils.dart';
import 'package:mobile/src/view.dart';
import 'package:simple_auth_flutter/simple_auth_flutter.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    SimpleAuthFlutter.init(context);

    MaterialColor colors = createMaterialColor(Color(0xFF88D5D1));
    const BorderRadius globalBorderRadius =
        BorderRadius.all(Radius.circular(30));
    return MaterialApp(
      title: 'Sofia',
      theme: ThemeData(
          primarySwatch: colors,
          iconTheme: IconThemeData(color: Color(0xFF5F6772)),
          textTheme: TextTheme(
              button: TextStyle(color: Color(0xFF5F6772)),
              bodyText2: TextStyle(color: Color(0xFF5F6772))),
          colorScheme: ColorScheme.fromSwatch(
              primarySwatch: colors, accentColor: colors.shade600),
          visualDensity: VisualDensity.adaptivePlatformDensity,
          hintColor: Color(0xFF4A6C80),
          outlinedButtonTheme: OutlinedButtonThemeData(
              style: OutlinedButton.styleFrom(
                primary: Color(0xFF4A6C80),
                shape: RoundedRectangleBorder(borderRadius: globalBorderRadius),
                backgroundColor: Colors.white,
                shadowColor: Colors.black,
                minimumSize: Size(85,40),
                elevation: 1.0
              )),
          buttonTheme: ButtonThemeData(
              shape: RoundedRectangleBorder(borderRadius: globalBorderRadius)),
          inputDecorationTheme: InputDecorationTheme(
              filled: true,
              contentPadding: const EdgeInsets.symmetric(vertical: 10.0, horizontal: 15.0),
              fillColor: Color(0x3088D5D2),
              focusedErrorBorder: OutlineInputBorder(borderSide: BorderSide(color: Colors.redAccent), borderRadius: globalBorderRadius),
              errorBorder: OutlineInputBorder(borderSide: BorderSide(color: Colors.redAccent), borderRadius: globalBorderRadius),
              enabledBorder: OutlineInputBorder(borderSide: BorderSide(color: Color(0xFF67A29E)), borderRadius: globalBorderRadius),
              focusedBorder: OutlineInputBorder(borderSide: BorderSide(color: Color(0xFF88D5D1)), borderRadius: globalBorderRadius))),
      home: Login(),
    );
  }
}
