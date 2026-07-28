import { BaseModel } from "./baseModel"
import { CoursePlanItemStatusEnum } from "../constants/coursePlanItemStatusEnum"

export interface CoursePlanItem extends BaseModel {
    id: number
    courseId: number
    status: CoursePlanItemStatusEnum
    remarks?: string
}