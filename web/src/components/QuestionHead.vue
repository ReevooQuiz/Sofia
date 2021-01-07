<template>
  <div id="questionHead">
    <a-card hoverable :title="ques.title" size="small" :headStyle="tstyle" style="border-radius : 3px">
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
              <a @click="gotoPerson(ques.raiser.uid)">{{ques.raiser.name}}</a>
            </template>
            <template #avatar>
              <a-avatar :src="ques.raiser.icon" :alt="ques.raiser.name" @click="gotoPerson(ques.raiser.uid)" />
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

              <span key="comment-basic-reply-to">
                <a-tooltip title="浏览量">
                  <FireOutlined />
                  {{ques.view_count}}
                </a-tooltip>
              </span>
              <span key="comment-basic-reply-to">
                <a-tooltip title="回答数">
                  <FileTextOutlined />
                  {{ques.answer_count}}
                </a-tooltip>
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
        <!-- <a-col :span="2" align="center">
          <h2>{{ques.category}}</h2>
          <h3>关注者</h3>
          <h3>{{ques.favorite_count}}</h3>
          <h3>热度</h3>
          <h3>{{ques.view_count}}</h3>
        </a-col>-->
      </a-row>
    </a-card>
    <br />
  </div>
</template>



<script >
import moment from "moment";
import { HeartFilled, HeartOutlined,FireOutlined ,FileTextOutlined} from "@ant-design/icons-vue";
import { postRequest } from "@/http/request";

export default {
  components: {
    HeartFilled,
    HeartOutlined,
    FireOutlined,
    FileTextOutlined
  },
  props: ["ques"],

  data() {
    return {
      time: moment(this.ques.time),
       tstyle: { "font-size": "21px", "font-weight": " bold", color: " #425050" }
    };
  },
  methods: {
    onFavorite() {
      console.log("a");
      this.ques.collected = !this.ques.collected;
      if (this.ques.collected) this.ques.favorite_count++;
      else this.ques.favorite_count--;
      postRequest(
        "/favorite",
        { qid: this.ques.qid, favorite: this.ques.collected },
        e => {
          console.log(e);
        },
        {
          errorCallback: e => {
            console.log(e);
          }
        }
      );
    },
    gotoPerson(id) {
      console.log(id)
      this.$router.push({
        path: "/personalSetOthers",
        query: { uId: id }
      });
    }
  }
};
</script>

<style >
#questionHead {
  min-width: 1022px;
}

.ant-comment-inner {
    display: flex;
    padding: 1px 0;
}
.v-md-editor-preview {
    padding: 2px;
    word-break: break-all;
}
.ant-divider-horizontal {
    display: block;
    clear: both;
    width: 100%;
    min-width: 100%;
    height: 1px;
    margin: 1px 0;
}
.ant-comment-actions {
    margin-top: 0px;
    padding-left: 0;
}
</style>
