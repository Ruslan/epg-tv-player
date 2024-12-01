import { makeAutoObservable } from "mobx";
import { FetchChannels } from "../../wailsjs/go/main/App";

class AppStore {
  totalChannels = 0;
  totalVideos = 0;
  channels = [];

  liveStreamUrlTemplate = "";
  vodUrlTemplate = "";

  constructor() {
    makeAutoObservable(this); // Makes state reactive
  }

  setChannels(channels) {
    this.channels = channels;
    this.totalChannels = channels.length;
  }

  setTotalVideos(count) {
    this.totalVideos = count;
  }


  setLiveStreamUrlTemplate(url) {
    this.liveStreamUrlTemplate = url;
  }

  setVodUrlTemplate(url) {
    this.vodUrlTemplate = url;
  }

  reloadChannels() {
    FetchChannels().then((result) => {
      this.setChannels(result.channels)
      this.setTotalVideos(result.totalVideos)
    })
  }
}

const appStore = new AppStore();
export default appStore;
