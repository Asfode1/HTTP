package service

import (
    "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
    "strings"
    "errors"
)

// Convert автоматически определяет формат данных и конвертирует их
func Convert(data string) (string, error) {
    if data == "" {
        return "", errors.New("пустые входные данные")
    }
    
    if isMorseCode(data) {
        return morse.ToText(data), nil
    }
    return morse.ToMorse(data), nil
}

// isMorseCode проверяет, содержит ли строка только символы Морзе
func isMorseCode(s string) bool {
    allowed := ".-/ "
    for _, r := range s {
        if !strings.ContainsRune(allowed, r) && r != ' ' {
            return false
        }
    }
    return true
}

