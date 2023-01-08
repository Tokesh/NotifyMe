package repositories

import (
	"context"
)

func (r *Repository) SelectUserSubscriptionRepo(userId int) ([]int, error) {
	// username, user_email, user_password, user_activation_status,status
	q := `
		SELECT subs_id from user_subscription where user_id = $1
	`
	subs := make([]int, 0)
	rows, err := r.client.Query(context.TODO(), q, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var sub int
		err = rows.Scan(&sub)
		if err != nil {
			return subs, err
		}
		subs = append(subs, sub)
	}
	return subs, nil
}
