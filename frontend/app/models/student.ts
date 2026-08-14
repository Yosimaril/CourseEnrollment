import type { BaseId } from "./baseId";
import type { BaseTimestamp } from "./baseTimestamp";
import type { CoursePlan } from "./coursePlan";

export interface Student extends BaseId, BaseTimestamp {
    username: string;
    email?: string;
    password?: string | null;
    nrp: string;
    maxCredits?: number;
    coursePlans?: CoursePlan[];
}
