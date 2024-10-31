import React from "react";
import { Card, CardContent, Typography } from "@mui/material";
import { User } from "../../models/user";
import { UsageStats } from "./usage-stats/UsageStats";
import { UserType } from "../../enums/userType.enum";

interface UserDashboardProps {
    user: User;
}

export const UserDashboard: React.FC<UserDashboardProps> = ({ user }) => {
    const { FirstName, UserRoles } = user;
    const maxRequestCount = UserRoles[0]?.Role?.MaxRequestCount || 20;

    return (
        <Card sx={{ maxWidth: 400, margin: "0 auto", mt: 5, p: 2 }}>
            <CardContent>
                <Typography variant="h5" component="div" gutterBottom>
                    Welcome back, {FirstName}
                </Typography>
                <UsageStats userId={user.ID} maxRequestCount={maxRequestCount} isAdmin={false} />
            </CardContent>
        </Card>
    );
};