<template>
  <div>
    <a-row justify="center">
      <a-col>
        <img
          slot="cover"
          alt="example"
          :src="this.info.icon"
          style="height: 150px; border-radius: 50%"
        />
        <!-- <img src="../assets/cactus2.png" height="345" width="180"    style="marginTop: 340px"/> -->
      </a-col>
    </a-row>

    <br />
     <!-- <a-row justify="center">
       <a-tag v-if="info.follow===true"  style="margin-bottom:10px" >取消关注</a-tag>
        <a-tag v-else style="margin-bottom:10px" ghost>关注</a-tag>
    </a-row> -->
    <a-row justify="center">
       <a-tag color="#88d5d1" style="margin:3px">
            <UserOutlined />
            <span v-if="info.account_type===1">管理员</span>
            <span v-else>普通用户</span>
          </a-tag>
           <a-tag  style="margin:3px" color="#e7bed3"> <StarOutlined />
              <span v-if="info.follow===true" @click="unfollow">取消关注</span>
             <span v-else  @click="follow">关注</span>
             
             </a-tag>
        
    </a-row>
    <a-row justify="center">
      <div  v-for="(item) in this.info.labels" v-bind:key="item.index"> <a-tag color="#acd3d0" style="margin:3px">
           {{item}}
          </a-tag>
      </div>
      
    </a-row>
    <!-- <br /> -->
    <a-row>
     <a-col :span="16" :offset="4">
       <a-divider />
    </a-col>
    </a-row>
    
    <a-row  justify="center">
      <a-col>
        <span>
            <TeamOutlined />
            {{info.following_count}} 关注 ·
          </span>
          <span>{{info.follower_count}} 粉丝 ·</span>
          <span>
            <LikeOutlined />
            {{info.like_count}} 赞
          </span>

          <span>
            <FormOutlined />
            {{info.question_count}} 问题 ·
          </span>
          <span>
            <CopyOutlined />
            {{info.answer_count}} 回答
          </span>
          <a-divider />
      </a-col>
    </a-row>

<br/>
    <a-row >
       <a-col :span="10" :offset="10">
         <!-- <div > -->
            <a-row>
             
              <a-col :span="3">
                <span class="set-lable">姓名</span>
              </a-col>
              <a-col :span="16" :offset="2">
                <span class="set-content">{{info.name}}</span>
              </a-col>
            </a-row>

            <br />
            <a-row>
              <a-col :span="3">
                <span class="set-lable">昵称</span>
              </a-col>
              <a-col :span="17" :offset="2">
                <span class="set-content">{{info.nickname}}</span>
              </a-col>
            </a-row>
            <br />
            <a-row>
              <a-col :span="3">
                <span class="set-lable">性别</span>
              </a-col>
              <a-col :span="17" :offset="2">
                <span v-if="info.gender==0" class="set-content">男</span>
                <span v-else-if="info.gender==1" class="set-content">女</span>
              </a-col>
            </a-row>
            <br />
            <a-row>
              <a-col :span="3">
                <span class="set-lable">邮箱</span>
              </a-col>
              <a-col :span="17" :offset="2">
                <span class="set-content">{{info.email}}</span>
              </a-col>
            </a-row>
            <br />
            <a-row>
              <a-col :span="3">
                <span class="set-lable">个人简介</span>
              </a-col>
              <a-col :span="17" :offset="2">
                <span class="set-content">{{info.profile}}</span>
              </a-col>
           </a-row>
       
       </a-col>
    </a-row>
     
    <br />

    <br />
  </div>
</template>



<script>
import { defineComponent } from "vue";
import { Options, Vue } from "vue-class-component";
import SubMenu from "../../components/PersonalNavigation";
import { message } from "ant-design-vue";
import { PlusOutlined, LoadingOutlined } from "@ant-design/icons-vue";
import {
  CameraOutlined,
  FireOutlined,
  VerifiedOutlined,
  UserOutlined,
  TeamOutlined,
  LikeOutlined,
  FormOutlined,
  CopyOutlined,
  StarOutlined
} from "@ant-design/icons-vue";
import { getRequest ,putRequest} from "@/http/request.js";

export default {
  components: {
    SubMenu,
    CameraOutlined,
    FireOutlined,
    VerifiedOutlined,
    UserOutlined,
    TeamOutlined,
    LikeOutlined,
    FormOutlined,
    CopyOutlined,
    PlusOutlined,
    LoadingOutlined,
    StarOutlined
  },
  data() {
    return {
      uid: 0,
      info: {}
    };
  },
  created() {
    console.log(this.$route.query.uId);
    this.uid = this.$route.query.uId;
    getRequest("/publicInfo", this.handleCallback, {
      errorCallback: e => {
        console.log(JSON.stringify(e));
      },
      params: { id: this.uid }
    });
  },

  methods: {
    handleCallback(response) {
      console.log(response);
      this.info = response.result;
    },
     unfollow()
    {
      putRequest("/follow",{uid:this.uid,follow:false},(res)=>{
        if(res.code===0)
        {
          //取关
          this.info.follow=false;
        }
         else{
          message.error("操作失败")
        }
      },{errorCallback:(e)=>{JSON.stringify(e)}})
    },
     follow(id)
    {
      putRequest("/follow",{uid:this.uid,follow:true},(res)=>{
        if(res.code===0)
        {
          //取关
          
         this.info.follow=true;
        
        }
        else{
          message.error("操作失败")
        }
      },{errorCallback:(e)=>{JSON.stringify(e)}})
    }
  }
};
</script>

<style>
.set-lable {
  font-size: 18px;
  color: #485355f5;

  font-weight: bold;
}

.set-content {
  font-size: 16px;
  color: #33393af5;

  /* font-weight: bold; */
}
.avatar-uploader > .ant-upload {
  width: 128px;
  height: 128px;
}
body {
  height: 100%;
  background-color: #edeeed;
}

.personal-nav .ant-menu-inline {
  width: 40%;
}

.ant-divider-horizontal {
  display: block;
  clear: both;
  width: 100%;
  min-width: 100%;
  height: 1px;
  margin: 16px 0;
  background-color: #acbab9;
}

.ant-btn-background-ghost {
  color: #98a7a6;
  background: transparent !important;
  border-color: #98a7a6;
}
</style>