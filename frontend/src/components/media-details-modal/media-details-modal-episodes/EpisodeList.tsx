import { Box, Divider, List, Skeleton } from '@mui/material';
import { EpisodeListItem } from '../media-details-modal-episodes/EpisodeListItem';
import { TV } from '../../../models/tv';
import { Episode } from '../../../models/episode';

interface EpisodeListProps {
    tv: TV;
    episodes: Episode[] | null | undefined;
    styles: any;
    loading?: boolean;
}

export const EpisodeList: React.FC<EpisodeListProps> = ({ tv, episodes, styles, loading = false }) => {
    if (loading) {
        return (
            <List sx={styles.episodeList}>
                {[...Array(3)].map((_, index) => (
                    <Box key={index}>
                        <Divider sx={styles.divider} />
                        <Box sx={{ display: 'flex', p: 2, alignItems: 'center' }}>
                            {/* Thumbnail with play icon skeleton */}
                            <Box sx={{ position: 'relative', mr: 2, width: 160, height: 90 }}>
                                <Skeleton variant="rectangular" width={160} height={90}
                                    sx={{ bgcolor: 'rgba(255, 255, 255, 0.1)' }} />
                                <Box sx={{
                                    position: 'absolute',
                                    top: '50%',
                                    left: '50%',
                                    transform: 'translate(-50%, -50%)'
                                }}>
                                    <Skeleton variant="circular" width={40} height={40}
                                        sx={{ bgcolor: 'rgba(255, 255, 255, 0.2)' }} />
                                </Box>
                            </Box>

                            {/* Episode details skeleton */}
                            <Box sx={{ flexGrow: 1 }}>
                                <Skeleton variant="text" width="40%" height={32}
                                    sx={{ bgcolor: 'rgba(255, 255, 255, 0.1)', mb: 1 }} />
                                <Skeleton variant="text" width="80%" height={20}
                                    sx={{ bgcolor: 'rgba(255, 255, 255, 0.1)', mb: 1 }} />
                                <Skeleton variant="text" width="15%" height={16}
                                    sx={{ bgcolor: 'rgba(255, 255, 255, 0.1)' }} />
                            </Box>
                        </Box>
                    </Box>
                ))}
            </List>
        );
    }

    if (!episodes || episodes.length === 0) {
        return <Box sx={{ color: 'white', p: 2 }}>No episodes available.</Box>;
    }

    return (
        <List sx={styles.episodeList}>
            {episodes.map((episode) => (
                <Box key={episode.EpisodeTMDBID}>
                    <Divider sx={styles.divider} />
                    <EpisodeListItem tv={tv} episode={episode} />
                </Box>
            ))}
        </List>
    );
};
