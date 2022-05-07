<script setup>
import { toRefs, onMounted, ref } from "vue";
import axios from "axios";
import NoteItem from "./NoteItem.vue";
import TrashIcon from "./icons/TrashIcon.vue";
import NewFolderIcon from "./icons/NewFolderIcon.vue";
import NewFileIcon from "./icons/NewFileIcon.vue";
import NewNoteForm from "../components/NewNoteForm.vue";

const props = defineProps(["folder"]);
const { folder } = toRefs(props);

const newNote = ref(false);
const open = ref(false);
const folders = ref([]);
const notes = ref([]);

onMounted(async () => {
  let res = await axios.get(`/api/folders?parent=${folder.value.id}`);
  folders.value = res.data;
  res = await axios.get(`/api/notes?folder=${folder.value.id}`);
  notes.value = res.data;
});

const addNote = (note) => {
  if (note === null) {
    newNote.value = false;
    return;
  }
  console.log(notes.value);
  notes.value = [...notes.value, note];
  newNote.value = false;
  console.log(notes.value);
};
</script>

<template>
  <div class="container">
    <div class="title">
      <div @click="open = !open" class="name">
        <p class="name-text">{{ open ? "v" : ">" }} {{ folder.name }}</p>
      </div>
      <div class="options">
        <button v-if="open" @click="newNote = true" class="btn">
          <NewFileIcon />
        </button>
        <button v-if="open" class="btn"><NewFolderIcon /></button>
        <button class="btn"><TrashIcon /></button>
      </div>
    </div>
    <ul v-if="open" class="indent">
      <li v-for="folder in folders" :key="folder.id">
        <FolderItem :folder="folder" />
      </li>

      <li v-for="note in notes" :key="note.id">
        <NoteItem :note="note" />
      </li>
      <li v-if="newNote">
        <NewNoteForm :folderId="folder.id" :addNote="addNote" />
      </li>
    </ul>
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
}

.title:hover {
  color: var(--orange-color);
}

.name {
  font-weight: bold;
  flex-grow: 1;
  height: 100%;
  padding: 0.25rem 0;
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
