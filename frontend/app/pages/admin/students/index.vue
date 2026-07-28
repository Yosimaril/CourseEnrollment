<template>
    <main class="min-h-screen bg-slate-50 px-4 py-6">
        <section class="mx-auto w-full max-w-6xl rounded-md bg-white p-6 shadow-md sm:p-8">
            <div class="flex flex-col gap-3 sm:flex-row sm:items-end sm:justify-between">
                <div>
                    <p class="text-sm font-bold uppercase tracking-[0.08em] text-slate-600">Admin students</p>
                    <h1 class="mt-2 text-3xl font-bold text-slate-900">Student list</h1>
                    <p class="mt-2 max-w-2xl text-base leading-7 text-slate-600">
                        Review the student accounts registered in the system.
                    </p>
                </div>

                <NuxtLink
                    to="/admin/dashboard"
                    class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"
                >
                    Back to dashboard
                </NuxtLink>
            </div>

            <div v-if="isLoading" class="mt-8 grid gap-4 sm:grid-cols-2 xl:grid-cols-3">
                <div v-for="index in 6" :key="index" class="rounded-md border border-slate-200 bg-slate-50 p-5">
                    <div class="h-4 w-24 rounded bg-slate-200"></div>
                    <div class="mt-4 h-6 w-40 rounded bg-slate-200"></div>
                    <div class="mt-3 h-4 w-20 rounded bg-slate-200"></div>
                </div>
            </div>

            <p v-else-if="errorMessage" class="mt-8 rounded-md border border-red-200 bg-red-50 px-4 py-3 text-red-700">
                {{ errorMessage }}
            </p>

            <div v-else class="mt-8 grid gap-4 sm:grid-cols-2 xl:grid-cols-3">
                <article
                    v-for="student in students"
                    :key="student.id"
                    class="flex h-full flex-col rounded-md border border-slate-200 bg-slate-50 p-5"
                >
                    <p class="text-sm font-semibold uppercase tracking-[0.08em] text-slate-500">
                        {{ student.role }}
                    </p>
                    <h2 class="mt-3 text-xl font-bold text-slate-900">
                        {{ student.username }}
                    </h2>
                    <p class="mt-2 text-sm text-slate-600">
                        {{ student.email }}
                    </p>
                    <p v-if="student.nrp" class="mt-2 text-sm text-slate-600">
                        NRP: {{ student.nrp }}
                    </p>
                    <p v-if="student.maxCredits !== undefined" class="mt-2 text-sm text-slate-600">
                        Max credits: {{ student.maxCredits }}
                    </p>
                </article>
            </div>
        </section>
    </main>
</template>

<script setup lang="ts">
import { StudentService } from "~/services/student"
import type { Student } from "~/models/student"
import { extractErrorMessage } from "~/utils/auth"

definePageMeta({
    middleware: ["auth", "admin"],
})

const students = ref<Student[]>([])
const isLoading = ref(true)
const errorMessage = ref("")

onMounted(async () => {
    try {
        const response = await StudentService.getStudents()
        students.value = Array.isArray(response) ? response : []
    } catch (error) {
        errorMessage.value = extractErrorMessage(error, "Unable to load students right now.")
    } finally {
        isLoading.value = false
    }
})
</script>