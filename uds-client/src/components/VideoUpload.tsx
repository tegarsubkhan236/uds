import React, { useState, DragEvent, ChangeEvent } from 'react';

interface VideoDropzoneProps {
    name: string;
    onFileChange: (file: File | null) => void;
}

const VideoDropzone: React.FC<VideoDropzoneProps> = ({ name, onFileChange }) => {
    const [previewURL, setPreviewURL] = useState<string | null>(null);
    const [isDragging, setIsDragging] = useState(false);
    const [videoFile, setVideoFile] = useState<File | null>(null);

    const handleDrop = (e: DragEvent<HTMLDivElement>) => {
        e.preventDefault();
        setIsDragging(false);
        const file = e.dataTransfer.files[0];
        handleFile(file);
    };

    const handleFileChange = (e: ChangeEvent<HTMLInputElement>) => {
        const file = e.target.files?.[0];
        handleFile(file);
    };

    const handleFile = (file?: File) => {
        if (file && file.type.startsWith('video/')) {
            setVideoFile(file);
            setPreviewURL(URL.createObjectURL(file));
            onFileChange(file);
        } else {
            setVideoFile(null);
            setPreviewURL(null);
            onFileChange(null);
        }
    };

    return (
        <div
            className={`border rounded p-4 text-center ${isDragging ? 'bg-light' : ''}`}
            style={{ borderStyle: 'dashed', cursor: 'pointer' }}
            onDragOver={(e) => {
                e.preventDefault();
                setIsDragging(true);
            }}
            onDragLeave={() => setIsDragging(false)}
            onDrop={handleDrop}
            onClick={() => document.getElementById(name)?.click()}
        >
            <i className="bi bi-cloud-upload fs-1 text-primary"></i>
            <p className="mb-0 mt-2">Drag & Drop video here or click to select</p>
            <input
                type="file"
                accept="video/*"
                id={name}
                hidden
                onChange={handleFileChange}
            />

            {previewURL && (
                <div className="mt-3">
                    <video controls width="100%">
                        <source src={previewURL} type={videoFile?.type} />
                        Your browser does not support the video tag.
                    </video>
                </div>
            )}
        </div>
    );
};

export default VideoDropzone;
