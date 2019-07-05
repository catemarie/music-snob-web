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