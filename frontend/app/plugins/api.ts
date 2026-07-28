export default defineNuxtPlugin(() => {
    const config = useRuntimeConfig()

    const api = $fetch.create({
        baseURL: config.public.apiUrl,

        onRequest({ options }) {
            if (process.client) {
                const token = localStorage.getItem("token")

                if (token) {
                    options.headers.set("Authorization", `Bearer ${token}`)
                }
            }

            options.headers.set("Content-Type", "application/json")
            options.headers.set("Accept", "application/json")
        },

        async onResponse({ response }) {
            console.log(
                `[${response.status}] ${response.url}`
            )
        },

        async onResponseError({ response }) {
            switch (response.status) {
                case 401:
                    localStorage.removeItem("token")
                    await navigateTo("/auth/login")
                    break

                case 403:
                    alert("Forbidden")
                    break

                case 500:
                    alert("Internal Server Error")
                    break
            }
        }
    })

    return {
        provide: {
            api
        }
    }
})