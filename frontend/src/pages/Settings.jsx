import React, { useContext } from "react";
import { observer } from "mobx-react-lite";
import { StoreContext } from "../main";

const Settings = observer(() => {
  const appStore = useContext(StoreContext);

  const handleReload = () => {
    appStore.reloadChannels();
  };

  const handleLiveStreamChange = (e) => {
    appStore.setLiveStreamUrlTemplate(e.target.value);
  };

  const handleVodChange = (e) => {
    appStore.setVodUrlTemplate(e.target.value);
  };

  return (
    <main>
      <h1>Settings</h1>
      <button onClick={handleReload}>Reload Videos</button>

      <div>
        <label>
          Live Stream URL Template:
          <input
            type="text"
            value={appStore.liveStreamUrlTemplate}
            onChange={handleLiveStreamChange}
            placeholder="Enter Live Stream URL Template"
          />
        </label>
      </div>

      <div>
        <label>
          VOD URL Template:
          <input
            type="text"
            value={appStore.vodUrlTemplate}
            onChange={handleVodChange}
            placeholder="Enter VOD URL Template"
          />
        </label>
      </div>
    </main>
  );
});

export default Settings;
