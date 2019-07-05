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