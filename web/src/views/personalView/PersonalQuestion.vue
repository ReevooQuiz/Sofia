<template>
  <div>
    <a-row>
      <a-col :span="1" :offset="1">
        <div class="personal-nav">
          <SubMenu />
        </div>
      </a-col>

      <a-col :span="15" :offset="2">
        <br />
        <br />

        <a-list
          item-layout="vertical"
          size="large"
          :pagination="pagination"
          :data-source="questionData"
        >
          <template #renderItem="{ item, index }">
            <a-list-item key="item.title">
              <QuestionForPersonal :ques="item" />
              {{ item.content }}
            </a-list-item>
          </template>
        </a-list>
        <!-- <QuestionForPersonal v-for="(item) in questionData" v-bind:key="item.id" :ques="item" /> -->
      </a-col>
    </a-row>

    <br />

    <br />
  </div>
</template>

<script>
import { defineComponent } from "vue";
import { Options, Vue } from "vue-class-component";
import {
  UserOutlined,
  StarOutlined,
  LikeOutlined,
  MessageOutlined
} from "@ant-design/icons-vue";
import SubMenu from "../../components/PersonalNavigation";
import QuestionForPersonal from "../../components/QuestionForPersonal";
const data = [
  {
    qid: 1,
    owner: {
      user_id: 1,
      user_name: "阿钪",
      user_icon:
        "https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
    },
    title: "如何看待上海交通大学花店事件",
    time: "2015-08-05T08:40:51.620Z",

    view_count: 10,
    favorite_count: 20,
    category: "study",
    labels: ["programming", "campus"],
    head: "What if we put",
    pictureUrls: [
      "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png",
      "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"
    ],

    head:
      "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打\
      近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打",
    answer_count: 4,

    follow_count: 234
  },
  {
    qid: 2,
    owner: {
      user_id: 1,
      user_name: "阿钪",
      user_icon:
        "https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
    },
    title: "如何看待上海交通大学花店事件",
    time: "2015-08-05T08:40:51.620Z",

    view_count: 10,
    favorite_count: 20,
    category: "study",
    labels: ["programming", "campus"],
    head: "What if we put",
    pictureUrls: [
      "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png",
      "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"
    ],

    head:
      "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打\
      近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打",
    answer_count: 4,

    follow_count: 234
  },
  {
    qid: 3,
    owner: {
      user_id: 1,
      user_name: "阿钪",
      user_icon:
        "https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
    },
    title: "如何看待上海交通大学花店事件",
    time: "2015-08-05T08:40:51.620Z",

    view_count: 10,
    favorite_count: 20,
    category: "study",
    labels: ["programming", "campus"],
    head: "What if we put",
    pictureUrls: [
      "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png",
      "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"
    ],

    head:
      "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打\
      近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打",
    answer_count: 4,

    follow_count: 234
  },
  {
    qid: 4,
    owner: {
      user_id: 1,
      user_name: "阿钪",
      user_icon:
        "https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
    },
    title: "如何看待上海交通大学花店事件",
    time: "2015-08-05T08:40:51.620Z",

    view_count: 10,
    favorite_count: 20,
    category: "study",
    labels: ["programming", "campus"],
    head: "What if we put",
    pictureUrls: [
      "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png",
      "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"
    ],

    head:
      "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打\
      近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打",
    answer_count: 4,

    follow_count: 234
  }
];
export default {
  components: {
    UserOutlined,
    SubMenu,
    QuestionForPersonal,
    StarOutlined,
    LikeOutlined,
    MessageOutlined
  },
  data() {
    return {
      questionData: data,
      pagination: {
        onChange: page => {
          console.log(page);
        },
        pageSize: 3
      }
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

.personal-nav .ant-menu-inline {
  width: 40%;
}
</style>