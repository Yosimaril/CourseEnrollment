import type { BaseId } from "./baseId";
import type { BaseTimestamp } from "./baseTimestamp";
import { CoursePlanStatusEnum } from "../constants/coursePlanStatusEnum";
import type { CoursePlanItem } from "./coursePlanItem";
import type { User } from "./user";

export interface CoursePlan extends BaseId, BaseTimestamp {
    studentId: number;
    status: CoursePlanStatusEnum;
    student?: User;
    items: CoursePlanItem[];
}
