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

        <a-list :loading="loading" item-layout="vertical" size="large" :data-source="answerData">
          <template #loadMore>
            <div
              v-if="showLoadingMore"
              :style="{ textAlign: 'center', marginTop: '12px', height: '32px', lineHeight: '32px' }"
            >
              <a-spin v-if="loadingMore" />
              <a-button v-else @click="onLoadMore">加载更多</a-button>
            </div>
            <div v-else align="center">已经到底了</div>
          </template>
          <template #renderItem="{ item, index }">
            <a-list-item key="item.title">
              <AnswerForPersonal :ans="item" />
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
import { getRequest } from "@/http/request.js";
import { Options, Vue } from "vue-class-component";
import {
  UserOutlined,
  StarOutlined,
  LikeOutlined,
  MessageOutlined
} from "@ant-design/icons-vue";
import SubMenu from "../../components/PersonalNavigation";
import AnswerForPersonal from "../../components/AnswerForPersonal";
export default {
  components: {
    UserOutlined,
    SubMenu,
    AnswerForPersonal,
    StarOutlined,
    LikeOutlined,
    MessageOutlined
  },
  data() {
    return {
      answerData: [],
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
      // this.loading = false;
      this.answerData = res.result;
      this.loadingMore=false;
      this.pageNum = this.pageNum + 1;
    });
  },
  methods: {
    getData(callback) {
      let uid = JSON.parse(sessionStorage.getItem("user")).uid;
      // this.loadingMore = true;

      getRequest("/userAnswers", callback, {
        errorCallback: e => {
          console.log(JSON.stringify(e));
        },
        params: { uid: uid, page: this.pageNum }
      });
    },
    onLoadMore() {
      this.loadingMore=true;
      this.getData(res => {
        this.answerData = this.answerData.concat(res.result);
        if(res.result.length==0)
          this.showLoadingMore=false;
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