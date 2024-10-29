import React from "react";
import { Episode } from "../../../models/episode";
import { Typography, Divider, List, Avatar, Box, ListItem, ListItemAvatar, ListItemText } from "@mui/material";
import { makeStyles } from '@mui/styles';
import { EpisodeListItem } from "./EpisodeListItem";
import { TV } from "../../../models/tv";

const useStyles = makeStyles({
    episodeList: {
        width: '100%',
        backgroundColor: 'black',
    },
    divider: {
        borderColor: 'white',
        marginBottom: 16,
    },
    episodeDivider: {
        borderColor: 'grey',
        marginTop: 24
    }
});

interface MediaDetailsModalEpisodesProps {
    episodes: Episode[];
    tv: TV;
}

export const MediaDetailsModalEpisodes: React.FC<MediaDetailsModalEpisodesProps> = (props) => {
    // Props
    const {
        episodes,
        tv
    } = props;

    // Hooks
    const classes = useStyles();

    return (
        <>
            <Typography variant="h5">
                Episodes
            </Typography>

            <List className={classes.episodeList}>
                {/* Example for iterating over episodes if media has them */}
                {episodes?.map((episode) => (
                    <Box key={episode.EpisodeTMDBID}>
                        <Divider className={classes.episodeDivider} />
                        <EpisodeListItem tv={tv} episode={episode} />
                    </Box>
                ))}
            </List>
        </>
    )
}