<template>
    <main class="min-h-screen bg-slate-50 px-4 py-6">
        <section class="mx-auto w-full max-w-6xl rounded-md bg-white p-6 shadow-md sm:p-8">
            <div class="flex flex-col gap-4 border-b border-slate-200 pb-6">
                <div>
                    <p class="text-sm font-bold uppercase tracking-[0.08em] text-slate-600">Student dashboard</p>
                    <h1 class="mt-2 text-3xl font-bold text-slate-900 sm:text-4xl">Signed in as student.</h1>
                    <p class="mt-4 max-w-3xl text-base leading-7 text-slate-600">Browse the course catalog, manage your basket, and track approval history from one overview.</p>
                </div>

                <div class="flex flex-col gap-3 sm:flex-row sm:flex-wrap">
                    <NuxtLink to="/student/courses" class="inline-flex min-h-11 items-center justify-center rounded-md bg-blue-900 px-4 py-2 font-semibold text-white transition hover:bg-blue-800"> Browse courses </NuxtLink>
                    <NuxtLink to="/student/course-plan" class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"> My course plan </NuxtLink>
                    <NuxtLink to="/student/approval-status" class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"> Approval history </NuxtLink>
                    <button type="button" class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900" @click="logout">Logout</button>
                </div>
            </div>

            <p v-if="isLoading" class="mt-8 text-slate-600">Loading dashboard summary...</p>
            <p v-else-if="errorMessage" class="mt-8 rounded-md border border-red-200 bg-red-50 px-4 py-3 text-red-700">
                {{ errorMessage }}
            </p>

            <div v-else class="mt-8 grid gap-4 lg:grid-cols-[minmax(0,1fr)_20rem]">
                <div class="grid gap-4 sm:grid-cols-2 xl:grid-cols-3">
                    <article class="rounded-md border border-slate-200 bg-slate-50 p-5">
                        <p class="text-sm font-semibold uppercase tracking-[0.08em] text-slate-500">Basket items</p>
                        <p class="mt-3 text-3xl font-bold text-slate-900">{{ draftItemsCount }}</p>
                        <p class="mt-2 text-sm text-slate-600">Courses currently in your draft basket.</p>
                    </article>

                    <article class="rounded-md border border-slate-200 bg-slate-50 p-5">
                        <p class="text-sm font-semibold uppercase tracking-[0.08em] text-slate-500">Draft credits</p>
                        <p class="mt-3 text-3xl font-bold text-slate-900">{{ draftCredits }}</p>
                        <p class="mt-2 text-sm text-slate-600">Total credits in your current course plan.</p>
                    </article>

                    <article class="rounded-md border border-slate-200 bg-slate-50 p-5 sm:col-span-2 xl:col-span-1">
                        <p class="text-sm font-semibold uppercase tracking-[0.08em] text-slate-500">Latest submission</p>
                        <p class="mt-3 break-all text-3xl font-bold text-slate-900">{{ latestSubmissionStatus }}</p>
                        <p class="mt-2 text-sm text-slate-600">Most recent submitted course plan or draft state.</p>
                    </article>
                </div>

                <aside class="rounded-md border border-slate-200 bg-slate-50 p-5">
                    <p class="text-sm font-bold uppercase tracking-[0.08em] text-slate-600">Quick actions</p>
                    <div class="mt-4 grid gap-3">
                        <NuxtLink to="/student/courses" class="inline-flex min-h-11 items-center justify-center rounded-md bg-blue-900 px-4 py-2 font-semibold text-white transition hover:bg-blue-800"> Add more courses </NuxtLink>
                        <NuxtLink to="/student/course-plan" class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"> Review basket </NuxtLink>
                        <NuxtLink to="/student/approval-status" class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"> Check approvals </NuxtLink>
                    </div>

                    <div class="mt-6 rounded-md border border-slate-200 bg-white p-4">
                        <p class="text-sm font-semibold uppercase tracking-[0.08em] text-slate-500">Recent plans</p>
                        <ul class="mt-3 space-y-3 text-sm text-slate-700">
                            <li v-for="plan in recentPlans" :key="plan.id" class="flex items-start justify-between gap-3">
                                <span>{{ CoursePlanStatusMapper[plan.status] }}</span>
                                <span class="whitespace-nowrap text-slate-500">{{ plan.items.length }} items</span>
                            </li>
                            <li v-if="recentPlans.length === 0" class="text-slate-500">No plan history yet.</li>
                        </ul>
                    </div>
                </aside>
            </div>
        </section>
    </main>
</template>

<script setup lang="ts">
import { CoursePlanStatusEnum } from "~/constants/coursePlanStatusEnum";
import { CoursePlanService } from "~/services/coursePlan";
import type { CoursePlan } from "~/models/coursePlan";
import { clearStoredToken, extractErrorMessage } from "~/utils/auth";
import { CoursePlanStatusMapper } from "~/mapper/frontend/coursePlanStatusMapper";

definePageMeta({
    middleware: ["auth", "student"],
});

const currentPlan = ref<CoursePlan | null>(null);
const plans = ref<CoursePlan[]>([]);
const isLoading = ref(true);
const errorMessage = ref("");

const draftItemsCount = computed(() => currentPlan.value?.items.length ?? 0);
const draftCredits = computed(() => currentPlan.value?.items.reduce((sum, item) => sum + (item.course?.credits ?? 0), 0) ?? 0);
const latestSubmissionStatus = computed(() => {
    const latestSubmittedPlan = plans.value.find((plan) => plan.status !== CoursePlanStatusEnum.DRAFT) ?? plans.value[0];

    if (!latestSubmittedPlan) {
        return currentPlan.value ? CoursePlanStatusMapper[currentPlan.value.status] : "No submissions yet";
    }

    return CoursePlanStatusMapper[latestSubmittedPlan.status];
});
const recentPlans = computed(() => plans.value.slice(0, 4));

onMounted(async () => {
    const [draftResult, historyResult] = await Promise.allSettled([CoursePlanService.getMyCoursePlan(), CoursePlanService.getMyCoursePlans()]);

    if (draftResult.status === "fulfilled") {
        currentPlan.value = draftResult.value;
    }

    if (historyResult.status === "fulfilled") {
        plans.value = historyResult.value;
    }

    if (draftResult.status === "rejected" && historyResult.status === "rejected") {
        errorMessage.value = extractErrorMessage(draftResult.reason, "Unable to load dashboard summary right now.");
    }

    isLoading.value = false;
});

async function logout() {
    clearStoredToken();
    await navigateTo("/auth/login");
}
</script>
