<template>
    <main class="min-h-screen bg-slate-50 px-4 py-6">
        <section class="mx-auto w-full max-w-2xl rounded-md bg-white p-6 shadow-md sm:p-8">
            <p class="text-sm font-bold uppercase tracking-[0.08em] text-slate-600">Admin courses</p>
            <h1 class="mt-2 text-3xl font-bold text-slate-900">Delete course</h1>

            <p v-if="isLoading" class="mt-8 text-slate-600">Loading course...</p>
            <p v-else-if="errorMessage" class="mt-8 rounded-md border border-red-200 bg-red-50 px-4 py-3 text-red-700">
                {{ errorMessage }}
            </p>

            <div v-else class="mt-8 space-y-4 rounded-md border border-slate-200 bg-slate-50 p-5">
                <p class="text-sm font-semibold uppercase tracking-[0.08em] text-slate-500">{{ course.code }}</p>
                <h2 class="text-2xl font-bold text-slate-900">{{ course.name }}</h2>
                <p class="text-sm text-slate-600">{{ course.credits }} credits</p>
                <p class="text-sm leading-7 text-slate-600">
                    This action removes the course from the catalog.
                </p>

                <p v-if="submitErrorMessage" class="mt-1 text-md text-red-700">{{ submitErrorMessage }}</p>

                <div class="flex flex-col gap-3 sm:flex-row">
                    <button
                        type="button"
                        class="inline-flex min-h-11 items-center justify-center rounded-md bg-blue-900 px-4 py-2 font-semibold text-white transition hover:bg-blue-800 disabled:cursor-not-allowed disabled:opacity-75"
                        :disabled="isDeleting"
                        @click="deleteCourse"
                    >
                        {{ isDeleting ? "Deleting..." : "Delete course" }}
                    </button>
                    <NuxtLink
                        to="/admin/courses"
                        class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"
                    >
                        Cancel
                    </NuxtLink>
                </div>
            </div>
        </section>
    </main>
</template>

<script setup lang="ts">
import { CourseService } from "~/services/course"
import type { Course } from "~/models/course"
import { extractErrorMessage } from "~/utils/auth"

definePageMeta({
    middleware: ["auth", "admin"],
})

const route = useRoute()
const courseId = Number(route.params.id)

const course = ref<Course>({
    id: courseId,
    code: "",
    name: "",
    credits: 0,
    createdAt: "",
    updatedAt: "",
})

const isLoading = ref(true)
const isDeleting = ref(false)
const errorMessage = ref("")
const submitErrorMessage = ref("")

onMounted(async () => {
    try {
        course.value = await CourseService.getCourse(courseId)
    } catch (error) {
        errorMessage.value = extractErrorMessage(error, "Unable to load course right now.")
    } finally {
        isLoading.value = false
    }
})

async function deleteCourse() {
    isDeleting.value = true
    submitErrorMessage.value = ""

    try {
        await CourseService.deleteCourse(courseId)
        await navigateTo("/admin/courses")
    } catch (error) {
        submitErrorMessage.value = extractErrorMessage(error, "Unable to delete course right now.")
    } finally {
        isDeleting.value = false
    }
}
</script>