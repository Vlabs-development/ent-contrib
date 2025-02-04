// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BillProductsColumns holds the columns for the "bill_products" table.
	BillProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "sku", Type: field.TypeString},
		{Name: "quantity", Type: field.TypeUint64},
	}
	// BillProductsTable holds the schema information for the "bill_products" table.
	BillProductsTable = &schema.Table{
		Name:       "bill_products",
		Columns:    BillProductsColumns,
		PrimaryKey: []*schema.Column{BillProductsColumns[0]},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"ENABLED", "DISABLED"}},
		{Name: "config", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"sqlite3": "json"}},
		{Name: "types", Type: field.TypeJSON, Nullable: true},
		{Name: "duration", Type: field.TypeInt64, Nullable: true},
		{Name: "count", Type: field.TypeUint64, Nullable: true},
		{Name: "strings", Type: field.TypeJSON, Nullable: true},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// FriendshipsColumns holds the columns for the "friendships" table.
	FriendshipsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "friend_id", Type: field.TypeUUID},
	}
	// FriendshipsTable holds the schema information for the "friendships" table.
	FriendshipsTable = &schema.Table{
		Name:       "friendships",
		Columns:    FriendshipsColumns,
		PrimaryKey: []*schema.Column{FriendshipsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "friendships_users_user",
				Columns:    []*schema.Column{FriendshipsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "friendships_users_friend",
				Columns:    []*schema.Column{FriendshipsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "friendship_user_id_friend_id",
				Unique:  true,
				Columns: []*schema.Column{FriendshipsColumns[2], FriendshipsColumns[3]},
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Default: "Unknown"},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// TodosColumns holds the columns for the "todos" table.
	TodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"IN_PROGRESS", "COMPLETED", "PENDING"}},
		{Name: "priority", Type: field.TypeInt, Default: 0},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "blob", Type: field.TypeBytes, Nullable: true},
		{Name: "init", Type: field.TypeJSON, Nullable: true},
		{Name: "custom", Type: field.TypeJSON, Nullable: true},
		{Name: "customp", Type: field.TypeJSON, Nullable: true},
		{Name: "value", Type: field.TypeInt, Default: 0},
		{Name: "category_id", Type: field.TypeUUID, Nullable: true},
		{Name: "todo_children", Type: field.TypeUUID, Nullable: true},
		{Name: "todo_secret", Type: field.TypeUUID, Nullable: true},
	}
	// TodosTable holds the schema information for the "todos" table.
	TodosTable = &schema.Table{
		Name:       "todos",
		Columns:    TodosColumns,
		PrimaryKey: []*schema.Column{TodosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "todos_categories_todos",
				Columns:    []*schema.Column{TodosColumns[10]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "todos_todos_children",
				Columns:    []*schema.Column{TodosColumns[11]},
				RefColumns: []*schema.Column{TodosColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "todos_very_secrets_secret",
				Columns:    []*schema.Column{TodosColumns[12]},
				RefColumns: []*schema.Column{VerySecretsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Default: "Anonymous"},
		{Name: "username", Type: field.TypeUUID},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "required_metadata", Type: field.TypeJSON},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// VerySecretsColumns holds the columns for the "very_secrets" table.
	VerySecretsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "password", Type: field.TypeString},
	}
	// VerySecretsTable holds the schema information for the "very_secrets" table.
	VerySecretsTable = &schema.Table{
		Name:       "very_secrets",
		Columns:    VerySecretsColumns,
		PrimaryKey: []*schema.Column{VerySecretsColumns[0]},
	}
	// CategorySubCategoriesColumns holds the columns for the "category_sub_categories" table.
	CategorySubCategoriesColumns = []*schema.Column{
		{Name: "category_id", Type: field.TypeUUID},
		{Name: "sub_category_id", Type: field.TypeUUID},
	}
	// CategorySubCategoriesTable holds the schema information for the "category_sub_categories" table.
	CategorySubCategoriesTable = &schema.Table{
		Name:       "category_sub_categories",
		Columns:    CategorySubCategoriesColumns,
		PrimaryKey: []*schema.Column{CategorySubCategoriesColumns[0], CategorySubCategoriesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "category_sub_categories_category_id",
				Columns:    []*schema.Column{CategorySubCategoriesColumns[0]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "category_sub_categories_sub_category_id",
				Columns:    []*schema.Column{CategorySubCategoriesColumns[1]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserGroupsColumns holds the columns for the "user_groups" table.
	UserGroupsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "group_id", Type: field.TypeUUID},
	}
	// UserGroupsTable holds the schema information for the "user_groups" table.
	UserGroupsTable = &schema.Table{
		Name:       "user_groups",
		Columns:    UserGroupsColumns,
		PrimaryKey: []*schema.Column{UserGroupsColumns[0], UserGroupsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_groups_user_id",
				Columns:    []*schema.Column{UserGroupsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_groups_group_id",
				Columns:    []*schema.Column{UserGroupsColumns[1]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BillProductsTable,
		CategoriesTable,
		FriendshipsTable,
		GroupsTable,
		TodosTable,
		UsersTable,
		VerySecretsTable,
		CategorySubCategoriesTable,
		UserGroupsTable,
	}
)

func init() {
	FriendshipsTable.ForeignKeys[0].RefTable = UsersTable
	FriendshipsTable.ForeignKeys[1].RefTable = UsersTable
	TodosTable.ForeignKeys[0].RefTable = CategoriesTable
	TodosTable.ForeignKeys[1].RefTable = TodosTable
	TodosTable.ForeignKeys[2].RefTable = VerySecretsTable
	CategorySubCategoriesTable.ForeignKeys[0].RefTable = CategoriesTable
	CategorySubCategoriesTable.ForeignKeys[1].RefTable = CategoriesTable
	UserGroupsTable.ForeignKeys[0].RefTable = UsersTable
	UserGroupsTable.ForeignKeys[1].RefTable = GroupsTable
}
