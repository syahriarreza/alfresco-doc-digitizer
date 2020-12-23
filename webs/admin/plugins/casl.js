import { Ability } from '@casl/ability'
import { abilitiesPlugin } from '@casl/vue'
import Vue from 'vue'

const ability = new Ability()
Vue.use(abilitiesPlugin, ability)

export default function ({ store }, inject) {
  inject('ability', ability)
}
