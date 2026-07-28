const { $api } = useNuxtApp()

export interface LoginRequest {
    email: string
    password: string
}

export interface LoginResponse {
    token: string
}

export const AuthService = {

    async login(request: LoginRequest) {

        return await $api<LoginResponse>(
            "/login",
            {
                method: "POST",
                body: request
            }
        )
    },

    async register(request: any) {

        return await $api(
            "/register",
            {
                method: "POST",
                body: request
            }
        )
    }

}