import { clearStoredToken, getSessionRedirectPath } from "~/utils/auth"

export default defineNuxtRouteMiddleware(() => {
	if (!process.client) {
		return
	}

	const redirectPath = getSessionRedirectPath()

	if (redirectPath) {
		return navigateTo(redirectPath)
	}

	const token = localStorage.getItem("token")

	if (token) {
		clearStoredToken()
	}
})
