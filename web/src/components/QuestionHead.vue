<template>
  <div id="questionHead">
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
        <a-col :span="22">
          <!-- <template> -->
          <a-comment>
            <template #author>
              <a>{{ques.raiser.name}}</a>
            </template>
            <template #avatar>
              <a-avatar
                :src="ques.raiser.icon"
                :alt="ques.raiser.name"
              />
            </template>
            <template #content>
              <v-md-editor mode="preview" v-model="ques.content"></v-md-editor>
            </template>
            <template #actions>
              <span key="comment-basic-approve">
                <a-tooltip title="favorite">
                  <template v-if="ques.collected">
                    <HeartFilled @click="onFavorite" />
                  </template>
                  <template v-else>
                    <HeartOutlined @click="onFavorite" />
                  </template>
                </a-tooltip>
                <span style="padding-left: '8px';cursor: 'auto'">{{ ques.favorite_count }}</span>
              </span>
            </template>
            <template #datetime>
              <a-tooltip :title="time.format('YYYY-MM-DD HH:mm:ss')">
                <span>{{ time.fromNow() }}</span>
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
        <a-col :span="2" align="center">
          <h2>{{ques.category}}</h2>
          <h3>关注者</h3>
          <h3>{{ques.favorite_count}}</h3>
          <h3>热度</h3>
          <h3>{{ques.view_count}}</h3>
        </a-col>
      </a-row>
    </a-card>
    <br />
  </div>
</template>



<script >
import moment from "moment";
import {
  HeartFilled,
  HeartOutlined,
} from "@ant-design/icons-vue";
import {postRequest} from "@/http/request";

export default {
  components: {
    HeartFilled,
    HeartOutlined,
  },
  props: ["ques"],

  data() {
    return {
      time:moment(this.ques.time),
    };
  },
  methods: {
    onFavorite(){
      console.log("a");
      this.ques.collected=!this.ques.collected;
      if (this.ques.collected)
        this.ques.favorite_count++;
      else this.ques.favorite_count--;
      postRequest("/favorite", {qid:this.ques.qid,favorite:this.ques.collected},(e)=>{
        console.log(e);
      },{errorCallback:(e)=>{
          console.log(e);
        }});
    },
  }
};
</script>

<style >
#questionHead{
  min-width: 1280px;
}
</style>
