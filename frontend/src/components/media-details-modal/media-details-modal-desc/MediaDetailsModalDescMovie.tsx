import { Box, Grid2, Typography } from '@mui/material';
import { makeStyles } from '@mui/styles';
import React from 'react';
import { Movie } from '../../../models/movie';

const useStyles = makeStyles(() => ({
    detailsContainer: {
        paddingBottom: '20px',
    },

}));

interface MediaDetailsModalDescMovieProps {
    movie: Movie;
}

export const MediaDetailsModalDescMovie: React.FC<MediaDetailsModalDescMovieProps> = ({ movie }) => {
    const classes = useStyles();

    const year = new Date(movie.ReleaseDate!).getUTCFullYear();
    const overview = movie.Media?.Overview;
    const runtimeHours = Math.floor(movie.Runtime / 60);
    const runtimeMinutes = movie.Runtime % 60;

    const genres = movie.Media?.Genres?.map(genre => genre.Name).join(', ');

    return (
        <Box className={classes.detailsContainer}>
            <Grid2 container spacing={6}>
                <Grid2 size={8}>
                    <Box>
                        <Typography>{year}&nbsp;&nbsp;{runtimeHours}h {runtimeMinutes}m</Typography>
                        <br />
                        <Typography>{overview}</Typography>
                    </Box>
                </Grid2>
                <Grid2 size={4}>
                    <Typography>Genres: {genres}</Typography>
                </Grid2>
            </Grid2>
        </Box>
    );
};
