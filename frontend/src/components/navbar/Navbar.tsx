import React, { useState } from 'react';
import { AppBar, Toolbar, Typography, Button, IconButton, Menu, MenuItem, Box } from '@mui/material';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import { useUser } from '../../hooks/useUser'; // Adjust path as needed
import { useNavigate } from 'react-router-dom';

export const Navbar: React.FC = () => {
    const { user, logout } = useUser();
    const navigate = useNavigate();
    const [anchorEl, setAnchorEl] = useState<HTMLAnchorElement | null>(null);
    const open = Boolean(anchorEl);

    const handleProfileMenuOpen = (event: React.MouseEvent<HTMLAnchorElement>) => {
        setAnchorEl(event.currentTarget);
    };

    const handleMenuClose = () => {
        setAnchorEl(null);
    };

    const handleLogout = () => {
        handleMenuClose();
        logout();
    };

    return (
        <AppBar position="static" sx={{ marginBottom: 2 }}>
            <Toolbar>
                <Box sx={{ display: 'flex', alignItems: 'center', flexGrow: 1 }}>
                    <Typography variant="h6" sx={{ mr: 2 }}>
                        Streambuster
                    </Typography>
                    {user && (
                        <Button color="inherit" onClick={() => navigate('/browse')}>
                            Browse
                        </Button>
                    )}
                </Box>

                {!user ? (
                    <Button color="inherit" onClick={() => navigate('/login')}>
                        Login
                    </Button>
                ) : (
                    <>
                        <IconButton
                            color="inherit"
                            onClick={handleProfileMenuOpen}
                            component="span"
                        >
                            <AccountCircleIcon />
                        </IconButton>
                        <Menu
                            anchorEl={anchorEl}
                            open={open}
                            onClose={handleMenuClose}
                            anchorOrigin={{
                                vertical: 'top',
                                horizontal: 'right',
                            }}
                            transformOrigin={{
                                vertical: 'top',
                                horizontal: 'right',
                            }}
                        >
                            <MenuItem onClick={() => { handleMenuClose(); navigate('/dashboard'); }}>Dashboard</MenuItem>
                            <MenuItem onClick={handleLogout}>Logout</MenuItem>
                        </Menu>
                    </>
                )}
            </Toolbar>
        </AppBar>
    );
};
