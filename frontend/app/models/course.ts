import type { BaseId } from "./baseId";
import type { BaseTimestamp } from "./baseTimestamp";

export interface Course extends BaseId, BaseTimestamp {
    code: string;
    name: string;
    credits: number;
}
