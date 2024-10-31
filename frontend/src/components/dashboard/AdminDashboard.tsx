import React, { useEffect, useState } from "react";
import { Box, Card, CardContent, Typography, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper } from "@mui/material";
import { User } from "../../models/user";  // Adjust the path as necessary
import { UsageStats } from "./usage-stats/UsageStats";  // Adjust the path as necessary
import { getAllUsers } from "../../api/services/user.service";
import { UserType } from "../../enums/userType.enum";

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
        <Box p={2}>
            {/* Welcome Card */}
            <Card sx={{ mb: 2 }}>
                <CardContent>
                    <Typography variant="h5">Welcome back, {user.FirstName}</Typography>
                    <UsageStats
                        userId={user.ID}
                        maxRequestCount={user.UserRoles[0]?.Role.MaxRequestCount}
                        isAdmin={true}
                    />
                </CardContent>
            </Card>

            {/* User Usage Card */}
            <Card>
                <CardContent>
                    <Typography variant="h6" gutterBottom>User Usage</Typography>
                    <TableContainer component={Paper}>
                        <Table aria-label="User Usage Table">
                            <TableHead>
                                <TableRow>
                                    <TableCell>User</TableCell>
                                    <TableCell align="center">Request Usage</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {users.filter(otherUser => otherUser.ID !== user.ID).map((user) => (
                                    <TableRow key={user.ID}>
                                        <TableCell>
                                            {user.FirstName} {user.LastName}
                                        </TableCell>
                                        <TableCell align="center">
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
        </Box>
    );
};
