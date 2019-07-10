<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <MusicSnob msg="only the music you like"/>
    <!-- <Options/> -->
    <Calendar/>
  </div>
</template>

<script>
import MusicSnob from './components/MusicSnob.vue'
import Options from './components/Options.vue'
import Calendar from './components/Calendar.vue'

export default {
  name: 'app',
  components: {
    MusicSnob,
    Options,
    Calendar
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
