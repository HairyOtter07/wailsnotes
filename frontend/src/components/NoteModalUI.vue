<template>
    <div
        class="fixed top-0 left-0 w-screen h-screen z-50 bg-black/50 flex justify-center items-center"
    >
        <div class="flex flex-col bg-white p-4 rounded-lg shadow-md space-y-4">
            <div class="flex w-full justify-between items-center">
                <h1 class="text-2xl font-bold">
                    {{ heading }}
                </h1>
                <div class="flex space-x-2">
                    <Icon
                        icon="mdi:close"
                        class="text-gray-600 cursor-pointer"
                        @click="$emit('close')"
                    />
                    <Icon
                        icon="mdi:check"
                        class="text-gray-600 cursor-pointer"
                        @click="validateAndSave"
                    />
                </div>
            </div>
            <div class="flex flex-col space-y-2">
                <div class="flex flex-col space-y-px">
                    <p class="text-sm text-red-500" v-if="!isTitleValid">
                        Please enter a valid title.
                    </p>
                    <input
                        v-model="title"
                        type="text"
                        placeholder="Title"
                        class="border rounded-md px-2 py-1"
                        :class="
                            isTitleValid ? 'border-gray-300' : 'border-red-500'
                        "
                    />
                </div>
                <textarea
                    v-model="content"
                    placeholder="Content"
                    class="border rounded-md px-2 py-1 border-gray-300"
                />
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import { is } from "@babel/types";
import { Icon } from "@iconify/vue";
import { ref } from "vue";

defineProps({
    heading: String,
});

const emit = defineEmits(["close", "save"]);

const title = defineModel<string>("title");
const content = defineModel<string>("content");

const isTitleValid = ref(true);

const validateAndSave = () => {
    if (title.value) {
        isTitleValid.value = true;
        emit("save");
    } else {
        isTitleValid.value = false;
    }
};
</script>
