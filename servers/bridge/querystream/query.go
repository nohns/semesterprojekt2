package querystream

import bridgepb "github.com/nohns/semesterprojekt2/proto/gen/go/cloud/bridge/v1"

type query struct {
	id   string
	send chan<- *bridgepb.StreamResponse
}

func (q *query) execute(sq *bridgepb.StreamQuery) {

}
