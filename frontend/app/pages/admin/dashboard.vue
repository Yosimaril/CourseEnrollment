<template>
    <main class="min-h-screen bg-slate-50 px-4 py-6">
        <section class="mx-auto w-full max-w-6xl rounded-md bg-white p-6 shadow-md sm:p-8">
            <div class="flex flex-col gap-4 border-b border-slate-200 pb-6">
                <div>
                    <p class="text-sm font-bold uppercase tracking-[0.08em] text-slate-600">Admin dashboard</p>
                    <h1 class="mt-2 text-3xl font-bold text-slate-900 sm:text-4xl">Signed in as admin.</h1>
                    <p class="mt-4 max-w-3xl text-base leading-7 text-slate-600">Manage the course catalog, student accounts, and course plan approvals from one place.</p>
                </div>

                <div class="flex flex-col gap-3 sm:flex-row sm:flex-wrap">
                    <NuxtLink to="/admin/courses" class="inline-flex min-h-11 items-center justify-center rounded-md bg-blue-900 px-4 py-2 font-semibold text-white transition hover:bg-blue-800"> Manage courses </NuxtLink>
                    <NuxtLink to="/admin/students" class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"> Student list </NuxtLink>
                    <NuxtLink to="/admin/approvals" class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"> Approvals </NuxtLink>
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
                        <p class="text-sm font-semibold uppercase tracking-[0.08em] text-slate-500">Pending reviews</p>
                        <p class="mt-3 text-3xl font-bold text-slate-900">{{ pendingApprovalsCount }}</p>
                        <p class="mt-2 text-sm text-slate-600">Submitted plans waiting in the approval queue.</p>
                    </article>

                    <article class="rounded-md border border-slate-200 bg-slate-50 p-5">
                        <p class="text-sm font-semibold uppercase tracking-[0.08em] text-slate-500">Student plans</p>
                        <p class="mt-3 text-3xl font-bold text-slate-900">{{ totalPendingItems }}</p>
                        <p class="mt-2 text-sm text-slate-600">Total course items currently waiting for review.</p>
                    </article>

                    <article class="rounded-md border border-slate-200 bg-slate-50 p-5 sm:col-span-2 xl:col-span-1">
                        <p class="text-sm font-semibold uppercase tracking-[0.08em] text-slate-500">Latest queue</p>
                        <p class="mt-3 text-3xl font-bold text-slate-900">{{ latestApprovalLabel }}</p>
                        <p class="mt-2 text-sm text-slate-600">Most recent submitted plan currently visible from this dashboard.</p>
                    </article>
                </div>

                <aside class="rounded-md border border-slate-200 bg-slate-50 p-5">
                    <p class="text-sm font-bold uppercase tracking-[0.08em] text-slate-600">Quick actions</p>
                    <div class="mt-4 grid gap-3">
                        <NuxtLink to="/admin/approvals" class="inline-flex min-h-11 items-center justify-center rounded-md bg-blue-900 px-4 py-2 font-semibold text-white transition hover:bg-blue-800"> Review approvals </NuxtLink>
                        <NuxtLink to="/admin/courses" class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"> Update course catalog </NuxtLink>
                        <NuxtLink to="/admin/students" class="inline-flex min-h-11 items-center justify-center rounded-md border border-slate-200 bg-white px-4 py-2 font-semibold text-slate-900 transition hover:border-blue-900"> Check student records </NuxtLink>
                    </div>

                    <div class="mt-6 rounded-md border border-slate-200 bg-white p-4">
                        <p class="text-sm font-semibold uppercase tracking-[0.08em] text-slate-500">Recent submissions</p>
                        <ul class="mt-3 space-y-3 text-sm text-slate-700">
                            <li v-for="plan in recentPendingPlans" :key="plan.id" class="flex items-start justify-between gap-3">
                                <span>{{ plan.student?.username ?? `Student #${plan.studentId}` }}</span>
                                <span class="whitespace-nowrap text-slate-500">{{ plan.items.length }} items</span>
                            </li>
                            <li v-if="recentPendingPlans.length === 0" class="text-slate-500">No plans waiting for review.</li>
                        </ul>
                    </div>
                </aside>
            </div>
        </section>
    </main>
</template>

<script setup lang="ts">
import { CoursePlanService } from "~/services/coursePlan";
import type { CoursePlan } from "~/models/coursePlan";
import { clearStoredToken } from "~/utils/auth";
import { extractErrorMessage } from "~/utils/auth";

definePageMeta({
    middleware: ["auth", "admin"],
});

const pendingPlans = ref<CoursePlan[]>([]);
const isLoading = ref(true);
const errorMessage = ref("");

const pendingApprovalsCount = computed(() => pendingPlans.value.length);
const totalPendingItems = computed(() => pendingPlans.value.reduce((sum, plan) => sum + plan.items.length, 0));
const latestApprovalLabel = computed(() => {
    const latestPlan = pendingPlans.value[0];

    if (!latestPlan) {
        return "No pending plans";
    }

    return latestPlan.student?.username ?? `Student #${latestPlan.studentId}`;
});
const recentPendingPlans = computed(() => pendingPlans.value.slice(0, 4));

onMounted(async () => {
    try {
        pendingPlans.value = await CoursePlanService.getPendingCoursePlans();
    } catch (error) {
        errorMessage.value = extractErrorMessage(error, "Unable to load dashboard summary right now.");
    } finally {
        isLoading.value = false;
    }
});

async function logout() {
    clearStoredToken();
    await navigateTo("/auth/login");
}
</script>
