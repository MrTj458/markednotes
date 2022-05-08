<script setup>
import { useNotificationStore } from "../stores/notification";

const notifications = useNotificationStore();
</script>

<template>
  <main>
    <h1>Marked Notes</h1>
    <button @click="notifications.error('There has been an error.')">
      Notify
    </button>
  </main>
</template>
