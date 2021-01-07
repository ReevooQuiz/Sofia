/* eslint-disable */
import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import Antd from 'ant-design-vue';

import VueMarkdownEditor from '@kangc/v-md-editor';//*****************npm i @kangc/v-md-editor@next -S
import '@kangc/v-md-editor/lib/style/base-editor.css';
import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js';
import 'ant-design-vue/dist/antd.less';
import moment from 'moment';
import 'moment/locale/zh-cn';
import { store, key } from './store'


moment.locale('zh-cn');

VueMarkdownEditor.use(vuepressTheme);
process.env.Mock && require('./mock.js')

createApp(App)
  .use(router)
  .use(Antd)
  .use(VueMarkdownEditor)
  .use(store, key)
  .mount("#app");

