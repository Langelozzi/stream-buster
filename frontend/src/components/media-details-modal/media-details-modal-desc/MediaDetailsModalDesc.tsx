import { Box, Grid2, Typography } from '@mui/material';
import { makeStyles } from '@mui/styles';
import React from 'react';
import { Movie } from '../../../models/movie';
import { TV } from '../../../models/tv';
import { Episode } from '../../../models/episode';

const useStyles = makeStyles(() => ({
    detailsContainer: {
        paddingBottom: '20px',
    },

}));

interface MediaDetailsModalDescProps {
    media: Movie | TV;
    currentEpisode?: Episode
}

export const MediaDetailsModalDesc: React.FC<MediaDetailsModalDescProps> = ({ media, currentEpisode }) => {
    const classes = useStyles();

    const endYear = 2021;
    const numSeasons = 8;

    const currentSeason = 3;
    const currentEpisodeNum = 31
    const currentEpisodeName = "Maximum Security"
    const currentEpisodeOverview = "To root out the crooked FBI agent who targeted Pimento, the team stages an elaborate fake funeral and sends Amy on a risky undercover mission."
    const genres = ["Sitcoms", "TV Comedies", "Crime TV Shows"]

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
                    <Typography>Genres: {genres.join(", ")}</Typography>
                </Grid2>
            </Grid2>
        </Box>
    );
};
