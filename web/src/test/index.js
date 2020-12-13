/* eslint-disable */
"use strict";
const express = require("express");
const Mock = require("mockjs");
const apiRoutes = express.Router();

// 启用mock : node test/test
apiRoutes.get("/question", function (req, res) {
    let jsonResponse = {};
    res.vary(
        "Origin",
        "Access-Control-Request-Headers",
        "Access-Control-Request-Method"
    );

    Object.assign(
        jsonResponse,
        Mock.mock({
            // data: {
            status: 200,
            qid: 234,
            owner: {
                user_id: 1,
                user_name: "嗷7777",
                user_icon: ""
            },
            title: "上海交通大学",
            content: "当事学生发帖称，其在点评网站上发布差评，商家还曾找人假冒外卖员套取其个人信息。9月14日，花店老板进入学校找其要求删除差评，当事学生随即向交大派出所民警和保卫处求助。经过协商后双方达成协议，商家不再骚扰当事学生同时退还买花的钱，当事学生删除相关差评和视频。该学生称，在协商过程中，花店老板对其进行污蔑和指责，在要签署调解协议书时，花店老板又拒绝退款。",
            answer_count: 4,
            follow_count: 234,
            view_count: 123,

            answer_list: [{
                aid: 234,
                owner: {
                    user_id: 1,
                    user_name: "zhc",
                    user_icon: ""
                },
                like_count: 2,
                criticism_count: 4,
                approval_count: 2,
                comment_count: 2,
                content: "当事学生发帖称，其在点评网站上发布差评，商家还曾找人假冒外卖员套取其个人信息。9月14日，花店老板进入学校找其要求删除差评，当事学生随即向交大派出所民警和保卫处求助。经过协商后双方达成协议，商家不再骚扰当事学生同时退还买花的钱，当事学生删除相关差评和视频。该学生称，在协商过程中，花店老板对其进行污蔑和指责，在要签署调解协议书时，花店老板又拒绝退款。"
            },
            {
                aid: 34,
                owner: {
                    user_id: 1,
                    user_name: "zhc",
                    user_icon: ""
                },
                like_count: 2,
                criticism_count: 4,
                approval_count: 2,
                comment_count: 2,
                content: "当事学生发帖称，其在点评网站上发布差评，商家还曾找人假冒外卖员套取其个人信息。9月14日，花店老板进入学校找其要求删除差评，当事学生随即向交大派出所民警和保卫处求助。经过协商后双方达成协议，商家不再骚扰当事学生同时退还买花的钱，当事学生删除相关差评和视频。该学生称，在协商过程中，花店老板对其进行污蔑和指责，在要签署调解协议书时，花店老板又拒绝退款。"
            },
            {
                aid: 235,
                owner: {
                    user_id: 1,
                    user_name: "zhc",
                    user_icon: ""
                },
                like_count: 2,
                criticism_count: 4,
                approval_count: 2,
                comment_count: 2,
                content: "当事学生发帖称，其在点评网站上发布差评，商家还曾找人假冒外卖员套取其个人信息。9月14日，花店老板进入学校找其要求删除差评，当事学生随即向交大派出所民警和保卫处求助。经过协商后双方达成协议，商家不再骚扰当事学生同时退还买花的钱，当事学生删除相关差评和视频。该学生称，在协商过程中，花店老板对其进行污蔑和指责，在要签署调解协议书时，花店老板又拒绝退款。"
            },
            {
                aid: 24,
                owner: {
                    user_id: 1,
                    user_name: "zhc",
                    user_icon: ""
                },
                like_count: 2,
                criticism_count: 4,
                approval_count: 2,
                comment_count: 2,
                content: "当事学生发帖称，其在点评网站上发布差评，商家还曾找人假冒外卖员套取其个人信息。9月14日，花店老板进入学校找其要求删除差评，当事学生随即向交大派出所民警和保卫处求助。经过协商后双方达成协议，商家不再骚扰当事学生同时退还买花的钱，当事学生删除相关差评和视频。该学生称，在协商过程中，花店老板对其进行污蔑和指责，在要签署调解协议书时，花店老板又拒绝退款。"
            }]
        })
    );

    setTimeout(() => {
        res.json(jsonResponse);
    }, 200);
});
apiRoutes.post("/register", function (req, res) {
    let jsonResponse = {};
    res.vary(
        "Origin",
        "Access-Control-Request-Headers",
        "Access-Control-Request-Method"
    );

    Object.assign(
        jsonResponse,
        Mock.mock({
            // data: {
            code: 2,
            status: 200,
        })
    );

    setTimeout(() => {
        res.json(jsonResponse);
    }, 200);
});
apiRoutes.post("/login", function (req, res) {
    let jsonResponse = {};
    res.vary(
        "Origin",
        "Access-Control-Request-Headers",
        "Access-Control-Request-Method"
    );

    Object.assign(
        jsonResponse,
        Mock.mock({
            // data: {
            status: 200,
            code: 0,
            result: {
                role: 0,
                uid: "hdfuighduhhhdfhu",
                icon: "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",
                name: "sk",
                nickname: "sk",
                token: "eyJhbGciOiJIUzI1NiJ9.eyJpZCI6MiwidXNlcm5hbWUiOiJ0ZXN0XzAwMDAwMSIsInJvbGUiOjAsImlzUmVmcmVzaCI6ZmFsc2UsInN1YiI6InRlc3RfMDAwMDAxIiwiaXNzIjoidXNlciIsImlhdCI6MTU5NjAwOTA4NCwiZXhwIjo5MjIzMzcyMDM2ODU0Nzc1fQ.PJWiCD-9cBvKdU2qFcjcabyNRCcZXT6B5pa9vDfPDvg",
                refresh_token: "eyJhbGciOiJIUzI1NiJ9.eyJpZCI6MiwidXNlcm5hbWUiOiJ0ZXN0XzAwMDAwMSIsInJvbGUiOjAsImlzUmVmcmVzaCI6ZmFsc2UsInN1YiI6InRlc3RfMDAwMDAxIiwiaXNzIjoidXNlciIsImlhdCI6MTU5NjAwOTA4NCwiZXhwIjo5MjIzMzcyMDM2ODU0Nzc1fQ.PJWiCD-9cBvKdU2qFcjcabyNRCcZXT6B5pa9vDfPDvg"
            }
        })
    );

    setTimeout(() => {
        res.json(jsonResponse);
    }, 200);
});

apiRoutes.post("/publicInfo", function (req, res) {
    let jsonResponse = {};
    res.vary(
        "Origin",
        "Access-Control-Request-Headers",
        "Access-Control-Request-Method"
    );

    Object.assign(
        jsonResponse,
        Mock.mock({
            // data: {
            status: 200,
            code: 0,
           
        })
    );

    setTimeout(() => {
        res.json(jsonResponse);
    }, 200);
});
apiRoutes.get("/followers", function (req, res) {
    let jsonResponse = {};
    res.vary(
        "Origin",
        "Access-Control-Request-Headers",
        "Access-Control-Request-Method"
    );

    Object.assign(
        jsonResponse,
        Mock.mock({
            // data: {
            code: 0,
            result: [
                {
                    icon: "https://tse2-mm.cn.bing.net/th/id/OIP.OCLuKoXlay8WIeNZPpCfcgHaHa?pid=Api&rs=1",
                    name: "akang",
                    nickname: "nickname of akang",
                    profile: "个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介"
                },
                {
                    icon: "https://tse2-mm.cn.bing.net/th/id/OIP.OCLuKoXlay8WIeNZPpCfcgHaHa?pid=Api&rs=1",
                    name: "akang",
                    nickname: "nickname of akang",
                    profile: "个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介"
                },
                {
                    icon: "https://tse2-mm.cn.bing.net/th/id/OIP.OCLuKoXlay8WIeNZPpCfcgHaHa?pid=Api&rs=1",
                    name: "akang",
                    nickname: "nickname of akang",
                    profile: "个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介"
                },
                {
                    icon: "https://tse2-mm.cn.bing.net/th/id/OIP.OCLuKoXlay8WIeNZPpCfcgHaHa?pid=Api&rs=1",
                    name: "akang",
                    nickname: "nickname of akang",
                    profile: "个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介"
                },
                {
                    icon: "https://tse2-mm.cn.bing.net/th/id/OIP.OCLuKoXlay8WIeNZPpCfcgHaHa?pid=Api&rs=1",
                    name: "akang",
                    nickname: "nickname of akang",
                    profile: "个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介"
                },
                {
                    icon: "https://tse2-mm.cn.bing.net/th/id/OIP.OCLuKoXlay8WIeNZPpCfcgHaHa?pid=Api&rs=1",
                    name: "akang",
                    nickname: "nickname of akang",
                    profile: "个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介"
                }]
        })
    );

    setTimeout(() => {
        res.json(jsonResponse);
    }, 200);
});
apiRoutes.get("/search", function (req, res) {
    let jsonResponse = {};
    res.vary(
        "Origin",
        "Access-Control-Request-Headers",
        "Access-Control-Request-Method"
    );

    Object.assign(
        jsonResponse,
        Mock.mock({
            // data: {
            status: 200,
            question_list: [
                {
                    qid: 1,
                    owner: {
                        user_id: 1,
                        user_name: "阿钪",
                        user_icon: "https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                    },
                    title: "如何看待上海交通大学花店事件",
                    description: "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打",
                    answer_count: 4,

                    follow_count: 234,

                },
                {
                    qid: 345,
                    owner: {
                        user_id: 3,
                        user_name: "violedo",
                        user_icon: "https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                    },
                    title: "如何看待上海交通大学花店事件",
                    description: "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打",
                    answer_count: 4,
                    follow_count: 234,
                },
                {
                    qid: 345,
                    owner: {
                        user_id: 3,
                        user_name: "violedo",
                        user_icon: "https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                    },
                    title: "如何看待上海交通大学花店事件",
                    description: "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打",
                    answer_count: 4,
                    follow_count: 234,
                },
                {
                    qid: 345,
                    owner: {
                        user_id: 3,
                        user_name: "violedo",
                        user_icon: "https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                    },
                    title: "如何看待上海交通大学花店事件",
                    description: "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打",
                    answer_count: 4,
                    follow_count: 234,
                }
            ],
            cardInfo: {
                kid: 123,
                title: "上海交通大学",
                attr: [{
                    name: "db",
                    value: "sad",
                    origin: 1234
                }]
            }

        })
    );

    setTimeout(() => {
        res.json(jsonResponse);
    }, 200);
});
module.exports = apiRoutes;
