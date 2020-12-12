import 'dart:io';

import 'package:http/http.dart' show Client, Request;
import '../controller.dart';
import '../model/user.dart';
import 'dart:convert';

class ApiProvider {
  final Client client;
  static ApiProvider _this;

  factory ApiProvider(Client client) {
    _this ??= ApiProvider._(client);
    return _this;
  }
  ApiProvider._(this.client);
  Future<User> login(dynamic body) async {
    final request = Request('POST', Uri.parse(hostUrl + 'login'));
    request.headers[HttpHeaders.contentTypeHeader] =
        "application/json;charset=utf-8";
    request.body =jsonEncode(body);
    final response = await client.send(request);
    final resJson = jsonDecode(await response.stream.bytesToString());
    if (resJson['code'] == 0) {
      return User.fromJson(resJson);
    } else
      return User.fromError(resJson['type']);
  }
}
