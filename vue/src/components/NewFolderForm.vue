<script setup>
import axios from "axios";
import { ref, toRefs, onMounted } from "vue";
import { useNotificationStore } from "../stores/notification";
import TrashIcon from "./icons/TrashIcon.vue";
import LoadingSpinner from "./LoadingSpinner.vue";

const props = defineProps(["folderId", "addFolder"]);
const { folderId } = toRefs(props);

const notify = useNotificationStore();

const loading = ref(false);
const inputRef = ref(null);
const name = ref("");

onMounted(() => {
  inputRef.value.focus();
});

const onSubmit = async () => {
  try {
    if (name.value.length < 1) {
      notify.error("Folder name must be at least 1 character.");
      return;
    } else if (name.value.length > 30) {
      notify.error("Folder name cannot be greater than 30 characters.");
      return;
    }

    loading.value = true;
    const res = await axios.post("/api/folders", {
      parent_id: folderId.value ? folderId.value : null,
      name: name.value,
    });
    props.addFolder(res.data);
    loading.value = false;
    notify.success("Folder created.");
    name.value = "";
  } catch (e) {
    console.error(e);
    loading.value = false;
    notify.error("Error creating new folder.");
  }
};

const cancel = () => {
  props.addFolder(null);
};
</script>

<template>
  <div v-if="!loading" class="container">
    <div class="title">
      <div class="name">
        <form @submit.prevent="onSubmit">
          <input
            v-model="name"
            class="name-input"
            placeholder="New Folder Name"
            ref="inputRef"
          />
        </form>
      </div>
      <div class="options">
        <button @click="cancel" class="btn"><TrashIcon /></button>
      </div>
    </div>
  </div>
  <LoadingSpinner v-else :white="true" />
</template>

<style scoped>
.container {
  margin-left: 1.5rem;
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
}

.name-input {
  padding: 0.25rem 0.5rem;
  background-color: var(--light-gray-color);
  border: 0;
  border-radius: 4px;
  color: var(--white-color);
  outline: none;
}

.name-input::placeholder {
  color: var(--white-color);
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
