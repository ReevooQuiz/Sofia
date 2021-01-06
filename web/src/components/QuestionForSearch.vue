<template>
  <div class="questionForSearch">
    <a-card hoverable :title="ques.title" size="small" style="border-radius : 3px">
      <a-row>
        <a-col :span="5">
          <img :src="ques.picture_urls" style="width:120px" @click="toQuestion"/>
        </a-col>
        <a-col :span="19">
          <a-comment>
            <template #author>
              <a> {{ques.raiser.nickname}}</a>
            </template>
            <template #avatar>
              <a-avatar
                :src="ques.raiser.icon"
                :alt="ques.raiser.nickname"
              />
            </template>
            <template #content>
              <v-md-editor mode="preview" v-model="ques.head" @click="toQuestion"></v-md-editor>
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
    toQuestion() {
      this.$router.push({ path:'/question' , query: { questionId: this.ques.qid } });
    }
  }
};
</script>
