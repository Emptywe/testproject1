package randomExamples

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Generate(input bool) string {

	rand.Seed(time.Now().Unix())
	if input {
		letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

		var res []string

		for i := 0; i < rand.Intn(5)+2; i++ {
			length := rand.Intn(5) + 1
			var b strings.Builder
			for i := 0; i < length; i++ {
				b.WriteRune(letters[rand.Intn(len(letters))])
			}
			str := b.String()

			block := fmt.Sprintf("%d-%s", rand.Intn(1000), str)
			res = append(res, block)
		}

		return strings.Join(res, "-")
	} else {
		letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

		var res []string

		for i := 0; i < rand.Intn(5)+2; i++ {
			length := rand.Intn(5) + 1
			var st strings.Builder
			for i := 0; i < length; i++ {
				st.WriteRune(letters[rand.Intn(len(letters))])
			}
			str := st.String()
			cond := rand.Intn(5)
			switch cond {
			case 0:
				var s strings.Builder
				for i := 0; i < length; i++ {
					st.WriteRune(letters[rand.Intn(len(letters))])
				}
				str1 := s.String()
				block := fmt.Sprintf("%s-%s", str, str1)
				res = append(res, block)
			case 1:
				block := fmt.Sprintf("%d-%d", rand.Intn(1000), rand.Intn(1000))
				res = append(res, block)
			case 2:
				block := fmt.Sprintf("%s-%d", str, rand.Intn(1000))
				res = append(res, block)

			case 3:
				block := fmt.Sprintf("%s", str)
				res = append(res, block)
			case 4:
				block := fmt.Sprintf("%d", rand.Intn(1000))
				res = append(res, block)
			}
		}

		return strings.Join(res, "-")
	}

	return ""
}
