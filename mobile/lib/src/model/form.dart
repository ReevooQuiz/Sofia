class LoginForm {
  final String name;
  final String password;
  LoginForm(this.name, this.password);
  Map<String, dynamic> toJson() => {"name": name, "password": password};
}

class SignInForm {}
