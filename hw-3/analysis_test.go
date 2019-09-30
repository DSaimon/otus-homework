package analysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAnaliseFrequency(t *testing.T) {
	text := "мама мыла раму раму раму мыла"
	expected := []string{"раму", "мыла", "мама"}
	require.Equal(t, expected, AnaliseFrequency(text), "Frequency analise "+text)

	text2 := "частотный частотный частотный анализ анализ текста"
	expected2 := []string{"частотный", "анализ", "текста"}
	require.Equal(t, expected2, AnaliseFrequency(text2), "Frequency analise "+text2)
}
