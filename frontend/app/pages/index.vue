<template>
    <main class="min-h-screen bg-slate-50 px-4 py-6 grid place-items-center">
        <section class="w-full max-w-3xl rounded-md bg-white p-8 shadow-md sm:p-10">
            <p class="mb-3 text-sm font-bold uppercase tracking-[0.08em] text-slate-600">Course Enrollment</p>
            <h1 class="text-4xl font-bold leading-tight text-slate-900 sm:text-5xl">
                Simple login and registration for students and admins.
            </h1>
            <p class="mt-4 max-w-2xl text-base leading-7 text-slate-600">
                Sign in, create an account, and move straight to the right dashboard based on your role.
            </p>

            <div class="mt-6 flex flex-col gap-3 sm:flex-row">
                <NuxtLink
                    to="/auth/login"
                    class="inline-flex min-h-11 items-center justify-center rounded-md bg-blue-900 px-4 py-2 font-semibold text-white transition hover:bg-blue-800"
                >
                    Login
                </NuxtLink>
                <NuxtLink
                    to="/auth/register"
                    class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"
                >
                    Register
                </NuxtLink>
            </div>

            <p class="mt-4 text-sm text-slate-500">{{ sessionState }}</p>
        </section>
    </main>
</template>

<script setup lang="ts">
import { getSessionRedirectPath } from "~/utils/auth";

const sessionState = ref("No active session");

onMounted(async () => {
    const redirectPath = getSessionRedirectPath();

    if (redirectPath) {
        sessionState.value = `Redirecting to ${redirectPath}`;
        await navigateTo(redirectPath);
    }
});
</script>
