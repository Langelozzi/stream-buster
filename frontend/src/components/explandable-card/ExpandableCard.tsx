import React, { ReactNode, useState } from 'react';
import { Card, CardContent, Collapse, IconButton, Box } from '@mui/material';
import { ExpandLess, ExpandMore } from '@mui/icons-material';

interface ExpandableCardProps {
    headerContent: ReactNode;    // Content shown when card is collapsed
    expandedContent: ReactNode;  // Content shown when card is expanded
    defaultOpen?: boolean;       // Optional: whether the card is expanded by default
}

export const ExpandableCard: React.FC<ExpandableCardProps> = ({
    headerContent,
    expandedContent,
    defaultOpen = false
}) => {
    const [open, setOpen] = useState(defaultOpen);

    const handleToggle = () => {
        setOpen((prevOpen) => !prevOpen);
    };

    return (
        <Card variant="outlined" style={{ margin: '10px' }}>
            <CardContent onClick={handleToggle} style={{ cursor: 'pointer', display: 'flex', alignItems: 'center' }}>
                <Box flexGrow={1}>{headerContent}</Box>
            </CardContent>
            <Collapse in={open} timeout="auto" unmountOnExit>
                <CardContent>
                    {expandedContent}

                </CardContent>
            </Collapse>
            {/* Icon button at the bottom */}
            <Box display="flex" justifyContent="center" onClick={handleToggle} style={{ cursor: 'pointer' }}>
                <IconButton>
                    {open ? <ExpandLess /> : <ExpandMore />}
                </IconButton>
            </Box>
        </Card>
    );
};