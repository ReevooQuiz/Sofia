import 'dart:async';
import 'dart:io';
import 'package:flutter/material.dart';
import 'package:http/http.dart' show Client, ClientException, Request, Response;
import 'package:retry/retry.dart';
import '../controller.dart';
import '../model.dart';
import '../model/user.dart';
import 'dart:convert';

class ApiProvider {
  final Client client;
  static ApiProvider _this;
  static String _token;
  static String _refreshToken;
  set token(String t) {
    _token = t;
  }

  factory ApiProvider(Client client) {
    _this ??= ApiProvider._(client);
    return _this;
  }
  ApiProvider._(this.client);
  Future<User> login(dynamic body, BuildContext context) async {
    final request = Request('POST', Uri.parse(hostUrl + 'login'));
    request.headers[HttpHeaders.contentTypeHeader] =
        "application/json;charset=utf-8";
    request.body = jsonEncode(body.toString());
    final response = await requestWithRetry(
        (request) => client.send(request), request, 3, context);
    final resJson = jsonDecode(await response.stream.bytesToString());
    if (resJson['code'] == 0) {
      _token = resJson['token'];
      _refreshToken = resJson['refresh_token'];
      return User.fromJson(resJson['result']);
    } else
      throw User.fromError(resJson['type']);
  }

  Future<List<Question>> hotlist(BuildContext context) async {
    final response = await requestWithRetry(
        (url) => client.get(url), hostUrl + 'hotlist', 3, context);
    final resJson = jsonDecode(response.body);
    List<Question> questions = List<Question>();
    if (resJson['code'] == 0) {
      final List<dynamic> lists = resJson['result'];
      lists.forEach((element) {
        questions.add(Question.fromJson(element));
      });
    }
    return questions;
  }

  Future<List<Question>> questions(BuildContext context, int page) async {
    final response = await requestWithRetry(
        (url) => client.get(url),
        hostUrl + 'questions?' + 'category=all&page=' + page.toString(),
        3,
        context);
    final resJson = jsonDecode(response.body);
    List<Question> questions = List<Question>();
    if (resJson['code'] == 0) {
      final List<dynamic> list = resJson['result'];
      list.forEach((element) {
        questions.add(Question.fromJson(element));
      });
    }
    return questions;
  }

  Future<Question> questionDetail(BuildContext context, String id) async {
    final response = await requestWithRetry((url) => client.get(url),
        hostUrl + 'question?' + 'qid=' + id, 3, context);
    final resJson = jsonDecode(response.body);
    return Question.fromJson(resJson);
  }

  Future<List<Answer>> answers(BuildContext context, String id,int page,int sort) async{
    
    final response = await requestWithRetry(
        (url) => client.get(url),
        hostUrl + 'answers?' + 'qid=' +id+'&page='+ page.toString()+'&sort='+sort.toString(),
        3,
        context);
    final resJson = jsonDecode(response.body);
    List<Answer> answers = List<Answer>();
    if (resJson['code'] == 0) {
      final List<dynamic> list = resJson['result'];
      list.forEach((element) {
        answers.add(Answer.fromJson(element));
      });
    }
    return answers;
  } 

  Future<Answer> answerDetail(BuildContext context,String id)async{
    final response = await requestWithRetry((url) => client.get(url),
        hostUrl + 'answer?' + 'aid=' + id, 3, context);
    final resJson = jsonDecode(response.body);
    return Answer.fromJson(resJson);
  }

  //Request the host with retry time of times, if failed finally, show a dialog.
  Future<dynamic> requestWithRetry(Function requestor, dynamic parameter,
      int times, BuildContext context) async {
    final r = RetryOptions(maxAttempts: 5);
    Response response;
    try {
      await r.retry(
          () async => requestor(parameter)
              .then((value) => response = value)
              .timeout(Duration(seconds: 3)),
          retryIf: (e) =>
              e is SocketException ||
              e is TimeoutException ||
              e is ClientException);
    } finally {
      if (response == null)
        showDialog(
            context: context,
            builder: (context) => AlertDialog(
                  content: Text("网络异常"),
                  actions: [
                    TextButton(
                        onPressed: () => Navigator.pop(context),
                        child: Text("确定"))
                  ],
                ));
    }
    return response;
  }
}
