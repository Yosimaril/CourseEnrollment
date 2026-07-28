<template>
    <main class="auth-shell">
        <section class="auth-card">
            <div class="hero-copy">
                <p class="eyebrow">Course Enrollment</p>
                <h1>Create your account.</h1>
            </div>

            <form class="auth-form" @submit.prevent="submitRegister">
                <label class="field">
                    <span>Username</span>
                    <input v-model.trim="form.username" type="text" placeholder="yosimaril" autocomplete="username" required />
                </label>

                <label class="field">
                    <span>Email</span>
                    <input v-model.trim="form.email" type="email" placeholder="yosimaril@example.com" autocomplete="email" required />
                </label>

                <label class="field">
                    <span>Password</span>
                    <input v-model="form.password" type="password" placeholder="Minimum 6 characters" autocomplete="new-password" required minlength="6" />
                </label>

                <label class="field">
                    <span>Role</span>
                    <select v-model="form.role" required>
                        <option :value="UserRoleEnum.STUDENT">Student</option>
                        <option :value="UserRoleEnum.ADMIN">Admin</option>
                    </select>
                </label>

                <template v-if="form.role === UserRoleEnum.STUDENT">
                    <label class="field">
                        <span>NRP</span>
                        <input v-model.trim="form.nrp" type="text" placeholder="160000001" autocomplete="off" />
                    </label>

                    <label class="field">
                        <span>Max credits</span>
                        <input v-model="form.maxCredits" type="number" min="0" step="1" placeholder="24" />
                    </label>
                </template>

                <p v-if="errorMessage" class="message error">{{ errorMessage }}</p>

                <button type="submit" class="submit-button" :disabled="isSubmitting">
                    {{ isSubmitting ? "Registering..." : "Register" }}
                </button>

                <p class="switch-link">
                    Already have an account?
                    <NuxtLink to="/auth/login">Login here</NuxtLink>
                </p>
            </form>
        </section>
    </main>
</template>

<script setup lang="ts">
import { UserRoleEnum } from "~/constants/userRoleEnum";
import { AuthService } from "~/services/auth";
import { extractErrorMessage, getDashboardPath, setStoredToken } from "~/utils/auth";

definePageMeta({
    middleware: "guest",
});

const form = reactive({
    username: "",
    email: "",
    password: "",
    role: UserRoleEnum.STUDENT,
    nrp: "",
    maxCredits: "",
});

const isSubmitting = ref(false);
const errorMessage = ref("");

async function submitRegister() {
    isSubmitting.value = true;
    errorMessage.value = "";

    const request = {
        username: form.username,
        email: form.email,
        password: form.password,
        role: form.role,
        ...(form.role === UserRoleEnum.STUDENT && form.nrp ? { nrp: form.nrp } : {}),
        ...(form.role === UserRoleEnum.STUDENT && form.maxCredits !== "" ? { max_credits: Number(form.maxCredits) } : {}),
    };

    try {
        const response = await AuthService.register(request);

        setStoredToken(response.token);
        await navigateTo(getDashboardPath(response.user.role));
    } catch (error) {
        errorMessage.value = extractErrorMessage(error, "Unable to register right now.");
    } finally {
        isSubmitting.value = false;
    }
}
</script>

<style scoped>
</style>
