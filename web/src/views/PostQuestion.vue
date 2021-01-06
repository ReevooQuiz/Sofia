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
          <a-textarea v-model:value="labelString" placeholder="格式：label1#label2#" showCount :maxlength="100" />
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
        <a-select-option value="fashion">
          时尚
        </a-select-option>
      </a-select>
      <br/>
      <br/>
      <a-button type="primary" shape="round" @click="onPostQuestion"><strong>&ensp;发&ensp;布&ensp;</strong></a-button>
    </a-col>
  </a-row>
</template>

<script>
import {postRequest} from "@/http/request";
import {message} from "ant-design-vue";

export default {
  data(){
    return {
      questionValue:"",
      category:"study",
      title:"",
      labelString:"",
      labels:[]
    };
  },
  methods: {
    onPostQuestion(){
      let start=0,end=0;
      while (start<this.labelString.length){
        end=this.labelString.indexOf("#",start);
        labels.concat(this.labelString.slice(start,end));
        start=end+1;
      }
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
  },
}
</script>

<style scoped>

</style>
