import type { BaseModel } from "./baseModel";
import { UserRoleEnum } from "../constants/userRoleEnum";

export interface User extends BaseModel {
    id: number;
    username: string;
    email: string;
    role: UserRoleEnum;
    nrp?: string;
    maxCredits?: number;
}
