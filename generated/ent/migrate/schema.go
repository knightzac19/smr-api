// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AnnouncementsColumns holds the columns for the "announcements" table.
	AnnouncementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "message", Type: field.TypeString},
		{Name: "importance", Type: field.TypeString},
	}
	// AnnouncementsTable holds the schema information for the "announcements" table.
	AnnouncementsTable = &schema.Table{
		Name:       "announcements",
		Columns:    AnnouncementsColumns,
		PrimaryKey: []*schema.Column{AnnouncementsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "announcement_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{AnnouncementsColumns[3]},
			},
		},
	}
	// GuidesColumns holds the columns for the "guides" table.
	GuidesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Size: 32},
		{Name: "short_description", Type: field.TypeString, Size: 128},
		{Name: "guide", Type: field.TypeString},
		{Name: "views", Type: field.TypeInt, Default: 0},
		{Name: "user_id", Type: field.TypeString, Nullable: true},
	}
	// GuidesTable holds the schema information for the "guides" table.
	GuidesTable = &schema.Table{
		Name:       "guides",
		Columns:    GuidesColumns,
		PrimaryKey: []*schema.Column{GuidesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "guides_users_guides",
				Columns:    []*schema.Column{GuidesColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "guide_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{GuidesColumns[3]},
			},
		},
	}
	// GuideTagsColumns holds the columns for the "guide_tags" table.
	GuideTagsColumns = []*schema.Column{
		{Name: "guide_id", Type: field.TypeString},
		{Name: "tag_id", Type: field.TypeString},
	}
	// GuideTagsTable holds the schema information for the "guide_tags" table.
	GuideTagsTable = &schema.Table{
		Name:       "guide_tags",
		Columns:    GuideTagsColumns,
		PrimaryKey: []*schema.Column{GuideTagsColumns[0], GuideTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "guide_tags_guides_guide",
				Columns:    []*schema.Column{GuideTagsColumns[0]},
				RefColumns: []*schema.Column{GuidesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "guide_tags_tags_tag",
				Columns:    []*schema.Column{GuideTagsColumns[1]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ModsColumns holds the columns for the "mods" table.
	ModsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Size: 32},
		{Name: "short_description", Type: field.TypeString, Size: 128},
		{Name: "full_description", Type: field.TypeString},
		{Name: "logo", Type: field.TypeString},
		{Name: "logo_thumbhash", Type: field.TypeString, Nullable: true},
		{Name: "source_url", Type: field.TypeString, Nullable: true},
		{Name: "creator_id", Type: field.TypeString},
		{Name: "approved", Type: field.TypeBool, Default: false},
		{Name: "views", Type: field.TypeUint, Default: 0},
		{Name: "hotness", Type: field.TypeUint, Default: 0},
		{Name: "popularity", Type: field.TypeUint, Default: 0},
		{Name: "downloads", Type: field.TypeUint, Default: 0},
		{Name: "denied", Type: field.TypeBool, Default: false},
		{Name: "last_version_date", Type: field.TypeTime, Nullable: true},
		{Name: "mod_reference", Type: field.TypeString, Size: 32},
		{Name: "hidden", Type: field.TypeBool, Default: false},
		{Name: "compatibility", Type: field.TypeJSON, Nullable: true},
		{Name: "toggle_network_use", Type: field.TypeBool, Default: false},
		{Name: "toggle_explicit_content", Type: field.TypeBool, Default: false},
	}
	// ModsTable holds the schema information for the "mods" table.
	ModsTable = &schema.Table{
		Name:       "mods",
		Columns:    ModsColumns,
		PrimaryKey: []*schema.Column{ModsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "mod_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{ModsColumns[3]},
			},
			{
				Name:    "mod_last_version_date",
				Unique:  false,
				Columns: []*schema.Column{ModsColumns[17]},
			},
			{
				Name:    "mod_mod_reference",
				Unique:  true,
				Columns: []*schema.Column{ModsColumns[18]},
			},
		},
	}
	// ModTagsColumns holds the columns for the "mod_tags" table.
	ModTagsColumns = []*schema.Column{
		{Name: "mod_id", Type: field.TypeString},
		{Name: "tag_id", Type: field.TypeString},
	}
	// ModTagsTable holds the schema information for the "mod_tags" table.
	ModTagsTable = &schema.Table{
		Name:       "mod_tags",
		Columns:    ModTagsColumns,
		PrimaryKey: []*schema.Column{ModTagsColumns[0], ModTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "mod_tags_mods_mod",
				Columns:    []*schema.Column{ModTagsColumns[0]},
				RefColumns: []*schema.Column{ModsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "mod_tags_tags_tag",
				Columns:    []*schema.Column{ModTagsColumns[1]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SatisfactoryVersionsColumns holds the columns for the "satisfactory_versions" table.
	SatisfactoryVersionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "version", Type: field.TypeInt, Unique: true},
		{Name: "engine_version", Type: field.TypeString, Size: 16, Default: "4.26"},
	}
	// SatisfactoryVersionsTable holds the schema information for the "satisfactory_versions" table.
	SatisfactoryVersionsTable = &schema.Table{
		Name:       "satisfactory_versions",
		Columns:    SatisfactoryVersionsColumns,
		PrimaryKey: []*schema.Column{SatisfactoryVersionsColumns[0]},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 24},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 512},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "tag_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{TagsColumns[3]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "email", Type: field.TypeString, Size: 256},
		{Name: "username", Type: field.TypeString, Size: 32},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "avatar_thumbhash", Type: field.TypeString, Nullable: true},
		{Name: "joined_from", Type: field.TypeString, Nullable: true},
		{Name: "banned", Type: field.TypeBool, Default: false},
		{Name: "rank", Type: field.TypeInt, Default: 1},
		{Name: "github_id", Type: field.TypeString, Nullable: true, Size: 16},
		{Name: "google_id", Type: field.TypeString, Nullable: true, Size: 16},
		{Name: "facebook_id", Type: field.TypeString, Nullable: true, Size: 16},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[3]},
			},
			{
				Name:    "user_email",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[4]},
			},
			{
				Name:    "user_github_id",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[11]},
			},
			{
				Name:    "user_google_id",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[12]},
			},
			{
				Name:    "user_facebook_id",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[13]},
			},
		},
	}
	// UserGroupsColumns holds the columns for the "user_groups" table.
	UserGroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "group_id", Type: field.TypeString, Size: 14},
		{Name: "user_id", Type: field.TypeString},
	}
	// UserGroupsTable holds the schema information for the "user_groups" table.
	UserGroupsTable = &schema.Table{
		Name:       "user_groups",
		Columns:    UserGroupsColumns,
		PrimaryKey: []*schema.Column{UserGroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_groups_users_groups",
				Columns:    []*schema.Column{UserGroupsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "usergroup_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{UserGroupsColumns[3]},
			},
			{
				Name:    "usergroup_user_id_group_id",
				Unique:  true,
				Columns: []*schema.Column{UserGroupsColumns[5], UserGroupsColumns[4]},
			},
		},
	}
	// UserModsColumns holds the columns for the "user_mods" table.
	UserModsColumns = []*schema.Column{
		{Name: "role", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
		{Name: "mod_id", Type: field.TypeString},
	}
	// UserModsTable holds the schema information for the "user_mods" table.
	UserModsTable = &schema.Table{
		Name:       "user_mods",
		Columns:    UserModsColumns,
		PrimaryKey: []*schema.Column{UserModsColumns[1], UserModsColumns[2]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_mods_users_user",
				Columns:    []*schema.Column{UserModsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_mods_mods_mod",
				Columns:    []*schema.Column{UserModsColumns[2]},
				RefColumns: []*schema.Column{ModsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UserSessionsColumns holds the columns for the "user_sessions" table.
	UserSessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "token", Type: field.TypeString, Size: 512},
		{Name: "user_agent", Type: field.TypeString, Nullable: true},
		{Name: "user_id", Type: field.TypeString},
	}
	// UserSessionsTable holds the schema information for the "user_sessions" table.
	UserSessionsTable = &schema.Table{
		Name:       "user_sessions",
		Columns:    UserSessionsColumns,
		PrimaryKey: []*schema.Column{UserSessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_sessions_users_sessions",
				Columns:    []*schema.Column{UserSessionsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "usersession_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{UserSessionsColumns[3]},
			},
			{
				Name:    "usersession_token",
				Unique:  true,
				Columns: []*schema.Column{UserSessionsColumns[4]},
			},
		},
	}
	// VersionsColumns holds the columns for the "versions" table.
	VersionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "version", Type: field.TypeString, Size: 16},
		{Name: "game_version", Type: field.TypeString},
		{Name: "required_on_remote", Type: field.TypeBool},
		{Name: "changelog", Type: field.TypeString, Nullable: true},
		{Name: "downloads", Type: field.TypeUint, Default: 0},
		{Name: "key", Type: field.TypeString, Nullable: true},
		{Name: "stability", Type: field.TypeEnum, Enums: []string{"release", "beta", "alpha"}, Default: "release"},
		{Name: "approved", Type: field.TypeBool, Default: false},
		{Name: "hotness", Type: field.TypeUint, Default: 0},
		{Name: "denied", Type: field.TypeBool, Default: false},
		{Name: "metadata", Type: field.TypeString, Nullable: true},
		{Name: "mod_reference", Type: field.TypeString, Size: 32},
		{Name: "version_major", Type: field.TypeInt, Nullable: true},
		{Name: "version_minor", Type: field.TypeInt, Nullable: true},
		{Name: "version_patch", Type: field.TypeInt, Nullable: true},
		{Name: "size", Type: field.TypeInt64, Nullable: true},
		{Name: "hash", Type: field.TypeString, Nullable: true, Size: 64},
		{Name: "mod_id", Type: field.TypeString},
	}
	// VersionsTable holds the schema information for the "versions" table.
	VersionsTable = &schema.Table{
		Name:       "versions",
		Columns:    VersionsColumns,
		PrimaryKey: []*schema.Column{VersionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "versions_mods_versions",
				Columns:    []*schema.Column{VersionsColumns[21]},
				RefColumns: []*schema.Column{ModsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "version_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{VersionsColumns[3]},
			},
			{
				Name:    "version_approved",
				Unique:  false,
				Columns: []*schema.Column{VersionsColumns[11]},
			},
			{
				Name:    "version_denied",
				Unique:  false,
				Columns: []*schema.Column{VersionsColumns[13]},
			},
			{
				Name:    "version_mod_id",
				Unique:  false,
				Columns: []*schema.Column{VersionsColumns[21]},
			},
		},
	}
	// VersionDependenciesColumns holds the columns for the "version_dependencies" table.
	VersionDependenciesColumns = []*schema.Column{
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "condition", Type: field.TypeString, Size: 64},
		{Name: "optional", Type: field.TypeBool},
		{Name: "version_id", Type: field.TypeString},
		{Name: "mod_id", Type: field.TypeString},
	}
	// VersionDependenciesTable holds the schema information for the "version_dependencies" table.
	VersionDependenciesTable = &schema.Table{
		Name:       "version_dependencies",
		Columns:    VersionDependenciesColumns,
		PrimaryKey: []*schema.Column{VersionDependenciesColumns[5], VersionDependenciesColumns[6]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "version_dependencies_versions_version",
				Columns:    []*schema.Column{VersionDependenciesColumns[5]},
				RefColumns: []*schema.Column{VersionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "version_dependencies_mods_mod",
				Columns:    []*schema.Column{VersionDependenciesColumns[6]},
				RefColumns: []*schema.Column{ModsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "versiondependency_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{VersionDependenciesColumns[2]},
			},
		},
	}
	// VersionTargetsColumns holds the columns for the "version_targets" table.
	VersionTargetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "target_name", Type: field.TypeString},
		{Name: "key", Type: field.TypeString, Nullable: true},
		{Name: "hash", Type: field.TypeString, Nullable: true},
		{Name: "size", Type: field.TypeInt64, Nullable: true},
		{Name: "version_id", Type: field.TypeString},
	}
	// VersionTargetsTable holds the schema information for the "version_targets" table.
	VersionTargetsTable = &schema.Table{
		Name:       "version_targets",
		Columns:    VersionTargetsColumns,
		PrimaryKey: []*schema.Column{VersionTargetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "version_targets_versions_targets",
				Columns:    []*schema.Column{VersionTargetsColumns[5]},
				RefColumns: []*schema.Column{VersionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "versiontarget_version_id_target_name",
				Unique:  true,
				Columns: []*schema.Column{VersionTargetsColumns[5], VersionTargetsColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AnnouncementsTable,
		GuidesTable,
		GuideTagsTable,
		ModsTable,
		ModTagsTable,
		SatisfactoryVersionsTable,
		TagsTable,
		UsersTable,
		UserGroupsTable,
		UserModsTable,
		UserSessionsTable,
		VersionsTable,
		VersionDependenciesTable,
		VersionTargetsTable,
	}
)

func init() {
	GuidesTable.ForeignKeys[0].RefTable = UsersTable
	GuideTagsTable.ForeignKeys[0].RefTable = GuidesTable
	GuideTagsTable.ForeignKeys[1].RefTable = TagsTable
	ModTagsTable.ForeignKeys[0].RefTable = ModsTable
	ModTagsTable.ForeignKeys[1].RefTable = TagsTable
	UserGroupsTable.ForeignKeys[0].RefTable = UsersTable
	UserModsTable.ForeignKeys[0].RefTable = UsersTable
	UserModsTable.ForeignKeys[1].RefTable = ModsTable
	UserSessionsTable.ForeignKeys[0].RefTable = UsersTable
	VersionsTable.ForeignKeys[0].RefTable = ModsTable
	VersionDependenciesTable.ForeignKeys[0].RefTable = VersionsTable
	VersionDependenciesTable.ForeignKeys[1].RefTable = ModsTable
	VersionTargetsTable.ForeignKeys[0].RefTable = VersionsTable
}
