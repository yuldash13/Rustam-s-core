package main

type Graph struct {
	Name  string
	Nodes []*Graph
}

func dfs(graph *Graph, order []string, visited map[string]bool) []string {
	if visited[graph.Name] == true {
		return order
	}
	visited[graph.Name] = true

	for i := 0; i < len(graph.Nodes); i++ {
		if visited[graph.Nodes[i].Name] != true {
			order = dfs(graph.Nodes[i], order, visited)
		}
	}

	order = append(order, graph.Name)
	return order
}

//socks := Graph{
//Name:  "носки",
//Nodes: nil,
//}
//shoes := Graph{
//Name:  "обувь",
//Nodes: nil,
//}
//jeans := Graph{
//Name:  "джинсы",
//Nodes: nil,
//}
//tShirt := Graph{
//Name:  "футболка",
//Nodes: nil,
//}
//sweater := Graph{
//Name:  "кофта",
//Nodes: nil,
//}
//coat := Graph{
//Name:  "куртка",
//Nodes: nil,
//}
//glasses := Graph{
//Name:  "очки",
//Nodes: nil,
//}
//
//socks.Nodes = []*Graph{&shoes}
//jeans.Nodes = []*Graph{&coat, &sweater, &shoes}
//tShirt.Nodes = []*Graph{&jeans, &sweater}
//sweater.Nodes = []*Graph{&coat}
//
//graphs := []*Graph{&tShirt, &sweater, &coat, &glasses, &socks, &shoes, &jeans}
//
//visited := map[string]bool{}
//order := make([]string, 0)
//
//for _, r := range graphs {
//order = dfs(r, order, visited)
//}
//
//for i := len(order) - 1; i >= 0; i-- {
//fmt.Println(order[i])
//}
