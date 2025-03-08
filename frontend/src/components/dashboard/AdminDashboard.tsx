import React, { useEffect, useState } from "react";
import { Card, CardContent, Typography, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper } from "@mui/material";
import { User } from "../../models/user";  // Adjust the path as necessary
import { getAllUsers } from "../../api/services/user.service";
import { UserUsageInfo } from "../usage-stats/UserUsageInfo";
import { isAdminUser } from "../../utils/user.helpers";
import { useTranslation } from "react-i18next";

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
    const { t } = useTranslation();

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
                    <Typography variant="h5">{t('dictionary.welcomeBack', { name: user.FirstName })}</Typography>
                    <br />
                    <UserUsageInfo user={user} isAdmin={true} />
                </CardContent>
            </Card>

            {/* User Usage Card */}
            <Card sx={styles.card}>
                <CardContent>
                    <Typography variant="h6" gutterBottom>{t('dictionary.userUsage')}</Typography>
                    <TableContainer component={Paper}>
                        <Table aria-label="User Usage Table">
                            <TableHead>
                                <TableRow>
                                    <TableCell>{t('columnHeader.user')}</TableCell>
                                    <TableCell>{t('columnHeader.email')}</TableCell>
                                    <TableCell>{t('columnHeader.usage')}</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {users.filter(otherUser => otherUser.ID !== user.ID).map((user) => (
                                    <TableRow key={user.ID}>
                                        <TableCell sx={{ verticalAlign: 'top', width: '150px' }}>
                                            <Typography variant="body1">
                                                {`${user.FirstName} ${user.LastName}`}
                                            </Typography>
                                        </TableCell>
                                        <TableCell sx={{ verticalAlign: 'top', width: '150px' }}>
                                            <Typography variant="body1">
                                                {user.Email}
                                            </Typography>
                                        </TableCell>
                                        <TableCell>
                                            <UserUsageInfo user={user} isAdmin={isAdminUser(user)} />
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
