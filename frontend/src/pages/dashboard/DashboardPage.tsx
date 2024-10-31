import { Typography } from "@mui/material";
import { UserType } from "../../enums/userType.enum";
import { useUser } from "../../hooks/useUser"
import { useState } from "react";

export const DashboardPage = () => {
    const { user } = useUser();
    console.log(user)
    const [isAdmin, setIsAdmin] = useState(() => {
        return user?.UserRoles.some(userRole => userRole.Role.ID === UserType.Admin)
    });

    return (
        <>
            {isAdmin && (
                <Typography>Admin dashboard</Typography>
            )}
            {!isAdmin && (
                <Typography>User dashboard</Typography>
            )}
        </>
    )
}