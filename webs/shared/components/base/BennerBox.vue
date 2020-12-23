<template>
  <div :class="classes" class="pt-2 gt-col">
      <div class="gt-inner">
        <div class="gt-logo">
          <base-img contain aspect-ratio="1.7" :src="require(`@/static/carlogo/${item.img}.png`)" color="grey lighten-1" width="100" height="100" :class="{ full: fullWidthImage }"  tile />
        </div>
        <div class="primary"></div>
      </div>
  </div>
</template>

<script>
  // Mixins
  import Heading from '@/mixins/heading'

  export default {
    name: 'BaseBannerBox',

    components: {
      BaseImg: () => import('./Img'),
    },

    mixins: [Heading],

    props: {
      item :  Object,
    },
    data: () => ({
      'fullWidthImage': false
    }),
    computed: {
      classes () {
        const classes = [
          `mb-${this.space}`,
        ]
        classes.push('text-center')
        return classes
      },
    },
  }
</script>

<style scoped>
.gt-col {
    width: 100%;
    position: relative;
    display: inline-block;
}
.gt-col>.gt-inner {
    display: inline-block;
    height: 100%;
    width: 100%;
}
.gt-logo {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-wrap: wrap;
    text-align: center;
    padding: 10px;
    box-shadow: 0px -4px 14px 0px rgba(215, 215, 215, 0.349019607843137);
    -webkit-box-shadow: 0px -4px 14px 0px rgba(215, 215, 215, 0.349019607843137);
    -moz-box-shadow: 0px -4px 14px 0px rgba(215, 215, 215, 0.349019607843137);
    height: 150px;
    width: 100%;
    border-radius: 5px 5px 0px 0px;
    overflow: hidden;
    background-color: #f9f9f9;
}
.primary{
    width: 100%;
    height: 5px;
    border-radius: 0px 0px 5px 5px;
}
.v-image.v-responsive{
  flex: none !important;
}
</style>
<style>
.gt-logo .v-responsive__sizer {
    padding-bottom: 0% !important;
}

</style>