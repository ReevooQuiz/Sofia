<template>
  <div>
      
    <a-row >
        <a-col :span="1" :offset="1">
            <SubMenu/>
        </a-col>
      
       
    </a-row>
    <br />
    
    <br />

   
  </div>
</template>

<script>
import { defineComponent } from "vue";
import { Options, Vue } from "vue-class-component";
import { UserOutlined } from "@ant-design/icons-vue";
import SubMenu from "../../components/PersonalNavigation"
const data1 = [
  {
    title: "Title 1"
  },
  {
    title: "Title 2"
  },
  {
    title: "Title 3"
  },
  {
    title: "Title 4"
  }
];
export default {
  components: {
    UserOutlined,
    SubMenu
  },
  data() {
    return {
      dataa: data1
    };
  },
  methods: {
    onSearch(value) {
      if (value != "") {
        this.$router.push({ path: "/search", query: { content: value } });
      }
    },
    goToFollowing() {
      this.$router.push({ path: "/personalFollowing" });
    }
  }
};
</script>

<style>
body {
  height: 100%;
  background-color: #edeeed;
}
</style>