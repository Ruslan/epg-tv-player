import React, { useEffect, useRef } from "react";
import Hls from "hls.js";

const VideoPlayer = ({ url }) => {
  const videoRef = useRef(null);

  if (!url) {
    return <p>URL for stream is not found</p>;
  }

  useEffect(() => {
    if (Hls.isSupported()) {
      const hls = new Hls();
      hls.loadSource(url);
      hls.attachMedia(videoRef.current);

      return () => hls.destroy();
    }
  }, [url]);

  return <video ref={videoRef} controls style={{ width: "100%" }} />;
};

export default VideoPlayer;
