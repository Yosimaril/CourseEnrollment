const { $api } = useNuxtApp()

export const CoursePlanItemService = {
	async addToPickedCourses(courseId: number) {
		return await $api("/student/picked-courses", {
			method: "POST",
			body: {
				course_id: courseId,
			},
		})
	},

	async removeFromPickedCourses(courseId: number) {
		return await $api(`/student/course-plan-items/${courseId}`, {
			method: "DELETE",
		})
	},
}
