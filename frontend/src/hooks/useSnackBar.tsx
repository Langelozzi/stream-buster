import { Snackbar } from '@mui/material';
import { useState, useCallback } from 'react';

interface UseSnackbarProps {
    autoClose?: boolean;
    autoHideDuration?: number;
    children?: React.ReactNode;
    onClose?: (event: Event, reason: string) => void
}

export const useSnackbar = ({
    autoClose = true,
    autoHideDuration = 10000,
    onClose: userOnClose,
    children,
}: UseSnackbarProps = {}) => {
    const [open, setOpen] = useState(false);
    const [message, setMessage] = useState('');

    const handleClose = useCallback(
        (event?: React.SyntheticEvent | Event, reason?: string) => {
            if (reason === 'clickaway' && !autoClose) return;

            setOpen(false);
            if (userOnClose) {
                userOnClose(event, reason);
            }
        },
        [userOnClose, autoClose]
    );

    const showSnackbar = useCallback((newMessage: string) => {
        setMessage(newMessage);
        setOpen(true);
    }, []);

    const SnackbarComponent = (
        <Snackbar
            open={open}
            autoHideDuration={autoClose ? autoHideDuration : null}
            message={children || message}
            onClose={handleClose}
        />
    );

    return {
        showSnackbar,
        SnackbarComponent,
    };
};
