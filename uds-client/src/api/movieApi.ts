import {AxiosResponse} from 'axios';
import axiosInstance from "../axios.ts";

export type Movie = {
    id: number;
    title: string;
    director: string;
    releaseDate: string;
}

type Meta = {
    currentPage: number,
    lastPage: number,
    totalPage: number,
}

type MovieApiResponse = {
    data: Movie[];
    meta: Meta
};

const MovieApi = () => {
    const fetchMovies = async (page: number, limit: number): Promise<Movie[]> => {
        const queryParams = new URLSearchParams({
            page: page.toString(),
            limit: limit.toString(),
        });

        const response: AxiosResponse<MovieApiResponse> = await axiosInstance.get(`/movie?${queryParams}`);
        return response.data.data;
    };

    const fetchMovieById = async (id: number): Promise<Movie> => {
        const response: AxiosResponse<Movie> = await axiosInstance.get(`/movies/${id}`);
        return response.data;
    };

    const createMovie = async (movie: Omit<Movie, 'id'>): Promise<Movie> => {
        const response: AxiosResponse<Movie> = await axiosInstance.post('/movies', movie);
        return response.data;
    };

    const updateMovie = async (id: number, movie: Partial<Movie>): Promise<Movie> => {
        const response: AxiosResponse<Movie> = await axiosInstance.put(`/movies/${id}`, movie);
        return response.data;
    };

    const deleteMovie = async (id: number): Promise<void> => {
        await axiosInstance.delete(`/movies/${id}`);
    };

    return {
        fetchMovies,
        fetchMovieById,
        createMovie,
        updateMovie,
        deleteMovie,
    };
};

export default MovieApi
