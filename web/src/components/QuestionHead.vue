<template>
  <div id="questionHead">
    <a-card hoverable :title="ques.title" size="small" style="border-radius : 3px">
      <template #extra>
        <!-- <a-tag v-for="(item) in ques.tags" v-bind:key="item.tid">
          {{111}}
        </a-tag> -->
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
              <p>{{ques.content}}</p>
            </template>
            <template #actions>
              <span key="comment-basic-approve">
                <a-tooltip title="favorite">
                  <template v-if="ans.favorited">
                    <HeartFilled @click="onFavorite" />
                  </template>
                  <template v-else>
                    <HeartOutlined @click="onFavorite" />
                  </template>
                </a-tooltip>
                <span style="padding-left: '8px';cursor: 'auto'">{{ ans.favorite_count }}</span>
              </span>
            </template>
            <template #datetime>
              <a-tooltip :title="item.datetime.format('YYYY-MM-DD HH:mm:ss')">
                <span>{{ time.fromNow() }}</span>
              </a-tooltip>
            </template>
          </a-comment>
        </a-col>
        <a-col :span="2" align="center">
          <h2>{{ques.category}}</h2>
          <br/>
          <h2>关注者</h2>
          <h3>{{ques.favorite_count}}</h3>
          <h2>热度</h2>
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
      /***********WAITING API**************/
      // this.ques.favorited=!this.ques.favorited;
      // if (this.ques.favorited)
      //   this.ques.favorite_count++;
      // else this.ques.favorite_count--;
      // postRequest("/favorite", {qid:this.ques.qid,favorite:this.ques.favorited},(e)=>{
      //   console.log(e);
      // },{errorCallback:(e)=>{
      //     console.log(e);
      //   }});
    },
  }
};
</script>

<style >
#questionHead{
  min-width: 1280px;
}
</style>
