<template>
  <div>
    <!-- <a-row justify="center"> -->


    <a-col :span="16" :offset="4">
      <a-row>
        <div>
          <QuestionHead :ques="questionHead" />
        </div>
      </a-row>
      <a-row justify="space-between">
        <a-tag> {{questionHead.answer_count}} 个回答</a-tag>
        <a-space :size="19">
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
          <span v-if="!questionHead.closed">
            <a-button v-if="writeAnswer" @click="onWriteAnswer">取消回答</a-button>
            <a-button v-else @click="onWriteAnswer">我要回答</a-button>
          </span>
        </a-space>
      </a-row>
        <br/>
        <v-md-editor v-if="writeAnswer" v-model="writeAnswerValue"></v-md-editor>
        <span v-if="writeAnswer">
          <br/>
              <a-row type="flex" justify="end">
                <a-button @click="onCommitAnswer" type="primary" shape="pill" size="small">提交回答</a-button>
              </a-row>
          <br/>
        </span>
      <a-row>
        <AnswerCard v-for="(item) in answerData" v-bind:key="item.aid" :ans="item" />
      </a-row>
      <a-row type="flex" justify="space-around" >
        <a-col>
          <div
              v-if="showLoadingMore"
              :style="{ textAlign: 'center', marginTop: '12px', height: '32px', lineHeight: '32px' }"
          >
            <a-spin v-if="loadingMore" />
            <a-button v-else @click="onLoadMore">
              加载更多
            </a-button>
          </div>
          <div v-else>已经到底了</div>
        </a-col>
      </a-row>
    </a-col>

    <!-- </a-row> -->
  </div>
</template>

<script >
import { Options, Vue } from "vue-class-component";
import QuestionHead from "@/components/QuestionHead.vue";
import AnswerCard from "@/components/AnswerCard.vue";
import { DownOutlined } from "@ant-design/icons-vue";
import { postRequest,getRequest,putRequest } from "@/http/request.js";

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
      questionHead:{raiser:{}},
      answerData:[],
      pageNow:0,
      showLoadingMore: true,
      loadingMore:true
    };
  },
  created() {
    //console.log(this.$route.query.questionId);
    let p= this.$route.query.questionId;
    this.qid=p;
    getRequest("/question",
        (response)=>{
      this.questionHead=response.result;
      console.log(response);
    }, {
      errorCallback:(e)=>{console.log(e)},
      params:{qid:p}
    });
    this.getAnswers();
  },
  methods: {
    getAnswers(){
      getRequest("/answers",
          (response)=>{
            this.answerData=this.answerData.concat(response.result);
            this.loadingMore=false;
            if (response.result.length==0)
              this.showLoadingMore=false;
          },
          {
        errorCallback:(e)=>{console.log(e)},
        params:{qid:this.questionHead.qid,page:this.pageNow++,sort:this.orderKey}
      });
    },
    onLoadMore(){
      this.loadingMore = true;
      this.getAnswers();
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
        params:{qid:this.questionHead.qid,page:0,sort:e.key}
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
      //console.log(this.writeAnswerValue);
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
