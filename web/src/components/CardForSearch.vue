<template>
  <div class="cardForSearch">
    <a-card hoverable :title="info.title" size="small" style="border-radius :5px">
      <!-- <a-divider /> -->
      <!-- <a-descriptions title="属性" size="small"> -->
      <!-- <a-descriptions > -->
        <div v-for="(item) in info.attr" v-bind:key="item.name">
        <!-- <a-descriptions-item v-for="(item) in info.attr" v-bind:key="item.name"> -->
          <!-- <a-tag color="#eeeeee"> -->
            <!-- <div style="background-color:#eeeeee"> -->
            <a-row >
              <a-col :span="8">
                <a-tag  color="#f7c5c5">
                <font color="black">{{item.name}}</font>
                </a-tag>
              </a-col>
              <a-col :span="12" :offset="4">
                <a-tag color="#a3dbd8" >
                  <font color="black">{{item.value}}</font>
                </a-tag>
              </a-col>
            </a-row><br/>
            </div>
            
            <!-- </div> -->
          <!-- </a-tag> -->
        <!-- </a-descriptions-item> -->
      <!-- </a-descriptions> -->
      <!-- <a-divider /> -->
      <!--      <a-descriptions title="标签" size="small">-->
      <!--        <a-descriptions-item v-for="(item) in info.tags" v-bind:key="item.index">-->
      <!--          <a-tag color="#a3dbd8">-->
      <!--            {{item}}-->

      <!--          </a-tag>-->
      <!--        </a-descriptions-item>-->
      <!--      </a-descriptions>-->
    </a-card>
    <br />
  </div>
</template>



<script >
import moment from "moment";
export default {
  props: ["info"],
  data() {
    return {};
  },
  created() {
    console.log(this.info);
  },
  methods: {}
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
#attribute {
  color: #000;
}


</style>
