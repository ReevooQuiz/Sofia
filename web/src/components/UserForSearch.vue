<template>
  <a-card hoverable>
    <a-comment  @click="toUser">
      <template v-if="admin" #actions>
        <a-button v-if="user.banned" type="danger" shape="round" size="small" @click="onBanUser">解禁用户</a-button>
        <a-button v-else type="danger" shape="round" size="small" @click="onBanUser">封禁用户</a-button>
      </template>
      <template #author><a>{{user.name}}</a></template>
      <template #avatar>
        <a-avatar
            :src="user.icon"
            :alt="user.nickname"
        />
      </template>
      <template #content>
        <p>
          {{user.profile}}
        </p>
      </template>
    </a-comment>
  </a-card>
  <br/>
</template>

<script>
import { postRequest,getRequest } from "@/http/request.js";
export default {
  props: ['user','admin'] ,
  data(){
    return {
    };
  },
  methods:{
    toUser(){
      console.log("toUser");//******************TODO*****************
    },
    onBanUser(){
      this.user.banned=!this.user.banned;
      postRequest("/ban", {uid:this.user.uid,ban:this.user.banned},(e)=>{
        console.log(e);
      },{errorCallback:(e)=>{
          console.log(e);
        }});
    }
  }
}
</script>

<style scoped>

</style>
