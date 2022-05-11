<script setup>
import { useNotificationStore } from "../stores/notification";
import StandardLayout from "../components/layouts/StandardLayout.vue";

const notifications = useNotificationStore();
</script>

<template>
  <StandardLayout>
    <h1>Marked Notes</h1>
    <button @click="notifications.success('This is a notification!')">
      Notify
    </button>
  </StandardLayout>
</template>
