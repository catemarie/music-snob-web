<template>
  <div class="options">
    <div class="dropdown">
      <button class="btn btn-default dropdown-toggle" type="button" id="dropdownMenu1" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
        city
        <span class="caret"></span>
      </button>
      <ul class="dropdown-menu" aria-labelledby="dropdownMenu1">
        <li role="separator" class="divider"></li>
        <li><a href="#">{{ msg }}</a></li>
      </ul>
    </div>

    <div class="dropdown">
      <button class="btn btn-default dropdown-toggle" type="button" id="dropdownMenu1" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
        genre
        <span class="caret"></span>
      </button>
      <ul class="dropdown-menu" aria-labelledby="dropdownMenu1">
        <li role="separator" class="divider"></li>
        <li><a href="#">{{ msg }}</a></li>
      </ul>
    </div>
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