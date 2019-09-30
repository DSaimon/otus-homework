package analysis

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAnaliseFrequency(t *testing.T) {
	text := "мама мыла раму раму раму мыла"
	expected := []string{"раму", "мыла", "мама"}
	require.Equal(t, expected, AnaliseFrequency(text), "Frequency analise " + text)

	text2 := "частотный частотный частотный анализ анализ тектса"
	expected2 := []string{"частотный", "анализ", "тектса"}
	require.Equal(t, expected2, AnaliseFrequency(text2), "Frequency analise " + text2)
}