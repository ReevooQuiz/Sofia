<template>
  <div>
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
          <a-button v-if="writeAnswer" @click="onWriteAnswer">取消回答</a-button>
                <a-button v-else @click="onWriteAnswer">我要回答</a-button>
        </a-space>
      </a-row>
        <br/>
        <v-md-editor v-if="writeAnswer" v-model="writeAnswerValue"></v-md-editor>
        <span v-if="writeAnswer">
          <br/>
              <a-row type="flex" justify="end">
                <a-button @click="onCommitAnswer" type="primary" shape="pill" size="small">提交回答</a-button>
              </a-row>
        </span>
      <br />
      <a-row>
        <AnswerCard v-for="(item) in data.answer_list" v-bind:key="item.aid" :ans="item" />
      </a-row>
    </a-col>
  </div>
</template>

<script >
import { Options, Vue } from "vue-class-component";
import QuestionHead from "@/components/QuestionHead.vue";
import AnswerCard from "@/components/AnswerCard.vue";
import { DownOutlined } from "@ant-design/icons-vue";
import {server} from "@/http/request.js";
const data = {
  code: 0,
  qid: 234,
  owner: {
    user_id: 1,
    user_name: "abc",
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
        user_name: "abc",
        user_icon: ""
      },
      like_count: 2,
      criticism_count: 4,
      approval_count: 2,
      comment_count: 2,
      content: "asadqwf"
    }
  ]

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
      orderNow:"按热度排序",
      writeAnswer:false,
      writeAnswerValue:"",
      questionHead:{}
    };
  },
  created() {
    console.log(this.$route.query.questionId);
    let p= this.$route.query.questionId;
    server
      .get("/question", {
        params: p
      })
      .catch(function(error) {
        console.log(error);
      })
      .then(response => {
        console.log("!111");
        console.log(response);
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
    },
    onWriteAnswer(){
      window.scrollTo(0,window.outerHeight);
      if (this.writeAnswer){
        this.writeAnswer=false;
        this.writeAnswerValue="";
        return;
      }
        this.writeAnswer=true;
    },
    onCommitAnswer(){

      this.writeAnswer=false;
      this.writeAnswerValue="";
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
