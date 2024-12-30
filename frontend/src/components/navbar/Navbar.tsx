import React, { useState } from 'react';
import { AppBar, Toolbar, Typography, Button, Menu, MenuItem, Box } from '@mui/material';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import HomeIcon from '@mui/icons-material/Home';
import { useUser } from '../../hooks/useUser'; // Adjust path as needed
import { useNavigate } from 'react-router-dom';
import { routes } from '../../router/Routes';
import { useTranslation } from 'react-i18next';
import { NavbarButton } from './navbar-button/NavbarButton';

export const Navbar: React.FC = () => {
    const { user, logout } = useUser();
    const navigate = useNavigate();
    const { t } = useTranslation();

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
                    <Typography variant="h6" sx={{ mr: 2, cursor: 'pointer' }} onClick={() => navigate(routes.home)}>
                        {t('dictionary.streambuster')}
                    </Typography>
                    {user && (
                        <NavbarButton
                            Icon={HomeIcon}
                            label={t('button.browse')}
                            onClick={() => navigate(routes.browse)}
                        />
                    )}
                </Box>

                {!user ? (
                    <Button color="inherit" onClick={() => navigate(routes.login)}>
                        {t('button.login')}
                    </Button>
                ) : (
                    <>
                        <Button color='inherit' onClick={handleProfileMenuOpen} component='span'>
                            <Typography>
                                {user.FirstName}
                            </Typography>
                            <AccountCircleIcon />
                        </Button>
                        <Menu
                            anchorEl={anchorEl}
                            open={open}
                            onClose={handleMenuClose}
                            anchorOrigin={{
                                vertical: 'bottom',
                                horizontal: 'right',
                            }}
                            transformOrigin={{
                                vertical: 'top',
                                horizontal: 'right',
                            }}
                        >
                            <MenuItem onClick={() => { handleMenuClose(); navigate(routes.dashboard); }}>{t('button.dashboard')}</MenuItem>
                            <MenuItem onClick={handleLogout}>{t('button.logout')}</MenuItem>
                        </Menu>
                    </>
                )}
            </Toolbar>
        </AppBar>
    );
};
