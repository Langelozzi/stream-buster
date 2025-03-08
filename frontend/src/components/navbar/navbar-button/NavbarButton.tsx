import { Box, Button, SvgIconTypeMap, useMediaQuery, useTheme } from "@mui/material";
import { OverridableComponent } from "@mui/material/OverridableComponent";

interface NavbarButtonProps {
    Icon: OverridableComponent<SvgIconTypeMap<{}, "svg">>,
    label: string,
    onClick: React.MouseEventHandler<HTMLButtonElement>
}

export const NavbarButton: React.FC<NavbarButtonProps> = ({ Icon, label, onClick }) => {
    const theme = useTheme();
    const isMobile = useMediaQuery(theme.breakpoints.down('sm'));

    const styles = {
        button: {
            position: 'relative',
            overflow: 'hidden',
            padding: 2,
            fontSize: isMobile ? '14px' : '16px',
            textTransform: 'uppercase',
            letterSpacing: '1px',
            minWidth: isMobile ? 'auto' : 'initial',
            '&:hover .underline': {
                width: '54%'
            }
        },
        label: {
            marginLeft: 1,
            fontSize: '1em',
            fontWeight: 'bold',
            display: isMobile ? 'none' : 'block',
        },
        underline: {
            position: 'absolute',
            marginRight: 2,
            height: '2px',
            width: 0,
            backgroundColor: 'white',
            transition: 'width 0.3s ease-in-out',
        }
    }

    return (
        <Button sx={styles.button} color="inherit" onClick={onClick}>
            <Icon />
            <Box sx={styles.label}>
                {label}
                <Box className="underline" sx={styles.underline} />
            </Box>
        </Button>
    )
}