/* eslint-disable */
import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import Antd from 'ant-design-vue';

// import store from '@store/index.js'
import 'ant-design-vue/dist/antd.less';
// import Mock from './mock.js';
// import './mock.js'
import moment from 'moment';
import 'moment/locale/zh-cn';
moment.locale('zh-cn');

process.env.Mock && require('./mock.js')

createApp(App)
  .use(router)
  .use(Antd) 
  .mount("#app");