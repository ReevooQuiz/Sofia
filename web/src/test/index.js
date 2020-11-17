/* eslint-disable */
"use strict";
const express = require("express");
const Mock = require("mockjs");
const apiRoutes = express.Router();

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
                code: 200,
                qid: 234,
                owner: {
                    user_id:1,
                    user_name:"嗷7777",
                    user_icon:""
                },
                title:"上海交通大学",
                content:"当事学生发帖称，其在点评网站上发布差评，商家还曾找人假冒外卖员套取其个人信息。9月14日，花店老板进入学校找其要求删除差评，当事学生随即向交大派出所民警和保卫处求助。经过协商后双方达成协议，商家不再骚扰当事学生同时退还买花的钱，当事学生删除相关差评和视频。该学生称，在协商过程中，花店老板对其进行污蔑和指责，在要签署调解协议书时，花店老板又拒绝退款。",
                answer_count:4,
                follow_count:234,
                view_count:123,
                
                answer_list:[{
                    aid: 234,
                    owner: {
                        user_id:1,
                        user_name:"zhc",
                        user_icon:""
                    },
                    like_count:2,
                    criticism_count:4,
                    approval_count:2,
                    comment_count:2,
                    content:"当事学生发帖称，其在点评网站上发布差评，商家还曾找人假冒外卖员套取其个人信息。9月14日，花店老板进入学校找其要求删除差评，当事学生随即向交大派出所民警和保卫处求助。经过协商后双方达成协议，商家不再骚扰当事学生同时退还买花的钱，当事学生删除相关差评和视频。该学生称，在协商过程中，花店老板对其进行污蔑和指责，在要签署调解协议书时，花店老板又拒绝退款。"
                },
                {
                    aid: 34,
                    owner: {
                        user_id:1,
                        user_name:"zhc",
                        user_icon:""
                    },
                    like_count:2,
                    criticism_count:4,
                    approval_count:2,
                    comment_count:2,
                    content:"当事学生发帖称，其在点评网站上发布差评，商家还曾找人假冒外卖员套取其个人信息。9月14日，花店老板进入学校找其要求删除差评，当事学生随即向交大派出所民警和保卫处求助。经过协商后双方达成协议，商家不再骚扰当事学生同时退还买花的钱，当事学生删除相关差评和视频。该学生称，在协商过程中，花店老板对其进行污蔑和指责，在要签署调解协议书时，花店老板又拒绝退款。"
                },
                {
                    aid: 235,
                    owner: {
                        user_id:1,
                        user_name:"zhc",
                        user_icon:""
                    },
                    like_count:2,
                    criticism_count:4,
                    approval_count:2,
                    comment_count:2,
                    content:"当事学生发帖称，其在点评网站上发布差评，商家还曾找人假冒外卖员套取其个人信息。9月14日，花店老板进入学校找其要求删除差评，当事学生随即向交大派出所民警和保卫处求助。经过协商后双方达成协议，商家不再骚扰当事学生同时退还买花的钱，当事学生删除相关差评和视频。该学生称，在协商过程中，花店老板对其进行污蔑和指责，在要签署调解协议书时，花店老板又拒绝退款。"
                },
                {
                    aid: 24,
                    owner: {
                        user_id:1,
                        user_name:"zhc",
                        user_icon:""
                    },
                    like_count:2,
                    criticism_count:4,
                    approval_count:2,
                    comment_count:2,
                    content:"当事学生发帖称，其在点评网站上发布差评，商家还曾找人假冒外卖员套取其个人信息。9月14日，花店老板进入学校找其要求删除差评，当事学生随即向交大派出所民警和保卫处求助。经过协商后双方达成协议，商家不再骚扰当事学生同时退还买花的钱，当事学生删除相关差评和视频。该学生称，在协商过程中，花店老板对其进行污蔑和指责，在要签署调解协议书时，花店老板又拒绝退款。"
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
            code: 200,
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
                        user_name:"violedo",
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
                        user_name:"violedo",
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
                        user_name:"violedo",
                        user_icon: "https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                    },
                    title: "如何看待上海交通大学花店事件",
                    description: "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打",
                    answer_count: 4,
                    follow_count: 234,
                }
            ],
            cardInfo:{
                kid:123,
                title:"上海交通大学",
                attr:[{
                    name:"db",
                    value:"sad",
                    origin:1234
                }]
            }
            
        })
    );

    setTimeout(() => {
        res.json(jsonResponse);
    }, 200);
});
module.exports = apiRoutes;