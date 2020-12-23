<template>
  <v-snackbar
    v-model="showDialog"
    :bottom="setSwl.y === 'bottom'"
    :color="setSwl.color"
    :left="setSwl.x === 'left'"
    :multi-line="setSwl.mode === 'multi-line'"
    :right="setSwl.x === 'right'"
    :timeout="setSwl.timeout"
    :top="setSwl.y === 'top'"
    :vertical="setSwl.mode === 'vertical'"
  >
    {{ setSwl.text }}
    <template v-slot:action="{ attrs }">
      <v-btn dark text v-bind="attrs" @click="showDialog = false">Close</v-btn>
    </template>
  </v-snackbar>
</template>

<script>
export default {
  props: {
    dialog: { type: Boolean, default: false }
  },
  components: {},
  data() {
    return {
      showDialog: false,
      setSwl: {
        color: "",
        mode: "",
        snackbar: false,
        text: "Hello, I'm a snackbar",
        timeout: 6000,
        x: null,
        y: "top"
      }
    };
  },
  mountad() {},
  computed: {
    compSwal() {
      return this.$store.state.account.swal;
    }
  },
  watch: {
    compSwal(newValue, oldValue) {
      this.setSwl = newValue;
      this.showDialog = newValue.snackbar;
    },
    dialog(newValue, oldValue) {
      this.showDialog = this.dialog;
    }
  }
};
</script>
<style scoped>
</style>