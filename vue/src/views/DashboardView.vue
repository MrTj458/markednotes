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
    <FileExplorer />
    <div class="dash-right">
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
      <NoteViewer v-if="viewing" />
      <NoteEditor v-else />
    </div>
  </main>
</template>

<style scoped>
.dash {
  display: flex;
}

.dash-right {
  width: 100%;
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
