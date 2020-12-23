export default ({ app, store, $config, $dialog }, inject) => {
  const tool = {
    int2time(i) {
      let h = Math.floor(i / 100)
      let m = i % 100

      let res = h < 10 ? '0' + h.toString() : h.toString()
      res += ':'
      res += m < 10 ? '0' + m.toString() : m.toString()
      return res
    },

    time2int(s) {
      let ss = s.split(':')
      if (ss.length != 2) {
        return 0
      }
      return parseInt(ss[0]) * 100 + parseInt(ss[1])
    },

    date2db(d) {
      return app.$moment(d).format('YYYY-MM-DDT00:00:00Z')
    },

    dateTime2db(d) {
      return app.$moment(d).format('YYYY-MM-DDTHH:mm:00Z')
    },

    humanizeDateTime(d) {
      return app.$moment(d).format('DD MMM YYYY HH:mm')
    },

    dateRelative: (date) => {
      return date ? app.$moment(date).fromNow() : '-'
    },

    tableSaveAttr(item, field, headers, saveAttrURL) {
      let value = item[field]
      let parm = { value: {} }
      let header = headers.find((h) => {
        return h.value == field
      })

      parm['_id'] = item['_id']
      if (!!header && header.type == 'time') {
        parm.value[field] = this.time2int(value)
      } else {
        parm.value[field] = value
      }
      console.log(
        `update data: '${value}' with reference: '${parm['_id']}|${field}'`
      )
      return app.$axios.post(saveAttrURL, parm)
    },

    uuid() {
      return ([1e7] + -1e3 + -4e3 + -8e3 + -1e11).replace(/[018]/g, (c) =>
        (
          c ^
          (crypto.getRandomValues(new Uint8Array(1))[0] & (15 >> (c / 4)))
        ).toString(16)
      )
    },

    info(msg) {
      store.commit('setInfoText', msg)
    },

    error(msg) {
      this.handleError(msg)
    },

    handleError(e) {
      let msg = ''
      if (e.response && e.response.data) {
        msg = e.response.data
      } else {
        msg = e
      }
      //console.log('shit happened boss, fail: ', msg)
      $dialog.notify.error(`Error : ${msg}`)
      store.commit('setErrorText', msg)
    },

    formatMoney(number, options) {
      let { thouSep, decSep, decimal } = options || {}
      thouSep = thouSep || ','
      decSep = decSep || '.'
      decimal = isNaN(decimal) ? 0 : Math.abs(decimal)

      let _nSplits = (Number(number) || 0).toFixed(decimal).split('.')
      let _numbStr = _nSplits[0]
      let _decStr = _nSplits.length > 1 ? _nSplits[1] : ''
      _numbStr = this.addNumbSep(_numbStr, thouSep)

      return _numbStr + (_decStr ? decSep + _decStr : '')
    },

    addNumbSep(number, thouSep) {
      return number.replace(/(\d)(?=(\d{3})+(?!\d))/g, '$1' + thouSep)
    },

    getAssetURL(assetID, isDownload) {
      if (!assetID) {
        return ''
      }
      let appendType = isDownload ? '&t=dl' : ''
      return `${$config.api.baseUrl}${$config.api.asset}view?id=${assetID}${appendType}`
    },

    formatDateFromTo(dateFrom, dateTo) {
      let from = app.$moment(dateFrom)
      let to = app.$moment(dateTo)
      if (from.format('MMMM') == to.format('MMMM')) {
        return from.format('D') + '-' + to.format('D MMMM YYYY')
      }
      return from.format('D MMMM YYYY') + ' - ' + to.format('D MMMM YYYY')
    },

    formatDenomination(number) {
      const denomList = ['rb', 'jt', 'm', 't']
      let denom = ''
      let res = 0
      number = typeof number == 'number' ? number : parseFloat(number)

      for (let denIdx in denomList) {
        let den = denomList[denIdx]
        res = number / 1000
        if (res < 1) {
          return { number: number, denom: denom }
        }

        denom = den
        number = res
      }

      return { number: number, denom: denom }
    },

    dateFormat(d, f) {
      return app.$moment(d).format(f)
    },

    popupWindow(pageURL, pageTitle, width, height) {
      let left = (screen.width - width) / 2
      let top = (screen.height - height) / 4

      return window.open(
        pageURL,
        pageTitle,
        `resizable=yes, width=${width}, height=${height}, top=${top}, left=${left}`
      )
    },

    slugText(title) {
      var slug = title.toLowerCase()
      return slug.replace(/ /g, '-')
    },

    randomString(len) {
      var charSet =
        'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
      var randomString = ''
      for (var i = 0; i < len; i++) {
        var randomPoz = Math.floor(Math.random() * charSet.length)
        randomString += charSet.substring(randomPoz, randomPoz + 1)
      }
      return randomString
    },

    uploadAsset(file, onSuccess, onError, onEmpty) {
      if (file) {
        let filename = 'TMP__' + file.name
        const reader = new FileReader()
        reader.addEventListener('load', () => {
          app.$axios
            .post('asset/write', {
              asset: { originalfilename: filename },
              content: btoa(reader.result),
            })
            .then((r) => {
              if (onSuccess && typeof onSuccess == 'function') {
                onSuccess(r)
              }
            })
            .catch((e) => {
              if (onError && typeof onError == 'function') {
                onError(e)
              }
            })
        })
        reader.readAsBinaryString(file)
      } else {
        if (onEmpty && typeof onEmpty == 'function') {
          onEmpty()
        }
      }
    },

    truncate(input, n) {
      n = n || 10
      if (input.length > n) {
        return input.substring(0, n) + '...'
      }
      return input
    },

    setPageLength(total, round) {
      return Math.ceil(total / round)
    },

    removePhoneCountryCode(phoneNumber) {
      let mobile = phoneNumber.replaceAll(' ', '')

      // If string starts with +, drop first 3 characters
      if (phoneNumber.slice(0, 1) == '+') {
        mobile = mobile.substring(3)
      }

      // If string starts with 0, drop first character
      if (phoneNumber.slice(0, 1) == '0') {
        mobile = mobile.substring(1)
      }

      return mobile
    },

    ensureCountryCode(phoneNumber) {
      let mobile = phoneNumber.replaceAll(' ', '').replaceAll('-', '')

      if (mobile.slice(0, 1) == '0') {
        mobile = '+62' + mobile.substring(1)
      } else if (mobile.slice(0, 1) != '0' && mobile.slice(0, 1) != '+') {
        mobile = '+62' + mobile
      }

      return mobile
    },

    getYoutubeVideoID(url) {
      var ID = ''
      url = url
        .replace(/(>|<)/gi, '')
        .split(/(vi\/|v=|\/v\/|youtu\.be\/|\/embed\/)/)
      if (url[2] !== undefined) {
        ID = url[2].split(/[^0-9a-z_\-]/i)
        ID = ID[0]
      } else {
        ID = url
      }
      return ID
    },

    ytThumbnail(url) {
      let videoID = this.getYoutubeVideoID(url)
      return `https://img.youtube.com/vi/${videoID}/sddefault.jpg`
    },

    /*!
     * Get the contrasting color for any hex color
     * (c) 2019 Chris Ferdinandi, MIT License, https://gomakethings.com
     * Derived from work by Brian Suda, https://24ways.org/2010/calculating-color-contrast/
     * @param  {String} A hexcolor value
     * @return {String} The contrasting color (black or white)
     */
    isDark(hexcolor) {
      // If a leading # is provided, remove it
      if (hexcolor.slice(0, 1) === '#') {
        hexcolor = hexcolor.slice(1)
      }

      // If a three-character hexcode, make six-character
      if (hexcolor.length === 3) {
        hexcolor = hexcolor
          .split('')
          .map(function (hex) {
            return hex + hex
          })
          .join('')
      }

      // Convert to RGB value
      var r = parseInt(hexcolor.substr(0, 2), 16)
      var g = parseInt(hexcolor.substr(2, 2), 16)
      var b = parseInt(hexcolor.substr(4, 2), 16)

      // Get YIQ ratio
      var yiq = (r * 299 + g * 587 + b * 114) / 1000

      // Check contrast
      return yiq >= 128
    },
  }

  inject('tool', tool)
}
