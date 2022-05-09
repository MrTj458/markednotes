import axios from "axios";
import { createRouter, createWebHashHistory } from "vue-router";
import { useUserStore } from "../stores/user";
import HomeView from "../views/HomeView.vue";

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/login",
      name: "login",
      component: () => import("../views/LoginView.vue"),
    },
    {
      path: "/register",
      name: "register",
      component: () => import("../views/RegisterView.vue"),
    },
    {
      path: "/dashboard",
      name: "dashboard",
      component: () => import("../views/DashboardView.vue"),
      meta: {
        requiresAuth: true,
      },
    },
  ],
});

router.beforeEach(async (to) => {
  const user = useUserStore();
  if (!user.user) {
    const token = localStorage.getItem("token");
    if (token) {
      try {
        const res = await axios.get("/api/users/me");
        user.user = res.data;
        user.error = null;
      } catch (e) {
        localStorage.removeItem("token");
        user.user = null;
        user.error = null;
      }
    }

    if (to.meta.requiresAuth && !user.user) {
      return { name: "login" };
    }
  }
});

export default router;
