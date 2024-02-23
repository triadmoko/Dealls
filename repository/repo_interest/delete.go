package repo_interest

import (
	"app/model"
	"log"
	"time"
)

func (r *Repository) DeletePartnerExpired() error {
	log.Println("Cron Start: DeletePartnerExpired")
	var (
		interest model.Interest
		where    string = "deleted_at IS NULL "
	)
	yesterday := time.Now().UTC().AddDate(0, 0, -1).Format("2006-01-02 15:04:05")
	where += " AND created_at < ?"
	err := r.db.Debug().Model(&model.Interest{}).Where(where, yesterday).Delete(&interest).Error
	if err != nil {
		return err
	}
	log.Println("End Cron: DeletePartnerExpired")
	return nil
}
