export const UserService = {
    async getUsers() {
        const { $api } = useNuxtApp();
        return await $api("/users");
    },

    async getUser(id: number) {
        const { $api } = useNuxtApp();
        return await $api(`/users/${id}`);
    },

    async createUser(user: any) {
        const { $api } = useNuxtApp();
        return await $api("/admin/user", {
            method: "POST",
            body: user,
        });
    },

    async updateUser(id: number, user: any) {
        const { $api } = useNuxtApp();
        return await $api(`/admin/user/${id}`, {
            method: "PUT",
            body: user,
        });
    },

    async deleteUser(id: number) {
        const { $api } = useNuxtApp();
        return await $api(`/admin/user/${id}`, {
            method: "DELETE",
        });
    },
};
