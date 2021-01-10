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

        <a-list :loading="loading" item-layout="vertical" size="large" :data-source="messageData">
          <template #loadMore>
            <div
              v-if="showLoadingMore"
              :style="{ textAlign: 'center', marginTop: '12px', height: '32px', lineHeight: '32px' }"
            >
              <a-spin v-if="loadingMore" />
              <a-button v-else @click="onLoadMore">加载更多</a-button>
            </div>
            <div v-else  align="center" > 已经到底了 </div>
          </template>
          <template #renderItem="{ item, index }">
            <a-list-item key="item.type">
              <div class="message-block">
                <a-comment v-if="item.type===0">
                  <template #content>
                    <span>
                      你的问题
                      <a @click="toQuestion(item.qid)">『{{item.question_title}}』</a>
                      新增了 {{item.new_answer_count}} 个回答
                    </span>
                  </template>
                  <template #datetime>
                    <a-tooltip :title="moment(item.time).format('YYYY-MM-DD HH:mm:ss')">
                      <span>{{ moment(item.time).fromNow() }}</span>
                    </a-tooltip>
                  </template>
                </a-comment>
                <a-comment v-if="item.type===1">
                  <template #author>
                    <a @click="toQuestion(item.qid)">问题 : {{item.question_title}}</a>
                  </template>
                  <template #content>
                    <span>你的回答</span>
                    <br/>
                    <span>
                      <a @click="toQuestion(item.qid)">『{{item.answer_head}}』</a>
                    </span>
                     <br/>
                    <span>新增了 {{item.new_like_count}} 个赞</span>
                  </template>
                  <template #datetime>
                    <a-tooltip :title="moment(item.time).format('YYYY-MM-DD HH:mm:ss')">
                      <span>{{ moment(item.time).fromNow() }}</span>
                    </a-tooltip>
                  </template>
                </a-comment>
                <a-comment v-if="item.type===3">
                  <template #author>
                    <a @click="toQuestion(item.qid)">问题 : {{item.question_title}}</a>
                  </template>
                  <template #content>
                    <span>你的回答</span>
                    <br/>
                    <span>
                      <a @click="toQuestion(item.qid)">『{{item.answer_head}}』</a>
                    </span>
                     <br/>
                    <span>
                      新增了 {{item.new_comment_count}} 个评论
                    </span>
                  </template>
                  <template #datetime>
                    <a-tooltip :title="moment(item.time).format('YYYY-MM-DD HH:mm:ss')">
                      <span>{{ moment(item.time).fromNow() }}</span>
                    </a-tooltip>
                  </template>
                </a-comment>
                <a-comment v-if="item.type===4">
                  <template #author>
                    <a @click="toQuestion(item.qid)">问题 : {{item.question_title}}</a>
                  </template>
                  <template #content>
                    <span>你的回答</span>
                    <br/>
                    <span>
                      <a @click="toQuestion(item.qid)">『{{item.answer_head}}』</a>
                    </span>
                     <br/>
                    <span>
                      新增了 {{item.new_criticism_count}} 个反驳评论
                    </span>
                  </template>
                  <template #datetime>
                    <a-tooltip :title="moment(item.time).format('YYYY-MM-DD HH:mm:ss')">
                      <span>{{ moment(item.time).fromNow() }}</span>
                    </a-tooltip>
                  </template>
                </a-comment>
                <a-comment v-if="item.type===2">
                  <template #author>
                    <a @click="toQuestion(item.qid)">问题 : {{item.question_title}}</a>
                  </template>
                  <template #content>
                    <span>你的回答</span>
                    <br/>
                    <span>
                      <a @click="toQuestion(item.qid)">『{{item.answer_head}}』</a>
                    </span>
                     <br/>
                    <span>
                      新增了 {{item.new_approval_count}} 个赞同
                    </span>
                  </template>
                  <template #datetime>
                    <a-tooltip :title="moment(item.time).format('YYYY-MM-DD HH:mm:ss')">
                      <span>{{ moment(item.time).fromNow() }}</span>
                    </a-tooltip>
                  </template>
                </a-comment>
                <a-comment v-if="item.type===5">
                  <template #content>
                    <span>
                      你
                      新增了 {{item.new_follower_count}} 个粉丝
                    </span>
                  </template>
                  <template #datetime>
                    <a-tooltip :title="moment(item.time).format('YYYY-MM-DD HH:mm:ss')">
                      <span>{{ moment(item.time).fromNow() }}</span>
                    </a-tooltip>
                  </template>
                </a-comment>

               <a-comment v-if="item.type===6">
                  <template #content>
                    <span>
                      你关注的问题
                      <a @click="toQuestion(item.qid)">『{{item.question_title}}』</a>
                      新增了 {{item.new_answer_count}} 个回答
                    </span>
                  </template>
                  <template #datetime>
                    <a-tooltip :title="moment(item.time).format('YYYY-MM-DD HH:mm:ss')">
                      <span>{{ moment(item.time).fromNow() }}</span>
                    </a-tooltip>
                  </template>
                </a-comment>

                <a-comment v-if="item.type===7">
                  <template #content>
                    <span>
                      你关注的用户
                      <a @click="toPerson(item.uid)">『{{item.name}}』</a>
                      发布了新的问题 <a @click="toQuestion(item.qid)">『{{item.question_title}}』</a>
                    </span>
                  </template>
                  <template #datetime>
                    <a-tooltip :title="moment(item.time).format('YYYY-MM-DD HH:mm:ss')">
                      <span>{{ moment(item.time).fromNow() }}</span>
                    </a-tooltip>
                  </template>
                </a-comment>

                 <a-comment v-if="item.type===8">
                  <template #content>
                    <span>
                      你关注的用户
                      <a  @click="toPerson(item.uid)">『{{item.name}}』</a>
                      新增了 {{item.new_answer_count}} 个回答
                    </span>
                  </template>
                  <template #datetime>
                    <a-tooltip :title="moment(item.time).format('YYYY-MM-DD HH:mm:ss')">
                      <span>{{ moment(item.time).fromNow() }}</span>
                    </a-tooltip>
                  </template>
                </a-comment>
                
              </div>
            </a-list-item>
          </template>
        </a-list>
      </a-col>
    </a-row>

    <br />

    <br />
  </div>
</template>

<script>
import moment from "moment";
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
      messageData: [],
      loading: true,
      loadingMore: false,
      showLoadingMore: true,
      pageNum: 0,
      moment
    };
  },
  mounted() {
    this.getData(res => {
      console.log("!");
      if(res.result.length==0)
          this.showLoadingMore=false;
      console.log(res.result);
      this.loading = false;
      this.messageData = res.result;
      this.pageNum = this.pageNum + 1;
    });
  },
  methods: {
    getData(callback) {
      let uid = JSON.parse(sessionStorage.getItem("user")).uid;
      // this.loadingMore = true;

      getRequest("/notifications", callback, {
        errorCallback: e => {
          console.log(JSON.stringify(e));
        },
        params: { id: uid, page: this.pageNum }
      });
    },

    onLoadMore() {
      this.getData(res => {
        this.messageData = this.messageData.concat(res.result);
        this.loadingMore = false;
        if(res.result.length==0)
          this.showLoadingMore=false;
        this.pageNum = this.pageNum + 1;
        this.$nextTick(() => {
          window.dispatchEvent(new Event("resize"));
        });
      });
    },
    toQuestion(id) {
      this.$router.push({ path:'/question' , query: { questionId:id } });
    },
    toPerson(id) {
      console.log(id)
      this.$router.push({
        path: "/personalSetOthers",
        query: { uId: id }
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

.message-block {
  top: 18%;
  /* position: fixed; */
  left: 27%;
  /* text-align: center; */
  background-color: #ffffff;
  width: 700px;
  border-radius: 40px;
  padding: 10px;
  margin: 0 auto;
  box-shadow: 3px 3px 2px #dcdfdf;
}
</style>