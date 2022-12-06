package services

import (
	"errors"
	"github.com/kamva/mgm/v3"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	db "mitte-task/models/db"
	"mitte-task/utils"
	"sort"
)

// CreateProfile create profile for a user
func CreateProfile(userId string, age int64, minAge int64, maxAge int64, interestIn []string, genderFilter string, preferredLocation []string, swipes []map[string]interface{}) (*db.Profile, error) {

	profile := db.NewProfile(userId, age, minAge, maxAge, interestIn, genderFilter, preferredLocation, swipes)
	err := mgm.Coll(profile).Create(profile)
	if err != nil {
		return nil, errors.New("cannot create profile for a user")
	}

	return profile, nil
}

// MatchProfiles return list of profile with potential match with user
func MatchProfiles(userId string) ([]db.Profile, error) {
	matchedProfiles, err := FindUserProfiles(userId)
	if err != nil {
		return []db.Profile{}, errors.New("cannot find matched profiles from user id")
	}

	return matchedProfiles, nil
}

// FindUserProfiles find user by id
func FindUserProfiles(userId string) ([]db.Profile, error) {
	var profile []db.Profile
	err := mgm.Coll(&db.Profile{}).SimpleFind(&profile, bson.M{})
	if err != nil {
		return []db.Profile{}, errors.New("cannot find user by id")
	}

	userProfile, err := FindProfileFromUserId(userId)
	if err != nil {
		return []db.Profile{}, errors.New("failed to find user profile from id")
	}

	objID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return []db.Profile{}, errors.New("failed to convert user id to object id")
	}
	userInfo, err := FindUserById(objID)

	if err != nil {
		return []db.Profile{}, errors.New("failed to find user profile from id")
	}

	//// exclude those profiles that user has swiped.
	for i, item := range profile {
		for _, userswipe := range item.Swipe {
			if userId == cast.ToString(userswipe["userid"]) {
				profile = append(profile[:i], profile[i+1:]...)
			}
		}
	}

	// find potential match profile for the given user id
	var matchProfiles []db.Profile
	for _, value := range profile {
		if value.UserId != userProfile.UserId {
			if utils.Contains(value.PreferredLocation, userInfo.Location) &&
				value.AgeFilterMin <= int64(userInfo.Age) && value.AgeFilterMax >= int64(userInfo.Age) &&
				utils.SubSlice(userProfile.InterestedIn, value.InterestedIn) {
				matchProfiles = append(matchProfiles, value)
			}
		}
	}

	// Apply sort on behalf of swipe statistics, the highest number of profile swipes will become first on potential match based on the ranking
	sort.Slice(matchProfiles[:], func(i, j int) bool {
		return len(matchProfiles[i].Swipe) > len(matchProfiles[j].Swipe)
	})

	return matchProfiles, nil
}

// FindProfileFromUserId find profile of user by id
func FindProfileFromUserId(userId string) (*db.Profile, error) {
	profile := &db.Profile{}
	err := mgm.Coll(profile).First(bson.M{"user_id": userId}, profile)
	if err != nil {
		return nil, errors.New("cannot find user")
	}

	return profile, nil
}

func FilteredProfiles(age int64, gender string) ([]db.Profile, error) {
	profile := []db.Profile{}

	if len(gender) != 0 && age != 0 {
		err := mgm.Coll(&db.Profile{}).SimpleFind(&profile, bson.M{"age": age, "gender": gender})
		if err != nil {
			return nil, errors.New("cannot find user")
		}
	} else if len(gender) != 0 {
		err := mgm.Coll(&db.Profile{}).SimpleFind(&profile, bson.M{"gender": gender})
		if err != nil {
			return nil, errors.New("cannot find user")
		}
	} else {
		err := mgm.Coll(&db.Profile{}).SimpleFind(&profile, bson.M{"age": age})
		if err != nil {
			return nil, errors.New("cannot find user")
		}
	}

	return profile, nil
}
