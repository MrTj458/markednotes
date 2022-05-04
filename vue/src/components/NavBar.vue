<script setup>
import { useRouter } from "vue-router";
import { useUserStore } from "../stores/user";

const router = useRouter();
const user = useUserStore();

const signOut = async () => {
  await router.push({ name: "home" });
  localStorage.removeItem("token");
  window.location.reload();
};
</script>

<template>
  <nav class="nav-bar">
    <ul class="nav-list">
      <li>
        <RouterLink :to="{ name: 'home' }" class="nav-item"
          >Marked Notes</RouterLink
        >
      </li>
    </ul>
    <ul v-if="user.user" class="nav-list">
      <li>
        <button @click="signOut" class="nav-item">Sign Out</button>
      </li>
    </ul>
    <ul v-else class="nav-list">
      <li>
        <RouterLink :to="{ name: 'login' }" class="nav-item">Login</RouterLink>
      </li>
      <li>
        <RouterLink :to="{ name: 'register' }" class="nav-item">
          Register
        </RouterLink>
      </li>
    </ul>
  </nav>
</template>

<style>
.nav-bar {
  background-color: var(--orange-color);
  display: flex;
  align-items: center;
  width: 100%;
  justify-content: space-between;
  padding: 1rem 2rem;
}

.nav-list {
  list-style: none;
  display: flex;
  justify-content: center;
  gap: 2em;
}

.nav-item {
  font-size: 1.2em;
  text-decoration: none;
  color: var(--white-color);
  border-radius: 4px;
  padding: 0.25rem 0.5rem;
  background: none;
  border: none;
  cursor: pointer;
  display: block;
  margin: 0;
  width: auto;
}
</style>
