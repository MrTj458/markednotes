<script setup>
import TrashIcon from "./icons/TrashIcon.vue";
import { toRefs } from "vue";
import axios from "axios";
import { useNotificationStore } from "../stores/notification";
import { useNoteStore } from "../stores/note";

const props = defineProps(["note", "deleteNote"]);
const { note } = toRefs(props);

const notify = useNotificationStore();
const noteStore = useNoteStore();

const deleteNote = async () => {
  try {
    await axios.delete(`/api/notes/${note.value.id}`);
    props.deleteNote(note.value.id);
    notify.success("Note deleted.");
  } catch (e) {
    console.error(e);
    notify.error("Error deleting note.");
  }
};
</script>

<template>
  <div class="container">
    <div class="title">
      <div class="name" @click="noteStore.note = note">
        <p class="name-text">{{ note.name }}</p>
      </div>
      <div class="options">
        <button @click="deleteNote" class="btn"><TrashIcon /></button>
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
  margin-right: 1rem;
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
