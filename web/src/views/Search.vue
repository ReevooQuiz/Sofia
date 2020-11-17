<template>
  <div id="search">
    <a-row>
      <a-col :span="12" :offset="2">
        <a-row>
          <a-col :span="20">
            <a-input-search
              placeholder="输入您的问题"
              v-model:value="inputValue"
              @search="onSearch"
              style="{'box-shadow': 5px 5px 10px gray}"
            />
          </a-col>
        </a-row>

        <br />

        <div>
          <a-tabs default-active-key="1" @change="callback" :size="small">
            <a-tab-pane key="1" tab="问题">
              <QuestionForSearch v-for="(item) in questionData" v-bind:key="item.id" :ques="item" />
            </a-tab-pane>
            <a-tab-pane key="2" tab="用户">用户</a-tab-pane>
          </a-tabs>
        </div>
      </a-col>
      <a-col :span="8" :offset="1">
        <CardForSearch :info="cardInfo" />
      </a-col>
    </a-row>
  </div>
</template>

<script>
import { Options, Vue } from "vue-class-component";
import QuestionForSearch from "@/components/QuestionForSearch.vue";
import CardForSearch from "@/components/CardForSearch.vue";
import server from "@/http/request.js";

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
    description:
      "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打",
    answer_count: 4,

    follow_count: 234
  },
  {
    qid: 345,
    owner: {
      user_id: 3,
      user_name: "violedo",
      user_icon:
        "https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
    },
    title: "如何看待上海交通大学花店事件",
    description:
      "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打",
    answer_count: 4,
    follow_count: 234
  },
  {
    qid: 345,
    owner: {
      user_id: 3,
      user_name: "violedo",
      user_icon:
        "https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
    },
    title: "如何看待上海交通大学花店事件",
    description:
      "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打",
    answer_count: 4,
    follow_count: 234
  },
  {
    qid: 345,
    owner: {
      user_id: 3,
      user_name: "violedo",
      user_icon:
        "https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
    },
    title: "如何看待上海交通大学花店事件",
    description:
      "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打",
    answer_count: 4,
    follow_count: 234
  }
];

const data2 = {
  title: "上海交通大学",
  keyWords: ["国际知名大学", "二月十三", "C9高校"],
  attributes: [
    ["前身", "南洋公学"],
    ["校长", "林忠钦"],
    ["位于", "上海"],
    ["书记", "姜斯宪"]
  ],
  tags: ["二月十三", "复旦"]
};
export default {
  components: { QuestionForSearch, CardForSearch },
  data() {
    return {
      questionData: data,
      searchValue: "",
      cardInfo: data2,
      inputValue: "wge"
    };
  },
  created() {
    this.searchValue=this.$route.query.content;
    this.inputValue=this.$route.query.content;
    console.log(this.inputValue);
    // server
    //   .post("/search", {
    //     value: this.searchValue
    //   })
    //   .catch(function(error) {
    //     console.log(error);
    //   })
    //   .then(response => {
    //     this.questionData= response.data.questionData;
    //   });
    //todo:search function
  },
  methods: {
    onSearch(value) {
      this.searchValue = value;
      console.log(this.searchValue);

      server
        .get("/search", {
          params: this.searchValue
        })
        .catch(function(error) {
          console.log(error);
        })
        .then(response => {
          console.log("!");
          console.log(response.data);
          console.log(response.data.question_list);
          this.questionData = response.data.question_list;
          console.log(this.questionData);
        });
    },

    handleInit(response) {
      this.questionData = response.data.questionData;
      this.cardInfo = response.data.cardInfo;
    }
  },
  // created: function() {
  //   server
  //     .get("/get", {
  //       params: {}
  //     })
  //     .then(response => this.handleInit(response))
  //     .catch(function(error) {
  //       console.log(error);
  //     });
  // }

};
</script>

<style>
#search {
  min-height: 667px;
  background-color: #edeeed;
}

.ant-input-affix-wrapper {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
  font-variant: tabular-nums;
  list-style: none;
  font-feature-settings: "tnum";
  position: relative;
  display: inline-flex;
  border: 1px solid #d9d9d9;
  border-radius: 20px;
  padding: 4px 11px;
  width: 100%;
  text-align: start;
  background-color: #fff;
  background-image: none;
  color: rgba(0, 0, 0, 0.65);
  font-size: 14px;
  line-height: 1.5715;
}
</style>
