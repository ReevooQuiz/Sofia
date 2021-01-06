<template>
  <a-card hoverable>
    <a-comment  @click="toUser">
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
export default {
  props: ['user'] ,
  data(){
    return {
    };
  },
  methods:{
    toUser(){
      console.log("toUser");//******************TODO*****************
    }
  }
}
</script>

<style scoped>

</style>
