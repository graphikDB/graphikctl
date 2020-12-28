package graph

import (
	"context"
	"fmt"
	apipb "github.com/graphikDB/graphik/gen/grpc/go"
	"github.com/graphikDB/graphikctl/helpers"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

func init() {
	Stream.Flags().StringVar(&channel, "channel", "", "the channel to publish a message to")
	Stream.Flags().StringVar(&expression, "expression", "", "CEL expression to filter streamed messages")
}

var Stream = &cobra.Command{
	Use:     "stream",
	Short:   "graphikDB stream operations",
	Example: `graphikctl stream --channel testing --expression "has(this.data.text)"`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helpers.GetClient(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if err := client.Stream(context.Background(), &apipb.StreamFilter{
			Channel:    channel,
			Expression: expression,
		}, func(msg *apipb.Message) bool {
			fmt.Println(protojson.Format(msg))
			return true
		}); err != nil {
			fmt.Println(err.Error())
			return
		}
	},
}
