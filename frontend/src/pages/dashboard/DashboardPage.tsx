import { UserType } from "../../enums/userType.enum";
import { useUser } from "../../hooks/useUser"
import { useState } from "react";
import { UserDashboard } from "../../components/dashboard/UserDashboard";
import { AdminDashboard } from "../../components/dashboard/AdminDashboard";
import { Box } from "@mui/material";

export const DashboardPage = () => {
    const { user } = useUser();
    const [isAdmin] = useState(() => {
        return user?.UserRoles.some(userRole => userRole.Role.ID === UserType.Admin)
    });

    // Make two dashboards, one for user and one for admin
    return (
        <Box p={2}>
            {isAdmin && user && (
                <AdminDashboard user={user} />
            )}
            {!isAdmin && user && (
                <UserDashboard user={user} />
            )}
        </Box>
    )
}