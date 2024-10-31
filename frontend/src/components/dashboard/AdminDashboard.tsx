import React, { useEffect, useState } from "react";
import { Card, CardContent, Typography, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper } from "@mui/material";
import { User } from "../../models/user";  // Adjust the path as necessary
import { UsageStats } from "./usage-stats/UsageStats";  // Adjust the path as necessary
import { getAllUsers } from "../../api/services/user.service";
import { UserType } from "../../enums/userType.enum";

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
        // Placeholder for fetching users from the API
        // Replace with actual fetch logic
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
                    <UsageStats
                        userId={user.ID}
                        maxRequestCount={user.UserRoles[0]?.Role.MaxRequestCount}
                        isAdmin={true}
                    />
                </CardContent>
            </Card>

            {/* User Usage Card */}
            <Card sx={styles.card}>
                <CardContent>
                    <Typography variant="h6" gutterBottom>User Usage</Typography>
                    <TableContainer component={Paper}>
                        <Table aria-label="User Usage Table">
                            <TableHead>
                                <TableRow>
                                    <TableCell>User</TableCell>
                                    <TableCell>Request Usage</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {users.filter(otherUser => otherUser.ID !== user.ID).map((user) => (
                                    <TableRow key={user.ID}>
                                        <TableCell>
                                            {user.FirstName} {user.LastName}
                                        </TableCell>
                                        <TableCell>
                                            <UsageStats
                                                userId={user.ID}
                                                maxRequestCount={user.UserRoles[0]?.Role.MaxRequestCount}
                                                isAdmin={user.UserRoles[0]?.RoleID === UserType.Admin}
                                            />
                                        </TableCell>
                                    </TableRow>
                                ))}
                            </TableBody>
                        </Table>
                    </TableContainer>
                </CardContent>
            </Card>
        </>
    );
};
