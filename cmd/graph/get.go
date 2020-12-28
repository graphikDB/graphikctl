package graph

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	apipb "github.com/graphikDB/graphik/gen/grpc/go"
	"github.com/graphikDB/graphikctl/helpers"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
	"time"
)

func init() {
	Get.PersistentFlags().StringVar(&gtype, "gtype", "", "the gtype of the doc/connection")
	Get.PersistentFlags().StringVar(&gid, "gid", "", "the gid of the doc/connection")
	Get.AddCommand(docsCmd, connectionCmd, schemaCmd)
}

var Get = &cobra.Command{
	Use:   "get",
	Short: "graphikDB get operations (doc, connection, schema)",
}

var docsCmd = &cobra.Command{
	Use:     "doc",
	Short:   "get a document",
	Example: "graphikctl get doc --gtype task --gid 1mGJsm5AFOsAKqUZXZu8OCXuIci",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		client, err := helpers.GetClient(ctx)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		doc, err := client.GetDoc(ctx, &apipb.Ref{
			Gtype: gtype,
			Gid:   gid,
		})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(protojson.Format(doc))
	},
}

var connectionCmd = &cobra.Command{
	Use:     "connection",
	Short:   "get a connection",
	Example: "graphikctl get connection --gtype edited --gid 1mGJsm5AFOsAKqUZXZu8OCXuIci",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		client, err := helpers.GetClient(ctx)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		doc, err := client.GetConnection(ctx, &apipb.Ref{
			Gtype: gtype,
			Gid:   gid,
		})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(protojson.Format(doc))
	},
}

var schemaCmd = &cobra.Command{
	Use:     "schema",
	Short:   "get the database schema",
	Example: "graphikctl get schema",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		client, err := helpers.GetClient(ctx)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		schema, err := client.GetSchema(ctx, &empty.Empty{})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(protojson.Format(schema))
	},
}
