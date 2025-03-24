package main

import (
	"testing"
)

func BenchmarkCountMeats(b *testing.B) {
	text := `Doner flank kielbasa shoulder shankle commodo culpa. Aute brisket cillum, ut pig dolore biltong quis jowl ad est. Tongue aliqua ribeye porchetta ball tip. Proident cupidatat tempor doner.
	Ut nostrud voluptate ad incididunt, ball tip picanha short ribs veniam pancetta. Brisket elit et chicken incididunt swine. Alcatra corned beef doner aute prosciutto in biltong short ribs commodo id pig kielbasa boudin. Ham hock short loin pork ham non.`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CountMeats(text)
	}
}
