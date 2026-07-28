<template>
    <main class="landing-shell">
        <section class="hero-card">
            <p class="eyebrow">Course Enrollment</p>
            <h1>Sign in, register, and get routed straight into the right student or admin flow.</h1>
            <p class="lede">The frontend now talks to the Go auth endpoints, stores the JWT, and sends users to the correct dashboard based on the role in the token.</p>

            <div class="actions">
                <NuxtLink to="/auth/login" class="primary-action">Login</NuxtLink>
                <NuxtLink to="/auth/register" class="secondary-action">Register</NuxtLink>
            </div>
        </section>

        <aside class="status-card">
            <p class="status-label">Current session</p>
            <p class="status-value">{{ sessionState }}</p>
            <p class="status-note">If a token exists, you will be forwarded automatically.</p>
        </aside>
    </main>
</template>

<script setup lang="ts">
import { getSessionRedirectPath } from "~/utils/auth";

const sessionState = ref("No active session");

onMounted(async () => {
    const redirectPath = getSessionRedirectPath();

    if (redirectPath) {
        sessionState.value = `Redirecting to ${redirectPath}`;
        await navigateTo(redirectPath);
    }
});
</script>

<style scoped>
</style>
