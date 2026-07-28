const { $api } = useNuxtApp()

export const CourseService = {
    async getCourses() {
        return await $api("/courses")
    },

    async getCourse(id: number) {
        return await $api(`/courses/${id}`)
    },

    async createCourse(course: any) {
        return await $api("/admin/courses", {
            method: "POST",
            body: course
        })
    },

    async updateCourse(id: number, course: any) {
        return await $api(`/admin/courses/${id}`, {
            method: "PUT",
            body: course
        })
    },

    async deleteCourse(id: number) {
        return await $api(`/admin/courses/${id}`, {
            method: "DELETE"
        })
    }
}