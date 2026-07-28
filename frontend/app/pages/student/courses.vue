<template>
  <main class="min-h-screen bg-slate-50 px-4 py-6">
    <section class="mx-auto w-full max-w-6xl rounded-md bg-white p-6 shadow-md sm:p-8">
      <div class="flex flex-col gap-3 sm:flex-row sm:items-end sm:justify-between">
        <div>
          <p class="text-sm font-bold uppercase tracking-[0.08em] text-slate-600">Student courses</p>
          <h1 class="mt-2 text-3xl font-bold text-slate-900">Course catalog</h1>
          <p class="mt-2 max-w-2xl text-base leading-7 text-slate-600">
            Load available courses and add the ones you want to your picked courses list.
          </p>
        </div>

        <NuxtLink
          to="/student/dashboard"
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
          <div class="mt-5 h-11 rounded bg-slate-200"></div>
        </div>
      </div>

      <p v-else-if="errorMessage" class="mt-8 rounded-md border border-red-200 bg-red-50 px-4 py-3 text-red-700">
        {{ errorMessage }}
      </p>

      <div v-else class="mt-8 grid gap-4 sm:grid-cols-2 xl:grid-cols-3">
        <article
          v-for="course in courses"
          :key="course.id"
          class="flex h-full flex-col rounded-md border border-slate-200 bg-slate-50 p-5"
        >
          <p class="text-sm font-semibold uppercase tracking-[0.08em] text-slate-500">
            {{ course.code }}
          </p>
          <h2 class="mt-3 text-xl font-bold text-slate-900">
            {{ course.name }}
          </h2>
          <p class="mt-2 text-sm text-slate-600">
            {{ course.credits }} credits
          </p>

          <button
            type="button"
            class="mt-5 inline-flex min-h-11 items-center justify-center rounded-md bg-blue-900 px-4 py-2 font-semibold text-white transition hover:bg-blue-800 disabled:cursor-not-allowed disabled:opacity-75"
            :disabled="pendingCourseIds.includes(course.id) || addedCourseIds.includes(course.id)"
            @click="addCourse(course.id)"
          >
            {{
              addedCourseIds.includes(course.id)
                ? "Added to cart"
                : pendingCourseIds.includes(course.id)
                  ? "Adding..."
                  : "Add to cart"
            }}
          </button>
        </article>
      </div>
    </section>
  </main>
</template>

<script setup lang="ts">
import { CourseService } from "~/services/course"
import { CoursePlanItemService } from "~/services/coursePlanItem"
import type { Course } from "~/models/course"
import { extractErrorMessage } from "~/utils/auth"

definePageMeta({
  middleware: ["auth", "student"],
})

const courses = ref<Course[]>([])
const isLoading = ref(true)
const errorMessage = ref("")
const pendingCourseIds = ref<number[]>([])
const addedCourseIds = ref<number[]>([])

onMounted(async () => {
  try {
    const response = await CourseService.getCourses()
    courses.value = Array.isArray(response) ? response : []
  } catch (error) {
    errorMessage.value = extractErrorMessage(error, "Unable to load courses right now.")
  } finally {
    isLoading.value = false
  }
})

async function addCourse(courseId: number) {
  if (pendingCourseIds.value.includes(courseId) || addedCourseIds.value.includes(courseId)) {
    return
  }

  pendingCourseIds.value = [...pendingCourseIds.value, courseId]

  try {
    await CoursePlanItemService.addToPickedCourses(courseId)
    addedCourseIds.value = [...addedCourseIds.value, courseId]
  } catch (error) {
    errorMessage.value = extractErrorMessage(error, "Unable to add this course right now.")
  } finally {
    pendingCourseIds.value = pendingCourseIds.value.filter((id) => id !== courseId)
  }
}
</script>