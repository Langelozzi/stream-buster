import { UserEndpointUsage } from "./user_endpoint_usage";
import { UserRole } from "./user_role";

export interface User {
    ID: number;
    Email: string;
    FirstName: string;
    LastName: string;
    Usage: UserEndpointUsage;
    UserRoles: UserRole[];
    DeletedAt: string;
    CreatedAt: string;
}