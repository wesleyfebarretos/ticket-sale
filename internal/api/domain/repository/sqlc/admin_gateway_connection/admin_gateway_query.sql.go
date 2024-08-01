// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: admin_gateway_query.sql

package admin_gateway_connection

import (
	"context"
)

const create = `-- name: Create :one
INSERT INTO fin.gateway
    ("name", description, client_id, client_secret, "order", active, test_environment, notif_user, notif_password, soft_descriptor, gateway_process_id, webhook_url, url, auth_type, use_3ds, adq_code_3ds, default_adq_code, use_antifraud, created_by, updated_by)
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
RETURNING id, uuid, name, description, client_id, client_secret, "order", active, is_deleted, test_environment, notif_user, notif_password, soft_descriptor, gateway_process_id, webhook_url, url, auth_type, use_3ds, adq_code_3ds, default_adq_code, use_antifraud, created_by, updated_by, created_at, updated_at
`

type CreateParams struct {
	Name             string          `json:"name"`
	Description      *string         `json:"description"`
	ClientID         *string         `json:"clientId"`
	ClientSecret     *string         `json:"clientSecret"`
	Order            int32           `json:"order"`
	Active           bool            `json:"active"`
	TestEnvironment  bool            `json:"testEnvironment"`
	NotifUser        *string         `json:"notifUser"`
	NotifPassword    *string         `json:"notifPassword"`
	SoftDescriptor   *string         `json:"softDescriptor"`
	GatewayProcessID int32           `json:"gatewayProcessId"`
	WebhookUrl       *string         `json:"webhookUrl"`
	Url              *string         `json:"url"`
	AuthType         GatewayAuthType `json:"authType"`
	Use3ds           bool            `json:"use3ds"`
	AdqCode3ds       *string         `json:"adqCode3ds"`
	DefaultAdqCode   *string         `json:"defaultAdqCode"`
	UseAntifraud     bool            `json:"useAntifraud"`
	CreatedBy        int32           `json:"createdBy"`
	UpdatedBy        *int32          `json:"updatedBy"`
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) (FinGateway, error) {
	row := q.db.QueryRow(ctx, create,
		arg.Name,
		arg.Description,
		arg.ClientID,
		arg.ClientSecret,
		arg.Order,
		arg.Active,
		arg.TestEnvironment,
		arg.NotifUser,
		arg.NotifPassword,
		arg.SoftDescriptor,
		arg.GatewayProcessID,
		arg.WebhookUrl,
		arg.Url,
		arg.AuthType,
		arg.Use3ds,
		arg.AdqCode3ds,
		arg.DefaultAdqCode,
		arg.UseAntifraud,
		arg.CreatedBy,
		arg.UpdatedBy,
	)
	var i FinGateway
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Name,
		&i.Description,
		&i.ClientID,
		&i.ClientSecret,
		&i.Order,
		&i.Active,
		&i.IsDeleted,
		&i.TestEnvironment,
		&i.NotifUser,
		&i.NotifPassword,
		&i.SoftDescriptor,
		&i.GatewayProcessID,
		&i.WebhookUrl,
		&i.Url,
		&i.AuthType,
		&i.Use3ds,
		&i.AdqCode3ds,
		&i.DefaultAdqCode,
		&i.UseAntifraud,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAll = `-- name: GetAll :many
SELECT id, uuid, name, description, client_id, client_secret, "order", active, test_environment, notif_user, notif_password, soft_descriptor, gateway_process_id, webhook_url, url, auth_type, use_3ds, adq_code_3ds, default_adq_code, use_antifraud, created_by, updated_by, created_at, updated_at, "gatewayProcess", "gatewayPaymentTypes" FROM gateway_details
`

func (q *Queries) GetAll(ctx context.Context) ([]GatewayDetail, error) {
	rows, err := q.db.Query(ctx, getAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GatewayDetail{}
	for rows.Next() {
		var i GatewayDetail
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.Name,
			&i.Description,
			&i.ClientID,
			&i.ClientSecret,
			&i.Order,
			&i.Active,
			&i.TestEnvironment,
			&i.NotifUser,
			&i.NotifPassword,
			&i.SoftDescriptor,
			&i.GatewayProcessID,
			&i.WebhookUrl,
			&i.Url,
			&i.AuthType,
			&i.Use3ds,
			&i.AdqCode3ds,
			&i.DefaultAdqCode,
			&i.UseAntifraud,
			&i.CreatedBy,
			&i.UpdatedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.GatewayProcess,
			&i.GatewayPaymentTypes,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOneById = `-- name: GetOneById :one
SELECT 
    id, uuid, name, description, client_id, client_secret, "order", active, test_environment, notif_user, notif_password, soft_descriptor, gateway_process_id, webhook_url, url, auth_type, use_3ds, adq_code_3ds, default_adq_code, use_antifraud, created_by, updated_by, created_at, updated_at, "gatewayProcess", "gatewayPaymentTypes"
FROM 
    gateway_details
WHERE id = $1
`

func (q *Queries) GetOneById(ctx context.Context, id int32) (GatewayDetail, error) {
	row := q.db.QueryRow(ctx, getOneById, id)
	var i GatewayDetail
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Name,
		&i.Description,
		&i.ClientID,
		&i.ClientSecret,
		&i.Order,
		&i.Active,
		&i.TestEnvironment,
		&i.NotifUser,
		&i.NotifPassword,
		&i.SoftDescriptor,
		&i.GatewayProcessID,
		&i.WebhookUrl,
		&i.Url,
		&i.AuthType,
		&i.Use3ds,
		&i.AdqCode3ds,
		&i.DefaultAdqCode,
		&i.UseAntifraud,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.GatewayProcess,
		&i.GatewayPaymentTypes,
	)
	return i, err
}

const softDelete = `-- name: SoftDelete :exec
UPDATE fin.gateway SET 
    active = false,
    is_deleted = true,
    updated_by = $2
WHERE id = $1
`

type SoftDeleteParams struct {
	ID        int32  `json:"id"`
	UpdatedBy *int32 `json:"updatedBy"`
}

func (q *Queries) SoftDelete(ctx context.Context, arg SoftDeleteParams) error {
	_, err := q.db.Exec(ctx, softDelete, arg.ID, arg.UpdatedBy)
	return err
}

const update = `-- name: Update :exec
UPDATE fin.gateway SET
    "name" = $2,
    description = $3,
    client_id = $4,
    client_secret = $5,
    "order" = $6,
    active = $7,
    test_environment = $8,
    notif_user = $9,
    notif_password = $10,
    soft_descriptor = $11,
    gateway_process_id = $12,
    webhook_url = $13,
    url = $14,
    auth_type = $15,
    use_3ds = $16,
    adq_code_3ds = $17,
    default_adq_code = $18,
    use_antifraud = $19,
    updated_by = $20
WHERE id = $1
`

type UpdateParams struct {
	ID               int32           `json:"id"`
	Name             string          `json:"name"`
	Description      *string         `json:"description"`
	ClientID         *string         `json:"clientId"`
	ClientSecret     *string         `json:"clientSecret"`
	Order            int32           `json:"order"`
	Active           bool            `json:"active"`
	TestEnvironment  bool            `json:"testEnvironment"`
	NotifUser        *string         `json:"notifUser"`
	NotifPassword    *string         `json:"notifPassword"`
	SoftDescriptor   *string         `json:"softDescriptor"`
	GatewayProcessID int32           `json:"gatewayProcessId"`
	WebhookUrl       *string         `json:"webhookUrl"`
	Url              *string         `json:"url"`
	AuthType         GatewayAuthType `json:"authType"`
	Use3ds           bool            `json:"use3ds"`
	AdqCode3ds       *string         `json:"adqCode3ds"`
	DefaultAdqCode   *string         `json:"defaultAdqCode"`
	UseAntifraud     bool            `json:"useAntifraud"`
	UpdatedBy        *int32          `json:"updatedBy"`
}

func (q *Queries) Update(ctx context.Context, arg UpdateParams) error {
	_, err := q.db.Exec(ctx, update,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.ClientID,
		arg.ClientSecret,
		arg.Order,
		arg.Active,
		arg.TestEnvironment,
		arg.NotifUser,
		arg.NotifPassword,
		arg.SoftDescriptor,
		arg.GatewayProcessID,
		arg.WebhookUrl,
		arg.Url,
		arg.AuthType,
		arg.Use3ds,
		arg.AdqCode3ds,
		arg.DefaultAdqCode,
		arg.UseAntifraud,
		arg.UpdatedBy,
	)
	return err
}
