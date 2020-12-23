<template>
  <v-container>
    <base-section id="news-grid">
      <v-row>
        <v-col cols="4" v-for="(n, i) in displayedNews" :key="i">
          <base-news :title="n.title" :text="n.text" :image="n.image" :url="n.url"></base-news>
        </v-col>
      </v-row>
      <v-row>
        <v-col class="mx-auto" style="text-align: center;">
          <a href="#" @click.prevent="readMore" v-show="paramNews.length > 0">Read More...</a>
        </v-col>
      </v-row>
    </base-section>
  </v-container>
</template>

<script>
function displayNews(_this) {
  for (let i = 0; i < _this.perPage; i++) {
    if (_this.paramNews.length > 0) {
      let n = _this.paramNews.shift();
      _this.displayedNews.push(n);
    }
  }
}

export default {
  components: {
    BaseSection: () => import("../base/Section"),
    BaseNews: () => import("../base/News")
  },
  props: {
    perPage: {
      type: Number,
      default: 3
    },
    news: {
      type: Array,
      default: []
    }
  },
  data() {
    return {
      proto: {
        title: "No Title",
        text: "No Content.", //TODO: implement helper to substring text to just only few sentences
        date: new Date(),
        image: "#",
        url: "#"
      },
      model: null,
      paramNews: [],
      displayedNews: []
    };
  },
  methods: {
    readMore() {
      displayNews(this);
    }
  },
  created() {
    this.paramNews = JSON.parse(JSON.stringify(this.news));
    displayNews(this);
  }
};
</script>

<style lang="scss" scoped>
</style>