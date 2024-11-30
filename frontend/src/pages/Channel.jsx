import React, { useContext } from "react";
import { observer } from "mobx-react-lite";
import { useParams } from "react-router-dom";
import { StoreContext } from "../main";
import VideoPlayer from "../components/VideoPlayer";

const Channel = observer(() => {
  const { slug } = useParams();
  const id = parseInt(slug)
  const appStore = useContext(StoreContext);
  const channel = appStore.channels.find((ch) => ch.id === id);
  window.channel = channel
  window.channelID = id

  if (!channel) {
    return <p>Channel not found</p>;
  }

  const streamUrl = appStore.liveStreamUrlTemplate.replace("{tvg_id}", channel.tvg_id)

  return (
    <main>
      <h1>{channel.title}</h1>
      <VideoPlayer url={streamUrl} />
    </main>
  );
});

export default Channel;
