package nats_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/itxor/tgsite/pkg/nats"
	"github.com/stretchr/testify/assert"

	nats_service "github.com/itxor/tgsite/internal/service/nats"
)

func TestDispatch(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	subject := "subject"
	object := "object"

	natsClient := nats.NewMockNatsClientInterface(ctrl)
	natsClient.
		EXPECT().
		Connect().
		Return(func() {}, nil)
	natsClient.
		EXPECT().
		Publish(gomock.Eq(subject), gomock.Eq(object)).
		Return(nil)

	service := nats_service.NewNatsTgUpdateService(natsClient)
	err := service.Dispatch(subject, object)

	assert.Nil(t, err)
}
