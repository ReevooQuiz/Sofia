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
                icon: "https://img.pc841.com/2018/0516/20180516050738880.jpg",
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
apiRoutes.get("/oauth/github", function (req, res) {
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
                icon: "https://img.pc841.com/2018/0516/20180516050738880.jpg",
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

apiRoutes.get("/publicInfo", function (req, res) {
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
            result: {
                username: "sxs",
                nickname: "ao",
                profile: "a brief profile of me",
                icon: "https://img.pc841.com/2018/0516/20180516050738880.jpg",
                level: 2,
                gender: 0,
                email: "test@sjtu.edu.cn",
                account_type: 1,
                label: "math",
                like_count: 10,
                question_count: 10,
                answer_count: 10,
                follower_count: 10,
                followed_count: 10
            }

        })
    );

    setTimeout(() => {
        res.json(jsonResponse);
    }, 200);
});

apiRoutes.get("/notifications", function (req, res) {
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
            result: {
                list: [
                    {
                        type: 0,
                        time: "2015-08-05T08:40:51.620Z",
                        qid: "657x575c7576",
                        title: "What is a bird?",
                        new_answer_count: 2
                    },

                    {
                        type: 1,
                        time: "2015-08-05T08:40:51.620Z",
                        qid: "384cb234cb",
                        question_title: "What is a board?",
                        aid: "234b2v34",
                        answer_head: "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打\
                        近日，有消息称上海一花店老板因差评进",
                        new_like_count: 2
                    },
                    {
                        type: 2,
                        time: "2015-08-05T08:40:51.620Z",
                        qid: "384cb234cb",
                        question_title: "What is a board?",
                        aid: "234b2v34",
                        answer_head: "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打\
                        近日，有消息称上海一花店老板因差评进",
                        new_comment_count: 2
                    },

                    {
                        type: 3,
                        time: "2015-08-05T08:40:51.620Z",
                        qid: "384cb234cb",
                        question_title: "What is a board?",
                        aid: "234b2v34",
                        answer_head: "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打\
                        近日，有消息称上海一花店老板因差评进",
                        new_criticism_count: 2
                    },
                    { type: 4, time: "2015-08-05T08:40:51.620Z", new_follower_count: 3 },
                    {
                        type: 0,
                        time: "2015-08-05T08:40:51.620Z",
                        qid: "657x575c7576",
                        title: "What is a bird?",
                        new_answer_count: 2
                    },
                    {
                        type: 3,
                        time: "2015-08-05T08:40:51.620Z",
                        qid: "384cb234cb",
                        question_title: "What is a board?",
                        aid: "234b2v34",
                        answer_head: "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打\
                        近日，有消息称上海一花店老板因差评进",
                        new_criticism_count: 2
                    },
                    { type: 4, time: "2015-08-05T08:40:51.620Z", new_follower_count: 3 },
                    {
                        type: 0,
                        time: "2015-08-05T08:40:51.620Z",
                        qid: "657x575c7576",
                        title: "What is a bird?",
                        new_answer_count: 2
                    },
                ]
            }

        })
    );

    setTimeout(() => {
        res.json(jsonResponse);
    }, 200);
});


apiRoutes.get("/userAnswers", function (req, res) {
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


                    question: {
                        qid: "234",
                        title: "Favourite programming language?",
                        category: "study",
                        labels: ["programming"],
                        head: "What if we put"
                    },
                    answer: {
                        aid: "234",
                        like_count: 2,
                        criticism_count: 4,
                        approval_count: 2,
                        comment_count: 2,
                        head: "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打\
近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前",
                        time: "2015-08-05T08:40:51.620Z",
                        pictureUrls: [
                            "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png",
                            "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"
                        ],
                        liked: false,
                        approved: true,
                        approvable: true
                    }
                },
            ]

        })
    );

    setTimeout(() => {
        res.json(jsonResponse);
    }, 200);
});
apiRoutes.get("/userQuestions", function (req, res) {
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
            result: [{
                qid: 1,

                title: "如何看待上海交通大学花店事件",
                time: "2015-08-05T08:40:51.620Z",
                answer_count: 5,
                view_count: 10,
                favorite_count: 20,
                category: "study",
                labels: ["programming", "campus"],
                pictureUrls: [
                    "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png",
                    "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"
                ],
                head:
                    "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打\
      近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打",



            },
            {
                qid: 1,

                title: "如何看待上海交通大学花店事件",
                time: "2015-08-05T08:40:51.620Z",
                answer_count: 5,
                view_count: 10,
                favorite_count: 20,
                category: "study",
                labels: ["programming", "campus"],
                pictureUrls: [
                    "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png",
                    "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"
                ],
                head:
                    "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打\
      近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打",



            },
            {
                qid: 1,

                title: "如何看待上海交通大学花店事件",
                time: "2015-08-05T08:40:51.620Z",
                answer_count: 5,
                view_count: 10,
                favorite_count: 20,
                category: "study",
                labels: ["programming", "campus"],
                pictureUrls: [
                    "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png",
                    "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"
                ],
                head:
                    "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打\
      近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打",



            }]
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
