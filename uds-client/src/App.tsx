import {useEffect, useState} from "react";
import useMovieApi from "./hooks/useMovieApi.ts";
import {Movie} from "./api/movieApi.ts";

function App() {
    const [page] = useState(1);
    const [limit] = useState(10);

    const {movies, loading, error, fetchMoviesList, createNewMovie} = useMovieApi()

    const [newMovie, setNewMovie] = useState({
        title: '',
        poster_url: null as File | null,
        video_url: null as File | null,
    });

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const {name, value, type, files} = e.target;
        if (type === "file") {
            setNewMovie((prevState) => ({
                ...prevState,
                [name]: files ? files[0] : null,
            }));
        } else {
            setNewMovie((prevState) => ({
                ...prevState,
                [name]: value,
            }));
        }
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        const {title, poster_url, video_url} = newMovie;

        if (!title || !poster_url || !video_url) {
            alert("All fields are required");
            return;
        }

        await createNewMovie({title, poster_url, video_url});
        setNewMovie({title: '', poster_url: null, video_url: null});
    };

    useEffect(() => {
        fetchMoviesList(page, limit)
    }, [limit, page]);

    const handleFetchMovie = async (id: number) => {
        console.log(id)
    }

    if (loading) return <div>Loading...</div>
    if (error) return <div>{error}</div>;

    return (
        <>
            <h1>Movie</h1>
            <hr/>
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    name="title"
                    value={newMovie.title}
                    onChange={handleInputChange}
                    placeholder="Title"
                />{" "}
                <br/>
                <label>Poster: </label>
                <input
                    type="file"
                    name="poster_url"
                    onChange={handleInputChange}
                />
                <br/>
                <label>Video: </label>
                <input
                    type="file"
                    name="video_url"
                    onChange={handleInputChange}
                />
                <br/>
                <button type="submit">Save</button>
            </form>
            <hr/>
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
