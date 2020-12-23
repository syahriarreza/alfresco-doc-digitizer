<template>
  <v-tooltip top>
    <template v-slot:activator="{ on, attrs }">
      <!-- <v-icon :color="color" small v-bind="attrs" v-on="on" @click="click">{{icon}}</v-icon> -->
      <v-btn
        small
        rounded
        :disabled="disabled"
        :dark="!disabled && dark"
        :color="color"
        v-bind="attrs"
        v-on="on"
        @click="click"
      >
        <v-icon small>
          {{ icon }}
        </v-icon>
      </v-btn>
    </template>
    <span>{{ title }}</span>
  </v-tooltip>
</template>

<script>
export default {
  name: 'BaseBtnIcon',

  props: {
    color: {
      type: String,
      default: 'primary',
    },
    icon: {
      type: String,
      default: 'mdi-cloud-download',
    },
    title: {
      type: String,
      default: 'Download',
    },
    disabled: {
      type: Boolean,
      default: false,
    },
    dark: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {}
  },
  methods: {
    click(calories) {
      this.$emit('click', true)
    },
  },
}
</script>
