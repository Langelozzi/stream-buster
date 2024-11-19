import React from 'react';
import { Box, IconButton, Button, Typography, Tooltip } from '@mui/material';
import { PlayArrow, Add, ThumbUp } from '@mui/icons-material';
import { makeStyles } from '@mui/styles';
import { Movie } from '../../../models/movie';
import { TV } from '../../../models/tv';
import { useNavigate } from 'react-router-dom';
import { Episode } from '../../../models/episode';
import { createCurrentlyWatching } from '../../../api/services/currentlyWatching.service';
import { CurrentlyWatching } from '../../../models/currently_watching';
import { createMedia } from '../../../api/services/media.service';
import { castToTvOrMovie } from '../../../api/services/search.service';
import { useUser } from '../../../hooks/useUser';

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
    const classes = useStyles();
    const navigate = useNavigate();
    const user = useUser();

    // Constants
    const defaultBackdropImage = "https://cdn.prod.website-files.com/5e261bc81db8f19fa664899d/64add0eb758ddc8d390ed4a0_out-0.png"
    const backgroundImage = !!media.BackdropImage ? media.BackdropImage : defaultBackdropImage;

    // Functions
    const onPlay = () => {
        if (currentEpisode) {
            navigate(`/watch/${media.Media?.TMDBID}/${currentEpisode.SeasonNumber}/${currentEpisode.EpisodeNumber}`, { state: { media, currentEpisode } });
        } else {

            navigate(`/watch/${media.Media?.TMDBID}`, { state: { media, currentEpisode } });
        }
    }

    const onAddToList = async () => {
        try {
            let mediaResponse;
            try {
                mediaResponse = await createMedia(media.Media!)
                console.log('mediaResponse', mediaResponse);

            } catch (error) {

            }

            const currentlyWatching: CurrentlyWatching = {
                MediaId: mediaResponse?.ID,
                UserID: user.user?.ID
            }

            const response = await createCurrentlyWatching(currentlyWatching)
            console.log('response', response);
        } catch (error) {
            console.error("Error addign to list")
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
                    Play
                </Button>
                <Tooltip title="Add to List" arrow>
                    <IconButton onClick={onAddToList} className={`${classes.roundButton}`} aria-label="Add to My List">
                        <Add />
                    </IconButton>
                </Tooltip>
                <Tooltip title="Like" arrow>
                    <IconButton className={`${classes.roundButton}`} aria-label="Rate">
                        <ThumbUp />
                    </IconButton>
                </Tooltip>
            </Box>
        </Box>
    );
};
