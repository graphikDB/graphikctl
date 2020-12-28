package graph

import (
	"context"
	"fmt"
	apipb "github.com/graphikDB/graphik/gen/grpc/go"
	"github.com/graphikDB/graphikctl/helpers"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
	"time"
)

func init() {
	Traverse.Flags().StringVar(&gtype, "root-gtype", "", "gtype of the root document of the traversal")
	Traverse.Flags().StringVar(&gid, "root-gid", "", "gid of the root document of the traversal")
	Traverse.Flags().StringVar(&docExpression, "doc-expression", "", "CEL expression used to determine which documents to return")
	Traverse.Flags().StringVar(&connectionExpression, "connection-expression", "", "CEL expression used to determine which connections to traverse")
	Traverse.Flags().StringVar(&sort, "sort", "", "sort documents on this attribute")
	Traverse.Flags().StringVar(&algorithm, "algorithm", apipb.Algorithm_DFS.String(), "traversal algorithm to use (DFS/BFS)")
	Traverse.Flags().IntVar(&limit, "limit", 0, "limit returned docs")
	Traverse.Flags().BoolVar(&reverse, "reverse", false, "traversal algorithm to use (DFS/BFS)")

}

var Traverse = &cobra.Command{
	Use:   "traverse",
	Short: "graphikDB traversal operations",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		client, err := helpers.GetClient(ctx)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		resp, err := client.Traverse(ctx, &apipb.TraverseFilter{
			Root: &apipb.Ref{
				Gtype: gtype,
				Gid:   gid,
			},
			DocExpression:        docExpression,
			ConnectionExpression: connectionExpression,
			Limit:                uint64(limit),
			Sort:                 sort,
			Reverse:              reverse,
			Algorithm:            apipb.Algorithm(apipb.Algorithm_value[algorithm]),
			MaxDepth:             uint64(maxDepth),
			MaxHops:              uint64(maxHops),
		})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(protojson.Format(resp))
	},
}
