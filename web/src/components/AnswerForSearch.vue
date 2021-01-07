<template>
  <div class="questionForSearch">
    <a-card hoverable :title="title" size="small" style="border-radius : 3px">
      <a-row>
        <a-col :span="5">
          <img :src="ans.picture_urls" style="width:120px" @click="toQuestion"/>
        </a-col>
        <a-col :span="19">
          <a-comment @click="toQuestion">
            <template #author>
              <a> {{ans.answerer.nickname}}</a>
            </template>
            <template #avatar>
              <a-avatar
                  :src="ans.answerer.icon"
                  :alt="ans.answerer.nickname"
              />
            </template>
            <template #content>
              <v-md-editor mode="preview" v-model="ans.head" @click="toQuestion"></v-md-editor>  <a>...查看全部</a>
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
  props: ['ans'] ,

  data() {
    return {
      title:"来自问题："+this.ans.question_title
    };
  },
  methods: {
    toQuestion() {
      this.$router.push({ path:'/question' , query: { questionId: this.ans.qid } });
    }
  }
};
</script>
