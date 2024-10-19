import { useState, useEffect, FormEvent } from 'react';
import { TextField, Button, Box, CircularProgress, Typography } from '@mui/material';
import { TV } from '../../models/tv';
import { Movie } from '../../models/movie';
import { searchMulti } from '../../api/search'; // Assuming this is the API function you've created

export const Search = () => {
    // State for the search query and the results
    const [query, setQuery] = useState<string>('');
    const [results, setResults] = useState<(TV | Movie)[]>([]);
    const [loading, setLoading] = useState<boolean>(false);

    // Function to fetch search results based on query
    const fetchResults = async (searchQuery: string) => {
        try {
            setLoading(true);
            const data = await searchMulti(searchQuery);
            setResults(data);
        } catch (error) {
            console.error('Failed to fetch search results:', error);
        } finally {
            setLoading(false);
        }
    };

    // Handle form submission
    const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        if (query.trim()) {
            fetchResults(query);
        }
    };

    return (
        <Box display="flex" flexDirection="column" alignItems="center" p={2}>
            {/* Search form */}
            <Box component="form" onSubmit={handleSubmit} display="flex" mb={2}>
                <TextField
                    label="Search"
                    variant="outlined"
                    value={query}
                    onChange={(e) => setQuery(e.target.value)}
                    sx={{ mr: 2 }}
                />
                <Button type="submit" variant="contained" color="primary">
                    Search
                </Button>
            </Box>

            {/* Loading spinner */}
            {loading && <CircularProgress />}

            {/* Search results */}
            {!loading && results.length > 0 && (
                <Box mt={2} width="100%">
                    <Typography variant="h6">Search Results:</Typography>
                    <pre>{JSON.stringify(results, null, 2)}</pre>
                </Box>
            )}

            {/* Message when no results */}
            {!loading && results.length === 0 && (
                <Typography variant="body1">No results found</Typography>
            )}
        </Box>
    );
};
