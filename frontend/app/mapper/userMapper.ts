import type { User } from "~/models/user";
import type { BaseIdDto } from "~/dto/baseIdDto";
import type { BaseTimestampDto } from "~/dto/baseTimestampDto";

export type RawUser = BaseIdDto &
    BaseTimestampDto & {
        username: string;
        email: string;
        role: User["role"] | string;
        nrp: string;
        max_credits?: number;
    };

export class UserMapper {
    static fromJson(json: Record<string, unknown>): RawUser {
        return {
            id: json.id as number,
            username: json.username as string,
            email: json.email as string,
            role: json.role as User["role"] | string,
            nrp: json.nrp as string,
            max_credits: json.max_credits as number,
            created_at: json.created_at as string,
            updated_at: json.updated_at as string,
            deleted_at: json.deleted_at as string,
        }
    }

    static toJson(dto: RawUser): object {
        return {
            id: dto.id,
            username: dto.username,
            email: dto.email,
            role: dto.role,
            nrp: dto.nrp,
            max_credits: dto.max_credits,
            created_at: dto.created_at,
            updated_at: dto.updated_at,
        };
    }
}
