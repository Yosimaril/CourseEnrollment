import type { BaseModel } from "./baseModel"
import { CoursePlanItemStatusEnum } from "../constants/coursePlanItemStatusEnum"
import type { Course } from "./course"

export interface CoursePlanItem extends BaseModel {
    coursePlanId: number
    courseId: number
    status: CoursePlanItemStatusEnum
    remarks?: string
    course?: Course
}