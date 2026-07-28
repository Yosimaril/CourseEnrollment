export default defineNuxtPlugin(() => {
    const config = useRuntimeConfig()

    const api = $fetch.create({
        baseURL: config.public.apiUrl,

        onRequest({ options }) {
            if (process.client) {
                const token = localStorage.getItem("token")

                if (token) {
                    options.headers.set(
                        "Authorization",
                        `Bearer ${token}`
                    )
                }
            }
        },

        async onResponseError({ response }) {
            if (response.status === 401) {
                await navigateTo("/login")
            }
        }
    })

    return {
        provide: {
            api
        }
    }
})