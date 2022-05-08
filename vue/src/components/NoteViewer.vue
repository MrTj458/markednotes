<script setup>
import { marked } from "marked";
import { computed } from "@vue/reactivity";
import { useNoteStore } from "../stores/note";

const noteStore = useNoteStore();

const parsedBody = computed(() => {
  if (noteStore.note?.body.length > 0) {
    return marked.parse(noteStore.note.body);
  } else {
    return "";
  }
});
</script>

<template>
  <div class="viewer" v-html="parsedBody"></div>
</template>

<style>
.viewer {
  padding: 1rem 2rem;
}

.viewer ul,
.viewer ol {
  padding-left: 20px;
}
</style>
