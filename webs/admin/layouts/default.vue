<template>
  <v-app>
    <app-drawer :defaultDrawer="!drawer.show"></app-drawer>
    <v-app-bar
      class="pl-md-5 pl-lg-5 pl-xl-5"
      elevation="0"
      color="#f6f4fc"
      app
    >
      <template v-if="false" v-slot:extension>
        <v-tabs class="ml-3" align-with-title>
          <v-tab>Tab 1</v-tab>
          <v-tab>Tab 2</v-tab>
          <v-tab>Tab 3</v-tab>
        </v-tabs>
      </template>
      <v-app-bar-nav-icon
        color="primary"
        @click.stop="setDrawer({ show: !drawer.show })"
      />
      <v-toolbar-title class="text-uppercase">{{
        pageInfo.title
      }}</v-toolbar-title>
      <span class="text--disabled mt-1 ml-2">{{ pageInfo.subtitle }}</span>
      <v-spacer />
      <v-menu bottom left min-width="150">
        <template v-slot:activator="{ on: onMenu }">
          <v-tooltip bottom>
            <template v-slot:activator="{ on: onTool }">
              <v-btn icon v-on="{ ...onMenu, ...onTool }">
                <v-avatar>
                  <v-icon color="primary">mdi-cog</v-icon>
                </v-avatar>
              </v-btn>
            </template>
            <span>Settings</span>
          </v-tooltip>
        </template>
        <v-list offset-y dense max-width="250px">
          <v-list-item :to="'/account/changepassword'">
            <v-list-item-icon>
              <v-icon>mdi-lock</v-icon>
            </v-list-item-icon>

            <v-list-item-content>
              <v-list-item-title>Ubah Password</v-list-item-title>
            </v-list-item-content>
          </v-list-item>

          <v-divider />
          <v-list-item :to="'/account/profile'">
            <v-list-item-icon>
              <v-icon>mdi-account</v-icon>
            </v-list-item-icon>

            <v-list-item-content>
              <v-list-item-title>Edit Profile</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-menu>
      <v-tooltip bottom>
        <template v-slot:activator="{ on, attrs }">
          <v-btn icon @click="logout" v-on="on" v-bind="attrs">
            <v-avatar>
              <v-icon color="primary">mdi-logout</v-icon>
            </v-avatar>
          </v-btn>
        </template>
        <span>Logout</span>
      </v-tooltip>
    </v-app-bar>
    <v-main>
      <v-container fluid class="pl-md-7 pt-0 pl-lg-7 pl-xl-7">
        <v-row>
          <v-col cols="12">
            <nuxt />
          </v-col>
        </v-row>
      </v-container>
    </v-main>
    <v-footer inset app>
      <span>&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>

<script>
import { mapState, mapActions } from 'vuex'
import { AbilityBuilder } from '@casl/ability'
export default {
  async middleware({ store, redirect, app, route }) {
    try {
      const { breadcrumbs, title, subtitle, module, submodule, backTo } =
        route.meta.length > 0 ? route.meta[0] : {}
      let res = await store.dispatch('account/authenticate')
      // await store.dispatch('account/fetchAccount')
      const { can, rules } = new AbilityBuilder()

      switch (res.data) {
        case 'appowner':
          can('access', 'all')
          can('complaint', 'close')
          can('complaint', 'refund')
          break
        case 'eventorganizer':
          can('access', 'all')
          break
        case 'tenant':
          can('access', 'all')
          break
        case 'visitor':
          can('complaint', 'close')
          can('access', 'all')
          break
        default:
          break
      }

      app.context.$ability.update(rules)
    } catch (error) {
      if (!!error.response && !!error.response.data && redirect) {
        return redirect('/account/login')
      }
    }
  },
  components: {
    AppDrawer: () => import('@/components/layout/AppDrawer'),
  },
  data() {
    return {
      isOnDevEnv: false,
      title: 'Vuetify.js',
    }
  },
  computed: {
    ...mapState({
      drawer: (state) => state.layout.drawer,
      pageInfo: (state) => state.layout.pageInfo,
    }),
  },
  methods: {
    logout(){},
  },
  mounted() {},
}
</script>

<style scoped>
</style>
