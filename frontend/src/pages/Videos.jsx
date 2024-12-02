import React, { useContext, useState, useCallback } from "react";
import { observer } from "mobx-react-lite";
import { StoreContext } from "../main";
import { Link } from "react-router-dom";

const debounce = (func, delay) => {
  let timer;
  return (...args) => {
    clearTimeout(timer);
    timer = setTimeout(() => func(...args), delay);
  };
};

const formatDateRange = (start, stop) => {
  const startDate = new Date(start);
  const stopDate = new Date(stop);

  const optionsTime = { hour: "2-digit", minute: "2-digit" };
  const optionsFull = { month: "short", day: "numeric", hour: "2-digit", minute: "2-digit" };

  const isSameDay =
    startDate.getFullYear() === stopDate.getFullYear() &&
    startDate.getMonth() === stopDate.getMonth() &&
    startDate.getDate() === stopDate.getDate();

  const formatter = new Intl.DateTimeFormat("en-US", isSameDay ? optionsTime : optionsFull);

  const startFormatted = new Intl.DateTimeFormat("en-US", optionsFull).format(startDate);
  const stopFormatted = formatter.format(stopDate);

  return isSameDay ? `${startFormatted} - ${stopFormatted}` : `${startFormatted} - ${stopFormatted}`;
};

const Videos = observer(() => {
  const appStore = useContext(StoreContext);
  const [query, setQuery] = useState(appStore.videosSearchString || ""); // Local state for the input

  const doSearch = useCallback(
    debounce((value) => appStore.loadVideos(value), 500),
    []
  );

  const handleInputChange = (e) => {
    const newQuery = e.target.value; // Read the input value
    setQuery(newQuery); // Update local state
    appStore.setVideosSearchString(newQuery); // Update the search string in the store
    doSearch(newQuery); // Trigger the debounced search
  };

  const handleSearchClick = () => {
    appStore.loadVideos(query); // Immediate search on button click
  };

  return (
    <main>
      <h1>Videos</h1>
      <input
        type="search"
        placeholder="Search videos..."
        value={appStore.videosSearchString}
        onChange={handleInputChange}
      />
      <button onClick={handleSearchClick}>Search</button>
      <ul>
        {appStore.videos.map((video) => {
          const channel = appStore.channels.find((ch) => ch.id === video.ChannelId);

          return (<li key={video.ID}>
            <h3>{video.Title}
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
            </h3>
            <div>{video.Desc}</div>
            <em>{formatDateRange(video.Start, video.Stop)}</em>
          </li>
        )})}
      </ul>
    </main>
  );
});

export default Videos;
