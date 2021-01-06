<template>
 
    
    <div id="background">
     <!-- <a-config-provider :locale="locale"> -->
        <Navigator></Navigator>
        <br />
        <router-view />
      <!-- </a-config-provider> -->
    </div>
  
</template>

<script >
import { Options, Vue } from "vue-class-component";
import Navigator from "@/components/Navigator.vue";

import zh_CN from "ant-design-vue/lib/locale-provider/zh_CN";
import moment from "moment";
import "moment/locale/zh-cn";

moment.locale("zh-cn");
export default {
  components: {
    Navigator
  },
  data() {
    return {
       locale: zh_CN
    };
  },
  

};
</script>


<style>
#background {
  min-height: 683px;
  background-color: #edeeed;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /* text-align: center; */
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
