<template>
  <div id="search">
    <a-row>
      <a-col :span="12" :offset="2">
        <a-row>
          <a-col :span="20">
            <a-input-search
              placeholder="请输入您想搜索的内容"
              v-model:value="inputValue"
              @search="onSearch"
              style="{'box-shadow': 5px 5px 10px gray}"
            />
          </a-col>
        </a-row>

        <br />

        <div>
          <a-tabs :default-active-key="tabNow" @change="onChangeTab" size="small">
            <a-tab-pane key="1" tab="问题">
              <QuestionForSearch v-for="(item) in questionData" v-bind:key="item.qid" :ques="item" />
            </a-tab-pane>
            <a-tab-pane key="2" tab="用户">
              <UserForSearch v-for="(item) in userData" v-bind:key="item.uid" :user="item"/>
            </a-tab-pane>
            <a-tab-pane key="3" tab="回答">
              <AnswerForSearch v-for="(item) in answerData" v-bind:key="item.aid" :ans="item"/>
            </a-tab-pane>
          </a-tabs>
        </div>
      </a-col>
      <a-col :span="8" :offset="1">
        <CardForSearch v-for="(item) in cardInfo" v-bind:key="item.title" :info="item" />
      </a-col>
    </a-row>
  </div>
</template>

<script>
import { Options, Vue } from "vue-class-component";
import QuestionForSearch from "@/components/QuestionForSearch.vue";
import CardForSearch from "@/components/CardForSearch.vue";
import server from "@/http/request.js";
import UserForSearch from "@/components/UserForSearch";
import AnswerForSearch from "@/components/AnswerForSearch";
import { postRequest,getRequest } from "@/http/request.js";


export default {
  components: {UserForSearch, QuestionForSearch, CardForSearch, AnswerForSearch },
  data() {
    return {
      questionData: [],
      userData:[],
      answerData:[],
      searchValue: "",
      cardInfo: [],
      inputValue: "wge",
      questionPageNow:0,
      answerPageNow:0,
      userPageNow:0,
      tabNow:1
    };
  },
  created() {
    this.searchValue=this.$route.query.content;
    this.inputValue=this.$route.query.content;
    console.log(this.inputValue);
    this.searchQuestion();
    this.searchCard();
  },
  methods: {
    onSearch(value) {
      this.searchValue = value;
      this.questionPageNow=0;
      this.userPageNow=0;
      this.answerPageNow=0;
      if (this.tabNow===1){
        this.searchQuestion();
      } else if (this.tabNow===2){
        this.searchUser();
      } else {
        this.searchAnswer();
      }
      this.searchCard();
    },
    searchQuestion(){
      getRequest("/searchQuestions",(e)=>{
        this.questionData=this.questionData.concat(e.result);
        this.questionPageNow++;
      }, {errorCallback:(e)=>{console.log(e)},
        params:{page:this.questionPageNow,text:this.searchValue}});
    },
    searchUser(){
      getRequest("/searchUsers",(e)=>{
        console.log(e);
        this.userData=this.userData.concat(e.result);
        this.userPageNow++;
      }, {errorCallback:(e)=>{console.log(e)},
        params:{page:this.userPageNow,text:this.searchValue}});
    },
    searchAnswer(){
      getRequest("/searchAnswers",(e)=>{
        this.answerData=this.answerData.concat(e.result);
        this.answerPageNow++;
      }, {errorCallback:(e)=>{console.log(e)},
        params:{page:this.answerPageNow,text:this.searchValue}});
    },
    searchCard(){
      getRequest("/search",(e)=>{
        this.cardInfo=e.result;
      }, {errorCallback:(e)=>{console.log(e)},
        params:{text:this.searchValue}});
    },
    onChangeTab(key){
      this.tabNow=key;
      if (this.tabNow==1){
        if (this.questionData.length==0)
          this.searchQuestion();
      } else if (this.tabNow==2){
        if (this.userData.length==0)
          this.searchUser();
      } else {
        if (this.answerData.length==0)
          this.searchAnswer();
      }
    }
  },
};
</script>

<style scoped>
#search {
  min-height: 667px;
  background-color: #edeeed;
}

.ant-input-affix-wrapper {
  border: 1px solid #d9d9d9;
  border-radius: 20px;
  padding: 4px 11px;
  width: 100%;
  text-align: start;
  background-color: #fff;
  background-image: none;
  color: rgba(0, 0, 0, 0.65);

}
</style>
