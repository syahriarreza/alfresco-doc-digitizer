<template>
  <div :class="classes" class="pt-2 gt-col">
    <div class="gt-inner">
      <div class="gt-logo">
        <base-img
          aspect-ratio="1.7"
          :src="require(`@/static/sponsor/${item.img}.jpg`)"
          color="grey lighten-1"
          width="300"
          height="200"
          :class="{ full: fullWidthImage }"
          tile
        />
      </div>
      <div class="primary"></div>
    </div>
  </div>
</template>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script>
// Mixins
import Heading from "@/mixins/heading";

export default {
  name: "BaseSponsor",

  components: {
    BaseImg: () => import("./Img")
  },

  mixins: [Heading],

  props: {
    space: {
      type: [Number, String],
      default: 2
    },
    item: Object
  },
  data: () => ({
    fullWidthImage: true
  }),
  computed: {
    classes() {
      const classes = [`mb-${this.space}`];
      classes.push("text-center");
      return classes;
    }
  }
};
</script>

<style scoped>
.gt-col {
  width: 100%;
  position: relative;
  display: inline-block;
}
.gt-col > .gt-inner {
  display: inline-block;
  height: 100%;
  width: 100%;
}
.primary {
  width: 100%;
  height: 5px;
  border-radius: 0px 0px 5px 5px;
}
.gt-logo {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-wrap: wrap;
  text-align: center;
  padding: 10px;
  box-shadow: 0px -4px 14px 0px rgba(215, 215, 215, 0.349019607843137);
  -webkit-box-shadow: 0px -4px 14px 0px rgba(215, 215, 215, 0.349019607843137);
  -moz-box-shadow: 0px -4px 14px 0px rgba(215, 215, 215, 0.349019607843137);
  height: 200px;
  width: 100%;
  background-color: #fff;
  border-radius: 5px 5px 0px 0px;
  overflow: hidden;
}
.img-spongsor {
  width: 200px;
  height: 400px;
  object-fit: cover !important;
}

.full {
  width: 100%;
  height: auto;
}
img {
  width: 150px;
  border-radius: 2px;
  box-shadow: 1px 1px 3px 1px rgba(0, 0, 0, 0.5);
  transition: width 1s;
}
img:hover {
  cursor: pointer;
}
</style>
<style>
.gt-logo .v-responsive__sizer {
  padding-bottom: 0% !important;
}

.gt-logo .v-image__image--cover {
  background-size: auto 90px !important;
}
</style>