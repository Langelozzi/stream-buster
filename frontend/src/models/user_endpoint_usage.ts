import { Endpoint } from "./endpoint";
import { User } from "./user";

export interface UserEndpointUsage {
    ID: number;
    UserID: number;
    EndpointID: number;
    RequestCount: number;
    LastAccess: string;
    CreatedAt: string;
    UpdatedAt: string;

    User: User;
    Endpoint: Endpoint;
}
