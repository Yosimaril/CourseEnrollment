<script setup lang="ts">
import Error401 from "~/components/error/Error401.vue";
import Error403 from "~/components/error/Error403.vue";
import Error404 from "~/components/error/Error404.vue";
import Error500 from "~/components/error/Error500.vue";

const error = useError();

const goHome = () => {
    clearError({ redirect: "/" });
};
</script>

<template>
    <Error404 v-if="error?.status === 404" @home="goHome" />
    <Error403 v-else-if="error?.status === 403" @home="goHome" />
    <Error401 v-else-if="error?.status === 401" @home="goHome" />
    <Error500 v-else @home="goHome" />
</template>
