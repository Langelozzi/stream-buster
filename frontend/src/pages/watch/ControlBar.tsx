import React from "react";
import { Box, IconButton, Tooltip } from "@mui/material";
import SkipNextIcon from '@mui/icons-material/SkipNext';
import SkipPreviousIcon from '@mui/icons-material/SkipPrevious';
interface ControlBarProps {
    goToNext: () => void;
    goToPrev: () => void;
}
const ControlBar: React.FC<ControlBarProps> = ({ goToNext, goToPrev }) => {
    return (
        <Box
            display="flex"
            justifyContent="space-between"
            alignItems="center"
            width="100%"
            padding={2}
        >
            {/* Watch Previous Episode */}
            <Tooltip title="Watch Previous Episode" arrow>
                <IconButton
                    onClick={goToPrev}
                    size="large"
                >
                    <SkipPreviousIcon />
                </IconButton>
            </Tooltip>

            {/* Watch Next Episode */}
            <Tooltip title="Watch Next Episode" arrow>
                <IconButton
                    onClick={goToNext}
                    size="large"
                >
                    <SkipNextIcon />
                </IconButton>
            </Tooltip>
        </Box>
    );
};

export default ControlBar;
