<template>
  <div class="questionForSearch">
    <a-card hoverable :title="ques.title" size="small" style="border-radius : 3px">
      <template #extra>
        <!-- <a-tag color="#68b0af" style="margin-top:10px">
          <a-tooltip title="分区">
            <DatabaseOutlined />
            {{ques.category}}
          </a-tooltip>
        </a-tag> -->
         <a-button type="primary" style="margin-right:10px">
          <a-tooltip title="分区">
            <DatabaseOutlined />
            {{ques.category}}
          </a-tooltip>
        </a-button>
      </template>
      <a-row>
        <a-col :span="5">
          <div style="align-items: center">
          <img :src="ques.picture_urls" style="width:100%" @click="toQuestion" />
          </div>
        </a-col>
        <a-col :span="17 " :offset="1">
          <a-comment>
            <template #author>
              <a  @click="gotoPerson(ques.raiser.uid)">{{ques.raiser.nickname}}</a>
            </template>
            <template #avatar>
              <a-avatar @click="gotoPerson(ques.raiser.uid)" :src="ques.raiser.icon" :alt="ques.raiser.nickname" />
            </template>
            <template #content>
              <v-md-editor mode="preview" v-model="ques.head" @click="toQuestion"></v-md-editor>
            </template>
            <template #datetime>
              <a-tooltip :title="moment(ques.time).format('YYYY-MM-DD HH:mm:ss')">
                <span>{{ moment(ques.time).fromNow() }}</span>
              </a-tooltip>
            </template>

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
  FileTextOutlined,
  LikeOutlined,
  FireOutlined
} from "@ant-design/icons-vue";
export default {
  components: {
    FileTextOutlined,
    LikeOutlined,
    FireOutlined
  },
  props: ["ques"],

  data() {
    return {
      action: null,
      moment
    };
  },
  created() {
    console.log(this.ques);
  },
  methods: {
    toQuestion() {
      this.$router.push({
        path: "/question",
        query: { questionId: this.ques.qid }
      });
    },
    gotoPerson(id) {
      console.log(id)
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
    background-color: #acbab9;
}
.v-md-editor-preview {
    padding: 1px;
    word-break: break-all;
}
ol, ul, dl {
    margin-top: 0;
    margin-bottom: 0em;
}
.ant-comment-inner {
    display: flex;
    padding: 4px 0;
}
.ant-comment-actions {
    margin-top: 1px;
    padding-left: 0;
}
</style>
