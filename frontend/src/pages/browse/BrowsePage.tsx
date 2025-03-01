import { useEffect, useState } from "react"
import { searchTrendingMovies } from "../../api/services/search.service";
import { MediaCarousel } from "../../components/media-carousel/MediaCarousel";
import { Container, Box, Typography } from "@mui/material";
import { MediaCard } from "../../components/media-card/MediaCard";

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

    const [moviePage, setMoviePage] = useState();

    const fetchTrendingMovies = async () => {
        const pageResults = await searchTrendingMovies();
        setMoviePage(pageResults);
    }

    useEffect(() => {
        fetchTrendingMovies();
    }, []);

    return (
        <Box sx={styles.pageContainer}>
            {moviePage?.Results && (
                <MediaCarousel
                    title="Trending Movies"
                    items={moviePage.Results}
                    renderItem={(movie) => (<MediaCard media={movie} />)}
                />
            )}
        </Box>
    )
}
