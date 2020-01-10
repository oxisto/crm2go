<template>
  <div class="contacts">
    <ContactDetail :contactId="contactId" v-if="contactId" class="p-2" />
    <ContactList v-else />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import ContactList from "@/components/ContactList.vue";
import ContactDetail from "@/components/ContactDetail.vue";
import { Route } from "vue-router";
import { Component, Watch } from "vue-property-decorator";

@Component({
  components: { ContactList, ContactDetail }
})
export default class Characters extends Vue {
  private contactId: number | "new" | null = null;

  mounted() {}

  @Watch("$route", { immediate: true, deep: true })
  routeChanged(route: Route) {
    const id = route.params["id"];
    this.contactId =
      id == undefined ? null : id == "new" ? "new" : parseInt(id);
  }
}
</script>
