<script setup>
import axios from "axios";
import { ref, toRefs, onMounted } from "vue";
import TrashIcon from "./icons/TrashIcon.vue";

const props = defineProps(["folderId", "addNote"]);
const { folderId } = toRefs(props);

const inputRef = ref(null);
const name = ref("");

onMounted(() => {
  inputRef.value.focus();
});

const onSubmit = async () => {
  try {
    const res = await axios.post("/api/notes", {
      folder_id: folderId.value,
      name: name.value,
    });
    props.addNote(res.data);
    name.value = "";
  } catch (e) {
    console.error(e);
  }
};

const cancel = () => {
  props.addNote(null);
};
</script>

<template>
  <div class="container">
    <div class="title">
      <div class="name">
        <form @submit.prevent="onSubmit">
          <input
            v-model="name"
            class="name-input"
            placeholder="New Note Name"
            ref="inputRef"
          />
        </form>
      </div>
      <div class="options">
        <button @click="cancel" class="btn"><TrashIcon /></button>
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
}

.name-input {
  padding: 0.25rem 0.5rem;
  background-color: var(--light-gray-color);
  border: 0;
  border-radius: 4px;
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
