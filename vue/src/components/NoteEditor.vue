<script setup>
import { onMounted, ref } from "vue";
import { useNoteStore } from "../stores/note";

const noteStore = useNoteStore();

const textArea = ref(null);

onMounted(() => {
  textArea.value.focus();
});
</script>

<template>
  <textarea
    v-model="noteStore.note.body"
    ref="textArea"
    class="editor"
  ></textarea>
</template>

<style scoped>
.editor {
  width: 100%;
  height: 100%;
  color: var(--white-color);
  background-color: var(--light-gray-color);
  font-size: 1rem;
  padding: 1rem;
  resize: none;
  border: none;
  outline: none;
  margin: 0;
}
</style>
