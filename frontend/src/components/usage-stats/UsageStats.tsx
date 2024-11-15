import { Box, Typography, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, TablePagination } from '@mui/material';
import React, { useState } from 'react';
import { UserEndpointUsage } from '../../models/user_endpoint_usage';

interface UsageStatsProps {
    usage: UserEndpointUsage[];
}

export const UsageStats: React.FC<UsageStatsProps> = ({ usage }) => {
    const [page, setPage] = useState(0);
    const [rowsPerPage, setRowsPerPage] = useState(5);

    // Handle page change
    const handleChangePage = (event: unknown, newPage: number) => {
        setPage(newPage);
    };

    // Handle rows per page change
    const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
        setRowsPerPage(parseInt(event.target.value, 10));
        setPage(0);
    };

    // Calculate rows to display on the current page
    const rowsToDisplay = usage.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage);

    return (
        <Box>
            <TableContainer component={Paper}>
                <Table stickyHeader aria-label="User Usage Table">
                    <TableHead>
                        <TableRow>
                            <TableCell>Method</TableCell>
                            <TableCell>Endpoint</TableCell>
                            <TableCell>Requests</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {rowsToDisplay.map((endpointUsage: UserEndpointUsage) => (
                            <TableRow key={endpointUsage.ID}>
                                <TableCell>
                                    <Typography variant="body1">
                                        {endpointUsage.Endpoint.Method}
                                    </Typography>
                                </TableCell>
                                <TableCell>
                                    <Typography>
                                        {endpointUsage.Endpoint.Path}
                                    </Typography>
                                </TableCell>
                                <TableCell>
                                    <Typography>
                                        {endpointUsage.RequestCount}
                                    </Typography>
                                </TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
            <TablePagination
                component="div"
                count={usage.length}
                page={page}
                onPageChange={handleChangePage}
                rowsPerPage={rowsPerPage}
                onRowsPerPageChange={handleChangeRowsPerPage}
                rowsPerPageOptions={[5, 10, 25]} // Customize options as needed
            />
        </Box>
    );
};
