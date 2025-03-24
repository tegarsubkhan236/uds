import {useCallback, useState} from "react";
import movieApi, {Movie} from "../api/movieApi.ts";
import {handleError} from "../utils.ts";

const useMovieApi = () => {
    const [data, setData] = useState<Movie[]>([]);
    const [loading, setLoading] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);

    const { fetchMovies, fetchMovieById, createMovie, updateMovie, deleteMovie } = movieApi();

    const fetchMoviesHandler = useCallback(async (page: number, limit: number) => {
        setLoading(true);
        setError(null);
        try {
            const fetchedMovies = await fetchMovies(page, limit);
            setData(fetchedMovies);
        } catch (err) {
            setError(handleError(err));
        } finally {
            setLoading(false);
        }
    }, [fetchMovies]);

    const fetchMovieByIdHandler = useCallback(async (id: number) => {
        setLoading(true);
        setError(null);
        try {
            return await fetchMovieById(id);
        } catch (err) {
            setError(handleError(err));
        } finally {
            setLoading(false);
        }
    }, [fetchMovieById]);

    const createMovieHandler = useCallback(async (movieData: Omit<Movie, 'id'>) => {
        setLoading(true);
        setError(null);
        try {
            const newMovieID = await createMovie(movieData);
            const newMovie = await fetchMovieById(newMovieID);
            setData((prevMovies) => [...prevMovies, newMovie]);
        } catch (err) {
            setError(handleError(err));
        } finally {
            setLoading(false);
        }
    }, [createMovie, fetchMovieById]);

    const updateMovieHandler = useCallback(async (id: number, movieData: Partial<Movie>) => {
        setLoading(true);
        setError(null);
        try {
            const updatedMovie = await updateMovie(id, movieData);
            setData((prevMovies) =>
                prevMovies.map((movie) =>
                    movie.id === updatedMovie.id ? updatedMovie : movie
                )
            );
        } catch (err) {
            setError(handleError(err));
        } finally {
            setLoading(false);
        }
    }, [updateMovie]);

    const deleteMovieHandler = useCallback(async (id: number) => {
        setLoading(true);
        setError(null);
        try {
            await deleteMovie(id);
            setData((prevMovies) => prevMovies.filter((movie) => movie.id !== id)); // Remove deleted movie from state
        } catch (err) {
            setError(handleError(err));
        } finally {
            setLoading(false);
        }
    }, [deleteMovie]);

    return {
        data,
        loading,
        error,
        fetchMoviesHandler,
        fetchMovieByIdHandler,
        createMovieHandler,
        updateMovieHandler,
        deleteMovieHandler,
    };
};

export default useMovieApi;
