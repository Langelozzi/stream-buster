import { CheckCircle, Cancel, Help } from "@mui/icons-material";
import { Box, Typography } from "@mui/material";
import React from "react";

interface AvailabilityInfoProps {
    available: number;
}

export const AvailabilityInfo: React.FC<AvailabilityInfoProps> = ({ available }) => {
    const styles = {
        availabilityContainer: {
            display: 'flex',
            alignItems: 'center',
        },
        availabilityText: {
            marginLeft: 1
        }
    }

    return (
        <div>
            {available === 1 && (
                <Box sx={styles.availabilityContainer}>
                    <CheckCircle color='success' />
                    <Typography sx={styles.availabilityText}>
                        Available
                    </Typography>
                </Box>
            )}
            {available === 0 && (
                <Box sx={styles.availabilityContainer}>
                    <Cancel color='error' />
                    <Typography sx={styles.availabilityText}>
                        Unavailable
                    </Typography>
                </Box>
            )}
            {available === -1 && (
                <Box sx={styles.availabilityContainer}>
                    <Help color='warning' />
                    <Typography sx={styles.availabilityText}>
                        Unknown
                    </Typography>
                </Box>
            )}
        </div>
    )
};