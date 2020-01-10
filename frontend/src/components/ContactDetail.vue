<template>
  <div class="contact-detail" container>
    <b-row>
      <b-col sm>
        <b-card no-body class="mb-2">
          <b-card-body>
            <b-card-title>{{ contact.name }}</b-card-title>
            <b-card-sub-title class="mb-2">
              {{ contact.company }}
            </b-card-sub-title>
            <b-card-text>Job Title</b-card-text>
          </b-card-body>
          <b-list-group flush>
            <b-list-group-item>
              this
              <div>that</div>
            </b-list-group-item>
          </b-list-group>
        </b-card>
      </b-col>
      <b-col cols="10">Interactions / Stuff</b-col>
    </b-row>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop, Watch } from "vue-property-decorator";

@Component
export default class ContactDetail extends Vue {
  private contact: any = null;
  @Prop() private contactId!: number;

  @Watch("contactId", { immediate: true })
  async changeContactId(contactId: number) {
    this.contact = await this.$http.get("/api/contacts/" + contactId);

    const numSlots = this.contact.inventory.slots.length;

    for (let i = numSlots; i < 16; i++) {
      this.contact.inventory.slots[i] = { typeId: -1, amount: 0 };
    }
  }
}
</script>
