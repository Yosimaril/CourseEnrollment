import type {
    LoginRequest,
    LoginResponse,
    RegisterRequest,
    RegisterResponse
} from "~/types/auth"

export class AuthService {

    static async login(request: LoginRequest) {
        const { $api } = useNuxtApp()

        return await $api<LoginResponse>(
            "/auth/login",
            {
                method: "POST",
                body: request
            }
        )
    }

    static async register(request: RegisterRequest) {
        const { $api } = useNuxtApp()

        return await $api<RegisterResponse>(
            "/auth/register",
            {
                method: "POST",
                body: request
            }
        )
    }

}