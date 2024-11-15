import React from 'react';
import { Box } from '@mui/material';
import { API_BASE_URL } from '../../utils/constants';

interface MediaPlayerProps {
    tmdbId: number;
    seasonNum?: number;
    episodeNum?: number;
}

export const MediaPlayer: React.FC<MediaPlayerProps> = (props) => {
    const {
        tmdbId,
        seasonNum,
        episodeNum
    } = props;

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
                height: '100vh', // Full height of the viewport
                width: '100%', // Full width of the parent
                overflow: 'hidden', // Hide overflow to maintain layout
            }}
        >
            <iframe
                src={src}
                allowFullScreen
                sandbox="allow-forms allow-pointer-lock allow-same-origin allow-scripts allow-top-navigation" // Don't add allow-popups to prevent
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