import { Box, LinearProgress, Typography } from "@mui/material";
import React, { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";

interface TotalUsageProgressProps {
    requestCount: number;
    maxRequests: number;
    isAdmin: boolean;
}

export const TotalUsageProgress: React.FC<TotalUsageProgressProps> = (props) => {
    const {
        requestCount,
        maxRequests,
        isAdmin
    } = props;

    const { t } = useTranslation();

    const [progressValue, setProgressValue] = useState<number>(0);
    const [progressColor, setProgressColor] = useState<string>("primary");

    const updateProgress = () => {
        if (isAdmin) {
            setProgressValue(0);
            setProgressColor("success");
            return;
        }

        let value = ((requestCount || 0) / maxRequests) * 100;
        if (value > 100) {
            value = 100;
            setProgressColor("error")
        } else if (value < 0) {
            value = 0;
        }

        setProgressValue(value);
    };

    useEffect(() => {
        updateProgress();
    }, [requestCount, maxRequests, isAdmin])

    return (
        <Box display="flex" flexDirection="column">
            <Typography variant="body1">
                {t('dictionary.requestsUsed')}: {requestCount || 0} / {!isAdmin ? (maxRequests || 0) : (<span>&infin;</span>)}
            </Typography >
            <LinearProgress
                variant="determinate"
                color={progressColor as any}
                value={progressValue}
                sx={{ mt: 1, height: 10, borderRadius: 5 }}
            />
            {
                progressValue === 100 && (
                    <Typography variant="body2" color="error">{t('error.exceededRequestCount')}</Typography>
                )
            }
        </Box >
    )
}