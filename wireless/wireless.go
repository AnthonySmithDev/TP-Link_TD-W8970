package wireless

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func authorization() string {
	user := "Anthony"
	pass := "Smith"
	bytes := []byte(fmt.Sprint(user, ":", pass))
	encode := base64.StdEncoding.EncodeToString(bytes)
	return fmt.Sprintf("Authorization=Basic %s", encode)
}

func request(params string, data string, debug bool) ([]byte, error) {
	req, err := http.NewRequest("POST", "http://192.168.0.1/cgi?"+params, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Origin", "http://192.168.0.1")
	req.Header.Set("Referer", "http://192.168.0.1")
	req.Header.Set("Cookie", authorization())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 500 {
		return nil, errors.New("Internal Server Error")
	}
	if debug {
		fmt.Println("Response Body:")
		fmt.Println(string(body))
	}
	return body, nil
}

func boolToInt(v bool) uint8 {
	if v {
		return 1
	}
	return 0
}

func strIntToBool(v string) bool {
	if v == "1" {
		return true
	}
	return false
}

type Slice struct {
	array []string
}

func NewSlice() *Slice {
	return &Slice{
		array: []string{},
	}
}

func (s *Slice) SetSSID() *Slice {
	s.array = append(s.array, "SSID")
	return s
}

func (s *Slice) SetEnable() *Slice {
	s.array = append(s.array, "Enable")
	return s
}

func (s *Slice) SetPassword() *Slice {
	s.array = append(s.array, "X_TP_PreSharedKey")
	return s
}

func (s *Slice) ToString() string {
	head := fmt.Sprintf("[LAN_WLAN#0,0,0,0,0,0#0,0,0,0,0,0]0,%d", len(s.array))
	lines := append([]string{head}, s.array...)
	var data string
	for _, line := range lines {
		data += line + "\r\n"
	}
	return data
}

type Map struct {
	array []string
}

func NewMap() *Map {
	return &Map{
		array: []string{},
	}
}

func (s *Map) SetSSID(ssid string) *Map {
	s.array = append(s.array, fmt.Sprintf("SSID=%s", ssid))
	return s
}

func (s *Map) SetEnable(enable bool) *Map {
	s.array = append(s.array, fmt.Sprintf("Enable=%d", boolToInt(enable)))
	return s
}

func (s *Map) SetPassword(password string) *Map {
	s.array = append(s.array, fmt.Sprintf("X_TP_PreSharedKey=%s", password))
	return s
}

func (s *Map) ToString() string {
	head := fmt.Sprintf("[LAN_WLAN#1,1,0,0,0,0#0,0,0,0,0,0]0,%d", len(s.array))
	lines := append([]string{head}, s.array...)
	var data string
	for _, line := range lines {
		data += line + "\r\n"
	}
	return data
}

type Info struct {
	SSID     string `json:"ssid"`
	Enable   bool   `json:"enable"`
	Password string `json:"password"`
}

func NewInfo(data []byte) *Info {
	info := new(Info)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		s := strings.Split(line, "=")
		switch s[0] {
		case "SSID":
			info.SSID = s[1]
		case "enable":
			info.Enable = strIntToBool(s[1])
		case "X_TP_PreSharedKey":
			info.Password = s[1]
		}
	}
	return info
}

func Get(s *Slice) (*Info, error) {
	data, err := request("5", s.ToString(), false)
	if err != nil {
		return nil, err
	}
	return NewInfo(data), nil
}

func Put(m *Map) error {
	_, err := request("2", m.ToString(), false)
	return err
}
