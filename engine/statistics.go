package engine

type Statistics struct {
	RoomCount         int         `json:"voiceRoomCount"`
	VoicePeerCount    int         `json:"voicePeerCount"`
	VoicePeerCountMap map[int]int `json:"VoicePeerCountMap"`
}

func statistics(r *R) (*Statistics, error) {
	rooms := r.router.engine.rooms
	rooms.RLock()
	defer rooms.RUnlock()
	voicePeerCountMap := make(map[int]int)
	voicePeerCount := 0
	for _, v := range rooms.m {
		peerCount := len(v.m)
		voicePeerCount += peerCount
		voicePeerCountMap[peerCount]++
	}
	return &Statistics{
		RoomCount:         len(rooms.m),
		VoicePeerCount:    voicePeerCount,
		VoicePeerCountMap: voicePeerCountMap,
	}, nil
}
