import MediaPlayer from "../../components/media-player/MediaPlayer";
import { Search } from "../../components/search/Search";
import { useUser } from "../../hooks/useUser";


const HomePage = () => {
    const user = useUser();
    console.log('user', user);

    const dummyId = 920;

    return (
        <>
            <Search />
            <MediaPlayer tmdbId={dummyId} />
        </>
    )
}

export default HomePage;