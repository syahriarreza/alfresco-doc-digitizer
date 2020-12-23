<template>
  <!-- <v-navigation-drawer absolute color="primary" dark :mini-variant="drawer.mini" app> -->
  <v-navigation-drawer
    v-model="showDrawer"
    color="primary darken-2"
    dark
    :mini-variant="drawer.mini"
    app
  >
    <v-list nav dense class="py-0 px-0">
      <!-- logo -->
      <v-list-item class="logo px-2" :class="!drawer.mini && 'py-2'">
        <v-list-item-avatar v-if="drawer.mini">
          <v-img :src="logoSM" contain></v-img>
        </v-list-item-avatar>
        <v-img
          v-if="!drawer.mini"
          :src="logo"
          contain
          max-height="42px"
        ></v-img>
      </v-list-item>
      <v-divider class="mx-3" />

      <!-- user info & avatar -->
      <v-tooltip right :disabled="!drawer.mini">
        <template v-slot:activator="{ on, attrs }">
          <v-list-item two-line v-bind="attrs" v-on="on">
            <v-list-item-avatar>
              <v-icon large>mdi-account-circle-outline</v-icon>
            </v-list-item-avatar>

            <v-list-item-content>
              <v-list-item-title>{{ 'Reza Dummy' }}</v-list-item-title>
              <v-list-item-subtitle>{{
                'dummy@gmail.com'
              }}</v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
        </template>
        <span>
          <strong class="d-block">{{ 'Reza Dummy' }}</strong>
          {{ 'dummy@gmail.com' }}
        </span>
      </v-tooltip>
      <v-divider class="mx-3"></v-divider>

      <!-- menu items -->
      <template v-for="(item, i) in menu.items">
        <template
          v-if="
            !!item.children &&
            item.children.length > 0
          "
        >
          <v-subheader
            :key="'subheader' + i"
            v-text="item.title"
            v-show="!drawer.mini"
          ></v-subheader>
          <v-divider :key="'divider' + i" v-show="drawer.mini"></v-divider>
          <template v-for="(item, j) in item.children">
            <v-tooltip
              :key="j + '-' + i"
              right
              :disabled="!drawer.mini"
            >
              <template v-slot:activator="{ on, attrs }">
                <v-list-item
                  class="menu px-3"
                  :to="item.to"
                  router
                  exact
                  v-bind="attrs"
                  v-on="on"
                >
                  <v-list-item-action>
                    <v-icon>{{ item.icon }}</v-icon>
                  </v-list-item-action>
                  <v-list-item-content>
                    <v-list-item-title class="fw-09">{{
                      item.title
                    }}</v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
              </template>
              <span v-text="item.title"></span>
            </v-tooltip>
          </template>
        </template>
        <template v-else>
          <v-tooltip :key="i" right :disabled="!drawer.mini">
            <template v-slot:activator="{ on, attrs }">
              <v-list-item
                class="menu px-3"
                :to="item.to"
                router
                exact
                v-bind="attrs"
                v-on="on"
              >
                <v-list-item-action>
                  <v-icon>{{ item.icon }}</v-icon>
                </v-list-item-action>
                <v-list-item-content>
                  <v-list-item-title class="fw-09">{{
                    item.title
                  }}</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
            </template>
            <span v-text="item.title"></span>
          </v-tooltip>
        </template>
      </template>
    </v-list>
  </v-navigation-drawer>
</template>

<script>
import { mapState, mapActions } from 'vuex'
import logoSM from '@/assets/auto-360-logo-sm.png'
import logo from '@/assets/auto-360-logo-md.png'
export default {
  name: 'AppDrawer',
  components: {
    BaseTitle: () => import('~shared/components/base/Title'),
    BaseDivider: () => import('~shared/components/base/Divider'),
  },

  props: {
    defaultDrawer: false,
  },
  data() {
    return {
      logoSM,
      logo,
    }
  },
  computed: {
    ...mapState({
      drawer: (state) => state.layout.drawer,
      menu: (state) => {
        let old_menu = JSON.parse(JSON.stringify(state.layout.menu))
        let new_menu = { items: [] }
        if (old_menu !== undefined) {
          old_menu.items.forEach((each) => {
            if (each.children !== undefined) {
              let new_children = []
              each.children.forEach((child) => {
                if (child.user_kind.includes('appowner')) {
                  new_children.push(child)
                }
              })
              each.children = new_children
              if (each.children.length > 0) {
                new_menu.items.push(each)
              }
            } else {
              if (each.user_kind.includes('appowner')) {
                new_menu.items.push(each)
              }
            }
          })
        }
        return new_menu
      },
    }),
    showDrawer: {
      get() {
        return this.drawer.show
      },
      set(nv) {
        this.$emit('input', nv)
      },
    },
  },

  methods: {
    ...mapActions('layout', ['setDrawer']),
    checkIncludes(menu, userkind) {
      if (userkind.kind !== undefined && menu.user_kind !== undefined) {
        // console.log(menu.user_kind.includes(userkind.kind), menu.module)
        return menu.user_kind.includes(userkind.kind)
      }
    },
  },
}
</script>

<style lang="scss">
.v-navigation-drawer {
  &__content {
    scrollbar-width: thin;

    // track
    &::-webkit-scrollbar-track {
      background: #f0f0f0;
    }

    // handle
    &::-webkit-scrollbar-thumb {
      background: #cdcdcd;
      border-radius: 50px;
    }

    // scrollbar
    &::-webkit-scrollbar {
      width: 5px;
      background-color: #f0f0f0;
    }

    // handle on hover
    &::-webkit-scrollbar-thumb:hover {
      background: #a6a6a6;
    }
  }
}
</style>
<style lang="scss" scoped>
.v-navigation-drawer {
  .v-subheader {
    text-transform: uppercase;
    font-family: $heading-font-family;
  }

  .v-list-item {
    &.menu {
      opacity: 0.9;
      border-radius: 0;

      .v-list-item__action {
        margin-right: 18px;
      }

      .v-icon {
        padding: 3px;
        padding: 3px;
        margin-left: -2px;
      }
    }

    &.logo {
      border-radius: 0;
      height: 64px;
    }

    &__content {
      .v-list-item__title {
        font-weight: 700;
      }
    }

    &--active.menu {
      opacity: 1;
      background-color: #2d4888;

      .v-list-item__content {
        margin-left: -3px;
      }

      &::after {
        opacity: 1;
        // background: var(--v-accent-lighten2);
        background: #f7ab1a;
        width: 5px;
        min-height: 50%;
        border-radius: 2px;
        position: absolute;
        right: 0px;
        left: auto;
      }

      .v-icon {
        border-radius: 5px;
        // color: var(--v-accent-lighten2);
        color: #f7ab1a;
      }
    }

    &--link {
      &::before {
        background: transparent;
        border-radius: 0;
      }
    }
  }
}
</style>

<style>
.v-navigation-drawer__border {
  background: transparent !important;
}
.fw-09 {
  font-size: 0.9125rem !important;
}
</style>
