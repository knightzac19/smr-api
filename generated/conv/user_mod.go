// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package conv

import (
	generated "github.com/satisfactorymodding/smr-api/generated"
	ent "github.com/satisfactorymodding/smr-api/generated/ent"
)

type UserModImpl struct{}

func (c *UserModImpl) Convert(source *ent.UserMod) *generated.UserMod {
	var pGeneratedUserMod *generated.UserMod
	if source != nil {
		var generatedUserMod generated.UserMod
		generatedUserMod.UserID = (*source).UserID
		generatedUserMod.ModID = (*source).ModID
		generatedUserMod.Role = (*source).Role
		pGeneratedUserMod = &generatedUserMod
	}
	return pGeneratedUserMod
}
func (c *UserModImpl) ConvertSlice(source []*ent.UserMod) []*generated.UserMod {
	var pGeneratedUserModList []*generated.UserMod
	if source != nil {
		pGeneratedUserModList = make([]*generated.UserMod, len(source))
		for i := 0; i < len(source); i++ {
			pGeneratedUserModList[i] = c.Convert(source[i])
		}
	}
	return pGeneratedUserModList
}
