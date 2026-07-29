<template>
    <main class="min-h-screen bg-slate-50 px-4 py-6">
        <section class="mx-auto w-full max-w-7xl rounded-md bg-white p-6 shadow-md sm:p-8">
            <div class="flex flex-col gap-3 sm:flex-row sm:items-end sm:justify-between">
                <div>
                    <p class="text-sm font-bold uppercase tracking-[0.08em] text-slate-600">Admin approvals</p>
                    <h1 class="mt-2 text-3xl font-bold text-slate-900">Pending course plans</h1>
                    <p class="mt-2 max-w-2xl text-base leading-7 text-slate-600">
                        Approve or reject all courses in a submission, or review each course item individually.
                    </p>
                </div>

                <NuxtLink
                    to="/admin/dashboard"
                    class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"
                >
                    Back to dashboard
                </NuxtLink>
            </div>

            <p v-if="isLoading" class="mt-8 text-slate-600">Loading pending approvals...</p>
            <p v-else-if="errorMessage" class="mt-8 rounded-md border border-red-200 bg-red-50 px-4 py-3 text-red-700">
                {{ errorMessage }}
            </p>

            <div v-else class="mt-8 space-y-4">
                <div v-if="plans.length === 0" class="rounded-md border border-slate-200 bg-slate-50 p-5 text-slate-600">
                    No submitted course plans to review.
                </div>

                <article
                    v-for="plan in plans"
                    :key="plan.id"
                    class="rounded-md border border-slate-200 bg-slate-50 p-5"
                >
                    <div class="flex flex-col gap-3 lg:flex-row lg:items-start lg:justify-between">
                        <div>
                            <p class="text-sm font-semibold uppercase tracking-[0.08em] text-slate-500">
                                Plan #{{ plan.id }}
                            </p>
                            <h2 class="mt-2 text-2xl font-bold text-slate-900">
                                {{ plan.student?.username ?? "Student" }}
                            </h2>
                            <p class="mt-2 text-sm text-slate-600">
                                {{ plan.student?.email }}
                            </p>
                            <p class="mt-2 text-sm text-slate-600">
                                Status: {{ CoursePlanStatusMapper[plan.status] }}
                            </p>
                        </div>

                        <div class="flex flex-col gap-3 sm:flex-row">
                            <button
                                type="button"
                                class="inline-flex min-h-11 items-center justify-center rounded-md bg-blue-900 px-4 py-2 font-semibold text-white transition hover:bg-blue-800 disabled:cursor-not-allowed disabled:opacity-75"
                                :disabled="pendingActions.has(actionKey(plan.id, 'approve-all'))"
                                @click="reviewPlan(plan.id, CoursePlanItemStatusEnum.APPROVED)"
                            >
                                Approve all
                            </button>
                            <button
                                type="button"
                                class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900 disabled:cursor-not-allowed disabled:opacity-75"
                                :disabled="pendingActions.has(actionKey(plan.id, 'reject-all'))"
                                @click="reviewPlan(plan.id, CoursePlanItemStatusEnum.REJECTED)"
                            >
                                Reject all
                            </button>
                        </div>
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

                            <div class="mt-4 flex gap-3">
                                <button
                                    type="button"
                                    class="inline-flex flex-1 min-h-11 items-center justify-center rounded-md bg-blue-900 px-4 py-2 font-semibold text-white transition hover:bg-blue-800 disabled:cursor-not-allowed disabled:opacity-75"
                                    :disabled="pendingActions.has(actionKey(plan.id, `approve-${item.courseId}`))"
                                    @click="reviewPlan(plan.id, CoursePlanItemStatusEnum.APPROVED, [item.courseId])"
                                >
                                    Approve
                                </button>
                                <button
                                    type="button"
                                    class="inline-flex flex-1 min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900 disabled:cursor-not-allowed disabled:opacity-75"
                                    :disabled="pendingActions.has(actionKey(plan.id, `reject-${item.courseId}`))"
                                    @click="reviewPlan(plan.id, CoursePlanItemStatusEnum.REJECTED, [item.courseId])"
                                >
                                    Reject
                                </button>
                            </div>
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
import { CoursePlanItemStatusEnum } from "~/constants/coursePlanItemStatusEnum"
import { extractErrorMessage } from "~/utils/auth"
import { CoursePlanStatusMapper } from "~/mapper/coursePlanStatusMapper"

definePageMeta({
    middleware: ["auth", "admin"],
})

const plans = ref<CoursePlan[]>([])
const isLoading = ref(true)
const errorMessage = ref("")
const pendingActions = ref(new Set<string>())

onMounted(async () => {
    await fetchPlans()
})

async function fetchPlans() {
    try {
        const response = await CoursePlanService.getPendingCoursePlans()
        plans.value = Array.isArray(response) ? response : []
    } catch (error) {
        errorMessage.value = extractErrorMessage(error, "Unable to load pending approvals right now.")
    } finally {
        isLoading.value = false
    }
}

function actionKey(planId: number, action: string) {
    return `${planId}:${action}`
}

async function reviewPlan(planId: number, itemStatus: CoursePlanItemStatusEnum, courseIds: number[] = []) {
    const action = courseIds.length === 0 ? `${itemStatus.toLowerCase()}-all` : `${itemStatus.toLowerCase()}-${courseIds.join("-")}`
    const key = actionKey(planId, action)

    if (pendingActions.value.has(key)) {
        return
    }

    pendingActions.value.add(key)
    pendingActions.value = new Set(pendingActions.value)

    try {
        await CoursePlanService.reviewCoursePlan(planId, {
            item_status: itemStatus,
            course_ids: courseIds,
        })

        await fetchPlans()
    } catch (error) {
        errorMessage.value = extractErrorMessage(error, "Unable to update this approval right now.")
    } finally {
        pendingActions.value.delete(key)
        pendingActions.value = new Set(pendingActions.value)
    }
}
</script>