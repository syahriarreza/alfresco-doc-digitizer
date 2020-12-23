<template>
  <base-dialog :dialog="dialogForgot" @change="getValue" :title="txtDialog" max-width="300">
    <template v-slot:body>
      <v-form ref="form" v-model="valid" lazy-validation>
        <v-row>
          <v-col cols="12" class="no-padding height-group">
            <base-title class="mb-0">Email</base-title>
            <base-text-field label="Email*" v-model="email" :rules="emailRules" required></base-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col col="12" class="padding-12-0">
            <base-btn
              :minWidth="80"
              height="30"
              width="100%"
              :tile="false"
              @click="sendEmail()"
              :btnClass="classBtnGroup"
              :xlarge="xlarge"
            >Send Email</base-btn>
          </v-col>
        </v-row>
      </v-form>
    </template>
  </base-dialog>
</template>

<script>
export default {
  props: {
    dialog: { type: Boolean, default: false },
  },
  components: {
    BaseDialog: () => import("../base/Dialog"),
    BaseTextField: () => import("../base/TextField"),
    BaseBtn: () => import("../base/Btn"),
    BaseTitle: () => import("../base/Title"),
  },
  data() {
    return {
      valid: "",
      dialogForgot: false,
      classBtnGroup: "font-weight-bold btn-custome",
      xlarge: false,
      txtDialog: "Forgot Password",
      email: "",
      emailRules: [
        (v) => !!v || "E-mail is required",
        (v) => /.+@.+\..+/.test(v) || "E-mail must be valid",
      ],
    };
  },
  watch: {
    dialog(newValue, oldValue) {
      this.dialogForgot = this.dialog;
    },
  },
  methods: {
    getValue(newValue) {
      this.$emit("change", newValue);
    },
    sendEmail() {
      if (this.$refs.form.validate()) {
        let self = this;
        this.$axios({
          method: "POST",
          url: "/account/forgotpassword",
          auth: {
            username: this.email,
          },
        }).then(
          (r) => {
            this.$emit("change", false);
            this.$store.dispatch("account/fetchSwal", {
              color: "success",
              mode: "",
              snackbar: true,
              text: "Success Login",
              timeout: 3000,
              x: null,
              y: "top",
            });
          },

          (e) => {
            this.$store.dispatch("account/fetchSwal", {
              color: "error",
              mode: "",
              snackbar: true,
              text: "Error Login",
              timeout: 3000,
              x: null,
              y: "top",
            });
          }
        );
      }
    },
    resetForm() {},
  },
  computed: {
    auth() {
      return this.$store.state.account.auth;
    },
  },
};
</script>
<style scoped>
.mb-0 {
  margin-bottom: 0px !important;
}
.pointer {
  cursor: pointer;
}
.no-padding {
  padding: 0px;
}
.height-group {
  height: 90px;
}
.padding-12-0 {
  padding: 12px 0px;
}
.height20 {
  height: 20px;
}
</style>