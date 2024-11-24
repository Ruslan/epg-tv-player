package main

import (
	"fmt"
	"iptv/epg"
	"log"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbEpg struct {
	tv          *epg.TV
	videoCache  map[string][]*Video
	channelsIds map[string]uint
	db          *gorm.DB
	newVideos   []*Video
	treshold    time.Time
}

const EPG_URL = "https://af-play.com/storage/epg/xmltv.xml.gz"
const ARCHIVE_DAYS = 7
const maxRetries = 5
const retryTimeout = 5 * time.Second

func (loader *DbEpg) LoadAndParse() {
	loader.LoadEpg()
	loader.ParseChannels()
	loader.ParseProgrammes()
}

func (loader *DbEpg) LoadEpg() {
	var err error

	log.Println("Downloading EPG:", EPG_URL)

	for retries := 0; retries < maxRetries; retries++ {
		loader.tv, err = epg.ParseEPG(EPG_URL)
		if err == nil {
			// Successfully parsed EPG, break out of the loop
			break
		}

		log.Printf("Error parsing EGP %s: %v. Retrying in %v...", EPG_URL, err, retryTimeout)
		time.Sleep(retryTimeout)
	}

	if err != nil {
		return
	}

	log.Println("EPG Downloaded. Channels:", len(loader.tv.Channels), "Programs:", len(loader.tv.Programmes))
}

func (loader *DbEpg) ParseChannels() {
	if len(loader.tv.Channels) == 0 {
		log.Println("No channels")
		return
	}

	loader.channelsIds = make(map[string]uint)

	for num, channel := range loader.tv.Channels {
		newChannel := Channel{
			Title:    channel.DisplayName.Value,
			Position: num,
			TvgID:    channel.ID,
			TvgLogo:  channel.Icon.Src,
		}
		existingChannel := Channel{}
		result := loader.db.Where("tvg_id = ?", newChannel.TvgID).First(&existingChannel)

		// If the channel exists, update its fields
		if result.RowsAffected > 0 {
			existingChannel.Title = channel.DisplayName.Value
			existingChannel.TvgLogo = channel.Icon.Src

			// Update the channel in the database
			updateResult := loader.db.Save(&existingChannel)
			loader.channelsIds[existingChannel.TvgID] = existingChannel.ID

			// Check for errors during the update
			if updateResult.Error != nil {
				fmt.Println("Error occurred during update:", updateResult.Error)
			}
		} else {
			// Create the new channel in the database
			createResult := loader.db.Create(&newChannel)

			// Check for errors during the create operation
			if createResult.Error != nil {
				fmt.Println("Error occurred during create:", createResult.Error)
			} else {
				loader.channelsIds[newChannel.TvgID] = newChannel.ID
			}
		}
	}
	log.Println("Channels loaded to DB")
}

func (loader *DbEpg) ParseProgrammes() {
	loader.treshold = time.Now().AddDate(0, 0, -ARCHIVE_DAYS)
	loader.db.Delete(&Video{}, "stop < ?", loader.treshold)

	if len(loader.tv.Programmes) == 0 {
		log.Println("No Programmes")
		return
	}

	minTime, err := loader.tv.Programmes[0].GetStart()
	if err != nil {
		return
	}
	maxTime, err := loader.tv.Programmes[0].GetStop()
	if err != nil {
		return
	}

	for _, programme := range loader.tv.Programmes {
		start, err := programme.GetStart()
		if err != nil {
			continue
		}
		if start.Before(minTime) {
			minTime = start
		}
		stop, err := programme.GetStop()
		if err != nil {
			continue
		}
		if stop.After(maxTime) {
			maxTime = stop
		}
	}

	// TODO: load channel-by-channel
	// TODO: remove not actual programs
	loader.videoCache = make(map[string][]*Video)
	var rawVideoCache []*Video
	loader.db.Where("start >= ? and stop <= ?", minTime, maxTime).Select("start", "stop", "channel_code").Find(&rawVideoCache)

	for _, video := range rawVideoCache {
		loader.videoCache[video.ChannelCode] = append(loader.videoCache[video.ChannelCode], video)
	}

	for _, programme := range loader.tv.Programmes {
		loader.LoadProgramme(&programme)
	}

	log.Println("Channels loaded to Batch")
	loader.flushVideos(true)
	log.Println("Channels loaded to DB")
}

func (loader *DbEpg) LoadProgramme(programme *epg.Programme) {
	// log.Println("loading video")
	start, err := programme.GetStart()
	if err != nil {
		return
	}
	stop, err := programme.GetStop()
	if err != nil {
		return
	}
	if stop.Before(loader.treshold) {
		return
	}

	newVideo := Video{
		Title:       programme.Title.Value,
		TitleLower:  strings.ToLower(programme.Title.Value),
		Desc:        programme.Desc.Value,
		DescLower:   strings.ToLower(programme.Desc.Value),
		Start:       start,
		Stop:        stop,
		ChannelCode: programme.Channel,
		ChannelId:   loader.channelsIds[programme.Channel],
	}
	if newVideo.ChannelId == 0 {
		log.Fatalf("%v is zero channeId", newVideo)
	}
	var existingVideo *Video

	// result := loader.db.Where("start = ? and stop = ? and channel_code = ?", newVideo.Start, newVideo.Stop, newVideo.ChannelCode).First(&existingVideo)
	// log.Println("It exists?")
	for _, cachedVideo := range loader.videoCache[newVideo.ChannelCode] {
		if cachedVideo.Start.Equal(start) && cachedVideo.Stop.Equal(stop) {
			existingVideo = cachedVideo
			break
		}
	}
	// log.Println("It exists = ", existingVideo)
	// If the channel exists, update its fields
	if existingVideo != nil { // do not update
	} else {
		loader.newVideos = append(loader.newVideos, &newVideo)
		loader.flushVideos(false)
	}
}

func (loader *DbEpg) flushVideos(force bool) {
	if force && len(loader.newVideos) > 0 || len(loader.newVideos) >= 1000 {
		log.Println("saving batch with size", len(loader.newVideos))
		loader.db.Logger.LogMode(logger.Silent)
		// Create the new channel in the database
		createResult := loader.db.Create(&loader.newVideos)

		// // Check for errors during the create operation
		if createResult.Error != nil {
			fmt.Println("Error occurred during create:", createResult.Error)
		}
		loader.newVideos = loader.newVideos[:0]
		loader.db.Logger.LogMode(logger.Info)
	}
}
