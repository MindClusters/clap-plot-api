package main

import "github.com/MindClusters/clap-plot-api/routes"

func main() {
	r := routes.Initialize()
	r.Run(":4000")
}
