<template>
  <div class="cardForSearch">
    <a-card hoverable :title="info.title" size="small" style="border-radius :5px">
      <a-descriptions title="关键字" size="small">
        <a-descriptions-item v-for="(item) in info.keyWords" v-bind:key="item.index">{{item}}</a-descriptions-item>
      </a-descriptions>
      <a-divider />
      <a-descriptions title="属性" size="small">
        <a-descriptions-item v-for="(item) in info.attributes" v-bind:key="item.index">
          <a-tag id="tagBack" color="#eeeeee">
            {{item[0]}}
            <!-- <span style="text-align:right">
{{item[1]}}
            </span> -->
            <a-tag color="#a3dbd8" style="align:right,margin-left ：100px">{{item[1]}}</a-tag>
          </a-tag>
        </a-descriptions-item>
      </a-descriptions>
      <a-divider />
      <a-descriptions title="标签" size="small">
        <a-descriptions-item v-for="(item) in info.tags" v-bind:key="item.index">
          <a-tag color="#a3dbd8">
            {{item}}

          </a-tag>
        </a-descriptions-item>
      </a-descriptions>
    </a-card>
    <br />
  </div>
</template>
       


<script >
import moment from "moment";
export default {
  props: ["info"],
  data() {
    return {
      likes: 0,
      dislikes: 0,
      action: null,

      moment
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
.ant-card-head-title {
  display: inline-block;
  flex: 1;
  padding: 16px 0;
  overflow: hidden;
  white-space: nowrap;
  font-weight: bold;
  text-overflow: ellipsis;
  font-size: 16px;
}

.ant-descriptions-title {
  flex: auto;
  overflow: hidden;
  color: rgba(0, 0, 0, 0.85);
  font-weight: bold;
  font-size: 14px;
  line-height: 1.2;
  white-space: nowrap;
  text-overflow: ellipsis;
}

#tagBack{
   width: 100px;
}
/* .ant-tag-has-color, .ant-tag-has-color a, .ant-tag-has-color a:hover, .ant-tag-has-color .anticon-close, .ant-tag-has-color .anticon-close:hover {
    width: 70px;
} */
</style>