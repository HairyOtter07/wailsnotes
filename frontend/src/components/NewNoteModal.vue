<template>
    <NoteModalUI
        heading="New Note"
        v-model:title="title"
        v-model:content="content"
        @save="createNote"
        @close="$emit('close')"
    />
</template>
<script setup lang="ts">
import NoteModalUI from "./NoteModalUI.vue";
import { ref } from "vue";

import { CreateNote } from "../../wailsjs/go/main/App";

const emit = defineEmits(["close"]);

const title = ref("");
const content = ref("");

const createNote = async () => {
    await CreateNote(title.value, content.value);
    emit("close");
};
</script>
