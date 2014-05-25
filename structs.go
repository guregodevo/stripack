package stripack




func Array(dx, dy int) [][]int64 {
    a := make([][]int64, dx) 
    for i := range a {
        a[i] = make([]int64, dy) 
        for j := range a[i] {
            a[i][j] = -1
        }
    }
    return a
}

