export const CourseService = {
    async getCourses() {
        const { $api } = useNuxtApp();
        return await $api("/courses");
    },

    async getCourse(id: number) {
        const { $api } = useNuxtApp();
        return await $api(`/courses/${id}`);
    },

    async createCourse(course: any) {
        const { $api } = useNuxtApp();
        return await $api("/admin/courses", {
            method: "POST",
            body: course,
        });
    },

    async updateCourse(id: number, course: any) {
        const { $api } = useNuxtApp();
        return await $api(`/admin/courses/${id}`, {
            method: "PUT",
            body: course,
        });
    },

    async deleteCourse(id: number) {
        const { $api } = useNuxtApp();
        return await $api(`/admin/courses/${id}`, {
            method: "DELETE",
        });
    },
};
