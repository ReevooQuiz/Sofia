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
