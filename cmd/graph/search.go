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
	Search.PersistentFlags().StringVar(&gtype, "gtype", "", "gtype of doc/connection to search for")
	Search.PersistentFlags().StringVar(&expression, "expression", "", "CEL filter expression")
	Search.PersistentFlags().IntVar(&limit, "limit", 0, "limit number of returned docs/connections")
	Search.PersistentFlags().StringVar(&sort, "sort", "", "sort returned docs/connections")
	Search.PersistentFlags().StringVar(&index, "index", "", "search in a specific index")
	Search.PersistentFlags().StringVar(&seek, "seek", "", "search in a specific index")
	Search.PersistentFlags().BoolVar(&reverse, "reverse", false, "reverse returned docs/connections")
	Search.AddCommand(searchDocs, searchConnections)
}

var Search = &cobra.Command{
	Use:   "search",
	Short: "graphikDB search operations  (docs, connections)",
}

var searchDocs = &cobra.Command{
	Use:   "docs",
	Short: "search for documents",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		client, err := helpers.GetClient(ctx)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		docs, err := client.SearchDocs(ctx, &apipb.Filter{
			Gtype:      gtype,
			Expression: expression,
			Limit:      uint64(limit),
			Sort:       sort,
			Seek:       seek,
			Reverse:    reverse,
			Index:      index,
		})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(protojson.Format(docs))
	},
}

var searchConnections = &cobra.Command{
	Use:   "connections",
	Short: "search for connections",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		client, err := helpers.GetClient(ctx)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		connections, err := client.SearchConnections(ctx, &apipb.Filter{
			Gtype:      gtype,
			Expression: expression,
			Limit:      uint64(limit),
			Sort:       sort,
			Seek:       seek,
			Reverse:    reverse,
			Index:      index,
		})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(protojson.Format(connections))
	},
}
