<template>
    <main class="min-h-screen bg-slate-50 px-4 py-6 grid place-items-center">
        <section class="w-full max-w-3xl rounded-md bg-white p-8 shadow-md sm:p-10">
            <p class="mb-3 text-sm font-bold uppercase tracking-[0.08em] text-slate-600">Student dashboard</p>
            <h1 class="text-3xl font-bold text-slate-900 sm:text-4xl">Signed in as student.</h1>
            <p class="mt-4 max-w-2xl text-base leading-7 text-slate-600">
                Browse the course catalog, manage your basket, and track approval history.
            </p>

            <div class="mt-6 flex flex-col gap-3 sm:flex-row sm:flex-wrap">
                <NuxtLink
                    to="/student/courses"
                    class="inline-flex min-h-11 items-center justify-center rounded-md bg-blue-900 px-4 py-2 font-semibold text-white transition hover:bg-blue-800"
                >
                    Browse courses
                </NuxtLink>
                <NuxtLink
                    to="/student/course-plan"
                    class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"
                >
                    My course plan
                </NuxtLink>
                <NuxtLink
                    to="/student/approval-status"
                    class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"
                >
                    Approval history
                </NuxtLink>
                <button
                    type="button"
                    class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"
                    @click="logout"
                >
                    Logout
                </button>
            </div>
        </section>
    </main>
</template>

<script setup lang="ts">
import { clearStoredToken } from "~/utils/auth";

definePageMeta({
    middleware: ["auth", "student"],
});

async function logout() {
    clearStoredToken();
    await navigateTo("/auth/login");
}
</script>
