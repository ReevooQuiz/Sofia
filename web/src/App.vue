<template>
  <div id="background">
    <Navigator></Navigator>
    <div id="nav">
      <router-link to="/">Home123</router-link>|
      <router-link to="/search">search</router-link>|
      <router-link to="/question">Question</router-link>
    </div>

    <router-view />
  </div>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
import Navigator from "@/components/Navigator.vue";

@Options({
  components: {
    Navigator
  }
})
export default class App extends Vue {}
</script>


<style>
#background {
  background-color: #edeeed;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /* text-align: center; */
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
