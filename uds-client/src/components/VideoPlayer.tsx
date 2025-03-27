import React, { useRef, useEffect } from 'react';
import Hls from 'hls.js';

interface VideoPlayerProps {
    videoUrl: string;
    posterUrl?: string;
}

const VideoPlayer: React.FC<VideoPlayerProps> = ({ videoUrl, posterUrl }) => {
    const videoRef = useRef<HTMLVideoElement | null>(null);

    useEffect(() => {
        let hls: Hls | undefined;

        if (videoRef.current) {
            if (Hls.isSupported()) {
                hls = new Hls();
                hls.loadSource(videoUrl);
                hls.attachMedia(videoRef.current);

                hls.on(Hls.Events.MANIFEST_PARSED, () => {
                    console.log('HLS manifest loaded');
                });
            } else if (videoRef.current.canPlayType('application/vnd.apple.mpegurl')) {
                // Safari or browsers with native HLS support
                videoRef.current.src = videoUrl;
            }
        }

        return () => {
            hls?.destroy();
        };
    }, [videoUrl]);

    return (
        <video
            ref={videoRef}
            controls
            poster={posterUrl}
            width={640}
            height={360}
            style={{ borderRadius: '10px', boxShadow: '0 4px 8px rgba(0,0,0,0.2)' }}
        />
    );
};

export default VideoPlayer;
