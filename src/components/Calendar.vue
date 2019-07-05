<template>
  <div class="calendar">
    <p>
      <h3>calendar</h3>
      <ul class="list-group">
        <li class="list-group-item">{{ msg }}</li>
        <li class="list-group-item">{{ msg }}</li>
        <li class="list-group-item">{{ msg }}</li>
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