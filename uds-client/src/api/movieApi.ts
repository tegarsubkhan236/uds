import {AxiosResponse} from 'axios';
import axiosInstance from "../axios.ts";

export type Movie = {
    id: number;
    title: string;
    video_url: File;
    poster_url: File;
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
        const formData = new FormData();
        formData.append('title', movie.title);
        formData.append('poster_file', movie.poster_url);
        formData.append('video_file', movie.video_url);

        const response = await axiosInstance.post('/movie/create', formData, {
            headers: {
                "Content-Type": "multipart/form-data",
            },
        });
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
