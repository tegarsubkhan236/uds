import React, {useEffect, useState} from "react";
import {useMovieApi} from "../hooks/useMovieApi.ts";
import VideoDropzone from "../components/VideoUpload.tsx";
import ImageDropzone from "../components/ImageUpload.tsx";

const VoD = () => {
    const [page] = useState(1);
    const [limit] = useState(10);
    const [videoFile, setVideoFile] = useState<File | null>(null);
    const [imageFile, setImageFile] = useState<File | null>(null);
    const [title, setTitle] = useState<string>("");
    const {
        data,
        loading,
        error,
        fetchMoviesHandler,
        fetchMovieByIdHandler,
        createMovieHandler,
        // updateMovieHandler,
        // deleteMovieHandler,
    } = useMovieApi();

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        if (!title || !videoFile || !imageFile) {
            alert("All fields are required");
            return;
        }
        await createMovieHandler({
            title,
            video_url: videoFile,
            poster_url: imageFile,
        });

        setTitle("");
        setVideoFile(null);
        setImageFile(null);
    };

    useEffect(() => {
        fetchMoviesHandler(page, limit)
    }, [limit, page]);

    const handleFetchMovie = async (id: number) => {
        return await fetchMovieByIdHandler(id);
    }

    if (loading) return <div>Loading...</div>

    if (error) return <div>{error}</div>;

    return (
        <div className="row mt-4">
            <div className="col-md-4 col-xs-12 mb-4">
                <div className="card">
                    <div className="card-header">
                        <h4>Upload Video</h4>
                    </div>
                    <div className="card-body">
                        <form onSubmit={handleSubmit}>
                            <div className="mb-3">
                                <VideoDropzone name="video_file" onFileChange={setVideoFile}/>
                            </div>
                            <div className="mb-3">
                                <ImageDropzone name="image_file" onFileChange={setImageFile}/>
                            </div>
                            <div className="mb-3">
                                <input name="title"
                                       type="text"
                                       className="form-control"
                                       placeholder="Video Title"
                                       value={title}
                                       onChange={(e) => setTitle(e.target.value)}/>
                            </div>
                            <button type="submit" className="btn btn-primary w-100">Upload</button>
                        </form>
                    </div>
                </div>
            </div>
            <div className="col-md-8 col-xs-12">
                <div className="card">
                    <div className="card-header">
                        <h4>Your Video</h4>
                    </div>
                    <div className="card-body">
                        <table>
                            <thead>
                            <tr>
                                <th>ID</th>
                                <th>Title</th>
                                <th>Action</th>
                            </tr>
                            </thead>
                            <tbody>
                            {data.map((movie) => (
                                <tr key={movie.id}>
                                    <td>{movie.id}</td>
                                    <td>{movie.title}</td>
                                    <td>
                                        <button onClick={() => handleFetchMovie(movie.id)}>
                                            View Detail
                                        </button>
                                        |
                                        <button onClick={() => handleFetchMovie(movie.id)}>
                                            Delete
                                        </button>|
                                        <button onClick={() => handleFetchMovie(movie.id)}>
                                            Update
                                        </button>
                                    </td>
                                </tr>
                            ))}
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default VoD;
