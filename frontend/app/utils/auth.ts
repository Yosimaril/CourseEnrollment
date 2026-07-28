import { UserRoleEnum } from "~/constants/userRoleEnum"

const TOKEN_KEY = "token"

type JwtPayload = {
    user_role?: string
    role?: string
    userRole?: string
}

export function getStoredToken() {
    if (!process.client) {
        return null
    }

    return localStorage.getItem(TOKEN_KEY)
}

export function setStoredToken(token: string) {
    if (!process.client) {
        return
    }

    localStorage.setItem(TOKEN_KEY, token)
}

export function clearStoredToken() {
    if (!process.client) {
        return
    }

    localStorage.removeItem(TOKEN_KEY)
}

function decodeBase64Url(value: string) {
    const normalized = value.replace(/-/g, "+").replace(/_/g, "/")
    const padded = normalized.padEnd(Math.ceil(normalized.length / 4) * 4, "=")

    if (typeof atob === "function") {
        return atob(padded)
    }

    if (typeof Buffer !== "undefined") {
        return Buffer.from(padded, "base64").toString("utf8")
    }

    throw new Error("Unable to decode token payload")
}

export function getRoleFromToken(token: string | null | undefined) {
    if (!token) {
        return null
    }

    try {
        const [, payload] = token.split(".")

        if (!payload) {
            return null
        }

        const decodedPayload = JSON.parse(decodeBase64Url(payload)) as JwtPayload
        const role = decodedPayload.user_role ?? decodedPayload.role ?? decodedPayload.userRole

        if (role === UserRoleEnum.ADMIN || role === UserRoleEnum.STUDENT) {
            return role
        }

        return null
    } catch {
        return null
    }
}

export function getDashboardPath(role: UserRoleEnum | null | undefined) {
    if (role === UserRoleEnum.ADMIN) {
        return "/admin/dashboard"
    }

    if (role === UserRoleEnum.STUDENT) {
        return "/student/dashboard"
    }

    return "/auth/login"
}

export function getSessionRedirectPath() {
    const token = getStoredToken()

    if (!token) {
        return null
    }

    const role = getRoleFromToken(token)

    if (!role) {
        return null
    }

    return getDashboardPath(role)
}

export function extractErrorMessage(error: unknown, fallbackMessage: string) {
    const fetchError = error as {
        data?: { message?: string }
        statusMessage?: string
        message?: string
    }

    return fetchError?.data?.message ?? fetchError?.statusMessage ?? fetchError?.message ?? fallbackMessage
}