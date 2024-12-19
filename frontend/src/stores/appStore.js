import { makeAutoObservable } from "mobx";
import { FetchChannels, FetchVideos, SetSetting, GetSetting } from "../../wailsjs/go/main/App";

class AppStore {
  totalChannels = 0;
  totalVideos = 0;
  channels = [];

  videos = [];
  videosSearchString = "";

  liveStreamUrlTemplate = "";
  vodUrlTemplate = "";

  constructor() {
    makeAutoObservable(this); // Makes state reactive
    GetSetting("LiveStreamUrlTemplate").then((v) => this.setLiveStreamUrlTemplate(v))
  }

  setChannels(channels) {
    this.channels = channels;
    this.totalChannels = channels.length;
  }

  setTotalVideos(count) {
    this.totalVideos = count;
  }

  setVideos(videos) {
    this.videos = videos
  }

  setVideosSearchString(string) {
    this.videosSearchString = string
  }

  setLiveStreamUrlTemplate(url) {
    this.liveStreamUrlTemplate = url;
    SetSetting("LiveStreamUrlTemplate", url)
  }

  setVodUrlTemplate(url) {
    this.vodUrlTemplate = url;
    SetSetting("VodUrlTemplate", url)
  }

  reloadChannels() {
    FetchChannels().then((result) => {
      this.setChannels(result.channels)
      this.setTotalVideos(result.totalVideos)
    })
  }

  loadVideos(query) {
    FetchVideos({ page: 1, per_page: 50 }, query).then((result) => {
      this.setVideos(result)
    })
  }
}

const appStore = new AppStore();
export default appStore;
