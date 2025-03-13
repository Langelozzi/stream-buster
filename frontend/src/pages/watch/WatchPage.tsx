
import { useLocation, useNavigate, useParams } from "react-router-dom";
import { MediaPlayer } from "../../components/media-player/MediaPlayer";
import { Box, IconButton, Typography } from "@mui/material";
import BackIcon from '@mui/icons-material/ArrowBack';
import Grid from "@mui/material/Grid2"
import { useState } from "react";
import { updateCurrentlyWatching } from "../../api/services/currentlyWatching.service";
import { CurrentlyWatching } from "../../models/currently_watching";
import { useUser } from "../../hooks/useUser";
import { getFormattedDate } from "../../utils/date.helpter";
import { Media } from "../../models/media";
import { MediaPlayerDescription } from "../../components/media-player/media-player-description/MediaPlayerDescription";

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
        }
        updateCurrentlyWatching(currentlyWatching)
        navigate(`/watch/${tmdbId}/${seasonNum}/${episodeNum + 1}`)
    }

    const goToPrev = (): undefined => {
        if (episodeNum <= 1)
            return;
        const currentlyWatching: CurrentlyWatching = {
            MediaId: media?.ID,
            UserID: user.user?.ID,
            SeasonNumber: seasonNum,
            EpisodeNumber: episodeNum - 1,
            UpdatedAt: getFormattedDate()
        }
        updateCurrentlyWatching(currentlyWatching)
        navigate(`/watch/${tmdbId}/${seasonNum}/${episodeNum - 1}`)
    }

    return (
        <Box sx={{ padding: 2 }}>
            <Grid container spacing={2} alignItems="center">
                <Grid size={1} component="div" sx={{ textAlign: 'left' }}> {/* Browse Button Section */}
                    <IconButton onClick={handleBrowseClick} aria-label="browse">
                        <BackIcon sx={{ color: 'white' }} />
                    </IconButton>
                </Grid>
                {media && (
                    <Grid size={11} component="div"> {/* Title Section */}
                        <Typography variant="h5" align="left" gutterBottom>
                            {media?.Title}
                        </Typography>
                    </Grid>
                )}
                <Grid size={12} component="div">
                    {tmdbId && !isTV && (
                        <Box sx={{ display: 'flex', justifyContent: 'center', flexDirection: "column" }}>
                            <MediaPlayer tmdbId={tmdbId} />
                            <MediaPlayerDescription
                                tmdbId={tmdbId}
                                seasonNum={seasonNum}
                                episodeNum={episodeNum}
                                goToNext={goToNext}
                                goToPrev={goToPrev}
                            />
                        </Box>
                    )}
                    {!!isTV && !!episodeNum && !!seasonNum && (
                        <Box sx={{ display: 'flex', justifyContent: 'center', flexDirection: "column" }}>
                            <MediaPlayer tmdbId={tmdbId} seasonNum={seasonNum} episodeNum={episodeNum} />

                            <MediaPlayerDescription
                                tmdbId={tmdbId}
                                seasonNum={seasonNum}
                                episodeNum={episodeNum}
                                goToNext={goToNext}
                                goToPrev={goToPrev}
                            />
                        </Box>
                    )}
                </Grid>

            </Grid>
        </Box>
    );
}
