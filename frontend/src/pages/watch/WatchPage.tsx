
import { useLocation, useNavigate } from "react-router-dom";
import { MediaPlayer } from "../../components/media-player/MediaPlayer";
import { useUser } from "../../hooks/useUser";
import { TV } from "../../models/tv";
import { Movie } from '../../models/movie';
import { Episode } from "../../models/episode";
import { Box, IconButton, Typography } from "@mui/material";
import BackIcon from '@mui/icons-material/ArrowBack';
import Grid from "@mui/material/Grid2"

export const WatchPage = () => {
    const user = useUser();
    const location = useLocation();
    const navigate = useNavigate();

    const media = location.state.media as TV | Movie;
    const episode = location.state.currentEpisode as Episode ?? null;
    const tmdbId = media.Media?.TMDBID;

    const isTV = media.Media?.MediaType?.Name.toLowerCase() === 'tv';

    const handleBrowseClick = () => {
        navigate('/browse');
    };

    return (
        <Box sx={{ padding: 2 }}>
            <Grid container spacing={2} alignItems="center">
                <Grid size={1} component="div" sx={{ textAlign: 'left' }}> {/* Browse Button Section */}
                    <IconButton onClick={handleBrowseClick} aria-label="browse">
                        <BackIcon sx={{ color: 'white' }} />
                    </IconButton>
                </Grid>
                <Grid size={11} component="div"> {/* Title Section */}
                    <Typography variant="h5" align="left" gutterBottom>
                        {media.Media?.Title}
                    </Typography>
                </Grid>
                <Grid size={12} component="div">
                    {tmdbId && !isTV && (
                        <Box sx={{ display: 'flex', justifyContent: 'center' }}>
                            <MediaPlayer tmdbId={tmdbId} />
                        </Box>
                    )}
                    {tmdbId && isTV && episode && (
                        <Box sx={{ display: 'flex', justifyContent: 'center' }}>
                            <MediaPlayer tmdbId={tmdbId} seasonNum={episode.SeasonNumber} episodeNum={episode.EpisodeNumber} />
                        </Box>
                    )}
                </Grid>
            </Grid>
        </Box>
    );
}
