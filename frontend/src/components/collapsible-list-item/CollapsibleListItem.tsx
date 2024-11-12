import React, { ReactNode } from 'react';
import {
    List,
    ListItem,
    ListItemText,
    Collapse,
    Card,
    CardContent,
    IconButton,
} from '@mui/material';
import { ExpandLess, ExpandMore } from '@mui/icons-material';

// Define the prop types for CollapsibleListItem
interface CollapsibleListItemProps {
    label: string;
    open: boolean;
    onToggle: () => void;
    children: ReactNode;
}

export const CollapsibleListItem: React.FC<CollapsibleListItemProps> = (props) => {
    const {
        label,
        open,
        onToggle,
        children
    } = props;

    return (
        <List>
            <ListItem onClick={onToggle} component='button'>
                <ListItemText primary={label} />
                <IconButton edge='end'>
                    {open ? <ExpandLess /> : <ExpandMore />}
                </IconButton>
            </ListItem>
            <Collapse in={open} timeout='auto' unmountOnExit>
                <Card variant='outlined' style={{ margin: '10px' }}>
                    <CardContent>
                        {children}
                    </CardContent>
                </Card>
            </Collapse>
        </List>
    );
};