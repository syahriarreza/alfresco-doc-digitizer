<template>
  <v-responsive>
    <v-row>
      <base-dialog :dialog="dialogShowHide" @hideDialog="refShowHideDialog">
        <template v-slot:append-outer>
          <div class="container">
            <v-card elevation="24" class="mx-auto">
              <v-system-bar lights-out></v-system-bar>
              <v-carousel :show-arrows="true">
                <v-carousel-item v-for="(item,i) in product" :key="i" :src="item.src"></v-carousel-item>
              </v-carousel>
              <v-list two-line>
                <v-list-item>
                  <v-list-item-content>
                    <base-title :title="'Mark Mobil'" class="text-uppercase" space="1" />
                  </v-list-item-content>
                  <base-subtitle>Lorem ipsum dolor sit amet, voluptua iracundia disputationi an pri, his utinam principes dignissim ad. Ne nec dolore oblique nusquam, cu luptatum volutpat delicatissimi has. Sed ad dicam platonem, mea eros illum elitr id, ei has similique constituto. Ea movet saperet rationibus sit, pri autem aliquip invidunt an. Consetetur omittantur consequuntur eos et. Eleifend praesent iudicabit no mea, tollit persequeris ex pri, tota splendide voluptaria in pri. Ad per tale aliquip, ei sit viris commune albucius. Eos aliquip scaevola ut, eum alii mentitum prodesset no, his ne suas atomorum. Et numquam deleniti ponderum vis, quod error at mei. Novum blandit adolescens sea te. Ea eum cetero scaevola.In his meis porro viris, illud imperdiet reprimique et vim. Feugiat atomorum reprehendunt vix ei, ei facete regione pri. Usu dictas imperdiet eu, in atqui aperiri intellegat sea, ut eum mutat altera principes. Te sit quaeque oportere, has modus inani ceteros ad. Impedit blandit deseruisse duo ea, ne graecis deleniti incorrupte usu. Ut mei splendide accommodare. An pri iisque meliore, eam ei splendide eloquentiam philosophia. Ne per meis eleifend electram. Ne eam porro aliquam invidunt. Minim docendi eloquentiam cum ad. Quo ea mazim ubique, ex est fuisset blandit scaevola. Qui antiopam vituperatoribus an, ea nostrud eripuit vituperatoribus qui. In eam diam nominati, per ea alia luptatum. Nam habemus electram democritum ut. Mei ea omnium admodum intellegat. Habeo atqui molestiae at mei, an nec ridens consequuntur. Quem nulla cum ei, his ipsum apeirian no, per at eius iriure aperiri. Sed dicam interesset ei. Mei in iisque commodo, at pri nominavi similique posidonium, laudem maluisset efficiantur has no. His verterem consectetuer consequuntur ne, no virtute atomorum usu. Eu quo nemore causae tacimates, eos viderer persequeris an. Cu molestie consulatu qui. Natum labores perfecto no ius, pri dico mundi inciderint id. Ei usu dico libris postea. Cu graeco doctus splendide qui, ei eum probo regione.</base-subtitle>
                </v-list-item>
              </v-list>
            </v-card>
          </div>
        </template>
      </base-dialog>
    </v-row>
  </v-responsive>
</template>


<script>
export default {
  props: {},
  components: {
    BaseDialog: () => import("./DialogFullscreen"),
    BaseTitle: () => import("./Title"),
    BaseSubtitle: () => import("./Subtitle")
  },
  data() {
    return {
      colors: [
        "green",
        "secondary",
        "yellow darken-4",
        "red lighten-2",
        "orange darken-1"
      ],
      cycle: false,
      slides: ["First", "Second", "Third", "Fourth", "Fifth"],

      dialogShowHide: false,
      notifications: false,
      sound: true,
      widgets: false,
      product: [
        {
          title: "",
          src: require("~/static/carlogo/logo-car-1.png"),
          link: "",
          description: "",
          showBorder: true
        },
        {
          title: "",
          src: require("~/static/carlogo/logo-car-2.png"),
          link: "",
          description: "",
          showBorder: true
        },
        {
          title: "",
          src: require("~/static/carlogo/logo-car-3.png"),
          link: "",
          description: "",
          showBorder: true
        },
        {
          title: "",
          src: require("~/static/carlogo/logo-car-4.png"),
          link: "",
          description: "",
          showBorder: true
        }
      ]
    };
  },
  methods: {
    refShowHideDialog(val) {
      this.dialogShowHide = val;
    }
  }
};
</script>