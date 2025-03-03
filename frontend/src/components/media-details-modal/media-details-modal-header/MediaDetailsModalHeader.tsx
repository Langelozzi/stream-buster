import React from 'react';
import { Box, IconButton, Button, Typography, Tooltip, useMediaQuery, useTheme } from '@mui/material';
import { PlayArrow, Add, ThumbUp, Close } from '@mui/icons-material';
import { Movie } from '../../../models/movie';
import { TV } from '../../../models/tv';
import { useNavigate } from 'react-router-dom';
import { Episode } from '../../../models/episode';
import { useTranslation } from 'react-i18next';
import { useUser } from '../../../hooks/useUser';
import { onAddToList } from '../../../api/services/currentlyWatching.service';
import { useSnackbar } from '../../../hooks/useSnackBar';
import { AvailabilityInfo } from './AvailabilityInfo';

interface MediaDetailsModalHeaderProps {
    media: Movie | TV;
    currentEpisode?: Episode;
    available: number;
    onClose: () => void;
}

export const MediaDetailsModalHeader: React.FC<MediaDetailsModalHeaderProps> = ({ media, currentEpisode, available, onClose }) => {
    // Hooks
    const { t } = useTranslation();
    const navigate = useNavigate();
    const user = useUser();
    const theme = useTheme();
    const isMobile = useMediaQuery(theme.breakpoints.down('sm'));
    const { showSnackbar, SnackbarComponent } = useSnackbar()

    // Define styles as a JSON object
    const styles = {
        modalContainer: {
            position: 'relative',
            width: '100%',
            height: isMobile ? '200px' : '500px',
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
            bottom: available === 0 ? '20px' : '70px',
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
        closeButton: {
            position: 'absolute',
            top: '16px',
            right: '16px',
            zIndex: 5,
            backgroundColor: 'rgba(0, 0, 0, 0.5)',
            color: 'white',
            '&:hover': {
                backgroundColor: 'rgba(0, 0, 0, 0.7)',
            },
        },
        availabilityContainer: {
            display: 'flex',
            alignItems: 'center',
        },
        availabilityText: {
            marginLeft: 1
        }
    };


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
        <Box sx={{
            ...styles.modalContainer,
        }}>
            {/* Close Button */}
            <IconButton
                onClick={onClose}
                aria-label="close"
                size={isMobile ? "small" : "medium"}
                sx={styles.closeButton}
            >
                <Close />
            </IconButton>

            {/* Image Overlay */}
            <Box
                sx={{
                    ...styles.imageOverlay,
                    backgroundImage: `url(${backgroundImage})`,
                    backgroundPosition: isMobile ? 'center top' : 'center',
                }}
            />
            {/* Title */}
            <Box sx={{
                ...styles.title,
                left: isMobile ? '10px' : '20px'
            }}>
                <Typography variant={isMobile ? "h5" : "h4"} fontWeight="bold">
                    {media.Media?.Title}
                </Typography>
                <AvailabilityInfo available={available} />
            </Box>
            {/* Controls */}

            {available != 0 && (
                <Box sx={{
                    ...styles.controls,
                    left: isMobile ? '10px' : '20px'
                }}>
                    <Button
                        variant="contained"
                        color="primary"
                        sx={styles.controlButton}
                        startIcon={<PlayArrow />}
                        onClick={onPlay}
                        size={isMobile ? "small" : "medium"}
                    >
                        {t('button.play')}
                    </Button>

                    <Tooltip title={t('dictionary.addToMyList')} arrow>
                        <IconButton
                            onClick={onAdd}
                            sx={styles.roundButton}
                            size={isMobile ? "small" : "medium"}
                        >
                            <Add />
                        </IconButton>
                    </Tooltip>
                    <Tooltip title={t('dictonary.rate')} arrow>
                        <IconButton
                            sx={styles.roundButton}
                            size={isMobile ? "small" : "medium"}
                        >
                            <ThumbUp />
                        </IconButton>
                    </Tooltip>
                </Box>
            )}
            {SnackbarComponent}
        </Box>
    );
};
