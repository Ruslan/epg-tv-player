import React, { useState } from "react";

const videoList = [
  { id: 1, title: "Video 1" },
  { id: 2, title: "Video 2" },
]; // Replace with real data

const Videos = () => {
  const [search, setSearch] = useState("");

  const filteredVideos = videoList.filter((video) =>
    video.title.toLowerCase().includes(search.toLowerCase())
  );

  return (
    <main>
      <h1>Videos</h1>
      <input
        type="text"
        placeholder="Search videos..."
        value={search}
        onChange={(e) => setSearch(e.target.value)}
      />
      <ul>
        {filteredVideos.map((video) => (
          <li key={video.id}>{video.title}</li>
        ))}
      </ul>
    </main>
  );
};

export default Videos;
