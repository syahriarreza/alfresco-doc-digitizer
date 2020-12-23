<template>
  <base-dialog :dialog="dialogRegister" @change="getValue" :title="txtDialog" max-width="350px">
    <template v-slot:body>
      <v-form ref="form" v-model="valid" lazy-validation>
        <v-row>
          <v-col cols="12" class="no-padding height-group">
            <base-title class="mb-0">Name</base-title>
            <base-text-field label="Name*" v-model="name" :rules="nameRules" required></base-text-field>
          </v-col>
          <v-col cols="12" class="no-padding height-group">
            <base-title class="mb-0">Email</base-title>
            <base-text-field label="Email*" v-model="email" :rules="emailRules" required></base-text-field>
          </v-col>
          <v-col cols="12" class="no-padding height-group">
            <base-title class="mb-0">Password</base-title>
            <base-text-field
              label="Password*"
              type="password"
              v-model="password"
              :rules="passwordRules"
              required
            ></base-text-field>
          </v-col>
          <v-col cols="12" class="no-padding height-group">
            <base-title class="mb-0">Reenter Password</base-title>
            <base-text-field
              label="Password*"
              type="password"
              v-model="reenterPassword"
              :rules="reenterPassworddRules"
              required
            ></base-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col col="12" class="padding-12-0">
            <base-btn
              :minWidth="80"
              height="30"
              width="100%"
              :tile="false"
              @click="register()"
              :btnClass="classBtnGroup"
              :xlarge="xlarge"
            >Register</base-btn>
          </v-col>
        </v-row>
      </v-form>
    </template>
  </base-dialog>
</template>

<script>
export default {
  props: {
    dialog: { type: Boolean, default: false }
  },
  components: {
    BaseDialog: () => import("../base/Dialog"),
    BaseTextField: () => import("../base/TextField"),
    BaseBtn: () => import("../base/Btn"),
    BaseTitle: () => import("../base/Title")
  },
  data() {
    return {
      valid: true,
      dialogRegister: false,
      classBtnGroup: "font-weight-bold btn-custome",
      xlarge: false,
      txtDialog: "Register",
      name: "",
      nameRules: [
        v => !!v || "User Name is required",
        v =>
          (v && v.length <= 10) || "User Name must be less than 10 characters"
      ],
      password: "",
      reenterPassword: "",
      passwordRules: [v => !!v || "Password is required"],
      reenterPassworddRules: [v => !!v || "Password is required"],
      email: "",
      emailRules: [
        v => !!v || "E-mail is required",
        v => /.+@.+\..+/.test(v) || "E-mail must be valid"
      ],
      swal: this.$store.state.account.swal
    };
  },
  watch: {
    dialog(newValue, oldValue) {
      this.dialogRegister = this.dialog;
    }
  },
  methods: {
    getValue(newValue) {
      this.$emit("change", newValue);
    },
    register() {
      if (this.$refs.form.validate()) {
        if (this.password != this.reenterPassword) {
          this.reenterPassworddRules = ["Password not Match"];
        } else {
          let self = this;
          this.$axios
            .post("/account/register", {
              email: self.email,
              name: self.name,
              password: self.password
            })
            .then(
              r => {
                this.$emit("change", false);
                this.$store.dispatch("account/fetchSwal", {
                  color: "success",
                  mode: "",
                  snackbar: true,
                  text: "Register Success",
                  timeout: 3000,
                  x: null,
                  y: "top"
                });
              },
              e => {
                this.$store.dispatch("account/fetchSwal", {
                  color: "",
                  mode: "",
                  snackbar: true,
                  text: "Password not Match",
                  timeout: 3000,
                  x: null,
                  y: "top"
                });
              }
            );
        }
      }
    },
    resetForm() {}
  },
  computed: {}
};
</script>
<style scoped>
.mb-0 {
  margin-bottom: 0px !important;
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
</style>