import React, {useEffect, useState} from "react";
import {useMovieApi} from "../hooks/useMovieApi.ts";
import VideoDropzone from "../components/VideoUpload.tsx";
import ImageDropzone from "../components/ImageUpload.tsx";
import VideoPlayer from "../components/VideoPlayer.tsx";

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
                    <div className="card-body">
                        <div className="card-title">
                            <h5>Upload Video</h5>
                        </div>
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
                    <div className="card-body">
                        <div className="d-flex card-title">
                            <h5>Your Video</h5>
                            <button className="btn btn-outline-dark btn-sm ms-auto">
                                <i className="bi bi-view-list"></i>
                            </button>
                        </div>
                        <div className="row row-cols-1 row-cols-md-2 g-4">
                            {data.map((movie) => (
                            <div className="col" key={`key for ${movie.id}`}>
                                <div className="card" onClick={() => handleFetchMovie(movie.id)}>
                                    <VideoPlayer videoUrl={movie.video_url} posterUrl={movie.poster_url}/>
                                    <div className="card-body">
                                        <h5 className="card-title">{movie.title}</h5>
                                    </div>
                                </div>
                            </div>
                            ))}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default VoD;
