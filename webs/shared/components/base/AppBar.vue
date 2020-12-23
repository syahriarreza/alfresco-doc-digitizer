<template>
  <div>
    <!-- TODO : Calculate primary color opacity -->
    <!-- TODO : In app link section -->
    <!-- TODO : Nested menu dropdown -->
    <!-- TODO : On menu change nuxt listener -->
    <v-app-bar id="home-app-bar" app :color="color" elevation="1" height="80">
      <base-img :src="srcLogo" class="mr-3 hidden-xs-only" contain height="60" width="1" />

      <!-- <base-img
        :src="srcLogo"
        class="d-none d-sm-flex d-md-flex d-lg-flex d-xl-flex"
        contain
        max-width="128"
        width="100%"
      />-->

      <v-spacer />

      <div>
        <v-tabs class="hidden-sm-and-down" v-model="activeMenuIndex">
          <v-menu v-for="(item, index) in items" :key="index" offset-y>
            <template v-slot:activator="{ attrs, on }">
              <v-tab
                class="white--text"
                @click="movePages(item)"
                v-bind="attrs"
                v-on="on"
              >{{ item.title }}</v-tab>
            </template>

            <v-list v-if="item.submenus.length > 0">
              <v-list-item v-for="(submenu, s) in item.submenus" :key="s" link>
                <v-list-item-title
                  class="text-uppercase sub-menu"
                  v-text="submenu.title"
                  @click="movePages(submenu)"
                ></v-list-item-title>
              </v-list-item>
            </v-list>
          </v-menu>
          <v-btn-toggle v-if="!auth" rounded class="btn-toggle">
            <base-btn
              :color="colorBtn"
              :minWidth="77"
              height="26"
              :tile="false"
              @click="register()"
              :btnClass="classBtnRegister"
              :xlarge="xlarge"
            >Register</base-btn>

            <base-btn
              :color="colorBtn"
              :minWidth="77"
              height="26"
              :tile="false"
              @click="login()"
              :btnClass="classBtnLogin"
              :xlarge="xlarge"
            >Login</base-btn>
          </v-btn-toggle>
          <div v-else class="d-flex">
            <base-btn
              :color="colorBtn"
              :minWidth="77"
              height="30"
              :tile="false"
              :btnClass="'primary--text my-ticket'"
              :xlarge="xlarge"
              @click="myTicket()"
            >My Ticket</base-btn>
            <div class="text-center d-flex">
              <base-title class="ma-2 white--text" v-text="aboutMe.name"></base-title>
              <v-menu
                v-model="optionProf"
                :close-on-content-click="false"
                :nudge-width="200"
                offset-y
              >
                <template v-slot:activator="{ on, attrs }">
                  <v-avatar :color="'white'" class="base-avatar__avatar" v-bind="attrs" v-on="on">
                    <base-icon :dark="false">{{ icon }}</base-icon>
                  </v-avatar>
                  <!-- <img
                    :src="require('@/static/logo-giias.png')"
                    cover
                    width="45"
                    height="45"
                    class="img-user"
                    v-bind="attrs"
                    v-on="on"
                  />-->
                </template>
                <v-card class="card-profile">
                  <v-list>
                    <v-list-item>
                      <!-- <v-list-item-avatar>
                        <img :src="require('@/static/logo-giias.png')" alt="Giias" />
                      </v-list-item-avatar>-->
                      <v-avatar :color="'white'" class="base-avatar__avatar mr-3">
                        <base-icon :dark="false">{{ icon }}</base-icon>
                      </v-avatar>
                      <v-list-item-content>
                        <v-list-item-title v-text="aboutMe.name">Giias 2020</v-list-item-title>
                        <v-list-item-subtitle v-text="aboutMe.email">giias@gmail.com</v-list-item-subtitle>
                      </v-list-item-content>
                    </v-list-item>
                  </v-list>

                  <v-divider></v-divider>
                  <v-list
                    :disabled="false"
                    :dense="false"
                    :two-line="false"
                    :three-line="false"
                    :shaped="false"
                    :flat="false"
                    :subheader="false"
                    :sub-group="false"
                    :nav="true"
                    :avatar="true"
                    :rounded="false"
                  >
                    <v-list-item-group color="primary">
                      <v-list-item
                        v-for="(item, i) in itemsProfile"
                        :key="i"
                        :inactive="false"
                        @click="myProfile(item.to)"
                      >
                        <v-list-item-avatar v-if="avatar">
                          <i :class="item.icon" :size="15"></i>
                        </v-list-item-avatar>
                        <v-list-item-content>
                          <v-list-item-title :color="'primary'" v-html="item.title"></v-list-item-title>
                          <v-list-item-subtitle
                            v-if="false || false"
                            v-html="item.subtitle"
                            :color="'primary'"
                          ></v-list-item-subtitle>
                        </v-list-item-content>
                      </v-list-item>
                    </v-list-item-group>
                  </v-list>

                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <!-- <v-btn text @click="optionProf = false">Log Out</v-btn> -->
                    <base-btn
                      :minWidth="80"
                      height="30"
                      :tile="false"
                      @click="logOut()"
                      :btnClass="'font-weight-bold btn-custome'"
                      :xlarge="false"
                    >Log Out</base-btn>
                  </v-card-actions>
                </v-card>
              </v-menu>
            </div>
          </div>
        </v-tabs>
      </div>

      <v-app-bar-nav-icon class="hidden-md-and-up white--text" @click="drawer = !drawer" />
    </v-app-bar>

    <drawer v-model="drawer" :items="items">
      <template v-slot:append-outer>
        <div class="btn-action-drawer">
          <div v-if="auth != ''" class="text-center">
            <base-img
              :src="require('@/static/logo-giias.png')"
              contain
              width="45"
              height="45"
              class="img-user"
            />
            <base-title class="ma-2 white--text text-center" v-text="'Giias 2020'"></base-title>

            <v-btn-toggle rounded class="btn-toggle-drawer">
              <base-btn
                :color="colorBtn"
                :minWidth="77"
                height="26"
                :tile="false"
                :btnClass="classBtnRegister"
                :xlarge="xlarge"
              >My Ticket</base-btn>

              <base-btn
                :color="colorBtn"
                :minWidth="77"
                height="26"
                :tile="false"
                :btnClass="classBtnLogin"
                :xlarge="xlarge"
              >Log Out</base-btn>
            </v-btn-toggle>
          </div>

          <v-btn-toggle v-else rounded class="btn-toggle-drawer">
            <base-btn
              :color="colorBtn"
              :minWidth="77"
              height="26"
              :tile="false"
              @click="register()"
              :btnClass="classBtnRegister"
              :xlarge="xlarge"
            >Register</base-btn>

            <base-btn
              :color="colorBtn"
              :minWidth="77"
              height="26"
              :tile="false"
              @click="login()"
              :btnClass="classBtnLogin"
              :xlarge="xlarge"
            >Login</base-btn>
          </v-btn-toggle>
        </div>
      </template>
    </drawer>
  </div>
</template>

<script>
export default {
  name: "BaseAppBar",

  components: {
    Drawer: () => import("./Drawer"),
    BaseImg: () => import("./Img"),
    BaseBtn: () => import("./Btn"),
    BaseAvatar: () => import("./Avatar"),
    BaseInfo: () => import("./Info"),
    BaseTitle: () => import("./Title"),
    BaseImg: () => import("./Img"),
    BaseIcon: () => import("./Icon"),
  },

  data: () => ({
    icon: "mdi-account-circle",
    activeMenuIndex: 0,
    colorBtn: "",
    drawer: null,
    classBtnLogin: "btn-login primary",
    classBtnRegister: "btn-register btn-no-active",
    xlarge: false,
    optionProf: false,
    menuprofile: "menu-profile",
    avatar: true,
    itemsProfile: [
      {
        icon: "mdi mdi-account icon-prof",
        title: "Edit Profile",
        subtitle: "you can change your profile",
        to: "/profile/",
      },
    ],
  }),

  props: {
    srcLogo: { type: String, default: "" },
    srcLogoSmall: { type: String, default: "" },
    color: { type: String, default: "white" },
  },

  methods: {
    register() {
      this.$emit("register", true);
      this.classBtnLogin = "btn-login btn-no-active primary--text";
      this.classBtnRegister = "btn-register primary white--text";
    },
    login() {
      this.$emit("login", true);
      this.classBtnLogin = "btn-login primary white--text";
      this.classBtnRegister = "btn-register btn-no-active primary--text";
    },
    goSection(val) {
      if (val.section != "") {
        setTimeout(() => {
          this.$vuetify.goTo(val.section);
        }, 200);
      }
    },
    movePages(val) {
      var currentRouter = this.$route.fullPath;
      var self = this;

      if (val.to != "#") {
        if (currentRouter == val.to) {
          if (val.section != "") {
            this.$vuetify.goTo(val.section);
          }
        } else {
          this.$emit("page-change", true);
          this.$router.push(val.to, () => {
            self.goSection(val);
          });
        }
      }
    },
    myTicket() {
      this.$router.push("/myticket/");
    },
    myProfile(url) {
      this.optionProf = false;
      this.$router.push(url);
    },
    logOut() {
      this.optionProf = false;
      this.$store.dispatch("account/fetchAuth", "");
      this.$router.push("/");
    },
  },

  async mounted() {
    await this.$store.dispatch("event/fetchEvents", {});

    let currentRouter = this.$route.fullPath;
    let activeMenuIndex = this.items.findIndex((x) => {
      let result = x.to == currentRouter;
      if (x.submenus.length > 0)
        result = x.submenus.findIndex((x) => x.to == currentRouter) > -1;
      return result;
    });
    this.activeMenuIndex = activeMenuIndex < 0 ? 0 : activeMenuIndex;
  },

  computed: {
    auth() {
      return this.$store.state.account.auth;
    },
    aboutMe() {
      return this.$store.state.account.current;
    },
    items() {
      let _items = [
        {
          to: "/",
          title: "home",
          section: "#hero",
          submenus: [],
        },
        {
          to: "/",
          title: "exhibitor",
          section: "#exhibitor",
          submenus: [],
        },
        {
          to: "/",
          title: "venue",
          section: "#venue",
          submenus: [],
        },
        {
          to: "/",
          title: "ticket",
          section: "#ticket",
          submenus: [],
        },
        {
          to: "/",
          title: "gallery",
          section: "#gallery",
          submenus: [],
        },
        {
          to: "#",
          title: "giias series",
          section: "",
          submenus: [
            {
              to: "/event/1",
              title: "giias surabaya",
              section: "",
              submenus: [],
            },
            {
              to: "/event/1",
              title: "giias medan",
              section: "",
              submenus: [],
            },
            {
              to: "/event/1",
              title: "giias makassar",
              section: "",
              submenus: [],
            },
          ],
        },
      ];
      let events = this.$store.state.event.events;
      let _giiasSeries = _items.find(
        (_) => _.title.toLowerCase() == "giias series"
      );
      _giiasSeries.submenus = events.map((_) => {
        _.to = `/event/${_._id}`;
        _.title = _.name;
        _.section = "";
        _.submenus = [];
        return _;
      });

      return _items;
    },
  },

  watch: {
    // compGetAuth(newValue) {
    //   if (newValue != "") {
    //     this.auth = true;
    //   } else {
    //     this.auth = false;
    //   }
    // }
  },
};
</script>

<style lang="scss">
#home-app-bar {
  .v-tabs {
    & > .v-tabs-bar {
      background: transparent !important;
    }
    .v-tabs-slider {
      max-width: 24px;
      margin: 0 auto;
      color: var(--v-accent-base) !important;
    }
    .v-tab {
      transition: opacity 0.3s ease;
      opacity: 0.5;

      &.v-tab--active {
        opacity: 1;
      }

      &:hover {
        opacity: 1;
      }
      &::before {
        display: none;
      }
    }
  }
}
.btn-toggle {
  height: 30px;
  margin-top: 10px;
}
.btn-login {
  font-size: 10px;
  margin: 2px 2px 0px 2px;
  font-size: 12px !important;
  text-transform: capitalize !important;
}
.btn-register {
  font-size: 10px;
  margin: 2px 0px 2px 2px;
  font-size: 12px !important;
  text-transform: capitalize !important;
}

.btn-no-active {
  background-color: white !important;
  border: 0px !important;
}
.btn-active {
  color: white !important;
}
.btn-toggle-drawer {
  width: auto;
}
.btn-action-drawer {
  text-align: center;
  height: 100%;
  padding: 30px;
  background-color: #373c4c;
}
.my-ticket {
  border-radius: 15px !important;
  text-transform: capitalize !important;
  margin-top: 10px;
  font-size: 12px !important;
}
.img-user {
  border: 1px solid #ffff;
  border-radius: 22.5px;
  display: inline-block !important;
}

.card-profile {
  margin-top: 20px;
}
.icon-prof {
  font-size: 25px;
}
.btn-custome {
  text-transform: capitalize;
}
</style>
