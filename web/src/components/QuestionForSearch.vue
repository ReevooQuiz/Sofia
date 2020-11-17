<template>
  <div class="questionForSearch">
    <a-card hoverable :title="ques.title" size="small" style="border-radius : 3px">
      <a-row>
        <a-col :span="5">
          <img src="https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png" style="width:120px" />
        </a-col>
        <a-col :span="19">
          <!-- <template> -->
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
                <span style="padding-left: '8px';cursor: 'auto'">{{ dislikes }}</span>
              </span>
              <span key="comment-basic-reply-to">Reply to</span>
            </template>
            <template #author>
              <a> {{ques.user}}</a>
            </template>
            <template #avatar>
              <a-avatar
                src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                alt="Han Solo"
              />
            </template>
            <template #content>
              <p>
                {{ques.title}}
                We supply a series of design principles, practical patterns and high quality design
                resources (Sketch and Axure), to help people create their product prototypes beautifully and
                efficiently.
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
export default {
  props: ['ques'] ,
  
  data() {
    return {
      likes :  this.ques.likeNum,
      dislikes:   this.ques.dislikeNum,
      action: null,
      moment
    };
  },
  methods: {
    like() {
      this.likes = 1;
      this.dislikes =  0;
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
