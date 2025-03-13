import React, { useEffect, useState } from 'react';
import { Box, Chip, Paper, Typography } from '@mui/material';
import Skeleton from '@mui/material/Skeleton';
import { API_BASE_URL } from '../../utils/constants';
import ControlBar from '../../pages/watch/ControlBar';
import { getEpisodesForSeason, getTVDetails } from '../../api/services/tv.service';
import { getMovieDetails } from '../../api/services/movie.service';
import { TV } from '../../models/tv';
import { Movie } from '../../models/movie';
import { Episode } from '../../models/episode';
import { ExpandableText } from '../expandable-text/ExpandableText';
import { Media } from '../../models/media';
import useKeyboardShortcuts from '../../hooks/useKeyboardShortcuts';

// Format functions
const formatDate = (dateString: String) => {
    if (!dateString) return '';
    const date = new Date(dateString);
    return date.getFullYear().toString();
};

const formatRuntime = (minutes: number) => {
    if (!minutes) return '';
    const hours = Math.floor(minutes / 60);
    const mins = minutes % 60;
    return hours > 0 ? `${hours}h ${mins}m` : `${mins}m`;
};

interface MediaPlayerProps {
    tmdbId: number;
    seasonNum?: number;
    episodeNum?: number;
    goToNext?: () => void
    goToPrev?: () => void
}

export const MediaPlayer: React.FC<MediaPlayerProps> = (props) => {
    const {
        tmdbId,
        seasonNum,
        episodeNum,
        goToNext,
        goToPrev,
    } = props;

    const isTv = seasonNum!! && episodeNum!!;
    const [loading, setLoading] = useState<boolean>(true)
    const [media, setMedia] = useState<TV | Movie>();
    const [, setCurrentSeason] = useState<Episode[]>()
    const [currentEpisode, setCurrentEpisode] = useState<Episode>();

    const handleNext = () => {
        if (isTv)
            goToNext!()
    };

    const handlePrevious = () => {
        if (isTv)
            goToPrev!()
    };

    useKeyboardShortcuts({
        n: handleNext,
        p: handlePrevious,
    });

    useEffect(() => {
        const fetchMedia = async () => {
            setLoading(true);
            if (isTv) {
                const m = await getTVDetails(tmdbId);
                console.log('m', m);
                setMedia(m as TV);

                const s = await getEpisodesForSeason(tmdbId, seasonNum);
                console.log('s', s);
                setCurrentSeason(s);
                try {
                    setCurrentEpisode(s[episodeNum - 1])
                } catch (e) {
                    console.error("error getting season and episode:", e)
                }

            } else {
                const m = await getMovieDetails(tmdbId);
                console.log('m', m);
                setMedia(m as Movie);
            }
            setLoading(false);
        }
        fetchMedia()
    }, [])

    const src = seasonNum && episodeNum ?
        `${API_BASE_URL}/cdn/tv/${tmdbId}/${seasonNum}/${episodeNum}`
        :
        `${API_BASE_URL}/cdn/movie/${tmdbId}`

    return (
        <Box
            sx={{
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                height: '100vh',
                overflow: 'hidden',
                flexDirection: "column"
            }}
        >
            <iframe
                src={src}
                allowFullScreen
                sandbox="allow-forms allow-pointer-lock allow-same-origin allow-scripts allow-top-navigation"
                style={{
                    border: 'none',
                    width: '60%',
                    height: '90%',
                    borderRadius: '8px',
                }}
            />

            <Box sx={{
                display: "flex",
                width: '60%',
                justifyContent: 'space-between',
                alignItems: 'center',
                mt: 2
            }}>
                {loading ? (
                    <Skeleton width={400} height={40} />
                ) : (
                    <Typography variant="h5" sx={{ fontWeight: 'medium' }}>
                        {isTv ?
                            `${media?.Media?.Title} - S${seasonNum.toString().padStart(2, '0')}E${episodeNum.toString().padStart(2, '0')} - ${currentEpisode?.Name}` :
                            `${media?.Media?.Title}`}
                    </Typography>
                )}

                {goToNext && (
                    <ControlBar goToNext={goToNext!} goToPrev={goToPrev!} />
                )}
            </Box>

            {loading ? (
                <Skeleton variant="rectangular" width="60%" height={150} sx={{ mt: 1 }} />
            ) : (
                <Box sx={{ width: '60%', mt: 1 }}>

                    <Paper
                        elevation={1}
                        sx={{
                            mt: 2,
                            p: 2,
                            backgroundColor: 'rgba(255, 255, 255, 0.05)',
                            backdropFilter: 'blur(5px)',
                            display: 'flex',
                            flexDirection: { xs: 'column', md: 'row' },
                            gap: 2
                        }}
                    >
                        <Typography variant="body2" color="text.secondary">
                            {isTv && formatDate(media?.FirstAirDate) + " - " + formatRuntime(currentEpisode?.Runtime)}
                        </Typography>
                        <Box sx={{ flex: 3 }}>
                            {isTv ? (
                                <Typography variant="body1">
                                    {currentEpisode?.Overview && currentEpisode.Overview.length > 150 ? (
                                        <ExpandableText text={currentEpisode.Overview} maxLength={150} />
                                    ) : (
                                        currentEpisode?.Overview
                                    )}
                                </Typography>
                            ) : (
                                <Typography variant="body1">
                                    {media?.Media?.Overview && media.Media.Overview.length > 150 ? (
                                        <ExpandableText text={media.Media.Overview} maxLength={150} />
                                    ) : (
                                        media?.Media?.Overview
                                    )}
                                </Typography>
                            )}
                        </Box>

                        <Box sx={{ flex: 1 }}>
                            <Typography variant="subtitle2" gutterBottom>Genres</Typography>
                            <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 1 }}>
                                {media?.Media?.Genres?.map((genre) => (
                                    <Chip
                                        key={genre.Name}
                                        label={genre.Name}
                                        size="small"
                                        sx={{
                                            backgroundColor: 'rgba(255, 255, 255, 0.1)',
                                            '&:hover': { backgroundColor: 'rgba(255, 255, 255, 0.2)' }
                                        }}
                                    />
                                ))}
                            </Box>
                        </Box>
                    </Paper>

                    {/*
                    <Paper
                        elevation={1}
                        sx={{
                            mt: 3,
                            p: 2,
                            backgroundColor: 'rgba(255, 255, 255, 0.05)',
                            backdropFilter: 'blur(5px)',
                        }}
                    >
                        <Typography variant="h6" gutterBottom>Comments</Typography>
                        <Typography color="text.secondary">No comments yet. Be the first to comment!</Typography>
                    </Paper>
                    */}
                </Box>
            )}
        </Box>
    );
};
