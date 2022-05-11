<script setup>
import { useNotificationStore } from "../stores/notification";
import ErrorIcon from "./icons/ErrorIcon.vue";
import CheckIcon from "./icons/CheckIcon.vue";

const notifications = useNotificationStore();
</script>

<template>
  <div class="notifications">
    <div
      v-for="n in notifications.notifications"
      :key="n.text"
      class="notification"
    >
      <ErrorIcon v-if="n.type === 'error'" />
      <CheckIcon v-if="n.type === 'success'" />
      <p>{{ n.text }}</p>
    </div>
  </div>
</template>

<style scoped>
.notifications {
  position: fixed;
  bottom: 1rem;
  right: 1rem;
  display: flex;
  flex-direction: column;
  gap: 1rem;
  z-index: 10;
}

.notification {
  color: (--white-color);
  padding: 1rem 2rem;
  display: flex;
  justify-content: center;
  gap: 1rem;
  background-color: var(--light-gray-color);
  border-radius: 4px;
  border: 2px solid var(--gray-color);
}
</style>
