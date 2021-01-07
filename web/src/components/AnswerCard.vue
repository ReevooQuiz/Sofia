<template>
  <div id="answerCard">
    <a-card  size="small" style="border-radius : 3px">
      <a-row>
        <a-col :span="22" :offset="1">
          <a-comment>
            <template #actions>
              <span key="comment-basic-approve" v-if="ans.approvable">
                <a-tooltip title="Approve">
                  <template v-if="ans.approved">
                    <LikeFilled @click="onApprove" />
                  </template>
                  <template v-else>
                    <LikeOutlined @click="onApprove" />
                  </template>
                </a-tooltip>
                <span style="padding-left: '8px';cursor: 'auto'">{{ ans.approval_count }}</span>
              </span>
              <span key="comment-basic-like">
                <a-tooltip title="Like">
                  <template v-if="ans.liked">
                    <HeartFilled @click="onLike" />
                  </template>
                  <template v-else>
                    <HeartOutlined @click="onLike" />
                  </template>
                </a-tooltip>
                <span style="padding-left: '8px';cursor: 'auto'">{{ ans.like_count }}</span>
              </span>

              <span key="comment-basic-comment" @click="clickComment" >
                <MessageTwoTone twoToneColor="#88d5d1" v-if="showComment" />
                <MessageOutlined v-else />
                <span style="padding-left: '8px';cursor: 'auto'">{{ ans.comment_count }}</span>
              </span>
              <span v-if="!writeCriticism">
                <a-button v-if="writeComment" @click="onWriteComment" type="primary"
                          style="background-color:#fbbdbd;border-color: #ecc7d4;" shape="pill" size="small">取消评论</a-button>
                <a-button v-else @click="onWriteComment" type="primary" shape="pill" size="small">我要评论</a-button>
              </span>
              <span v-if="!writeComment">
                <a-button v-if="writeCriticism" @click="onWriteCriticism" type="primary"
                          style="background-color:#fbbdbd;border-color: #ecc7d4;" shape="pill" size="small">取消反驳</a-button>
                <a-button v-else @click="onWriteCriticism" type="primary"
                          style="background-color:#fbbdbd;border-color: #ecc7d4;" shape="pill"
                size="small">我要反驳</a-button>
              </span>
            </template>
            <template #author>
              <a @click="gotoPerson(ans.answerer.uid)"> {{ans.answerer.name}}</a>
            </template>
            <template #avatar>
              <a-avatar
              @click="gotoPerson(ans.answerer.uid)"
                :src="ans.answerer.icon"
                :alt="ans.answerer.name"
              />
            </template>
            <template #content>
              <p v-if="full">
                <v-md-editor mode="preview" v-model="ans.content"></v-md-editor>
              </p>
              <p v-else>
                <v-md-editor mode="preview" v-model="ans.head"></v-md-editor>
                <a @click="getAnswerDetail">...查看全部</a>
              </p>
            </template>
            <template #datetime>
              <a-tooltip :title="time.format('YYYY-MM-DD HH:mm:ss')">
                <span>{{ time.fromNow() }}</span>
              </a-tooltip>
            </template>
          </a-comment>
          <span v-if="writeComment||writeCriticism">
              <a-textarea v-model:value="writeCommentValue" placeholder="您的评论或反驳..." :rows="4" />
              <br/>
              <a-row type="flex" justify="end">
                <a-button v-if="writeComment" @click="onCommitComment" type="primary" shape="pill" size="small">提交评论</a-button>
                <a-button v-if="writeCriticism" @click="onCommitCriticism" type="primary" shape="pill" size="small">提交反驳</a-button>
              </a-row>
            </span>
          <span v-if="showComment">
            <a-tabs :default-active-key="tabNow" @change="onChangeTab" size="small">
              <a-tab-pane key="1" tab="评论">
                <a-list
                    class="comment-list"
                    item-layout="horizontal"
                    :data-source="comments"
                >
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
                  <div v-else>已经到底了</div>
                </template>
                <template #renderItem="{ item }">
                  <a-list-item>
                    <a-comment :author="item.nickname" :avatar="item.icon">
                      <template #content>
                        <p>
                          {{ item.content }}
                        </p>
                      </template>
                      <template #datetime>
                        <a-tooltip :title="time.format('YYYY-MM-DD HH:mm:ss')">
                         <span>{{ time.fromNow() }}</span>
                        </a-tooltip>
                     </template>

                    </a-comment>
                  </a-list-item>
                </template>
                </a-list>
              </a-tab-pane>
              <a-tab-pane key="2" tab="反驳">
                <a-list
                    class="comment-list"
                    item-layout="horizontal"
                    :data-source="criticisms"
                >
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
                  <div v-else>已经到底了</div>
                </template>
                <template #renderItem="{ item }">
                  <a-list-item>
                    <a-comment :author="item.nickname" :avatar="item.icon">
                      <template #content>
                        <p>
                          {{ item.content }}
                        </p>
                      </template>
                      <template #datetime>
                        <a-tooltip :title="time.format('YYYY-MM-DD HH:mm:ss')">
                         <span>{{ time.fromNow() }}</span>
                        </a-tooltip>
                     </template>

                    </a-comment>
                  </a-list-item>
                </template>
                </a-list>
              </a-tab-pane>
            </a-tabs>

            </span>
        </a-col>
      </a-row>
    </a-card>
    <br/>
  </div>
</template>



<script >
import moment from "moment";
import { LikeFilled, LikeOutlined,MessageOutlined,MessageTwoTone,HeartOutlined,HeartFilled } from '@ant-design/icons-vue';
import { notification } from 'ant-design-vue';
import { postRequest,getRequest } from "@/http/request.js";

export default {
  components: {
     LikeFilled, LikeOutlined,MessageOutlined,MessageTwoTone,HeartOutlined,HeartFilled
  },
  props: ['ans'] ,

  data() {
    return {
      action: null,
      showComment:false,
      comments:[],
      criticisms:[],
      showLoadingMore: true,
      writeComment:false,
      writeCriticism:false,
      writeCommentValue:"",
      full:false,
      pageNow:0,
      pageNow2:0,
      time:null,
      loadingMore:true,
      tabNow:1
    };
  },
  created(){
    this.time=moment(this.ans.time);
  },
  methods: {
    getAnswerDetail() {
      getRequest("/answer", (e) => {
        this.ans.content = e.result.content;
        this.full = true;
      }, {
        errorCallback: (e) => {
          console.log(e)
        },
        params: {aid: this.ans.aid}
      });
    },
    onLike() {
      this.ans.liked = !this.ans.liked;
      if (this.ans.liked)
        this.ans.like_count++;
      else this.ans.like_count--;
      postRequest("/like", {aid: this.ans.aid, like: this.ans.liked}, (e) => {
        console.log(e);
      }, {
        errorCallback: (e) => {
          console.log(e);
        }
      });
    },
    onApprove() {
      this.ans.approved = !this.ans.approved;
      if (this.ans.approved)
        this.ans.approval_count++;
      else this.ans.approval_count--;
      postRequest("/approve", {aid: this.ans.aid, approve: this.ans.approved}, (e) => {
        console.log(e);
      }, {
        errorCallback: (e) => {
          console.log(e);
        }
      });
    },
    clickComment() {
      this.showComment = !this.showComment;
      if (this.showComment) {
        if (this.pageNow == 0) {
          this.getData();
        }
      }
    },
    getData() {
      this.loadingMore=true;
      getRequest("/comments",
          (response) => {
            this.comments = this.comments.concat(response.result);
            this.pageNow++;
            this.loadingMore = false;
            if (response.result.length == 0)
              this.showLoadingMore = false;
          },
          {
            errorCallback: (e) => {
              console.log(e);
            },
            params: {aid: this.ans.aid, page: this.pageNow}
          });
    },
    getDataCriticism() {
      this.loadingMore=true;
      getRequest("/criticisms",
          (response) => {
            this.criticisms = this.criticisms.concat(response.result);
            this.pageNow2++;
            this.loadingMore = false;
            if (response.result.length == 0)
              this.showLoadingMore = false;
          },
          {
            errorCallback: (e) => {
              console.log(e);
            },
            params: {aid: this.ans.aid, page: this.pageNow2}
          });
    },
    onLoadMore() {
      this.loadingMore = true;
      if (this.tabNow==1)
        this.getData();
      else this.getDataCriticism();
    },
    onWriteComment() {
      if (!this.writeComment)
        this.writeComment = true;
      else {
        this.writeComment = false;
        this.writeCommentValue = "";
      }
    },
    onCommitComment() {
      if (this.writeCommentValue.length == 0) {
        notification.open({
          message: '您的评论不能为空',
        });
        return;
      }
      if (this.writeCommentValue.length > 150) {
        notification.open({
          message: '您的评论过长',
          description:
              '评论的最大长度为150字',
        });
        return;
      }
      let body = {aid: this.ans.aid, content: this.writeCommentValue};
      postRequest("/comments", body, (e) => {
        console.log(e);
        this.writeComment = false;
        this.writeCommentValue = "";
        this.ans.comment_count++;
        if (this.showComment) {
          this.comments = [];
          this.pageNow = 0;
          this.getData();
        }
      }, {
        errorCallback: (e) => {
          console.log(e);
        }
      });
    },
    onWriteCriticism() {
      if (!this.writeCriticism)
        this.writeCriticism = true;
      else {
        this.writeCriticism = false;
        this.writeCriticismValue = "";
      }
    },
    onCommitCriticism() {
      if (this.writeCommentValue.length == 0) {
        notification.open({
          message: '您的反驳不能为空',
        });
        return;
      }
      if (this.writeCommentValue.length > 150) {
        notification.open({
          message: '您的反驳过长',
          description:
              '反驳的最大长度为150字',
        });
        return;
      }
      let body = {aid: this.ans.aid, content: this.writeCommentValue};
      postRequest("/criticisms", body, (e) => {
        console.log(e);
        this.writeCriticism = false;
        this.writeCommentValue = "";
        this.ans.criticism_count++;
        if (this.showComment) {
          this.criticisms = [];
          this.pageNow2 = 0;
          this.getData();
        }
      }, {
        errorCallback: (e) => {
          console.log(e);
        }
      });
    },
    gotoPerson(id) {
      console.log(id)
      this.$router.push({
        path: "/personalSetOthers",
        query: {uId: id}
      });
    },
    onChangeTab(key) {
      this.tabNow = key;
      if (this.tabNow == 1) {
        if (this.comments.length == 0){
          this.getData();

        }
      } else if (this.tabNow == 2) {
        if (this.criticisms.length == 0){
          this.getDataCriticism();
        }
      }
    }
  }
};
</script>

<style>
#answerCard
{
  min-width :1022px;
}

p {
    margin-top: 0;
    margin-bottom: 0.2em;
}
</style>
