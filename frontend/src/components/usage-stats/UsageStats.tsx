import { Box, Typography, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from "@mui/material";
import React from "react";
import { UserEndpointUsage } from "../../models/user_endpoint_usage";

interface UsageStatsProps {
    usage: UserEndpointUsage[];
}

export const UsageStats: React.FC<UsageStatsProps> = (props) => {
    const {
        usage,
    } = props;

    return (
        <Box>
            <TableContainer component={Paper}>
                <Table aria-label="User Usage Table">
                    <TableHead>
                        <TableRow>
                            <TableCell>Method</TableCell>
                            <TableCell>Endpoint</TableCell>
                            <TableCell>Requests</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {usage !== undefined && (
                            usage.map((endpointUsage: UserEndpointUsage) => (
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
                            ))
                        )}
                    </TableBody>
                </Table>
            </TableContainer>
        </Box>
    )
}