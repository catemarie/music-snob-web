<template>
  <div class="calendar">
    <p>
      <h3>calendar</h3>
      <ul>
        <li>{{ msg }}</li>
        <li>{{ msg }}</li>
        <li>{{ msg }}</li>
      </ul>
    </p>
  </div>
</template>

<script>
export default {
  name: 'Calendar',
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>