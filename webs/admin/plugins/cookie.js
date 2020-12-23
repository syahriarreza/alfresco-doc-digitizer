export default function ({ store, app }) {
  const auth = app.$cookie.get('auth_admin')
  // store.commit('account/SET_AUTH', auth)
}
