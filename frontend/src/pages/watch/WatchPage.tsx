
import { useLocation, useNavigate, useParams } from "react-router-dom";
import { MediaPlayer } from "../../components/media-player/MediaPlayer";
import { TV } from "../../models/tv";
import { Movie } from '../../models/movie';
import { Box, IconButton, Typography } from "@mui/material";
import BackIcon from '@mui/icons-material/ArrowBack';
import Grid from "@mui/material/Grid2"
import { useState } from "react";

export const WatchPage = () => {
    // Hooks
    const location = useLocation();
    const navigate = useNavigate();

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
    const [media] = useState<Movie | TV | undefined>(location.state?.media);
    // const [episode, setEpisode] = useState<Episode | undefined>(location.state?.currentEpisode);
    // const media = location.state.media as TV | Movie ?? null;
    // const episode = location.state.currentEpisode as Episode ?? null;

    // Constants
    const isTV = seasonNum && episodeNum;

    // Functions
    const handleBrowseClick = () => {
        navigate(-1);
    };

    // Effects
    // Write an effect to fetch the media and episode information from api if not passed as router state

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
                            {media.Media?.Title}
                        </Typography>
                    </Grid>
                )}
                <Grid size={12} component="div">
                    {tmdbId && !isTV && (
                        <Box sx={{ display: 'flex', justifyContent: 'center' }}>
                            <MediaPlayer tmdbId={tmdbId} />
                        </Box>
                    )}
                    {tmdbId && isTV && episodeNum && seasonNum && (
                        <Box sx={{ display: 'flex', justifyContent: 'center' }}>
                            <MediaPlayer tmdbId={tmdbId} seasonNum={seasonNum} episodeNum={episodeNum} />
                        </Box>
                    )}
                </Grid>
            </Grid>
        </Box>
    );
}
