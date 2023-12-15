package services

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"project/source/domain/entity"
)

func (s Service) SelectEventBySubId(subs []int) ([]string, error) {
	return s.Repository.SelectEventBySubIdsRepo(subs)
}
func (s Service) SelectEventByEventIds(eventIds []string) ([]entity.Event, error) {
	return s.Repository.SelectEventsByIdRepo(eventIds)
}
func (s Service) SelectEventsByUserId(userId int) ([]entity.Event, error) {
	ctx := context.Background()

	// Попробуем получить данные из кеша Redis
	cachedData, err := s.RedisClient.Get(ctx, fmt.Sprintf("user:%d:events", userId)).Result()
	if err == nil {
		// Если данные есть в кеше, преобразуем их и вернем
		var events []entity.Event
		fmt.Println("Cached")

		if err := json.Unmarshal([]byte(cachedData), &events); err == nil {
			fmt.Println(events);
			return events, nil
		}
	}
	fmt.Println("Not cached")
	// Если данные отсутствуют в кеше, получаем их из репозитория
	events, err := s.Repository.SelectEventsByUserIdRepo(userId)
	if err != nil {
		return nil, err
	}

	// Сохраняем данные в кеше на будущее
	eventData, err := json.Marshal(events)
	if err != nil {
		// Обрабатываем ошибку маршалинга (например, записываем в лог)
	} else {
		//err = s.RedisClient.Set(ctx, fmt.Sprintf("user:%d:events", userId), eventData, 5*time.Minute).Err()
		fmt.Println(eventData)
		err = s.RedisClient.Set(ctx, fmt.Sprintf("user:%d:events", userId), eventData, 0).Err()
		if err != nil {
			// Обрабатываем ошибку сохранения в кеше (например, записываем в лог)
		}
	}

	return events, nil
}
