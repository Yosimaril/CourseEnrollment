<template>
    <main class="dashboard-shell">
        <section class="dashboard-card">
            <p class="eyebrow">Student dashboard</p>
            <h1>Signed in as student.</h1>

            <div class="actions">
                <NuxtLink to="/student/courses" class="secondary-action">Browse courses</NuxtLink>
                <button type="button" class="primary-action" @click="logout">Logout</button>
            </div>
        </section>
    </main>
</template>

<script setup lang="ts">
import { clearStoredToken } from "~/utils/auth";

definePageMeta({
    middleware: ["auth", "student"],
});

async function logout() {
    clearStoredToken();
    await navigateTo("/auth/login");
}
</script>

<style scoped>
</style>
