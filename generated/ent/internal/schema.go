// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = "{\"Schema\":\"github.com/satisfactorymodding/smr-api/db/schema\",\"Package\":\"github.com/satisfactorymodding/smr-api/generated/ent\",\"Schemas\":[{\"name\":\"Announcement\",\"config\":{\"Table\":\"\"},\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}},{\"name\":\"message\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"importance\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"id\"]},{\"fields\":[\"deleted_at\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}]},{\"name\":\"Guide\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"user\",\"type\":\"User\",\"field\":\"user_id\",\"ref_name\":\"guides\",\"unique\":true,\"inverse\":true},{\"name\":\"tags\",\"type\":\"Tag\",\"through\":{\"N\":\"guide_tags\",\"T\":\"GuideTag\"}}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}},{\"name\":\"user_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":32,\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"short_description\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":128,\"validators\":1,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"guide\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"views\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":2,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"id\"]},{\"fields\":[\"deleted_at\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}]},{\"name\":\"GuideTag\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"guide\",\"type\":\"Guide\",\"field\":\"guide_id\",\"unique\":true,\"required\":true},{\"name\":\"tag\",\"type\":\"Tag\",\"field\":\"tag_id\",\"unique\":true,\"required\":true}],\"fields\":[{\"name\":\"guide_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"tag_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}}],\"annotations\":{\"Fields\":{\"ID\":[\"guide_id\",\"tag_id\"],\"StructTag\":null}}},{\"name\":\"Mod\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"versions\",\"type\":\"Version\"},{\"name\":\"authors\",\"type\":\"User\",\"ref_name\":\"mods\",\"through\":{\"N\":\"user_mods\",\"T\":\"UserMod\"},\"inverse\":true},{\"name\":\"tags\",\"type\":\"Tag\",\"through\":{\"N\":\"mod_tags\",\"T\":\"ModTag\"}},{\"name\":\"dependents\",\"type\":\"Version\",\"ref_name\":\"dependencies\",\"through\":{\"N\":\"version_dependencies\",\"T\":\"VersionDependency\"},\"inverse\":true}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":32,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"short_description\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":128,\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"full_description\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"logo\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"source_url\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"creator_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"approved\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":false,\"default_kind\":1,\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"views\",\"type\":{\"Type\":17,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":7,\"position\":{\"Index\":7,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"hotness\",\"type\":{\"Type\":17,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":7,\"position\":{\"Index\":8,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"popularity\",\"type\":{\"Type\":17,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":7,\"position\":{\"Index\":9,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"downloads\",\"type\":{\"Type\":17,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":7,\"position\":{\"Index\":10,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"denied\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":false,\"default_kind\":1,\"position\":{\"Index\":11,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"last_version_date\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":12,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"mod_reference\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":32,\"unique\":true,\"validators\":1,\"position\":{\"Index\":13,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"hidden\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":false,\"default_kind\":1,\"position\":{\"Index\":14,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"compatibility\",\"type\":{\"Type\":3,\"Ident\":\"*util.CompatibilityInfo\",\"PkgPath\":\"github.com/satisfactorymodding/smr-api/util\",\"PkgName\":\"util\",\"Nillable\":true,\"RType\":{\"Name\":\"CompatibilityInfo\",\"Ident\":\"util.CompatibilityInfo\",\"Kind\":22,\"PkgPath\":\"github.com/satisfactorymodding/smr-api/util\",\"Methods\":{}}},\"optional\":true,\"position\":{\"Index\":15,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"id\"]},{\"fields\":[\"deleted_at\"]},{\"fields\":[\"last_version_date\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}]},{\"name\":\"ModTag\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"mod\",\"type\":\"Mod\",\"field\":\"mod_id\",\"unique\":true,\"required\":true},{\"name\":\"tag\",\"type\":\"Tag\",\"field\":\"tag_id\",\"unique\":true,\"required\":true}],\"fields\":[{\"name\":\"mod_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"tag_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}}],\"annotations\":{\"Fields\":{\"ID\":[\"mod_id\",\"tag_id\"],\"StructTag\":null}}},{\"name\":\"SmlVersion\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"targets\",\"type\":\"SmlVersionTarget\"}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}},{\"name\":\"version\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":32,\"unique\":true,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"satisfactory_version\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"stability\",\"type\":{\"Type\":6,\"Ident\":\"util.Stability\",\"PkgPath\":\"github.com/satisfactorymodding/smr-api/util\",\"PkgName\":\"util\",\"Nillable\":false,\"RType\":{\"Name\":\"Stability\",\"Ident\":\"util.Stability\",\"Kind\":24,\"PkgPath\":\"github.com/satisfactorymodding/smr-api/util\",\"Methods\":{\"Values\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"enums\":[{\"N\":\"release\",\"V\":\"release\"},{\"N\":\"beta\",\"V\":\"beta\"},{\"N\":\"alpha\",\"V\":\"alpha\"}],\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"date\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"link\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"changelog\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"bootstrap_version\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":14,\"optional\":true,\"validators\":1,\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"engine_version\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":16,\"default\":true,\"default_value\":\"4.26\",\"default_kind\":24,\"validators\":1,\"position\":{\"Index\":7,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"id\"]},{\"fields\":[\"deleted_at\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}]},{\"name\":\"SmlVersionTarget\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"sml_version\",\"type\":\"SmlVersion\",\"field\":\"version_id\",\"ref_name\":\"targets\",\"unique\":true,\"inverse\":true,\"required\":true}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"version_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"target_name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"link\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"id\"]},{\"unique\":true,\"fields\":[\"version_id\",\"target_name\"]}]},{\"name\":\"Tag\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"mods\",\"type\":\"Mod\",\"ref_name\":\"tags\",\"through\":{\"N\":\"mod_tags\",\"T\":\"ModTag\"},\"inverse\":true},{\"name\":\"guides\",\"type\":\"Guide\",\"ref_name\":\"tags\",\"through\":{\"N\":\"guide_tags\",\"T\":\"GuideTag\"},\"inverse\":true}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":24,\"unique\":true,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"description\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":512,\"unique\":true,\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"id\"]},{\"fields\":[\"deleted_at\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}]},{\"name\":\"User\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"guides\",\"type\":\"Guide\"},{\"name\":\"sessions\",\"type\":\"UserSession\",\"storage_key\":{\"Table\":\"\",\"Symbols\":null,\"Columns\":[\"user_id\"]}},{\"name\":\"mods\",\"type\":\"Mod\",\"through\":{\"N\":\"user_mods\",\"T\":\"UserMod\"}},{\"name\":\"groups\",\"type\":\"UserGroup\"}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}},{\"name\":\"email\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":256,\"unique\":true,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"username\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":32,\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"avatar\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"joined_from\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"banned\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":false,\"default_kind\":1,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"rank\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":1,\"default_kind\":2,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"github_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":16,\"unique\":true,\"optional\":true,\"validators\":1,\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"google_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":16,\"unique\":true,\"optional\":true,\"validators\":1,\"position\":{\"Index\":7,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"facebook_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":16,\"unique\":true,\"optional\":true,\"validators\":1,\"position\":{\"Index\":8,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"id\"]},{\"fields\":[\"deleted_at\"]},{\"fields\":[\"email\"]},{\"fields\":[\"github_id\"]},{\"fields\":[\"google_id\"]},{\"fields\":[\"facebook_id\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}]},{\"name\":\"UserGroup\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"user\",\"type\":\"User\",\"field\":\"user_id\",\"ref_name\":\"groups\",\"unique\":true,\"inverse\":true,\"required\":true}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}},{\"name\":\"user_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":14,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"group_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":14,\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"id\"]},{\"fields\":[\"deleted_at\"]},{\"unique\":true,\"fields\":[\"user_id\",\"group_id\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}]},{\"name\":\"UserMod\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"user\",\"type\":\"User\",\"field\":\"user_id\",\"unique\":true,\"required\":true},{\"name\":\"mod\",\"type\":\"Mod\",\"field\":\"mod_id\",\"unique\":true,\"required\":true}],\"fields\":[{\"name\":\"user_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"mod_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"role\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}}],\"annotations\":{\"Fields\":{\"ID\":[\"user_id\",\"mod_id\"],\"StructTag\":null}}},{\"name\":\"UserSession\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"user\",\"type\":\"User\",\"ref_name\":\"sessions\",\"unique\":true,\"inverse\":true,\"required\":true}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}},{\"name\":\"token\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":512,\"unique\":true,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"user_agent\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"id\"]},{\"fields\":[\"deleted_at\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}]},{\"name\":\"Version\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"mod\",\"type\":\"Mod\",\"field\":\"mod_id\",\"ref_name\":\"versions\",\"unique\":true,\"inverse\":true,\"required\":true},{\"name\":\"dependencies\",\"type\":\"Mod\",\"through\":{\"N\":\"version_dependencies\",\"T\":\"VersionDependency\"}},{\"name\":\"targets\",\"type\":\"VersionTarget\"}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}},{\"name\":\"mod_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"version\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":16,\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"sml_version\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":16,\"validators\":1,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"changelog\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"downloads\",\"type\":{\"Type\":17,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":7,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"key\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"stability\",\"type\":{\"Type\":6,\"Ident\":\"util.Stability\",\"PkgPath\":\"github.com/satisfactorymodding/smr-api/util\",\"PkgName\":\"util\",\"Nillable\":false,\"RType\":{\"Name\":\"Stability\",\"Ident\":\"util.Stability\",\"Kind\":24,\"PkgPath\":\"github.com/satisfactorymodding/smr-api/util\",\"Methods\":{\"Values\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"enums\":[{\"N\":\"release\",\"V\":\"release\"},{\"N\":\"beta\",\"V\":\"beta\"},{\"N\":\"alpha\",\"V\":\"alpha\"}],\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"approved\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":false,\"default_kind\":1,\"position\":{\"Index\":7,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"hotness\",\"type\":{\"Type\":17,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":0,\"default_kind\":7,\"position\":{\"Index\":8,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"denied\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":false,\"default_kind\":1,\"position\":{\"Index\":9,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"metadata\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":10,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"mod_reference\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":32,\"validators\":1,\"position\":{\"Index\":11,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"version_major\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":12,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"version_minor\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":13,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"version_patch\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":14,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"size\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":15,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"hash\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":64,\"optional\":true,\"validators\":2,\"position\":{\"Index\":16,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"id\"]},{\"fields\":[\"deleted_at\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":2}]},{\"name\":\"VersionDependency\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"version\",\"type\":\"Version\",\"field\":\"version_id\",\"unique\":true,\"required\":true},{\"name\":\"mod\",\"type\":\"Mod\",\"field\":\"mod_id\",\"unique\":true,\"required\":true}],\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"version_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"mod_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"condition\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":64,\"validators\":1,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"optional\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"deleted_at\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}],\"annotations\":{\"Fields\":{\"ID\":[\"version_id\",\"mod_id\"],\"StructTag\":null}}},{\"name\":\"VersionTarget\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"sml_version\",\"type\":\"Version\",\"field\":\"version_id\",\"ref_name\":\"targets\",\"unique\":true,\"inverse\":true,\"required\":true}],\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"version_id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"target_name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"key\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"hash\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"size\",\"type\":{\"Type\":13,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"id\"]},{\"unique\":true,\"fields\":[\"version_id\",\"target_name\"]}]}],\"Features\":[\"sql/modifier\",\"intercept\",\"schema/snapshot\",\"sql/execquery\",\"sql/upsert\"]}"
