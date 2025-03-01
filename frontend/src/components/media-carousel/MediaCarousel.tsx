import { useState, useRef } from 'react';
import {
    Box,
    Typography,
    IconButton,
    useTheme,
    Stack
} from '@mui/material';
import ArrowBackIosNewIcon from '@mui/icons-material/ArrowBackIosNew';
import ArrowForwardIosIcon from '@mui/icons-material/ArrowForwardIos';

interface MediaCarouselProps {
    title: string;
    items: any[];
    renderItem: (item: any) => JSX.Element;
}

const styles = {
    container: {
        position: 'relative'
    },
    title: {
        mb: 2,
        ml: 2,
        fontWeight: 'bold'
    },
    wrapper: {
        position: 'relative',
        display: 'flex',
        alignItems: 'center'
    },
    arrowButton: {
        position: 'absolute',
        zIndex: 2,
        height: '100%',
        borderRadius: 0,
        p: 1,
        backgroundColor: 'rgba(0, 0, 0, 0.5)',
        '&:hover': {
            backgroundColor: 'rgba(0, 0, 0, 0.7)',
        }
    },
    leftArrow: {
        left: 0
    },
    rightArrow: {
        right: 0
    },
    arrowIcon: {
        color: 'white'
    },
    carouselContainer: {
        display: 'flex',
        width: '100%',
        overflowX: 'hidden',
        scrollBehavior: 'smooth',
        py: 1,
        px: 2,
        '&::-webkit-scrollbar': {
            display: 'none'
        }
    },
    itemContainer: {
        flex: '0 0 auto',
    }
};

export const MediaCarousel: React.FC<MediaCarouselProps> = ({ title, items, renderItem }) => {
    const theme = useTheme();
    const [scrollPosition, setScrollPosition] = useState(0);
    const carouselRef = useRef<HTMLDivElement>(null);
    const [showLeftArrow, setShowLeftArrow] = useState(false);
    const [showRightArrow, setShowRightArrow] = useState(true);

    const handleScroll = (direction: 'left' | 'right') => {
        if (!carouselRef.current) return;

        const scrollAmount = carouselRef.current.clientWidth * 0.8;
        const maxScroll = carouselRef.current.scrollWidth - carouselRef.current.clientWidth;

        let newPosition;
        if (direction === 'left') {
            newPosition = Math.max(0, scrollPosition - scrollAmount);
        } else {
            newPosition = Math.min(maxScroll, scrollPosition + scrollAmount);
        }

        setScrollPosition(newPosition);
        setShowLeftArrow(newPosition > 0);
        setShowRightArrow(newPosition < maxScroll);

        carouselRef.current.scrollTo({
            left: newPosition,
            behavior: 'smooth'
        });
    };

    return (
        <Box sx={styles.container}>
            <Typography
                variant="h5"
                sx={{
                    ...styles.title,
                    color: theme.palette.mode === 'dark' ? 'white' : 'inherit'
                }}
            >
                {title}
            </Typography>

            <Box sx={styles.wrapper}>
                {showLeftArrow && (
                    <IconButton
                        sx={{ ...styles.arrowButton, ...styles.leftArrow }}
                        onClick={() => handleScroll('left')}
                    >
                        <ArrowBackIosNewIcon sx={styles.arrowIcon} />
                    </IconButton>
                )}

                <Box
                    ref={carouselRef}
                    sx={styles.carouselContainer}
                >
                    <Stack direction="row" spacing={1}>
                        {items.map((item, index) => (
                            <Box
                                key={index}
                                sx={styles.itemContainer}
                            >
                                {renderItem(item)}
                            </Box>
                        ))}
                    </Stack>
                </Box>

                {showRightArrow && (
                    <IconButton
                        sx={{ ...styles.arrowButton, ...styles.rightArrow }}
                        onClick={() => handleScroll('right')}
                    >
                        <ArrowForwardIosIcon sx={styles.arrowIcon} />
                    </IconButton>
                )}
            </Box>
        </Box>
    );
};