<template>
  <div>
    <a-card
      hoverable
      :title="ques.question_title"
      size="small"
      style="border-radius : 3px"
      :headStyle="tstyle"
    >
      <template #extra>
        <a-button type="primary" style="margin-right:10px">
          <a-tooltip title="分区">
            <DatabaseOutlined />
            {{ques.category}}
          </a-tooltip>
        </a-button>
        <a-button
          v-if="this.status"
          type="primary"
          @click="cancle"
        >取消收藏</a-button>
         <a-button
          v-else
          type="primary"
          style="background-color:#fbbdbd;border-color: #ecc7d4;"
          @click="collect"
        >收藏</a-button>
      </template>

      <a-row>
        <a-col :span="5">
          <div style="align-items: center">
            <div v-for="(item) in ques.picture_urls" v-bind:key="item.index">
              <img :src="item" @click="toQuestion" style="width:100%" />
            </div>
          </div>
        </a-col>
        <a-col :span="18" :offset="1">
          <a-comment>
            <template #actions>
              <span key="comment-basic-reply-to">
                <a-tooltip title="赞">
                  <LikeOutlined />
                  {{ques.favorite_count}}
                </a-tooltip>
              </span>

              <span key="comment-basic-reply-to">
                <a-tooltip title="浏览量">
                  <FireOutlined />
                  {{ques.view_count}}
                </a-tooltip>
              </span>
              <span key="comment-basic-reply-to">
                <a-tooltip title="回答数">
                  <FileTextOutlined />
                  {{ques.answer_count}}
                </a-tooltip>
              </span>
              <!-- <span key="comment-basic-reply-to">
                <a-tooltip title="关注人数">
                  <TeamOutlined />
                  {{ques.follow_count}}
                </a-tooltip>
              </span>-->
            </template>
            <template #author>
              <a @click="gotoPerson(ques.raiser.uid)">{{ques.raiser.name}}</a>
            </template>
            <template #avatar>
              <a-avatar @click="gotoPerson(ques.raiser.uid)" :src="ques.raiser.icon" alt="avatar" />
            </template>
            <template #content>
              <v-md-editor mode="preview" v-model="ques.question_head" @click="toQuestion"></v-md-editor>
            </template>
            <!-- <template #content>
              <p @click="toQuestion">{{ques.head}}</p>
            </template>-->
            <!-- <template #datetime>
              <a-tooltip :title="moment(ques.time).format('YYYY-MM-DD HH:mm:ss')">
                <span>{{ moment(ques.time).fromNow() }}</span>
              </a-tooltip>
            </template>-->
          </a-comment>

          <a-divider />

          <a-tag
            v-for="(item) in ques.labels"
            v-bind:key="item.index"
            color="#88d5d1"
            style="margin-top:10px"
          >{{item}}</a-tag>
        </a-col>
      </a-row>
    </a-card>
    <br />
  </div>
</template>



<script >
import moment from "moment";
import {
  LikeFilled,
  LikeOutlined,
  DislikeFilled,
  DislikeOutlined,
  FireOutlined,
  FileTextOutlined,
  TeamOutlined,
  LeftCircleOutlined,
  RightCircleOutlined,
  DatabaseOutlined
} from "@ant-design/icons-vue";
import { putRequest } from "../http/request";
export default {
  components: {
    LikeFilled,
    LikeOutlined,
    DislikeFilled,
    DislikeOutlined,
    FireOutlined,
    TeamOutlined,
    FileTextOutlined,
    LeftCircleOutlined,
    RightCircleOutlined,
    DatabaseOutlined
  },
  props: ["ques"],

  data() {
    return {
      status: true,
      action: null,
      moment,
      modifiable: false,
      tstyle: { "font-size": "21px", "font-weight": " bold", color: " #425050" }
    };
  },
  created() {
    if (
      moment(this.ques.time).format("YYYY-MM-DD") >
      moment()
        .subtract(1, "days")
        .format("YYYY-MM-DD")
    )
      this.modifiable = true;
  },
  methods: {
    toQuestion() {
      console.log("??");
      console.log(this.ques.qid);
      this.$router.push({
        path: "/question",
        query: { questionId: this.ques.qid }
      });
    },

    cancle() {
      putRequest("/favorite", { qid: this.ques.qid, favorite: false }, res => {
        this.status = false;
      },{errorCallback:(e)=>{JSON.stringify(e)}});
    },
    collect() {
      putRequest("/favorite", { qid: this.ques.qid, favorite: true }, res => {
        this.status = true;
      },{errorCallback:(e)=>{JSON.stringify(e)}});
    },
    gotoPerson(id) {
      console.log(id);
      this.$router.push({
        path: "/personalSetOthers",
        query: { uId: id }
      });
    }
  }
};
</script>

<style >
.ant-divider-horizontal {
  display: block;
  clear: both;
  width: 100%;
  min-width: 100%;
  height: 1px;
  margin: 1px 0;
}
.v-md-editor-preview {
    padding: 6px;
    word-break: break-all;
}
/* .ant-comment-inner {
    display: flex;
    padding: 1px 0;
} */
.ant-carousel ::v-deep(.slick-slide) {
  text-align: center;
  height: 160px;
  line-height: 160px;
  background: #5e8884be;
  overflow: hidden;
}

.ant-carousel ::v-deep(.slick-arrow.custom-slick-arrow) {
  width: 25px;
  height: 25px;
  font-size: 25px;
  color: #fff;
  background-color: rgba(31, 45, 61, 0.11);
  opacity: 0.3;
  align-self: center;
}
.ant-carousel ::v-deep(.custom-slick-arrow:before) {
  display: none;
}
.ant-carousel ::v-deep(.custom-slick-arrow:hover) {
  opacity: 0.5;
}

.ant-carousel ::v-deep(.slick-slide h3) {
  color: #fff;
}
</style>
