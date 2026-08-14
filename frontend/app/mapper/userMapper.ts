import type { User } from "~/models/user";
import type { UserDto } from "~/dto/userDto";

export class UserMapper {
    static toModel(dto: UserDto): User {
        return {
            id: dto.id,
            username: dto.username,
            email: dto.email,
            role: dto.role,
            nrp: dto.nrp,
            maxCredits: dto.max_credits,
            createdAt: new Date(dto.created_at),
            updatedAt: new Date(dto.updated_at),
            deletedAt: dto.deleted_at ? new Date(dto.deleted_at) : null,
        };
    }

    static toDto(model: User): UserDto {
        return {
            id: model.id,
            username: model.username,
            email: model.email,
            role: model.role,
            nrp: model.nrp,
            max_credits: model.maxCredits,
            created_at: model.createdAt.toISOString(),
            updated_at: model.updatedAt.toISOString(),
            deleted_at: model.deletedAt ? model.deletedAt.toISOString() : null,
        };
    }
}
