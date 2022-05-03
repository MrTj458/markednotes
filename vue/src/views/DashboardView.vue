<template>
  <h1>Dashboard</h1>
  <p>{{ user.user }}</p>
</template>

<script setup>
import { useUserStore } from "../stores/user";

const user = useUserStore();
</script>

<style></style>
