import {useEffect, useState} from "react";
import useMovieApi from "./hooks/useMovieApi.ts";
import {Movie} from "./api/movieApi.ts";

function App() {
    const [page] = useState(1);
    const [limit] = useState(10);

    const {movies, loading, error, fetchMoviesList, fetchMovie} = useMovieApi()

    useEffect(() => {
        fetchMoviesList(page, limit)
    }, [limit, page]);

    const handleFetchMovie = async (id: number) => {
        const movie = await fetchMovie(id)
        console.log(movie)
    }

    if (loading) return <div>Loading...</div>
    if (error) return <div>{error}</div>;

    return (
        <>
            <h1>Movie</h1>
            <hr/>
            <form action="">
                <input type="text" name="title" placeholder="Title"/> <br/>
                <label>Poster : </label><input type="file" name="poster_url"/><br/>
                <label>Video : </label><input type="file" name="video_url"/><br/>
                <button type="submit">Save</button>
            </form>
            <hr/>
            <ol>
                {movies.map((movie: Movie) => (
                    <li key={movie.id}>
                        {movie.title} |
                        <button onClick={() => handleFetchMovie(movie.id)}>
                            View Detail
                        </button> |
                        <button onClick={() => handleFetchMovie(movie.id)}>
                            Delete
                        </button> |
                        <button onClick={() => handleFetchMovie(movie.id)}>
                            Update
                        </button>
                    </li>
                ))}
            </ol>
            <hr/>
        </>
    )
}

export default App
