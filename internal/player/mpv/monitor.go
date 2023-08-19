package mpv

import (
	"bufio"
	"encoding/json"
	"io"
	"math/rand"
	"time"

	"git.sr.ht/~poldi1405/glog"
)

func (p *MPV) monitor() {
	defer p.conn.Close()
	defer close(p.notifySeek)
	defer close(p.notifyPause)

	rd := bufio.NewReader(p.conn)
	for {
		// message, err := bufio.NewReader(p.conn).ReadBytes('\n')
		message, err := rd.ReadBytes('\n')
		if err == io.EOF {
			return
		}
		var res response
		err = json.Unmarshal(message, &res)
		glog.Tracef("received message: %s", message)
		if err != nil {
			glog.Warnf("received non-understood message %q", message)
			continue
		}

		switch {
		case res.Event == "seek":
			select {
			case p.notifySeekInternal <- struct{}{}:
			default:
			}
		case res.Event == "idle":
			select {
			case p.notifyIdle <- struct{}{}:
			default:
			}
		case res.Event == "end-file":
			if res.Reason == "quit" {
				// TODO: kill parent
				return
			}
		case res.RequestID != 0:
			p.responsesMtx.Lock()
			p.responses[res.RequestID] = res
			p.responsesMtx.Unlock()
		}
	}
}

func (p *MPV) pollPause() {
	var pauseState bool
	for {
		<-time.After(50 * time.Millisecond)
		req := rand.Int()
		p.send(command{
			Command:   []any{"get_property", "pause"},
			RequestID: req,
		})
		res := p.getResponse(req)
		if pause, ok := res.Data.(bool); ok {
			if pause != pauseState {
				pauseState = pause
				select {
				case p.notifyPause <- pauseState:
				default:
				}
			}
		}
	}
}

func (p *MPV) handleSeekEvents() {
	for range p.notifySeekInternal {
		pos, err := p.GetPlaybackPos()
		if err == nil {
			select {
			case p.notifySeek <- pos:
			default:
			}
		}
	}
}

type response struct {
	Error     string `json:"error"`
	Event     string `json:"event"`
	ID        int    `json:"id"`
	Data      any    `json:"data"`
	Name      string `json:"name"`
	RequestID int    `json:"request_id"`
	Reason    string `json:"reason"`
}
