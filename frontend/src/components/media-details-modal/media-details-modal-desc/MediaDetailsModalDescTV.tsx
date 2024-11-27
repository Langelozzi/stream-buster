import { Box, Grid2, Typography } from '@mui/material';
import { makeStyles } from '@mui/styles';
import React from 'react';
import { TV } from '../../../models/tv';
import { Episode } from '../../../models/episode';
import { useTranslation } from 'react-i18next';

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
    const { t } = useTranslation();
    const classes = useStyles();

    const endYear = new Date(tv.LastAirDate!).getUTCFullYear();
    const numSeasons = tv.SeasonCount;

    return (
        <Box className={classes.detailsContainer}>
            <Grid2 container spacing={6}>
                <Grid2 size={8}>
                    <Box>
                        <Typography>{endYear}&nbsp;&nbsp;{numSeasons} {numSeasons > 1 ? t('dictionary.seasons') : t('dictionary.season')}</Typography>
                        <br />
                    </Box>
                </Grid2>
                <Grid2 size={4}>
                    <Typography>{t('dictionary.genres')}: {tv.Media?.Genres?.map(genre => genre.Name).join(', ')}</Typography>
                </Grid2>

                <Grid2 size={12}>
                    <Typography variant='h5'>{t('dictionary.seasonLetter')}{currentEpisode?.SeasonNumber}:{t('dictionary.episodeLetter')}{currentEpisode?.EpisodeNumber} "{currentEpisode?.Name}"</Typography>
                    <Typography>{currentEpisode?.Overview}</Typography>
                </Grid2>
            </Grid2>
        </Box>
    );
};
