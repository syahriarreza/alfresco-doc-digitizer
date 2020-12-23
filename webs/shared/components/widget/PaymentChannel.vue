<template>
  <v-card>
    <v-card-title>
      <h4>Payment</h4>
    </v-card-title>
    <v-divider />
    <v-card-text class="pa-0">
      <v-tabs vertical hide-slider background-color="#172b79">
        <v-tab>
          <h4>Bank Transfer</h4>
        </v-tab>
        <v-tab>
          <h4>Virtual Account</h4>
        </v-tab>
        <v-tab>
          <h4>Internet Banking</h4>
        </v-tab>

        <v-tab-item>
          <v-card flat>
            <v-card-title>
              <h3>Bank Transfer</h3>
            </v-card-title>
            <v-card-text>
              <v-sheet
                v-for="(met, metIdx) in paymentMethods[0].list"
                :key="metIdx"
                :elevation="0"
                outlined
                rounded
                class="pa-3 my-2"
              >
                <v-radio :label="met.name" :value="met.isSelected"></v-radio>
              </v-sheet>
            </v-card-text>
            <v-divider />
            <v-card-text>
              <v-row>
                <v-col>
                  <h3 class="color-black">Price Details</h3>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="8">
                  <h3>
                    <span class="color-black">4x</span> GIIAS Jakarta - Personal Tickets
                  </h3>
                </v-col>
                <v-col cols="4" class="pr-5">
                  <h3 class="text-right">Rp 100.000</h3>
                </v-col>
              </v-row>
            </v-card-text>
            <v-divider class="mx-auto" width="93%"></v-divider>
            <v-card-text>
              <v-row>
                <v-col cols="8">
                  <h3 class="color-black">Total</h3>
                </v-col>
                <v-col cols="4" class="pr-5">
                  <h3 class="text-right color-black">Rp 400.000</h3>
                </v-col>
              </v-row>
            </v-card-text>
          </v-card>
          <v-row class="px-4">
            <v-col cols="8">
              <span
                class="color-black"
              >By clicking this button, you acknowledge that you have read and agreed to the Terms & Conditions and Privacy Policy of GAIKINDO</span>
            </v-col>
            <v-col cols="4">
              <base-btn width="90%" height="30" class="btn-buy accent float-right rounded-pill">Pay</base-btn>
            </v-col>
          </v-row>
        </v-tab-item>
        <v-tab-item>
          <v-card flat>
            <v-card-text>
              <p>Morbi nec metus. Suspendisse faucibus, nunc et pellentesque egestas, lacus ante convallis tellus, vitae iaculis lacus elit id tortor. Sed mollis, eros et ultrices tempus, mauris ipsum aliquam libero, non adipiscing dolor urna a orci. Curabitur ligula sapien, tincidunt non, euismod vitae, posuere imperdiet, leo. Nunc sed turpis.</p>

              <p>Suspendisse feugiat. Suspendisse faucibus, nunc et pellentesque egestas, lacus ante convallis tellus, vitae iaculis lacus elit id tortor. Proin viverra, ligula sit amet ultrices semper, ligula arcu tristique sapien, a accumsan nisi mauris ac eros. In hac habitasse platea dictumst. Fusce ac felis sit amet ligula pharetra condimentum.</p>

              <p>Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, quis gravida magna mi a libero. Nam commodo suscipit quam. In consectetuer turpis ut velit. Sed cursus turpis vitae tortor. Aliquam eu nunc.</p>

              <p>Etiam ut purus mattis mauris sodales aliquam. Ut varius tincidunt libero. Aenean viverra rhoncus pede. Duis leo. Fusce fermentum odio nec arcu.</p>

              <p
                class="mb-0"
              >Donec venenatis vulputate lorem. Aenean viverra rhoncus pede. In dui magna, posuere eget, vestibulum et, tempor auctor, justo. Fusce commodo aliquam arcu. Suspendisse enim turpis, dictum sed, iaculis a, condimentum nec, nisi.</p>
            </v-card-text>
          </v-card>
        </v-tab-item>
        <v-tab-item>
          <v-card flat>
            <v-card-text>
              <p>Fusce a quam. Phasellus nec sem in justo pellentesque facilisis. Nam eget dui. Proin viverra, ligula sit amet ultrices semper, ligula arcu tristique sapien, a accumsan nisi mauris ac eros. In dui magna, posuere eget, vestibulum et, tempor auctor, justo.</p>

              <p
                class="mb-0"
              >Cras sagittis. Phasellus nec sem in justo pellentesque facilisis. Proin sapien ipsum, porta a, auctor quis, euismod ut, mi. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nam at tortor in tellus interdum sagittis.</p>
            </v-card-text>
          </v-card>
        </v-tab-item>
      </v-tabs>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  components: {
    BaseSubheading: () => import("~shared/components/base/Subheading"),
    BaseBoxDisplayItem: () => import("~shared/components/base/BoxDisplayItem"),
    BaseBody: () => import("~shared/components/base/Body"),
    BaseView: () => import("~shared/components/base/View"),
    BaseBtn: () => import("~shared/components/base/Btn"),
  },
  data() {
    return {
      paymentMethods: [
        {
          name: "Bank Transfer",
          list: [
            {
              name: "BCA Transfer",
              src: "",
              isSelected: false,
            },
            {
              name: "Mandiri Transfer",
              src: "",
              isSelected: false,
            },
            {
              name: "BNI Transfer",
              src: "",
              isSelected: false,
            },
            {
              name: "BRI Transfer",
              src: "",
              isSelected: false,
            },
          ],
        },
      ],
    };
  },
};
</script>

<style scoped>
.v-tab--active {
  background-color: white;
}
.v-tab:not(.v-tab--active) {
  color: white !important;
}

.color-black {
  color: black;
}
</style>