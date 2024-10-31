import React from "react";
import { Card, CardContent, Typography } from "@mui/material";
import { User } from "../../models/user";
import { UsageStats } from "./usage-stats/UsageStats";

const styles = {
    card: {
        padding: 2,
        backgroundColor: '#424242', // Slightly lighter grey for contrast
        color: '#ffffff',
        borderRadius: 2,
        boxShadow: '0px 4px 20px rgba(0, 0, 0, 0.5)',
        marginBottom: 2
    },
};

interface UserDashboardProps {
    user: User;
}

export const UserDashboard: React.FC<UserDashboardProps> = ({ user }) => {
    const { FirstName, UserRoles } = user;
    const maxRequestCount = UserRoles[0]?.Role?.MaxRequestCount || 20;

    return (
        <Card sx={styles.card}>
            <CardContent>
                <Typography variant="h5" component="div" gutterBottom>
                    Welcome back, {FirstName}
                </Typography>
                <br />
                <UsageStats userId={user.ID} maxRequestCount={maxRequestCount} isAdmin={false} />
            </CardContent>
        </Card>
    );
};