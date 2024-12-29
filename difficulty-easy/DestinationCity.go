package difficultyeasy

// Question :

// You are given the array paths, where paths[i] = [cityAi, cityBi] means there exists a direct path going from cityAi to cityBi. Return the destination city, that is, the city without any path outgoing to another city.
// It is guaranteed that the graph of paths forms a line without any loop, therefore, there will be exactly one destination city.
// Example 1:
// Input: paths = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
// Output: "Sao Paulo"
// Explanation: Starting at "London" city you will reach "Sao Paulo" city which is the destination city. Your trip consist of: "London" -> "New York" -> "Lima" -> "Sao Paulo".
// Example 2:
// Input: paths = [["B","C"],["D","B"],["C","A"]]
// Output: "A"
// Explanation: All possible trips are:
// "D" -> "B" -> "C" -> "A".
// "B" -> "C" -> "A".
// "C" -> "A".
// "A".
// Clearly the destination city is "A".
// Example 3:
// Input: paths = [["A","Z"]]
// Output: "Z"
// Constraints:
// 1 <= paths.length <= 100
// paths[i].length == 2
// 1 <= cityAi.length, cityBi.length <= 10
// cityAi != cityBi
// All strings consist of lowercase and uppercase English letters and the space character.

// Answer :
func DestinationCity(paths [][]string) string {
    m := map[string]bool{}
    for i := 0; i < len(paths); i++{
        for j := 0; j < len(paths[i]) - 1; j++{
            m[paths[i][j]] = true
        }
    }
    r := ""
    for i := 0; i < len(paths); i++{
        end := len(paths[i]) - 1
        if !m[paths[i][end]] {
            r = paths[i][end]
        }   
    }
	return r
}


// my ans before:

// 1. attempt : 
// func destCity(paths [][]string) string {
//     m := map[string]bool{}
//     r := ""
//     for i := 0; i < len(paths); i++{
//         for j := 0; j < len(paths[i]) - 1; j++{
//             m[paths[i][j]] = true
//         }
//         last := paths[i][len(paths[i]) - 1]
//         if !m[last]{
//             r = last
//         }
//     }    
// 	return r
// }


// 2. attempt:
// func destCity(paths [][]string) string {
//     return paths[len(paths)- 1] [len(paths[len(paths) - 1]) - 1]
// }



// 3. attempt:
// func destCity(paths [][]string) string {
// 	x := map[string]bool{}
// 	t := []string{}
// 	for _, path := range paths {
// 		for _, p := range path {
// 			if !x[p] {
// 				t = append(t, p)
// 			}
// 			x[p] = true
// 		}	
// 	}
// 	return t[len(t) - 1]
// }