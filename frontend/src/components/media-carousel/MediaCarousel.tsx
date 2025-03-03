import { useState, useRef, useEffect } from 'react';
import {
    Box,
    Typography,
    IconButton,
    useTheme,
    Stack,
    Divider
} from '@mui/material';
import ArrowBackIosNewIcon from '@mui/icons-material/ArrowBackIosNew';
import ArrowForwardIosIcon from '@mui/icons-material/ArrowForwardIos';

interface MediaCarouselProps {
    title: string;
    items: any[];
    renderItem: (item: any) => JSX.Element;
}

const styles = {
    container: { position: 'relative' },
    title: { mb: 2, mx: 2, fontWeight: 'bold' },
    wrapper: { position: 'relative', display: 'flex', alignItems: 'center' },
    arrowButton: {
        position: 'absolute',
        zIndex: 2,
        height: '100%',
        borderRadius: 0,
        p: 1,
        backgroundColor: 'rgba(0, 0, 0, 0.5)',
        '&:hover': { backgroundColor: 'rgba(0, 0, 0, 0.7)' },
    },
    leftArrow: { left: 0 },
    rightArrow: { right: 0 },
    arrowIcon: { color: 'white' },
    carouselContainer: {
        display: 'flex',
        width: '100%',
        overflowX: 'hidden',
        scrollBehavior: 'smooth',
        py: 1,
        px: 2,
        '&::-webkit-scrollbar': { display: 'none' },
    },
    itemContainer: { flex: '0 0 auto' },
    divider: { borderColor: 'gray', marginTop: 1 },
};

export const MediaCarousel: React.FC<MediaCarouselProps> = ({ title, items, renderItem }) => {
    const theme = useTheme();
    const [scrollPosition] = useState(0);
    const carouselRef = useRef<HTMLDivElement>(null);
    const [showLeftArrow, setShowLeftArrow] = useState(false);
    const [showRightArrow, setShowRightArrow] = useState(true);

    const updateArrows = () => {
        if (carouselRef.current) {
            const maxScroll = carouselRef.current.scrollWidth - carouselRef.current.clientWidth;
            setShowLeftArrow(scrollPosition > 0);
            setShowRightArrow(maxScroll > 0 && scrollPosition < maxScroll);
        }
    };

    // Update arrows on mount
    useEffect(() => {
        updateArrows();
    }, [items]);

    // Update arrows on resize
    useEffect(() => {
        const handleResize = () => updateArrows();
        window.addEventListener('resize', handleResize);
        return () => window.removeEventListener('resize', handleResize);
    }, [scrollPosition]);

    // Update arrows on scroll event
    useEffect(() => {
        const handleScroll = () => {
            if (!carouselRef.current) return;
            const { scrollLeft, scrollWidth, clientWidth } = carouselRef.current;
            setShowLeftArrow(scrollLeft > 0);
            setShowRightArrow(scrollLeft < scrollWidth - clientWidth);
        };

        const carousel = carouselRef.current;
        carousel?.addEventListener('scroll', handleScroll);
        handleScroll();

        return () => {
            carousel?.removeEventListener('scroll', handleScroll);
        };
    }, []);

    const handleScrollClick = (direction: 'left' | 'right') => {
        if (!carouselRef.current) return;
        const scrollAmount = carouselRef.current.clientWidth * 0.8;
        carouselRef.current.scrollBy({
            left: direction === 'left' ? -scrollAmount : scrollAmount,
            behavior: 'smooth'
        });
    };

    return (
        <Box sx={styles.container}>
            <Typography
                variant="h5"
                sx={{ ...styles.title, color: theme.palette.mode === 'dark' ? 'white' : 'inherit' }}
            >
                {title}
                <Divider sx={styles.divider} />
            </Typography>

            <Box sx={styles.wrapper}>
                {showLeftArrow && (
                    <IconButton
                        sx={{ ...styles.arrowButton, ...styles.leftArrow }}
                        onClick={() => handleScrollClick('left')}
                    >
                        <ArrowBackIosNewIcon sx={styles.arrowIcon} />
                    </IconButton>
                )}

                <Box ref={carouselRef} sx={styles.carouselContainer}>
                    <Stack direction="row" spacing={1}>
                        {items.map((item, index) => (
                            <Box key={index} sx={styles.itemContainer}>
                                {renderItem(item)}
                            </Box>
                        ))}
                    </Stack>
                </Box>

                {showRightArrow && (
                    <IconButton
                        sx={{ ...styles.arrowButton, ...styles.rightArrow }}
                        onClick={() => handleScrollClick('right')}
                    >
                        <ArrowForwardIosIcon sx={styles.arrowIcon} />
                    </IconButton>
                )}
            </Box>
        </Box>
    );
};
