import React, { useEffect, useState } from "react";
import { Card, CardContent, Typography, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper, List } from "@mui/material";
import { User } from "../../models/user";  // Adjust the path as necessary
import { getAllUsers } from "../../api/services/user.service";
import { UserUsageInfo } from "../usage-stats/UserUsageInfo";

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

interface AdminDashboardProps {
    user: User;
}

export const AdminDashboard: React.FC<AdminDashboardProps> = ({ user }) => {
    const [users, setUsers] = useState<User[]>([]);

    useEffect(() => {
        const fetchUsers = async () => {
            const allUsers = await getAllUsers(true);
            setUsers(allUsers);
        };
        fetchUsers();
    }, []);

    return (
        <>
            {/* Welcome Card */}
            <Card sx={styles.card}>
                <CardContent>
                    <Typography variant="h5">Welcome back, {user.FirstName}</Typography>
                    <br />
                    <UserUsageInfo user={user} isAdmin={true} />
                </CardContent>
            </Card>

            {/* User Usage Card */}
            <Card sx={styles.card}>
                <CardContent>
                    <Typography variant="h6" gutterBottom>User Usage</Typography>
                    <List>
                        {users.filter(otherUser => otherUser.ID !== user.ID).map((user) => (
                            <UserUsageInfo user={user} isAdmin />
                        ))}
                    </List>
                </CardContent>
            </Card>
        </>
    );
};
