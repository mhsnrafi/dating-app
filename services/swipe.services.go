package services

import (
	"errors"
	"github.com/kamva/mgm/v3"
	"github.com/spf13/cast"
	db "mitte-task/models/db"
	"mitte-task/utils"
)

// SwipeProfile swipe profile to get a perfect match
func SwipeProfile(userId string, profileId string, preference string) (bool, error) {
	var match bool

	profileToMatch, err := FindUserProfileFromId(profileId)
	if err != nil {
		return false, errors.New("cannot find user profile from id")
	}

	userWhoSwipe, err := FindProfileFromUserId(userId)
	if err != nil {
		return false, errors.New("failed to find user profile from user id")
	}

	var pref string
	if len(profileToMatch.Swipe) != 0 {
		for _, swipe := range profileToMatch.Swipe {
			pref = cast.ToString(swipe["preference"])
			if cast.ToString(swipe["userid"]) != userId {
				swipeRecord := map[string]interface{}{
					"userid":     userId,
					"preference": preference,
				}
				profileToMatch.Swipe = append(profileToMatch.Swipe, swipeRecord)
				err = mgm.Coll(profileToMatch).Update(profileToMatch)
				if err != nil {
					return false, errors.New("failed to update user profile")
				}
			}
		}
	} else {
		swipeRecord := map[string]interface{}{
			"userid":     userId,
			"preference": preference,
		}
		profileToMatch.Swipe = append(profileToMatch.Swipe, swipeRecord)
		err = mgm.Coll(profileToMatch).Update(profileToMatch)
		if err != nil {
			return false, errors.New("failed to update user profile")
		}
	}

	if len(userWhoSwipe.Swipe) != 0 {
		for _, prof := range userWhoSwipe.Swipe {
			if cast.ToString(prof["userid"]) == profileToMatch.UserId && cast.ToString(prof["preference"]) == pref {
				match = true
			} else {
				match = false
			}
		}
	}

	return match, nil
}

// FindUserProfileFromId find user profile from id
func FindUserProfileFromId(profileId string) (*db.Profile, error) {
	id, err := utils.ConvertStringToObjectId(profileId)
	if err != nil {
		return nil, errors.New("failed to convert id from string to object id")
	}

	profile := &db.Profile{}
	err = mgm.Coll(profile).FindByID(id, profile)
	if err != nil {
		return nil, errors.New("cannot find user profile form id")
	}

	return profile, nil
}
