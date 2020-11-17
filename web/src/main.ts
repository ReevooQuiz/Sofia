/* eslint-disable */
import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import Antd from 'ant-design-vue';

import 'ant-design-vue/dist/antd.less';
// import Mock from './mock.js';
// import './mock.js'

process.env.Mock && require('./mock.js')

createApp(App)
  .use(router)
  .use(Antd) 
  .mount("#app");