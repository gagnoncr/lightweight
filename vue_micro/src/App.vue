<template>
  <div id="app">
    <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
    <Deploy msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
  import Deploy from './components/Deployments.vue'

  export default {
    name: 'app',
    components: {
      Deploy
    }
  }
</script>

<style>
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: black;
    margin-top: 20px;
    margin-bottom: 10px;
  }
</style>