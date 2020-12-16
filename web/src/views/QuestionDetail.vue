<template>
  <div>
    <a-col :span="20" :offset="2">
      <a-row>
        <div>
          <QuestionHead :ques="questionHead" />
        </div>
      </a-row>
      <a-row justify="space-between">
        <a-tag>{{questionHead.answer_count}} 个回答</a-tag>
        <a-space :size="20">
          <a-dropdown>
            <template #overlay>
              <a-menu @click="handleOrderClick">
                <a-menu-item key="0">按时间排序</a-menu-item>
                <a-menu-item key="1">按热度排序</a-menu-item>
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
        <AnswerCard v-for="(item) in answerData" v-bind:key="item.aid" :ans="item" />
      </a-row>
    </a-col>
  </div>
</template>

<script >
import { Options, Vue } from "vue-class-component";
import QuestionHead from "@/components/QuestionHead.vue";
import AnswerCard from "@/components/AnswerCard.vue";
import { DownOutlined } from "@ant-design/icons-vue";
import { postRequest,getRequest } from "@/http/request.js";

const orderBy = ["按时间排序", "按热度排序"];

export default {
  components: {
    QuestionHead,
    DownOutlined,
    AnswerCard
  },
  data() {
    return {
      orderNow:"按时间排序",
      orderKey:0,
      writeAnswer:false,
      writeAnswerValue:"",
      questionHead:{},
      answerData:[],
      pageNow:0
    };
  },
  created() {
    console.log(this.$route.query.questionId);
    let p= this.$route.query.questionId;
    getRequest("/question",
        (response)=>{
      this.questionHead=response.result
    }, {
      errorCallback:(e)=>{console.log(e)},
      params:{qid:p}
    });
    this.getAnswers();
  },
  methods: {
    getAnswers(){
      getRequest("/answers",
          (response)=>{this.answerData.append(response.result);},
          {
        errorCallback:(e)=>{console.log(e)},
        params:{qid:p,page:this.pageNow++,sort:this.orderKey}
      });

    },
    reorderAnswers(response){
      this.answerData=response.result;
      this.pageNow=1;
    },
    handleOrderClick(e) {
      this.orderNow = orderBy[e.key];
      this.orderKey=e.key;
      getRequest("/answers",this.reorderAnswers, {
        errorCallback:(event)=>{console.log(event)},
        params:{qid:p,page:0,sort:e.key}
      });
    },
    onWriteAnswer(){
      if (this.writeAnswer){
        this.writeAnswer=false;
        this.writeAnswerValue="";
        return;
      }
      this.writeAnswer=true;
    },
    onCommitAnswer(){
      postRequest("/answers", {qid:this.questionHead.qid,content:this.writeAnswerValue},(e)=>{
        console.log(e);
        this.writeAnswer=false;
        this.writeAnswerValue="";
        this.questionHead.answer_count++;
        this.answerData=[];
        this.pageNow=0;
        this.getAnswers();
      },{errorCallback:(e)=>{
          console.log(e);
        }});
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
