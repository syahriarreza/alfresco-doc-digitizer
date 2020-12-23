export const state = () => ({
  auth: '',
  account: {
    isLoading: false,
    data: {},
    kind: {},
  },
})

export const mutations = {
}

export const actions = {
  fetchAccount({ commit }) {
    
  },
  authenticate({ commit, dispatch }) {
    return 'appowner'
  },
  updateRules({ commit }, rules) {
    console.log('1.', this)
    console.log('2.', this.$ability)
    console.log('3.', rules)
  },
}
