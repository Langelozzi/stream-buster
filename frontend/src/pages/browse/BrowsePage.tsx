import { useEffect, useState } from "react";
import { searchTrendingMovies, searchTrendingTv } from "../../api/services/search.service";
import { MediaCarousel } from "../../components/media-carousel/MediaCarousel";
import { Box } from "@mui/material";
import { MediaCard } from "../../components/media-card/MediaCard";
import { SearchPage } from "../../models/search_page";
import { deleteCurrentlyWatching, getWatchList } from "../../api/services/currentlyWatching.service";

const styles = {
    pageContainer: {
        py: 2
    },
    pageTitle: {
        mb: 4
    },
    movieCard: {
        width: 220,
        height: 330
    },
    cardMedia: {
        height: 280
    },
    cardContent: {
        p: 1,
        height: 50
    }
};

export const BrowsePage = () => {
    const [currentlyWatching, setCurrentlyWatching] = useState<any[]>([]);
    const [moviePage, setMoviePage] = useState<SearchPage>();
    const [tvPage, setTvPage] = useState<SearchPage>();

    const onRemoveCurrentlyWatching = async (mediaId: number) => {
        try {
            const removeMediaItem = (mediaId: number) => {
                setCurrentlyWatching(prevItems => {
                    return prevItems!.filter(tvOrMovie => tvOrMovie!.MediaId != mediaId)
                })
            }
            removeMediaItem(mediaId)

            await deleteCurrentlyWatching(mediaId)
        } catch (error) {
            console.error("Error deleteing currently watching" + error)
            throw error
        }
    }

    const fetchCurrentlyWatching = async () => {
        const currentlyWatchings = await getWatchList();
        setCurrentlyWatching(currentlyWatchings);
    }

    const fetchTrendingMovies = async () => {
        const pageResults = await searchTrendingMovies();
        setMoviePage(pageResults);
    }

    const fetchTrendingTv = async () => {
        const pageResults = await searchTrendingTv();
        setTvPage(pageResults);
    }

    useEffect(() => {
        fetchCurrentlyWatching();
        fetchTrendingMovies();
        fetchTrendingTv();
    }, []);

    return (
        <Box sx={styles.pageContainer}>
            {!!currentlyWatching.length && (
                <>
                    <MediaCarousel
                        title="Continue Watching"
                        items={currentlyWatching}
                        renderItem={(item) => (
                            <MediaCard
                                media={item}
                                currentlyWatching={item}
                                onDelete={onRemoveCurrentlyWatching}
                                search={false}
                            />
                        )}
                    />
                    <br />
                </>
            )}
            {moviePage?.Results && (
                <>
                    <MediaCarousel
                        title="Trending Movies"
                        items={moviePage.Results}
                        renderItem={(movie) => (<MediaCard media={movie} />)}
                    />
                    <br />
                </>
            )}
            {tvPage?.Results && (
                <MediaCarousel
                    title="Trending TV"
                    items={tvPage.Results}
                    renderItem={(tv) => (<MediaCard media={tv} />)}
                />
            )}
        </Box>
    )
}
