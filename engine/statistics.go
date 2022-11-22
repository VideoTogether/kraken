package engine

type Statistics struct {
	RoomCount int `json:"voiceRoomCount"`
}

func statistics(r *R) (*Statistics, error) {
	rooms := r.router.engine.rooms
	rooms.RLock()
	defer rooms.RUnlock()
	return &Statistics{
		RoomCount: len(rooms.m),
	}, nil
}
