
import { useEffect } from "react";

type ShortcutMap = {
    [key: string]: () => void;
};

const useKeyboardShortcuts = (shortcuts: ShortcutMap) => {
    useEffect(() => {
        const handleKeyDown = (event: KeyboardEvent) => {
            if (shortcuts[event.key]) {
                event.preventDefault(); // Prevent default action if needed
                shortcuts[event.key]();
            }
        };

        window.addEventListener("keydown", handleKeyDown);
        return () => {
            window.removeEventListener("keydown", handleKeyDown);
        };
    }, [shortcuts]);
};

export default useKeyboardShortcuts;
