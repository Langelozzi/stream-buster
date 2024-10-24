import React from 'react';
import { Box } from '@mui/material';

interface MediaPlayerProps {
    tmdbId: number;
    seasonNum?: number;
    episodeNum?: number;
}

const BASE_URL = import.meta.env.VITE_API_BASE_URL

const MediaPlayer: React.FC<MediaPlayerProps> = (props) => {
    const {
        tmdbId,
        seasonNum,
        episodeNum
    } = props;

    const src = seasonNum && episodeNum ?
        `${BASE_URL}/cdn/tv/${tmdbId}/${seasonNum}/${episodeNum}`
        :
        `${BASE_URL}/cdn/movie/${tmdbId}`

    return (
        <Box
            sx={{
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                height: '100vh', // Full height of the viewport
                width: '100%', // Full width of the parent
                overflow: 'hidden', // Hide overflow to maintain layout
            }}
        >
            <iframe
                src={src}
                allowFullScreen
                style={{
                    border: 'none',
                    width: '80%', // Set width to 80% of the parent
                    height: '80%', // Set height to 80% of the viewport
                    borderRadius: '8px', // Optional: Add rounded corners
                }}
            />
        </Box>
    );
};

export default MediaPlayer;