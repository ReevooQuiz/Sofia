<template>
  <a-row id="search">
    <a-col :span="8">
        <img src="../assets/cactus1.png" height="500" width="360"    style="marginTop: 95px"/>
    </a-col>
    <a-col :span="8">
      <a-row justify="center" style="marginTop: 80px">
          <h1>这里一片荒芜</h1>
      </a-row>
      <a-row justify="center">
          <h1>直到有了绿洲……</h1>
      </a-row>
      <a-row justify="center">
          <h1>Sofia</h1>
      </a-row>
      <a-row justify="center">
        <div class="col-7">
             <a-input-search placeholder="输入您的问题" style="width: 450px" @search="onSearch" size="large" v-model:value="value"/>
        </div>
      </a-row>    
    </a-col>
    <a-col :span="5" :offset="3">
        <img src="../assets/cactus2.png" height="345" width="180"    style="marginTop: 250px"/>
    </a-col>
  </a-row>
</template>

<script>
export default {
  data() {
    return {
      value: '',
    };
  },
  methods: {
    onSearch(value) {
      if (value!=''){
        this.$router.push({ path:'/search' , query: { content: value } });
      }
    },
  },
};
</script>

<style>
#search {
  background-color: #edeeed;
}
</style>