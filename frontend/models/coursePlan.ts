import { BaseModel } from "./baseModel"
import { CoursePlanStatusEnum } from "../constants/coursePlanStatusEnum"

export interface CoursePlan extends BaseModel {
    id: number
    status: CoursePlanStatusEnum
}