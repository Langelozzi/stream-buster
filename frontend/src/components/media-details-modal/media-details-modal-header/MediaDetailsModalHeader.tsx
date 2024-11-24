import React from 'react';
import { Box, IconButton, Button, Typography, Tooltip } from '@mui/material';
import { PlayArrow, Add, ThumbUp } from '@mui/icons-material';
import { makeStyles } from '@mui/styles';
import { Movie } from '../../../models/movie';
import { TV } from '../../../models/tv';
import { useNavigate } from 'react-router-dom';
import { Episode } from '../../../models/episode';
import { useTranslation } from 'react-i18next';
import { useUser } from '../../../hooks/useUser';
import { onAddToList } from '../../../api/services/currentlyWatching.service';

const useStyles = makeStyles(() => ({
    modalContainer: {
        position: 'relative',
        width: '100%',
        height: '600px',
        overflow: 'hidden',
    },
    imageOverlay: {
        position: 'absolute',
        width: '100%',
        height: '100%',
        backgroundSize: 'cover',
        opacity: 0.4,
    },
    title: {
        position: 'absolute',
        bottom: '70px', // adjust this value if needed
        left: '20px',
        zIndex: 3,
    },
    controls: {
        position: 'absolute',
        bottom: '20px',
        left: '20px',
        display: 'flex',
        gap: '10px',
        zIndex: 3,
    },
    controlButton: {
        display: 'flex',
        alignItems: 'center',
        gap: '4px',
    },
    roundButton: {
        borderRadius: '50%',
        backgroundColor: 'rgba(255, 255, 255, 0.3)',
        '&:hover': {
            backgroundColor: 'rgba(255, 255, 255, 0.5)',
        },
    },
}));

interface MediaDetailsModalHeaderProps {
    media: Movie | TV;
    currentEpisode?: Episode
}

export const MediaDetailsModalHeader: React.FC<MediaDetailsModalHeaderProps> = ({ media, currentEpisode }) => {
    // Hooks
    const { t } = useTranslation();
    const classes = useStyles();
    const navigate = useNavigate();
    const user = useUser();

    // Constants
    const defaultBackdropImage = "https://cdn.prod.website-files.com/5e261bc81db8f19fa664899d/64add0eb758ddc8d390ed4a0_out-0.png"
    const backgroundImage = !!media.BackdropImage ? media.BackdropImage : defaultBackdropImage;

    // Functions
    const onPlay = () => {
        if (currentEpisode) {
            onAddToList(media, user, currentEpisode.SeasonNumber, currentEpisode.EpisodeNumber)
            navigate(`/watch/${media.Media?.TMDBID}/${currentEpisode.SeasonNumber}/${currentEpisode.EpisodeNumber}`, { state: { media, currentEpisode } });
        } else if (media.Media?.MediaType?.Name == "TV") {
            onAddToList(media, user, 1, 1)
            navigate(`/watch/${media.Media?.TMDBID}/1/1`,
                // todo get the media
                { state: { media } }
            );
        } else {

            onAddToList(media, user)
            navigate(`/watch/${media.Media?.TMDBID}`, { state: { media, currentEpisode } });
        }
    }


    return (
        <Box className={classes.modalContainer}>
            {/* Image Overlay */}
            <Box
                className={classes.imageOverlay}
                sx={{
                    backgroundImage: `url(${backgroundImage})`
                }}
            />
            {/* Title */}
            <Box className={classes.title}>
                <Typography variant="h4" fontWeight="bold">
                    {media.Media?.Title}
                </Typography>
            </Box>
            {/* Controls */}
            <Box className={classes.controls}>
                <Button
                    variant="contained"
                    color="primary"
                    className={classes.controlButton}
                    startIcon={<PlayArrow />}
                    onClick={onPlay}
                >
                    {t('button.play')}
                </Button>
                <IconButton className={`${classes.roundButton}`} aria-label={t('dictionary.addToMyList')}>
                    <Add />
                </IconButton>
                <IconButton className={`${classes.roundButton}`} aria-label={t('dictionary.rate')}>
                    <ThumbUp />
                </IconButton>

                <Tooltip title={t('dictionary.addToMyList')} arrow>
                    <IconButton onClick={() => { onAddToList(media, user, currentEpisode?.SeasonNumber, currentEpisode?.EpisodeNumber) }} className={`${classes.roundButton}`} aria-label={t('dictionary.addToMyList')}>
                        <Add />
                    </IconButton>
                </Tooltip>
                <Tooltip title={t('dictonary.rate')} arrow>
                    <IconButton className={`${classes.roundButton}`} aria-label={t('dictonary.rate')}>
                        <ThumbUp />
                    </IconButton>
                </Tooltip>
            </Box>
        </Box>
    );
};
