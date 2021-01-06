<template>
  <div>
    <a-card
      hoverable
      :title="ans.question.title"
      size="small"
      style="border-radius : 3px"
      :headStyle="tstyle"
    >
      <template #extra>
        <a-tag color="#68b0af" style="margin-top:10px">
          <a-tooltip title="分区">
            <DatabaseOutlined />
            {{ans.question.category}}
          </a-tooltip>
        </a-tag>
      </template>

      <a-row>
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
              <div v-for="(item) in ans.answer.pictureUrls" v-bind:key="item.index">
                <img :src="item" style="width:100%" />
              </div>
            </a-carousel>
          </div>
        </a-col>
        <a-col :span="18" :offset="1">
          <a-comment>
            <template #actions>
              <!-- <span key="comment-basic-like">
                <a-tooltip title="赞">
                  <template v-if="action === 'liked'">
                    <LikeFilled @click="like" />
                    {{ans.answer.like_count}}
                  </template>
                  <template v-else>
                    <LikeOutlined @click="like" />
                    {{ans.answer.like_count}}
                  </template>
                </a-tooltip>
                <span style="padding-left: '8px';cursor: 'auto'">{{ likes }}</span>
              </span>-->
              <!-- <span key="comment-basic-dislike">
                <a-tooltip title="反对">
                  <template v-if="action === 'disliked'">
                    <DislikeFilled @click="dislike" />
                    {{ans.answer.criticism_count}}
                  </template>
                  <template v-else>
                    <DislikeOutlined @click="dislike" />
                    {{ans.answer.criticism_count}}
                  </template>
                </a-tooltip>
                <span style="padding-left: '8px';cursor: 'auto'">{{ dislikes }}</span>
              </span>-->
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
            <!-- <template #author>
              <a>{{ans.owner.user_name}}</a>
            </template>
            <template #avatar>
              <a-avatar :src="ans.owner.user_icon" alt="avatar" />
            </template>-->
            <template #content>
              <p>{{ans.answer.head}}</p>
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
  FrownOutlined,
  SmileOutlined
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
    SmileOutlined
  },
  props: ["ans"],

  data() {
    return {
      likes: 0,
      dislikes: 0,
      action: null,
      moment,
      tstyle: { "font-size": "21px", "font-weight": " bold", color: " #425050" }
    };
  },
  methods: {
    like() {
      this.likes = 1;
      this.dislikes = 0;
      this.action = "liked";
    },
    dislike() {
      this.likes = 0;
      this.dislikes = 1;
      this.action = "disliked";
    }
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