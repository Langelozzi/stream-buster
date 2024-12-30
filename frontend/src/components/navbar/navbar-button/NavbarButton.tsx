import { Box, Button, SvgIconTypeMap, Typography } from "@mui/material";
import { OverridableComponent } from "@mui/material/OverridableComponent";

interface NavbarButtonProps {
    Icon: OverridableComponent<SvgIconTypeMap<{}, "svg">>,
    label: string,
    onClick: React.MouseEventHandler<HTMLButtonElement>
}

export const NavbarButton: React.FC<NavbarButtonProps> = ({ Icon, label, onClick }) => {
    const styles = {
        button: {
            position: 'relative',
            overflow: 'hidden',
            padding: 2,
            fontSize: '16px',
            textTransform: 'uppercase',
            letterSpacing: '1px',
            '&:hover .underline': {
                width: '54%'
            }
        },
        label: {
            marginLeft: 1,
            fontSize: '1em',
            fontWeight: 'bold'
        },
        underline: {
            position: 'absolute',
            // bottom: 0,
            // left: 0,
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