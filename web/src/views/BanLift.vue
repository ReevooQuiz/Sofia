<template>
  <div>
    <a-row>
      <a-col :span="19" :offset="2">
        <br />
        <a-row>
          <!-- <a-col :span="6">
            <a-input-search
              placeholder="输入查询关键词"
              v-model:value="inputValue"
              @search="onSearch"
              style=" border-radius: 25px;box-shadow: 3px 3px 2px #dcdfdf"
            />
          </a-col>-->
          <a-col :span="6">
            <a-input-search
              placeholder="输入封禁关键词"
              v-model:value="banValue"
              enter-button="封禁"
              @search="ban"
            />
            <!-- style=" border-radius: 25px;box-shadow: 3px 3px 2px #dcdfdf" -->
          </a-col>
        </a-row>

        <br />
        <br />
        <a-list
          :grid="{ gutter:8, column: 6 }"
          :data-source="data"
          :loading="loading"
          item-layout="vertical"
        >
          <template #loadMore>
            <div
              v-if="showLoadingMore"
              :style="{ textAlign: 'center', marginTop: '12px', height: '32px', lineHeight: '32px' }"
            >
              <a-spin v-if="loadingMore" />
              <a-button v-else @click="onLoadMore">加载更多</a-button>
            </div>
          </template>

          <template #renderItem="{ item, index  }">
            <a-list-item>
              <!-- <a-card :title="item.keyWord" :headStyle="tstyle"> -->
              <a-card bodyStyle="padding:10px">
                <!-- <template #extra> -->
                <a-row>
                  <a-col :span="14">
                    <span
                      style="font-size: 20px;
    font-weight: bold; overflow: hidden;
  white-space: nowrap;

  text-overflow: ellipsis;"
                    >{{item.word}}</span>
                  </a-col>

                  <a-col :span="9" :offset="1">
                    <a-button v-if="item.ban===true" type="primary" @click="lift(item.word)">解禁</a-button>
                    <a-button
                      v-if="item.ban===false"
                      type="primary"
                      style="background-color:#fbbdbd;border-color: #ecc7d4;
"
                      @click="ban2(item.word)"
                    >封禁</a-button>
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

import { putRequest, getRequest } from "../http/request";
import { message } from "ant-design-vue";
export default {
  components: { QuestionForSearch, CardForSearch },
  data() {
    return {
      loading: true,
      loadingMore: false,
      showLoadingMore: true,

      inputValue: "",
      data: [],
      tstyle: { padding: "2 12px" },
      banValue: "",
      page: 0
    };
  },
  mounted() {
    this.$store.commit('changeTarget','ban');
    this.getData(res => {
      console.log("!");
      console.log(res.result);
      this.loading = false;
      let r = [];
      let newData = res.result;
      for (var i = 0; i < newData.length; i++) {
        r.push({ ban: true, word: newData[i] });
      }
      this.data = r;
      this.page = this.page + 1;
    });
  },
  methods: {
    lift(word) {
      putRequest(
        "/wordBan",
        {
          word: word,
          ban: false
        },
        res => {
          if (res.code === 0) {
            // 解禁成功
            console.log(word);

            for (var i = 0; i < this.data.length; i++) {
              if (this.data[i].word == word) {
                message.success("解禁成功");
                this.data[i].ban = false;
                break;
              }
            }

            console.log(this.data);
          } else {
            message.error("解禁失败");
          }
        },
        {
          errorCallback: error => {
            console.log(error);
          }
        }
      );
    },
    getData(callback) {
      getRequest("/wordsBanned", callback, {
        errorCallback: e => {
          JSON.stringify(e);
        },
        params: { page: this.page }
      });
    },

    onLoadMore() {
      this.getData(res => {
        let r = [];
        let newData = res.result;
        for (var i = 0; i < newData.length; i++) {
          r.push({ ban: true, word: newData[i] });
        }
        this.data = this.data.concat(r);

        this.loadingMore = false;
        this.page = this.page + 1;
        this.$nextTick(() => {
          window.dispatchEvent(new Event("resize"));
        });
      });
    },
    ban2(word) {
      let flag = true;

      if (flag) {
        putRequest(
          "/wordBan",
          {
            word: word,
            ban: true
          },
          res => {
            if (res.code === 0) {
              // 禁用成功
              message.success("禁用成功");

              for (var i = 0; i < this.data.length; i++) {
                if (this.data[i].word == word) {
                  message.success("禁用成功");
                  this.data[i].ban = true;
                  break;
                }
              }
            } else {
              message.error("禁用失败");
            }
          },
          {
            errorCallback: error => {
              console.log(error);
            }
          }
        );
      }
    },
    ban() {
      let flag = true;
      if (this.banValue === "") {
        flag = false;
        message.error("请输入关键字");
      }
      if (this.banValue.length > 5) {
        flag = false;
        message.error("关键字不超过五个字");
      }
      for (var i = 0; i < this.data.length; i++) {
        if (this.data[i].word === this.banValue) {
          message.error("该禁用关键字已存在");
          flag = false;
          break;
        }
      }
      if (flag) {
        putRequest(
          "/wordBan",
          {
            word: this.banValue,
            ban: true
          },
          res => {
            if (res.code === 0) {
              // 禁用成功
              message.success("禁用成功");
              this.data.push({ban:true,word:this.banValue});
              this.banValue = "";
            } else {
              message.error("禁用失败");
            }
          },
          {
            errorCallback: error => {
              console.log(error);
            }
          }
        );
      }
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
