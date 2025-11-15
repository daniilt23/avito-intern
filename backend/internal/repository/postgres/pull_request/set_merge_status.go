package pullrequest

import (
	"time"
)

func (r *PullRequestRepoSQL) SetMergeStatus(prId string, mergedAt time.Time) error {
	query := `
	UPDATE pull_requests
	SET status = 'MERGED', merged_at = $1
	WHERE pull_request_id = $2`

	rows, err := r.db.Exec(query, mergedAt, prId)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return err
	}

	return nil
}
