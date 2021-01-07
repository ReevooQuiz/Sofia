<template>
  <div>
    <a-card
      hoverable
      :title="ques.title"
      size="small"
      style="border-radius : 3px"
      :headStyle="tstyle"
      v-if="!deleted"
    >
      <template #extra>
        <a-tag color="#68b0af" style="margin-top:10px">
          <a-tooltip title="分区">
            <DatabaseOutlined />
            {{ques.category}}
          </a-tooltip>
        </a-tag>
        <a-space size="small">
          <a-button size="small" shape="pill" type="primary" v-if="modifiable" @click="onModify">修改问题</a-button>
          <a-button v-if="!ques.closed" @click="onClose" size="small" shape="pill" type="primary">关闭问题</a-button>
          <a-avatar @click="onDelete" size="small" shape="round" type="primary" style="background-color:#fbbdbd;border-color: #ecc7d4;"><DeleteOutlined /></a-avatar>

        </a-space>
      </template>

      <a-row  @click="toQuestion">
        <a-col :span="4" >
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
              <span key="comment-basic-reply-to">
                <a-tooltip title="关注人数">
                  <TeamOutlined />
                  {{ques.follow_count}}
                </a-tooltip>
              </span>
            </template>
            <template #content>
              <v-md-editor mode="preview" v-model="ques.head" @click="toQuestion"></v-md-editor>
            </template>
            <template #datetime>
              <a-tooltip :title="moment(ques.time).format('YYYY-MM-DD HH:mm:ss')">
                <span>{{ moment(ques.time).fromNow() }}</span>
              </a-tooltip>
            </template>
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
  DatabaseOutlined,
    DeleteOutlined,
    CloseOutlined
} from "@ant-design/icons-vue";
import {postRequest,putRequest} from "../http/request";
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
    DeleteOutlined,
    CloseOutlined

  },
  props: ["ques"],

  data() {
    return {
      likes: this.ques.likeNum,
      dislikes: this.ques.dislikeNum,
      action: null,
      moment,
      modifiable:false,
      tstyle: { "font-size": "21px", "font-weight": " bold", color: " #425050" },
      deleted:false
    };
  },
  created() {
    if (moment(this.ques.time).format('YYYY-MM-DD')>moment().subtract(1, 'days').format('YYYY-MM-DD'))
      this.modifiable=true;
  },
  methods: {
    toQuestion() {
      this.$router.push({
        path: "/question",
        query: { questionId: this.ques.qid }
      });
    },
    onModify(){
      this.$router.push({ path:'/postQuestion' , query: { questionId: this.ques.qid } });
    },
    onDelete(){
      postRequest("/disable_question", {qid:this.ques.qid},(e)=>{
        console.log(e);
        if (e.code==0){
          this.deleted=true;
        }
      },{errorCallback:(e)=>{
          console.log(e);
        }});
    },
    onClose(){
      this.ques.closed=true;
      putRequest("/disable_question", {qid:this.ques.qid},(e)=>{
        console.log(e);
        if (e.code==0){
        }
      },{errorCallback:(e)=>{
          console.log(e);
        }});
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
