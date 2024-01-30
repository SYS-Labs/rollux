package sources

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

// JSONMarshalerV2 is used for marshalling requests to newer Syscoin Type RPC interfaces
type JSONMarshalerV2 struct{}

var addressLabel string = "podalabel"

// Marshal converts struct passed by parameter to JSON
func (JSONMarshalerV2) Marshal(v interface{}) ([]byte, error) {
	d, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// SyscoinRPC is an interface to JSON-RPC syscoind service.
type SyscoinRPC struct {
	client       http.Client
	rpcURL       string
	user         string
	password     string
	podaurl      string
	RPCMarshaler JSONMarshalerV2
}
type SyscoinClient struct {
	client *SyscoinRPC
}

func NewSyscoinClient(podaurl string) (*SyscoinClient, error) {
	transport := &http.Transport{
		Dial:                (&net.Dialer{KeepAlive: 600 * time.Second}).Dial,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100, // necessary to not to deplete ports
	}
	s := &SyscoinRPC{
		client:       http.Client{Timeout: time.Duration(600) * time.Second, Transport: transport},
		rpcURL:       "http://54.187.137.60:8370",
		user:         "u",
		password:     "p",
		podaurl:      podaurl,
		RPCMarshaler: JSONMarshalerV2{},
	}
	client := SyscoinClient{s}
	// if its empty it means we are in batcher mode which needs SYS gas for PoDA transactions
	if podaurl == "" {
		log.Info("NewSyscoinClient loading wallet...")
		walletName := "wallet"
		var err error = errors.New("")
		for err != nil {
			err = client.CreateOrLoadWallet(walletName)
			if err != nil {
				log.Info("NewSyscoinClient", "err", err)
			}
			time.Sleep(1 * time.Second)
		}
		client.client.rpcURL += "/wallet/" + walletName
		log.Info("NewSyscoinClient checking balance...")
		balance, err := client.GetBalance()
		if err != nil {
			return &client, err
		}
		var address string
		if balance <= 0.0 {
			log.Info("NewSyscoinClient balance is empty, fetching funding address", "label", addressLabel)
			address, err = client.FetchAddressByLabel(addressLabel)
			if address == "" {
				log.Info("NewSyscoinClient label does not exist, creating new funding address")
				address, err = client.GetNewAddress(addressLabel)
				if err != nil {
					return &client, err
				}
			}
			log.Info("NewSyscoinClient please fund SYS", "address", address)
		}
		for balance <= 0.0 {
			balance, err = client.GetBalance()
			if err != nil {
				return &client, err
			}
			time.Sleep(10 * time.Second)
			log.Info("NewSyscoinClient waiting for funds at funding destination", "address", address)
		}
	}

	log.Info("NewSyscoinClient loaded!")
	return &client, nil
}

// RPCError defines rpc error returned by backend
type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
func safeDecodeResponse(body io.ReadCloser, res interface{}) (err error) {
	var data []byte
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			if len(data) > 0 && len(data) < 2048 {
				err = errors.New(fmt.Sprintf("Error %v", string(data)))
			} else {
				err = errors.New("Internal error")
			}
		}
	}()
	data, err = ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &res)
}

// Call calls Backend RPC interface, using RPCMarshaler interface to marshall the request
func (s *SyscoinClient) Call(req interface{}, res interface{}) error {
	httpData, err := s.client.RPCMarshaler.Marshal(req)
	if err != nil {
		return err
	}
	httpReq, err := http.NewRequest("POST", s.client.rpcURL, bytes.NewBuffer(httpData))
	if err != nil {
		return err
	}
	httpReq.SetBasicAuth(s.client.user, s.client.password)
	httpRes, err := s.client.client.Do(httpReq)
	// in some cases the httpRes can contain data even if it returns error
	// see http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/
	if httpRes != nil {
		defer httpRes.Body.Close()
	}
	if err != nil {
		return err
	}
	// if server returns HTTP error code it might not return json with response
	// handle both cases
	if httpRes.StatusCode != 200 {
		err = safeDecodeResponse(httpRes.Body, &res)
		if err != nil {
			return errors.New(fmt.Sprintf("Error %v %v", httpRes.Status, err))
		}
		return nil
	}
	return safeDecodeResponse(httpRes.Body, &res)
}

func (s *SyscoinClient) CreateBlob(data []byte) (common.Hash, error) {
	type ResCreateBlob struct {
		Error  *RPCError `json:"error"`
		Result struct {
			VH string `json:"versionhash"`
		} `json:"result"`
	}

	res := ResCreateBlob{}
	type CmdCreateBlob struct {
		Method string `json:"method"`
		Params struct {
			Data string `json:"data"`
		} `json:"params"`
	}
	req := CmdCreateBlob{Method: "syscoincreatenevmblob"}
	req.Params.Data = hex.EncodeToString(data)
	err := s.Call(&req, &res)
	if err != nil {
		return common.Hash{}, err
	}
	if res.Error != nil {
		return common.Hash{}, res.Error
	}
	return common.HexToHash(res.Result.VH), err
}
func (s *SyscoinClient) CreateOrLoadWallet(walletName string) error {
	type ResCreateWallet struct {
		Error  *RPCError `json:"error"`
		Result struct {
			Warning string `json:"warning"`
		} `json:"result"`
	}

	res := ResCreateWallet{}
	type CmdCreateWallet struct {
		Method string `json:"method"`
		Params struct {
			WalletName string `json:"wallet_name"`
		} `json:"params"`
	}
	req := CmdCreateWallet{Method: "createwallet"}
	req.Params.WalletName = walletName
	err := s.Call(&req, &res)
	if err != nil {
		return err
	}
	// might actually be created already so just load it
	if res.Error != nil {
		log.Info("CreateOrLoadWallet wallet exists, loading it...")
		type ResLoadWallet struct {
			Error  *RPCError `json:"error"`
			Result struct {
				Warning string `json:"warning"`
			} `json:"result"`
		}

		res := ResLoadWallet{}
		type CmdLoadWallet struct {
			Method string `json:"method"`
			Params struct {
				WalletName string `json:"filename"`
			} `json:"params"`
		}
		req := CmdLoadWallet{Method: "loadwallet"}
		req.Params.WalletName = walletName
		err = s.Call(&req, &res)
		if err != nil {
			if strings.Contains(err.Error(), "is already loaded") {
				log.Info("CreateOrLoadWallet wallet already loaded...")
				return nil
			}
			return err
		}
		if res.Error != nil {
			if strings.Contains(res.Error.Error(), "is already loaded") {
				log.Info("CreateOrLoadWallet wallet already loaded...")
				return nil
			}
			return res.Error
		}
		return nil
	}
	if len(res.Result.Warning) > 0 {
		log.Info("CreateOrLoadWallet", "warning", res.Result.Warning)
	}
	return nil
}
func (s *SyscoinClient) GetBalance() (float64, error) {
	type ResGetBalance struct {
		Error   *RPCError `json:"error"`
		Balance float64   `json:"result"`
	}

	res := ResGetBalance{}
	type CmdGetBalance struct {
		Method string `json:"method"`
		Params struct {
		} `json:"params"`
	}
	req := CmdGetBalance{Method: "getbalance"}
	err := s.Call(&req, &res)
	if err != nil {
		return 0, err
	}
	if res.Error != nil {
		return 0, res.Error
	}
	return res.Balance, nil
}
func (s *SyscoinClient) GetNewAddress(addresslabel string) (string, error) {
	type ResGetAddress struct {
		Error   *RPCError `json:"error"`
		Address string    `json:"result"`
	}

	res := ResGetAddress{}
	type CmdGetAddress struct {
		Method string `json:"method"`
		Params struct {
			Label string `json:"label"`
		} `json:"params"`
	}
	req := CmdGetAddress{Method: "getnewaddress"}
	req.Params.Label = addresslabel
	err := s.Call(&req, &res)
	if err != nil {
		return "", err
	}
	if res.Error != nil {
		return "", res.Error
	}
	return res.Address, nil
}
func (s *SyscoinClient) FetchAddressByLabel(addresslabel string) (string, error) {
	type GetAddressesByLabelRespElement struct {
		// Purpose of address ("send" for sending address, "receive" for receiving address)
		Purpose string `json:"purpose"`
	}
	type ResGetAddress struct {
		Error  *RPCError                                 `json:"error"`
		Result map[string]GetAddressesByLabelRespElement `json:"result"`
	}

	res := ResGetAddress{}
	type CmdGetAddress struct {
		Method string `json:"method"`
		Params struct {
			Label string `json:"label"`
		} `json:"params"`
	}
	req := CmdGetAddress{Method: "getaddressesbylabel"}
	req.Params.Label = addresslabel
	err := s.Call(&req, &res)
	if err != nil {
		return "", err
	}
	if res.Error != nil {
		return "", res.Error
	}
	for key, _ := range res.Result {
		return key, nil
	}
	return "", nil
}

// SYSCOIN used to get blob confirmation by checking block number then tx receipt below to get block height of blob confirmation
func (s *SyscoinClient) BlockNumber(ctx context.Context) (uint64, error) {
	type ResGetBlockNumber struct {
		Error       *RPCError `json:"error"`
		BlockNumber uint64    `json:"result"`
	}
	res := ResGetBlockNumber{}
	type CmdGetBlockNumber struct {
		Method string `json:"method"`
		Params struct {
		} `json:"params"`
	}
	req := CmdGetBlockNumber{Method: "getblockcount"}
	err := s.Call(&req, &res)
	if err != nil {
		return 0, err
	}
	if res.Error != nil {
		return 0, res.Error
	}
	return res.BlockNumber, err
}

// SYSCOIN used to get blob receipt
func (s *SyscoinClient) TransactionReceipt(ctx context.Context, vh common.Hash) (*types.Receipt, error) {
	type ResGetBlobReceipt struct {
		Error  *RPCError `json:"error"`
		Result struct {
			MPT int64 `json:"mpt"`
		} `json:"result"`
	}
	res := ResGetBlobReceipt{}
	type CmdGetBlobReceipt struct {
		Method string `json:"method"`
		Params struct {
			TXID string `json:"versionhash_or_txid"`
		} `json:"params"`
	}
	req := CmdGetBlobReceipt{Method: "getnevmblobdata"}
	req.Params.TXID = vh.String()[2:]
	err := s.Call(&req, &res)
	if err != nil {
		return nil, err
	}
	if res.Error != nil {
		return nil, res.Error
	}
	receipt := types.Receipt{}
	if res.Result.MPT > 0 {
		// store VH in TxHash used by driver to put into the batch
		receipt = types.Receipt{
			TxHash: vh,
			// store MPT in BlockNumber to be used in caller
			BlockNumber: big.NewInt(res.Result.MPT),
			Status:      types.ReceiptStatusSuccessful,
		}
	}
	return &receipt, err
}

func (s *SyscoinClient) GetBlobFromRPC(vh common.Hash) ([]byte, error) {
	type ResGetBlobData struct {
		Error  *RPCError `json:"error"`
		Result struct {
			Data string `json:"data"`
		} `json:"result"`
	}
	res := ResGetBlobData{}
	type CmdGetBlobData struct {
		Method string `json:"method"`
		Params struct {
			VersionHash string `json:"versionhash_or_txid"`
			Verbose     bool   `json:"getdata"`
		} `json:"params"`
	}
	req := CmdGetBlobData{Method: "getnevmblobdata"}
	req.Params.VersionHash = vh.String()[2:]
	req.Params.Verbose = true
	err := s.Call(&req, &res)
	if err != nil {
		return nil, err
	}
	if res.Error != nil {
		return nil, res.Error
	}
	data, err := hex.DecodeString(res.Result.Data)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *SyscoinClient) GetBlobFromCloud(vh common.Hash) ([]byte, error) {
	url := s.client.podaurl + vh.String()[2:]
	var res *http.Response
	var err error
	// try 4 times incase of timeout or reset/hanging socket with 5+i second expiry each attempt
	for i := 0; i < 4; i++ {
		client := http.Client{
			Timeout: (5 + time.Duration(i)) * time.Second,
		}
		res, err = client.Get(url)
		if err != nil {
			continue
		} else {
			err = nil
			break
		}
	}
	if err != nil {
		return nil, err
	}
	defer res.Body.Close() // we need to close the connection
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	txBytes, err := hex.DecodeString(string(body))
	if err != nil {
		return nil, err
	}
	return txBytes, nil
}
