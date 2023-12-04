package main

import (
	"fmt"
	"sync"
	"time"
)

// TurnstileCounter представляет счетчик людей, проходящих через турникет
type TurnstileCounter struct {
	mu    sync.RWMutex
	count int
}

// Turnstile представляет турникет с счетчиком
type Turnstile struct {
	counter *TurnstileCounter
}

// NewTurnstile создает новый турникет с счетчиком
func NewTurnstile() *Turnstile {
	return &Turnstile{
		counter: &TurnstileCounter{},
	}
}

// Pass проход человека через турникет
func (t *Turnstile) Pass() {
	t.counter.mu.Lock()
	defer t.counter.mu.Unlock()
	t.counter.count++
}

// Count возвращает текущее количество людей, прошедших через турникет
func (t *Turnstile) Count() int {
	t.counter.mu.RLock()
	defer t.counter.mu.RUnlock()
	return t.counter.count
}

func main() {
	// Демонстрация работы с несколькими турникетами и использованием RWMutex

	turnstile1 := NewTurnstile()
	turnstile2 := NewTurnstile()

	// Используем горутины для параллельного прохождения людей через турникеты
	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				turnstile1.Pass()
				turnstile2.Pass()
			}
		}()
	}

	// Ждем завершения всех горутин
	time.Sleep(time.Second)

	fmt.Printf("With synchronization using RWMutex:\n")
	fmt.Printf("Turnstile 1: Total people passed: %d\n", turnstile1.Count())
	fmt.Printf("Turnstile 2: Total people passed: %d\n", turnstile2.Count())
	// Демонстрация работы без использования примитива синхронизации

	// Создаем новые счетчики для турникетов без синхронизации
	nonSyncTurnstile1 := NewTurnstile()
	nonSyncTurnstile2 := NewTurnstile()

	// Используем горутины для параллельного прохождения людей через турникеты без синхронизации
	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				nonSyncTurnstile1.counter.count++
				nonSyncTurnstile2.counter.count++
			}
		}()
	}

	// Ждем завершения всех горутин
	time.Sleep(time.Second)

	// Результат без синхронизации может быть непредсказуемым
	fmt.Printf("\nWithout synchronization:\n")
	fmt.Printf("Turnstile 1: Total people passed: %d\n", nonSyncTurnstile1.counter.count)
	fmt.Printf("Turnstile 2: Total people passed: %d\n", nonSyncTurnstile2.counter.count)
}
