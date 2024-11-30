import React, { useContext } from "react";
import { observer } from "mobx-react-lite";
import { StoreContext } from "../main";
import { Link } from "react-router-dom";

const Header = observer(() => {
  const appStore = useContext(StoreContext);
  const channelsCount = appStore.totalChannels;
  const videosCount = appStore.totalVideos;

  return (
    <header>
      <nav>
        <Link to="/">Player</Link> |
        <Link to="/">Channels ({channelsCount})</Link> |
        <Link to="/videos">Videos ({videosCount})</Link> |
        <Link to="/settings">Settings</Link>
      </nav>
    </header>
  );
});

export default Header;
