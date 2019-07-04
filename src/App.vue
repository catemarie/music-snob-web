<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <MusicSnob msg="find only the music you like"/>
  </div>
</template>

<script>
import MusicSnob from './components/MusicSnob.vue'

export default {
  name: 'app',
  components: {
    MusicSnob
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
