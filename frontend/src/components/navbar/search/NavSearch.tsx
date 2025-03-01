import { styled, alpha } from '@mui/material/styles';
import InputBase from '@mui/material/InputBase';
import { useNavigate, useLocation, useSearchParams } from 'react-router-dom';
import { useState, useEffect, useCallback } from 'react';
import { Search, Close } from '@mui/icons-material';
import debounce from 'lodash/debounce';
import { IconButton } from '@mui/material';

const SearchInput = styled('div')(({ theme }) => ({
    position: 'relative',
    borderRadius: theme.shape.borderRadius,
    backgroundColor: alpha(theme.palette.common.white, 0.15),
    '&:hover': {
        backgroundColor: alpha(theme.palette.common.white, 0.25),
    },
    marginLeft: 0,
    width: '100%',
    [theme.breakpoints.up('sm')]: {
        marginLeft: theme.spacing(1),
        width: 'auto',
    },
}));

const SearchIconWrapper = styled('div')(({ theme }) => ({
    padding: theme.spacing(0, 2),
    height: '100%',
    position: 'absolute',
    pointerEvents: 'none',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
}));

const StyledInputBase = styled(InputBase)(({ theme }) => ({
    color: 'inherit',
    width: '100%',
    '& .MuiInputBase-input': {
        padding: theme.spacing(1, 1, 1, 0),
        // vertical padding + font size from searchIcon
        paddingLeft: `calc(1em + ${theme.spacing(4)})`,
        transition: theme.transitions.create('width'),
        [theme.breakpoints.up('sm')]: {
            width: '12ch',
            '&:focus': {
                width: '20ch',
            },
        },
    },
}));

const ClearButton = styled(IconButton)(({ theme }) => ({
    padding: 4,
    position: 'absolute',
    right: theme.spacing(1),
    top: '50%',
    transform: 'translateY(-50%)',
    visibility: 'visible',
    color: alpha(theme.palette.common.white, 0.7),
    '&:hover': {
        color: theme.palette.common.white,
    },
}));

export const NavSearch = () => {
    const [searchValue, setSearchValue] = useState('');
    const navigate = useNavigate();
    const location = useLocation();
    const [searchParams] = useSearchParams();
    const [previousPath, setPreviousPath] = useState<string>('/browse'); // Default fallback
    
    // Effect to track previous path
    useEffect(() => {
        // Only store paths that are not search or watch pages as previous paths
        if (location.pathname !== '/search' && !location.pathname.startsWith('/watch')) {
            setPreviousPath(location.pathname);
        }
    }, [location.pathname]);
    
    // Create a debounced version of the navigation function with shorter delay
    // eslint-disable-next-line react-hooks/exhaustive-deps
    const debouncedNavigate = useCallback(
        debounce((query: string) => {
            if (query.trim() !== '') {
                navigate(`/search?q=${encodeURIComponent(query.trim())}&page=1`);
            } else if (query.trim() === '') {
                // Navigate back to previous route instead of hardcoded /browse
                navigate(previousPath);
            }
        }, 300), // Reduced from 500ms to 300ms for more immediate feedback
        [navigate, previousPath]
    );

    // Handle input change
    const handleSearchChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const newValue = event.target.value;
        setSearchValue(newValue);
        debouncedNavigate(newValue);
    };
    
    // Handle clear button click
    const handleClearSearch = () => {
        setSearchValue('');
        debouncedNavigate('');
    };
    
    // Effect to sync search input with URL query parameter
    useEffect(() => {
        // Check if we're on the search page
        if (location.pathname === '/search') {
            const queryParam = searchParams.get('q');
            if (queryParam) {
                setSearchValue(queryParam);
            }
        } else {
            // Clear search value when not on search page
            setSearchValue('');
        }
    }, [location.pathname, searchParams]);

    // Clean up debounced function on unmount
    useEffect(() => {
        return () => {
            debouncedNavigate.cancel();
        };
    }, [debouncedNavigate]);

    return (
        <SearchInput>
            <SearchIconWrapper>
                <Search />
            </SearchIconWrapper>
            <StyledInputBase
                placeholder="Search…"
                inputProps={{ 'aria-label': 'search' }}
                value={searchValue}
                onChange={handleSearchChange}
            />
            {searchValue && (
                <ClearButton 
                    size="small" 
                    aria-label="clear search" 
                    onClick={handleClearSearch}
                >
                    <Close fontSize="small" />
                </ClearButton>
            )}
        </SearchInput>
    );
}