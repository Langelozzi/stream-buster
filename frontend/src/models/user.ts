import { Usage } from "./usage";
import { UserRole } from "./user_role";

export interface User {
    ID: number;
    Email: string;
    FirstName: string;
    LastName: string;
    Usage: Usage;
    UserRoles: UserRole[];
    DeletedAt: string;
    CreatedAt: string;
}