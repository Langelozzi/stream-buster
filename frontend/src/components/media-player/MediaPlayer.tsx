// MediaPlayer.tsx
import React from 'react';
import { Box } from '@mui/material';
import { API_BASE_URL } from '../../utils/constants';

interface MediaPlayerProps {
    tmdbId: number;
    seasonNum?: number;
    episodeNum?: number;
}

export const MediaPlayer: React.FC<MediaPlayerProps> = ({ tmdbId, seasonNum, episodeNum }) => {
    const src = seasonNum && episodeNum
        ? `${API_BASE_URL}/cdn/tv/${tmdbId}/${seasonNum}/${episodeNum}`
        : `${API_BASE_URL}/cdn/movie/${tmdbId}`;

    return (
        <Box
            sx={{
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                width: '100%',
                height: '70vh',
            }}
        >
            <iframe
                src={src}
                allowFullScreen
                sandbox="allow-forms allow-pointer-lock allow-same-origin allow-scripts allow-top-navigation"
                style={{
                    border: 'none',
                    width: '100%',
                    height: '100%',
                    borderRadius: '8px',
                }}
            />
        </Box>
    );
};
