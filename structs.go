package stripack




func Array(dx, dy int) [][]int {
    a := make([][]int, dx) 
    for i := range a {
        a[i] = make([]int, dy) 
        for j := range a[i] {
            a[i][j] = -1
        }
    }
    return a
}

