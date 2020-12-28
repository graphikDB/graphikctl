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
	editDocCmd.PersistentFlags().StringVar(&gtype, "gtype", "", "the gtype of the doc")
	editDocCmd.PersistentFlags().StringVar(&gid, "gid", "", "the gid of the doc")
	editDocCmd.PersistentFlags().StringVar(&attributes, "attributes", "", "json attributes of the doc")

	editConnectionCmd.PersistentFlags().StringVar(&gtype, "gtype", "", "the gtype of the connection")
	editConnectionCmd.PersistentFlags().StringVar(&gid, "gid", "", "the gid of the connection")
	editConnectionCmd.PersistentFlags().StringVar(&attributes, "attributes", "", "json attributes of the connection")

	Edit.AddCommand(editDocCmd, editConnectionCmd)
}

var Edit = &cobra.Command{
	Use:   "edit",
	Short: "graphikDB edit operations (doc, connection)",
}

var editDocCmd = &cobra.Command{
	Use:   "doc",
	Short: "edit a document",
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
		doc, err := client.EditDoc(ctx, &apipb.Edit{
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

var editConnectionCmd = &cobra.Command{
	Use:   "connection",
	Short: "edit a connection",
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
		connection, err := client.EditConnection(ctx, &apipb.Edit{
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
		fmt.Println(protojson.Format(connection))
	},
}
