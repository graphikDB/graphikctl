package graph

import (
	"context"
	"encoding/json"
	"fmt"
	apipb "github.com/graphikDB/graphik/gen/grpc/go"
	"github.com/graphikDB/graphikctl/helpers"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	Broadcast.Flags().StringVar(&channel, "channel", "", "the channel to publish a message to")
	Broadcast.Flags().StringVar(&attributes, "data", "", "json attributes of the message to send")
}

var Broadcast = &cobra.Command{
	Use:     "broadcast",
	Short:   "graphikDB broadcast operations",
	Example: `graphikctl broadcast --channel testing --data '{"text": "testing!"}'`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		client, err := helpers.GetClient(ctx)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		var attributeMap = map[string]interface{}{}
		err = json.Unmarshal([]byte(attributes), &attributeMap)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if err := client.Broadcast(ctx, &apipb.OutboundMessage{
			Channel: channel,
			Data:    apipb.NewStruct(attributeMap),
		}); err != nil {
			fmt.Println(err.Error())
			return
		}
	},
}
