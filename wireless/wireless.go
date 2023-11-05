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

func getValue(s string) string {
	array := strings.Split(s, "=")
	return array[1]
}

func arrayToTemplate(lines []string) string {
	var s string
	for _, line := range lines {
		s += line + "\r\n"
	}
	return s
}

type Template interface {
	ToString() string
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

type Class struct {
	SSID     string `json:"ssid"`
	Enable   bool   `json:"enable"`
	Password string `json:"password"`
}

func NewClass(data []byte) *Class {
	class := new(Class)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		s := strings.Split(line, "=")
		switch s[0] {
		case "SSID":
			class.SSID = s[1]
		case "enable":
			class.Enable = strIntToBool(s[1])
		case "X_TP_PreSharedKey":
			class.Password = s[1]
		}
	}
	return class
}

func Get(s *Slice) (*Class, error) {
	data, err := request("5", s.ToString(), true)
	if err != nil {
		return nil, err
	}
	return NewClass(data), nil
}

func Put(m *Map) error {
	_, err := request("2", m.ToString(), true)
	return err
}
