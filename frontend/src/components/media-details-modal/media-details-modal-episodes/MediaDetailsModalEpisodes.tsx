import React from "react";
import { Episode } from "../../../models/episode";
import { Typography, Divider, List, Box, Select, MenuItem, SelectChangeEvent } from "@mui/material";
import { makeStyles } from '@mui/styles';
import { EpisodeListItem } from "./EpisodeListItem";
import { TV } from "../../../models/tv";
import { Season } from "../../../models/season";

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
        marginTop: 24,
    },
    headerContainer: {
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        marginBottom: 16,
    },
    select: {
        color: 'white',
        backgroundColor: '#333', // Dark grey background for selector
        borderRadius: 4, // Rounded corners for a better visual effect
        padding: '8px 16px',
        '& .MuiSelect-icon': {
            color: 'white',
        },
    },
    menuItem: {
        color: 'black',
        backgroundColor: 'white',
        '&:hover': {
            backgroundColor: '#f0f0f0',
        },
    },
});

interface MediaDetailsModalEpisodesProps {
    tv: TV;
    episodes: Episode[];
    currentSeason: Season;
    setCurrentSeason: (season: Season) => void;
}

export const MediaDetailsModalEpisodes: React.FC<MediaDetailsModalEpisodesProps> = (props) => {
    const { tv, episodes, currentSeason, setCurrentSeason } = props;

    const classes = useStyles();

    const handleSeasonChange = (event: SelectChangeEvent<number>) => {
        const selectedSeasonNumber = event.target.value as number;
        const selectedSeason = tv.Seasons.find(season => season.SeasonNumber === selectedSeasonNumber);
        if (selectedSeason) setCurrentSeason(selectedSeason);
    };

    return (
        <>
            <Box className={classes.headerContainer}>
                <Typography variant="h5">Episodes</Typography>

                {/* Right-aligned Season Selector */}
                <Select
                    value={currentSeason.SeasonNumber}
                    onChange={handleSeasonChange}
                    variant="standard" // No outline variant
                    className={classes.select}
                    sx={{ minWidth: 120 }}
                >
                    {tv.Seasons?.map((season, index) => (
                        // Only display actual seasons, not specials (for now)
                        season.SeasonNumber > 0 && (
                            <MenuItem key={index} value={season.SeasonNumber} className={classes.menuItem}>
                                {season.Name}
                            </MenuItem>
                        )
                    ))}
                </Select>
            </Box>

            <List className={classes.episodeList}>
                {episodes?.map((episode) => (
                    <Box key={episode.EpisodeTMDBID}>
                        <Divider className={classes.episodeDivider} />
                        <EpisodeListItem tv={tv} episode={episode} />
                    </Box>
                ))}
            </List>
        </>
    );
};
