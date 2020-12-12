<template>
  <div style="background-color:#FFFFFF">
    <a-row type="flex" justify="space-between" align="middle">
      <img alt="Sofia" src="../assets/Sofia.png" height="25" width="65" style="marginLeft:20px" />
      <a-menu v-model:selectedKeys="current" mode="horizontal" inlineIndent="0" align="middle">
        <a-menu-item key="home">
          <router-link to="/">首页</router-link>
        </a-menu-item>
        <a-menu-item key="category">分类</a-menu-item>
        <a-menu-item key="recommend">推荐</a-menu-item>
        <a-menu-item key="explore">探索</a-menu-item>
        <a-menu-item key="ban">封禁</a-menu-item>
        <a-menu-item key="mine">
          <router-link to="/personal">我的</router-link>
        </a-menu-item>
      </a-menu>
      <a-space :size="10" style="marginRight:20px">
        <router-link to="/">
          <a-button type="primary" shape="circle">
            <template #icon>
              <SearchOutlined style="color:'#FFFFFF'" />
            </template>
          </a-button>
        </router-link>
        <a-button type="primary" shape="round" size="small">提问</a-button>
        <a-avatar @click="goToPersonal">
          <template #icon>
            <UserOutlined />
          </template>
        </a-avatar>
        <a-button type="primary" shape="round" size="small" @click="goToLogin">登录</a-button>
      </a-space>
    </a-row>
  </div>
</template>

<script >
import { defineComponent } from "vue";
import { Options, Vue } from "vue-class-component";
import { UserOutlined, SearchOutlined } from "@ant-design/icons-vue";

import { Button } from 'ant-design-vue';
export default {
  components: {
    UserOutlined,
    SearchOutlined,
     'a-button': Button

  },
  data() {
    return {
      current: ["home"]
    };
  },
  methods: {
    goToPersonal() {
      this.$router.push({ path: "/personal" });
    },
    goToLogin() {
      console.log("!");
      this.$router.push({ path: "/login" });
    }
  }
};
</script>

