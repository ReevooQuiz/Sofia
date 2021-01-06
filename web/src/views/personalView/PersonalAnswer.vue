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
const data = [
  {
    question: {
      qid: "234",
      title: "Favourite programming language?",
      category: "study",
      labels: ["programming"],
      head: "What if we put"
    },
    answer: {
      aid: "234",
      like_count: 2,
      criticism_count: 4,
      approval_count: 2,
      comment_count: 2,
      head:
        "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打\
      近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前f",
      time: "2015-08-05T08:40:51.620Z",
      pictureUrls: [
        "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png",
        "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"
      ]
    }
  },
  {
    question: {
      qid: "234",
      title: "Favourite programming language?",
      category: "study",
      labels: ["programming"],
      head: "What if we put"
    },
    answer: {
      aid: "234",
      like_count: 2,
      criticism_count: 4,
      approval_count: 2,
      comment_count: 2,
      head:
        "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打\
      近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前f",
      time: "2015-08-05T08:40:51.620Z",
      pictureUrls: [
        "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png",
        "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"
      ]
    }
  },
  {
    question: {
      qid: "234",
      title: "Favourite programming language?",
      category: "study",
      labels: ["programming"],
      head: "What if we put"
    },
    answer: {
      aid: "234",
      like_count: 2,
      criticism_count: 4,
      approval_count: 2,
      comment_count: 2,
      head:
        "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打\
      近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前f",
      time: "2015-08-05T08:40:51.620Z",
      pictureUrls: [
        "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png",
        "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"
      ]
    }
  },
  {
    question: {
      qid: "234",
      title: "Favourite programming language?",
      category: "study",
      labels: ["programming"],
      head: "What if we put"
    },
    answer: {
      aid: "234",
      like_count: 2,
      criticism_count: 4,
      approval_count: 2,
      comment_count: 2,
      head:
        "近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前往花店协商未果，拍照发差评却被打\
      近日，有消息称上海一花店老板因差评进校骚扰上海交通大学密西根学院学生一事引发关注。当事学生在网上发帖称其买到的花与预定样子不符，前f",
      time: "2015-08-05T08:40:51.620Z",
      pictureUrls: [
        "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png",
        "https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"
      ]
    }
  }
];
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
      loading: true,
      loadingMore: false,
      showLoadingMore: true,
      pageNum: 0
    };
  },
  mounted() {
    this.getData(res => {
      console.log("!");
      console.log(res.result);
      this.loading = false;
      this.answerData = res.result;
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
        params: { id: uid, page: this.pageNum }
      });
    },
    onLoadMore() {
      this.getData(res => {
        this.answerData = this.answerData.concat(res.result);
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