package firefisharchive_test

import (
	"testing"

	"encoding/json"

	"sourcecode.social/reiver/go-opt"

	"github.com/reiver/firefish2atom/firefisharchive"
)

func TestFile_unmarshalJSON(t *testing.T) {

	tests := []struct{
		JSON string
		Expected firefisharchive.File
	}{
		{
			JSON:
     `{
        "id": "9km5k6ctizy9ins5",
        "createdAt": "2023-10-09T07:45:18.894Z",
        "name": "Screenshot_20231009_004416_Telegram~2.jpg",
        "type": "image/jpeg",
        "md5": "1f477b4a5cb774836139eccdc123c6f5",
        "size": 126453,
        "isSensitive": false,
        "blurhash": "yM7_NoRhRNt3M_t8WCRwkFt9t8Roobf5tWa]V=V@oba$flIRt9j^RkxuRij[WGocaxR%t6R*j]oNWBj=kCaykCayoYWEkDt7Rkofay",
        "properties": {
          "width": 1080,
          "height": 865
        },
        "url": "https://firefish.lol/files/c0e33b7d-b612-45be-9eca-11d6883c594e",
        "thumbnailUrl": "https://firefish.lol/files/thumbnail-7bd86835-7601-426a-8ed3-498da98c5e30",
        "comment": "WIFE:\n\n\" https://www.google.ca/search?q=dec2+short+sleep&ie=UTF-8&oe=UTF-8&hl=en-ca&client=safari \"\n\n\" You think you might have this gene? \"\n\nME:\n\n\" Because I can function with a lot less sleep than you? 🙂 \"",
        "folderId": null,
        "folder": null,
        "userId": null,
        "user": null
      }`,
			Expected: firefisharchive.File{
				ID:           opt.Something("9km5k6ctizy9ins5"),
				CreatedAt:    opt.Something("2023-10-09T07:45:18.894Z"),
				Name:         opt.Something("Screenshot_20231009_004416_Telegram~2.jpg"),
				Type:         opt.Something("image/jpeg"),
				MD5:          opt.Something("1f477b4a5cb774836139eccdc123c6f5"),
				Size:         opt.Something(uint64(126453)),
				IsSensitive:  opt.Something(false),
				BlurHash:     opt.Something("yM7_NoRhRNt3M_t8WCRwkFt9t8Roobf5tWa]V=V@oba$flIRt9j^RkxuRij[WGocaxR%t6R*j]oNWBj=kCaykCayoYWEkDt7Rkofay"),
				Properties: struct{
					Width  opt.Optional[uint64] `json:"width"`
					Height opt.Optional[uint64] `json:"height"`
				}{
					Width:  opt.Something(uint64(1080)),
					Height: opt.Something(uint64(865)),
				},
				URL:          opt.Something("https://firefish.lol/files/c0e33b7d-b612-45be-9eca-11d6883c594e"),
				ThumbnailURL: opt.Something("https://firefish.lol/files/thumbnail-7bd86835-7601-426a-8ed3-498da98c5e30"),
				Comment:      opt.Something("WIFE:\n\n\" https://www.google.ca/search?q=dec2+short+sleep&ie=UTF-8&oe=UTF-8&hl=en-ca&client=safari \"\n\n\" You think you might have this gene? \"\n\nME:\n\n\" Because I can function with a lot less sleep than you? 🙂 \""),
			},
		},
	}

	for testNumber, test := range tests {

		var p []byte = []byte(test.JSON)

		var actual firefisharchive.File

		err := json.Unmarshal(p, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("JSON:\n%s", test.JSON)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual 'file' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("JSON:\n%s", test.JSON)
				continue
			}
		}
	}
}
