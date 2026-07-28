<template>
    <main class="dashboard-shell">
        <section class="dashboard-card">
            <p class="eyebrow">Student dashboard</p>
            <h1>Signed in as student.</h1>
            <p>Your auth token has been accepted and you can now continue to the student flow.</p>

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
.dashboard-shell {
    min-height: 100vh;
    display: grid;
    place-items: center;
    padding: 1rem;
    box-sizing: border-box;
}

.dashboard-card {
    width: min(100%, 760px);
    padding: clamp(1.5rem, 3vw, 2.5rem);
    border-radius: 30px;
    border: 1px solid var(--panel-border);
    background: var(--panel);
    backdrop-filter: blur(18px);
    box-shadow: var(--shadow);
}

.eyebrow {
    margin: 0 0 1rem;
    text-transform: uppercase;
    letter-spacing: 0.18em;
    font-size: 0.78rem;
    color: var(--accent-deep);
    font-weight: 700;
}

h1 {
    margin: 0;
    font-size: clamp(2rem, 4vw, 3.5rem);
}

p {
    color: var(--muted);
    line-height: 1.7;
}

.actions {
    display: flex;
    flex-wrap: wrap;
    gap: 0.9rem;
    margin-top: 1.5rem;
}

.primary-action,
.secondary-action {
    min-height: 48px;
    padding: 0 1.3rem;
    border-radius: 999px;
    border: 0;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    text-decoration: none;
    font-weight: 700;
    cursor: pointer;
}

.primary-action {
    color: white;
    background: linear-gradient(135deg, var(--accent), var(--accent-deep));
}

.secondary-action {
    border: 1px solid rgba(76, 54, 41, 0.16);
    background: rgba(255, 255, 255, 0.55);
}
</style>
