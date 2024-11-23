import { useEffect, useState } from "react";
import { getWatchList } from "../../api/services/currentlyWatching.service";
import MediaList from "../../components/media-list/medialist";
import { useUser } from "../../hooks/useUser";
import { Box, Typography } from "@mui/material";

export const HomePage = () => {
    const user = useUser()
    const [currentlyWatchings, setCurrentlyWatchings] = useState<any[]>([])
    useEffect(() => {
        // const test = async () => {
        //     const res = await searchMulti("How to train your dragon")
        //     
        //     const himym = res[0]
        //     
        //     const createdMedia = await createMedia(himym.Media!)
        // }
        //
        // if (first == 0) {
        //     setFirst(1)
        //     test()
        // }

        const getMediaList = async () => {
            const currentlyWatchingList = await getWatchList()
            setCurrentlyWatchings(currentlyWatchingList)
        }
        getMediaList()

    }, [])

    return (
        <>
            {user ? (
                <>
                    <Box padding={2}>

                        <Typography>Continue Watching</Typography>
                        <MediaList currentlyWatchings={currentlyWatchings}></MediaList>
                    </Box>
                </>
            ) : (
                <>
                    <div>Welcome to stream buster</div>
                    <div>login or sign up</div>
                </>
            )}
        </>
    )
}
