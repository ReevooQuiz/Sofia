<template>
  <div>
    <a-row>
      <a-col :span="19" :offset="2">
        <br />
        <a-row>
          <a-col :span="8">
            <a-input-search
              placeholder="输入封禁关键词"
              v-model:value="inputValue"
              @search="onSearch"
              style=" border-radius: 25px;box-shadow: 3px 3px 2px #dcdfdf"
            />
          </a-col>
        </a-row>

        <br />
        <br />
        <a-list :grid="{ gutter:8, column: 6 }" :data-source="data">
          <template #renderItem="{ item }">
            <a-list-item>
              <!-- <a-card :title="item.keyWord" :headStyle="tstyle"> -->
              <a-card bodyStyle="padding:10px">
                <!-- <template #extra> -->
                <a-row>
                  <a-col :span="14">
                    <span  style="font-size: 20px;
    font-weight: bold; overflow: hidden;
  white-space: nowrap;
 
  text-overflow: ellipsis;">{{item}}</span>
                  </a-col>

                  <a-col :span="9" :offset="1">
                    <a-button type="primary" @click="lift(item)">解禁</a-button>
                  </a-col>
                </a-row>

                <!-- </template> -->
                <!-- <span>已封禁 {{item.num}} 词条</span> -->
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
const data = ["s", "sdf", "fghjkxxxxxxxl"];

export default {
  components: { QuestionForSearch, CardForSearch },
  data() {
    return {
      inputValue: "",
      data,
      tstyle: { padding: "2 12px" }
    };
  },
  created() {
    this.searchValue = this.$route.query.content;
    this.inputValue = this.$route.query.content;
    console.log(this.inputValue);
  },
  methods: {
    handleInit(response) {
      this.questionData = response.data.questionData;
      this.cardInfo = response.data.cardInfo;
    },
    lift(item) {
      postRequest(
        "/wordBan",
        {
          word: item.keyWord,
          ban: false
        },
        this.handleLogin,
        {
          errorCallback: error => {
            console.log(error);
          }
        }
      );
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
</style>
