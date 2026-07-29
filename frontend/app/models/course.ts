import type { BaseModel } from "./baseModel.ts"

export interface Course extends BaseModel {
    id: number
    code: string
    name: string
    credits: number
}