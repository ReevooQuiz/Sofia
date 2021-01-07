<template>
  <div class="questionForSearch">
    <a-card hoverable :title="ques.title" size="small" style="border-radius : 3px">
      <template #extra>
        <a-tag color="#68b0af" style="margin-top:10px">
          <a-tooltip title="分区">
            <DatabaseOutlined />
            {{ques.category}}
          </a-tooltip>
        </a-tag>
      </template>
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
            <template #datetime>
              <a-tooltip :title="moment(ques.time).format('YYYY-MM-DD HH:mm:ss')">
                <span>{{ moment(ques.time).fromNow() }}</span>
              </a-tooltip>
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
