import React from 'react';
import { Box, IconButton, Button, Typography, Tooltip } from '@mui/material';
import { PlayArrow, Add, ThumbUp, CheckCircle, Cancel } from '@mui/icons-material';
import { Movie } from '../../../models/movie';
import { TV } from '../../../models/tv';
import { useNavigate } from 'react-router-dom';
import { Episode } from '../../../models/episode';
import { useTranslation } from 'react-i18next';
import { useUser } from '../../../hooks/useUser';
import { onAddToList } from '../../../api/services/currentlyWatching.service';
import { useSnackbar } from '../../../hooks/useSnackBar';

interface MediaDetailsModalHeaderProps {
    media: Movie | TV;
    currentEpisode?: Episode;
    available: boolean;
}

export const MediaDetailsModalHeader: React.FC<MediaDetailsModalHeaderProps> = ({ media, currentEpisode, available }) => {
    const styles = {
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
            bottom: available ? '70px' : '20px', // adjust this value if needed
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
    }

    // Hooks
    const { t } = useTranslation();
    const navigate = useNavigate();
    const user = useUser();

    const { showSnackbar, SnackbarComponent } = useSnackbar()

    // Constants
    const defaultBackdropImage = "https://cdn.prod.website-files.com/5e261bc81db8f19fa664899d/64add0eb758ddc8d390ed4a0_out-0.png"
    const backgroundImage = !!media.BackdropImage ? media.BackdropImage : defaultBackdropImage;

    // Functions
    const onPlay = async () => {
        if (currentEpisode) {
            const mediaResponse = await onAddToList(media, user, currentEpisode.SeasonNumber, currentEpisode.EpisodeNumber)
            media.MediaID = mediaResponse.ID
            navigate(`/watch/${media.Media?.TMDBID}/${currentEpisode.SeasonNumber}/${currentEpisode.EpisodeNumber}`, { state: { media: mediaResponse, currentEpisode } });
        } else if (media.Media?.MediaType?.Name == "TV") {
            const mediaResponse = await onAddToList(media, user, 1, 1)
            media.MediaID = mediaResponse.ID;
            navigate(`/watch/${media.Media?.TMDBID}/1/1`,
                { state: { media: mediaResponse } }
            );
        } else {
            const mediaResponse = await onAddToList(media, user);
            media.MediaID = mediaResponse.ID;
            navigate(`/watch/${media.Media?.TMDBID}`, { state: { media: mediaResponse, currentEpisode } });
        }
    }

    const onAdd = async () => {
        try {
            await onAddToList(media, user, currentEpisode?.SeasonNumber, currentEpisode?.EpisodeNumber)
            showSnackbar("Successfully added to watchlist")
        } catch (error) {
            showSnackbar("Error added to watchlist")
        }
    }


    return (
        <Box sx={styles.modalContainer}>
            {/* Image Overlay */}
            <Box
                sx={{
                    backgroundImage: `url(${backgroundImage})`,
                    ...styles.imageOverlay
                }}
            />
            {/* Title */}
            <Box sx={styles.title}>
                <Typography variant="h4" fontWeight="bold">
                    {media.Media?.Title}
                </Typography>
                {available ? (
                    <Typography>
                        <CheckCircle />
                        Available
                    </Typography>
                ) : (
                    <Typography>
                        <Cancel />
                        Unavailable
                    </Typography>
                )}
            </Box>

            {/* Controls */}
            {available && (
                <Box sx={styles.controls}>
                    <Button
                        variant="contained"
                        color="primary"
                        sx={styles.controlButton}
                        startIcon={<PlayArrow />}
                        onClick={onPlay}
                    >
                        {t('button.play')}
                    </Button>

                    <Tooltip title={t('dictionary.addToMyList')} arrow>
                        <IconButton onClick={onAdd} sx={styles.roundButton} aria-label={t('dictionary.addToMyList')}>
                            <Add />
                        </IconButton>
                    </Tooltip>
                    <Tooltip title={t('dictonary.rate')} arrow>
                        <IconButton sx={styles.roundButton} aria-label={t('dictonary.rate')}>
                            <ThumbUp />
                        </IconButton>
                    </Tooltip>
                </Box>
            )}
            {SnackbarComponent}
        </Box>
    );
};
