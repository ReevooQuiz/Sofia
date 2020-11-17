import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import Antd from 'ant-design-vue';

import 'ant-design-vue/dist/antd.less';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap';
import $ from 'jquery';


createApp(App)
  .use(router)
  .use(Antd) 
  .mount("#app");