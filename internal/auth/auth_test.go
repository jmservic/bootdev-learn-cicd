package auth

import (
	"net/http"
	"testing"
	//"errors"
)

func Test_GetApiKey(t *testing.T) {
	type testcase struct {
		headers http.Header
		authToken string
		err error
	}
	testcases := []testcase{
		{
			headers: http.Header{
				"Authorization": []string{"ApiKey abc88efgh"},
			},
			authToken: "abc88efgh",
			err: nil,
		},
		{
			headers: http.Header{},
			authToken: "",
			err: ErrNoAuthHeaderIncluded,
		},
		{
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			authToken: "",
			err: ErrMalformedAuthHeader,

		},
		{
			headers: http.Header{
				"Authorization": []string{"Apey abc88efgh"},
			},
			authToken: "",
			err: ErrMalformedAuthHeader,
		},
	}

	for _, tc := range testcases {
		authToken, err := GetAPIKey(tc.headers)
		if err != tc.err {
			t.Errorf("Received err (%v) doesn't match (%v)", 
				err, tc.err) 
		}
		if authToken != tc.authToken {
			t.Errorf("Received auth token %s doesn't equal %s",
				authToken,
				tc.authToken)
		}
	}

}
