package main

func main() {
	filename := "test.log.gz"
	result := GetMetaInfo(filename)
	PrintArray(result)
}
