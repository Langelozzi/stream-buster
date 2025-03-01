import React from "react";
import { Episode } from "../../../models/episode";
import { Typography, Divider, List, Box, Select, MenuItem, SelectChangeEvent, useMediaQuery, useTheme } from "@mui/material";
import { EpisodeListItem } from "./EpisodeListItem";
import { TV } from "../../../models/tv";
import { Season } from "../../../models/season";
import { useTranslation } from "react-i18next";

// Define styles as a JSON object
const styles = {
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
        marginTop: 2,
    },
    headerContainer: {
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
    },
    select: {
        color: 'white',
        backgroundColor: '#333', // Dark grey background for selector
        borderRadius: 2, // Rounded corners for a better visual effect
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
};

interface MediaDetailsModalEpisodesProps {
    tv: TV;
    episodes: Episode[];
    currentSeason: Season;
    setCurrentSeason: (season: Season) => void;
}

export const MediaDetailsModalEpisodes: React.FC<MediaDetailsModalEpisodesProps> = (props) => {
    const { tv, episodes, currentSeason, setCurrentSeason } = props;

    const { t } = useTranslation();
    const theme = useTheme();
    const isMobile = useMediaQuery(theme.breakpoints.down('sm'));

    const handleSeasonChange = (event: SelectChangeEvent<number>) => {
        const selectedSeasonNumber = event.target.value as number;
        const selectedSeason = tv.Seasons.find(season => season.SeasonNumber === selectedSeasonNumber);
        if (selectedSeason) setCurrentSeason(selectedSeason);
    };

    return (
        <>
            <Box sx={{
                ...styles.headerContainer,
                flexDirection: isMobile ? 'column' : 'row',
                alignItems: isMobile ? 'flex-start' : 'center',
                gap: isMobile ? '16px' : '0px'
            }}
            >
                <Typography variant={isMobile ? "h6" : "h5"}>{t('dictionary.episodes')}</Typography>

                {/* Season Selector */}
                <Select
                    value={currentSeason.SeasonNumber}
                    onChange={handleSeasonChange}
                    variant="standard"
                    sx={{
                        ...styles.select,
                        minWidth: 120,
                        width: isMobile ? '100%' : 'auto'
                    }}
                >
                    {tv.Seasons?.map((season, index) => (
                        season.SeasonNumber > 0 && (
                            <MenuItem key={index} value={season.SeasonNumber} sx={styles.menuItem}>
                                {season.Name}
                            </MenuItem>
                        )
                    ))}
                </Select>
            </Box>

            <List sx={styles.episodeList}>
                {episodes?.map((episode) => (
                    <Box key={episode.EpisodeTMDBID}>
                        <Divider sx={styles.episodeDivider} />
                        <EpisodeListItem tv={tv} episode={episode} />
                    </Box>
                ))}
            </List>
        </>
    );
};
