import React, { useState } from 'react';
import { AppBar, Toolbar, Typography, Button, Menu, MenuItem, Box, IconButton, useMediaQuery, useTheme } from '@mui/material';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import HomeIcon from '@mui/icons-material/Home';
import SearchIcon from '@mui/icons-material/Search';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import { useUser } from '../../hooks/useUser';
import { useLocation, useNavigate } from 'react-router-dom';
import { routes } from '../../router/Routes';
import { useTranslation } from 'react-i18next';
import { NavbarButton } from './navbar-button/NavbarButton';
import { NavSearch } from './search/NavSearch';

export const Navbar: React.FC = () => {
    const { user, logout } = useUser();
    const navigate = useNavigate();
    const location = useLocation();
    const { t } = useTranslation();
    const theme = useTheme();
    const isMobile = useMediaQuery(theme.breakpoints.down('sm'));
    const [searchExpanded, setSearchExpanded] = useState(false);

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

    const toggleSearch = () => {
        setSearchExpanded(!searchExpanded);
    };

    // Show search only on non-watch pages and when user is logged in
    const shouldShowSearch = user && !location.pathname.startsWith('/watch');

    return (
        <AppBar position="static" sx={{ marginBottom: 2 }}>
            <Toolbar>
                {isMobile && searchExpanded ? (
                    // Mobile expanded search view
                    <Box sx={{ display: 'flex', width: '100%', alignItems: 'center' }}>
                        <IconButton
                            color="inherit"
                            onClick={toggleSearch}
                            edge="start"
                            sx={{ mr: 1 }}
                        >
                            <ArrowBackIcon />
                        </IconButton>
                        <Box sx={{ flexGrow: 1 }}>
                            <NavSearch autoFocus={true} />
                        </Box>
                    </Box>
                ) : (
                    // Normal navbar view
                    <>
                        <Box sx={{ display: 'flex', alignItems: 'center', flexGrow: 1 }}>
                            {!isMobile && (
                                <Typography
                                    variant={isMobile ? "body1" : "h6"}
                                    sx={{ mr: isMobile ? 1 : 2, cursor: 'pointer' }}
                                    onClick={() => navigate(routes.browse)}
                                    noWrap
                                >
                                    {t('dictionary.streambuster')}
                                </Typography>
                            )}
                            {user && !isMobile && (
                                <NavbarButton
                                    Icon={HomeIcon}
                                    label={t('button.browse')}
                                    onClick={() => navigate(routes.browse)}
                                />
                            )}
                            {user && isMobile && (
                                <IconButton color="inherit" onClick={() => navigate(routes.browse)}>
                                    <HomeIcon />
                                </IconButton>
                            )}
                        </Box>

                        {!user ? (
                            <Button
                                color="inherit"
                                onClick={() => navigate(routes.login)}
                                size={isMobile ? "small" : "medium"}
                            >
                                {t('button.login')}
                            </Button>
                        ) : (
                            <>
                                {shouldShowSearch && !isMobile && (
                                    <Box mr={2}>
                                        <NavSearch />
                                    </Box>
                                )}

                                {shouldShowSearch && isMobile && (
                                    <IconButton
                                        color="inherit"
                                        onClick={toggleSearch}
                                        sx={{ mr: 1 }}
                                    >
                                        <SearchIcon />
                                    </IconButton>
                                )}

                                <Button
                                    color='inherit'
                                    onClick={handleProfileMenuOpen}
                                    component='span'
                                    size={isMobile ? "small" : "medium"}
                                >
                                    {!isMobile && (
                                        <Typography mr={1}>
                                            {user.FirstName}
                                        </Typography>
                                    )}
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
                    </>
                )}
            </Toolbar>
        </AppBar>
    );
};
