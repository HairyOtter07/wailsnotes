<template>
    <div class="flex flex-col w-full h-full items-center">
        <div class="flex flex-col w-full h-full p-12 items-center">
            <div class="flex flex-col w-full max-w-lg p-4 space-y-4">
                <div
                    class="flex flex-col w-full items-center space-y-4"
                    v-if="notes.length === 0"
                >
                    <h1 class="text-4xl text-gray-800">No Notes</h1>
                    <p class="text-lg text-gray-600">
                        Create a new note to get started.
                    </p>
                </div>
                <NoteCard
                    v-for="note in notes"
                    :id="note.id"
                    :title="note.title"
                    :content="note.content"
                    @edit="toggleEditModal"
                    @delete="deleteNote"
                />
            </div>
        </div>
    </div>
    <div
        class="fixed flex items-center justify-center bottom-4 right-4 bg-gray-200 rounded-full shadow-lg p-4 w-16 h-16 z-10 cursor-pointer"
        @click="toggleNewModal"
    >
        <Icon icon="mdi:plus" width="24" height="24" class="text-gray-600" />
    </div>
    <NewNoteModal v-if="isNewModalActive" @close="toggleNewModal" />
    <EditNoteModal
        v-if="isEditModalActive"
        :note-id="selectedNoteId"
        @close="() => toggleEditModal('')"
    />
</template>
<script lang="ts" setup>
import NoteCard from "./components/NoteCard.vue";
import NewNoteModal from "./components/NewNoteModal.vue";
import EditNoteModal from "./components/EditNoteModal.vue";

import { GetNotes, DeleteNote } from "../wailsjs/go/main/App";
import { main } from "../wailsjs/go/models";
import { EventsOn } from "../wailsjs/runtime/runtime";

import { Icon } from "@iconify/vue";
import { ref, onMounted } from "vue";

const isNewModalActive = ref(false);
const isEditModalActive = ref(false);

const notes = ref([] as main.Note[]);
const selectedNoteId = ref("");

const fetchNotes = async () => {
    notes.value = await GetNotes();
};

onMounted(() => {
    fetchNotes();
    EventsOn("updateNotes", fetchNotes);
});

const toggleNewModal = () => {
    isNewModalActive.value = !isNewModalActive.value;
};

const toggleEditModal = (id: string) => {
    selectedNoteId.value = id;
    isEditModalActive.value = !isEditModalActive.value;
};

const deleteNote = async (id: string) => {
    await DeleteNote(id);
};
</script>
