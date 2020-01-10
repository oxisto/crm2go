<template>
  <div class="contact-detail" container>
    <b-row v-if="contact">
      <b-col sm>
        <b-card no-body class="mb-2">
          <b-card-body>
            <b-card-title v-if="!edit" v-on:dblclick="enableEdit">{{ contact.name }}</b-card-title>
            <b-form-input v-else v-model="contact.name" placeholder="Enter name" valid></b-form-input>
            <b-card-sub-title
              v-if="!edit"
              v-on:dblclick="enableEdit"
              class="mb-2"
            >{{ contact.company }}</b-card-sub-title>
            <b-form-input
              v-else
              v-model="contact.company"
              class="mb-2 mt-2"
              placeholder="Enter company"
              valid
            ></b-form-input>
            <b-card-text>Job Title</b-card-text>
          </b-card-body>
          <b-list-group flush>
            <b-list-group-item>
              <div v-if="!edit" v-on:dblclick="enableEdit">
                E-Mail
                <div>{{ contact.email }}</div>
              </div>
              <b-form-group
                v-else
                id="fieldset-horizontal"
                label-cols-sm="4"
                label-cols-lg="3"
                description="The primary E-Mail for this contact."
                label="E-Mail"
                label-for="input-email"
              >
                <b-form-input v-model="contact.email" id="input-email" type="email"></b-form-input>
              </b-form-group>
            </b-list-group-item>
          </b-list-group>
          <b-card-body v-if="edit">
            <b-button v-on:click="saveContact" type="Submit">
              <span v-if="contactId=='new'">Add</span>
              <span v-else>Save</span>
            </b-button>
          </b-card-body>
        </b-card>
      </b-col>
      <b-col cols="8">
        <span>{{ notes }}</span>
        <span>{{ note }}</span>
        <b-card no-body>
          <b-list-group flush>
            <b-list-group-item :key="note.id" v-for="note in notes">{{ note.text }}</b-list-group-item>
            <b-list-group-item v-on:click="enableNewNote" v-if="!newNote" class="big-plus">
              <font-awesome-icon icon="plus-square" />
            </b-list-group-item>
            <b-form-group v-else>
              <b-form-textarea
                id="textarea"
                v-model="note.text"
                placeholder="Enter something..."
                class="border-0"
                rows="3"
                max-rows="6"
              ></b-form-textarea>
              <b-button class="mt-2 ml-2" v-on:click="addNewNote">Add</b-button>
            </b-form-group>
          </b-list-group>
        </b-card>
      </b-col>
    </b-row>
    {{ contact }}
  </div>
</template>

<style scoped>
.big-plus {
  font-size: 32px;
  text-align: center;
}
</style>

<script lang="ts">
import { Component, Vue, Prop, Watch } from "vue-property-decorator";

@Component
export default class ContactDetail extends Vue {
  private contact: any = null;
  @Prop() private contactId!: number | "new";
  private notes: any[] = [];
  private edit: boolean = false;
  private newNote: boolean = false;
  private note: any = {};

  @Watch("contactId", { immediate: true })
  async changeContactId(contactId: number | "new") {
    if (contactId == "new") {
      // new contact
      this.contact = {};
      this.edit = true;
    } else {
      this.contact = await this.$http.get("/api/contacts/" + contactId);
      this.notes = await this.$http.get(
        "/api/contacts/" + contactId + "/notes/"
      );
    }
  }

  enableEdit() {
    this.edit = true;
  }

  enableNewNote() {
    this.newNote = true;
  }

  async addNewNote() {
    // skip empty notes
    if (this.note.text === "" || this.note.text === undefined) {
      return;
    }

    // set contact id
    this.note.contactId = this.contact.id;

    let note = await this.$http.post(
      "/api/contacts/" + this.contactId + "/notes/",
      this.note
    );
    this.notes.push(note);
  }

  async saveContact() {
    if (this.contactId == "new") {
      this.contact = await this.$http.post("/api/contacts", this.contact);
    } else {
      this.contact = await this.$http.put(
        "/api/contacts/" + this.contact.id,
        this.contact
      );
    }

    // turn off edit
    this.contactId = this.contact.id;
    this.edit = false;
  }
}
</script>
