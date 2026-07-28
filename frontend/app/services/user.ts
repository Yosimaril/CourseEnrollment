const { $api } = useNuxtApp()

export const UserService = {
    async getUsers() {
        return await $api("/users")
    },

    async getUser(id: number) {
        return await $api(`/users/${id}`)
    },

    async createUser(user: any) {
        return await $api("/admin/user", {
            method: "POST",
            body: user
        })
    },

    async updateUser(id: number, user: any) {
        return await $api(`/admin/user/${id}`, {
            method: "PUT",
            body: user
        })
    },

    async deleteUser(id: number) {
        return await $api(`/admin/user/${id}`, {
            method: "DELETE"
        })
    }
}