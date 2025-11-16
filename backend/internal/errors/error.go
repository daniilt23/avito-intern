package errors

import "errors"

var ErrTeamExists = errors.New("team already exist")
var ErrTeamNotFound = errors.New("team not found")
var ErrUserNotFound = errors.New("user not found")
var ErrPrNotFound = errors.New("pull request not found")
var ErrInternalServer = errors.New("internal server error")
var ErrPrExists = errors.New("pull request already exist")
var ErrPrMerged = errors.New("cannot reassign on merged PR")
var ErrNotAssigned = errors.New("reviewer is not assigned to this PR")
var ErrNoCandidate = errors.New("no active replacement candidate in team")
var ErrNoTeamQuery = errors.New("query parameter need to be team_name")
var ErrNoUserQuery = errors.New("query parameter need to be user_id")
var ErrMissingHeader = errors.New("missing authorization header")
var ErrInvalidFormat = errors.New("invalid authorization header format")
var ErrUnauthorized = errors.New("invalid token")
