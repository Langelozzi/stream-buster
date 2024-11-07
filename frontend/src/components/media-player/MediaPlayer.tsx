import React, { useEffect } from 'react';
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

    useEffect(() => {
        // Save the original window.open function
        const originalWindowOpen = window.open;

        // Override window.open to block new tabs for external URLs
        window.open = function (url, ...args) {
            const isExternal = new URL(url as any, window.location.origin).origin !== window.location.origin;
            if (isExternal) {
                console.log('Blocked external popup:', url);
                return null; // Block the popup
            }
            // Allow internal links
            return originalWindowOpen.call(window, url, ...args);
        };

        // Clean up and restore the original window.open function on component unmount
        return () => {
            window.open = originalWindowOpen;
        };
    }, []);

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