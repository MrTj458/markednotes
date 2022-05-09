<script setup>
import { useUserStore } from "../stores/user";
import FolderItem from "./FolderItem.vue";

const user = useUserStore();

const rootFolder = {
  name: user.user.username,
  id: "",
};
</script>

<template>
  <div class="explorer">
    <FolderItem :folder="rootFolder" />
  </div>
</template>

<style scoped>
.explorer {
  border-right: 1px solid var(--light-gray-color);
  min-width: 300px;
  height: 100%;
  gap: 2rem;
}

ul {
  list-style-type: none;
}
</style>
