import { defineStore } from "pinia";
import axios from "axios";

export const useNoteStore = defineStore({
  id: "note",
  state: () => ({
    note: null,
    error: null,
  }),
  getters: {},
  actions: {
    async save() {
      try {
        const res = await axios.put(`/api/notes/${this.note.id}`, this.note);
        this.note = res.data;
        this.error = null;
      } catch (e) {
        this.note = null;
        this.error = "Error saving note.";
      }
    },
  },
});
