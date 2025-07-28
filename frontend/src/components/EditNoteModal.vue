<template>
    <NoteModalUI
        heading="Edit Note"
        v-model:title="title"
        v-model:content="content"
        @save="updateNote"
        @close="$emit('close')"
    />
</template>
<script setup lang="ts">
import NoteModalUI from "./NoteModalUI.vue";
import { ref, onMounted } from "vue";

import { GetNote, UpdateNote } from "../../wailsjs/go/main/App";

const props = defineProps({
    noteId: {
        type: String,
        required: true,
    },
});

const emit = defineEmits(["close"]);

const title = ref("");
const content = ref("");

onMounted(async () => {
    const note = await GetNote(props.noteId);
    title.value = note.title;
    content.value = note.content;
});

const updateNote = async () => {
    await UpdateNote(props.noteId, title.value, content.value);
    emit("close");
};
</script>
