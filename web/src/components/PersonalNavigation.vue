<template>
  <div style="width: 256px">
    <a-button type="primary" @click="toggleCollapsed" style="margin-bottom: 16px">
      <MenuUnfoldOutlined v-if="collapsed" />
      <MenuFoldOutlined v-else />
    </a-button>
    <a-menu
      :default-selected-keys="['1']"
      mode="inline"
      :inline-collapsed="collapsed"
    >
     <a-menu-item key="1"  >
       <HeartOutlined />
        
        <span><router-link to="/personalFollowing"> 关注</router-link></span>
        
      
      </a-menu-item>
      <a-menu-item key="2">
        <SmileOutlined />
       
        <span>   <router-link to="/personalFollower"> 粉丝 </router-link></span>
        
      </a-menu-item>
      <a-menu-item key="3">
        <PicLeftOutlined />
        <span> <router-link to="/personalQuestion"> 问题 </router-link></span>
      
      </a-menu-item>
      <a-menu-item key="4">
        <BulbOutlined />
        <span>  <router-link to="/personalAnswer"> 回答 </router-link></span>
      
      </a-menu-item>
      <a-menu-item key="5">
        <MailOutlined />
        <span>  <router-link to="/personalMessage"> 消息 </router-link></span>
      
      </a-menu-item>
      <a-menu-item key="6">
        <SettingOutlined />
        <span>  <router-link to="/personalSet"> 设置 </router-link></span>
      
      </a-menu-item>
     
    </a-menu>
  </div>
</template>

<script>
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  PieChartOutlined,
  MailOutlined,
  DesktopOutlined,
  InboxOutlined,
  AppstoreOutlined,
  HeartOutlined,
  SmileOutlined,
  BulbOutlined,
  PicLeftOutlined,
  SettingOutlined
 
} from '@ant-design/icons-vue';
import { Menu } from 'ant-design-vue';

// you can rewrite it to a single file component, if not, you should config vue alias to vue/dist/vue.esm-bundler.js
const SubMenu = {
  components: {
    PieChartOutlined,
    MailOutlined,
    
  },
 
  name: 'SubMenu',
  props: {
    menuInfo: {
      type: Object,
      default: () => ({}),
    },
  },
};
export default {
  components: {
    'sub-menu': SubMenu,
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    PieChartOutlined,
    MailOutlined,
    DesktopOutlined,
    InboxOutlined,
    AppstoreOutlined,
   HeartOutlined,
   SmileOutlined,
   BulbOutlined,
   PicLeftOutlined,
   SettingOutlined
    

  },
  data() {
    return {
      collapsed:true,
    };
  },
  methods: {
    toggleCollapsed() {
      this.collapsed = !this.collapsed;
    },
  },
};
</script>