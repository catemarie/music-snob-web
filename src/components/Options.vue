<template>
  <div class="options">
    <div class="container">
      <div class="row">

        <!-- city options -->
        <div class="col-sm">
          <div class="dropdown" align="right">
            <button class="btn btn-default dropdown-toggle" type="button" id="dropdown_city" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
              city
              <span class="caret"></span>
            </button>
            <ul class="dropdown-menu" aria-labelledby="dropdown_city" id="city_menu">
              <li><a href="#">San Diego</a></li>
            </ul>
          </div>
        </div>

        <!-- genre options -->
        <div class="col-sm">
          <div class="dropdown" align="left">
            <button class="btn btn-default dropdown-toggle" type="button" id="dropdown_genre" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
              genre
              <span class="caret"></span>
            </button>
            <ul class="dropdown-menu" aria-labelledby="dropdown_genre" id="genre_menu">
              <li><a href="#">trance</a></li>
            </ul>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Options',
  data() {
    return {
      city_selected: "",
      genre_selected: ""
    }
  },

  mounted() {
    $("#city_menu a").click(function(e){
      e.preventDefault();
      this.city_selected = $(this).text();
      $("#dropdown_city").text(this.city_selected);
      console.log(this.city_selected);
    });

    $("#genre_menu a").click(function(e){
      e.preventDefault();
      this.genre_selected = $(this).text();
      $("#dropdown_genre").text(this.genre_selected);
      console.log(this.genre_selected);
    });
  },
}
</script>