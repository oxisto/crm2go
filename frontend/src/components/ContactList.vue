<template>
  <div class="contact-list">
    <b-carousel
      :interval="4000"
      controls
      indicators
      background="#ababab"
      img-width="1024"
      img-height="480"
    >
      <router-link
        :to="{ name: 'contacts-with-id', params: { id: contact.id } }"
        :key="contact.id"
        v-for="contact in contacts"
      >
        <b-carousel-slide :caption="contact.name" img-blank>
          <p>{{ contact }}</p>
        </b-carousel-slide>
      </router-link>
    </b-carousel>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class ContactList extends Vue {
  private contacts: any[] = [];

  async mounted() {
    this.contacts = await this.$http.get("/api/contacts/");
  }
}
</script>
