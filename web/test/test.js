/* eslint-disable */
"use strict";
const express = require("express");
const cors = require("cors");
const app = express();
const testSever = require("../src/test");
app.use(cors({
    origin: "*",
    methods: ["GET", "POST", "PUT", "DELETE"],
    allowedHeaders: ["content-Type", "Authorization"]
}));
// port
let NODE_PORT = process.env.PORT || 4000;
// 监听 /realTimeStat

app.use(express.static("public"));
app.use("/", testSever);
app.listen(NODE_PORT, function() {
    console.log("mock服务在" + NODE_PORT + "端口上已启用！");
});
