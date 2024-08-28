package app

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	iot "github.com/alibabacloud-go/iot-api-gateway/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/super1-chen/fxoss/logger"
	"github.com/super1-chen/fxoss/utils"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type operationCode int

const (
	stop operationCode = iota
	start
)

type SOServer struct {
	Host       string
	HTTPClient *http.Client
	logger     *log.Logger
}

func NewSOServer(verbose bool) (*SOServer, error) {
	// skip ssl verification.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	so := &SOServer{
		Host:       os.Getenv(sohostKey),
		HTTPClient: &http.Client{Timeout: time.Minute, Transport: tr},
		logger:     logger.Mylogger(verbose),
	}

	return so, nil
}

func (so *SOServer) StopSO(deviceName string) error {
	return so.operateSO(deviceName, stop)
}

func (so *SOServer) StartSO(deviceName string) error {
	return so.operateSO(deviceName, start)
}

func (so *SOServer) RestartSO(deviceName string) error {
	err := so.StopSO(deviceName)
	if err != nil {
		return err
	}
	utils.SuccessPrintln("stopped the smart outlet")
	time.Sleep(10 * time.Second)
	err = so.StartSO(deviceName)
	if err != nil {
		return err
	}
	utils.SuccessPrintln("started the smart outlet")
	time.Sleep(10 * time.Second)

	return nil
}

func (so *SOServer) operateSO(deviceName string, opCode operationCode) error {
	iotSecretKeyId := os.Getenv("FXOSS_IOT_KEY")
	iotSecretKeySecretKey := os.Getenv("FXOSS_IOT_KEY_SECRET")
	errorMsg := "restart smart outlet by aliyun iot api failed"
	//successMsg := "restart smart outlet by aliyun iot api successfully"
	productKeyList := []string{
		"a1Ug5Ju6PeZ", // PWE4GW0011~PWE5GL0001
		"a1ez4hPF43N", // <= PWE4GW0010
		"a1mjUUWlC16", // else
	}

	so.logger.Println("Creating AliIotApiClient")
	client, err := utils.CreateAliIotApiClient(iotSecretKeyId, iotSecretKeySecretKey)
	if err != nil {
		utils.ErrorPrintln(errorMsg, false)
		return fmt.Errorf("%s, %v", errorMsg, err)
	}
	if client == nil {
		return errors.New("client creation failed, client is nil")
	}

	for _, key := range productKeyList {
		params := map[string]interface{}{
			"items": map[string]operationCode{
				"PowerSwitch": opCode,
			},
			"productKey": key,
			"deviceName": deviceName,
			"iotId":      "",
		}
		req := new(iot.CommonParams).
			SetApiVer("1.0.2")
		body := new(iot.IoTApiRequest).
			SetParams(params).
			SetRequest(req)
		runtime := new(util.RuntimeOptions)

		so.logger.Println("Sending API request")
		resp, err := client.DoRequest(tea.String("/cloud/thing/properties/set"), tea.String("HTTPS"), tea.String("POST"), nil, body, runtime)
		if err != nil {
			return fmt.Errorf("API request failed, %v", err)
		}
		if resp == nil {
			return errors.New("API response is nil")
		}
		if resp.Body != nil {
			b, _ := io.ReadAll(resp.Body)
			var data = &soApiResponse{}
			err = json.Unmarshal(b, data)
			if err != nil {
				so.logger.Printf("failed to decode response data", false)
				return err
			}

			// 调用成功直接退出 for 循环
			if data.Code == 200 {
				break
			}

			// productKey 不对则尝试下一个
			if data.Code == 6100 {
				continue
			}

			if data.Code != 200 {
				return errors.New("ErrorCode: " + strconv.Itoa(data.Code) + " - " + data.LocalizedMsg)
			}
		} else {
			return errors.New("API response body is nil")
		}
	}

	return nil
}

func (so *SOServer) ShowSOList(option string) error {
	api := "/smartOutlet/list"
	errorMsg := "get smart outlet list from api failed"
	successMsg := "get smart outlet list from api successfully"
	data := new(soList)
	var soList []*soInfo

	var headers []string
	var content [][]string

	b, err := so.get(api)
	if err != nil {
		utils.ErrorPrintln(errorMsg, false)
		return fmt.Errorf("%s, %v", errorMsg, err)
	}

	if err = json.Unmarshal(b, &data); err != nil {
		so.logger.Printf("decode list failed %v", err)
		utils.ErrorPrintln("decode smart outlet list failed", false)
		return fmt.Errorf("decode smart outlet list failed, %v", err)
	}

	utils.SuccessPrintln(successMsg)
	if len(data.SO) == 0 {
		so.logger.Printf("decode list failed %v", err)
		utils.ColorPrintln("smart outlet list is empty", utils.Yellow)
		return nil
	}

	if option == "" {
		soList = data.SO
	} else {
		for _, so := range data.SO {
			if strings.Contains(so.ProbeSN, option) || strings.Contains(so.Customer, option) || strings.Contains(so.OutletSN, option) {
				soList = append(soList, so)
			}
		}
	}
	if len(soList) == 0 {
		utils.ColorPrintln("smart outlet list is empty", utils.Yellow)
		return nil
	}

	headers = []string{"#", "ProbeSN", "OutletSN", "OutletICCID", "Customer"}
	for index, so := range soList {
		index++
		content = append(content, []string{
			strconv.Itoa(index),
			so.ProbeSN,
			so.OutletSN,
			so.OutletICCID,
			so.Customer,
		})
	}
	utils.PrintTable(headers, content)
	return nil
}

func (so *SOServer) get(api string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", so.Host, api)
	req, err := http.NewRequest("GET", url, nil)
	so.logger.Printf("start request api %s", url)
	if err != nil {
		return nil, fmt.Errorf("create new get request %s failed %v", api, err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := so.HTTPClient.Do(req)

	if err != nil {
		so.logger.Printf("request get %s failed %v", api, err)
		return nil, fmt.Errorf("reqest get %s failed %v", api, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		so.logger.Printf("request get %s status %s", api, resp.Status)
		return nil, fmt.Errorf("request get %s status %s", api, resp.Status)
	}
	return ioutil.ReadAll(resp.Body)
}
