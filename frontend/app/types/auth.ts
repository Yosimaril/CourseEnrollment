import type { User } from "../models/user"
import type { UserRoleEnum } from "../constants/userRoleEnum"

export interface LoginRequest {
    email: string
    password: string
}

export interface LoginResponse {
    message: string
    token: string
}

export interface RegisterRequest {
    username: string
    email: string
    password: string
    role: UserRoleEnum
    nrp?: string
    max_credits?: number
}

export interface RegisterResponse {
    message: string
    token: string
    user: User
}