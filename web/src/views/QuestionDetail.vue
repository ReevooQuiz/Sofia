<template>
  <div id="Question">
    <a-col :span="20" :offset="2">
        <a-row>
            <div>
                <QuestionHead :ques="data.questionHead"/>
            </div>
        </a-row>
        <a-row justify="space-between">
            <a-tag>{{data.questionHead.answers}} 个回答</a-tag>
            <a-space :size=20>
                <a-dropdown>
                    <template #overlay>
                        <a-menu @click="handleOrderClick">
                            <a-menu-item key="0">
                                按热度排序
                            </a-menu-item>
                            <a-menu-item key="1">
                                按时间排序
                            </a-menu-item>
                        </a-menu>
                    </template>
                    <a-button> {{orderNow}} <DownOutlined /> </a-button>
                </a-dropdown>
                <a-button>我要回答</a-button>
            </a-space>
        </a-row>
        <br/>
        <a-row>
            <AnswerCard v-for="(item) in data.answers" v-bind:key="item.id" :ans="item"/>
        </a-row>
    </a-col>
  </div>
</template>

<script>
import { Options, Vue } from "vue-class-component";
import QuestionHead from "@/components/QuestionHead.vue";
import AnswerCard from "@/components/AnswerCard.vue";
import { DownOutlined } from '@ant-design/icons-vue';
const data = {
  questionHead:{
    id: 1,
    title: "Ant Design Title 1",
    user:"akvfcdg",
    description:"dfghjklfcghjm,dfgbndghjkrewqwertyuiytr",
    likeNum :123,
    dislikeNum:4567,
    commentNum:7890,
    content:"We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully andefficiently.",
    followers: 1551,
    answers:114514
  },
  answers:[
      {
          id:1,
          user:"akangakang",
          likeNum :123,
          dislikeNum:4567,
          commentNum:7890,
          content:"We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully andefficiently.",
      },
      {
          id:2,
          user:"violedo",
          likeNum :12524,
          dislikeNum:427,
          commentNum:785460,
          content:"We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully andefficiently.",
      },
      {
          id:3,
          user:"zhc",
          likeNum :1465,
          dislikeNum:478,
          commentNum:7250,
          content:"We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully andefficiently.",
      },
  ]

};

const orderBy=["按热度排序","按时间排序"];

export default{
  components: {
    QuestionHead,
    DownOutlined,
    AnswerCard
  },
  data() {
    return {
      data,
      orderNow:"按热度排序",
    };
  },
  created(){
      console.log(this.$route.query.questionId)
  },
  methods: {
    handleOrderClick(e) {
        this.orderNow=orderBy[e.key];
    },
  },

}
</script>

<style>
#search {
  min-height: 667px;
  background-color: #edeeed;
}
</style>
