'use strict';
const express = require('express');
const Mock = require('mockjs');
const apiRoutes = express.Router();

let random = Math.random() * 500 + 500;
// 访问 /user/ 时
apiRoutes.get('/', function(req, res) {
  setTimeout(() => {
    res.json({
      status: 1,
      msg: '查询成功',
      data: {
          name: '张三'
      }
    });
  }, random);
});
// 访问 /user/idList 时
apiRoutes.get('/idList', function(req, res) {
  setTimeout(() => {
    res.json({
      status: 1,
      msg: 'OK',
      data: Mock.mock({
            'list|1-10': [{
              'id|+1': 1
            }]
        })
    });
  }, random);
});

module.exports = apiRoutes;
