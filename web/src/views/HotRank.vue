<template>
  <a-row type="flex" justify="space-around" align="middle">
    <a-col :span="8">
      <a-button type="primary" :disabled="!hasSelected" :loading="loading" @click="deleteHot">
        撤榜
      </a-button>
      <br/>
      <br/>
      <a-table
          :columns="columns"
          :data-source="data"
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
    <a-col :span="7">
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
    </a-col>
  </a-row>
</template>

<script>
const columns=[
  {
    dataIndex: 'index',
    key: 'index',
    slots: {  title: 'customTitle' },
    width:70
  },{
    dataIndex: 'name',
    key: 'name',
    ellipsis:true,
  },{
    dataIndex: 'heat',
    key: 'heat',
    width:70
  },
];
const data=[{
  index:1,
  name:"wg23rgeeb",
  heat:"2346"
},{
  index:2,
  name:"wg23rgeeb",
  heat:"2346"
},{
  index:3,
  name:"wg23rgeeb",
  heat:"2346"
},{
  index:4,
  name:"wg23rgeeb",
  heat:"2346"
},{
  index:5,
  name:"wg23rgeeafsssssssgeqwhrnjetkry,tu.lkjehwrgeb",
  heat:"2346"
},{
  index:6,
  name:"wg23rgeerkdjehsktdulyktjhrstykfjtretjsytstjhrsjheratjb",
  heat:"2346"
},{
  index:7,
  name:"wg23rgeeb",
  heat:"2346"
},{
  index:8,
  name:"wg23rgeeb",
  heat:"2346"
},{
  index:9,
  name:"wg23rgeeb",
  heat:"2346"
},{
  index:10,
  name:"wg23rgeeb",
  heat:"2346"
}
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
      data,
      apealData,
      columns,
      selectedRowKeys: [],
      loading:false,
    };
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

<style scoped>

</style>
