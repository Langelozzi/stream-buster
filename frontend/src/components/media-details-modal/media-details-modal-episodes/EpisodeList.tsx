import { Box, Divider, List } from '@mui/material';
import { EpisodeListItem } from '../media-details-modal-episodes/EpisodeListItem';
import { TV } from '../../../models/tv';
import { Episode } from '../../../models/episode';


interface EpisodeListProps {
    tv: TV;
    episodes: Episode[] | null | undefined;
}

export const EpisodeList: React.FC<EpisodeListProps> = ({ tv, episodes }) => {
    const styles = {
        episodeList: {
            width: '100%',
            backgroundColor: 'black',
        },
        divider: {
            borderColor: 'gray',
            my: 1,
        },
    };

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

