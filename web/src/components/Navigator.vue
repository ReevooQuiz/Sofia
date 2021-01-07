<template>
  <div style="background-color:#FFFFFF">
    <a-row type="flex" justify="space-between" align="middle">
      <img alt="Sofia" src="../assets/Sofia.png" height="25" width="65" style="marginLeft:20px" />
      <a-menu v-model:selectedKeys="current" mode="horizontal" inlineIndent="0" align="middle">
        <a-menu-item key="home"><router-link to="/">首页</router-link></a-menu-item>
        <a-sub-menu key="category">
          <template #title>
            <span class="submenu-title-wrapper">
              分类
            </span>
          </template>
          <a-menu-item key="study">
            <router-link to="/categoryStudy"><BookOutlined />学习</router-link>
          </a-menu-item>
          <a-menu-item key="life">
            <router-link to="/categoryLife"><CoffeeOutlined />生活</router-link>
          </a-menu-item>
        </a-sub-menu>
        <a-menu-item key="recommend"><router-link to="/recommend">推荐</router-link></a-menu-item>
        <!-- <a-menu-item key="explore">探索</a-menu-item> -->
        <a-menu-item key="ban">
          <router-link v-if="admin" to="/ban">封禁</router-link></a-menu-item>
        <a-menu-item key="mine">
          <router-link to="/personalSet">我的</router-link></a-menu-item>
        <a-menu-item v-if="admin" key="hotRank">
          <router-link to="/hotRank">热榜管理</router-link></a-menu-item>
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
        <a-avatar @click="goToPersonal">
          <img v-if="this.logStatus" src="'data:image/png;base64,'+this.avtar" alt="图片未上传" />
          <template #icon>
            <UserOutlined />
          </template>
        </a-avatar>
        <a-button >{{logButton}}</a-button>
        <a-button
          v-if="this.loginStatus"
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
import { UserOutlined, SearchOutlined , BookOutlined, CoffeeOutlined } from "@ant-design/icons-vue";
import { message } from "ant-design-vue";
import { Button } from "ant-design-vue";
export default {
  components: {
    UserOutlined,
    SearchOutlined,
    "a-button": Button,
    BookOutlined,
    CoffeeOutlined
  },
  data() {
    return {
      current: ["home"],
      logStatus: false,
      avtar: "",
      admin:false
    };
  },
  created() {
    if (sessionStorage.getItem("user") !== null) {
      this.logStatus = true;
      this.avtar = JSON.parse(sessionStorage.getItem("user")).icon;
      if (JSON.parse(sessionStorage.getItem("user")).role==0)
        this.admin=true;
    } else {
    }
  },

  methods: {
    goToPersonal() {
      this.$router.push({ path: "/personalSet" });
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
  computed:{
            // userType:function(){
            //     return this.$store.state.user.type;
            // },
            logButton:function () {
                if(sessionStorage.getItem("user")){
                    return 1;
                }
                else return 0;
            },
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

