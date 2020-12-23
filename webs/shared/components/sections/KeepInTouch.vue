<template>
  <base-section id="info-alt" space="56">
    <v-container>
      <v-row justify="space-between">
        <v-col cols="12" md="5">
          <widget-business-info :title="title">{{ info }}</widget-business-info>
        </v-col>

        <v-col cols="12" md="6">
          <widget-contact-form :title="contactFormTitle" />
        </v-col>
      </v-row>
    </v-container>
  </base-section>
</template>

<script>
export default {
  name: "SectionKeepInTouch",

  components: {
    WidgetBusinessInfo: () => import("../widget/BusinessInfo"),
    WidgetContactForm: () => import("../widget/ContactForm"),
    BaseSection: () => import("../base/Section")
  },

  props: {
    title: {
      type: String,
      default: "KEEP IN TOUCH WITH US"
    },
    info: {
      type: String,
      default:
        "Lorem ipsum dolor, sit amet consectetur adipisicing elit. Illo pariatur quidem, quis modi iure nihil commodi suscipit, corporis dicta adipisci, voluptatum quia cupiditate."
    },
    contactFormTitle: {
      type: String,
      default: "MAIL US YOUR MESSAGE"
    }
  }
};
</script>
