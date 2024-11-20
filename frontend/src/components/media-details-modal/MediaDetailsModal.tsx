import React, { useCallback, useEffect, useState } from 'react';
import {
    Box,
} from '@mui/material';
import { makeStyles } from '@mui/styles';
import { TV } from '../../models/tv';
import { Movie } from '../../models/movie';
import { MediaDetailsModalHeader } from './media-details-modal-header/MediaDetailsModalHeader';
import { MediaDetailsModalDescTV } from './media-details-modal-desc/MediaDetailsModalDescTV';
import { MediaDetailsModalDescMovie } from './media-details-modal-desc/MediaDetailsModalDescMovie';
import { getEpisodesForSeason, getTVDetails } from '../../api/services/tv.service';
import { getMovieDetails } from '../../api/services/movie.service';
import { Episode } from '../../models/episode';
import { MediaDetailsModalEpisodes } from './media-details-modal-episodes/MediaDetailsModalEpisodes';
import { Season } from '../../models/season';

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
        overflowY: 'auto', // Enable vertical scrolling in the modal
        maxHeight: '90vh',
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
    currentSeasonNumber?: number | undefined;
    currentEpisodeNumber?: number | undefined;
}

/*
NOTES:
- When user goes back from media player it should remember state of browse page
*/
const MediaDetailsModal: React.FC<MediaDetailsModalProps> = (props) => {
    // Props
    const {
        media,
        isOpen,
        onClose,
        currentSeasonNumber = 1,
        currentEpisodeNumber = 1,
    } = props;


    // Hooks
    const classes = useStyles();

    // Constants
    const isTV = media.Media?.MediaType?.Name.toLowerCase() === 'tv';

    // States
    const [detailedMedia, setDetailedMedia] = useState<Movie | TV | null>(null);
    const [currentSeason, setCurrentSeason] = useState<Season | null>(null);
    const [episodes, setEpisodes] = useState<Episode[] | null>(null);
    const [currentEpisode, setCurrentEpisode] = useState<Episode | null>(null);

    // Functions
    const fetchDetailedTV = async () => {
        const tv: TV = await getTVDetails(media.Media?.TMDBID!);

        setDetailedMedia(tv);
    }

    const fetchDetailedMovie = async () => {
        const movie: Movie = await getMovieDetails(media.Media?.TMDBID!);

        setDetailedMedia(movie);
    }

    const determineCurrentSeason = () => {
        const currentSeason: Season = (detailedMedia as TV).Seasons.filter(season => season.SeasonNumber === currentSeasonNumber)[0];
        setCurrentSeason(currentSeason);
    }

    const fetchEpisodesForCurrentSeason = async () => {
        const episodes: Episode[] = await getEpisodesForSeason(media.Media?.TMDBID!, currentSeason?.SeasonNumber!);

        setEpisodes(episodes);
    }

    // Callbacks
    const determineCurrentEpisode = useCallback(() => {
        if (!episodes || currentEpisode) return;

        setCurrentEpisode(episodes[currentEpisodeNumber - 1])
    }, [episodes]);

    // Effects
    useEffect(() => { // Runs when modal component is opened
        if (!isOpen) {
            document.body.style.overflow = 'unset'; // Restore body scroll
            return;
        }

        // Fetch the details of the media clicked
        if (isTV) {
            fetchDetailedTV();
        } else {
            fetchDetailedMovie();
        }

        // Prevent body from scrolling when modal open
        document.body.style.overflow = 'hidden';
        return () => {
            document.body.style.overflow = 'unset'; // Restore body scroll
        }
    }, [isOpen, media]);

    useEffect(() => {
        if (!isOpen) return;
        if (isTV && detailedMedia) {
            determineCurrentSeason();
        }
    }, [detailedMedia])

    useEffect(() => {
        if (!isOpen) return;
        if (currentSeason) {
            fetchEpisodesForCurrentSeason();
        }
    }, [currentSeason])

    useEffect(() => {
        if (!isOpen) return;
        if (isTV) {
            determineCurrentEpisode();
        }
    }, [episodes])

    // Render nothing if modal is not open
    if (!isOpen) return null;
    return (
        <Box onClick={onClose} className={classes.overlay}>
            <Box onClick={(e) => e.stopPropagation()} className={classes.modalContainer}>
                {/* Header Section with Background Image (will need to pass current episode in for tv shows) */}
                {detailedMedia && (
                    <MediaDetailsModalHeader media={detailedMedia} currentEpisode={currentEpisode ?? undefined} />
                )}

                <Box p={6}>
                    {detailedMedia && isTV && currentEpisode && (
                        <MediaDetailsModalDescTV tv={detailedMedia as TV} currentEpisode={currentEpisode} />
                    )}
                    {detailedMedia && !isTV && (
                        <MediaDetailsModalDescMovie movie={detailedMedia as Movie} />
                    )}

                    {/* Episode List Section (should be conditionally rendered if it's a tv show)*/}
                    {detailedMedia && isTV && episodes && currentSeason && (
                        <MediaDetailsModalEpisodes
                            tv={detailedMedia as TV}
                            episodes={episodes}
                            currentSeason={currentSeason}
                            setCurrentSeason={setCurrentSeason}
                        />
                    )}
                </Box>
            </Box>
        </Box>
    );
};

export default MediaDetailsModal;
