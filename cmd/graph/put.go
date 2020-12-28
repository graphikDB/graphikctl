package graph

import (
	"context"
	"encoding/json"
	"fmt"
	apipb "github.com/graphikDB/graphik/gen/grpc/go"
	"github.com/graphikDB/graphikctl/helpers"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
	"time"
)

func init() {
	putDocCmd.PersistentFlags().StringVar(&gtype, "gtype", "", "the gtype of the doc")
	putDocCmd.PersistentFlags().StringVar(&gid, "gid", "", "the gid of the doc")
	putDocCmd.PersistentFlags().StringVar(&attributes, "attributes", "", "json attributes of the doc")

	putConnectionCmd.PersistentFlags().StringVar(&gtype, "gtype", "", "the gtype of the connection")
	putConnectionCmd.PersistentFlags().StringVar(&gid, "gid", "", "the gid of the connection")
	putConnectionCmd.PersistentFlags().StringVar(&attributes, "attributes", "", "json attributes of the connection")
	putConnectionCmd.PersistentFlags().StringVar(&from_gid, "from-gid", "", "the gid of the root doc")
	putConnectionCmd.PersistentFlags().StringVar(&from_gtype, "from-gtype", "", "the gtype of the root doc")
	putConnectionCmd.PersistentFlags().StringVar(&to_gid, "to-gid", "", "the gid of the destintion doc")
	putConnectionCmd.PersistentFlags().StringVar(&to_gtype, "to-gtype", "", "the gtype of the destintion doc")
	putConnectionCmd.PersistentFlags().BoolVar(&directed, "directed", false, "is the connection unidirectional?")

	Put.AddCommand(putDocCmd, putConnectionCmd)
}

var Put = &cobra.Command{
	Use:   "put",
	Short: "graphikDB put operations (doc, connection)",
}

var putDocCmd = &cobra.Command{
	Use:     "doc",
	Short:   "create-or-replace a document",
	Example: "graphikctl put doc --gtype task --attributes '{ \"title\": \"this is a title\", \"description\": \"this is a description\", \"priority\": \"low\"}'",
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
		doc, err := client.PutDoc(ctx, &apipb.Doc{
			Ref: &apipb.Ref{
				Gtype: gtype,
				Gid:   gid,
			},
			Attributes: apipb.NewStruct(attributeMap),
		})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(protojson.Format(doc))
	},
}

var putConnectionCmd = &cobra.Command{
	Use:   "connection",
	Short: "create-or-replace a connection",
	//Example: "graphikctl put connection --gtype category --attributes '{ \"title\": \"this is a title\", \"description\": \"this is a description\", \"priority\": \"low\"}'",
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
		connection, err := client.PutConnection(ctx, &apipb.Connection{
			Ref: &apipb.Ref{
				Gtype: gtype,
				Gid:   gid,
			},
			Attributes: apipb.NewStruct(attributeMap),
			Directed:   directed,
			From: &apipb.Ref{
				Gtype: from_gtype,
				Gid:   from_gid,
			},
			To: &apipb.Ref{
				Gtype: to_gtype,
				Gid:   to_gid,
			},
		})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(protojson.Format(connection))
	},
}
