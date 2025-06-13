import React, { useContext } from "react";
import { observer } from "mobx-react-lite";
import { useParams } from "react-router-dom";
import { StoreContext } from "../main";
import { Link } from "react-router-dom";
import VideoPlayer from "../components/VideoPlayer";
import { formatDateRange } from '../lib/timeHelpers'

const VideosShow = observer(() => {
  const { slug } = useParams();
  const id = parseInt(slug)
  const appStore = useContext(StoreContext);
  const video = appStore.videos.find((video) => video.ID === id);
  const channel = appStore.channels.find((ch) => ch.id === video.ChannelId);
  window.video = video

  if (!video) {
    return <p>Video not found</p>;
  }

  const streamUrl = appStore.liveStreamUrlTemplate.replace("{tvg_id}", channel.tvg_id) + `?utc=${Date.parse(video.Start) / 1000}`

  return (
    <main>
      <h1>
        <Link to="/videos">&larr;</Link>
        {video.Title}
      </h1>
      <VideoPlayer url={streamUrl} />
      <div>{video.Desc}</div>
      <div>
        {channel && (
          <>
            <img
              src={channel.logo}
              alt={`${channel.title} logo`}
              style={{ maxHeight: "1em", marginLeft: "0.5em" }}
            />
            <Link to={`/channel/${channel.id}`} style={{ marginLeft: "0.5em" }}>
              {channel.title}
            </Link>
          </>
        )}
        <em>{formatDateRange(video.Start, video.Stop)}</em>
      </div>
      <p>{streamUrl}</p>
    </main>
  );
});

export default VideosShow;
