<template>
  <div id="questionHead">
    <a-card hoverable :title="ques.title" size="small" style="border-radius : 3px">
      <template #extra>
        <a-tag v-for="(item) in ques.tags" v-bind:key="item.tid">
          {{111}}
        </a-tag>
      </template>
      <a-row>
        <a-col :span="3">
          <img src="https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png" style="width:120px" />
        </a-col>
        <a-col :span="19">
          <!-- <template> -->
          <a-comment>
            <template #author>
              <a>{{ques.user}}</a>
            </template>
            <template #avatar>
              <a-avatar
                src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                alt="Han Solo"
              />
            </template>
            <template #content>
              <p>{{ques.content}}</p>
            </template>
          </a-comment>
        </a-col>
        <a-col :span="2" align="center">
          <h2>关注者</h2>
          <h3>{{ques.followers}}</h3>
          <h2>热度</h2>
          <h3>{{ques.viewCount}}</h3>
        </a-col>
      </a-row>
    </a-card>
    <br />
  </div>
</template>



<script >
import moment from "moment";
import {
  LikeFilled,
  LikeOutlined,
  DislikeFilled,
  DislikeOutlined
} from "@ant-design/icons-vue";
export default {
  components: {
    LikeFilled,
    LikeOutlined,
    DislikeFilled,
    DislikeOutlined
  },
  props: ["ques"],

  data() {
    return {
      // likes: this.ques.likeNum,
      // dislikes: this.ques.dislikeNum,
      action: null,
      moment,
      followers: this.ques.followers,
      content: this.ques.content
    };
  },
  methods: {
    like() {
      this.likes = 1;
      this.dislikes = 0;
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

<style >
#questionHead{
  min-width: 1280px;
}
</style>