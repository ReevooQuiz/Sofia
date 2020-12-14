<template>
  <a-row>
    <a-col :span="18" :offset="2">
      <v-md-editor v-model="questionValue" height="400px"></v-md-editor>
    </a-col>
    <a-col :span="2" :offset="1" >
      <a-select
          v-model:value="category"
          style="width: 80px"
          ref="select"
      >
        <a-select-option value="study">
          学习
        </a-select-option>
        <a-select-option value="life">
          生活
        </a-select-option>
        <a-select-option value="fashion">
          时尚
        </a-select-option>
      </a-select>
      <br/>
      <br/>
      <a-button type="primary" shape="round" @click="onPostQuestion"><strong>&ensp;发&ensp;布&ensp;</strong></a-button>
    </a-col>
  </a-row>
</template>

<script>
export default {
  data(){
    return {
      questionValue:"",
      category:"study",
    };
  },
  methods: {
    onPostQuestion(){
      console.log(this.questionValue);
    }
  },
}
</script>

<style scoped>

</style>
