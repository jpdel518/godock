package foo

// スコープ
const (
    // 頭文字が大文字だとpublic
    Max = 100
    // 頭文字が小文字だとprivate
    min = 1
)

func ReturnMin() int {
    return min
}
