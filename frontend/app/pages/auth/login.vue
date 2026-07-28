<template>
    <main class="min-h-screen bg-slate-50 px-4 py-6 grid place-items-center">
        <section class="w-full max-w-md rounded-md bg-white p-8 shadow-md sm:p-10">
            <h1 class="mb-6 text-center text-3xl font-bold text-slate-900">Login</h1>

            <form class="space-y-4" @submit.prevent="submitLogin">
                <label class="space-y-2 block">
                    <span class="block text-md font-semibold text-slate-900">Email</span>
                    <input
                        v-model.trim="form.email"
                        type="email"
                        placeholder="yosimaril@example.com"
                        autocomplete="email"
                        required
                        class="w-full rounded-md border border-slate-200 bg-slate-100 px-4 py-2 outline-none transition focus:border-blue-900"
                    />
                </label>

                <label class="space-y-2 block">
                    <span class="block text-md font-semibold text-slate-900">Password</span>
                    <input
                        v-model="form.password"
                        type="password"
                        placeholder="Enter your password"
                        autocomplete="current-password"
                        required
                        minlength="6"
                        class="w-full rounded-md border border-slate-200 bg-slate-100 px-4 py-2 outline-none transition focus:border-blue-900"
                    />
                </label>

                <p v-if="errorMessage" class="mt-1 text-md text-red-700">{{ errorMessage }}</p>

                <button
                    type="submit"
                    class="mt-4 w-full cursor-pointer rounded-md bg-blue-900 px-4 py-2 font-semibold text-white transition hover:bg-blue-800 disabled:cursor-not-allowed disabled:opacity-75"
                    :disabled="isSubmitting"
                >
                    {{ isSubmitting ? "Logging in..." : "Login" }}
                </button>

                <p class="text-center text-sm text-slate-500">
                    No account yet?
                    <NuxtLink to="/auth/register" class="font-bold text-blue-900 hover:underline">Register here</NuxtLink>
                </p>
            </form>
        </section>
    </main>
</template>

<script setup lang="ts">
import { AuthService } from "~/services/auth";
import { extractErrorMessage, getDashboardPath, getRoleFromToken, setStoredToken } from "~/utils/auth";

definePageMeta({
    middleware: "guest",
});

const form = reactive({
    email: "",
    password: "",
});

const isSubmitting = ref(false);
const errorMessage = ref("");

async function submitLogin() {
    isSubmitting.value = true;
    errorMessage.value = "";

    try {
        const response = await AuthService.login({
            email: form.email,
            password: form.password,
        });

        setStoredToken(response.token);
        await navigateTo(getDashboardPath(getRoleFromToken(response.token)));
    } catch (error) {
        errorMessage.value = extractErrorMessage(error, "Unable to log in right now.");
    } finally {
        isSubmitting.value = false;
    }
}
</script>
