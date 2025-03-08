import { UserType } from "../enums/userType.enum";
import { User } from "../models/user";

export const isAdminUser = (user: User): boolean => {
    return user.UserRoles.some(userRole => userRole.Role.ID === UserType.Admin)
}