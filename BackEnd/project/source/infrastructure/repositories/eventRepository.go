package repositories

import (
	"context"
	"fmt"
	"github.com/jackc/pgtype"
	"project/source/domain/entity"
)

func (r *Repository) SelectEventBySubIdsRepo(subs []int) ([]string, error) {
	// username, user_email, user_password, user_activation_status,status
	events := make([]string, 0)
	for i := 0; i < len(subs); i++ {
		q := `
			SELECT event_id from event_sub where subs_id = $1
		`
		rows, err := r.client.Query(context.TODO(), q, subs[i])
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var event pgtype.UUID
			err = rows.Scan(&event)
			//fmt.Println(event)
			if err != nil {
				return nil, err
			}
			eventString := fmt.Sprintf("%x-%x-%x-%x-%x", event.Bytes[0:4], event.Bytes[4:6], event.Bytes[6:8], event.Bytes[8:10], event.Bytes[10:16])
			//fmt.Println(eventString)
			events = append(events, eventString)
		}

	}
	//fmt.Println(events)
	return events, nil
}

func (r *Repository) SelectEventsByIdRepo(eventIds []string) ([]entity.Event, error) {
	events := make([]entity.Event, 0)
	var event entity.Event
	for i := 0; i < len(eventIds); i++ {
		//events_id, event_name, event_timestart, event_timeend, event_result
		q := `
			SELECT events_id, event_name, event_timestart, event_timeend, event_result
			FROM events WHERE events_id = $1
		`

		err := r.client.QueryRow(context.TODO(), q, eventIds[i]).Scan(&event.Id, &event.Name,
			&event.TimeStart, &event.TimeEnd, &event.Result)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (r *Repository) SelectEventsByUserIdRepo(userId int) ([]entity.Event, error) {
	q := `
			select event_id, event_name, event_timestart, event_timeend, event_result from user_subscription
				left join subscriptions on user_subscription.subs_id = subscriptions.subs_id
				left join event_sub on subscriptions.subs_id = event_sub.subs_id
				left join events e on event_sub.event_id = e.events_id
				where user_id = $1
		`
	events := make([]entity.Event, 0)
	rows, err := r.client.Query(context.TODO(), q, userId)
	for rows.Next() {
		var event entity.Event
		err = rows.Scan(&event.Id, &event.Name, &event.TimeStart, &event.TimeEnd, &event.Result)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
