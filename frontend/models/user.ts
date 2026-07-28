import { UserRoleEnum } from "../constants/userRoleEnum"
import { BaseModel } from "./baseModel"

export interface User extends BaseModel {
    id: number
    username: string
    email: string
    role: UserRoleEnum
    nrp?: string
    maxCredits?: number
}