<template>
  <div style="width: 256px">
    <!--<a-button type="primary"  style="margin-bottom: 16px">
      <MenuUnfoldOutlined v-if="collapsed" />
      <MenuFoldOutlined v-else />
    </a-button>-->
    <a-menu
      :default-selected-keys="['1']"
      mode="inline"
      :inline-collapsed="collapsed"
    ><a-menu-item key="1"> <router-link to="/personalSet">
        <SettingOutlined />
          设置 </router-link>

      </a-menu-item>
     <a-menu-item key="2"  >
       <router-link to="/personalFollowing">
       <HeartOutlined /> 关注
       </router-link>
      </a-menu-item>
      <a-menu-item key="3"><router-link to="/personalFollower">
      <SmileOutlined />  粉丝
      </router-link>
      </a-menu-item>
      <a-menu-item key="4"><router-link to="/personalQuestion">
        <PicLeftOutlined /> 问题
      </router-link>
      </a-menu-item>
      <a-menu-item key="5"> <router-link to="/personalAnswer">
        <BulbOutlined /> 回答
      </router-link>
      </a-menu-item>
      <a-menu-item key="6">  <router-link to="/personalMessage">
      <MailOutlined /> 消息
      </router-link>
      </a-menu-item>


      <a-menu-item key="7"> <router-link to="/personalCollection">
       <StarOutlined />
       收藏 </router-link>

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
  SettingOutlined,
  StarOutlined

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
   SettingOutlined,
   StarOutlined


  },
  data() {
    return {
      collapsed:false,
    };
  },
  methods: {
    // toggleCollapsed() {
    //   this.collapsed = !this.collapsed;
    // },
  },
};
</script>
