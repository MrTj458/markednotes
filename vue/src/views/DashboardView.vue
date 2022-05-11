<script setup>
import FileExplorer from "../components/FileExplorer.vue";
import NoteViewer from "../components/NoteViewer.vue";
import { useNoteStore } from "../stores/note";
import EditIcon from "../components/icons/EditIcon.vue";
import { ref } from "vue";
import ViewIcon from "../components/icons/ViewIcon.vue";
import NoteEditor from "../components/NoteEditor.vue";
import SaveIcon from "../components/icons/SaveIcon.vue";
import { useNotificationStore } from "../stores/notification";
import NavBar from "../components/NavBar.vue";

const noteStore = useNoteStore();
const notify = useNotificationStore();

const viewing = ref(true);

const saveNote = async () => {
  await noteStore.save();
  if (noteStore.error) {
    notify.error("Error saving note.");
    return;
  }

  notify.success("Note saved.");
};
</script>

<template>
  <main class="dash">
    <NavBar class="nav" />
    <FileExplorer class="file-explorer" />
    <div class="main">
      <div class="dash-ribbon">
        <p class="dash-title">{{ noteStore.note?.name }}</p>
        <div class="dash-options">
          <button v-if="!viewing" @click="saveNote" class="dash-btn">
            <SaveIcon />
          </button>
          <button
            v-if="noteStore.note"
            @click="viewing = !viewing"
            class="dash-btn"
          >
            <EditIcon v-if="viewing" /><ViewIcon v-else />
          </button>
        </div>
      </div>
      <div class="viewing-area">
        <NoteViewer v-if="viewing" />
        <NoteEditor v-else />
      </div>
    </div>
  </main>
</template>

<style scoped>
.dash {
  display: grid;
  grid-template-areas:
    "nav nav"
    "side main";
  grid-template-columns: auto 1fr;
  grid-template-rows: auto 1fr;
  height: 100%;
  overflow-y: hidden;
}

.nav {
  grid-area: nav;
}

.file-explorer {
  grid-area: side;
}

.main {
  grid-area: main;
  display: grid;
  grid-template-rows: auto 1fr;
}

.viewing-area {
  position: relative;
}

.dash-title {
  font-size: 2rem;
  font-weight: bold;
}

.dash-ribbon {
  border-bottom: 1px solid var(--light-gray-color);
  padding: 0 1rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.dash-options {
  display: flex;
  gap: 1rem;
}

.dash-btn {
  background: 0;
  border: 0;
  color: var(--white-color);
  cursor: pointer;
}

.dash-btn:hover {
  color: var(--orange-color);
}
</style>
