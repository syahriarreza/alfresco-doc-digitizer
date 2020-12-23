import { extend, ValidationObserver, ValidationProvider } from 'vee-validate'
import { messages } from 'vee-validate/dist/locale/id.json'
import * as rules from 'vee-validate/dist/rules'
import Vue from 'vue'

// overriode validation message with locale id message
Object.keys(rules).forEach((rule) => {
  extend(rule, {
    ...rules[rule],
    message: messages[rule],
  })
})

// if you want to add custom rule validation, add here
extend('equal_with', {
  params: ['target'],
  validate(value, { target }) {
    return value === target
  },
  message: (field) => {
    return `nilai tidak sama dengan ${field}`
  },
})
extend('is_url', {
  validate(value) {
    // source: https://stackoverflow.com/questions/5717093/check-if-a-javascript-string-is-a-url
    let pattern = new RegExp(
      '^(https?:\\/\\/)?' + // protocol
      '((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.)+[a-z]{2,}|' + // domain name
      '((\\d{1,3}\\.){3}\\d{1,3}))' + // OR ip (v4) address
      '(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*' + // port and path
      '(\\?[;&a-z\\d%_.~+=-]*)?' + // query string
        '(\\#[-a-z\\d_]*)?$',
      'i'
    ) // fragment locator
    return !!pattern.test(value)
  },
  message: (field) => {
    return `invalid URL: ${field}`
  },
})

Vue.component('ValidationProvider', ValidationProvider)
Vue.component('ValidationObserver', ValidationObserver)
