<script setup>
import { useNotificationStore } from "../stores/notification";
import ErrorIcon from "./icons/ErrorIcon.vue";
import CheckIcon from "./icons/CheckIcon.vue";

const notifications = useNotificationStore();
</script>

<template>
  <div class="screen">
    <div class="container">
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
    </div>
  </div>
</template>

<style scoped>
.screen {
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  z-index: -1;
}

.container {
  position: relative;
  width: 100%;
  height: 100%;
}

.notifications {
  position: absolute;
  bottom: 1rem;
  right: 1rem;
}

.notification {
  color: (--white-color);
  padding: 1rem 2rem;
  z-index: 10;
  display: flex;
  justify-content: center;
  gap: 1rem;
  background-color: var(--light-gray-color);
  border-radius: 4px;
}
</style>
