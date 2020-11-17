<template>
  <div class="answerCard">
    <a-card  size="small" style="border-radius : 3px">
        <!-- <a-card-meta title="ans.user">
            <template #avatar>
                <a-avatar src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png" />
            </template>
        </a-card-meta> -->
      <a-row>
        <a-col :span="4">
          <img src="https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png" style="width:120px" />
        </a-col>
        <a-col :span="20">
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
              <span key="comment-basic-comment">
                <MessageOutlined @click="comment" />
                <span style="padding-left: '8px';cursor: 'auto'">{{ ans.comment_count }}</span>
              </span>
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
            <!-- <template #datetime>
                <a-tooltip :title="moment().format('YYYY-MM-DD HH:mm:ss')">
                  <span>{{ moment().fromNow() }}</span>
                </a-tooltip>
            </template>-->
            </a-comment>
            <!-- </template> -->
        </a-col>
        
      </a-row>
    </a-card>
    <br/>
  </div>
</template>
       


<script >
import moment from "moment";
import { LikeFilled, LikeOutlined,DislikeFilled,DislikeOutlined,MessageOutlined } from '@ant-design/icons-vue';
export default {
  components: {
     LikeFilled, LikeOutlined,DislikeFilled,DislikeOutlined,MessageOutlined
  },
  props: ['ans'] ,
  
  data() {
    return {
      action: null,
      likes:this.ans.like_count,
      // dislikes:this.ans.dislike_count,
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
    comment() {
        console.log("comment");
    }
  }
};
</script>
