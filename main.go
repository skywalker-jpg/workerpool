package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

// Pool структура для управления воркерами
type Pool struct {
	JobQueue     chan string
	Workers      map[int]*Worker
	SyncGroup    sync.WaitGroup
	Logger       *log.Logger
	nextWorkerID int
}

// NewPool создает новый pool
func NewPool(logger *log.Logger) *Pool {
	return &Pool{
		JobQueue: make(chan string),
		Workers:  make(map[int]*Worker),
		Logger:   logger,
	}
}

// Worker структура
type Worker struct {
	ID       int
	JobQueue chan string
	Quit     chan struct{}
}

// NewWorker создает новый воркер
func NewWorker(id int, jobQueue chan string) *Worker {
	return &Worker{
		ID:       id,
		JobQueue: jobQueue,
		Quit:     make(chan struct{}),
	}
}

// Start запускает воркера для обработки данных
func (w *Worker) Start(wg *sync.WaitGroup, logger *log.Logger) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-w.JobQueue:
			if !ok {
				logger.Printf("Воркер %d: канал закрыт, завершаю работу\n", w.ID)
				return
			}
			logger.Printf("Воркер %d: выполняет задание: %s\n", w.ID, job)
			time.Sleep(time.Duration(rand.Intn(2)+1) * time.Second)       // имитация выполнения
			logger.Printf("Воркер %d: завершил задание: %s\n", w.ID, job) // Лог завершения работы над заданием
		case <-w.Quit:
			logger.Printf("Воркер %d: завершает работу\n", w.ID)
			return
		}
	}
}

// Stop останавливает воркера
func (w *Worker) Stop() {
	close(w.Quit)
}

// AddWorker добавляет нового воркера в pool
func (p *Pool) AddWorker() {
	newID := len(p.Workers) + 1
	worker := NewWorker(newID, p.JobQueue)
	p.Workers[p.nextWorkerID] = worker
	p.nextWorkerID++
	p.SyncGroup.Add(1)
	go worker.Start(&p.SyncGroup, p.Logger)
	fmt.Printf("Воркер %d запущен\n", newID)
}

// RemoveWorker удаляет воркера из pool по ID
func (p *Pool) RemoveWorker(id int) {
	log.Println(p.Workers, id)
	worker, exists := p.Workers[id-1]
	if !exists {
		fmt.Printf("Воркер с ID %d не найден\n", id)
		return
	}
	worker.Stop()
	delete(p.Workers, id)
	fmt.Printf("Воркер %d удален\n", id)
}

// Stop завершает работу всех воркеров и закрывает pool
func (p *Pool) Stop() {
	for id, worker := range p.Workers {
		worker.Stop()
		delete(p.Workers, id)
		p.Logger.Printf("Воркер %d: отправлен сигнал на остановку\n", id)
	}
	p.SyncGroup.Wait()
	close(p.JobQueue)
}

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	pool := NewPool(logger)

	// Канал для завершения программы
	done := make(chan struct{})

	// Запускаем обработку команд пользователя
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Команды:")
		fmt.Println("add    - добавить воркер")
		fmt.Println("remove - удалить воркер по ID")
		fmt.Println("exit   - выйти из программы")

		for scanner.Scan() {
			command := scanner.Text()
			switch command {
			case "add":
				pool.AddWorker()
			case "remove":
				fmt.Print("Введите ID воркера для удаления: ")
				if scanner.Scan() {
					idStr := scanner.Text()
					id, err := strconv.Atoi(idStr)
					if err == nil {
						pool.RemoveWorker(id)
					} else {
						fmt.Println("Некорректный ID воркера")
					}
				}
			case "exit":
				fmt.Println("Завершение работы...")
				pool.Stop()
				close(done)
				return
			default:
				fmt.Println("Неизвестная команда")
			}
		}
	}()

	// Имитируем поток заданий, отправляемых в канал JobQueue
	go func() {
		jobNumber := 1
		for {
			select {
			case <-done:
				return
			default:
				job := fmt.Sprintf("Задание %d", jobNumber)
				pool.JobQueue <- job
				jobNumber++
				time.Sleep(time.Second) // Пауза между заданиями
			}
		}
	}()

	<-done
}
