package analysis

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAnaliseFrequency(t *testing.T) {
	text := "мама мыла раму раму раму мыла"
	expected := []string{"раму", "мыла", "мама"}
	require.Equal(t, expected, AnaliseFrequency(text), "Frequency analise " + text)

	text2 := "Частотный анализ текста, это анализ текста, при котором считается количество вхождений слов в текст"
	expected2 := []string{"анализ", "текста", "частотный", "это", "при", "котором", "считается", "количество", "вхождений", "слов"}
	require.Equal(t, expected2, AnaliseFrequency(text2), "Frequency analise " + text2)
}