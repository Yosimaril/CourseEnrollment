import type { CoursePlan } from "~/models/coursePlan"
import type { CoursePlanItemStatusEnum } from "~/constants/coursePlanItemStatusEnum"

const { $api } = useNuxtApp()

export const CoursePlanService = {
    async getMyCoursePlan(): Promise<CoursePlan> {
        return await $api("/student/course-plan")
    },

    async getMyCoursePlans(): Promise<CoursePlan[]> {
        return await $api("/student/course-plans")
    },

    async submitMyCoursePlan(): Promise<CoursePlan> {
        return await $api("/student/course-plan/submit", {
            method: "POST",
        })
    },

    async cancelMyCoursePlan(id: number) {
        return await $api(`/student/course-plans/${id}`, {
            method: "DELETE",
        })
    },

    async getPendingCoursePlans(): Promise<CoursePlan[]> {
        return await $api("/admin/course-plans?status=SUBMITTED")
    },

    async reviewCoursePlan(
        id: number,
        payload: {
            item_status: CoursePlanItemStatusEnum
            course_ids?: number[]
        }
    ) {
        return await $api(`/admin/course-plans/${id}/review`, {
            method: "PUT",
            body: payload,
        })
    },
}