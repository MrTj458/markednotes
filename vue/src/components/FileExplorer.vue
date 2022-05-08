<script setup>
import axios from "axios";
import { onMounted, ref } from "vue";
import FolderItem from "./FolderItem.vue";
import NoteItem from "./NoteItem.vue";

const folders = ref([]);
const notes = ref([]);

onMounted(async () => {
  let res = await axios.get("/api/folders");
  folders.value = res.data;
  res = await axios.get("/api/notes");
  notes.value = res.data;
});
</script>

<template>
  <div class="explorer">
    <h2>File Explorer</h2>
    <ul>
      <li v-for="folder in folders" :key="folder.id">
        <FolderItem :folder="folder" />
      </li>
      <li v-for="note in notes" :key="note.id"><NoteItem :note="note" /></li>
    </ul>
  </div>
</template>

<style scoped>
.explorer {
  border-right: 1px solid var(--light-gray-color);
  min-width: 300px;
  height: 100%;
  padding: 1rem;
  gap: 2rem;
}

ul {
  list-style-type: none;
}
</style>
