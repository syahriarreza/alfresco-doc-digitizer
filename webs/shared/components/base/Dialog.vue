<template>
  <v-row justify="center">
    <v-dialog v-model="showDialog" v-bind="$attrs" v-on="$listeners">
      <v-card>
        <v-card-title class="headline">{{title}}</v-card-title>
        <v-card-text>
          <v-container>
            <slot name="body" />
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <slot name="action" />
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
export default {
  props: {
    dialog: { type: Boolean, default: false },
    title: { type: String, default: "" }
  },
  components: {},
  data() {
    return {
      showDialog: false
    };
  },
  watch: {
    dialog(newValue, oldValue) {
      this.showDialog = this.dialog;
    },
    showDialog(newValue, oldValue) {
      this.$emit("change", newValue);
    }
  },
  methods: {},
  created() {},
  mounted() {}
};
</script>
<style scoped>
</style>