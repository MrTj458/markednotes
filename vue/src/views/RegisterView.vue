<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useUserStore } from "../stores/user";
import LoadingSpinner from "../components/LoadingSpinner.vue";

const user = useUserStore();
const router = useRouter();

const email = ref("");
const username = ref("");
const password = ref("");
const password2 = ref("");

const handleSubmit = async () => {
  if (password.value !== password2.value) {
    return;
  }

  await user.register(email.value, username.value, password.value);
  if (user.user !== null) {
    router.push({ name: "dashboard" });
  }
};
</script>

<template>
  <main>
    <h1>Register</h1>
    <form @submit.prevent="handleSubmit">
      <div class="input-group">
        <label for="email">Email</label>
        <input type="email" name="email" v-model="email" required autofocus />
      </div>
      <div class="input-group">
        <label for="username">Username</label>
        <input
          type="text"
          name="username"
          v-model="username"
          required
          autofocus
        />
      </div>
      <div class="input-group">
        <label for="password">Password</label>
        <input type="password" name="password" v-model="password" required />
      </div>
      <div class="input-group">
        <label for="password2">Confirm Password</label>
        <input type="password" name="password2" v-model="password2" required />
      </div>
      <p v-if="user.error" class="error">{{ user.error }}</p>
      <button type="submit" :disabled="user.loading">
        <LoadingSpinner v-if="user.loading" />
        <p v-else>Register</p>
      </button>
      <small
        >Already have an account?
        <RouterLink :to="{ name: 'login' }">Log in Here</RouterLink></small
      >
    </form>
  </main>
</template>

<style>
main {
  margin-top: 2rem;
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
}

form {
  max-width: 95%;
  width: 600px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.error {
  color: red;
  font-weight: bold;
  font-size: 1.2rem;
}

.input-group {
  width: 100%;
}

label {
  font-size: 1.4rem;
}

input {
  margin: 1rem 0;
  display: block;
  width: 100%;
  padding: 0.5rem 1rem;
  font-size: 1.2rem;
  background-color: var(--white-color);
  border: 0;
  border-radius: 4px;
}

button {
  display: block;
  margin-top: 1rem;
  width: 70%;
  border: 2px solid var(--orange-color);
  padding: 1rem;
  color: var(--gray-color);
  background-color: var(--orange-color);
  border-radius: 4px;
  cursor: pointer;
}

small {
  margin-top: 0.5rem;
}
</style>
