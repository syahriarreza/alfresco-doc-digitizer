<template>
  <base-section id="about-our-product">
    <base-section-heading color="grey lighten-2" icon="mdi-vuetify" title="About our Product" />
    <v-container>
      <v-row>
        <v-col v-for="card in cards" :key="card.title" cols="12" md="4">
          <base-info v-bind="card"></base-info>
        </v-col>
      </v-row>
    </v-container>
  </base-section>
</template>

<script>
  export default {
    name: 'SectionAboutOurProduct',

    components: {
      BaseSection: () => import('../base/Section'),
      BaseSectionHeading:()=>import('../base/SectionHeading'),
      BaseInfo:()=>import('../base/Info'),
    },


    data: () => ({
      cards: [
        {
          title: 'Grow your Revenue',
          subtitle: 'Best Productivity',
          text: 'Similique sunt in culpa qui officia deserunt mollitia animi, id est laborut dolorum fuga.harum quidem rerum facilis estexpedita distinctio.',
          callout: '01',
        },
        {
          title: 'Affordable Prices',
          subtitle: 'Special Offers',
          text: 'Similique sunt in culpa qui officia deserunt mollitia animi, id est laborut dolorum fuga.harum quidem rerum facilis estexpedita distinctio.',
          callout: '02',
        },
        {
          title: 'Target Setting',
          subtitle: 'Income Flow',
          text: 'Similique sunt in culpa qui officia deserunt mollitia animi, id est laborut dolorum fuga.harum quidem rerum facilis estexpedita distinctio.',
          callout: '03',
        },
      ],
    }),
  }
</script>
