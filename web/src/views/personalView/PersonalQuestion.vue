<template>
  <div>
    <a-row>
      <a-col :span="1" :offset="1">
        <div class="personal-nav">
          <SubMenu />
        </div>
      </a-col>

      <a-col :span="15" :offset="2">
        <br />
        <br />

        <a-list :loading="loading" item-layout="vertical" size="large" :data-source="questionData">
          <template #loadMore>
            <div
              v-if="showLoadingMore"
              :style="{ textAlign: 'center', marginTop: '12px', height: '32px', lineHeight: '32px' }"
            >
              
              <a-spin v-if="loadingMore" />
            <a-button v-else @click="onLoadMore">
              加载更多
            </a-button>
              
            </div>
            <div v-else  align="center">已经到底了</div>
          </template>
          <template #renderItem="{ item, index }">
            <a-list-item key="item.title">
              <QuestionForPersonal :ques="item" />
              {{ item.content }}
            </a-list-item>
          </template>
        </a-list>
        <!-- <QuestionForPersonal v-for="(item) in questionData" v-bind:key="item.id" :ques="item" /> -->
      </a-col>
    </a-row>

    <br />

    <br />
  </div>
</template>

<script>
import { defineComponent } from "vue";
import { Options, Vue } from "vue-class-component";
import { getRequest } from "@/http/request.js";
import {
  UserOutlined,
  StarOutlined,
  LikeOutlined,
  MessageOutlined
} from "@ant-design/icons-vue";
import SubMenu from "../../components/PersonalNavigation";
import QuestionForPersonal from "../../components/QuestionForPersonal";

export default {
  components: {
    UserOutlined,
    SubMenu,
    QuestionForPersonal,
    StarOutlined,
    LikeOutlined,
    MessageOutlined
  },
  data() {
    return {
      questionData: [],
      // loading: true,
      loadingMore:true,
      showLoadingMore: true,
      pageNum: 0
    };
  },

  mounted() {
    this.getData(res => {
      console.log("!");
      console.log(res.result);
      if(res.result.length==0)
          this.showLoadingMore=false;
      this.loading = false;
      this.questionData = res.result;
      this.pageNum = this.pageNum + 1;
      this.loadingMore=false;
    });
  },
  methods: {
    getData(callback) {
      let uid = JSON.parse(sessionStorage.getItem("user")).uid;
      // this.loadingMore = true;

      getRequest("/userQuestions", callback, {
        errorCallback: e => {
          console.log(JSON.stringify(e));
        },
        params: { uid: uid, page: this.pageNum }
      });
    },

    onLoadMore() {
      this.loadingMore=true;
      this.getData(res => {
         if(res.result.length==0)
          this.showLoadingMore=false;
        this.questionData = this.questionData.concat(res.result);
        this.loadingMore = false;
        this.pageNum = this.pageNum + 1;
        this.$nextTick(() => {
          window.dispatchEvent(new Event("resize"));
        });
      });
    }
  }
};
</script>

<style>
body {
  height: 100%;
  background-color: #edeeed;
}

.personal-nav .ant-menu-inline {
  width: 40%;
}
</style>