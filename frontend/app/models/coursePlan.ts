import { BaseModel } from "./baseModel"
import { CoursePlanStatusEnum } from "../constants/coursePlanStatusEnum"
import type { CoursePlanItem } from "./coursePlanItem"
import type { User } from "./user"

export interface CoursePlan extends BaseModel {
    id: number
    studentId: number
    status: CoursePlanStatusEnum
    student?: User
    items: CoursePlanItem[]
}