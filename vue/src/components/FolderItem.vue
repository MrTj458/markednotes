<script setup>
import { toRefs, onMounted, ref } from "vue";
import axios from "axios";
import NoteItem from "./NoteItem.vue";
import TrashIcon from "./icons/TrashIcon.vue";
import NewFolderIcon from "./icons/NewFolderIcon.vue";
import NewFileIcon from "./icons/NewFileIcon.vue";
import NewNoteForm from "../components/NewNoteForm.vue";
import OpenIcon from "./icons/OpenIcon.vue";
import ClosedIcon from "./icons/ClosedIcon.vue";
import NewFolderForm from "../components/NewFolderForm.vue";
import { useNotificationStore } from "../stores/notification";
import { useNoteStore } from "../stores/note";
import LoadingSpinner from "./LoadingSpinner.vue";

const props = defineProps(["folder", "removeChild"]);
const { folder } = toRefs(props);

const notify = useNotificationStore();
const noteStore = useNoteStore();

const loading = ref(false);
const newNote = ref(false);
const newFolder = ref(false);
const open = ref(false);
const folders = ref([]);
const notes = ref([]);

onMounted(async () => {
  loading.value = true;
  try {
    let res = await axios.get(`/api/folders?parent=${folder.value.id}`);
    folders.value = res.data;
    res = await axios.get(`/api/notes?folder=${folder.value.id}`);
    notes.value = res.data;
    loading.value = false;
  } catch (e) {
    notify.error("Error loading files");
    loading.value = false;
  }
});

const addNote = (note) => {
  if (note === null) {
    newNote.value = false;
    return;
  }
  notes.value = [...notes.value, note];
  newNote.value = false;
};

const addFolder = (folder) => {
  if (folder === null) {
    newFolder.value = false;
    return;
  }
  folders.value = [...folders.value, folder];
  newFolder.value = false;
};

const removeChild = (id) => {
  folders.value = folders.value.filter((f) => f.id !== id);
};

const deleteFolder = async () => {
  try {
    loading.value = true;
    await axios.delete(`/api/folders/${folder.value.id}`);
    props.removeChild(folder.value.id);
    noteStore.note = null;
    notify.success("Folder deleted.");
  } catch (e) {
    console.error(e);
    loading.value = false;
    notify.error("Error deleting folder.");
  }
};

const deleteNote = (id) => {
  notes.value = notes.value.filter((note) => note.id !== id);
};
</script>

<template>
  <div v-if="!loading" class="container">
    <div class="title">
      <div @click="open = !open" class="name">
        <p :class="{ 'name-text': true, root: folder.id === '' }">
          <OpenIcon v-if="open" />
          <ClosedIcon v-else />
          {{ folder.name }}
        </p>
      </div>

      <div class="options">
        <button
          @click="
            newNote = true;
            open = true;
          "
          class="btn"
        >
          <NewFileIcon />
        </button>
        <button
          @click="
            newFolder = true;
            open = true;
          "
          class="btn"
        >
          <NewFolderIcon />
        </button>
        <button v-if="folder.id !== ''" @click="deleteFolder" class="btn">
          <TrashIcon />
        </button>
      </div>
    </div>

    <ul v-if="open && !loading">
      <li v-for="folder in folders" :key="folder.id">
        <FolderItem :folder="folder" :removeChild="removeChild" />
      </li>

      <li v-for="note in notes" :key="note.id">
        <NoteItem :note="note" :deleteNote="deleteNote" />
      </li>

      <li v-if="newNote">
        <NewNoteForm :folderId="folder.id" :addNote="addNote" />
      </li>

      <li v-if="newFolder">
        <NewFolderForm :folderId="folder.id" :addFolder="addFolder" />
      </li>
    </ul>
  </div>
  <LoadingSpinner v-else :white="true" class="container" />
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

.root {
  font-size: 1.4rem;
}

.title:hover {
  color: var(--orange-color);
}

.name {
  font-weight: bold;
  flex-grow: 1;
  height: 100%;
  padding: 0.25rem 0;
  padding-right: 2rem;
}

.name-text {
  display: flex;
  align-items: center;
}

.options {
  display: flex;
  justify-content: space-between;
  gap: 1rem;
  margin-left: 1rem;
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

ul {
  list-style-type: none;
}
</style>
