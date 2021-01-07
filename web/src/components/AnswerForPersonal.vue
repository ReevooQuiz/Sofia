<template>
  <div>
    <a-card
      hoverable
      :title="ans.question.title"
      size="small"
      style="border-radius : 3px"
      :headStyle="tstyle"
      v-if="!deleted"
    >
      <template #extra>
          <a-tag color="#68b0af" style="margin-top:10px">
            <a-tooltip title="分区">
              <DatabaseOutlined />
              {{ans.question.category}}
            </a-tooltip>
          </a-tag>
          <a-button v-if="modifying" size="small" shape="pill" type="primary" @click="onModify">取消修改</a-button>
          <a-space v-else size="small">
            <a-button size="small" shape="pill" type="primary" @click="onModify">修改回答</a-button>
            <a-avatar @click="onDelete" size="small" shape="round" type="primary" style="background-color:#fbbdbd;border-color: #ecc7d4;"><DeleteOutlined /></a-avatar>
          </a-space>
      </template>
      <a-row v-if="!modifying" @click="toQuestion" >
        <a-col :span="4">
          <div style="align-items: center">
            <a-carousel arrows>
              <template #prevArrow>
                <div class="custom-slick-arrow" style="left: 10px;zIndex: 1">
                  <left-circle-outlined />
                </div>
              </template>
              <template #nextArrow>
                <div class="custom-slick-arrow" style="right: 10px">
                  <right-circle-outlined />
                </div>
              </template>
              <div v-for="(item) in ans.answer.picture_urls" v-bind:key="item.index">
                <img :src="item" style="width:100%" />
              </div>
            </a-carousel>
          </div>
        </a-col>
        <a-col :span="18" :offset="1">
          <a-comment>
            <template #actions>
              <span key="comment-basic-reply-to">
                <a-tooltip title="点赞数">
                  <LikeOutlined />
                  {{ans.answer.like_count}} 点赞
                </a-tooltip>
              </span>
              <span key="comment-basic-reply-to">
                <a-tooltip title="反对数">
                  <FrownOutlined />

                  {{ans.answer.criticism_count}} 反对
                </a-tooltip>
              </span>
              <span key="comment-basic-reply-to">
                <a-tooltip title="赞同数">
                  <SmileOutlined />
                  {{ans.answer.approval_count}} 赞同
                </a-tooltip>
              </span>

              <span key="comment-basic-reply-to">
                <a-tooltip title="评论数">
                  <FileTextOutlined />
                  {{ans.answer.comment_count}} 评论
                </a-tooltip>
              </span>
            </template>
            <template #content>
              <v-md-editor mode="preview" v-model="ans.answer.head"></v-md-editor>
            </template>
            <template #datetime>
              <a-tooltip :title="moment(ans.answer.time).format('YYYY-MM-DD HH:mm:ss')">
                <span>{{ moment(ans.answer.time).fromNow() }}</span>
              </a-tooltip>
            </template>
          </a-comment>

          <a-divider />

          <a-tag
            v-for="(item) in ans.question.labels"
            v-bind:key="item.index"
            color="#88d5d1"
            style="margin-top:10px"
          >{{item}}</a-tag>
        </a-col>
      </a-row>
      <a-row v-else>
        <a-col :span="24" >
          <v-md-editor v-model="answerValue" height="400px"></v-md-editor>
          <br/>
          <a-row type="flex" justify="end">
            <a-button @click="onCommitAnswer" type="primary" shape="pill" size="small">提交回答</a-button>
          </a-row>
        </a-col>
      </a-row>
    </a-card>
    <br />
  </div>
</template>



<script >
import moment from "moment";
import { putRequest,getRequest } from "@/http/request.js";
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
  DatabaseOutlined,
  FrownOutlined,
  SmileOutlined,
  DeleteOutlined
} from "@ant-design/icons-vue";
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
    DatabaseOutlined,
    FrownOutlined,
    SmileOutlined,
    DeleteOutlined
  },
  props: ["ans"],

  data() {
    return {
      likes: 0,
      dislikes: 0,
      action: null,
      moment,
      tstyle: { "font-size": "21px", "font-weight": " bold", color: " #425050" },
      deleted:false,
      modifying:false,
      answerValue:""
    };
  },
  methods: {
    onModify(){
      this.modifying=!this.modifying;
      getRequest("/answer", (e) => {
        this.answerValue = e.result.content;
      }, {
        errorCallback: (e) => {
          console.log(e)
        },
        params: {aid: this.ans.answer.aid}
      });
    },
    onDelete(){
      postRequest("/delete_answer", {aid:this.ans.answer.aid},(e)=>{
        console.log(e);
        if (e.code==0){
          this.deleted=true;
        }
      },{errorCallback:(e)=>{
          console.log(e);
        }});
    },
    onCommitAnswer(){
      putRequest("/answers", {aid:this.ans.answer.aid,content:this.answerValue},(e)=>{
        console.log(e);
        if (e.code==0){
          this.ans.answer.head=this.answerValue;
          this.answerValue="";
          this.modifying=false;
        }
      },{errorCallback:(e)=>{
          console.log(e);
        }});
    },
    toQuestion() {
      this.$router.push({
        path: "/question",
        query: { questionId: this.ans.question.qid }
      });
    },
  }
};
</script>

<style scoped>
.ant-divider-horizontal {
  display: block;
  clear: both;
  width: 100%;
  min-width: 100%;
  height: 1px;
  margin: 1px 0;
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
