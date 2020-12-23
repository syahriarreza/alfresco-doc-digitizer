<template>
  <base-section id="news-slider">
    <v-slide-group v-model="model" class="pa-4" active-class="success" show-arrows>
      <v-slide-item v-for="(n, i) in news" :key="i">
        <v-card class="ma-4">
          <v-row class="fill-height" align="center" justify="center">
            <v-scale-transition>
              <base-news :title="n.title" :text="n.text" :image="n.image" :url="n.url"></base-news>
            </v-scale-transition>
          </v-row>
        </v-card>
      </v-slide-item>
    </v-slide-group>
  </base-section>
</template>

<script>
export default {
  components: {
    BaseSection:()=>import('../base/Section'),
    BaseNews:()=>import('../base/News'),
  },
  props: {
    news: {
      type: Array,
      default: [
        {
          title: 'No Title',
          text: 'No Content.', //TODO: implement helper to substring text to just only few sentences
          date: new Date(),
          image: '#',
          url: '#',  
        }
      ]
    },
  },
  data() {
    return {
      model: null,
    }
  },
}
</script>

<style lang="scss" scoped>
</style>