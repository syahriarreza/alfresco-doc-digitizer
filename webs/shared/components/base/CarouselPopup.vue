<template>
  <v-dialog
    class="gallery--dialog"
    :value="isDialogShown"
    :overlay-opacity="0.9"
    :width="$vuetify.breakpoint.mdAndUp ? '50%' : '100%'"
  >
    <v-carousel :value="mediaSelIdx" hide-delimiters>
      <v-carousel-item v-for="(item,i) in medias" :key="i" :src="item.src" contain></v-carousel-item>
    </v-carousel>
  </v-dialog>
</template>

<script>
export default {
  props: {
    isShow: {
      type: Boolean,
      default: false
    },
    selIdx: Number,
    medias: {
      type: Array,
      default: () => [{ _id: 0, src: require("~/static/u856.jpg") }]
    }
  },
  // data() {
  //   return {
  //     isDialogShown: false,
  //     mediaSelIdx: 0,
  //   }
  // },
  // watch: {
  //   isDialogShown(newValue, oldValue) {
  //     this.$emit('update-isShow', newValue)
  //   },
  //   mediaSelIdx(newValue, oldValue) {
  //     this.$emit('update-selIdx', newValue)
  //   },
  // },
  computed: {
    isDialogShown: {
      get: function() {
        return this.isShow;
      },
      set: function(newValue) {
        console.log("isDialogShown newValue:", newValue);
        this.$emit("emitIsShow", newValue);
      }
    },
    mediaSelIdx: {
      get: function() {
        return this.selIdx;
      },
      set: function(newValue) {
        console.log("mediaSelIdx newValue:", newValue);
        this.$emit("emitSelIdx", newValue);
      }
    }
  }
};
</script>

<style scoped>
</style>