import { defineStore } from "pinia";
import axios from "axios";

export const useUserStore = defineStore({
  id: "user",
  state: () => ({
    user: null,
    error: null,
    loading: false,
  }),
  getters: {},
  actions: {
    async logIn(email, password) {
      try {
        this.loading = true;
        const res = await axios.post("/api/users/login", {
          email,
          password,
        });
        localStorage.setItem("token", res.data.token);
        this.user = res.data.user;
        this.error = null;
        this.loading = false;
      } catch (e) {
        this.user = null;
        this.error = "Invalid email or password.";
        this.loading = false;
      }
    },
  },
});
