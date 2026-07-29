<script setup lang="ts">
defineEmits<{
    home: [];
}>();
</script>

<template>
    <main class="flex min-h-screen items-center justify-center bg-slate-50 px-4 py-6">
        <section class="w-full max-w-md rounded-md bg-white p-8 text-center shadow-md sm:p-10">
            <p class="text-sm font-bold uppercase tracking-[0.08em] text-slate-500">Error 403</p>
            <h1 class="mt-2 text-4xl font-bold text-slate-900">Access Denied</h1>
            <p class="mt-4 text-base leading-7 text-slate-600">
                You don't have permission to access this page.
            </p>

            <button
                type="button"
                class="mt-8 inline-flex min-h-11 w-full items-center justify-center rounded-md bg-blue-900 px-4 py-2 font-semibold text-white transition hover:bg-blue-800 sm:w-auto sm:px-6"
                @click="$emit('home')"
            >
                Go Home
            </button>
        </section>
    </main>
</template>