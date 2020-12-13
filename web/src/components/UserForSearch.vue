<template>
  <a-card hoverable>
    <a-comment  @click="toUser">
      <template #author><a>{{name}}</a></template>
      <template #avatar>
        <a-avatar
            :src="icon"
            :alt="name"
        />
      </template>
      <template #content>
        <p>
          {{profile}}
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
      name:this.user.name,
      icon:this.user.icon,
      profile:this.user.profile
    };
  },
  methods:{
    toUser(){
      console.log("toUser");
    }
  }
}
</script>

<style scoped>

</style>
