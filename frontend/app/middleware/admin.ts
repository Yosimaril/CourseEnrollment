import { UserRoleEnum } from "~/constants/userRoleEnum"
import { getRoleFromToken, getSessionRedirectPath } from "~/utils/auth"

export default defineNuxtRouteMiddleware(() => {
	if (!process.client) {
		return
	}

	const token = localStorage.getItem("token")
	const role = getRoleFromToken(token)

	if (!token || !role) {
		return navigateTo("/auth/login")
	}

	if (role !== UserRoleEnum.ADMIN) {
		return navigateTo(getSessionRedirectPath() ?? "/auth/login")
	}
})
