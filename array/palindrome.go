package array

import "math"

func isPalindromeSol1(x int) bool {
    if x < 0 {
        return false
    }

    l := int(math.Log10(float64(x)))

    for i := 0; i <= (l - 1) / 2; i++ {
        if x / int(math.Pow10(l - i)) % 10 != x / int(math.Pow10(i)) % 10 {
            return false
        }
    }

    return true
}

func isPalindromeSol2(x int) bool {

    if x < 0 {
        return false
    }

    reverse := 0
    xcopy := x

    for x > 0 {
        reverse = reverse * 10 + x % 10
        x /= 10
    }

    return xcopy == reverse
}
