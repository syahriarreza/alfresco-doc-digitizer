<template>
  <v-menu
    v-model="showDialog"
    :close-on-content-click="false"
    :nudge-right="40"
    transition="scale-transition"
    offset-y
    min-width="290px"
  >
    <template v-slot:activator="{ on, attrs }">
      <v-text-field v-model="valDate" prepend-icon="mdi-calendar" readonly v-bind="attrs" v-on="on"></v-text-field>
    </template>
    <v-date-picker v-model="valDate" @input="showDialog = false"></v-date-picker>
  </v-menu>
</template>

<script>
export default {
  name: "BaseDatePicker",
  props: {
    menu: false,
  },
  data() {
    return {
      showDialog: false,
      valDate: new Date().toISOString().substr(0, 10),
    };
  },
  watch: {
    menu(newValue) {
      this.showDialog = this.dialog;
    },
    valDate(newValue) {
      this.$emit("input", newValue);
    },
  },
};
</script>

<style lang="sass">
  //
</style>
