<template>
  <section :id="id">
    <v-row no-gutters>
      <v-col cols="12">
        <slot></slot>
      </v-col>
    </v-row>
  </section>
</template>

<script>
  export default {
    name: 'BaseView',

    props: {
      id: {
        type: String,
        default: 'view',
      },
    },

    data: () => ({ sections: [] }),
  }
</script>
