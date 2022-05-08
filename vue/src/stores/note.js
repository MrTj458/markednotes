import { defineStore } from "pinia";

export const useNoteStore = defineStore({
  id: "user",
  state: () => ({
    note: null,
  }),
  getters: {},
  actions: {},
});
