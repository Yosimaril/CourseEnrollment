import type { RawCourse } from "~/mapper/courseMapper"

export type RawCoursePlanItem = {
    coursePlanId?: number
    course_plan_id?: number
    CoursePlanID?: number
    courseId?: number
    course_id?: number
    CourseID?: number
    status?: string
    remarks?: string | null
    createdAt?: string
    created_at?: string
    updatedAt?: string
    updated_at?: string
    course?: RawCourse
    Course?: RawCourse
}