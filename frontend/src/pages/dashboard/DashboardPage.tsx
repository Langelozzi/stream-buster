import { UserType } from "../../enums/userType.enum";
import { useUser } from "../../hooks/useUser"
import { useState } from "react";
import { UserDashboard } from "../../components/dashboard/UserDashboard";
import { AdminDashboard } from "../../components/dashboard/AdminDashboard";

export const DashboardPage = () => {
    const { user } = useUser();
    console.log(user)
    const [isAdmin] = useState(() => {
        return user?.UserRoles.some(userRole => userRole.Role.ID === UserType.Admin)
    });

    // Make two dashboards, one for user and one for admin
    return (
        <>
            {isAdmin && user && (
                <AdminDashboard user={user} />
            )}
            {!isAdmin && user && (
                <UserDashboard user={user} />
            )}
        </>
    )
}