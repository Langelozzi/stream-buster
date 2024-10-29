import React from 'react';
import { Box, IconButton, Button, Typography } from '@mui/material';
import { PlayArrow, Add, ThumbUp } from '@mui/icons-material';
import { makeStyles } from '@mui/styles';
import { Movie } from '../../../models/movie';
import { TV } from '../../../models/tv';
import { useNavigate } from 'react-router-dom';

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
}

export const MediaDetailsModalHeader: React.FC<MediaDetailsModalHeaderProps> = ({ media }) => {
    // Hooks
    const classes = useStyles();
    const navigate = useNavigate();

    // Constants
    const defaultBackdropImage = "https://cdn.prod.website-files.com/5e261bc81db8f19fa664899d/64add0eb758ddc8d390ed4a0_out-0.png"
    const backgroundImage = !!media.BackdropImage ? media.BackdropImage : defaultBackdropImage;

    // Functions
    const onPlay = () => {
        navigate('/watch', { state: { media } });
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
                <IconButton className={`${classes.roundButton}`} aria-label="Add to My List">
                    <Add />
                </IconButton>
                <IconButton className={`${classes.roundButton}`} aria-label="Rate">
                    <ThumbUp />
                </IconButton>
            </Box>
        </Box>
    );
};
