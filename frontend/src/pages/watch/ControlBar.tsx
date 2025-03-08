
import React from "react";
import { Box, Button, Tooltip } from "@mui/material";
import SkipNextIcon from "@mui/icons-material/SkipNext";
import SkipPreviousIcon from "@mui/icons-material/SkipPrevious";

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
            width="60%"
            padding={2}
            sx={{
                "@media (max-width: 600px)": {
                    width: "100%", // Adjust width for mobile if needed
                },
            }}
        >
            {/* Watch Previous Episode */}
            <Tooltip title="Watch Previous Episode" arrow>
                <Button
                    variant="contained"
                    onClick={goToPrev}
                    size="large"
                    startIcon={<SkipPreviousIcon />}
                    sx={{
                        width: "120px",
                        "@media (max-width: 600px)": {
                            width: "80px",
                            justifyContent: "center",
                        },
                    }}
                >
                    <Box
                        component="span"
                        sx={{
                            "@media (max-width: 600px)": {
                                display: "none",
                            },
                        }}
                    >
                        Previous
                    </Box>
                </Button>
            </Tooltip>

            {/* Watch Next Episode */}
            <Tooltip title="Watch Next Episode" arrow>
                <Button
                    variant="contained"
                    onClick={goToNext}
                    size="large"
                    endIcon={<SkipNextIcon />}
                    sx={{
                        width: "120px",
                        "@media (max-width: 600px)": {
                            width: "80px",
                            justifyContent: "center",
                        },
                    }}
                >
                    <Box
                        component="span"
                        sx={{
                            "@media (max-width: 600px)": {
                                display: "none",
                            },
                        }}
                    >
                        Next
                    </Box>
                </Button>
            </Tooltip>
        </Box>
    );
};

export default ControlBar;
