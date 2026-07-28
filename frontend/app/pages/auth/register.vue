<template>
    <main class="min-h-screen bg-slate-50 px-4 py-6 grid place-items-center">
        <section class="w-full max-w-md rounded-md bg-white p-8 shadow-md sm:p-10">
            <h1 class="mb-6 text-center text-3xl font-bold text-slate-900">Register</h1>

            <form class="space-y-4" @submit.prevent="submitRegister">
                <label class="space-y-2 block">
                    <span class="block text-md font-semibold text-slate-900">Username</span>
                    <input
                        v-model.trim="form.username"
                        type="text"
                        placeholder="yosimaril"
                        autocomplete="username"
                        required
                        class="w-full rounded-md border border-slate-200 bg-slate-100 px-4 py-2 outline-none transition focus:border-blue-900"
                    />
                </label>

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
                        placeholder="password must be at least 6 characters"
                        autocomplete="new-password"
                        required
                        minlength="6"
                        class="w-full rounded-md border border-slate-200 bg-slate-100 px-4 py-2 outline-none transition focus:border-blue-900"
                    />
                </label>

                <label class="space-y-2 block">
                    <span class="block text-md font-semibold text-slate-900">Role</span>
                    <select
                        v-model="form.role"
                        required
                        class="w-full rounded-md border border-slate-200 bg-slate-100 px-4 py-2 outline-none transition focus:border-blue-900"
                    >
                        <option :value="UserRoleEnum.STUDENT">Student</option>
                        <option :value="UserRoleEnum.ADMIN">Admin</option>
                    </select>
                </label>

                <template v-if="form.role === UserRoleEnum.STUDENT">
                    <label class="space-y-2 block">
                        <span class="block text-md font-semibold text-slate-900">NRP</span>
                        <input
                            v-model.trim="form.nrp"
                            type="text"
                            placeholder="160000001"
                            autocomplete="off"
                            class="w-full rounded-md border border-slate-200 bg-slate-100 px-4 py-2 outline-none transition focus:border-blue-900"
                        />
                    </label>

                    <label class="space-y-2 block">
                        <span class="block text-md font-semibold text-slate-900">Max credits</span>
                        <input
                            v-model="form.maxCredits"
                            type="number"
                            min="0"
                            step="1"
                            placeholder="24"
                            class="w-full rounded-md border border-slate-200 bg-slate-100 px-4 py-2 outline-none transition focus:border-blue-900"
                        />
                    </label>
                </template>

                <p v-if="errorMessage" class="mt-1 text-md text-red-700">{{ errorMessage }}</p>

                <button
                    type="submit"
                    class="mt-4 w-full cursor-pointer rounded-md bg-blue-900 px-4 py-2 font-semibold text-white transition hover:bg-blue-800 disabled:cursor-not-allowed disabled:opacity-75"
                    :disabled="isSubmitting"
                >
                    {{ isSubmitting ? "Registering..." : "Register" }}
                </button>

                <p class="text-center text-sm text-slate-500">
                    Already have an account?
                    <NuxtLink to="/auth/login" class="font-bold text-blue-900 hover:underline">Login here</NuxtLink>
                </p>
            </form>
        </section>
    </main>
</template>

<script setup lang="ts">
import { UserRoleEnum } from "~/constants/userRoleEnum";
import { AuthService } from "~/services/auth";
import { extractErrorMessage, getDashboardPath, setStoredToken } from "~/utils/auth";

definePageMeta({
    middleware: "guest",
});

const form = reactive({
    username: "",
    email: "",
    password: "",
    role: UserRoleEnum.STUDENT,
    nrp: "",
    maxCredits: "",
});

const isSubmitting = ref(false);
const errorMessage = ref("");

async function submitRegister() {
    isSubmitting.value = true;
    errorMessage.value = "";

    const request = {
        username: form.username,
        email: form.email,
        password: form.password,
        role: form.role,
        ...(form.role === UserRoleEnum.STUDENT && form.nrp ? { nrp: form.nrp } : {}),
        ...(form.role === UserRoleEnum.STUDENT && form.maxCredits !== "" ? { max_credits: Number(form.maxCredits) } : {}),
    };

    try {
        const response = await AuthService.register(request);

        setStoredToken(response.token);
        await navigateTo(getDashboardPath(response.user.role));
    } catch (error) {
        errorMessage.value = extractErrorMessage(error, "Unable to register right now.");
    } finally {
        isSubmitting.value = false;
    }
}
</script>
