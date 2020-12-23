export default function ({ store, app }) {
  const auth = app.$cookie.get('auth')
  store.dispatch("account/fetchAuth", auth);
}