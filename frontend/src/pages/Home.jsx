import React, { useContext } from "react";
import { observer } from "mobx-react-lite"
import { StoreContext } from "../main";
import { Link } from "react-router-dom";

const Home = observer(() => {
  const appStore = useContext(StoreContext);
  const channels = appStore.channels;

  return (
    <main>
      <h1>Channels</h1>
      <ul>
        {channels.map((channel) => (
          <li key={channel.id}>
            <img src={channel.logo} style={{maxHeight: "1em"}}></img><Link to={`/channel/${channel.id}`}>{channel.title}</Link>
          </li>
        ))}
      </ul>
    </main>
  );
});

export default Home;
