<template>
  <div class="questionForSearch">
    <a-card hoverable :title="ques.title" size="small" style="border-radius : 3px">
      <a-row>
        <a-col :span="5">
          <img src="https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png" style="width:120px" @click="toQuestion"/>
        </a-col>
        <a-col :span="19">
          <a-comment>
            <template #author>
              <a> {{ques.owner.user_name}}</a>
            </template>
            <template #avatar>
              <a-avatar
                src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                alt="Han Solo"
              />
            </template>
            <template #content>
              <p  @click="toQuestion">
                {{ques.description}}
              </p>
            </template>
          </a-comment>
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
      // likes :  this.ques.likeNum,
      // dislikes:   this.ques.dislikeNum,
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
    },
    toQuestion() {
      console.log("??")
      console.log( this.ques.qid )
      this.$router.push({ path:'/question' , query: { questionId: this.ques.qid } });
    }
  }
};
</script>
