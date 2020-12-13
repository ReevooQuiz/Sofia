<template>
  <div style="background-color:#FFFFFF">
    <a-row type="flex" justify="space-between" align="middle">
      <img alt="Sofia" src="../assets/Sofia.png" height="25" width="65" style="marginLeft:20px" />
      <a-menu v-model:selectedKeys="current" mode="horizontal" inlineIndent="0" align="middle">

                <a-menu-item key="home"><router-link to="/">首页</router-link></a-menu-item>
                <a-menu-item key="category">分类</a-menu-item>
              <a-menu-item key="recommend"><router-link to="/recommend">推荐</router-link></a-menu-item>
        <a-menu-item key="explore">探索</a-menu-item>
        <a-menu-item key="ban">
          <router-link to="/ban">封禁</router-link></a-menu-item>
        <a-menu-item key="mine">
          <router-link to="/personal">我的</router-link></a-menu-item>
            </a-menu>
            <a-space :size="10" style="marginRight:20px">
                <router-link to="/">
                    <a-button type="primary" shape="circle" >
                        <template #icon>
                            <SearchOutlined style="color:'#FFFFFF'"/>
                        </template>
                    </a-button>
                </router-link>
                <a-button type="primary" shape="round" size="small">
                  <router-link to="/postQuestion">提问</router-link>
                </a-button>
                <a-avatar@click="goToPersonal">
          <img v-if="this.logStatus" src="'data:image/png;base64,'+this.avtar" alt="图片未上传" />
          <template #icon>
            <UserOutlined />
          </template>
        </a-avatar>
        <a-button
          v-if="this.logStatus"
          type="primary"
          shape="round"
          size="small"
          @click="goToLogout"
        >登出</a-button>
        <a-button v-else type="primary" shape="round" size="small" @click="goToLogin">登录</a-button>
      </a-space>
    </a-row>
  </div>
</template>

<script >
import { defineComponent } from "vue";
import { Options, Vue } from "vue-class-component";
import { UserOutlined, SearchOutlined } from "@ant-design/icons-vue";
import { message } from "ant-design-vue";
import { Button } from "ant-design-vue";
export default {
  components: {
    UserOutlined,
    SearchOutlined,
    "a-button": Button
  },
  data() {
    return {
      current: ["home"],
      logStatus: false,
      avtar: ""
    };
  },
  created() {
    if (sessionStorage.getItem("user") !== null) {
      this.logStatus = true;
      this.avtar = JSON.parse(sessionStorage.getItem("user")).icon;
    } else {
    }
  },

  methods: {
    goToPersonal() {
      this.$router.push({ path: "/personal" });
    },
    goToLogin() {
      console.log("!");
      this.$router.push({ path: "/login" });
    },
    goToLogout() {
      console.log("!");
      if (sessionStorage.getItem("user") !== null) {
        sessionStorage.removeItem("user");
      }
      this.logStatus = false;
      // this.$store.state.loginStatus=false;
      message.success("成功登出");
      this.$router.push({ path: "/" });
    }
  },
  // computed: {
  //   user: function() {
  //     return this.$store.state.loginStatus;
  //   }
  // },
  // watch: {
  //   user: function(newType) {
  //     if (newType) {
  //       this.logStatus = true;
  //       this.avtar = JSON.parse(newType).icon;
  //     } else {
  //       this.logStatus = true;
  //     }
  //   }
  // }
};
</script>

