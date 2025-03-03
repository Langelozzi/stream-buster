import { useState, useEffect } from 'react';
import { Box, CircularProgress, Typography, Pagination, useMediaQuery, useTheme } from '@mui/material';
import { useSearchParams } from 'react-router-dom';
import { TV } from '../../models/tv';
import { Movie } from '../../models/movie';
import { searchMovies, searchTv } from '../../api/services/search.service';
import { MediaCard } from '../../components/media-card/MediaCard';
import { useTranslation } from 'react-i18next';
import { MediaCarousel } from '../../components/media-carousel/MediaCarousel';

export const SearchResultPage = () => {
    const { t } = useTranslation();
    const [searchParams, setSearchParams] = useSearchParams();
    const theme = useTheme();
    const isMobile = useMediaQuery(theme.breakpoints.down('sm'));

    // Get query and page from URL parameters
    const query = searchParams.get('q') || '';
    const paramPage = Number(searchParams.get('page')) || 1;

    // State for the search query and the results
    const [movies, setMovies] = useState<Movie[]>([]);
    const [tv, setTv] = useState<TV[]>([]);
    const [loading, setLoading] = useState<boolean>(false);
    const [totalPages, setTotalPages] = useState<number>(1);

    const fetchResults = async (searchQuery: string, pageNum: number) => {
        try {
            setLoading(true);

            const movieSearchPage = await searchMovies(searchQuery, pageNum);
            const tvSearchPage = await searchTv(searchQuery, pageNum);

            setMovies(movieSearchPage.Results ?? []);
            setTv(tvSearchPage.Results ?? []);
            setTotalPages(Math.max(movieSearchPage.TotalPages, tvSearchPage.TotalPages));
        } catch (error) {
            console.error('Failed to fetch search results:', error);
        } finally {
            setLoading(false);
        }
    };

    // Handle page change with debounce to prevent rapid requests
    const handlePageChange = (_event: React.ChangeEvent<unknown>, value: number) => {
        if (loading) return;
        setSearchParams({ q: query, page: value.toString() });
    };

    useEffect(() => {
        if (query) {
            fetchResults(query, paramPage);
        } else {
            setMovies([]);
            setTv([]);
        }
    }, [query, paramPage]);

    return (
        <Box display="flex" flexDirection="column" alignItems="center" pt={2}>
            {loading && <CircularProgress />}

            {(!!movies || !!tv) && (
                <Box width="100%">
                    <MediaCarousel
                        title="Movie Search Results"
                        emptyMessage={t('dictionary.noResultsFound')}
                        items={movies}
                        renderItem={(movie) => <MediaCard media={movie} />}
                    />
                    <br />
                    <MediaCarousel
                        title="TV Search Results"
                        emptyMessage={t('dictionary.noResultsFound')}
                        items={tv}
                        renderItem={(tvShow) => <MediaCard media={tvShow} />}
                    />
                    {!loading && (
                        <Box
                            display="flex"
                            justifyContent="center"
                            width="100%"
                            mt={3}
                            sx={{ overflowX: 'auto', maxWidth: '100%' }}
                        >
                            <Pagination
                                onChange={handlePageChange}
                                page={paramPage}
                                count={totalPages}
                                size={isMobile ? 'small' : 'large'}
                                showFirstButton
                                showLastButton
                                color='primary'
                                variant='text'
                                sx={{
                                    '& .MuiPaginationItem-root': { color: 'white' },
                                    '& .MuiPaginationItem-icon': { color: 'white' },
                                    display: 'flex',
                                    flexWrap: 'wrap',
                                    justifyContent: 'center'
                                }}
                            />
                        </Box>
                    )}
                </Box>
            )}

            {!loading && query && !movies && !tv && (
                <Typography variant="body1">{t('dictionary.noResultsFound')}</Typography>
            )}
        </Box>
    );
};
