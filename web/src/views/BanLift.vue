<template>
  <div >
    <a-row>
      <a-col :span="19" :offset="2">
          <br/>
        <a-row>
          <a-col :span="8">
            <a-input-search
            id="ban-lift-input"
              placeholder="输入封禁关键词"
              v-model:value="inputValue"
              @search="onSearch"
              style="{'box-shadow': 5px 5px 10px gray}"
            />
          </a-col>

          
        </a-row>

        <br /><br/>
        <a-list :grid="{ gutter: 6, column: 6 }" :data-source="data">
          <template #renderItem="{ item }">
            <a-list-item>
              <a-card :title="item.keyWord" :headStyle="tstyle">
                <template #extra>
                 <a-button type="primary">解禁</a-button>
                </template>
                <span>已封禁 {{item.num}} 词条</span>
              </a-card>
            </a-list-item>
          </template>
        </a-list>
      </a-col>
    </a-row>
  </div>
</template>

<script>
import { Options, Vue } from "vue-class-component";
import QuestionForSearch from "@/components/QuestionForSearch.vue";
import CardForSearch from "@/components/CardForSearch.vue";
import server from "@/http/request.js";
const data = [
  {
    keyWord: "上海",
    num: 568
  },
  {
    keyWord: "交通",
    num: 568
  },
  {
    keyWord: "大学",
    num: 568
  },
   {
    keyWord: "上海",
    num: 568
  },
  {
    keyWord: "交通",
    num: 568
  },
  {
    keyWord: "大学",
    num: 568
  }
];

export default {
  components: { QuestionForSearch, CardForSearch },
  data() {
    return {
      inputValue: "wge",
      data,
      tstyle: { "padding":"2 12px"}
    };
  },
  created() {
    this.searchValue = this.$route.query.content;
    this.inputValue = this.$route.query.content;
    console.log(this.inputValue);
  },
  methods: {
    onSearch(value) {
      this.searchValue = value;
      console.log(this.searchValue);

      server
        .get("/search", {
          params: this.searchValue
        })
        .catch(function(error) {
          console.log(error);
        })
        .then(response => {
          console.log("!");
          console.log(response.data);
          console.log(response.data.question_list);
          this.questionData = response.data.question_list;
          console.log(this.questionData);
        });
    },

    handleInit(response) {
      this.questionData = response.data.questionData;
      this.cardInfo = response.data.cardInfo;
    }
  }
  // created: function() {
  //   server
  //     .get("/get", {
  //       params: {}
  //     })
  //     .then(response => this.handleInit(response))
  //     .catch(function(error) {
  //       console.log(error);
  //     });
  // }
};
</script>

<style >
body {
  height: 100%;
  background-color: #edeeed;
}

.ban-lift-input .ant-input-affix-wrapper {
  border: 1px solid #d9d9d9;
  border-radius: 16px;
  padding: 4px 11px;
  width: 100%;
  text-align: start;
  background-color: #fff;
  background-image: none;
  color: rgba(0, 0, 0, 0.65);
}

</style>
