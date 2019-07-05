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
export default {
  name: 'Calendar',
  data: () => ({
    items: ['result 1', 'result 2', 'result 3']
  })
}
</script>