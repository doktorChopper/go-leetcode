package array

// complexity:  O(n^2)
// memmory:     O(1) 

func twoSumSol1(nums []int, target int) []int {

    for i := 0; i < len(nums); i++ {
        for j := i + 1; i < len(nums); j++ {
            if (nums[i] + nums[j]) == target {
                return []int{i, j}
            }
        }
    }
    
    return []int{}
}

// HashTable impl

// complexity:  O(n)
// memmory:     O(n)

func twoSumSol2(nums []int, target int) []int {

    m := make(map[int]int)

    for i, v := range nums {
        m[v] = i
    }
    
    for i, v := range nums {
        diff := target - v
        if j, ok := m[diff]; ok && j != i {
            return []int{i, j}
        }
    }

    return []int{}
}

// also HT impl, but with one loop

// complexity:  O(n)
// memmory:     O(n)

func twoSumSol3(nums []int, target int) []int {

    m := make(map[int]int)

    for i, v := range nums {
        diff := target - v
        if j, ok := m[diff]; ok {
            return []int{i, j}
        }
        m[v] = i
    }
    
    return []int{}
}
