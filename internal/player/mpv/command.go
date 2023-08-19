package mpv

import (
	"encoding/json"
	"fmt"

	"git.sr.ht/~poldi1405/glog"
)

func (p *MPV) send(cmd command) {
	cmdJSON, _ := json.Marshal(cmd)
	_, err := fmt.Fprintf(p.conn, "%s\n", cmdJSON)
	glog.Tracef("sent: %s: %v", cmdJSON, err)
}

func (p *MPV) getResponse(id int) response {
	for {
		p.responsesMtx.RLock()
		if res, exists := p.responses[id]; exists {
			p.responsesMtx.RUnlock()
			p.responsesMtx.Lock()
			delete(p.responses, id)
			p.responsesMtx.Unlock()

			return res
		}
		p.responsesMtx.RUnlock()
	}
}

type command struct {
	Command   []any `json:"command"`
	RequestID int   `json:"request_id,omitempty"`
}
