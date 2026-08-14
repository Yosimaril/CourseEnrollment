import type { BaseId } from "./baseId";
import type { BaseTimestamp } from "./baseTimestamp";
import { UserRoleEnum } from "../constants/userRoleEnum";

export interface User extends BaseId, BaseTimestamp {
    username: string;
    email: string;
    role: UserRoleEnum;
    nrp?: string | null;
    maxCredits?: number | null;
}
