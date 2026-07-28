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
}
