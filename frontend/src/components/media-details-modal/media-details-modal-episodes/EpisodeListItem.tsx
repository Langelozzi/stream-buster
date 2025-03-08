import React from "react";
import { Episode } from "../../../models/episode";
import { Avatar, Box, ListItem, ListItemAvatar, ListItemText, Typography, IconButton, useMediaQuery, useTheme } from "@mui/material";
import PlayCircleFilledWhiteIcon from '@mui/icons-material/PlayCircleOutline';
import { TV } from "../../../models/tv";
import { useNavigate } from "react-router-dom";
import { useTranslation } from "react-i18next";
import { onAddToList } from "../../../api/services/currentlyWatching.service";
import { useUser } from "../../../hooks/useUser";

// Define styles as a JSON object
const styles = {
    listItem: {
        padding: '12px 0',
        display: 'flex',
        flexDirection: 'row',
        alignItems: 'flex-start',
        gap: '16px',
    },
    avatarContainer: {
        position: 'relative', // Set relative position to contain the overlay
    },
    playButton: {
        position: 'absolute',
        top: '50%',
        left: '50%',
        transform: 'translate(-50%, -50%)', // Center the play button
        backgroundColor: 'rgba(0, 0, 0, 0.6)', // Optional: Add a background for visibility
        color: 'white',
        borderRadius: '50%',
    }
};

interface EpisodeListItemProps {
    episode: Episode;
    tv: TV;
}

export const EpisodeListItem: React.FC<EpisodeListItemProps> = (props) => {
    const user = useUser()
    const { episode, tv } = props;

    const { t } = useTranslation();
    const navigate = useNavigate();
    const theme = useTheme();
    const isMobile = useMediaQuery(theme.breakpoints.down('sm'));

    const onPlayEpisode = () => {
        onAddToList(tv, user, episode.SeasonNumber, episode.EpisodeNumber)
        navigate(`/watch/${tv.Media?.TMDBID}/${episode.SeasonNumber}/${episode.EpisodeNumber}`, { state: { media: tv, currentEpisode: episode } });
    }

    return (
        <ListItem key={episode.EpisodeTMDBID}
            sx={{
                ...styles.listItem,
                flexDirection: isMobile ? 'column' : 'row',
                alignItems: isMobile ? 'center' : 'flex-start',
                padding: isMobile ? '16px 8px' : '12px 0',
            }}
        >
            <ListItemAvatar>
                <Box sx={{
                    ...styles.avatarContainer,
                    width: isMobile ? '100%' : 200,
                    height: isMobile ? 'auto' : 100,
                    aspectRatio: isMobile ? '16/9' : 'auto'
                }}>
                    <Avatar
                        variant="square"
                        src={episode.StillPath}
                        sx={{ width: '100%', height: '100%' }}
                    />
                    <IconButton sx={styles.playButton} onClick={onPlayEpisode}>
                        <PlayCircleFilledWhiteIcon fontSize={isMobile ? "medium" : "large"} />
                    </IconButton>
                </Box>
            </ListItemAvatar>
            <ListItemText
                primary={
                    <Box ml={isMobile ? 0 : 2} mt={isMobile ? 2 : 0} sx={{
                        display: 'flex',
                        flexDirection: 'column',
                        gap: 1,
                        width: '100%'
                    }}>
                        <Typography variant="body1" fontWeight="bold">
                            {`${episode.EpisodeNumber}. ${episode.Name}`}
                        </Typography>
                        <Typography variant="body2">{episode.Overview}</Typography>
                        <Typography variant="body2" color="grey">{episode.Runtime}{t('dictionary.minuteLetter')}</Typography>
                    </Box>
                }
            />
        </ListItem>
    )
}
