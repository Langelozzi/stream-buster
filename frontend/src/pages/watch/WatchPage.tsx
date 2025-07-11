import { useLocation, useNavigate, useParams } from "react-router-dom";
import { MediaPlayer } from "../../components/media-player/MediaPlayer";
import {
    Box,
    IconButton,
    Link,
    Typography,
    Card
} from "@mui/material";
import BackIcon from '@mui/icons-material/ArrowBack';
import Grid from "@mui/material/Grid2";
import { useEffect, useState } from "react";
import { updateCurrentlyWatching } from "../../api/services/currentlyWatching.service";
import { CurrentlyWatching } from "../../models/currently_watching";
import { useUser } from "../../hooks/useUser";
import { getFormattedDate } from "../../utils/date.helpter";
import { Media } from "../../models/media";
import { MediaPlayerDescription } from "../../components/media-player/media-player-description/MediaPlayerDescription";
import { EpisodeList } from "../../components/media-details-modal/media-details-modal-episodes/EpisodeList";
import { TV } from "../../models/tv";
import useKeyboardShortcuts from "../../hooks/useKeyboardShortcuts";
import { Movie } from "../../models/movie";
import { Episode } from "../../models/episode";
import { getEpisodesForSeason, getTVDetails } from "../../api/services/tv.service";
import { getMovieDetails } from "../../api/services/movie.service";

export const WatchPage = () => {
    // Hooks
    const location = useLocation();
    const navigate = useNavigate();
    const user = useUser();

    // Params
    const {
        tmdbId: tmdbIdStr,
        seasonNum: seasonNumStr,
        episodeNum: episodeNumStr
    } = useParams<{ tmdbId: string, seasonNum?: string, episodeNum?: string }>();
    const tmdbId: number = Number(tmdbIdStr);
    const seasonNum: number = Number(seasonNumStr);
    const episodeNum: number = Number(episodeNumStr);

    // State
    const [media] = useState<Media | undefined>(location.state?.media);
    const [loading, setLoading] = useState<boolean>(true);
    const [currentMedia, setCurrentMedia] = useState<TV | Movie>();
    const [currentSeason, setCurrentSeason] = useState<Episode[]>();
    const [currentEpisode, setCurrentEpisode] = useState<Episode>();
    const [error, setError] = useState<string | null>(null);
    const [showAdblockMessage, setShowAdblockMessage] = useState(true);

    // Constants
    const isTV = !!seasonNum && !!episodeNum;

    // Functions
    const handleBrowseClick = () => {
        navigate(-1);
    };

    const goToNext = (): undefined => {
        const currentlyWatching: CurrentlyWatching = {
            MediaId: media?.ID,
            UserID: user.user?.ID,
            SeasonNumber: seasonNum,
            EpisodeNumber: episodeNum + 1,
            UpdatedAt: getFormattedDate()
        };
        updateCurrentlyWatching(currentlyWatching);
        navigate(`/watch/${tmdbId}/${seasonNum}/${episodeNum + 1}`);
    };

    const goToPrev = (): undefined => {
        if (episodeNum <= 1) return;
        const currentlyWatching: CurrentlyWatching = {
            MediaId: media?.ID,
            UserID: user.user?.ID,
            SeasonNumber: seasonNum,
            EpisodeNumber: episodeNum - 1,
            UpdatedAt: getFormattedDate()
        };
        updateCurrentlyWatching(currentlyWatching);
        navigate(`/watch/${tmdbId}/${seasonNum}/${episodeNum - 1}`);
    };

    const handleNext = () => {
        if (isTV) {
            goToNext();
        }
    };

    const handlePrevious = () => {
        if (isTV) {
            goToPrev();
        }
    };

    useKeyboardShortcuts({
        n: handleNext,
        p: handlePrevious,
    });

    useEffect(() => {
        const fetchMedia = async () => {
            setLoading(true);
            setError(null);
            try {
                if (isTV) {
                    const tvDetails = await getTVDetails(tmdbId);
                    setCurrentMedia(tvDetails as TV);

                    const episodes = await getEpisodesForSeason(tmdbId, seasonNum);
                    setCurrentSeason(episodes);

                    if (episodes && episodes.length >= episodeNum) {
                        setCurrentEpisode(episodes[episodeNum - 1]);
                    } else {
                        throw new Error("Episode not found");
                    }
                } else {
                    const movieDetails = await getMovieDetails(tmdbId);
                    setCurrentMedia(movieDetails as Movie);
                }
            } catch (e) {
                console.error("Error fetching media:", e);
                setError(e instanceof Error ? e.message : "Unknown error occurred");
            } finally {
                setLoading(false);
            }
        };

        fetchMedia();
    }, [tmdbId, seasonNum, episodeNum, isTV]);

    if (error) {
        return <div>Error loading media: {error}</div>;
    }

    return (
        <Box sx={{ padding: 2 }}>
            <Grid container spacing={2} alignItems="center">
                <IconButton onClick={handleBrowseClick} aria-label="browse">
                    <BackIcon sx={{ color: 'white' }} />
                </IconButton>
                {media && (
                    <Typography variant="h5" align="left" gutterBottom>
                        {media?.Title}
                    </Typography>
                )}

                <Grid size={12} component="div">
                    {showAdblockMessage && (
                        <Box sx={{ marginBottom: 2 }}>
                            <Card sx={{ position: 'relative', padding: 2, backgroundColor: '#fff8dc' }}>
                                <IconButton
                                    size="small"
                                    onClick={() => setShowAdblockMessage(false)}
                                    sx={{ position: 'absolute', top: 4, right: 4 }}
                                    aria-label="close"
                                >
                                    ×
                                </IconButton>
                                <Typography variant="body1">
                                    We are sorry to announce that Streambuster must now adhere to ad limitations.
                                    We recommend downloading the following ad blocker:{' '}
                                    <Link
                                        href="https://chromewebstore.google.com/detail/gighmmpiobklfepjocnamgkkbiglidom?utm_source=item-share-cb."
                                        target="_blank"
                                        rel="noopener noreferrer"
                                    >
                                        AdBlock (Chrome Web Store)
                                    </Link>
                                </Typography>
                            </Card>
                        </Box>
                    )}

                    {tmdbId && !isTV && (
                        <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', flexDirection: "column" }}>
                            <Box sx={{
                                width: {
                                    xs: '100%', // 100% width on extra small screens and up
                                    lg: '75%',  // 75% width on large and up
                                }
                            }}>
                                <MediaPlayer tmdbId={tmdbId} />
                            </Box>
                            <MediaPlayerDescription
                                media={currentMedia}
                                loading={loading}
                                seasonNum={seasonNum}
                                episodeNum={episodeNum}
                                goToNext={goToNext}
                                goToPrev={goToPrev}
                            />
                        </Box>
                    )}

                    {!!isTV && !!episodeNum && !!seasonNum && (
                        <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', flexDirection: "column" }}>
                            <Box sx={{
                                width: {
                                    xs: '100%', // 100% width on extra small screens and up
                                    lg: '75%',  // 75% width on large and up
                                }
                            }}>
                                <MediaPlayer tmdbId={tmdbId} seasonNum={seasonNum} episodeNum={episodeNum} />
                            </Box>
                            <MediaPlayerDescription
                                media={currentMedia}
                                currentSeason={currentSeason}
                                currentEpisode={currentEpisode}
                                seasonNum={seasonNum}
                                episodeNum={episodeNum}
                                goToNext={goToNext}
                                goToPrev={goToPrev}
                                loading={loading}
                            />
                        </Box>
                    )}
                </Grid>
            </Grid>

            {isTV && (
                <EpisodeList tv={currentMedia as TV} episodes={currentSeason} styles={{}} loading={loading} />
            )}
        </Box>
    );
};
