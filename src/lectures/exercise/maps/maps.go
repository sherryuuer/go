//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printStatus(servers map[string]int) {
	fmt.Println("------")
	fmt.Println("Total num of servers is", len(servers))
	statusCount := make(map[int]int)
	for _, status := range servers {
		statusCount[status] += 1
	}
	// 这个写法比较酷，同时const的东西可以直接用，也很酷，这个应该是吸收了java的功能
	fmt.Println(statusCount[Online], "servers are online")
	fmt.Println(statusCount[Offline], "servers are offline")
	fmt.Println(statusCount[Maintenance], "servers are undergoint maintenance")
	fmt.Println(statusCount[Retired], "servers are retired")
	fmt.Println("------")
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	// create a map for servers
	serversMap := make(map[string]int)
	for _, server := range servers {
		// 定数可以直接使用
		serversMap[server] = Online
	}
	fmt.Println(serversMap)
	//  - call display server info function
	printStatus(serversMap)
	//  - change server status of `darkstar` to `Retired`
	serversMap["darkstar"] = Retired
	//  - change server status of `aiur` to `Offline`
	serversMap["aiur"] = Offline
	//  - call display server info function
	printStatus(serversMap)
	//  - change server status of all servers to `Maintenance`
	for server, status := range serversMap {
		fmt.Println("Change status of", server, "from", status, "to", Maintenance)
		serversMap[server] = Maintenance
	}
	printStatus(serversMap)
	//  - call display server info function
}
