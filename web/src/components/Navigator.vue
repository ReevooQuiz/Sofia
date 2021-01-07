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
        <a-menu-item  v-if="admin" key="ban">
          <router-link to="/ban">封禁</router-link></a-menu-item>
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
        <a-avatar @click="goToPersonal" v-if="logStatus" :src="avatar" alt="图片未上传" >
        </a-avatar>
        <a-avatar v-else>
          <template #icon>
            <UserOutlined />
          </template>
        </a-avatar>
        <a-button
          v-if="logStatus"
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
      current: ["home"]
    };
  },
  created() {
    if (sessionStorage.getItem("user") !== null) {
      this.$store.commit("changeLogStatus",true);
      if (JSON.parse(sessionStorage.getItem("user")).role==0)
        this.$store.commit("changeAdmin",true);
    } else {
    }
    // const store = useStore();
    // store.commit('increment');
    // console.log(store.state.count);

  },

  methods: {
    goToPersonal() {
      this.$router.push({ path: "/personalSet" });
    },
    goToLogin() {
      this.$router.push({ path: "/login" });
    },
    goToLogout() {
      if (sessionStorage.getItem("user") !== null) {
        sessionStorage.removeItem("user");
      }
      this.$store.commit("changeLogStatus",false);
      this.$store.commit("changeIcon","");
        this.$store.commit("changeAdmin",false);
      message.success("成功登出");
      this.$router.push({ path: "/" });
    }
  },
  computed: {
    current:function () {
      return [this.$store.state.navTarget];
    },
    logStatus:function (){
      return this.$store.state.logStatus;
    },
    avatar:function () {
      return this.$store.state.navIcon;
    },
    admin:function(){
        return this.$store.state.admin;
    }
  }
};
</script>

