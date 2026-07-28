import { clearStoredToken, getSessionRedirectPath } from "~/utils/auth"

export default defineNuxtRouteMiddleware(() => {
	if (!process.client) {
		return
	}

	const redirectPath = getSessionRedirectPath()

	if (!redirectPath) {
		clearStoredToken()
		return navigateTo("/auth/login")
	}
})
