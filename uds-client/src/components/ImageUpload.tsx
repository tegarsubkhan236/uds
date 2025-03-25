import React, { useState, DragEvent, ChangeEvent } from 'react';

interface ImageDropzoneProps {
    name: string;
    onFileChange: (file: File | null) => void;
}

const ImageDropzone: React.FC<ImageDropzoneProps> = ({ name, onFileChange }) => {
    const [previewURL, setPreviewURL] = useState<string | null>(null);
    const [isDragging, setIsDragging] = useState(false);
    const [imageFile, setImageFile] = useState<File | null>(null);

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
        if (file && file.type.startsWith('image/')) {
            setImageFile(file);
            setPreviewURL(URL.createObjectURL(file));
            onFileChange(file);
        } else {
            setImageFile(null);
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
            <i className="bi bi-cloud-upload fs-1 text-success"></i>
            <p className="mb-0 mt-2">Drag & Drop image here or click to select</p>
            <input
                type="file"
                accept="image/*"
                id={name}
                hidden
                onChange={handleFileChange}
            />

            {previewURL && (
                <div className="mt-3">
                    <img
                        src={previewURL}
                        alt="Preview"
                        className="img-fluid rounded"
                        style={{ maxHeight: '300px' }}
                    />
                </div>
            )}
        </div>
    );
};

export default ImageDropzone;
