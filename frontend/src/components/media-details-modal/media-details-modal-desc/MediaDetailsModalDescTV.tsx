import { Box, Grid2, Typography, useMediaQuery, useTheme } from '@mui/material';
import React from 'react';
import { TV } from '../../../models/tv';
import { Episode } from '../../../models/episode';
import { useTranslation } from 'react-i18next';

// Define styles as a JSON object
const styles = {
    detailsContainer: {
        paddingBottom: '20px',
    },
};

interface MediaDetailsModalDescTVProps {
    tv: TV;
    currentEpisode?: Episode
    available: number
}

export const MediaDetailsModalDescTV: React.FC<MediaDetailsModalDescTVProps> = ({ tv, currentEpisode, available }) => {
    const { t } = useTranslation();
    const theme = useTheme();
    const isMobile = useMediaQuery(theme.breakpoints.down('sm'));

    const endYear = new Date(tv.LastAirDate!).getUTCFullYear();
    const numSeasons = tv.SeasonCount;

    return (
        <Box sx={styles.detailsContainer}>
            <Grid2 container spacing={isMobile ? 2 : 6}>
                <Grid2 size={{ xs: 12, md: 8 }}>
                    <Box>
                        <Typography>{endYear}&nbsp;&nbsp;{numSeasons} {numSeasons > 1 ? t('dictionary.seasons') : t('dictionary.season')}</Typography>
                        <br />
                    </Box>
                </Grid2>
                <Grid2 size={{ xs: 12, md: 4 }}>
                    <Typography>{t('dictionary.genres')}: {tv.Media?.Genres?.map(genre => genre.Name).join(', ')}</Typography>
                </Grid2>

                {available !== 0 && (
                    <Grid2 size={12}>
                        <Typography variant={isMobile ? 'h6' : 'h5'}>{t('dictionary.seasonLetter')}{currentEpisode?.SeasonNumber}:{t('dictionary.episodeLetter')}{currentEpisode?.EpisodeNumber} "{currentEpisode?.Name}"</Typography>
                        <Typography>{currentEpisode?.Overview}</Typography>
                    </Grid2>
                )}
            </Grid2>
        </Box>
    );
};
