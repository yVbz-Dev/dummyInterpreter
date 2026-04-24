package main

import "fmt"

func runner(nodes []ASTNode) {
	for _, node := range nodes {
		switch {
		case node.NodeAction == "print":
			fmt.Println(node.NodeArgs["value"])
		}
	}
}
