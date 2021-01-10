<template>
  <a-row  justify="center" >
    <a-col :span="12" :offset="6" >
      <a-button type="primary" :disabled="!hasSelected" :loading="loading" @click="deleteHot">
        撤榜
      </a-button>
      <br/>
      <br/>
      <a-table
          :columns="columns"
          :data-source="hotRankData"
          :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
          :pagination="false"
          size="middle"
      >
        <template #customTitle>
          <h2><strong>热榜</strong></h2>
        </template>
      </a-table>
      <br/>
      <br/>

    </a-col>
    <a-col :span="5" :offset="1">
        <img src="../assets/cactus2.png" height="345" width="180"    style="marginTop: 340px"/>
    </a-col>
    <!-- <a-col :span="7">
      <a-card title="申诉" :bordered="false">
        <a-list
            class="comment-list"
            item-layout="horizontal"
            :data-source="apealData"
        >
          <template #renderItem="{ item }">
            <a-list-item>
              <a-comment :author="item.author" :avatar="item.avatar">
                <template #content>
                  <p>
                    {{ item.content }}
                  </p>
                </template>
              </a-comment>
            </a-list-item>
          </template>
        </a-list>
      </a-card>
    </a-col> -->
  </a-row>
</template>

<script>
import { postRequest,getRequest } from "@/http/request.js";
const columns=[
  {
    dataIndex: 'index',
    key: 'index',
    slots: {  title: 'customTitle' },
    width:70
  },{
    dataIndex: 'title',
    key: 'title',
    ellipsis:true,
  },{
    dataIndex: 'view_count',
    key: 'view_count',
    width:70
  },
];
const apealData=[{
  author:"akangakang",
  avatar:"https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png",
  content:"为什么我的问题被封禁了为什么我的问题被封禁了为什么我的问题被封禁了为什么我的问题被封禁了为什么我的问题被封禁了为什么我的问题被封禁了",
},{
  author:"akangakang",
  avatar:"https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png",
  content:"为什么我的问题被封禁了为什么我的问题被封禁了为什么我的问题被封禁了为什么我的问题被封禁了为什么我的问题被封禁了为什么我的问题被封禁了",
},{
  author:"akangakang",
  avatar:"https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png",
  content:"为什么我的问题被封禁了为什么我的问题被封禁了为什么我的问题被封禁了为什么我的问题被封禁了为什么我的问题被封禁了为什么我的问题被封禁了",
},
];
export default {
  components :{

  },
  data(){
    return {
      apealData,
      columns,
      selectedRowKeys: [],
      loading:false,
      hotRankData:[]
    };
  },
  created() {
    this.$store.commit('changeTarget','hotRank');
    getRequest("/hotlist",
        (response)=>{
          this.hotRankData=response.result;
          let i=0;
          for (;i<10;){
            if (i>=this.hotRankData.length)
              break;
            this.hotRankData[i].index=++i;
          }
        }, {
          errorCallback:(e)=>{console.log(e)},
          params:{}
        });
  },
  computed: {
    hasSelected() {
      return this.selectedRowKeys.length > 0;
    },
  },
  methods:{
    onSelectChange(selectedRowKeys) {
      this.selectedRowKeys = selectedRowKeys;
    },
    deleteHot(){
      this.loading = true;
      //************************************TODO******************************
      this.loading=false;
      this.selectedRowKeys = [];
    }
  },
}
</script>

<style>
body {
  height: 100%;
  background-color: #edeeed;
}
</style>
