import type { User } from "~/models/user"

export type RawUser = {
    id?: number
    username?: string
    Username?: string
    email?: string
    role?: User["role"] | string
    nrp?: string
    maxCredits?: number
    max_credits?: number
    createdAt?: string
    created_at?: string
    updatedAt?: string
    updated_at?: string
}