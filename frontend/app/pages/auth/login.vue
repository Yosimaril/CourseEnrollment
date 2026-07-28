<template>
    <main class="auth-shell">
        <section class="auth-card">
            <div class="hero-copy">
                <p class="eyebrow">Course Enrollment</p>
                <h1>Welcome back.</h1>
            </div>

            <form class="auth-form" @submit.prevent="submitLogin">
                <label class="field">
                    <span>Email</span>
                    <input v-model.trim="form.email" type="email" placeholder="yosimaril@example.com" autocomplete="email" required />
                </label>

                <label class="field">
                    <span>Password</span>
                    <input v-model="form.password" type="password" placeholder="Enter your password" autocomplete="current-password" required minlength="6" />
                </label>

                <p v-if="errorMessage" class="message error">{{ errorMessage }}</p>

                <button type="submit" class="submit-button" :disabled="isSubmitting">
                    {{ isSubmitting ? "Logging in..." : "Login" }}
                </button>

                <p class="switch-link">
                    No account yet?
                    <NuxtLink to="/auth/register">Register here</NuxtLink>
                </p>
            </form>
        </section>
    </main>
</template>

<script setup lang="ts">
import { AuthService } from "~/services/auth";
import { extractErrorMessage, getDashboardPath, getRoleFromToken, setStoredToken } from "~/utils/auth";

definePageMeta({
    middleware: "guest",
});

const form = reactive({
    email: "",
    password: "",
});

const isSubmitting = ref(false);
const errorMessage = ref("");

async function submitLogin() {
    isSubmitting.value = true;
    errorMessage.value = "";

    try {
        const response = await AuthService.login({
            email: form.email,
            password: form.password,
        });

        setStoredToken(response.token);
        await navigateTo(getDashboardPath(getRoleFromToken(response.token)));
    } catch (error) {
        errorMessage.value = extractErrorMessage(error, "Unable to log in right now.");
    } finally {
        isSubmitting.value = false;
    }
}
</script>

<style scoped>
</style>
