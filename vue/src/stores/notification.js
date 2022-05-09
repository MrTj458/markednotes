import { defineStore } from "pinia";
import { v4 as uuid } from "uuid";

export const useNotificationStore = defineStore({
  id: "notification",
  state: () => ({
    notifications: [],
  }),
  getters: {},
  actions: {
    error(text) {
      const id = uuid();
      this.notifications.push({ id, type: "error", text });
      setTimeout(() => {
        this.notifications = this.notifications.filter((n) => n.id != id);
      }, 5000);
    },
    success(text) {
      const id = uuid();
      this.notifications.push({ id, type: "success", text });
      setTimeout(() => {
        this.notifications = this.notifications.filter((n) => n.id != id);
      }, 5000);
    },
  },
});
