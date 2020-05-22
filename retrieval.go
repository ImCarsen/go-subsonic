package subsonic

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)

// Stream returns the contents of a song, optionally transcoded, from the server.
//
// Optional Parameters:
//   maxBitRate:             (Since 1.2.0) If specified, the server will attempt to limit the bitrate to this value, in kilobits per second. If set to zero, no limit is imposed.
//   format:                 (Since 1.6.0) Specifies the preferred target format (e.g., "mp3" or "flv") in case there are multiple applicable transcodings.  Starting with 1.9.0 you can use the special value "raw" to disable transcoding.
//   timeOffset:             Only applicable to video streaming. If specified, start streaming at the given offset (in seconds) into the video. Typically used to implement video skipping.
//   size:                   (Since 1.6.0) Only applicable to video streaming. Requested video size specified as WxH, for instance "640x480".
//   estimateContentLength:  (Since 1.8.0). If set to "true", the Content-Length HTTP header will be set to an estimated value for transcoded or downsampled media.
//   converted:              (Since 1.14.0) Only applicable to video streaming. Subsonic can optimize videos for streaming by converting them to MP4. If a conversion exists for the video in question, then setting this parameter to "true" will cause the converted video to be returned instead of the original.
func (s *Client) Stream(id string, parameters map[string]string) ([]byte, error) {
	params := make(map[string]string)
	params["id"] = id
	for k, v := range parameters {
		params[k] = v
	}
	response, err := s.Request("GET", "stream", params)
	if err != nil {
		return nil, err
	}
	contentType := response.Header.Get("Content-Type")
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(contentType, "text/xml") {
		// An error was returned
		resp := Response{}
		err = xml.Unmarshal(responseBody, &resp)
		if err != nil {
			return nil, err
		}
		if resp.Error != nil {
			err = fmt.Errorf("Error #%d: %s\n", resp.Error.Code, resp.Error.Message)
		} else {
			err = fmt.Errorf("An error occurred: %#v\n", resp)
		}
		return nil, err
	}
	return responseBody, nil
}
