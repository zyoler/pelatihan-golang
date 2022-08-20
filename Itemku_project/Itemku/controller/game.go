package controller

import (
	"context"
	"encoding/json"
	"net/http"
)

func (c *ctrl) GetDataGame(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	redisData, err := REDIS.Get(ctx, "game").Result()
	if err != nil {
		listuser, err := c.us.GetDataGame()
		if err != nil {
			ResponseApi(w, 500, nil, "Internal Server Error")
			return
		}
		datajson, err := json.Marshal(listuser)
		if err != nil {
			ResponseApi(w, 500, nil, "Internal Server Error")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(redisData))
	w.WriteHeader(200)
}
