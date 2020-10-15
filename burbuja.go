package main

func Burbuja(s []int64) {
	var tam int
	for x, _ := range s {
		tam = x
	}
	tam++
	var temporal int64

	for i := 0; i < tam; i++ {
		for j := 0; j < tam-1; j++ {
			if s[j] > s[j+1] {
				temporal = s[j]
				s[j] = s[j+1]
				s[j+1] = temporal
			}
		}
	}
}

func main() {

}
