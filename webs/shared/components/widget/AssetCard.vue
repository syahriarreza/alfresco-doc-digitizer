<template>
  <v-card elevation="0">
    <v-card-title class="pb-0 px-0">
      Aset
      <v-spacer />
      <v-btn color="primary" text @click="addAsset" v-if="!hideAddAsset">
        <v-icon>mdi-plus</v-icon>Tambah
      </v-btn>
    </v-card-title>
    <v-card-text class="px-0 py-0">
      <v-alert
        outlined
        color="grey"
        text
        class="mt-4"
        icon="mdi-folder-multiple-image"
        v-if="!content_assets || (content_assets && content_assets.length == 0)"
        >Asset untuk lineup masih belum tersedia, silahkan upload aset
        lineup</v-alert
      >
      <v-slide-group
        v-model="model"
        class="pt-5"
        active-class="success"
        show-arrows
        v-if="content_assets && content_assets.length > 0"
      >
        <v-slide-item
          v-for="(item, i) in content_assets"
          :key="i"
          v-slot:default="{ active }"
        >
          <v-card
            :color="active ? undefined : 'grey lighten-1'"
            class="ma-4"
            height="200"
            width="200"
            @click="editAsset(item, i)"
          >
            <!-- TODO : change image with youtube preview and placeholder loading -->
            <v-img
              class="white--text align-end fill-height"
              :src="
                item.content_asset_kind == 'videoyoutube'
                  ? $tool.ytThumbnail(item.url_source)
                  : $tool.getAssetURL(item.asset_id)
              "
            >
              <template v-slot:placeholder>
                <v-row class="fill-height ma-0" align="center" justify="center">
                  <v-progress-circular
                    indeterminate
                    color="grey lighten-5"
                  ></v-progress-circular>
                </v-row>
              </template>
              <v-card-title>{{ item.title }}</v-card-title>
            </v-img>
          </v-card>
        </v-slide-item>
      </v-slide-group>
    </v-card-text>

    <!-- Aset Input Dialog -->
    <v-dialog
      v-model="showAssetDialog"
      fullscreen
      hide-overlay
      transition="dialog-bottom-transition"
    >
      <v-card tile>
        <v-toolbar dark color="primary">
          <v-btn icon dark @click="closeFormAsset" :disabled="uploading">
            <v-icon>mdi-close</v-icon>
          </v-btn>
          <v-toolbar-title>Tambah Asset</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-toolbar-items>
            <v-btn dark text @click="removeAsset" :disabled="uploading">
              <v-icon left>mdi-trash-can</v-icon>Hapus
            </v-btn>
            <v-btn dark text @click="saveAsset" :loading="uploading">
              <v-icon left>mdi-content-save</v-icon>Simpan
            </v-btn>
          </v-toolbar-items>
        </v-toolbar>
        <v-card-text>
          <ValidationObserver ref="formAssetValidator">
            <v-form v-model="formAsset" ref="formAsset">
              <v-row>
                <v-col cols="9" class="pr-5">
                  <v-row>
                    <v-col cols="6">
                      <KInput
                        v-model="asset.title"
                        rules="required"
                        label="Judul"
                        dense
                        placeholder="Judul media aset"
                      />
                    </v-col>
                    <v-col cols="6">
                      <v-checkbox
                        v-model="asset.is_featured"
                        dense
                        class="compact mx-2"
                        label="Featured"
                      ></v-checkbox>
                    </v-col>
                    <v-col cols="12">
                      <ValidationProvider
                        name="Jenis asset"
                        rules="required"
                        #default="{ errors }"
                      >
                        <v-combobox
                          ref="comboAssetKind"
                          dense
                          v-model="asset.content_asset_kind"
                          :items="assetKinds"
                          item-text="text"
                          :error-messages="errors"
                          item-value="value"
                          label="Jenis"
                          hide-details="auto"
                          :return-object="false"
                        ></v-combobox>
                      </ValidationProvider>
                    </v-col>
                    <v-col cols="12">
                      <ValidationProvider
                        name="Asset"
                        :rules="
                          (asset.content_asset_kind == 'image' &&
                            !asset.asset_id) ||
                          (asset.content_asset_kind == 'video' &&
                            !asset.asset_id)
                            ? 'required'
                            : null
                        "
                        #default="{ errors }"
                      >
                        <v-file-input
                          v-if="asset.content_asset_kind == 'image'"
                          v-model="assetFile"
                          dense
                          accept="image/*"
                          label="Gambar"
                          :error-messages="errors"
                          placeholder="Pilih gambar untuk lineup Anda"
                          prepend-icon="mdi-image"
                          hide-details="auto"
                          @change="uploadAndPreviewFeaturedImage"
                        ></v-file-input>
                        <v-file-input
                          v-if="asset.content_asset_kind == 'video'"
                          dense
                          v-model="assetFile"
                          accept="video/*"
                          label="Video"
                          :error-messages="errors"
                          placeholder="Pilih video untuk lineup Anda"
                          prepend-icon="mdi-video-box"
                          hide-details="auto"
                          @change="uploadAndPreviewFeaturedImage"
                        ></v-file-input>
                      </ValidationProvider>
                      <KInput
                        v-if="asset.content_asset_kind == 'videoyoutube'"
                        dense
                        rules="required"
                        v-model="asset.url_source"
                        label="Youtube URL"
                        placeholder="Masukkan URL untuk lineup Anda"
                        hide-details="auto"
                      />
                      <KInput
                        v-if="asset.content_asset_kind == 'videostream'"
                        dense
                        rules="required"
                        v-model="asset.url_source"
                        label="URL Stream"
                      />
                    </v-col>
                    <v-col cols="12">
                      <KInput
                        v-model="asset.description"
                        label="Deskirpsi"
                        placeholder="Deskripsi media aset"
                        :multiRow="3"
                      />
                    </v-col>
                  </v-row>
                </v-col>
                <v-divider vertical />
                <v-col cols="3" class="ml-n1 pl-5">
                  <div class="mb-4">Preview Aset</div>
                  <v-alert
                    outlined
                    color="grey"
                    text
                    icon="mdi-folder-multiple-image"
                    v-if="!assetPreview"
                    >Asset masih kosong, silahkan isi aset terlebih
                    dahulu</v-alert
                  >
                  <v-img
                    v-if="assetPreview"
                    max-height="250"
                    cover
                    :src="assetPreview"
                  ></v-img>
                </v-col>
              </v-row>
            </v-form>
          </ValidationObserver>
        </v-card-text>
      </v-card>
    </v-dialog>
  </v-card>
</template>

<script>
import { cloneDeep } from 'lodash'
import { mapActions, mapState } from 'vuex'

export default {
  components: {
    KInput: () => import('~shared/components/k-vue/input'),
  },

  props: {
    assets: {
      type: Array,
      default: () => [],
    },
    types: {
      type: Array,
      default: () => ['video', 'videoyoutube', 'videostream', 'image'],
    },
    hideAddAsset: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {
      id: this.$route.params.id,
      model: null,
      content_assets: [],
      showAssetDialog: false,
      asset: {},
      assetFile: null,
      assetPreview: '',
      assetIndex: -1,
      formAsset: true,
      uploading: false,
      proto: {
        // TODO: get model from backend
        asset: {
          _id: '',
          title: '',
          description: '',
          content_asset_kind: '',
          asset_id: '',
          url_source: '',
          mime: '',
          extras: '',
          is_featured: false,
        },
      },
      asset_type_image: 'image',
      asset_type_youtube: 'videoyoutube',
    }
  },
  computed: {
    ...mapState({
      content: (state) => state.content,
      event: (state) => state.event,
    }),
    assetKinds() {
      let data = this.$store.state.content.assetKinds

      if (this.types.length == 0) {
        if (data.length > 0 && !this.asset.content_asset_kind) {
          this.asset.content_asset_kind = data[0].value
        }
        return data
      }

      let availableKinds = data.filter((x) => {
        return this.types.indexOf(x.value) > -1
      })

      if (availableKinds.length == 1 && !this.asset.content_asset_kind) {
        this.asset.content_asset_kind = availableKinds[0].value
      }

      return availableKinds
    },
  },
  methods: {
    addAsset() {
      this.resetFormAsset()
      this.showAssetDialog = true
    },
    editAsset(item, index) {
      this.assetIndex = index
      this.showAssetDialog = true
      this.asset = item
      this.assetPreview = this.$tool.getAssetURL(item.asset_id)
    },
    resetFormAsset() {
      this.asset = cloneDeep(this.proto.asset)

      this.assetPreview = ''
      this.assetFile = null
      this.assetIndex = -1
      if (this.$refs.formAssetValidator) {
        this.$refs.formAssetValidator.reset()
      }
    },
    closeFormAsset() {
      this.showAssetDialog = false
      this.resetFormAsset()
    },
    uploadAndPreviewFeaturedImage(file) {
      let self = this
      if (file) {
        let filename = 'TMP__' + file.name
        self.uploading = true
        const reader = new FileReader()
        reader.addEventListener(
          'load',
          function () {
            self.$axios
              .post('asset/write', {
                asset: {
                  originalfilename: filename,
                },
                content: btoa(reader.result),
              })
              .then((r) => {
                self.asset.asset_id = r.data._id
                self.asset.mime = r.data.contenttype
              })
              .catch((e) => {
                let errorText = e
                if (e.response && e.response.data) {
                  errorText = e.response.data
                }
                self.asset.asset_id = ''
                self.asset.mime = ''
                self.assetPreview = ''
                self.$dialog.notify.error(errorText)
              })
              .finally(() => {
                self.uploading = false
              })
          },
          false
        )
        reader.readAsBinaryString(file)
        self.assetPreview = URL.createObjectURL(file)
      } else {
        self.assetPreview = ''
      }
    },
    async saveAsset() {
      const valid = await this.$refs.formAssetValidator.validate()
      console.log('>>>', this.asset)
      if (!valid) return

      if (!this.asset._id) this.asset._id = this.$tool.uuid()

      if (!this.asset._id) {
        this.content_assets.push(this.asset)
      } else {
        const index = this.content_assets.findIndex((_) => {
          return _._id == this.asset._id
        })

        if (index == -1) this.content_assets.push(this.asset)
        else this.$set(this.content_assets, index, this.asset)
      }

      if (this.asset.is_featured) {
        // turn off is_featured for all other assets
        const index = this.content_assets.indexOf(this.asset)
        this.content_assets.forEach((as, idx) => {
          if (idx != index) {
            as.is_featured = false
          }
        })
      }

      this.closeFormAsset()
      this.$emit('update-asset', this.content_assets, this.asset)
    },
    removeAsset() {
      this.$delete(this.content_assets, this.assetIndex)
      this.closeFormAsset()
      this.$emit('update-asset', this.content_assets)
    },
    getIdFromUrl(url) {
      // ex: https://www.youtube.com/watch?v=PTA24XJmaN8&t=100s
      let splits1 = url.split('v=')
      if (splits1.length < 2) return ''
      let splits2 = splits1[1].split('&')
      let videoID = splits2[0]
      return videoID
    },
    getYoutubeThumbnail(url) {
      if (!url) {
        return ''
      }
      let videoID = this.getIdFromUrl(url)
      return `https://img.youtube.com/vi/${videoID}/sddefault.jpg`
    },
  },
  watch: {
    assets: {
      deep: true,
      handler(newValue) {
        this.content_assets = cloneDeep(newValue || [])
      },
    },
    asset: {
      deep: true,
      handler(newValue) {
        if (newValue.content_asset_kind == this.asset_type_youtube) {
          this.assetPreview = this.getYoutubeThumbnail(this.asset.url_source)
        }
      },
    },
  },

  created() {
    this.asset = cloneDeep(this.proto.asset)
  },
  mounted() {
    this.content_assets = cloneDeep(this.assets || [])
  },
}
</script>

<style scoped>
.compact {
  margin-top: 0;
}
.compact >>> .v-input__slot {
  margin-bottom: 0;
}
.compact >>> .v-messages {
  display: none;
}
</style>