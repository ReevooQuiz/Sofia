<template>
    <div style="background-color:#FFFFFF">
        <a-row type="flex" justify="space-between" align="middle">
            <img alt="Sofia" src="../assets/Sofia.png" height="25" width="65" style="marginLeft:20px"/>
            <a-menu v-model:selectedKeys="current" mode="horizontal" inlineIndent=0 align="middle">
                <a-menu-item key="home">首页</a-menu-item>
                <a-menu-item key="category">分类</a-menu-item>
                <a-menu-item key="recommend">推荐</a-menu-item>
                <a-menu-item key="explore">探索</a-menu-item>
                <a-menu-item key="ban">封禁</a-menu-item>
                <a-menu-item key="mine">我的</a-menu-item>
            </a-menu>
            <a-space :size="10" style="marginRight:20px">
                <a-button type="primary" shape="circle" >
                    <template #icon>
                        <SearchOutlined style="color:'#FFFFFF'"/>
                    </template>
                </a-button>
                <a-button type="primary" shape="round" size="small">
                    提问
                </a-button>
                <a-avatar>
                    <template #icon>
                        <UserOutlined />
                    </template>
                </a-avatar>
                <a-button type="primary" shape="round" size="small">
                    登录
                </a-button>
            </a-space>
        </a-row>
    </div>
</template>
<script lang="ts">
import { defineComponent } from 'vue';
import { Options, Vue } from "vue-class-component";
import { UserOutlined, SearchOutlined } from '@ant-design/icons-vue';

export default {
  components: {
    UserOutlined,
    SearchOutlined
  },
  data() {
    return {
      current: ['home'],
    };
  },
};
</script>

