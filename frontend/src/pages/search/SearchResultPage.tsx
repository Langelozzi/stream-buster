import { useState, useEffect } from 'react';
import { Box, CircularProgress, Typography, Pagination } from '@mui/material';
import { useSearchParams } from 'react-router-dom';
import { TV } from '../../models/tv';
import { Movie } from '../../models/movie';
import { searchMulti } from '../../api/services/search.service';
import { MediaCard } from '../../components/media-card/MediaCard';
import { useTranslation } from 'react-i18next';

export const SearchResultPage = () => {
    const { t } = useTranslation();
    const [searchParams, setSearchParams] = useSearchParams();

    // Get query and page from URL parameters
    const query = searchParams.get('q') || '';
    const paramPage = Number(searchParams.get('page')) || 1;

    // State for the search query and the results
    const [results, setResults] = useState<(TV | Movie)[]>([]);
    const [loading, setLoading] = useState<boolean>(false);
    const [totalPages, setTotalPages] = useState<number>(1);

    // Function to fetch search results based on query
    const fetchResults = async (searchQuery: string, pageNum: number) => {
        try {
            setLoading(true);
            const searchPage = await searchMulti(searchQuery, pageNum);
            setResults(searchPage.Results);
            setTotalPages(searchPage.TotalPages);
        } catch (error) {
            console.error('Failed to fetch search results:', error);
        } finally {
            // Only update loading state if component is still mounted
            setLoading(false);
        }
    };

    // Handle page change with debounce to prevent rapid requests
    const handlePageChange = (_event: React.ChangeEvent<unknown>, value: number) => {
        // Prevent rapid page changes by disabling if already loading
        if (loading) return;

        // Update URL with new page number
        setSearchParams({ q: query, page: value.toString() });
    };

    // Effect to fetch results when URL parameters change
    useEffect(() => {
        if (query) {
            fetchResults(query, paramPage);
        } else {
            setResults([]);
        }
    }, [query, paramPage]);

    return (
        <Box display="flex" flexDirection="column" alignItems="center" p={2}>
            {/* Loading spinner */}
            {loading && <CircularProgress />}

            {/* Search results */}
            {results && results.length > 0 && (
                <Box width="100%">
                    <Typography variant="h5" sx={{ fontWeight: "bold" }}>{t('dictionary.searchResults')}:</Typography>
                    <Box display="flex" flexWrap="wrap" justifyContent="flex-start" gap={2} mt={2}>
                        {results.map((media, index) => (
                            <MediaCard media={media} key={`${media.MediaID}-${index}`} />
                        ))}
                    </Box>
                    {/* Pagination controls - hidden during loading */}
                    {!loading && (
                        <Box display="flex" justifyContent="center" width="100%" mt={3}>
                            <Pagination
                                onChange={handlePageChange}
                                page={paramPage}
                                count={totalPages}
                                size='large'
                                showFirstButton
                                showLastButton
                                color='primary'
                                sx={{
                                    '& .MuiPaginationItem-root': {
                                        color: 'white'
                                    },
                                    '& .MuiPaginationItem-icon': {
                                        color: 'white'
                                    }
                                }}
                            />
                        </Box>
                    )}
                </Box>
            )}

            {/* Message when no results */}
            {!loading && query && !results && (
                <Typography variant="body1">{t('dictionary.noResultsFound')}</Typography>
            )}
        </Box>
    );
};
