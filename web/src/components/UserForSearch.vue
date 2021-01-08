<template>
<div style="margin-bottom:10px">
  <a-card hoverable >
    <a-comment  @click="toUser">
      <template v-if="admin" #actions>
        <a-button v-if="user.banned"  type="primary" style="background-color:#fbbdbd;border-color: #ecc7d4;" shape="round" size="small" @click="onBanUser">解禁用户</a-button>
        <a-button v-else  type="primary"  shape="round" size="small" @click="onBanUser">封禁用户</a-button>
      </template>
      <template #author><a>{{user.name}}</a></template>
      <template #avatar>
        <a-avatar
        @click="gotoPerson(user.uid)"
            :src="user.icon"
            :alt="user.nickname"
        />
      </template>
      <template #content>
        <p @click="gotoPerson(user.uid)">
          {{user.profile}}
        </p>
      </template>
    </a-comment>
  </a-card>
  <!-- <br/> -->
</div>
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
      postRequest("/ban", {uid:""+this.user.uid,ban:this.user.banned},(e)=>{
        console.log(e);
      },{errorCallback:(e)=>{
          console.log(e);
        }});
    },
     gotoPerson(id) {
      console.log(id)
      this.$router.push({
        path: "/personalSetOthers",
        query: { uId: id }
      });
    }
  }
}
</script>

<style scoped>

</style>
