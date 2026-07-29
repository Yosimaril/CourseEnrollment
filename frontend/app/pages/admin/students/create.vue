<template>
    <main class="min-h-screen bg-slate-50 px-4 py-6">
        <section class="mx-auto w-full max-w-2xl rounded-md bg-white p-6 shadow-md sm:p-8">
            <div class="flex flex-col gap-3 sm:flex-row sm:items-end sm:justify-between">
                <div>
                    <p class="text-sm font-bold uppercase tracking-[0.08em] text-slate-600">Admin students</p>
                    <h1 class="mt-2 text-3xl font-bold text-slate-900">Create student</h1>
                    <p class="mt-2 max-w-2xl text-base leading-7 text-slate-600">Add a new student account and set the initial profile details.</p>
                </div>

                <NuxtLink to="/admin/students" class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"> Back to list </NuxtLink>
            </div>

            <form class="mt-8 space-y-4" @submit.prevent="submitCreate">
                <label class="block space-y-2">
                    <span class="block text-md font-semibold text-slate-900">Username</span>
                    <input v-model.trim="form.username" type="text" placeholder="jane.doe" class="w-full rounded-md border border-slate-200 bg-slate-100 px-4 py-2 outline-none transition focus:border-blue-900" required />
                </label>

                <label class="block space-y-2">
                    <span class="block text-md font-semibold text-slate-900">Email</span>
                    <input v-model.trim="form.email" type="email" placeholder="jane@example.com" class="w-full rounded-md border border-slate-200 bg-slate-100 px-4 py-2 outline-none transition focus:border-blue-900" required />
                </label>

                <label class="block space-y-2">
                    <span class="block text-md font-semibold text-slate-900">Password</span>
                    <input v-model="form.password" type="password" placeholder="Minimum 6 characters" class="w-full rounded-md border border-slate-200 bg-slate-100 px-4 py-2 outline-none transition focus:border-blue-900" required />
                </label>

                <label class="block space-y-2">
                    <span class="block text-md font-semibold text-slate-900">NRP</span>
                    <input v-model.trim="form.nrp" type="text" placeholder="12345678" class="w-full rounded-md border border-slate-200 bg-slate-100 px-4 py-2 outline-none transition focus:border-blue-900" />
                </label>

                <label class="block space-y-2">
                    <span class="block text-md font-semibold text-slate-900">Max credits</span>
                    <input v-model="form.maxCredits" type="number" min="0" step="1" placeholder="24" class="w-full rounded-md border border-slate-200 bg-slate-100 px-4 py-2 outline-none transition focus:border-blue-900" />
                </label>

                <p v-if="errorMessage" class="mt-1 text-md text-red-700">{{ errorMessage }}</p>

                <button type="submit" class="mt-2 w-full rounded-md bg-blue-900 px-4 py-2 font-semibold text-white transition hover:bg-blue-800 disabled:cursor-not-allowed disabled:opacity-75" :disabled="isSubmitting">
                    {{ isSubmitting ? "Creating..." : "Create student" }}
                </button>
            </form>
        </section>
    </main>
</template>

<script setup lang="ts">
import { UserRoleEnum } from "~/constants/userRoleEnum";
import { StudentService } from "~/services/student";
import { extractErrorMessage } from "~/utils/auth";

definePageMeta({
    middleware: ["auth", "admin"],
});

const form = reactive({
    username: "",
    email: "",
    password: "",
    nrp: "",
    maxCredits: "",
});

const isSubmitting = ref(false);
const errorMessage = ref("");

async function submitCreate() {
    isSubmitting.value = true;
    errorMessage.value = "";

    try {
        await StudentService.createStudent({
            username: form.username,
            email: form.email,
            password: form.password,
            role: UserRoleEnum.STUDENT,
            nrp: form.nrp ? form.nrp : undefined,
            max_credits: form.maxCredits !== "" ? Number(form.maxCredits) : undefined,
        });

        await navigateTo({ path: "/admin/students", query: { refresh: Date.now().toString() } });
    } catch (error) {
        errorMessage.value = extractErrorMessage(error, "Unable to create student right now.");
    } finally {
        isSubmitting.value = false;
    }
}
</script>
