<template>
  <div class="calendar">
    <p>
      <h3>calendar</h3>
      <ul class="list-group">
        <li v-for="item in items" class="list-element">{{ item }}</li>
      </ul>
    </p>
  </div>
</template>

<script>

import axios from 'axios';
export default {
  name: 'Calendar',
  data: () => ({
    items: [],
  }),
  mounted() {
    axios.get('http://localhost:3000/results/').then(response => (this.items = response.data.Events))
  }
};
</script>