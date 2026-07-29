import type { BaseModel } from "./baseModel";
import { UserRoleEnum } from "../constants/userRoleEnum";

export interface Student extends BaseModel {
    id: number;
    username: string;
    email: string;
    role: UserRoleEnum;
    nrp?: string;
    maxCredits?: number;
}

export interface StudentCreate {
    username: string;
    email: string;
    password: string;
    role: UserRoleEnum;
    nrp?: string;
    max_credits?: number;
}

export interface StudentUpdate {
    username?: string;
    email?: string;
    role?: UserRoleEnum;
    nrp?: string;
    max_credits?: number;
}
