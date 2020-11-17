<template>
  <div class="questionHead">
    <a-card hoverable :title="ques.title" size="small" style="border-radius : 3px">
      <a-row>
        <a-col :span="4">
          <img src="https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png" style="width:120px" />
        </a-col>
        <a-col :span="18">
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
            <template #extra>
                <a-tag>
                    #tag
                </a-tag>
            </template>
            <template #content>
              <p>
                {{ques.content}}
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
        <a-col :span="2" align="center">
            <h2>关注者</h2>
            <h3>{{ques.followers}}</h3>
            <h2>热度</h2>
            <h3>{{ques.likeNum}}</h3>
        </a-col>
      </a-row>
    </a-card>
    <br/>
  </div>
</template>
       


<script >
import moment from "moment";
import { LikeFilled, LikeOutlined,DislikeFilled,DislikeOutlined } from '@ant-design/icons-vue';
export default {
  components: {
     LikeFilled, LikeOutlined,DislikeFilled,DislikeOutlined
  },
  props: ['ques'] ,
  
  data() {
    return {
      likes :  this.ques.likeNum,
      dislikes:   this.ques.dislikeNum,
      action: null,
      moment,
      followers: this.ques.followers,
      content: this.ques.content
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
