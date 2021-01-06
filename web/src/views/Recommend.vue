<template>
  <a-row type="flex" justify="space-around" align="top">
    <a-col :span="4">
      <a-card  size="small" :bordered="false">用户推荐</a-card>
      <a-card  size="small"  v-for="(item) in userData" v-bind:key="item.uid">
        <a-card-meta :title="item.name">
          <template #avatar>
            <a-avatar :src="item.icon" />
          </template>
        </a-card-meta>
      </a-card>
    </a-col>
    <a-col :span="12">
      <QuestionForSearch v-for="(item) in recommendedQuestionData" v-bind:key="item.qid" :ques="item" />
    </a-col>
    <a-col :span="4">
      <a-table
          :columns="columns"
          :data-source="hotRankData"
          :pagination="false"
          size="small"
      >
        <template #customTitle>
          <h4><strong>热榜</strong></h4>
        </template>
      </a-table>
    </a-col>
  </a-row>
</template>

<script>
import QuestionForSearch from "@/components/QuestionForSearch.vue";
import UserForSearch from "@/components/UserForSearch";
import CardForSearch from "@/components/CardForSearch";
import { postRequest,getRequest } from "@/http/request.js";

const userData=[{
  uid:"sdfw",
  name:"dsfwfwg",
  nickname:"nick",
  profile:"sdg3g2gbrgefgrwwgbfrrwfgerwger",
  icon:"https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"
},{
  uid:"sdfw",
  name:"dsfwfwg",
  nickname:"nick",
  profile:"sdg3g2gbrgefgrwwgbfrrwfgerwger",
  icon:"https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"
},{
  uid:"sdfw",
  name:"dsfwfwg",
  nickname:"nick",
  profile:"sdg3g2gbrgefgrwwgbfrrwfgerwger",
  icon:"https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"
},
];
const columns=[
  {
    dataIndex: 'index',
    key: 'index',
    slots: {  title: 'customTitle' },
    width:50
  },{
    dataIndex: 'title',
    key: 'title',
    ellipsis:true,
  },{
    dataIndex: 'view_count',
    key: 'view_count',
    width:70
  },
];
export default {
  components: {QuestionForSearch,},
  data() {
    return {
      userData:userData,
      recommendedQuestionData: [],
      hotRankData:[],
      columns,
    };
  },
  created() {
    getRequest("/questions",
        (response)=>{
          this.recommendedQuestionData=response.result;
        }, {
          errorCallback:(e)=>{console.log(e)},
          params:{}
        });
    getRequest("/hotlist",
        (response)=>{
          this.hotRankData=response.result;
          let i=0;
          for (;i<10;){
            if (i>=this.hotRankData.length)
              break;
            this.hotRankData[i].index=++i;
          }
        }, {
          errorCallback:(e)=>{console.log(e)},
          params:{}
        });
  }
}
</script>

<style scoped>

</style>
