package moneyconnect

import (
	"fmt"
	"github.com/cyjaysong/lakala-go/moneyconnect/models"
)

func (t *Client) NotifyVerify(notifyBody models.NotifyBody) (verifyPass bool) {
	var str = fmt.Sprintf("%s%s%s%s%s", notifyBody.AppId, notifyBody.Timestamp, notifyBody.Version, notifyBody.EventType, notifyBody.Event)
	return t.Verify(notifyBody.Sign, []byte(str))
}
