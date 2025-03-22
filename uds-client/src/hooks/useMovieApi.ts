import {useCallback, useState} from "react";
import movieApi, {Movie} from "../api/movieApi.ts";
import {handleError} from "../utils.ts";

const useMovieApi = () => {
    const [movies, setMovies] = useState<Movie[]>([]);
    const [loading, setLoading] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);

    const { fetchMovies, fetchMovieById, createMovie, updateMovie, deleteMovie } = movieApi();

    const fetchMoviesList = useCallback(async (page: number, limit: number) => {
        setLoading(true);
        setError(null);
        try {
            const fetchedMovies = await fetchMovies(page, limit);
            setMovies(fetchedMovies);
        } catch (err) {
            setError(handleError(err));
        } finally {
            setLoading(false);
        }
    }, [fetchMovies]);

    const fetchMovie = useCallback(async (id: number) => {
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

    const createNewMovie = useCallback(async (movieData: Omit<Movie, 'id'>) => {
        setLoading(true);
        setError(null);
        try {
            const newMovie = await createMovie(movieData);
            setMovies((prevMovies) => [...prevMovies, newMovie]);
        } catch (err) {
            setError(handleError(err));
        } finally {
            setLoading(false);
        }
    }, [createMovie]);

    const updateMovieDetails = useCallback(async (id: number, movieData: Partial<Movie>) => {
        setLoading(true);
        setError(null);
        try {
            const updatedMovie = await updateMovie(id, movieData);
            setMovies((prevMovies) =>
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

    const deleteMovieById = useCallback(async (id: number) => {
        setLoading(true);
        setError(null);
        try {
            await deleteMovie(id);
            setMovies((prevMovies) => prevMovies.filter((movie) => movie.id !== id)); // Remove deleted movie from state
        } catch (err) {
            setError(handleError(err));
        } finally {
            setLoading(false);
        }
    }, [deleteMovie]);

    return {
        movies,
        loading,
        error,
        fetchMoviesList,
        fetchMovie,
        createNewMovie,
        updateMovieDetails,
        deleteMovieById,
    };
};

export default useMovieApi;
