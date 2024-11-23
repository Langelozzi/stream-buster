import React from "react";
import { Episode } from "../../../models/episode";
import { Avatar, Box, ListItem, ListItemAvatar, ListItemText, Typography, IconButton } from "@mui/material";
import { makeStyles } from "@mui/styles";
import PlayCircleFilledWhiteIcon from '@mui/icons-material/PlayCircleOutline';
import { TV } from "../../../models/tv";
import { useNavigate } from "react-router-dom";
import { onAddToList } from "../../../api/services/currentlyWatching.service";
import { useUser } from "../../../hooks/useUser";

const useStyles = makeStyles({
    listItem: {
        padding: '12px 0',
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
});

interface EpisodeListItemProps {
    episode: Episode;
    tv: TV;
}

export const EpisodeListItem: React.FC<EpisodeListItemProps> = (props) => {
    const user = useUser()
    const { episode, tv } = props;

    const classes = useStyles();
    const navigate = useNavigate();

    const onPlayEpisode = () => {
        onAddToList(tv, user, episode.SeasonNumber, episode.EpisodeNumber)
        navigate(`/watch/${tv.Media?.TMDBID}/${episode.SeasonNumber}/${episode.EpisodeNumber}`, { state: { media: tv, currentEpisode: episode } });
    }

    return (
        <ListItem key={episode.EpisodeTMDBID} className={classes.listItem}>
            <ListItemAvatar>
                <Box className={classes.avatarContainer} sx={{ width: 200, height: 100 }}>
                    <Avatar
                        variant="square"
                        src={episode.StillPath}
                        sx={{ width: '100%', height: '100%' }}
                    />
                    <IconButton className={classes.playButton} onClick={onPlayEpisode}>
                        <PlayCircleFilledWhiteIcon fontSize="large" />
                    </IconButton>
                </Box>
            </ListItemAvatar>
            <ListItemText
                primary={
                    <Box ml={2} sx={{ display: 'flex', flexDirection: 'column', gap: 1 }}>
                        <Typography variant="body1" fontWeight="bold">
                            {`${episode.EpisodeNumber}. ${episode.Name}`}
                        </Typography>
                        <Typography variant="body2">{episode.Overview}</Typography>
                        <Typography variant="body2" color="grey">{episode.Runtime}m</Typography>
                    </Box>
                }
            />
        </ListItem>
    )
}
