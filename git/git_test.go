package git_test

import (
	"github.com/pcarranza/sh-tools/git"
	"reflect"
	"testing"
)

func Test_ParseGitHubURL(t *testing.T) {
	tt := []struct {
		name           string
		url            string
		expectedGoPath string
		expectedError  error
	}{
		{
			"Invalid URL",
			"",
			"",
			git.ErrInvalidURL,
		},
		{
			"GitHub Git URL",
			"git@github.com:gomeeseeks/meeseeks-box.git",
			"github.com/gomeeseeks/meeseeks-box",
			nil,
		},
		{
			"GitHub Git URL without ending",
			"git@github.com:gomeeseeks/meeseeks-box",
			"github.com/gomeeseeks/meeseeks-box",
			nil,
		},
		{
			"GitHub HTTP URL",
			"https://github.com/gomeeseeks/meeseeks-box.git",
			"github.com/gomeeseeks/meeseeks-box",
			nil,
		},
		{
			"GitHub HTTP URL",
			"http://github.com/gomeeseeks/meeseeks-box.git",
			"github.com/gomeeseeks/meeseeks-box",
			nil,
		},
		{
			"GitHub HTTP URL without ending",
			"http://github.com/gomeeseeks/meeseeks-box",
			"github.com/gomeeseeks/meeseeks-box",
			nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			url, err := git.Parse(tc.url)
			assertEquals(t, tc.expectedError, err)
			if err == nil {
				assertEquals(t, tc.expectedGoPath, url.ToGoPath())
			}
		})
	}
}

func assertEquals(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Value %s is not as expected %s", actual, expected)
	}
}

func assertErr(t *testing.T, expected, actual error) {
	if expected != nil && actual != nil {
		assertEquals(t, expected.Error(), actual.Error())
	} else if expected != nil || actual != nil {
		t.Fatalf("Error %s is not as expected %s", actual, expected)
	}
}
