// https://leetcode.com/problems/search-suggestions-system/

// 1. Greedy Solution

// func suggestedProducts(products []string, searchWord string) [][]string {
//     sort.Strings(products)

//     res := make([][]string, len(searchWord))

//     for i, _ := range searchWord {
//         subst := searchWord[0:i + 1]
//         var partRes []string

//         for _, product := range products {
//             if len(partRes) == 3 {
//                 break;
//             }
//             if strings.HasPrefix(product, subst) {
//                 partRes = append(partRes, product)
//                 res[i] = partRes
//             }
//         }
//     }
//     return res
// }

// 2. Binary Search
func suggestedProducts(products []string, searchWord string) [][]string {
    var j int
    res := make([][]string, len(searchWord))
    sort.Strings(products)

    for i := range searchWord {
        prefix := searchWord[:i+1]
        j = j + sort.SearchStrings(products[j:], prefix)

        var cur []string
        for k := 0; k < 3 && j+k < len(products); k++ {
            if strings.HasPrefix(products[j+k], prefix) {
                cur = append(cur, products[j+k])
            }
        }
        res[i] = cur
    }
    return res
}
