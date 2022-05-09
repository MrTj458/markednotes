import { defineStore } from "pinia";

export const useNoteStore = defineStore({
  id: "note",
  state: () => ({
    note: null,
  }),
  getters: {},
  actions: {},
});
