import { UserType } from "../../enums/userType.enum";
import { useUser } from "../../hooks/useUser";
import { useState } from "react";
import { UserDashboard } from "../../components/dashboard/UserDashboard";
import { AdminDashboard } from "../../components/dashboard/AdminDashboard";
import { Box, Grid2 } from "@mui/material";
import { AskQuery } from "../../components/chat/AskQuery";

export const DashboardPage = () => {
    const { user } = useUser();
    const [isAdmin] = useState(() => {
        return user?.UserRoles.some(userRole => userRole.Role.ID === UserType.Admin);
    });

    return (
        <Box p={2} sx={{ display: 'flex', height: '100vh' }}>
            <Grid2 container spacing={2} sx={{ flexGrow: 1 }}>
                <Grid2 size={8} sx={{ height: '100%' }}>
                    {isAdmin && user && <AdminDashboard user={user} />}
                    {!isAdmin && user && <UserDashboard user={user} />}
                </Grid2>
                <Grid2 size={4} sx={{ height: '100%' }}>
                    <AskQuery />
                </Grid2>
            </Grid2>
        </Box>
    );
};
