import { Box, Grid2, Typography } from '@mui/material';
import { makeStyles } from '@mui/styles';
import React from 'react';
import { TV } from '../../../models/tv';
import { Episode } from '../../../models/episode';

const useStyles = makeStyles(() => ({
    detailsContainer: {
        paddingBottom: '20px',
    },
}));

interface MediaDetailsModalDescTVProps {
    tv: TV;
    currentEpisode?: Episode
}

export const MediaDetailsModalDescTV: React.FC<MediaDetailsModalDescTVProps> = ({ tv, currentEpisode }) => {
    const classes = useStyles();

    const endYear = new Date(tv.LastAirDate!).getUTCFullYear();
    const numSeasons = tv.SeasonCount;

    return (
        <Box className={classes.detailsContainer}>
            <Grid2 container spacing={6}>
                <Grid2 size={8}>
                    <Box>
                        <Typography>{endYear}&nbsp;&nbsp;{numSeasons} Season{numSeasons > 1 ? 's' : ''}</Typography>
                        <br />
                        <Typography variant='h5'>S{currentEpisode?.SeasonNumber}:E{currentEpisode?.EpisodeNumber} "{currentEpisode?.Name}"</Typography>
                        <Typography>{currentEpisode?.Overview}</Typography>
                    </Box>
                </Grid2>
                <Grid2 size={4}>
                    <Typography>Genres: {tv.Media?.Genres?.map(genre => genre.Name).join(', ')}</Typography>
                </Grid2>
            </Grid2>
        </Box>
    );
};
