<script setup>
import StandardLayout from "../components/layouts/StandardLayout.vue";
import { useUserStore } from "../stores/user";

const user = useUserStore();
</script>

<template>
  <StandardLayout>
    <div class="center">
      <h1>Marked Notes</h1>
      <p>Write your notes in Markdown and view them as HTML</p>
      <RouterLink v-if="!user.user" :to="{ name: 'register' }" class="cta">
        Register Now
      </RouterLink>
      <RouterLink v-else :to="{ name: 'dashboard' }" class="cta">
        Go to Dashboard
      </RouterLink>
    </div>
  </StandardLayout>
</template>

<style scoped>
.center {
  text-align: center;
}

.cta {
  color: var(--white-color);
  background-color: var(--orange-color);
  border-radius: 4px;
  padding: 0.5rem 1rem;
  text-decoration: none;
  display: inline-block;
  margin-top: 1rem;
}
</style>
