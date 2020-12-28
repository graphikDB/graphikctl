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
	createDocCmd.PersistentFlags().StringVar(&gtype, "gtype", "", "the gtype of the doc")
	createDocCmd.PersistentFlags().StringVar(&gid, "gid", "", "the gid of the doc")
	createDocCmd.PersistentFlags().StringVar(&attributes, "attributes", "", "json attributes of the doc")

	createConnectionCmd.PersistentFlags().StringVar(&gtype, "gtype", "", "the gtype of the connection")
	createConnectionCmd.PersistentFlags().StringVar(&gid, "gid", "", "the gid of the connection")
	createConnectionCmd.PersistentFlags().StringVar(&attributes, "attributes", "", "json attributes of the connection")
	createConnectionCmd.PersistentFlags().StringVar(&from_gid, "from-gid", "", "the gid of the root doc")
	createConnectionCmd.PersistentFlags().StringVar(&from_gtype, "from-gtype", "", "the gtype of the root doc")
	createConnectionCmd.PersistentFlags().StringVar(&to_gid, "to-gid", "", "the gid of the destintion doc")
	createConnectionCmd.PersistentFlags().StringVar(&to_gtype, "to-gtype", "", "the gtype of the destintion doc")
	createConnectionCmd.PersistentFlags().BoolVar(&directed, "directed", false, "is the connection unidirectional?")

	Create.AddCommand(createDocCmd, createConnectionCmd)
}

var Create = &cobra.Command{
	Use:   "create",
	Short: "graphikDB create operations (doc, connection)",
}

var createDocCmd = &cobra.Command{
	Use:     "doc",
	Short:   "create a document",
	Example: "graphikctl create doc --gtype task --attributes '{ \"title\": \"this is a title\", \"description\": \"this is a description\", \"priority\": \"low\"}'",
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
		doc, err := client.CreateDoc(ctx, &apipb.DocConstructor{
			Ref: &apipb.RefConstructor{
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

var createConnectionCmd = &cobra.Command{
	Use:   "connection",
	Short: "create a connection",
	//Example: "graphikctl create connection --gtype category --attributes '{ \"title\": \"this is a title\", \"description\": \"this is a description\", \"priority\": \"low\"}'",
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
		connection, err := client.CreateConnection(ctx, &apipb.ConnectionConstructor{
			Ref: &apipb.RefConstructor{
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
