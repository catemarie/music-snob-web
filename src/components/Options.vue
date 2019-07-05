<template>
  <div class="options">
    <p>
      city: {{ msg }}
    </p>
    <p>
      genre: {{ msg }}
    </p>
  </div>
</template>

<script>
export default {
  name: 'Options',
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
