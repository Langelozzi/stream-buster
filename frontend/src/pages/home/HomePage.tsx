import { Search } from "../../components/search/Search";
import { useUser } from "../../hooks/useUser";

const BASE_URL = import.meta.env.VITE_API_BASE_URL

const HomePage = () => {
    const user = useUser();
    console.log('user', user);

    const dummyId = 920;

    return (
        <>
            <Search />
            <iframe src={`${BASE_URL}/cdn/movie/${dummyId}`} allowFullScreen></iframe>
        </>
    )
}

export default HomePage;