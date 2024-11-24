import React from 'react';
import { Box } from '@mui/material';
import { API_BASE_URL } from '../../utils/constants';
import ControlBar from '../../pages/watch/ControlBar';

interface MediaPlayerProps {
    tmdbId: number;
    seasonNum?: number;
    episodeNum?: number;
    goToNext: () => void
    goToPrev: () => void
}

export const MediaPlayer: React.FC<MediaPlayerProps> = (props) => {
    const {
        tmdbId,
        seasonNum,
        episodeNum,
        goToNext,
        goToPrev
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
                height: '70vh', // Full height of the viewport
                overflow: 'hidden', // Hide overflow to maintain layout
                flexDirection: "column"
            }}
        >
            <iframe
                src={src}
                allowFullScreen
                sandbox="allow-forms allow-pointer-lock allow-same-origin allow-scripts allow-top-navigation" // Don't add allow-popups to prevent
                style={{
                    border: 'none',
                    width: '60%', // Set width to 80% of the parent
                    height: '90%', // Set height to 80% of the viewport
                    borderRadius: '8px', // Optional: Add rounded corners
                }}
            />
            <ControlBar goToNext={goToNext} goToPrev={goToPrev}></ControlBar>
        </Box>
    );
};
