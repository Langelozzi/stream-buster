import { Box, Grid2, Typography, useMediaQuery, useTheme } from '@mui/material';
import React from 'react';
import { Movie } from '../../../models/movie';
import { useTranslation } from 'react-i18next';

// Define styles as a JSON object
const styles = {
    detailsContainer: {
        paddingBottom: '20px',
    }
};

interface MediaDetailsModalDescMovieProps {
    movie: Movie;
}

export const MediaDetailsModalDescMovie: React.FC<MediaDetailsModalDescMovieProps> = ({ movie }) => {
    const { t } = useTranslation();
    const theme = useTheme();
    const isMobile = useMediaQuery(theme.breakpoints.down('sm'));

    const year = new Date(movie.ReleaseDate!).getUTCFullYear();
    const overview = movie.Media?.Overview;
    const runtimeHours = Math.floor(movie.Runtime / 60);
    const runtimeMinutes = movie.Runtime % 60;

    const genres = movie.Media?.Genres?.map(genre => genre.Name).join(', ');

    return (
        <Box sx={styles.detailsContainer}>
            <Grid2 container spacing={isMobile ? 2 : 6}>
                <Grid2 xs={12} md={8}>
                    <Box>
                        <Typography>{year}&nbsp;&nbsp;{runtimeHours}{t('dictionary.hourLetter')} {runtimeMinutes}{t('dictionary.minuteLetter')}</Typography>
                        <br />
                        <Typography>{overview}</Typography>
                    </Box>
                </Grid2>
                <Grid2 xs={12} md={4}>
                    <Typography>{t('dictionary.genres')}: {genres}</Typography>
                </Grid2>
            </Grid2>
        </Box>
    );
};
