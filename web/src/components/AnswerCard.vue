<template>
  <div id="answerCard">
    <a-card  size="small" style="border-radius : 3px">
        <!-- <a-card-meta title="ans.user">
            <template #avatar>
                <a-avatar src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png" />
            </template>
        </a-card-meta> -->
      <a-row>
        <a-col :span="3">
          <img src="https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png" style="width:120px" />
        </a-col>
        <a-col :span="21">
          <a-comment>
            <template #actions>
              <span key="comment-basic-like">
                <a-tooltip title="Like">
                  <template v-if="action === 'liked'">
                    <LikeFilled @click="like" />
                  </template>
                  <template v-else>
                    <LikeOutlined @click="like" />
                  </template>
                </a-tooltip>
                <span style="padding-left: '8px';cursor: 'auto'">{{ likes }}</span>
              </span>
              <span key="comment-basic-dislike">
                <a-tooltip title="Dislike">

                  <template v-if="action === 'disliked'">
                    <DislikeFilled @click="dislike" />
                  </template>
                  <template v-else>
                    <DislikeOutlined @click="dislike" />
                  </template>
                </a-tooltip>
                <!-- <span style="padding-left: '8px';cursor: 'auto'">{{ dislikes }}</span> -->
              </span>
              <span key="comment-basic-comment" @click="clickComment" >
                <MessageTwoTone v-if="showComment" />
                <MessageOutlined v-else />
                <span style="padding-left: '8px';cursor: 'auto'">{{ ans.comment_count }}</span>
              </span>
              <a-button v-if="writeComment" @click="onWriteComment" type="primary" shape="pill" size="small">取消评论</a-button>
              <a-button v-else @click="onWriteComment" type="primary" shape="pill" size="small">我要评论</a-button>
            </template>
            <template #author>
              <a> {{ans.owner.user_name}}</a>
            </template>
            <template #avatar>
              <a-avatar
                src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                alt="Han Solo"
              />
            </template>
            <template #extra>
                <a-tag>
                    #tag
                </a-tag>
            </template>
            <template #content>
              <p>
                {{ans.content}}
              </p>
            </template>
            </a-comment>
            <span v-if="writeComment">
              <a-textarea v-model:value="writeCommentValue" placeholder="您的评论" :rows="4" />
              <br/>
              <a-row type="flex" justify="end">
                <a-button @click="onCommitComment" type="primary" shape="pill" size="small">提交评论</a-button>
              </a-row>
            </span>
            <span v-if="showComment">
<!--              <a-divider/>-->
              <a-list
                  class="comment-list"
                  :header="'评论'"
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
                    <a-comment :author="item.user" :avatar="item.avatar">
<!--                      <template #actions>-->
<!--                        <span v-for="action in item.actions">{{ action }}</span>-->
<!--                      </template>-->
                      <template #content>
                        <p>
                          {{ item.content }}
                        </p>
                      </template>
<!--                      <template #datetime>-->
<!--                        <a-tooltip :title="item.datetime.format('YYYY-MM-DD HH:mm:ss')">-->
<!--                          <span>{{ item.datetime.fromNow() }}</span>-->
<!--                        </a-tooltip>-->
<!--                      </template>-->
                    </a-comment>
                  </a-list-item>
                </template>
              </a-list>

            </span>
        </a-col>

      </a-row>
    </a-card>
    <br/>
  </div>
</template>



<script >
import moment from "moment";
import { LikeFilled, LikeOutlined,DislikeFilled,DislikeOutlined,MessageOutlined,MessageTwoTone } from '@ant-design/icons-vue';
import { notification } from 'ant-design-vue';

const comments=[{
    id:1,
    user:"violedo",
    avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
    content:"We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully andefficiently.",

},{
  id:2,
  user:"violedo",
  avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
  content:"We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully andefficiently.",

},{
  id:3,
  user:"violedo",
  avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
  content:"We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully andefficiently.",

},
];

export default {
  components: {
     LikeFilled, LikeOutlined,DislikeFilled,DislikeOutlined,MessageOutlined,MessageTwoTone
  },
  props: ['ans'] ,

  data() {
    return {
      action: null,
      likes:this.ans.like_count,
      // dislikes:this.ans.dislike_count,
      showComment:false,
      comments,
      showLoadingMore: true,
      writeComment:false,
      writeCommentValue:"",
      moment
    };
  },
  methods: {
    like() {
      this.likes = 1;
      // this.dislikes =  0;
      this.action = "liked";
    },
    dislike() {
      this.likes = 0;
      // this.dislikes = 1;
      this.action = "disliked";
    },
    clickComment() {
        this.showComment=!this.showComment;
        if (this.showComment){
          console.log(this.showComment);

        }
    },
    getData(callback) {/****************************************TODO getData********************************/
      callback({results:null});
      // reqwest({
      //   url: fakeDataUrl,
      //   type: 'json',
      //   method: 'get',
      //   contentType: 'application/json',
      //   success: res => {
      //     callback(res);
      //   },
      // });
    },
    onLoadMore() {
      this.loadingMore = true;
      this.getData(res => {
        //this.comments = this.comments.concat(res.results);
        this.loadingMore = false;
        // if (xxxxx){
        //   this.showLoadingMore=false;
        // }
        // this.$nextTick(() => {
        //   window.dispatchEvent(new Event('resize'));
        // });
      });
    },
    onWriteComment(){
      if (!this.writeComment)
        this.writeComment=true;
      else {
        this.writeComment=false;
        this.writeCommentValue="";
      }
    },
    onCommitComment(){
      if (this.writeCommentValue.length==0){
        notification.open({
          message: '您的评论不能为空',
          // description:
          //     'This is the content of the notification. This is the content of the notification. This is the content of the notification.',
          // onClick: () => {
          //   console.log('Notification Clicked!');
          // },
        });
        return;
      }
      if (this.writeCommentValue.length>150){
        notification.open({
          message: '您的评论过长',
          description:
              '评论的最大长度为150字',
        });
        return;
      }
      /*********************************TODO commit comment*********************************************/

      this.writeComment=false;
      this.writeCommentValue="";
    }
  }
};
</script>

<style>
#answerCard
{
  min-width :1280px;
}
</style>
