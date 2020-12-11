'use strict';

const express = require('express');
const app = express();
// port
let NODE_PORT = process.env.PORT || 4000;
// 监听 /user
app.use('/user', function(req, res) {
  // 让接口 500-1000ms 返回 好让页面有个loading
  setTimeout(() => {
    res.json({
      status: 1,
      msg: '查询成功',
      data: {
          name: '张三'
      }
    });
  }, Math.random() * 500 + 500);
});

app.listen(NODE_PORT, function() {
  console.log('mock服务在' + NODE_PORT + '端口上已启用！');
});