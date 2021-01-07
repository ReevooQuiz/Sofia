import 'package:json_annotation/json_annotation.dart';

part 'form.g.dart';
@JsonSerializable()
class LoginForm {
  final String name;
  final String password;
  LoginForm(this.name, this.password);
  Map<String, String> toJson() => _$LoginFormToJson(this);
}

class SignInForm {}
