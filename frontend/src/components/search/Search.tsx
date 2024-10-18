import { useEffect, useState } from "react";
import { searchMulti } from "../../api/search";

const Search = () => {
    const [query, setQuery] = useState('how I met your mother');
    const [results, setResults] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchResults = async () => {
            try {
                const data = await searchMulti(query);
                setResults(data);
            } catch (error) {
                console.error('Failed to fetch search results:', error);
            } finally {
                setLoading(false);
            }
        };

        fetchResults();
    }, [query]);

    return (
        <div>
            {loading ? <p>Loading...</p> : <pre>{JSON.stringify(results, null, 2)}</pre>}
        </div>
    );
}

export default Search