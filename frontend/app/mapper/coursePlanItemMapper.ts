import type { RawCourse } from "~/mapper/courseMapper"

export type RawCoursePlanItem = {
    course_plan_id?: number
    course_id?: number
    status?: string
    remarks?: string | null
    created_at?: string
    updated_at?: string
    course?: RawCourse
}