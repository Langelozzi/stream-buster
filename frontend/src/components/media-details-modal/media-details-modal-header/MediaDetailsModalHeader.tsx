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
        backgroundImage: `url(https://image.tmdb.org/t/p/w500/pa4UM9lTaYLhi7RuBuPOejAoNfu.jpg)`,
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
    const classes = useStyles();
    const navigate = useNavigate();

    const onPlay = () => {
        navigate('/watch', { state: { media } });
    }

    return (
        <Box className={classes.modalContainer}>
            {/* Image Overlay */}
            <Box className={classes.imageOverlay} />
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
