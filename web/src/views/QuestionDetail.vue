<template>
  <div id="Question">
    <a-col :span="20" :offset="2">
      <a-row>
        <div>
          <QuestionHead :ques="questionHead" />
        </div>
      </a-row>
      <a-row justify="space-between">
        <a-tag>{{data.answer_count}} 个回答</a-tag>
        <a-space :size="20">
          <a-dropdown>
            <template #overlay>
              <a-menu @click="handleOrderClick">
                <a-menu-item key="0">按热度排序</a-menu-item>
                <a-menu-item key="1">按时间排序</a-menu-item>
              </a-menu>
            </template>
            <a-button>
              {{orderNow}}
              <DownOutlined />
            </a-button>
          </a-dropdown>
          <a-button>我要回答</a-button>
        </a-space>
      </a-row>
      <br />
      <a-row>
        <AnswerCard v-for="(item) in data.answer_list" v-bind:key="item.aid" :ans="item" />
      </a-row>
    </a-col>
  </div>
</template>

<script>
import { Options, Vue } from "vue-class-component";
import QuestionHead from "@/components/QuestionHead.vue";
import AnswerCard from "@/components/AnswerCard.vue";
import { DownOutlined } from "@ant-design/icons-vue";
import server from "@/http/request.js";
const data = {
  code: 0,
  qid: 234,
  owner: {
    user_id: 1,
    user_name: abc,
    user_icon: ""
  },
  title: "ababa",
  content: "abaaba",
  answer_count: 4,
  follow_count: 234,
  view_count: 123,

  answer_list: [
    {
      aid: 234,
      owner: {
        user_id: 1,
        user_name: abc,
        user_icon: ""
      },
      like_count: 2,
      criticism_count: 4,
      approval_count: 2,
      comment_count: 2,
      content: "asadqwf"
    }
  ]
  // questionHead:{
  //   id: 1,
  //   title: "Ant Design Title 1",
  //   user:"akvfcdg",
  //   description:"dfghjklfcghjm,dfgbndghjkrewqwertyuiytr",
  //   likeNum :123,
  //   dislikeNum:4567,
  //   commentNum:7890,
  //   content:"We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully andefficiently.",
  //   followers: 1551,
  //   answers:114514
  // },
  // answers:[
  //     {
  //         id:1,
  //         user:"akangakang",
  //         likeNum :123,
  //         dislikeNum:4567,
  //         commentNum:7890,
  //         content:"We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully andefficiently.",
  //     },
  //     {
  //         id:2,
  //         user:"violedo",
  //         likeNum :12524,
  //         dislikeNum:427,
  //         commentNum:785460,
  //         content:"We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully andefficiently.",
  //     },
  //     {
  //         id:3,
  //         user:"zhc",
  //         likeNum :1465,
  //         dislikeNum:478,
  //         commentNum:7250,
  //         content:"We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully andefficiently.",
  //     },
  // ]
};

const orderBy = ["按热度排序", "按时间排序"];

export default {
  components: {
    QuestionHead,
    DownOutlined,
    AnswerCard
  },
  data() {
    return {
      data,
      orderNow: "按热度排序",
      questionHead
    };
  },
  created() {
    console.log(this.$route.query.questionId);
    server
      .get("/question", {
        params: this.$route.query.questionId
      })
      .catch(function(error) {
        console.log(error);
      })
      .then(response => {
        console.log("!");
        // console.log(response.data);
        // console.log(response.data.question_list);
        this.data = response.data;
        // console.log(this.questionData);

        this.questionHead.title = response.data.title;
        this.questionHead.tags = response.data.tags;
        this.questionHead.user = response.data.owner.user_name;
        this.questionHead.content = response.data.content;
        this.questionHead.followers = response.data.follower_count;
        this.questionHead.viewCount = response.data.view_count;
      });
  },
  methods: {
    handleOrderClick(e) {
      this.orderNow = orderBy[e.key];
    }
  }
};
</script>

<style>
#search {
  min-height: 667px;
  background-color: #edeeed;
}
</style>
