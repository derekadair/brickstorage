package part

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
==> colors.csv <==
id,name,rgb,is_transparent
-1,[Unknown],0033B2,False

==> elements.csv <==
element_id,part_num,color_id,design_id
4153515,40905,10,

==> inventories.csv <==
id,version,set_num
1,1,7922-1

==> inventory_minifigs.csv <==
inventory_id,fig_num,quantity
3,fig-001549,1

==> inventory_parts.csv <==
inventory_id,part_num,color_id,quantity,is_spare,img_url
1,48379c04,72,1,False,https://cdn.rebrickable.com/media/parts/photos/1/48379c01-1-839cbcec-62de-4733-ba23-20f35f4dd5d5.jpg

==> inventory_sets.csv <==
inventory_id,set_num,quantity
35,75911-1,1

==> minifigs.csv <==
fig_num,name,num_parts,img_url
fig-000001,Toy Store Employee,4,https://cdn.rebrickable.com/media/sets/fig-000001.jpg

==> part_categories.csv <==
id,name
1,Baseplates

==> part_relationships.csv <==
rel_type,child_part_num,parent_part_num
P,3626cpr3662,3626c

==> parts.csv <==
part_num,name,part_cat_id,part_material
003381,Sticker Sheet for Set 663-1,58,Plastic

==> sets.csv <==
set_num,name,year,theme_id,num_parts,img_url
0003977811-1,Ninjago: Book of Adventures,2022,761,1,https://cdn.rebrickable.com/media/sets/0003977811-1.jpg

==> themes.csv <==
id,name,parent_id
1,Technic,

==> brick architect_part.csv <==
ba_part_id,ba_part_subcategory_id,ba_part_description,part_start_year,part_end_year,ba_has_label

==> brick architect to pab.csv <==
pab_part_id,ba_part_id,pab_part_description

==> brick architect to bricklink.csv <==
bricklink_part_id,ba_part_id,bricklink_part_description

==> brick architect to rebrickable.csv <==
rebrickable_part_id,ba_part_id,rebrickable_part_description

==> brick architect to brickset.csv <==
brickset_part_id,ba_part_id,brickset_part_description

==> brick architect to ldraw.csv <==
ldraw_part_id,ba_part_id,ldraw_part_description

==> brick architect categories <==
category_id,ba_category_id,ba_category_name,ba_category_parent_id,ba_category_description
*/
type DTO struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date"`
	ImageURL      string `json:"image_url"`
	Description   string `json:"description"`
}

type Form struct {
	Title         string `json:"title" form:"required,max=255"`
	Author        string `json:"author" form:"required,alpha_space,max=255"`
	PublishedDate string `json:"published_date" form:"required,datetime=2006-01-02"`
	ImageURL      string `json:"image_url" form:"url"`
	Description   string `json:"description"`
}

type Book struct {
	ID            uuid.UUID `gorm:"primarykey"`
	Title         string
	Author        string
	PublishedDate time.Time
	ImageURL      string
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type Books []*Book

func (b *Book) ToDto() *DTO {
	return &DTO{
		ID:            b.ID.String(),
		Title:         b.Title,
		Author:        b.Author,
		PublishedDate: b.PublishedDate.Format("2006-01-02"),
		ImageURL:      b.ImageURL,
		Description:   b.Description,
	}
}

func (bs Books) ToDto() []*DTO {
	dtos := make([]*DTO, len(bs))
	for i, v := range bs {
		dtos[i] = v.ToDto()
	}

	return dtos
}

func (f *Form) ToModel() *Book {
	pubDate, _ := time.Parse("2006-01-02", f.PublishedDate)

	return &Book{
		Title:         f.Title,
		Author:        f.Author,
		PublishedDate: pubDate,
		ImageURL:      f.ImageURL,
		Description:   f.Description,
	}
}
