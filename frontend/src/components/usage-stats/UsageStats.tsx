import { Box, Typography, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, TablePagination } from '@mui/material';
import React, { useState } from 'react';
import { UserEndpointUsage } from '../../models/user_endpoint_usage';
import { Endpoint } from '../../models/endpoint';

interface UsageStatsProps {
    usage: UserEndpointUsage[];
    endpoints: Endpoint[];
}

export const UsageStats: React.FC<UsageStatsProps> = ({ usage, endpoints }) => {
    const [page, setPage] = useState(0);
    const [rowsPerPage, setRowsPerPage] = useState(5);

    // Handle page change
    const handleChangePage = (_: unknown, newPage: number) => {
        setPage(newPage);
    };

    // Handle rows per page change
    const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
        setRowsPerPage(parseInt(event.target.value, 10));
        setPage(0);
    };

    // Calculate rows to display on the current page
    const rowsToDisplay = endpoints.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage);

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
                        {rowsToDisplay.map((endpoint: Endpoint) => (
                            <TableRow key={endpoint.ID}>
                                <TableCell>
                                    <Typography variant="body1">
                                        {endpoint.Method}
                                    </Typography>
                                </TableCell>
                                <TableCell>
                                    <Typography>
                                        {endpoint.Path}
                                    </Typography>
                                </TableCell>
                                <TableCell>
                                    <Typography>
                                        {usage.find(endpointUsage => endpointUsage.EndpointID === endpoint.ID)?.RequestCount ?? 0}
                                    </Typography>
                                </TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
            <TablePagination
                component="div"
                count={endpoints.length}
                page={page}
                onPageChange={handleChangePage}
                rowsPerPage={rowsPerPage}
                onRowsPerPageChange={handleChangeRowsPerPage}
                rowsPerPageOptions={[5, 10, 25]} // Customize options as needed
            />
        </Box>
    );
};
