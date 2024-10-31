import { Box, Typography, LinearProgress } from "@mui/material";
import React, { useEffect, useState } from "react";
import { Usage } from "../../../models/usage";
import { getUserUsage } from "../../../api/services/user.service";

interface UsageStatsProps {
    userId: number;
    isAdmin: boolean;
    maxRequestCount: number;
}

export const UsageStats: React.FC<UsageStatsProps> = (props) => {
    const {
        userId,
        maxRequestCount,
        isAdmin
    } = props;

    const [usage, setUsage] = useState<Usage | undefined>();
    const [progressValue, setProgressValue] = useState<number>(0);
    const [progressColor, setProgressColor] = useState<string>("primary");

    // We need to refetch the usage for the user everytime this component reloads
    const fetchUsage = async () => {
        try {
            const fetchedUsage = await getUserUsage(userId);
            setUsage(fetchedUsage);
        } catch (e) {
            setUsage(undefined);
        }
    };

    const updateProgress = () => {
        if (isAdmin) {
            setProgressValue(0);
            setProgressColor("success");
            return;
        }

        let value = ((usage?.RequestCount || 0) / maxRequestCount) * 100;
        if (value > 100) {
            value = 100;
            setProgressColor("error")
        } else if (value < 0) {
            value = 0;
        }

        setProgressValue(value);
    };

    useEffect(() => {
        fetchUsage();
    }, [userId, maxRequestCount])

    useEffect(() => {
        updateProgress();
    }, [usage])

    return (
        <Box display="flex" flexDirection="column">
            <Typography variant="body1">
                Requests used: {usage?.RequestCount || 0} / {!isAdmin ? (maxRequestCount || 0) : (<span>&infin;</span>)}
            </Typography>
            <LinearProgress
                variant="determinate"
                color={progressColor as any}
                value={progressValue}
                sx={{ mt: 1, height: 10, borderRadius: 5 }}
            />
            {progressValue === 100 && (
                <Typography variant="body2" color="error">You have exceeded your free credits</Typography>
            )}
        </Box>
    )
}