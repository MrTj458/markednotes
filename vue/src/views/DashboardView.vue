<script setup>
import FileExplorer from "../components/FileExplorer.vue";
import NoteViewer from "../components/NoteViewer.vue";
</script>

<template>
  <main>
    <FileExplorer />
    <NoteViewer />
  </main>
</template>

<style scoped>
main {
  display: flex;
}

.input {
  width: 100%;
}
</style>
