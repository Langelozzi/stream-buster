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
    emptyMessage?: string;
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

export const MediaCarousel: React.FC<MediaCarouselProps> = ({ title, emptyMessage, items, renderItem }) => {
    const theme = useTheme();
    const carouselRef = useRef<HTMLDivElement>(null);
    const [showLeftArrow, setShowLeftArrow] = useState(false);
    const [showRightArrow, setShowRightArrow] = useState(false);
    const [contentLoaded, setContentLoaded] = useState(false);

    const checkOverflow = () => {
        if (!carouselRef.current) return;

        const { scrollWidth, clientWidth } = carouselRef.current;
        // Initial right arrow visibility check - show if content overflows
        setShowRightArrow(scrollWidth > clientWidth);
    };

    const handleScroll = () => {
        if (!carouselRef.current) return;

        const { scrollLeft, scrollWidth, clientWidth } = carouselRef.current;
        setShowLeftArrow(scrollLeft > 0);
        setShowRightArrow(scrollLeft < scrollWidth - clientWidth - 1); // Small buffer for rounding errors
    };

    // Check for overflow when items change
    useEffect(() => {
        setContentLoaded(false);
        // First set a small timeout to check after initial render
        const initialTimer = setTimeout(() => {
            checkOverflow();
        }, 50);

        // Add a longer timeout to account for image loading
        const imageLoadTimer = setTimeout(() => {
            checkOverflow();
            setContentLoaded(true);
        }, 500);

        return () => {
            clearTimeout(initialTimer);
            clearTimeout(imageLoadTimer);
        };
    }, [items]);

    // Set up image load detection
    useEffect(() => {
        if (!carouselRef.current || items.length === 0) return;

        // Find all images within the carousel
        const images = carouselRef.current.querySelectorAll('img');
        if (images.length === 0) {
            // No images, just check for overflow immediately
            checkOverflow();
            setContentLoaded(true);
            return;
        }

        let loadedCount = 0;
        const totalImages = images.length;

        const handleImageLoad = () => {
            loadedCount++;
            if (loadedCount === totalImages) {
                // All images loaded, now check for overflow
                checkOverflow();
                setContentLoaded(true);
            }
        };

        // Add load event listeners to all images
        images.forEach(img => {
            if (img.complete) {
                handleImageLoad();
            } else {
                img.addEventListener('load', handleImageLoad);
            }
        });

        // Cleanup
        return () => {
            images.forEach(img => {
                img.removeEventListener('load', handleImageLoad);
            });
        };
    }, [items]);

    // Use ResizeObserver to detect content changes
    useEffect(() => {
        if (!carouselRef.current) return;

        const resizeObserver = new ResizeObserver(() => {
            checkOverflow();
            handleScroll();
        });

        resizeObserver.observe(carouselRef.current);

        return () => {
            resizeObserver.disconnect();
        };
    }, [contentLoaded]); // Re-run when content loaded changes

    // Update on window resize
    useEffect(() => {
        const handleResize = () => {
            checkOverflow();
            handleScroll();
        };

        window.addEventListener('resize', handleResize);
        return () => window.removeEventListener('resize', handleResize);
    }, []);

    // Update on scroll
    useEffect(() => {
        const carousel = carouselRef.current;
        if (!carousel) return;

        carousel.addEventListener('scroll', handleScroll);
        // Initial check
        handleScroll();

        return () => {
            carousel.removeEventListener('scroll', handleScroll);
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
                    {items.length === 0 ? (
                        <Box sx={{ display: 'flex', justifyContent: 'center', width: '100%' }}>
                            <Typography variant="h6">{emptyMessage}</Typography>
                        </Box>
                    ) : (
                        <Stack direction="row" spacing={1}>
                            {items.map((item, index) => (
                                <Box key={index} sx={styles.itemContainer}>
                                    {renderItem(item)}
                                </Box>
                            ))}
                        </Stack>
                    )}
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
