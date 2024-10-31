import { Role } from "./role";
import { User } from "./user";

export interface UserRole {
    UserID: number;
    RoleID: number;
    Role: Role;
    User: User;
}