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
  <div class="note-viewer" v-html="parsedBody"></div>
</template>

<style>
.note-viewer {
  padding: 1rem;
  width: 100%;
  height: 100%;
  overflow-y: auto;
  position: absolute;
}

.note-viewer ul,
.note-viewer ol {
  padding-left: 20px;
}

.note-viewer code {
  display: inline-block;
  background-color: var(--light-gray-color);
  padding: 0.25rem;
  border-radius: 4px;
}
</style>
