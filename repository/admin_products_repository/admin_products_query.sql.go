// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: admin_products_query.sql

package admin_products_repository

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const create = `-- name: Create :one
INSERT INTO products 
(name, description, uuid, price, discount_price, active, image, image_mobile, image_thumbnail, category_id, created_by)
VALUES 
($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) 
RETURNING id, name, description, uuid, price, discount_price, active, is_deleted, image, image_mobile, image_thumbnail, category_id, created_by, updated_by, created_at, updated_at
`

type CreateParams struct {
	Name           string    `json:"name"`
	Description    *string   `json:"description"`
	Uuid           uuid.UUID `json:"uuid"`
	Price          float64   `json:"price"`
	DiscountPrice  *float64  `json:"discountPrice"`
	Active         bool      `json:"active"`
	Image          *string   `json:"image"`
	ImageMobile    *string   `json:"imageMobile"`
	ImageThumbnail *string   `json:"imageThumbnail"`
	CategoryID     int32     `json:"categoryId"`
	CreatedBy      int32     `json:"createdBy"`
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) (Product, error) {
	row := q.db.QueryRow(ctx, create,
		arg.Name,
		arg.Description,
		arg.Uuid,
		arg.Price,
		arg.DiscountPrice,
		arg.Active,
		arg.Image,
		arg.ImageMobile,
		arg.ImageThumbnail,
		arg.CategoryID,
		arg.CreatedBy,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Uuid,
		&i.Price,
		&i.DiscountPrice,
		&i.Active,
		&i.IsDeleted,
		&i.Image,
		&i.ImageMobile,
		&i.ImageThumbnail,
		&i.CategoryID,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAll = `-- name: GetAll :many
SELECT id, name, description, uuid, price, discount_price, active, is_deleted, image, image_mobile, image_thumbnail, category_id, created_by, updated_by, created_at, updated_at FROM products 
WHERE 
    is_deleted IS FALSE 
AND
    active IS TRUE
ORDER BY 
    created_at DESC
`

func (q *Queries) GetAll(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, getAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Uuid,
			&i.Price,
			&i.DiscountPrice,
			&i.Active,
			&i.IsDeleted,
			&i.Image,
			&i.ImageMobile,
			&i.ImageThumbnail,
			&i.CategoryID,
			&i.CreatedBy,
			&i.UpdatedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getAllWithRelations = `-- name: GetAllWithRelations :many
SELECT 
    p.id, p.name, p.description, p.uuid, p.price, p.discount_price, p.active, p.is_deleted, p.image, p.image_mobile, p.image_thumbnail, p.category_id, p.created_by, p.updated_by, p.created_at, p.updated_at,
    CASE
        WHEN ps.id IS NULL THEN NULL
        ELSE
            json_build_object(
                'id', ps.id,
                'productId', ps.product_id,
                'qty', ps.qty,
                'minQty', ps.min_qty
            )
    END as stock,
    CASE
        WHEN pc.id IS NULL THEN NULL
        ELSE
            json_build_object(
                'id', pc.id,
                'name', pc.name,
                'description', pc.description
            )
    END as category
FROM 
    products as p
LEFT JOIN 
    product_stocks as ps 
ON 
    ps.product_id = p.id
LEFT JOIN
    product_categories as pc
ON
    pc.id = p.category_id
WHERE 
    p.is_deleted IS FALSE 
AND
    p.active IS TRUE
ORDER BY 
    p.created_at DESC
`

type GetAllWithRelationsRow struct {
	ID             int32       `json:"id"`
	Name           string      `json:"name"`
	Description    *string     `json:"description"`
	Uuid           uuid.UUID   `json:"uuid"`
	Price          float64     `json:"price"`
	DiscountPrice  *float64    `json:"discountPrice"`
	Active         bool        `json:"active"`
	IsDeleted      bool        `json:"isDeleted"`
	Image          *string     `json:"image"`
	ImageMobile    *string     `json:"imageMobile"`
	ImageThumbnail *string     `json:"imageThumbnail"`
	CategoryID     int32       `json:"categoryId"`
	CreatedBy      int32       `json:"createdBy"`
	UpdatedBy      *int32      `json:"updatedBy"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      *time.Time  `json:"updatedAt"`
	Stock          interface{} `json:"stock"`
	Category       interface{} `json:"category"`
}

func (q *Queries) GetAllWithRelations(ctx context.Context) ([]GetAllWithRelationsRow, error) {
	rows, err := q.db.Query(ctx, getAllWithRelations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllWithRelationsRow{}
	for rows.Next() {
		var i GetAllWithRelationsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Uuid,
			&i.Price,
			&i.DiscountPrice,
			&i.Active,
			&i.IsDeleted,
			&i.Image,
			&i.ImageMobile,
			&i.ImageThumbnail,
			&i.CategoryID,
			&i.CreatedBy,
			&i.UpdatedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Stock,
			&i.Category,
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
    p.id, p.name, p.description, p.uuid, p.price, p.discount_price, p.active, p.is_deleted, p.image, p.image_mobile, p.image_thumbnail, p.category_id, p.created_by, p.updated_by, p.created_at, p.updated_at,
    CASE
        WHEN ps.id IS NULL THEN NULL
        ELSE
            json_build_object(
                'id', ps.id,
                'productId', ps.product_id,
                'qty', ps.qty,
                'minQty', ps.min_qty
            )
    END as stock,
    CASE
        WHEN pc.id IS NULL THEN NULL
        ELSE
            json_build_object(
                'id', pc.id,
                'name', pc.name,
                'description', pc.description
            )
    END as category
FROM 
    products as p
LEFT JOIN 
    product_stocks as ps 
ON 
    ps.product_id = p.id
LEFT JOIN
    product_categories as pc
ON
    pc.id = p.category_id
WHERE 
    p.id = $1
LIMIT 1
`

type GetOneByIdRow struct {
	ID             int32       `json:"id"`
	Name           string      `json:"name"`
	Description    *string     `json:"description"`
	Uuid           uuid.UUID   `json:"uuid"`
	Price          float64     `json:"price"`
	DiscountPrice  *float64    `json:"discountPrice"`
	Active         bool        `json:"active"`
	IsDeleted      bool        `json:"isDeleted"`
	Image          *string     `json:"image"`
	ImageMobile    *string     `json:"imageMobile"`
	ImageThumbnail *string     `json:"imageThumbnail"`
	CategoryID     int32       `json:"categoryId"`
	CreatedBy      int32       `json:"createdBy"`
	UpdatedBy      *int32      `json:"updatedBy"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      *time.Time  `json:"updatedAt"`
	Stock          interface{} `json:"stock"`
	Category       interface{} `json:"category"`
}

func (q *Queries) GetOneById(ctx context.Context, id int32) (GetOneByIdRow, error) {
	row := q.db.QueryRow(ctx, getOneById, id)
	var i GetOneByIdRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Uuid,
		&i.Price,
		&i.DiscountPrice,
		&i.Active,
		&i.IsDeleted,
		&i.Image,
		&i.ImageMobile,
		&i.ImageThumbnail,
		&i.CategoryID,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Stock,
		&i.Category,
	)
	return i, err
}

const getOneByUuid = `-- name: GetOneByUuid :one
SELECT 
    p.id, p.name, p.description, p.uuid, p.price, p.discount_price, p.active, p.is_deleted, p.image, p.image_mobile, p.image_thumbnail, p.category_id, p.created_by, p.updated_by, p.created_at, p.updated_at,
    CASE
        WHEN ps.id IS NULL THEN NULL
        ELSE
            json_build_object(
                'id', ps.id,
                'productId', ps.product_id,
                'qty', ps.qty,
                'minQty', ps.min_qty
            )
    END as stock,
    CASE
        WHEN pc.id IS NULL THEN NULL
        ELSE
            json_build_object(
                'id', pc.id,
                'name', pc.name,
                'description', pc.description
            )
    END as category
FROM 
    products as p
LEFT JOIN 
    product_stocks as ps 
ON 
    ps.product_id = p.id
LEFT JOIN
    product_categories as pc
ON
    pc.id = p.category_id
WHERE 
    p.uuid = $1
LIMIT 1
`

type GetOneByUuidRow struct {
	ID             int32       `json:"id"`
	Name           string      `json:"name"`
	Description    *string     `json:"description"`
	Uuid           uuid.UUID   `json:"uuid"`
	Price          float64     `json:"price"`
	DiscountPrice  *float64    `json:"discountPrice"`
	Active         bool        `json:"active"`
	IsDeleted      bool        `json:"isDeleted"`
	Image          *string     `json:"image"`
	ImageMobile    *string     `json:"imageMobile"`
	ImageThumbnail *string     `json:"imageThumbnail"`
	CategoryID     int32       `json:"categoryId"`
	CreatedBy      int32       `json:"createdBy"`
	UpdatedBy      *int32      `json:"updatedBy"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      *time.Time  `json:"updatedAt"`
	Stock          interface{} `json:"stock"`
	Category       interface{} `json:"category"`
}

func (q *Queries) GetOneByUuid(ctx context.Context, argUuid uuid.UUID) (GetOneByUuidRow, error) {
	row := q.db.QueryRow(ctx, getOneByUuid, argUuid)
	var i GetOneByUuidRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Uuid,
		&i.Price,
		&i.DiscountPrice,
		&i.Active,
		&i.IsDeleted,
		&i.Image,
		&i.ImageMobile,
		&i.ImageThumbnail,
		&i.CategoryID,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Stock,
		&i.Category,
	)
	return i, err
}

const softDelete = `-- name: SoftDelete :exec
UPDATE products SET
    is_deleted = true,
    updated_by = $2
WHERE 
    id = $1
AND 
    is_deleted IS FALSE
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
UPDATE products SET
    name = $1,
    description = $2,
    price = $3,
    discount_price = $4,
    active = $5,
    image = $6,
    image_mobile = $7,
    image_thumbnail = $8,
    category_id = $9,
    updated_by = $10
WHERE 
    id = $11
`

type UpdateParams struct {
	Name           string   `json:"name"`
	Description    *string  `json:"description"`
	Price          float64  `json:"price"`
	DiscountPrice  *float64 `json:"discountPrice"`
	Active         bool     `json:"active"`
	Image          *string  `json:"image"`
	ImageMobile    *string  `json:"imageMobile"`
	ImageThumbnail *string  `json:"imageThumbnail"`
	CategoryID     int32    `json:"categoryId"`
	UpdatedBy      *int32   `json:"updatedBy"`
	ID             int32    `json:"id"`
}

func (q *Queries) Update(ctx context.Context, arg UpdateParams) error {
	_, err := q.db.Exec(ctx, update,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.DiscountPrice,
		arg.Active,
		arg.Image,
		arg.ImageMobile,
		arg.ImageThumbnail,
		arg.CategoryID,
		arg.UpdatedBy,
		arg.ID,
	)
	return err
}
