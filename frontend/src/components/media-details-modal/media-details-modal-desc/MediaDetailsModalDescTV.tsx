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

    const currentSeason = 3;
    const currentEpisodeNum = 31
    const currentEpisodeName = "Maximum Security"
    const currentEpisodeOverview = "To root out the crooked FBI agent who targeted Pimento, the team stages an elaborate fake funeral and sends Amy on a risky undercover mission."
    const genres = tv.Media?.Genres?.map(genre => genre.Name).join(', ');

    return (
        <Box className={classes.detailsContainer}>
            <Grid2 container spacing={6}>
                <Grid2 size={8}>
                    <Box>
                        <Typography>{endYear}&nbsp;&nbsp;{numSeasons} Season{numSeasons > 1 ? 's' : ''}</Typography>
                        <br />
                        <Typography variant='h5'>S{currentSeason}:E{currentEpisodeNum} "{currentEpisodeName}"</Typography>
                        <Typography>{currentEpisodeOverview}</Typography>
                    </Box>
                </Grid2>
                <Grid2 size={4}>
                    <Typography>Genres: {genres}</Typography>
                </Grid2>
            </Grid2>
        </Box>
    );
};
