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
                aspectRatio: '16 / 9'
            }}
        >
            <iframe
                src={src}
                allowFullScreen
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
