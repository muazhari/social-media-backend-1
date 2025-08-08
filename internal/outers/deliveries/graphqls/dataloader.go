package graphqls

import (
	"social-media-backend-1/internal/outers/deliveries/graphqls/model"
	"time"

	"github.com/vikstrous/dataloadgen"
)

type Dataloader struct {
	AccountDataloader *dataloadgen.Loader[string, *model.Account]
}

func NewDataloader(resolver *Resolver) *Dataloader {
	return &Dataloader{
		AccountDataloader: dataloadgen.NewLoader[string, *model.Account](
			resolver.GetAccountsByIDs,
			dataloadgen.WithWait(10*time.Millisecond),
		),
	}
}
