package main

import "fmt"

func runner(nodes []ASTNode) {
	// iterate node by node
	for _, node := range nodes {
		switch {
		case node.NodeAction == KW_PRINT:
			var varInMemory any = memory["var_"+node.NodeArgs["value"].(string)]
			if varInMemory != nil {
				fmt.Println(varInMemory)
			} else {
				fmt.Println(node.NodeArgs["value"])
			}
		}
	}
}
