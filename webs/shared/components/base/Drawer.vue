<template>
  <v-navigation-drawer
    bottom
    color="transparent"
    fixed
    height="auto"
    overlay-color="secondary"
    overlay-opacity=".8"
    temporary
    v-bind="$attrs"
    v-on="$listeners"
  >
    <!-- <v-list color="white" shaped>
      <v-list-item v-for="(item, i) in items" :key="i" :to="item.to" color="primary">
        <v-list-item-content>
          <v-list-item-title class="text-uppercase" v-text="item.title" />
        </v-list-item-content>
      </v-list-item>
    </v-list>-->
    <v-list color="white" shaped>
      <template v-for="(item, i) in items">
        <v-list-group v-if="item.submenus.length > 0" :key="i" color="primary">
          <template v-slot:activator>
            <v-list-item-title class="text-uppercase" v-text="item.title"></v-list-item-title>
          </template>
          <v-list-item v-for="(submenu, i) in item.submenus" :key="i" color="primary">
            <v-list-item-title class="text-uppercase sub-menu" v-text="submenu.title" :to="item.to"></v-list-item-title>
          </v-list-item>
        </v-list-group>
        <v-list-item v-else :key="i" color="primary">
          <v-list-item-title class="text-uppercase" :to="item.to" v-text="item.title"></v-list-item-title>
        </v-list-item>
      </template>
    </v-list>
    <slot name="append-outer" />
  </v-navigation-drawer>
</template>

<script>
export default {
  name: "Drawer",

  props: {
    items: {
      type: Array,
      default: () => []
    }
  },
  data: () => ({})
};
</script>
<style scoped>
.sub-menu {
  padding: 0px 40px;
  font-size: 13px;
}
</style>