// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

type Announcement struct {
	ID         string                 `json:"id"`
	Message    string                 `json:"message"`
	Importance AnnouncementImportance `json:"importance"`
}

type Compatibility struct {
	State CompatibilityState `json:"state"`
	Note  *string            `json:"note,omitempty"`
}

type CompatibilityInfo struct {
	Ea  *Compatibility `json:"EA"`
	Exp *Compatibility `json:"EXP"`
}

type CompatibilityInfoInput struct {
	Ea  *CompatibilityInput `json:"EA"`
	Exp *CompatibilityInput `json:"EXP"`
}

type CompatibilityInput struct {
	State CompatibilityState `json:"state"`
	Note  *string            `json:"note,omitempty"`
}

type CreateVersionResponse struct {
	AutoApproved bool     `json:"auto_approved"`
	Version      *Version `json:"version,omitempty"`
}

type GetGuides struct {
	Guides []*Guide `json:"guides"`
	Count  int      `json:"count"`
}

type GetMods struct {
	Mods  []*Mod `json:"mods"`
	Count int    `json:"count"`
}

type GetMyMods struct {
	Mods  []*Mod `json:"mods"`
	Count int    `json:"count"`
}

type GetMyVersions struct {
	Versions []*Version `json:"versions"`
	Count    int        `json:"count"`
}

type GetSMLVersions struct {
	SmlVersions []*SMLVersion `json:"sml_versions"`
	Count       int           `json:"count"`
}

type GetVersions struct {
	Versions []*Version `json:"versions"`
	Count    int        `json:"count"`
}

type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Guide struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Guide            string `json:"guide"`
	Views            int    `json:"views"`
	UserID           string `json:"user_id"`
	UpdatedAt        string `json:"updated_at"`
	CreatedAt        string `json:"created_at"`
	Tags             []*Tag `json:"tags"`
	User             *User  `json:"user"`
}

type LatestVersions struct {
	Alpha   *Version `json:"alpha,omitempty"`
	Beta    *Version `json:"beta,omitempty"`
	Release *Version `json:"release,omitempty"`
}

type Mod struct {
	ID                    string             `json:"id"`
	Name                  string             `json:"name"`
	ShortDescription      string             `json:"short_description"`
	FullDescription       *string            `json:"full_description,omitempty"`
	Logo                  *string            `json:"logo,omitempty"`
	LogoThumbhash         *string            `json:"logo_thumbhash,omitempty"`
	SourceURL             *string            `json:"source_url,omitempty"`
	CreatorID             string             `json:"creator_id"`
	Approved              bool               `json:"approved"`
	Views                 int                `json:"views"`
	Downloads             int                `json:"downloads"`
	Hotness               int                `json:"hotness"`
	Popularity            int                `json:"popularity"`
	UpdatedAt             string             `json:"updated_at"`
	CreatedAt             string             `json:"created_at"`
	LastVersionDate       *string            `json:"last_version_date,omitempty"`
	ModReference          string             `json:"mod_reference"`
	Hidden                bool               `json:"hidden"`
	Tags                  []*Tag             `json:"tags,omitempty"`
	Compatibility         *CompatibilityInfo `json:"compatibility,omitempty"`
	ToggleNetworkUse      bool               `json:"toggle_network_use"`
	ToggleExplicitContent bool               `json:"toggle_explicit_content"`
	Authors               []*UserMod         `json:"authors"`
	Version               *Version           `json:"version,omitempty"`
	Versions              []*Version         `json:"versions"`
	LatestVersions        *LatestVersions    `json:"latestVersions"`
}

type ModVersion struct {
	ID           string     `json:"id"`
	ModReference string     `json:"mod_reference"`
	Versions     []*Version `json:"versions"`
}

type ModVersionConstraint struct {
	ModIDOrReference string `json:"modIdOrReference"`
	Version          string `json:"version"`
}

type Mutation struct {
}

type NewAnnouncement struct {
	Message    string                 `json:"message"`
	Importance AnnouncementImportance `json:"importance"`
}

type NewGuide struct {
	Name             string   `json:"name"`
	ShortDescription string   `json:"short_description"`
	Guide            string   `json:"guide"`
	TagIDs           []string `json:"tagIDs,omitempty"`
}

type NewSatisfactoryVersion struct {
	Version       int    `json:"version"`
	EngineVersion string `json:"engine_version"`
}

type NewTag struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type NewVersion struct {
	Changelog string             `json:"changelog"`
	Stability VersionStabilities `json:"stability"`
}

type OAuthOptions struct {
	Github   string `json:"github"`
	Google   string `json:"google"`
	Facebook string `json:"facebook"`
}

type Query struct {
}

type SMLVersion struct {
	ID                  string              `json:"id"`
	Version             string              `json:"version"`
	SatisfactoryVersion int                 `json:"satisfactory_version"`
	Stability           VersionStabilities  `json:"stability"`
	Link                string              `json:"link"`
	Targets             []*SMLVersionTarget `json:"targets"`
	Changelog           string              `json:"changelog"`
	Date                string              `json:"date"`
	BootstrapVersion    *string             `json:"bootstrap_version,omitempty"`
	EngineVersion       string              `json:"engine_version"`
	UpdatedAt           string              `json:"updated_at"`
	CreatedAt           string              `json:"created_at"`
}

type SMLVersionTarget struct {
	VersionID  string     `json:"VersionID"`
	TargetName TargetName `json:"targetName"`
	Link       string     `json:"link"`
}

type SatisfactoryVersion struct {
	ID            string `json:"id"`
	Version       int    `json:"version"`
	EngineVersion string `json:"engine_version"`
}

type Tag struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TagFilter struct {
	Limit  *int     `json:"limit,omitempty"`
	Offset *int     `json:"offset,omitempty"`
	Order  *Order   `json:"order,omitempty"`
	Search *string  `json:"search,omitempty"`
	Ids    []string `json:"ids,omitempty"`
}

type UpdateAnnouncement struct {
	Message    *string                 `json:"message,omitempty"`
	Importance *AnnouncementImportance `json:"importance,omitempty"`
}

type UpdateGuide struct {
	Name             *string  `json:"name,omitempty"`
	ShortDescription *string  `json:"short_description,omitempty"`
	Guide            *string  `json:"guide,omitempty"`
	TagIDs           []string `json:"tagIDs,omitempty"`
}

type UpdateSatisfactoryVersion struct {
	Version       *int    `json:"version,omitempty"`
	EngineVersion *string `json:"engine_version,omitempty"`
}

type UpdateUser struct {
	Avatar   *graphql.Upload `json:"avatar,omitempty"`
	Groups   []string        `json:"groups,omitempty"`
	Username *string         `json:"username,omitempty"`
}

type UpdateUserMod struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

type UpdateVersion struct {
	Changelog *string             `json:"changelog,omitempty"`
	Stability *VersionStabilities `json:"stability,omitempty"`
}

type User struct {
	ID              string     `json:"id"`
	Email           *string    `json:"email,omitempty"`
	Username        string     `json:"username"`
	Avatar          *string    `json:"avatar,omitempty"`
	AvatarThumbhash *string    `json:"avatar_thumbhash,omitempty"`
	CreatedAt       string     `json:"created_at"`
	GithubID        *string    `json:"github_id,omitempty"`
	GoogleID        *string    `json:"google_id,omitempty"`
	FacebookID      *string    `json:"facebook_id,omitempty"`
	Roles           *UserRoles `json:"roles"`
	Groups          []*Group   `json:"groups"`
	Mods            []*UserMod `json:"mods"`
	Guides          []*Guide   `json:"guides"`
}

type UserMod struct {
	UserID string `json:"user_id"`
	ModID  string `json:"mod_id"`
	Role   string `json:"role"`
	User   *User  `json:"user"`
	Mod    *Mod   `json:"mod"`
}

type UserRoles struct {
	ApproveMods              bool `json:"approveMods"`
	ApproveVersions          bool `json:"approveVersions"`
	DeleteContent            bool `json:"deleteContent"`
	EditContent              bool `json:"editContent"`
	EditUsers                bool `json:"editUsers"`
	EditSatisfactoryVersions bool `json:"editSatisfactoryVersions"`
	EditBootstrapVersions    bool `json:"editBootstrapVersions"`
	EditAnyModCompatibility  bool `json:"editAnyModCompatibility"`
}

type UserSession struct {
	Token string `json:"token"`
}

type Version struct {
	ID               string               `json:"id"`
	ModID            string               `json:"mod_id"`
	Version          string               `json:"version"`
	SmlVersion       string               `json:"sml_version"`
	GameVersion      string               `json:"game_version"`
	RequiredOnRemote bool                 `json:"required_on_remote"`
	Changelog        string               `json:"changelog"`
	Downloads        int                  `json:"downloads"`
	Stability        VersionStabilities   `json:"stability"`
	Approved         bool                 `json:"approved"`
	UpdatedAt        string               `json:"updated_at"`
	CreatedAt        string               `json:"created_at"`
	Link             string               `json:"link"`
	Targets          []*VersionTarget     `json:"targets"`
	Metadata         *string              `json:"metadata,omitempty"`
	Size             *int                 `json:"size,omitempty"`
	Hash             *string              `json:"hash,omitempty"`
	Mod              *Mod                 `json:"mod"`
	Dependencies     []*VersionDependency `json:"dependencies"`
}

type VersionDependency struct {
	VersionID    string   `json:"version_id"`
	ModID        string   `json:"mod_id"`
	Condition    string   `json:"condition"`
	Optional     bool     `json:"optional"`
	ModReference string   `json:"mod_reference"`
	Mod          *Mod     `json:"mod,omitempty"`
	Version      *Version `json:"version,omitempty"`
}

type VersionTarget struct {
	VersionID  string     `json:"VersionID"`
	TargetName TargetName `json:"targetName"`
	Link       string     `json:"link"`
	Size       *int       `json:"size,omitempty"`
	Hash       *string    `json:"hash,omitempty"`
}

type AnnouncementImportance string

const (
	AnnouncementImportanceFix     AnnouncementImportance = "Fix"
	AnnouncementImportanceInfo    AnnouncementImportance = "Info"
	AnnouncementImportanceWarning AnnouncementImportance = "Warning"
	AnnouncementImportanceAlert   AnnouncementImportance = "Alert"
)

var AllAnnouncementImportance = []AnnouncementImportance{
	AnnouncementImportanceFix,
	AnnouncementImportanceInfo,
	AnnouncementImportanceWarning,
	AnnouncementImportanceAlert,
}

func (e AnnouncementImportance) IsValid() bool {
	switch e {
	case AnnouncementImportanceFix, AnnouncementImportanceInfo, AnnouncementImportanceWarning, AnnouncementImportanceAlert:
		return true
	}
	return false
}

func (e AnnouncementImportance) String() string {
	return string(e)
}

func (e *AnnouncementImportance) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AnnouncementImportance(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AnnouncementImportance", str)
	}
	return nil
}

func (e AnnouncementImportance) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CompatibilityState string

const (
	CompatibilityStateWorks   CompatibilityState = "Works"
	CompatibilityStateDamaged CompatibilityState = "Damaged"
	CompatibilityStateBroken  CompatibilityState = "Broken"
)

var AllCompatibilityState = []CompatibilityState{
	CompatibilityStateWorks,
	CompatibilityStateDamaged,
	CompatibilityStateBroken,
}

func (e CompatibilityState) IsValid() bool {
	switch e {
	case CompatibilityStateWorks, CompatibilityStateDamaged, CompatibilityStateBroken:
		return true
	}
	return false
}

func (e CompatibilityState) String() string {
	return string(e)
}

func (e *CompatibilityState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CompatibilityState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CompatibilityState", str)
	}
	return nil
}

func (e CompatibilityState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type GuideFields string

const (
	GuideFieldsName      GuideFields = "name"
	GuideFieldsCreatedAt GuideFields = "created_at"
	GuideFieldsUpdatedAt GuideFields = "updated_at"
	GuideFieldsViews     GuideFields = "views"
)

var AllGuideFields = []GuideFields{
	GuideFieldsName,
	GuideFieldsCreatedAt,
	GuideFieldsUpdatedAt,
	GuideFieldsViews,
}

func (e GuideFields) IsValid() bool {
	switch e {
	case GuideFieldsName, GuideFieldsCreatedAt, GuideFieldsUpdatedAt, GuideFieldsViews:
		return true
	}
	return false
}

func (e GuideFields) String() string {
	return string(e)
}

func (e *GuideFields) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GuideFields(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid GuideFields", str)
	}
	return nil
}

func (e GuideFields) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ModFields string

const (
	ModFieldsCreatedAt       ModFields = "created_at"
	ModFieldsUpdatedAt       ModFields = "updated_at"
	ModFieldsName            ModFields = "name"
	ModFieldsViews           ModFields = "views"
	ModFieldsDownloads       ModFields = "downloads"
	ModFieldsHotness         ModFields = "hotness"
	ModFieldsPopularity      ModFields = "popularity"
	ModFieldsLastVersionDate ModFields = "last_version_date"
	ModFieldsSearch          ModFields = "search"
)

var AllModFields = []ModFields{
	ModFieldsCreatedAt,
	ModFieldsUpdatedAt,
	ModFieldsName,
	ModFieldsViews,
	ModFieldsDownloads,
	ModFieldsHotness,
	ModFieldsPopularity,
	ModFieldsLastVersionDate,
	ModFieldsSearch,
}

func (e ModFields) IsValid() bool {
	switch e {
	case ModFieldsCreatedAt, ModFieldsUpdatedAt, ModFieldsName, ModFieldsViews, ModFieldsDownloads, ModFieldsHotness, ModFieldsPopularity, ModFieldsLastVersionDate, ModFieldsSearch:
		return true
	}
	return false
}

func (e ModFields) String() string {
	return string(e)
}

func (e *ModFields) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ModFields(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ModFields", str)
	}
	return nil
}

func (e ModFields) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Order string

const (
	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)

var AllOrder = []Order{
	OrderAsc,
	OrderDesc,
}

func (e Order) IsValid() bool {
	switch e {
	case OrderAsc, OrderDesc:
		return true
	}
	return false
}

func (e Order) String() string {
	return string(e)
}

func (e *Order) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Order(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Order", str)
	}
	return nil
}

func (e Order) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SMLVersionFields string

const (
	SMLVersionFieldsName                SMLVersionFields = "name"
	SMLVersionFieldsCreatedAt           SMLVersionFields = "created_at"
	SMLVersionFieldsUpdatedAt           SMLVersionFields = "updated_at"
	SMLVersionFieldsSatisfactoryVersion SMLVersionFields = "satisfactory_version"
	SMLVersionFieldsDate                SMLVersionFields = "date"
)

var AllSMLVersionFields = []SMLVersionFields{
	SMLVersionFieldsName,
	SMLVersionFieldsCreatedAt,
	SMLVersionFieldsUpdatedAt,
	SMLVersionFieldsSatisfactoryVersion,
	SMLVersionFieldsDate,
}

func (e SMLVersionFields) IsValid() bool {
	switch e {
	case SMLVersionFieldsName, SMLVersionFieldsCreatedAt, SMLVersionFieldsUpdatedAt, SMLVersionFieldsSatisfactoryVersion, SMLVersionFieldsDate:
		return true
	}
	return false
}

func (e SMLVersionFields) String() string {
	return string(e)
}

func (e *SMLVersionFields) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SMLVersionFields(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SMLVersionFields", str)
	}
	return nil
}

func (e SMLVersionFields) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TargetName string

const (
	TargetNameWindows       TargetName = "Windows"
	TargetNameWindowsServer TargetName = "WindowsServer"
	TargetNameLinuxServer   TargetName = "LinuxServer"
)

var AllTargetName = []TargetName{
	TargetNameWindows,
	TargetNameWindowsServer,
	TargetNameLinuxServer,
}

func (e TargetName) IsValid() bool {
	switch e {
	case TargetNameWindows, TargetNameWindowsServer, TargetNameLinuxServer:
		return true
	}
	return false
}

func (e TargetName) String() string {
	return string(e)
}

func (e *TargetName) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TargetName(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TargetName", str)
	}
	return nil
}

func (e TargetName) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type VersionFields string

const (
	VersionFieldsCreatedAt VersionFields = "created_at"
	VersionFieldsUpdatedAt VersionFields = "updated_at"
	VersionFieldsDownloads VersionFields = "downloads"
)

var AllVersionFields = []VersionFields{
	VersionFieldsCreatedAt,
	VersionFieldsUpdatedAt,
	VersionFieldsDownloads,
}

func (e VersionFields) IsValid() bool {
	switch e {
	case VersionFieldsCreatedAt, VersionFieldsUpdatedAt, VersionFieldsDownloads:
		return true
	}
	return false
}

func (e VersionFields) String() string {
	return string(e)
}

func (e *VersionFields) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VersionFields(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid VersionFields", str)
	}
	return nil
}

func (e VersionFields) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type VersionStabilities string

const (
	VersionStabilitiesAlpha   VersionStabilities = "alpha"
	VersionStabilitiesBeta    VersionStabilities = "beta"
	VersionStabilitiesRelease VersionStabilities = "release"
)

var AllVersionStabilities = []VersionStabilities{
	VersionStabilitiesAlpha,
	VersionStabilitiesBeta,
	VersionStabilitiesRelease,
}

func (e VersionStabilities) IsValid() bool {
	switch e {
	case VersionStabilitiesAlpha, VersionStabilitiesBeta, VersionStabilitiesRelease:
		return true
	}
	return false
}

func (e VersionStabilities) String() string {
	return string(e)
}

func (e *VersionStabilities) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VersionStabilities(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid VersionStabilities", str)
	}
	return nil
}

func (e VersionStabilities) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
