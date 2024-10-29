import React, { useEffect, useState } from 'react';
import {
    Box,
    Typography,
    Divider,
    List,
} from '@mui/material';
import { makeStyles } from '@mui/styles';
import { TV } from '../../models/tv';
import { Movie } from '../../models/movie';
import { MediaDetailsModalHeader } from './media-details-modal-header/MediaDetailsModalHeader';
import { MediaDetailsModalDesc } from './media-details-modal-desc/MediaDetailsModalDesc';
import { getTVDetails } from '../../api/services/tv';

// Defining styles using makeStyles
const useStyles = makeStyles({
    overlay: {
        position: 'fixed',
        top: 0,
        left: 0,
        width: '100%',
        height: '100%',
        backgroundColor: 'rgba(0, 0, 0, 0.5)',
        zIndex: 1000,
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
    },
    modalContainer: {
        width: '100%',
        maxWidth: 1200,
        backgroundColor: 'black',
        color: 'white',
        borderRadius: 8,
        overflow: 'hidden',
        margin: '0 auto',
    },
    header: {
        position: 'relative',
        height: 400,
        backgroundImage: `url(https://image.tmdb.org/t/p/w500/pa4UM9lTaYLhi7RuBuPOejAoNfu.jpg)`,
        backgroundSize: 'cover',
        backgroundPosition: 'center',
    },
    headerContent: {
        position: 'absolute',
        bottom: 0,
        left: 0,
        padding: 24,
        width: '100%',
        background: 'linear-gradient(to top, rgba(0, 0, 0, 0.8), transparent)',
    },
    buttonContainer: {
        display: 'flex',
        marginTop: 16,
        gap: 16,
    },
    episodeList: {
        width: '100%',
        backgroundColor: 'black',
    },
    listItem: {
        padding: '16px 0',
    },
    episodeTitleContainer: {
        display: 'flex',
        justifyContent: 'space-between',
    },
    divider: {
        borderColor: 'white',
        marginBottom: 16,
    },
});

interface MediaDetailsModalProps {
    media: Movie | TV;
    isOpen: boolean;
    onClose: () => void;
}

/*
NOTES:
- Need to create backend endpoints for fetching details of movie or show + seasons
- Use the extra data to dynamically populate modal
- TV shows don't work to play yet
- I want to make the modal have it's own url route so that we can copy paste it
- When user goes back from media player it should remember state of browse page
- When scrolling with modal open it should scroll modal not background
*/
const MediaDetailsModal: React.FC<MediaDetailsModalProps> = (props) => {
    // Props
    const {
        media,
        isOpen,
        onClose
    } = props;

    // Hooks
    const classes = useStyles();

    // Constants
    const isTV = media.Media?.MediaType?.Name.toLowerCase() === 'tv';

    // States
    const [detailedMedia, setDetailedMedia] = useState<Movie | TV | null>(null);

    // Effects
    useEffect(() => { // Runs when component loads the first time
        if (isTV && isOpen) {
            const fetchDetailedTV = async () => {
                const tv: TV = await getTVDetails(media.Media?.TMDBID!);
                console.log(tv);
                setDetailedMedia(tv);
            }
            fetchDetailedTV();
        } else {
            setDetailedMedia(media);
        }
    }, [isOpen])

    // Render nothing if modal is not open
    if (!isOpen) return null;
    return (
        <Box onClick={onClose} className={classes.overlay}>
            <Box onClick={(e) => e.stopPropagation()} className={classes.modalContainer}>
                {/* Header Section with Background Image */}
                {detailedMedia && (
                    <MediaDetailsModalHeader media={detailedMedia} />
                )}


                <Box p={6}>
                    {detailedMedia && (
                        <MediaDetailsModalDesc media={detailedMedia} />
                    )}

                    {/* Episode List Section (should be conditionally rendered if it's a tv show)*/}
                    <Typography variant="h5" mb={2}>
                        Episodes
                    </Typography>
                    <Divider className={classes.divider} />

                    <List className={classes.episodeList}>
                        {/* Example for iterating over episodes if media has them */}
                        {/* {media.episodes?.map((episode) => (
                            <ListItem key={episode.number} className={classes.listItem}>
                                <ListItemAvatar>
                                    <Avatar
                                        variant="square"
                                        src={episode.thumbnailUrl}
                                        sx={{ width: 100, height: 60 }}
                                    />
                                </ListItemAvatar>
                                <ListItemText
                                    primary={
                                        <Box className={classes.episodeTitleContainer}>
                                            <Typography variant="body1" fontWeight="bold">
                                                {`${episode.number}. ${episode.title}`}
                                            </Typography>
                                            <Typography variant="body2">{episode.duration}</Typography>
                                        </Box>
                                    }
                                    secondary={episode.description}
                                />
                            </ListItem>
                        ))} */}
                    </List>
                </Box>
            </Box>
        </Box>
    );
};

export default MediaDetailsModal;
