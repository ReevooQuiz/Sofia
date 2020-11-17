<template>
  <div id="search">
    <a-row>
      <a-col :span="12" :offset="2">
        <a-row>
          <a-col :span="20">
            <a-input-search
              placeholder="input search text"
              style="{'box-shadow': 5px 5px 10px gray}"
            />
          </a-col>
        </a-row>

        <br />

        <div>
          <a-tabs default-active-key="1" @change="callback" :size="small">
            <a-tab-pane key="1" tab="问题">
            
              <QuestionForSearch v-for="(item) in data" v-bind:key="item.id" :ques="item"/>
              
            </a-tab-pane>
            <a-tab-pane key="2" tab="用户">用户</a-tab-pane>
          </a-tabs>
        </div>
      </a-col>
      <a-col :span="8" :offset="1">
<CardForSearch/>


      </a-col>
    </a-row>
  </div>
</template>

<script>
import { Options, Vue } from "vue-class-component";
import QuestionForSearch from "@/components/QuestionForSearch.vue";
import CardForSearch from "@/components/CardForSearch.vue";
const data = [
  {
    id: 1,
    title: "Ant Design Title 1",
    user:"akvfcdg",
    description:"dfghjklfcghjm,dfgbndghjkrewqwertyuiytr",
    likeNum :123,
    dislikeNum:4567,
    commentNum:7890,

  },
  {
    id: 2,
    title: "A7654 Title 1",
    user:"dfsfskang",
    description:"dfghjklfcghjm,dfgbndghjkrewqwertyuiytr",
    likeNum :123,
    dislikeNum:4567,
    commentNum:7890,
  },
  {
     id: 3,
    title: "gfds 1",
    user:"ererang",
    description:"dfghjklfcghjm,dfgbndghjkrewqwertyuiytr",
    likeNum :123,
    dislikeNum:4567,
    commentNum:7890,
  },
  {
     id: 4,
    title: "ytre",
    user:"akgfds",
    description:"dfghjklfcghjm,dfgbndghjkrewqwertyuiytr",
    likeNum :123,
    dislikeNum:4567,
    commentNum:7890,
  }
];

@Options({
  components: {
    QuestionForSearch,
     CardForSearch 
  }
})
export default class Search extends Vue {
  data() {
    return {
      data
    };
  }
}
// export default {
//   data() {
//     return {
//       data
//     };
//   }

// };
// export default class Search extends Vue {}
</script>

<style>
#search {
  min-height: 667px;
  background-color: #edeeed;
}
</style>