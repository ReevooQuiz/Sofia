<template>
  <a-row>
    <a-col :span="18" :offset="2">
      <a-row>
        <a-col :span="14">
          请输入标题
          <a-textarea v-model:value="title" placeholder="请输入标题" showCount :maxlength="100" />
        </a-col>
        <a-col :span="9" :offset="1">
          请输入您想标注的标签
          <a-textarea v-model:value="labelString" placeholder="格式：label1#label2#" showCount :maxlength="200" />
        </a-col>
      </a-row>
      <v-md-editor v-model="questionValue" height="400px"></v-md-editor>
    </a-col>
    <a-col :span="2" :offset="1" >
      <a-select
          v-model:value="category"
          style="width: 80px"
          ref="select"
      >
        <a-select-option value="study">
          学习
        </a-select-option>
        <a-select-option value="life">
          生活
        </a-select-option>
      </a-select>
      <br/>
      <br/>
      <a-button type="primary" shape="round" @click="onPostQuestion"><strong>&ensp;发&ensp;布&ensp;</strong></a-button>
    </a-col>
  </a-row>
</template>

<script>
import {message} from "ant-design-vue";
import { postRequest,getRequest } from "@/http/request.js";

export default {
  data(){
    return {
      questionValue:"",
      category:"study",
      title:"",
      labelString:"",
      labels:[],
      qid:null
    };
  },
  created() {
    let p= this.$route.query.questionId;
    if (p!=null){
      this.qid=p;
      getRequest("/question",
          (response)=>{
            this.questionValue=response.result.content;
            this.title=response.result.title;
            this.category=response.result.category;
            for (var i=0;i<response.result.labels.length;++i){
              this.labelString=this.labelString.concat(response.result.labels[i]).concat("#");
            }
            console.log(response);
          }, {
            errorCallback:(e)=>{console.log(e)},
            params:{qid:p}
          });
    }

  },
  methods: {
    onPostQuestion(){
      let start=0,end=0,count=0;
      while (start<this.labelString.length){
        if (count>=5){
          message.error("label个数不应超过5个");
          return;
        }
        end=this.labelString.indexOf("#",start);
        if (end-start>32){
          message.error("单个label不应超过32个字符");
          return;
        }
        this.labels=this.labels.concat(this.labelString.slice(start,end));
        start=end+1;
        count++;
      }
      if (this.qid!=null){
        postRequest("/questions",
            {
              qid:this.qid,
              title:this.title,
              content:this.questionValue,
              labels:this.labels,
              category:this.category,
            },(e)=>{
              //console.log(e);
              if (e.code==0){
                message.success("发布成功");
                this.$router.push({ path:'/question' , query: { questionId: this.qid } });
              }
            },{errorCallback:(e)=>{
                console.log(e);
              }});
      }
      else {
        postRequest("/questions",
            {
              title:this.title,
              content:this.questionValue,
              labels:this.labels,
              category:this.category,
            },(e)=>{
              //console.log(e);
              if (e.code==0){
                message.success("发布成功");
                this.$router.push({ path:'/question' , query: { questionId: e.result.qid } });
              }
            },{errorCallback:(e)=>{
                console.log(e);
              }});
      }
    }
  },
}
</script>

<style scoped>

</style>
