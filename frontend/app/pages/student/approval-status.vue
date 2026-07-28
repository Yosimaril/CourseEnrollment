<template>
    <main class="min-h-screen bg-slate-50 px-4 py-6">
        <section class="mx-auto w-full max-w-6xl rounded-md bg-white p-6 shadow-md sm:p-8">
            <div class="flex flex-col gap-3 sm:flex-row sm:items-end sm:justify-between">
                <div>
                    <p class="text-sm font-bold uppercase tracking-[0.08em] text-slate-600">Approval history</p>
                    <h1 class="mt-2 text-3xl font-bold text-slate-900">Submission status</h1>
                    <p class="mt-2 max-w-2xl text-base leading-7 text-slate-600">
                        Review every course plan and cancel any request that has not been approved yet.
                    </p>
                </div>

                <NuxtLink
                    to="/student/dashboard"
                    class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"
                >
                    Back to dashboard
                </NuxtLink>
            </div>

            <p v-if="isLoading" class="mt-8 text-slate-600">Loading approval history...</p>
            <p v-else-if="errorMessage" class="mt-8 rounded-md border border-red-200 bg-red-50 px-4 py-3 text-red-700">
                {{ errorMessage }}
            </p>

            <div v-else class="mt-8 space-y-4">
                <div v-if="plans.length === 0" class="rounded-md border border-slate-200 bg-slate-50 p-5 text-slate-600">
                    No course plan history yet.
                </div>

                <article
                    v-for="plan in plans"
                    :key="plan.id"
                    class="rounded-md border border-slate-200 bg-slate-50 p-5"
                >
                    <div class="flex flex-col gap-3 sm:flex-row sm:items-start sm:justify-between">
                        <div>
                            <p class="text-sm font-semibold uppercase tracking-[0.08em] text-slate-500">
                                Plan #{{ plan.id }}
                            </p>
                            <h2 class="mt-2 text-2xl font-bold text-slate-900">
                                {{ plan.status }}
                            </h2>
                            <p class="mt-2 text-sm text-slate-600">
                                {{ plan.items.length }} item{{ plan.items.length === 1 ? "" : "s" }}
                            </p>
                        </div>

                        <button
                            v-if="canCancel(plan.status)"
                            type="button"
                            class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900 disabled:cursor-not-allowed disabled:opacity-75"
                            :disabled="pendingCancelIds.includes(plan.id)"
                            @click="cancelPlan(plan.id)"
                        >
                            {{ pendingCancelIds.includes(plan.id) ? "Cancelling..." : "Cancel request" }}
                        </button>
                    </div>

                    <div class="mt-5 grid gap-3 sm:grid-cols-2 xl:grid-cols-3">
                        <div
                            v-for="item in plan.items"
                            :key="`${plan.id}-${item.courseId}`"
                            class="rounded-md border border-slate-200 bg-white p-4"
                        >
                            <p class="text-sm font-semibold uppercase tracking-[0.08em] text-slate-500">
                                {{ item.course?.code }}
                            </p>
                            <h3 class="mt-2 text-lg font-bold text-slate-900">
                                {{ item.course?.name }}
                            </h3>
                            <p class="mt-2 text-sm text-slate-600">
                                {{ item.status }}
                            </p>
                        </div>
                    </div>
                </article>
            </div>
        </section>
    </main>
</template>

<script setup lang="ts">
import { CoursePlanService } from "~/services/coursePlan"
import type { CoursePlan } from "~/models/coursePlan"
import { CoursePlanStatusEnum } from "~/constants/coursePlanStatusEnum"
import { extractErrorMessage } from "~/utils/auth"

definePageMeta({
    middleware: ["auth", "student"],
})

const plans = ref<CoursePlan[]>([])
const isLoading = ref(true)
const errorMessage = ref("")
const pendingCancelIds = ref<number[]>([])

onMounted(async () => {
    try {
        const response = await CoursePlanService.getMyCoursePlans()
        plans.value = Array.isArray(response) ? response : []
    } catch (error) {
        errorMessage.value = extractErrorMessage(error, "Unable to load approval history right now.")
    } finally {
        isLoading.value = false
    }
})

function canCancel(status: CoursePlanStatusEnum) {
    return status !== CoursePlanStatusEnum.APPROVED
}

async function cancelPlan(id: number) {
    if (pendingCancelIds.value.includes(id)) {
        return
    }

    pendingCancelIds.value = [...pendingCancelIds.value, id]

    try {
        await CoursePlanService.cancelMyCoursePlan(id)
        plans.value = plans.value.filter((plan) => plan.id !== id)
    } catch (error) {
        errorMessage.value = extractErrorMessage(error, "Unable to cancel this request right now.")
    } finally {
        pendingCancelIds.value = pendingCancelIds.value.filter((planId) => planId !== id)
    }
}
</script>