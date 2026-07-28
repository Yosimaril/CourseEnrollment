<template>
    <main class="dashboard-shell">
        <section class="dashboard-card">
            <p class="eyebrow">Admin dashboard</p>
            <h1>Signed in as admin.</h1>

            <div class="actions">
                <NuxtLink to="/admin/courses" class="secondary-action">Manage courses</NuxtLink>
                <button type="button" class="primary-action" @click="logout">Logout</button>
            </div>
        </section>
    </main>
</template>

<script setup lang="ts">
import { clearStoredToken } from "~/utils/auth";

definePageMeta({
    middleware: ["auth", "admin"],
});

async function logout() {
    clearStoredToken();
    await navigateTo("/auth/login");
}
</script>

<style scoped>
</style>
