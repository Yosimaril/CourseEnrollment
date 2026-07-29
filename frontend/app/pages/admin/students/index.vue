<template>
    <main class="min-h-screen bg-slate-50 px-4 py-6">
        <section class="mx-auto w-full max-w-6xl rounded-md bg-white p-6 shadow-md sm:p-8">
            <div class="flex flex-col gap-3 sm:flex-row sm:items-end sm:justify-between">
                <div>
                    <p class="text-sm font-bold uppercase tracking-[0.08em] text-slate-600">Admin students</p>
                    <h1 class="mt-2 text-3xl font-bold text-slate-900">Student list</h1>
                    <p class="mt-2 max-w-2xl text-base leading-7 text-slate-600">Review every student account, then edit, remove, or create a new one.</p>
                </div>

                <div class="flex flex-col gap-3 sm:flex-row">
                    <NuxtLink to="/admin/dashboard" class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"> Back to dashboard </NuxtLink>
                    <NuxtLink to="/admin/students/create" class="inline-flex min-h-11 items-center justify-center rounded-md bg-blue-900 px-4 py-2 font-semibold text-white transition hover:bg-blue-800"> Add student </NuxtLink>
                </div>
            </div>

            <div v-if="isLoading" class="mt-8 grid gap-4 sm:grid-cols-2 xl:grid-cols-3">
                <div v-for="index in 6" :key="index" class="rounded-md border border-slate-200 bg-slate-50 p-5">
                    <div class="h-4 w-24 rounded bg-slate-200"></div>
                    <div class="mt-4 h-6 w-40 rounded bg-slate-200"></div>
                    <div class="mt-3 h-4 w-28 rounded bg-slate-200"></div>
                    <div class="mt-3 h-4 w-20 rounded bg-slate-200"></div>
                    <div class="mt-5 h-11 rounded bg-slate-200"></div>
                </div>
            </div>

            <p v-else-if="errorMessage" class="mt-8 rounded-md border border-red-200 bg-red-50 px-4 py-3 text-red-700">
                {{ errorMessage }}
            </p>

            <div v-else class="mt-8 grid gap-4 sm:grid-cols-2 xl:grid-cols-3">
                <article v-for="student in students" :key="student.id" class="flex h-full flex-col rounded-md border border-slate-200 bg-slate-50 p-5">
                    <p class="text-sm font-semibold uppercase tracking-[0.08em] text-slate-500">
                        {{ student.role }}
                    </p>
                    <h2 class="mt-3 text-xl font-bold text-slate-900">
                        {{ student.username }}
                    </h2>
                    <div class="mt-3 space-y-2 text-sm text-slate-600">
                        <p>{{ student.email }}</p>
                        <p v-if="student.nrp">NRP: {{ student.nrp }}</p>
                        <p v-if="student.maxCredits !== undefined && student.maxCredits !== null">Max credits: {{ student.maxCredits }}</p>
                    </div>

                    <div class="mt-5 flex gap-3">
                        <NuxtLink :to="`/admin/students/${student.id}`" class="inline-flex flex-1 min-h-11 items-center justify-center rounded-md bg-blue-900 px-4 py-2 font-semibold text-white transition hover:bg-blue-800"> Edit </NuxtLink>
                        <button type="button" @click="confirmDelete(student)" class="inline-flex flex-1 min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900">Delete</button>
                    </div>
                </article>
            </div>
        </section>

        <Modal :is-open="isModalOpen" @close="closeModal">
            <template #header>Delete Student</template>
            <template #body>
                <p>
                    Are you sure you want to delete the student "{{ studentToDelete?.username }}" <span v-if="studentToDelete?.nrp">({{ studentToDelete?.nrp }})</span>?
                </p>
                <p class="mt-2 text-sm text-red-600">This action cannot be undone.</p>
                <p v-if="submitErrorMessage" class="mt-4 text-md text-red-700">{{ submitErrorMessage }}</p>
            </template>
            <template #footer>
                <button type="button" @click="closeModal" class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900">Cancel</button>
                <button type="button" @click="deleteStudent" class="inline-flex min-h-11 items-center justify-center rounded-md bg-red-600 px-4 py-2 font-semibold text-white transition hover:bg-red-500 disabled:cursor-not-allowed disabled:opacity-75" :disabled="isDeleting">
                    {{ isDeleting ? "Deleting..." : "Delete" }}
                </button>
            </template>
        </Modal>
    </main>
</template>

<script setup lang="ts">
import Modal from "~/components/Modal.vue";
import type { Student } from "~/models/student";
import { StudentService } from "~/services/student";
import { extractErrorMessage } from "~/utils/auth";

definePageMeta({
    middleware: ["auth", "admin"],
});

const route = useRoute();
const students = ref<Student[]>([]);
const isLoading = ref(true);
const errorMessage = ref("");

const isModalOpen = ref(false);
const studentToDelete = ref<Student | null>(null);
const isDeleting = ref(false);
const submitErrorMessage = ref("");

onMounted(async () => {
    await fetchStudents();
});

watch(
    () => route.query.refresh,
    async () => {
        await fetchStudents();
    },
);

async function fetchStudents() {
    isLoading.value = true;
    errorMessage.value = "";

    try {
        const response = await StudentService.getStudents();
        students.value = Array.isArray(response) ? response : [];
    } catch (error) {
        errorMessage.value = extractErrorMessage(error, "Unable to load students right now.");
    } finally {
        isLoading.value = false;
    }
}

function confirmDelete(student: Student) {
    studentToDelete.value = student;
    submitErrorMessage.value = "";
    isModalOpen.value = true;
}

function closeModal() {
    isModalOpen.value = false;
    studentToDelete.value = null;
    isDeleting.value = false;
    submitErrorMessage.value = "";
}

async function deleteStudent() {
    if (!studentToDelete.value) return;

    isDeleting.value = true;
    submitErrorMessage.value = "";

    try {
        await StudentService.deleteStudent(studentToDelete.value.id);
        await fetchStudents();
        closeModal();
    } catch (error) {
        submitErrorMessage.value = extractErrorMessage(error, "Unable to delete student right now.");
    } finally {
        isDeleting.value = false;
    }
}
</script>
