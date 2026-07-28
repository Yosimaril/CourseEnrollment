<template>
    <main class="min-h-screen bg-slate-50 px-4 py-6">
        <section class="mx-auto w-full max-w-6xl rounded-md bg-white p-6 shadow-md sm:p-8">
            <div class="flex flex-col gap-3 sm:flex-row sm:items-end sm:justify-between">
                <div>
                    <p class="text-sm font-bold uppercase tracking-[0.08em] text-slate-600">Student course plan</p>
                    <h1 class="mt-2 text-3xl font-bold text-slate-900">Basket</h1>
                    <p class="mt-2 max-w-2xl text-base leading-7 text-slate-600">
                        Review picked courses, remove any item you no longer want, then submit the plan for approval.
                    </p>
                </div>

                <NuxtLink
                    to="/student/dashboard"
                    class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"
                >
                    Back to dashboard
                </NuxtLink>
            </div>

            <p v-if="isLoading" class="mt-8 text-slate-600">Loading course plan...</p>
            <p v-else-if="successMessage" class="mt-8 rounded-md border border-emerald-200 bg-emerald-50 px-4 py-3 text-emerald-700">
                {{ successMessage }}
            </p>
            <p v-else-if="errorMessage" class="mt-8 rounded-md border border-red-200 bg-red-50 px-4 py-3 text-red-700">
                {{ errorMessage }}
            </p>

            <div v-else>
                <div v-if="!coursePlan || courseItems.length === 0" class="mt-8 rounded-md border border-slate-200 bg-slate-50 p-5 text-slate-600">
                    Your basket is empty. Pick courses first.
                </div>

                <div v-else class="mt-8 grid gap-4 xl:grid-cols-[minmax(0,1fr)_18rem]">
                    <div class="grid gap-4">
                        <article
                            v-for="item in courseItems"
                            :key="`${item.coursePlanId}-${item.courseId}`"
                            class="rounded-md border border-slate-200 bg-slate-50 p-5"
                        >
                            <div class="flex flex-col gap-4 sm:flex-row sm:items-start sm:justify-between">
                                <div>
                                    <p class="text-sm font-semibold uppercase tracking-[0.08em] text-slate-500">
                                        {{ item.course?.code }}
                                    </p>
                                    <h2 class="mt-2 text-xl font-bold text-slate-900">
                                        {{ item.course?.name }}
                                    </h2>
                                    <p class="mt-2 text-sm text-slate-600">
                                        {{ item.course?.credits }} credits
                                    </p>
                                </div>

                                <button
                                    type="button"
                                    class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900 disabled:cursor-not-allowed disabled:opacity-75"
                                    :disabled="pendingDeleteIds.includes(item.courseId)"
                                    @click="removeItem(item.courseId)"
                                >
                                    {{ pendingDeleteIds.includes(item.courseId) ? "Removing..." : "Remove" }}
                                </button>
                            </div>
                        </article>
                    </div>

                    <aside class="rounded-md border border-slate-200 bg-slate-50 p-5">
                        <p class="text-sm font-bold uppercase tracking-[0.08em] text-slate-600">Summary</p>
                        <p class="mt-3 text-3xl font-bold text-slate-900">{{ totalCredits }} credits</p>
                        <p class="mt-2 text-sm text-slate-600">{{ courseItems.length }} course{{ courseItems.length === 1 ? "" : "s" }} in basket</p>

                        <button
                            type="button"
                            class="mt-6 w-full rounded-md bg-blue-900 px-4 py-2 font-semibold text-white transition hover:bg-blue-800 disabled:cursor-not-allowed disabled:opacity-75"
                            :disabled="isSubmitting || courseItems.length === 0"
                            @click="submitPlan"
                        >
                            {{ isSubmitting ? "Submitting..." : "Submit for approval" }}
                        </button>
                    </aside>
                </div>
            </div>
        </section>
    </main>
</template>

<script setup lang="ts">
import { CoursePlanService } from "~/services/coursePlan"
import { CoursePlanItemService } from "~/services/coursePlanItem"
import type { CoursePlan } from "~/models/coursePlan"
import { extractErrorMessage } from "~/utils/auth"

definePageMeta({
    middleware: ["auth", "student"],
})

const coursePlan = ref<CoursePlan | null>(null)
const isLoading = ref(true)
const isSubmitting = ref(false)
const successMessage = ref("")
const errorMessage = ref("")
const pendingDeleteIds = ref<number[]>([])

const totalCredits = computed(() => {
    return courseItems.value.reduce((sum, item) => sum + (item.course?.credits ?? 0), 0)
})

const courseItems = computed(() => coursePlan.value?.items ?? [])

onMounted(async () => {
    try {
        coursePlan.value = await CoursePlanService.getMyCoursePlan()
    } catch (error) {
        const message = extractErrorMessage(error, "Unable to load course plan right now.")
        if (!message.toLowerCase().includes("not found")) {
            errorMessage.value = message
        }
    } finally {
        isLoading.value = false
    }
})

async function removeItem(courseId: number) {
    if (pendingDeleteIds.value.includes(courseId)) {
        return
    }

    pendingDeleteIds.value = [...pendingDeleteIds.value, courseId]

    try {
        await CoursePlanItemService.removeFromPickedCourses(courseId)
        coursePlan.value = await CoursePlanService.getMyCoursePlan().catch(() => null)
    } catch (error) {
        errorMessage.value = extractErrorMessage(error, "Unable to remove this course right now.")
    } finally {
        pendingDeleteIds.value = pendingDeleteIds.value.filter((id) => id !== courseId)
    }
}

async function submitPlan() {
    isSubmitting.value = true
    errorMessage.value = ""
    successMessage.value = ""

    try {
        await CoursePlanService.submitMyCoursePlan()
        coursePlan.value = null
        successMessage.value = "Course plan submitted for approval."
    } catch (error) {
        errorMessage.value = extractErrorMessage(error, "Unable to submit course plan right now.")
    } finally {
        isSubmitting.value = false
    }
}
</script>