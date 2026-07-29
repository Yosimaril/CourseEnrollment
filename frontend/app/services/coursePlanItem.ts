
export const CoursePlanItemService = {
	async addToPickedCourses(courseId: number) {
		const { $api } = useNuxtApp()
		return await $api("/student/picked-courses", {
			method: "POST",
			body: {
				course_id: courseId,
			},
		})
	},

	async removeFromPickedCourses(courseId: number) {
		const { $api } = useNuxtApp()
		return await $api(`/student/course-plan-items/${courseId}`, {
			method: "DELETE",
		})
	},
}
