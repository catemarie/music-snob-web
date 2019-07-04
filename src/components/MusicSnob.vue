<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <p>
      city: [options go here]
    </p>
    <p>
      genre: [options go here]
    </p>
    <p>
      <h3>calendar</h3>
      <ul>
        <li>[results]</li>
        <li>[go]</li>
        <li>[here]</li>
      </ul>
    </p>
  </div>
</template>

<script>
export default {
  name: 'MusicSnob',
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
