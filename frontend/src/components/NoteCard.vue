<template>
    <div class="flex flex-col p-4 items-start space-y-2 bg-gray-200 rounded-xl">
        <div class="flex w-full justify-between items-center">
            <h1 class="text-4xl font-bold">{{ title }}</h1>
            <div class="flex items-center">
                <Icon
                    icon="mdi:delete"
                    width="24"
                    height="24"
                    class="cursor-pointer text-gray-600"
                    @click="$emit('delete', id)"
                />
                <Icon
                    icon="mdi:edit"
                    width="24"
                    height="24"
                    class="cursor-pointer text-gray-600"
                    @click="$emit('edit', id)"
                />
            </div>
        </div>
        <p class="text-lg text-gray-600">
            {{ content ? content : "No content" }}
        </p>
    </div>
</template>
<script setup lang="ts">
import { Icon } from "@iconify/vue";

defineProps({
    id: { type: String, required: true },
    title: { type: String, required: true },
    content: { type: String, required: true },
});
defineEmits(["delete", "edit"]);
</script>
