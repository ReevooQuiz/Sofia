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
      <QuestionForSearch v-for="(item) in recommendedQuestionData" v-bind:key="item.id" :ques="item" />
    </a-col>
    <a-col :span="4">
<!--      <a-card  size="small" :bordered="false">热榜</a-card>-->
<!--      <a-card  size="small"  v-for="(item) in hotRankData" v-bind:key="item.index">-->
<!--        <a-row type="flex" justify="space-between">-->
<!--          <a style="color: #000000">{{item.index}} {{ item.name}}</a>-->
<!--          <a style="color: #000000">1242535</a>-->
<!--        </a-row>-->
<!--      </a-card>-->
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
const recommendedQuestionData = [
  {
    id: 1,
    title: "Ant Design Title 1",
    user: "akvfcdg",
    description: "dfghjklfcghjm,dfgbndghjkrewqwertyuiytr",
    likeNum: 123,
    dislikeNum: 4567,
    commentNum: 7890
  },
  {
    id: 2,
    title: "A7654 Title 1",
    user: "dfsfskang",
    description: "dfghjklfcghjm,dfgbndghjkrewqwertyuiytr",
    likeNum: 123,
    dislikeNum: 4567,
    commentNum: 7890
  },
  {
    id: 3,
    title: "gfds 1",
    user: "ererang",
    description: "dfghjklfcghjm,dfgbndghjkrewqwertyuiytr",
    likeNum: 123,
    dislikeNum: 4567,
    commentNum: 7890
  },
  {
    id: 4,
    title: "ytre",
    user: "akgfds",
    description: "dfghjklfcghjm,dfgbndghjkrewqwertyuiytr",
    likeNum: 123,
    dislikeNum: 4567,
    commentNum: 7890
  }
];
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
const hotRankData=[{
  index:1,
  name:"wg23rgeeb",
  heat:"2346"
},{
  index:2,
  name:"wg23rgeeb",
  heat:"2346"
},{
  index:3,
  name:"wg23rgeeb",
  heat:"2346"
},{
  index:4,
  name:"wg23rgeeb",
  heat:"2346"
},{
  index:5,
  name:"wg23rgeeafsssssssgeqwhrnjetkry,tu.lkjehwrgeb",
  heat:"2346"
},{
  index:6,
  name:"wg23rgeerkdjehsktdulyktjhrstykfjtretjsytstjhrsjheratjb",
  heat:"2346"
},{
  index:7,
  name:"wg23rgeeb",
  heat:"2346"
},{
  index:8,
  name:"wg23rgeeb",
  heat:"2346"
},{
  index:9,
  name:"wg23rgeeb",
  heat:"2346"
},{
  index:10,
  name:"wg23rgeeb",
  heat:"2346"
}
];
const columns=[
  {
    dataIndex: 'index',
    key: 'index',
    slots: {  title: 'customTitle' },
    width:50
  },{
    dataIndex: 'name',
    key: 'name',
    ellipsis:true,
  },{
    dataIndex: 'heat',
    key: 'heat',
    width:70
  },
];
export default {
  components: {QuestionForSearch,},
  data() {
    return {
      userData:userData,
      recommendedQuestionData:recommendedQuestionData,
      hotRankData:hotRankData,
      columns,
    };
  },
}
</script>

<style scoped>

</style>
