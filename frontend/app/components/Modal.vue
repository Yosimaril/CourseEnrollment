<template>
    <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center overflow-y-auto bg-gray-600 bg-opacity-50">
        <div class="relative w-full max-w-md p-6 my-8 overflow-hidden bg-white rounded-lg shadow-xl">
            <div class="flex items-center justify-between">
                <h3 class="text-lg font-bold text-gray-900">
                    <slot name="header"></slot>
                </h3>
                <button @click="closeModal" class="text-gray-500 hover:text-gray-700">
                    <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                    </svg>
                </button>
            </div>
            <div class="mt-4 text-gray-700">
                <slot name="body"></slot>
            </div>
            <div class="mt-5 flex justify-end gap-3">
                <slot name="footer"></slot>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
const props = defineProps<{
    isOpen: boolean;
}>();

const emit = defineEmits(["close"]);

const closeModal = () => {
    emit("close");
};
</script>
