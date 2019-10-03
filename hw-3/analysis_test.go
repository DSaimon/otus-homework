package analysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAnaliseFrequency(t *testing.T) {
	text := "мама мыла раму раму раму мыла"
	expected := []string{"раму", "мыла", "мама"}
	require.Equal(t, expected, AnaliseFrequency(text), "Frequency analise "+text)

	text2 := "частотный, частотный! частотный. анализ; анализ- текста'"
	expected2 := []string{"частотный", "анализ", "текста"}
	require.Equal(t, expected2, AnaliseFrequency(text2), "Frequency analise "+text2)

	text3 := "Ехал Грека через реку, Видит Грека – в реке рак. Сунул Грека руку в реку, - Рак за руку Греку ЦАП !"
	expected3 := []string{"грека", "в", "рак", "реку", "руку", "видит", "греку", "ехал", "за", "реке"}
	require.Equal(t, expected3, AnaliseFrequency(text3), "Frequency analise "+text3)
}
