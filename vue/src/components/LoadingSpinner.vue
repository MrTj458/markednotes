<script setup>
import { defineProps } from "vue";

const props = defineProps(["white"]);
</script>

<template>
  <div :class="{ spinner: true, white: props.white }"></div>
</template>

<style scoped>
.spinner {
  width: 25px;
  height: 25px;
  border: 4px solid var(--gray-color);
  border-right-color: transparent;
  border-radius: 50%;
  animation: spin 0.5s linear infinite;
  margin: auto;
}

.white {
  border-color: var(--white-color);
  border-right-color: transparent;
}

@keyframes spin {
  0% {
    transform: rotate(0);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>
