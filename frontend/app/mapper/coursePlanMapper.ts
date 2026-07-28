import type { RawUser } from "~/mapper/userMapper"
import type { RawCoursePlanItem } from "./coursePlanItemMapper"

export type RawCoursePlan = {
    id?: number
    studentId?: number
    student_id?: number
    status?: string
    items?: RawCoursePlanItem[]
    Items?: RawCoursePlanItem[]
    student?: RawUser
    Student?: RawUser
    createdAt?: string
    created_at?: string
    updatedAt?: string
    updated_at?: string
}