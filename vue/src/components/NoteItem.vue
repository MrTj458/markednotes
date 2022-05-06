<script setup>
import { toRefs } from "vue";
const props = defineProps(["note"]);
const { note } = toRefs(props);
</script>

<template>
  <div class="container">
    <div class="title">
      <div @click="open = !open" class="name">
        <p class="name-text">{{ note.name }}</p>
      </div>
      <div class="options">
        <button class="btn">+</button>
        <button class="btn">+</button>
        <button class="btn">x</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  padding-left: 1rem;
}

.title {
  display: flex;
  justify-content: space-between;
  align-items: center;
  user-select: none;
  cursor: pointer;
  padding: 0.25rem 0;
}

.title:hover {
  color: var(--orange-color);
}

.name {
  font-weight: bold;
  flex-grow: 1;
  height: 100%;
}

.options {
  display: flex;
  justify-content: space-between;
  gap: 1rem;
}

.btn {
  color: var(--white-color);
  background: 0;
  border: 0;
  cursor: pointer;
  font-size: 1rem;
}

.btn:hover {
  color: var(--orange-color);
}
</style>
