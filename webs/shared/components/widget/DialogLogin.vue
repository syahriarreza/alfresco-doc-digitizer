<template>
  <base-dialog :dialog="dialogLogin" @change="getValue" :title="txtDialog" max-width="300">
    <template v-slot:body>
      <v-form ref="form" v-model="valid" lazy-validation>
        <v-row>
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
          <v-col cols="12" class="no-padding height20">
            <small>*Forgot Password ?</small>
            <small class="primary--text pointer" @click="forgotPassword()">Click Here</small>
          </v-col>
        </v-row>
        <v-row>
          <v-col col="12" class="padding-12-0">
            <base-btn
              :minWidth="80"
              height="30"
              width="100%"
              :tile="false"
              @click="login()"
              :btnClass="classBtnGroup"
              :xlarge="xlarge"
            >Login</base-btn>
          </v-col>
        </v-row>
        <v-row>
          <v-col col="12" class="padding-12-0 text-center">
            <small class="primary--text pointer" @click="createAccount()">Create Account</small>
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
    BaseDialog: () => import('../base/Dialog'),
    BaseTextField: () => import('../base/TextField'),
    BaseBtn: () => import('../base/Btn'),
    BaseTitle: () => import('../base/Title'),
  },
  data() {
    return {
      valid: '',
      dialogLogin: false,
      classBtnGroup: 'font-weight-bold btn-custome',
      xlarge: false,
      txtDialog: 'Login',
      password: '',
      passwordRules: [(v) => !!v || 'Password is required'],
      email: '',
      emailRules: [
        (v) => !!v || 'E-mail is required',
        (v) => /.+@.+\..+/.test(v) || 'E-mail must be valid',
      ],
    }
  },
  watch: {
    dialog(newValue, oldValue) {
      this.dialogLogin = this.dialog
    },
  },
  methods: {
    getValue(newValue) {
      this.$emit('change', newValue)
    },
    login() {
      if (this.$refs.form.validate()) {
        let self = this
        this.$axios({
          method: 'POST',
          url: '/account/login',
          auth: {
            username: this.email,
            password: this.password,
          },
        }).then(
          (r) => {
            this.$emit('change', false)
            this.$store.dispatch('account/fetchAuth', r.data)
            this.$store.dispatch('account/fetchSwal', {
              color: 'success',
              mode: '',
              snackbar: true,
              text: 'Success Login',
              timeout: 3000,
              x: null,
              y: 'top',
            })
            this.$store.dispatch('account/fetchCurrent')
          },

          (e) => {
            this.$store.dispatch('account/fetchSwal', {
              color: 'error',
              mode: '',
              snackbar: true,
              text: 'Error Login',
              timeout: 3000,
              x: null,
              y: 'top',
            })
          }
        )
      }
    },
    createAccount() {
      this.$emit('change', false)
      this.$emit('register', true)
    },
    forgotPassword() {
      this.$emit('change', false)
      this.$emit('forgot', true)
    },
    resetForm() {},
  },
  computed: {
    auth() {
      return this.$store.state.account.auth
    },
  },
}
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