import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import { Vue } from 'vue-class-component';

Vue.use(BootstrapVue)
Vue.use(IconsPlugin)

createApp(App)
  .use(router)
  .mount("#app");
