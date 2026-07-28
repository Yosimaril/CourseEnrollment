<template>
  <main class="min-h-screen bg-slate-50 px-4 py-6">
    <section class="mx-auto w-full max-w-2xl rounded-md bg-white p-6 shadow-md sm:p-8">
      <div class="flex flex-col gap-3 sm:flex-row sm:items-end sm:justify-between">
        <div>
          <p class="text-sm font-bold uppercase tracking-[0.08em] text-slate-600">Admin courses</p>
          <h1 class="mt-2 text-3xl font-bold text-slate-900">Create course</h1>
        </div>

        <NuxtLink
          to="/admin/courses"
          class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"
        >
          Back to list
        </NuxtLink>
      </div>

      <form class="mt-8 space-y-4" @submit.prevent="submitCreate">
        <label class="block space-y-2">
          <span class="block text-md font-semibold text-slate-900">Code</span>
          <input
            v-model.trim="form.code"
            type="text"
            placeholder="CS101"
            class="w-full rounded-md border border-slate-200 bg-slate-100 px-4 py-2 outline-none transition focus:border-blue-900"
            required
          />
        </label>

        <label class="block space-y-2">
          <span class="block text-md font-semibold text-slate-900">Name</span>
          <input
            v-model.trim="form.name"
            type="text"
            placeholder="Introduction to Programming"
            class="w-full rounded-md border border-slate-200 bg-slate-100 px-4 py-2 outline-none transition focus:border-blue-900"
            required
          />
        </label>

        <label class="block space-y-2">
          <span class="block text-md font-semibold text-slate-900">Credits</span>
          <input
            v-model.number="form.credits"
            type="number"
            min="0"
            step="1"
            placeholder="3"
            class="w-full rounded-md border border-slate-200 bg-slate-100 px-4 py-2 outline-none transition focus:border-blue-900"
            required
          />
        </label>

        <p v-if="errorMessage" class="mt-1 text-md text-red-700">{{ errorMessage }}</p>

        <button
          type="submit"
          class="mt-2 w-full rounded-md bg-blue-900 px-4 py-2 font-semibold text-white transition hover:bg-blue-800 disabled:cursor-not-allowed disabled:opacity-75"
          :disabled="isSubmitting"
        >
          {{ isSubmitting ? "Creating..." : "Create course" }}
        </button>
      </form>
    </section>
  </main>
</template>

<script setup lang="ts">
import { CourseService } from "~/services/course"
import { extractErrorMessage } from "~/utils/auth"

definePageMeta({
  middleware: ["auth", "admin"],
})

const form = reactive({
  code: "",
  name: "",
  credits: 0,
})

const isSubmitting = ref(false)
const errorMessage = ref("")

async function submitCreate() {
  isSubmitting.value = true
  errorMessage.value = ""

  try {
    await CourseService.createCourse({
      code: form.code,
      name: form.name,
      credits: Number(form.credits),
    })

    await navigateTo({ path: "/admin/courses", query: { refresh: Date.now().toString() } })
  } catch (error) {
    errorMessage.value = extractErrorMessage(error, "Unable to create course right now.")
  } finally {
    isSubmitting.value = false
  }
}
</script>