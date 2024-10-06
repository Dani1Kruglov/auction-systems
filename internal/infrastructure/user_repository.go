package infrastructure

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"auction-systems/internal/domain"

	_ "github.com/lib/pq"
)

type UserRepository struct {
	db     *sql.DB
	buyers map[int]*domain.User
}

func NewUserRepository(db *sql.DB) (*UserRepository, error) {
	userRepo := &UserRepository{db: db}

	var err error
	userRepo.buyers, err = userRepo.getBuyers()
	if err != nil {
		return nil, err
	}

	go func() {
		ticker := time.NewTicker(time.Minute)
		for range ticker.C {
			buyers, err := userRepo.getBuyers()
			if err != nil {
				continue
			}
			userRepo.buyers = buyers
		}
	}()

	return userRepo, nil
}

func (ur *UserRepository) getBuyers() (map[int]*domain.User, error) {
	rows, err := ur.db.Query(`SELECT id, name, email, password, notification_url FROM users WHERE role = 'buyer'`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var buyers = make(map[int]*domain.User)

	for rows.Next() {
		var buyer domain.User
		err := rows.Scan(&buyer.Id, &buyer.Name, &buyer.Email, &buyer.Password, &buyer.NotificationUrl)
		if err != nil {
			return nil, err
		}
		buyers[buyer.Id] = &buyer
	}
	return buyers, nil
}

func (ur *UserRepository) GetBuyers() map[int]*domain.User {
	return ur.buyers
}

func (ur *UserRepository) Create(user *domain.User) (int, error) {
	query := `
        INSERT INTO users (name, role, email, password, notification_url) 
        VALUES ($1, $2, $3, $4, $5) 
        RETURNING id`

	err := ur.db.QueryRow(query, user.Name, user.Role, user.Email, user.Password, user.NotificationUrl).Scan(&user.Id)
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (ur *UserRepository) GetUserById(userId int, role string) (*domain.User, error) {
	var user domain.User
	query := `SELECT id, name, email, password FROM users WHERE id = $1 AND role = $2`

	err := ur.db.QueryRow(query, userId, role).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	query := `SELECT id, name, email, password FROM users WHERE email = $1`

	err := ur.db.QueryRow(query, email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) AddUserBalance(userID int, amount float64) error {
	_, err := ur.db.Exec(`INSERT INTO user_balances (user_id, balance) VALUES ($1, $2)`, userID, amount)
	if err != nil {
		return fmt.Errorf("error inserting balance: %w", err)
	}
	return nil
}

func (ur *UserRepository) UpdateUserBalance(userID int, amount float64, symbol string) (bool, error) {
	res, err := ur.db.Exec(fmt.Sprintf(`UPDATE user_balances SET balance = balance %s $1 WHERE user_id = $2`, symbol), amount, userID)
	if err != nil {
		return false, fmt.Errorf("error updating balance: %w", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("error updating balance: %w", err)
	}
	if rowsAffected != 0 {
		return true, nil
	}
	return false, nil
}

func (ur *UserRepository) GetBalanceByUserId(userID int) (float64, error) {
	var balance float64

	err := ur.db.QueryRow(`SELECT balance FROM public.user_balances WHERE user_id = $1`, userID).Scan(&balance)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, fmt.Errorf("error getting balance for user ID %d: %w", userID, err)
	}

	return balance, nil
}
