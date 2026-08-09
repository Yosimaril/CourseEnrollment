import type { BaseIdDto } from "~/dto/baseIdDto";
import type { BaseTimestampDto } from "~/dto/baseTimestampDto";

export type RawStudent = BaseIdDto &
    BaseTimestampDto & {
        username: string;
        email: string;
        nrp: string;
        max_credits: number;
    };

export class StudentMapper {
    static fromJson(json: unknown): RawStudent {
        return json as RawStudent;
    }

    static toJson(dto: RawStudent): object {
        return {
            id: dto.id,
            username: dto.username,
            email: dto.email,
            nrp: dto.nrp,
            max_credits: dto.max_credits,
            created_at: dto.created_at,
            updated_at: dto.updated_at,
        };
    }
}
