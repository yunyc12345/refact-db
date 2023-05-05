// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ILayerZeroUltraLightNodeV2ApplicationConfiguration is an auto generated low-level Go binding around an user-defined struct.
type ILayerZeroUltraLightNodeV2ApplicationConfiguration struct {
	InboundProofLibraryVersion uint16
	InboundBlockConfirmations  uint64
	Relayer                    common.Address
	OutboundProofType          uint16
	OutboundBlockConfirmations uint64
	Oracle                     common.Address
}

// UltraLightNodeMetaData contains all meta data concerning the UltraLightNode contract.
var UltraLightNodeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_endpoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nonceContract\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_localChainId\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"}],\"name\":\"AddInboundProofLibraryForChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userApplication\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"configType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newConfig\",\"type\":\"bytes\"}],\"name\":\"AppConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"proofType\",\"type\":\"uint16\"}],\"name\":\"EnableSupportedOutboundProof\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"srcChainId\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lookupHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockData\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"confirmations\",\"type\":\"uint256\"}],\"name\":\"HashReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"srcChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srcAddress\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dstAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"payloadHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidDst\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"Packet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"srcChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srcAddress\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dstAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"payloadHash\",\"type\":\"bytes32\"}],\"name\":\"PacketReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"adapterParams\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"outboundProofType\",\"type\":\"uint16\"}],\"name\":\"RelayerParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"SetChainAddressSize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"proofType\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"adapterParams\",\"type\":\"bytes\"}],\"name\":\"SetDefaultAdapterParamsForChainId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"inboundProofLib\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"inboundBlockConfirm\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"outboundProofType\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"outboundBlockConfirm\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"SetDefaultConfigForChainId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"SetLayerZeroToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"uln\",\"type\":\"bytes32\"}],\"name\":\"SetRemoteUln\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasuryAddress\",\"type\":\"address\"}],\"name\":\"SetTreasury\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawNative\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawZRO\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CONFIG_TYPE_INBOUND_BLOCK_CONFIRMATIONS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONFIG_TYPE_INBOUND_PROOF_LIBRARY_VERSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONFIG_TYPE_ORACLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONFIG_TYPE_OUTBOUND_BLOCK_CONFIRMATIONS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONFIG_TYPE_OUTBOUND_PROOF_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONFIG_TYPE_RELAYER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"accruedNativeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_library\",\"type\":\"address\"}],\"name\":\"addInboundProofLibraryForChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"appConfig\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"inboundProofLibraryVersion\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"inboundBlockConfirmations\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"outboundProofType\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"outboundBlockConfirmations\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"chainAddressSizeMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"defaultAdapterParams\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"defaultAppConfig\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"inboundProofLibraryVersion\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"inboundBlockConfirmations\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"outboundProofType\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"outboundBlockConfirmations\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_proofType\",\"type\":\"uint16\"}],\"name\":\"enableSupportedOutboundProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_ua\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"_payInZRO\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"}],\"name\":\"estimateFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zroFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_ua\",\"type\":\"address\"}],\"name\":\"getAppConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"inboundProofLibraryVersion\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"inboundBlockConfirmations\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"outboundProofType\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"outboundBlockConfirmations\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"internalType\":\"structILayerZeroUltraLightNodeV2.ApplicationConfiguration\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_ua\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_configType\",\"type\":\"uint256\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_path\",\"type\":\"bytes\"}],\"name\":\"getOutboundNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hashLookup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"inboundProofLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"layerZeroToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localChainId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"maxInboundProofLibrary\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nativeFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonceContract\",\"outputs\":[{\"internalType\":\"contractNonceContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ua\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_path\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"_refundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zroPaymentAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"}],\"name\":\"setChainAddressSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_ua\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_configType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_config\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_proofType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_adapterParams\",\"type\":\"bytes\"}],\"name\":\"setDefaultAdapterParamsForChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_chainId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_inboundProofLibraryVersion\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"_inboundBlockConfirmations\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_outboundProofType\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"_outboundBlockConfirmations\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setDefaultConfigForChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_layerZeroToken\",\"type\":\"address\"}],\"name\":\"setLayerZeroToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_remoteChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"_remoteUln\",\"type\":\"bytes32\"}],\"name\":\"setRemoteUln\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"supportedOutboundProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryContract\",\"outputs\":[{\"internalType\":\"contractILayerZeroTreasury\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryZROFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"ulnLookup\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"_lookupHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_confirmations\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_blockData\",\"type\":\"bytes32\"}],\"name\":\"updateHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_dstAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_lookupHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_blockData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_transactionProof\",\"type\":\"bytes\"}],\"name\":\"validateTransactionProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawZRO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UltraLightNodeABI is the input ABI used to generate the binding from.
// Deprecated: Use UltraLightNodeMetaData.ABI instead.
var UltraLightNodeABI = UltraLightNodeMetaData.ABI

// UltraLightNode is an auto generated Go binding around an Ethereum contract.
type UltraLightNode struct {
	UltraLightNodeCaller     // Read-only binding to the contract
	UltraLightNodeTransactor // Write-only binding to the contract
	UltraLightNodeFilterer   // Log filterer for contract events
}

// UltraLightNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type UltraLightNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UltraLightNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UltraLightNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UltraLightNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UltraLightNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UltraLightNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UltraLightNodeSession struct {
	Contract     *UltraLightNode   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UltraLightNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UltraLightNodeCallerSession struct {
	Contract *UltraLightNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UltraLightNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UltraLightNodeTransactorSession struct {
	Contract     *UltraLightNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UltraLightNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type UltraLightNodeRaw struct {
	Contract *UltraLightNode // Generic contract binding to access the raw methods on
}

// UltraLightNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UltraLightNodeCallerRaw struct {
	Contract *UltraLightNodeCaller // Generic read-only contract binding to access the raw methods on
}

// UltraLightNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UltraLightNodeTransactorRaw struct {
	Contract *UltraLightNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUltraLightNode creates a new instance of UltraLightNode, bound to a specific deployed contract.
func NewUltraLightNode(address common.Address, backend bind.ContractBackend) (*UltraLightNode, error) {
	contract, err := bindUltraLightNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UltraLightNode{UltraLightNodeCaller: UltraLightNodeCaller{contract: contract}, UltraLightNodeTransactor: UltraLightNodeTransactor{contract: contract}, UltraLightNodeFilterer: UltraLightNodeFilterer{contract: contract}}, nil
}

// NewUltraLightNodeCaller creates a new read-only instance of UltraLightNode, bound to a specific deployed contract.
func NewUltraLightNodeCaller(address common.Address, caller bind.ContractCaller) (*UltraLightNodeCaller, error) {
	contract, err := bindUltraLightNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UltraLightNodeCaller{contract: contract}, nil
}

// NewUltraLightNodeTransactor creates a new write-only instance of UltraLightNode, bound to a specific deployed contract.
func NewUltraLightNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*UltraLightNodeTransactor, error) {
	contract, err := bindUltraLightNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UltraLightNodeTransactor{contract: contract}, nil
}

// NewUltraLightNodeFilterer creates a new log filterer instance of UltraLightNode, bound to a specific deployed contract.
func NewUltraLightNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*UltraLightNodeFilterer, error) {
	contract, err := bindUltraLightNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UltraLightNodeFilterer{contract: contract}, nil
}

// bindUltraLightNode binds a generic wrapper to an already deployed contract.
func bindUltraLightNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UltraLightNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UltraLightNode *UltraLightNodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UltraLightNode.Contract.UltraLightNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UltraLightNode *UltraLightNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UltraLightNode.Contract.UltraLightNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UltraLightNode *UltraLightNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UltraLightNode.Contract.UltraLightNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UltraLightNode *UltraLightNodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UltraLightNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UltraLightNode *UltraLightNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UltraLightNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UltraLightNode *UltraLightNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UltraLightNode.Contract.contract.Transact(opts, method, params...)
}

// CONFIGTYPEINBOUNDBLOCKCONFIRMATIONS is a free data retrieval call binding the contract method 0x49148c37.
//
// Solidity: function CONFIG_TYPE_INBOUND_BLOCK_CONFIRMATIONS() view returns(uint256)
func (_UltraLightNode *UltraLightNodeCaller) CONFIGTYPEINBOUNDBLOCKCONFIRMATIONS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "CONFIG_TYPE_INBOUND_BLOCK_CONFIRMATIONS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CONFIGTYPEINBOUNDBLOCKCONFIRMATIONS is a free data retrieval call binding the contract method 0x49148c37.
//
// Solidity: function CONFIG_TYPE_INBOUND_BLOCK_CONFIRMATIONS() view returns(uint256)
func (_UltraLightNode *UltraLightNodeSession) CONFIGTYPEINBOUNDBLOCKCONFIRMATIONS() (*big.Int, error) {
	return _UltraLightNode.Contract.CONFIGTYPEINBOUNDBLOCKCONFIRMATIONS(&_UltraLightNode.CallOpts)
}

// CONFIGTYPEINBOUNDBLOCKCONFIRMATIONS is a free data retrieval call binding the contract method 0x49148c37.
//
// Solidity: function CONFIG_TYPE_INBOUND_BLOCK_CONFIRMATIONS() view returns(uint256)
func (_UltraLightNode *UltraLightNodeCallerSession) CONFIGTYPEINBOUNDBLOCKCONFIRMATIONS() (*big.Int, error) {
	return _UltraLightNode.Contract.CONFIGTYPEINBOUNDBLOCKCONFIRMATIONS(&_UltraLightNode.CallOpts)
}

// CONFIGTYPEINBOUNDPROOFLIBRARYVERSION is a free data retrieval call binding the contract method 0xb77d22ad.
//
// Solidity: function CONFIG_TYPE_INBOUND_PROOF_LIBRARY_VERSION() view returns(uint256)
func (_UltraLightNode *UltraLightNodeCaller) CONFIGTYPEINBOUNDPROOFLIBRARYVERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "CONFIG_TYPE_INBOUND_PROOF_LIBRARY_VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CONFIGTYPEINBOUNDPROOFLIBRARYVERSION is a free data retrieval call binding the contract method 0xb77d22ad.
//
// Solidity: function CONFIG_TYPE_INBOUND_PROOF_LIBRARY_VERSION() view returns(uint256)
func (_UltraLightNode *UltraLightNodeSession) CONFIGTYPEINBOUNDPROOFLIBRARYVERSION() (*big.Int, error) {
	return _UltraLightNode.Contract.CONFIGTYPEINBOUNDPROOFLIBRARYVERSION(&_UltraLightNode.CallOpts)
}

// CONFIGTYPEINBOUNDPROOFLIBRARYVERSION is a free data retrieval call binding the contract method 0xb77d22ad.
//
// Solidity: function CONFIG_TYPE_INBOUND_PROOF_LIBRARY_VERSION() view returns(uint256)
func (_UltraLightNode *UltraLightNodeCallerSession) CONFIGTYPEINBOUNDPROOFLIBRARYVERSION() (*big.Int, error) {
	return _UltraLightNode.Contract.CONFIGTYPEINBOUNDPROOFLIBRARYVERSION(&_UltraLightNode.CallOpts)
}

// CONFIGTYPEORACLE is a free data retrieval call binding the contract method 0x31bd2430.
//
// Solidity: function CONFIG_TYPE_ORACLE() view returns(uint256)
func (_UltraLightNode *UltraLightNodeCaller) CONFIGTYPEORACLE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "CONFIG_TYPE_ORACLE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CONFIGTYPEORACLE is a free data retrieval call binding the contract method 0x31bd2430.
//
// Solidity: function CONFIG_TYPE_ORACLE() view returns(uint256)
func (_UltraLightNode *UltraLightNodeSession) CONFIGTYPEORACLE() (*big.Int, error) {
	return _UltraLightNode.Contract.CONFIGTYPEORACLE(&_UltraLightNode.CallOpts)
}

// CONFIGTYPEORACLE is a free data retrieval call binding the contract method 0x31bd2430.
//
// Solidity: function CONFIG_TYPE_ORACLE() view returns(uint256)
func (_UltraLightNode *UltraLightNodeCallerSession) CONFIGTYPEORACLE() (*big.Int, error) {
	return _UltraLightNode.Contract.CONFIGTYPEORACLE(&_UltraLightNode.CallOpts)
}

// CONFIGTYPEOUTBOUNDBLOCKCONFIRMATIONS is a free data retrieval call binding the contract method 0x904d3b8d.
//
// Solidity: function CONFIG_TYPE_OUTBOUND_BLOCK_CONFIRMATIONS() view returns(uint256)
func (_UltraLightNode *UltraLightNodeCaller) CONFIGTYPEOUTBOUNDBLOCKCONFIRMATIONS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "CONFIG_TYPE_OUTBOUND_BLOCK_CONFIRMATIONS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CONFIGTYPEOUTBOUNDBLOCKCONFIRMATIONS is a free data retrieval call binding the contract method 0x904d3b8d.
//
// Solidity: function CONFIG_TYPE_OUTBOUND_BLOCK_CONFIRMATIONS() view returns(uint256)
func (_UltraLightNode *UltraLightNodeSession) CONFIGTYPEOUTBOUNDBLOCKCONFIRMATIONS() (*big.Int, error) {
	return _UltraLightNode.Contract.CONFIGTYPEOUTBOUNDBLOCKCONFIRMATIONS(&_UltraLightNode.CallOpts)
}

// CONFIGTYPEOUTBOUNDBLOCKCONFIRMATIONS is a free data retrieval call binding the contract method 0x904d3b8d.
//
// Solidity: function CONFIG_TYPE_OUTBOUND_BLOCK_CONFIRMATIONS() view returns(uint256)
func (_UltraLightNode *UltraLightNodeCallerSession) CONFIGTYPEOUTBOUNDBLOCKCONFIRMATIONS() (*big.Int, error) {
	return _UltraLightNode.Contract.CONFIGTYPEOUTBOUNDBLOCKCONFIRMATIONS(&_UltraLightNode.CallOpts)
}

// CONFIGTYPEOUTBOUNDPROOFTYPE is a free data retrieval call binding the contract method 0x8140666e.
//
// Solidity: function CONFIG_TYPE_OUTBOUND_PROOF_TYPE() view returns(uint256)
func (_UltraLightNode *UltraLightNodeCaller) CONFIGTYPEOUTBOUNDPROOFTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "CONFIG_TYPE_OUTBOUND_PROOF_TYPE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CONFIGTYPEOUTBOUNDPROOFTYPE is a free data retrieval call binding the contract method 0x8140666e.
//
// Solidity: function CONFIG_TYPE_OUTBOUND_PROOF_TYPE() view returns(uint256)
func (_UltraLightNode *UltraLightNodeSession) CONFIGTYPEOUTBOUNDPROOFTYPE() (*big.Int, error) {
	return _UltraLightNode.Contract.CONFIGTYPEOUTBOUNDPROOFTYPE(&_UltraLightNode.CallOpts)
}

// CONFIGTYPEOUTBOUNDPROOFTYPE is a free data retrieval call binding the contract method 0x8140666e.
//
// Solidity: function CONFIG_TYPE_OUTBOUND_PROOF_TYPE() view returns(uint256)
func (_UltraLightNode *UltraLightNodeCallerSession) CONFIGTYPEOUTBOUNDPROOFTYPE() (*big.Int, error) {
	return _UltraLightNode.Contract.CONFIGTYPEOUTBOUNDPROOFTYPE(&_UltraLightNode.CallOpts)
}

// CONFIGTYPERELAYER is a free data retrieval call binding the contract method 0x2cfacb06.
//
// Solidity: function CONFIG_TYPE_RELAYER() view returns(uint256)
func (_UltraLightNode *UltraLightNodeCaller) CONFIGTYPERELAYER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "CONFIG_TYPE_RELAYER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CONFIGTYPERELAYER is a free data retrieval call binding the contract method 0x2cfacb06.
//
// Solidity: function CONFIG_TYPE_RELAYER() view returns(uint256)
func (_UltraLightNode *UltraLightNodeSession) CONFIGTYPERELAYER() (*big.Int, error) {
	return _UltraLightNode.Contract.CONFIGTYPERELAYER(&_UltraLightNode.CallOpts)
}

// CONFIGTYPERELAYER is a free data retrieval call binding the contract method 0x2cfacb06.
//
// Solidity: function CONFIG_TYPE_RELAYER() view returns(uint256)
func (_UltraLightNode *UltraLightNodeCallerSession) CONFIGTYPERELAYER() (*big.Int, error) {
	return _UltraLightNode.Contract.CONFIGTYPERELAYER(&_UltraLightNode.CallOpts)
}

// AccruedNativeFee is a free data retrieval call binding the contract method 0x69412bfa.
//
// Solidity: function accruedNativeFee(address _address) view returns(uint256)
func (_UltraLightNode *UltraLightNodeCaller) AccruedNativeFee(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "accruedNativeFee", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedNativeFee is a free data retrieval call binding the contract method 0x69412bfa.
//
// Solidity: function accruedNativeFee(address _address) view returns(uint256)
func (_UltraLightNode *UltraLightNodeSession) AccruedNativeFee(_address common.Address) (*big.Int, error) {
	return _UltraLightNode.Contract.AccruedNativeFee(&_UltraLightNode.CallOpts, _address)
}

// AccruedNativeFee is a free data retrieval call binding the contract method 0x69412bfa.
//
// Solidity: function accruedNativeFee(address _address) view returns(uint256)
func (_UltraLightNode *UltraLightNodeCallerSession) AccruedNativeFee(_address common.Address) (*big.Int, error) {
	return _UltraLightNode.Contract.AccruedNativeFee(&_UltraLightNode.CallOpts, _address)
}

// AppConfig is a free data retrieval call binding the contract method 0xddfdef5a.
//
// Solidity: function appConfig(address , uint16 ) view returns(uint16 inboundProofLibraryVersion, uint64 inboundBlockConfirmations, address relayer, uint16 outboundProofType, uint64 outboundBlockConfirmations, address oracle)
func (_UltraLightNode *UltraLightNodeCaller) AppConfig(opts *bind.CallOpts, arg0 common.Address, arg1 uint16) (struct {
	InboundProofLibraryVersion uint16
	InboundBlockConfirmations  uint64
	Relayer                    common.Address
	OutboundProofType          uint16
	OutboundBlockConfirmations uint64
	Oracle                     common.Address
}, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "appConfig", arg0, arg1)

	outstruct := new(struct {
		InboundProofLibraryVersion uint16
		InboundBlockConfirmations  uint64
		Relayer                    common.Address
		OutboundProofType          uint16
		OutboundBlockConfirmations uint64
		Oracle                     common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InboundProofLibraryVersion = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.InboundBlockConfirmations = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Relayer = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.OutboundProofType = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.OutboundBlockConfirmations = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.Oracle = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// AppConfig is a free data retrieval call binding the contract method 0xddfdef5a.
//
// Solidity: function appConfig(address , uint16 ) view returns(uint16 inboundProofLibraryVersion, uint64 inboundBlockConfirmations, address relayer, uint16 outboundProofType, uint64 outboundBlockConfirmations, address oracle)
func (_UltraLightNode *UltraLightNodeSession) AppConfig(arg0 common.Address, arg1 uint16) (struct {
	InboundProofLibraryVersion uint16
	InboundBlockConfirmations  uint64
	Relayer                    common.Address
	OutboundProofType          uint16
	OutboundBlockConfirmations uint64
	Oracle                     common.Address
}, error) {
	return _UltraLightNode.Contract.AppConfig(&_UltraLightNode.CallOpts, arg0, arg1)
}

// AppConfig is a free data retrieval call binding the contract method 0xddfdef5a.
//
// Solidity: function appConfig(address , uint16 ) view returns(uint16 inboundProofLibraryVersion, uint64 inboundBlockConfirmations, address relayer, uint16 outboundProofType, uint64 outboundBlockConfirmations, address oracle)
func (_UltraLightNode *UltraLightNodeCallerSession) AppConfig(arg0 common.Address, arg1 uint16) (struct {
	InboundProofLibraryVersion uint16
	InboundBlockConfirmations  uint64
	Relayer                    common.Address
	OutboundProofType          uint16
	OutboundBlockConfirmations uint64
	Oracle                     common.Address
}, error) {
	return _UltraLightNode.Contract.AppConfig(&_UltraLightNode.CallOpts, arg0, arg1)
}

// ChainAddressSizeMap is a free data retrieval call binding the contract method 0x959f5943.
//
// Solidity: function chainAddressSizeMap(uint16 ) view returns(uint256)
func (_UltraLightNode *UltraLightNodeCaller) ChainAddressSizeMap(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "chainAddressSizeMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainAddressSizeMap is a free data retrieval call binding the contract method 0x959f5943.
//
// Solidity: function chainAddressSizeMap(uint16 ) view returns(uint256)
func (_UltraLightNode *UltraLightNodeSession) ChainAddressSizeMap(arg0 uint16) (*big.Int, error) {
	return _UltraLightNode.Contract.ChainAddressSizeMap(&_UltraLightNode.CallOpts, arg0)
}

// ChainAddressSizeMap is a free data retrieval call binding the contract method 0x959f5943.
//
// Solidity: function chainAddressSizeMap(uint16 ) view returns(uint256)
func (_UltraLightNode *UltraLightNodeCallerSession) ChainAddressSizeMap(arg0 uint16) (*big.Int, error) {
	return _UltraLightNode.Contract.ChainAddressSizeMap(&_UltraLightNode.CallOpts, arg0)
}

// DefaultAdapterParams is a free data retrieval call binding the contract method 0x2a819bbf.
//
// Solidity: function defaultAdapterParams(uint16 , uint16 ) view returns(bytes)
func (_UltraLightNode *UltraLightNodeCaller) DefaultAdapterParams(opts *bind.CallOpts, arg0 uint16, arg1 uint16) ([]byte, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "defaultAdapterParams", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DefaultAdapterParams is a free data retrieval call binding the contract method 0x2a819bbf.
//
// Solidity: function defaultAdapterParams(uint16 , uint16 ) view returns(bytes)
func (_UltraLightNode *UltraLightNodeSession) DefaultAdapterParams(arg0 uint16, arg1 uint16) ([]byte, error) {
	return _UltraLightNode.Contract.DefaultAdapterParams(&_UltraLightNode.CallOpts, arg0, arg1)
}

// DefaultAdapterParams is a free data retrieval call binding the contract method 0x2a819bbf.
//
// Solidity: function defaultAdapterParams(uint16 , uint16 ) view returns(bytes)
func (_UltraLightNode *UltraLightNodeCallerSession) DefaultAdapterParams(arg0 uint16, arg1 uint16) ([]byte, error) {
	return _UltraLightNode.Contract.DefaultAdapterParams(&_UltraLightNode.CallOpts, arg0, arg1)
}

// DefaultAppConfig is a free data retrieval call binding the contract method 0x2f813464.
//
// Solidity: function defaultAppConfig(uint16 ) view returns(uint16 inboundProofLibraryVersion, uint64 inboundBlockConfirmations, address relayer, uint16 outboundProofType, uint64 outboundBlockConfirmations, address oracle)
func (_UltraLightNode *UltraLightNodeCaller) DefaultAppConfig(opts *bind.CallOpts, arg0 uint16) (struct {
	InboundProofLibraryVersion uint16
	InboundBlockConfirmations  uint64
	Relayer                    common.Address
	OutboundProofType          uint16
	OutboundBlockConfirmations uint64
	Oracle                     common.Address
}, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "defaultAppConfig", arg0)

	outstruct := new(struct {
		InboundProofLibraryVersion uint16
		InboundBlockConfirmations  uint64
		Relayer                    common.Address
		OutboundProofType          uint16
		OutboundBlockConfirmations uint64
		Oracle                     common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InboundProofLibraryVersion = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.InboundBlockConfirmations = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Relayer = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.OutboundProofType = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.OutboundBlockConfirmations = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.Oracle = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// DefaultAppConfig is a free data retrieval call binding the contract method 0x2f813464.
//
// Solidity: function defaultAppConfig(uint16 ) view returns(uint16 inboundProofLibraryVersion, uint64 inboundBlockConfirmations, address relayer, uint16 outboundProofType, uint64 outboundBlockConfirmations, address oracle)
func (_UltraLightNode *UltraLightNodeSession) DefaultAppConfig(arg0 uint16) (struct {
	InboundProofLibraryVersion uint16
	InboundBlockConfirmations  uint64
	Relayer                    common.Address
	OutboundProofType          uint16
	OutboundBlockConfirmations uint64
	Oracle                     common.Address
}, error) {
	return _UltraLightNode.Contract.DefaultAppConfig(&_UltraLightNode.CallOpts, arg0)
}

// DefaultAppConfig is a free data retrieval call binding the contract method 0x2f813464.
//
// Solidity: function defaultAppConfig(uint16 ) view returns(uint16 inboundProofLibraryVersion, uint64 inboundBlockConfirmations, address relayer, uint16 outboundProofType, uint64 outboundBlockConfirmations, address oracle)
func (_UltraLightNode *UltraLightNodeCallerSession) DefaultAppConfig(arg0 uint16) (struct {
	InboundProofLibraryVersion uint16
	InboundBlockConfirmations  uint64
	Relayer                    common.Address
	OutboundProofType          uint16
	OutboundBlockConfirmations uint64
	Oracle                     common.Address
}, error) {
	return _UltraLightNode.Contract.DefaultAppConfig(&_UltraLightNode.CallOpts, arg0)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_UltraLightNode *UltraLightNodeCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_UltraLightNode *UltraLightNodeSession) Endpoint() (common.Address, error) {
	return _UltraLightNode.Contract.Endpoint(&_UltraLightNode.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_UltraLightNode *UltraLightNodeCallerSession) Endpoint() (common.Address, error) {
	return _UltraLightNode.Contract.Endpoint(&_UltraLightNode.CallOpts)
}

// EstimateFees is a free data retrieval call binding the contract method 0x40a7bb10.
//
// Solidity: function estimateFees(uint16 _dstChainId, address _ua, bytes _payload, bool _payInZRO, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_UltraLightNode *UltraLightNodeCaller) EstimateFees(opts *bind.CallOpts, _dstChainId uint16, _ua common.Address, _payload []byte, _payInZRO bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "estimateFees", _dstChainId, _ua, _payload, _payInZRO, _adapterParams)

	outstruct := new(struct {
		NativeFee *big.Int
		ZroFee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NativeFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ZroFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateFees is a free data retrieval call binding the contract method 0x40a7bb10.
//
// Solidity: function estimateFees(uint16 _dstChainId, address _ua, bytes _payload, bool _payInZRO, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_UltraLightNode *UltraLightNodeSession) EstimateFees(_dstChainId uint16, _ua common.Address, _payload []byte, _payInZRO bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _UltraLightNode.Contract.EstimateFees(&_UltraLightNode.CallOpts, _dstChainId, _ua, _payload, _payInZRO, _adapterParams)
}

// EstimateFees is a free data retrieval call binding the contract method 0x40a7bb10.
//
// Solidity: function estimateFees(uint16 _dstChainId, address _ua, bytes _payload, bool _payInZRO, bytes _adapterParams) view returns(uint256 nativeFee, uint256 zroFee)
func (_UltraLightNode *UltraLightNodeCallerSession) EstimateFees(_dstChainId uint16, _ua common.Address, _payload []byte, _payInZRO bool, _adapterParams []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _UltraLightNode.Contract.EstimateFees(&_UltraLightNode.CallOpts, _dstChainId, _ua, _payload, _payInZRO, _adapterParams)
}

// GetAppConfig is a free data retrieval call binding the contract method 0xa4662222.
//
// Solidity: function getAppConfig(uint16 _remoteChainId, address _ua) view returns((uint16,uint64,address,uint16,uint64,address))
func (_UltraLightNode *UltraLightNodeCaller) GetAppConfig(opts *bind.CallOpts, _remoteChainId uint16, _ua common.Address) (ILayerZeroUltraLightNodeV2ApplicationConfiguration, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "getAppConfig", _remoteChainId, _ua)

	if err != nil {
		return *new(ILayerZeroUltraLightNodeV2ApplicationConfiguration), err
	}

	out0 := *abi.ConvertType(out[0], new(ILayerZeroUltraLightNodeV2ApplicationConfiguration)).(*ILayerZeroUltraLightNodeV2ApplicationConfiguration)

	return out0, err

}

// GetAppConfig is a free data retrieval call binding the contract method 0xa4662222.
//
// Solidity: function getAppConfig(uint16 _remoteChainId, address _ua) view returns((uint16,uint64,address,uint16,uint64,address))
func (_UltraLightNode *UltraLightNodeSession) GetAppConfig(_remoteChainId uint16, _ua common.Address) (ILayerZeroUltraLightNodeV2ApplicationConfiguration, error) {
	return _UltraLightNode.Contract.GetAppConfig(&_UltraLightNode.CallOpts, _remoteChainId, _ua)
}

// GetAppConfig is a free data retrieval call binding the contract method 0xa4662222.
//
// Solidity: function getAppConfig(uint16 _remoteChainId, address _ua) view returns((uint16,uint64,address,uint16,uint64,address))
func (_UltraLightNode *UltraLightNodeCallerSession) GetAppConfig(_remoteChainId uint16, _ua common.Address) (ILayerZeroUltraLightNodeV2ApplicationConfiguration, error) {
	return _UltraLightNode.Contract.GetAppConfig(&_UltraLightNode.CallOpts, _remoteChainId, _ua)
}

// GetConfig is a free data retrieval call binding the contract method 0x52d2871f.
//
// Solidity: function getConfig(uint16 _remoteChainId, address _ua, uint256 _configType) view returns(bytes)
func (_UltraLightNode *UltraLightNodeCaller) GetConfig(opts *bind.CallOpts, _remoteChainId uint16, _ua common.Address, _configType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "getConfig", _remoteChainId, _ua, _configType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0x52d2871f.
//
// Solidity: function getConfig(uint16 _remoteChainId, address _ua, uint256 _configType) view returns(bytes)
func (_UltraLightNode *UltraLightNodeSession) GetConfig(_remoteChainId uint16, _ua common.Address, _configType *big.Int) ([]byte, error) {
	return _UltraLightNode.Contract.GetConfig(&_UltraLightNode.CallOpts, _remoteChainId, _ua, _configType)
}

// GetConfig is a free data retrieval call binding the contract method 0x52d2871f.
//
// Solidity: function getConfig(uint16 _remoteChainId, address _ua, uint256 _configType) view returns(bytes)
func (_UltraLightNode *UltraLightNodeCallerSession) GetConfig(_remoteChainId uint16, _ua common.Address, _configType *big.Int) ([]byte, error) {
	return _UltraLightNode.Contract.GetConfig(&_UltraLightNode.CallOpts, _remoteChainId, _ua, _configType)
}

// GetOutboundNonce is a free data retrieval call binding the contract method 0xb9a99bed.
//
// Solidity: function getOutboundNonce(uint16 _chainId, bytes _path) view returns(uint64)
func (_UltraLightNode *UltraLightNodeCaller) GetOutboundNonce(opts *bind.CallOpts, _chainId uint16, _path []byte) (uint64, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "getOutboundNonce", _chainId, _path)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetOutboundNonce is a free data retrieval call binding the contract method 0xb9a99bed.
//
// Solidity: function getOutboundNonce(uint16 _chainId, bytes _path) view returns(uint64)
func (_UltraLightNode *UltraLightNodeSession) GetOutboundNonce(_chainId uint16, _path []byte) (uint64, error) {
	return _UltraLightNode.Contract.GetOutboundNonce(&_UltraLightNode.CallOpts, _chainId, _path)
}

// GetOutboundNonce is a free data retrieval call binding the contract method 0xb9a99bed.
//
// Solidity: function getOutboundNonce(uint16 _chainId, bytes _path) view returns(uint64)
func (_UltraLightNode *UltraLightNodeCallerSession) GetOutboundNonce(_chainId uint16, _path []byte) (uint64, error) {
	return _UltraLightNode.Contract.GetOutboundNonce(&_UltraLightNode.CallOpts, _chainId, _path)
}

// HashLookup is a free data retrieval call binding the contract method 0x759c5b3b.
//
// Solidity: function hashLookup(address , uint16 , bytes32 , bytes32 ) view returns(uint256)
func (_UltraLightNode *UltraLightNodeCaller) HashLookup(opts *bind.CallOpts, arg0 common.Address, arg1 uint16, arg2 [32]byte, arg3 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "hashLookup", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashLookup is a free data retrieval call binding the contract method 0x759c5b3b.
//
// Solidity: function hashLookup(address , uint16 , bytes32 , bytes32 ) view returns(uint256)
func (_UltraLightNode *UltraLightNodeSession) HashLookup(arg0 common.Address, arg1 uint16, arg2 [32]byte, arg3 [32]byte) (*big.Int, error) {
	return _UltraLightNode.Contract.HashLookup(&_UltraLightNode.CallOpts, arg0, arg1, arg2, arg3)
}

// HashLookup is a free data retrieval call binding the contract method 0x759c5b3b.
//
// Solidity: function hashLookup(address , uint16 , bytes32 , bytes32 ) view returns(uint256)
func (_UltraLightNode *UltraLightNodeCallerSession) HashLookup(arg0 common.Address, arg1 uint16, arg2 [32]byte, arg3 [32]byte) (*big.Int, error) {
	return _UltraLightNode.Contract.HashLookup(&_UltraLightNode.CallOpts, arg0, arg1, arg2, arg3)
}

// InboundProofLibrary is a free data retrieval call binding the contract method 0xdb00719b.
//
// Solidity: function inboundProofLibrary(uint16 , uint16 ) view returns(address)
func (_UltraLightNode *UltraLightNodeCaller) InboundProofLibrary(opts *bind.CallOpts, arg0 uint16, arg1 uint16) (common.Address, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "inboundProofLibrary", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InboundProofLibrary is a free data retrieval call binding the contract method 0xdb00719b.
//
// Solidity: function inboundProofLibrary(uint16 , uint16 ) view returns(address)
func (_UltraLightNode *UltraLightNodeSession) InboundProofLibrary(arg0 uint16, arg1 uint16) (common.Address, error) {
	return _UltraLightNode.Contract.InboundProofLibrary(&_UltraLightNode.CallOpts, arg0, arg1)
}

// InboundProofLibrary is a free data retrieval call binding the contract method 0xdb00719b.
//
// Solidity: function inboundProofLibrary(uint16 , uint16 ) view returns(address)
func (_UltraLightNode *UltraLightNodeCallerSession) InboundProofLibrary(arg0 uint16, arg1 uint16) (common.Address, error) {
	return _UltraLightNode.Contract.InboundProofLibrary(&_UltraLightNode.CallOpts, arg0, arg1)
}

// LayerZeroToken is a free data retrieval call binding the contract method 0x07b9ca7c.
//
// Solidity: function layerZeroToken() view returns(address)
func (_UltraLightNode *UltraLightNodeCaller) LayerZeroToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "layerZeroToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LayerZeroToken is a free data retrieval call binding the contract method 0x07b9ca7c.
//
// Solidity: function layerZeroToken() view returns(address)
func (_UltraLightNode *UltraLightNodeSession) LayerZeroToken() (common.Address, error) {
	return _UltraLightNode.Contract.LayerZeroToken(&_UltraLightNode.CallOpts)
}

// LayerZeroToken is a free data retrieval call binding the contract method 0x07b9ca7c.
//
// Solidity: function layerZeroToken() view returns(address)
func (_UltraLightNode *UltraLightNodeCallerSession) LayerZeroToken() (common.Address, error) {
	return _UltraLightNode.Contract.LayerZeroToken(&_UltraLightNode.CallOpts)
}

// LocalChainId is a free data retrieval call binding the contract method 0x5b056da5.
//
// Solidity: function localChainId() view returns(uint16)
func (_UltraLightNode *UltraLightNodeCaller) LocalChainId(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "localChainId")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LocalChainId is a free data retrieval call binding the contract method 0x5b056da5.
//
// Solidity: function localChainId() view returns(uint16)
func (_UltraLightNode *UltraLightNodeSession) LocalChainId() (uint16, error) {
	return _UltraLightNode.Contract.LocalChainId(&_UltraLightNode.CallOpts)
}

// LocalChainId is a free data retrieval call binding the contract method 0x5b056da5.
//
// Solidity: function localChainId() view returns(uint16)
func (_UltraLightNode *UltraLightNodeCallerSession) LocalChainId() (uint16, error) {
	return _UltraLightNode.Contract.LocalChainId(&_UltraLightNode.CallOpts)
}

// MaxInboundProofLibrary is a free data retrieval call binding the contract method 0xb8e7e3e0.
//
// Solidity: function maxInboundProofLibrary(uint16 ) view returns(uint16)
func (_UltraLightNode *UltraLightNodeCaller) MaxInboundProofLibrary(opts *bind.CallOpts, arg0 uint16) (uint16, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "maxInboundProofLibrary", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaxInboundProofLibrary is a free data retrieval call binding the contract method 0xb8e7e3e0.
//
// Solidity: function maxInboundProofLibrary(uint16 ) view returns(uint16)
func (_UltraLightNode *UltraLightNodeSession) MaxInboundProofLibrary(arg0 uint16) (uint16, error) {
	return _UltraLightNode.Contract.MaxInboundProofLibrary(&_UltraLightNode.CallOpts, arg0)
}

// MaxInboundProofLibrary is a free data retrieval call binding the contract method 0xb8e7e3e0.
//
// Solidity: function maxInboundProofLibrary(uint16 ) view returns(uint16)
func (_UltraLightNode *UltraLightNodeCallerSession) MaxInboundProofLibrary(arg0 uint16) (uint16, error) {
	return _UltraLightNode.Contract.MaxInboundProofLibrary(&_UltraLightNode.CallOpts, arg0)
}

// NativeFees is a free data retrieval call binding the contract method 0xf58589a2.
//
// Solidity: function nativeFees(address ) view returns(uint256)
func (_UltraLightNode *UltraLightNodeCaller) NativeFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "nativeFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NativeFees is a free data retrieval call binding the contract method 0xf58589a2.
//
// Solidity: function nativeFees(address ) view returns(uint256)
func (_UltraLightNode *UltraLightNodeSession) NativeFees(arg0 common.Address) (*big.Int, error) {
	return _UltraLightNode.Contract.NativeFees(&_UltraLightNode.CallOpts, arg0)
}

// NativeFees is a free data retrieval call binding the contract method 0xf58589a2.
//
// Solidity: function nativeFees(address ) view returns(uint256)
func (_UltraLightNode *UltraLightNodeCallerSession) NativeFees(arg0 common.Address) (*big.Int, error) {
	return _UltraLightNode.Contract.NativeFees(&_UltraLightNode.CallOpts, arg0)
}

// NonceContract is a free data retrieval call binding the contract method 0x02bd9743.
//
// Solidity: function nonceContract() view returns(address)
func (_UltraLightNode *UltraLightNodeCaller) NonceContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "nonceContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NonceContract is a free data retrieval call binding the contract method 0x02bd9743.
//
// Solidity: function nonceContract() view returns(address)
func (_UltraLightNode *UltraLightNodeSession) NonceContract() (common.Address, error) {
	return _UltraLightNode.Contract.NonceContract(&_UltraLightNode.CallOpts)
}

// NonceContract is a free data retrieval call binding the contract method 0x02bd9743.
//
// Solidity: function nonceContract() view returns(address)
func (_UltraLightNode *UltraLightNodeCallerSession) NonceContract() (common.Address, error) {
	return _UltraLightNode.Contract.NonceContract(&_UltraLightNode.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UltraLightNode *UltraLightNodeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UltraLightNode *UltraLightNodeSession) Owner() (common.Address, error) {
	return _UltraLightNode.Contract.Owner(&_UltraLightNode.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UltraLightNode *UltraLightNodeCallerSession) Owner() (common.Address, error) {
	return _UltraLightNode.Contract.Owner(&_UltraLightNode.CallOpts)
}

// SupportedOutboundProof is a free data retrieval call binding the contract method 0xd543c774.
//
// Solidity: function supportedOutboundProof(uint16 , uint16 ) view returns(bool)
func (_UltraLightNode *UltraLightNodeCaller) SupportedOutboundProof(opts *bind.CallOpts, arg0 uint16, arg1 uint16) (bool, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "supportedOutboundProof", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportedOutboundProof is a free data retrieval call binding the contract method 0xd543c774.
//
// Solidity: function supportedOutboundProof(uint16 , uint16 ) view returns(bool)
func (_UltraLightNode *UltraLightNodeSession) SupportedOutboundProof(arg0 uint16, arg1 uint16) (bool, error) {
	return _UltraLightNode.Contract.SupportedOutboundProof(&_UltraLightNode.CallOpts, arg0, arg1)
}

// SupportedOutboundProof is a free data retrieval call binding the contract method 0xd543c774.
//
// Solidity: function supportedOutboundProof(uint16 , uint16 ) view returns(bool)
func (_UltraLightNode *UltraLightNodeCallerSession) SupportedOutboundProof(arg0 uint16, arg1 uint16) (bool, error) {
	return _UltraLightNode.Contract.SupportedOutboundProof(&_UltraLightNode.CallOpts, arg0, arg1)
}

// TreasuryContract is a free data retrieval call binding the contract method 0x18da0011.
//
// Solidity: function treasuryContract() view returns(address)
func (_UltraLightNode *UltraLightNodeCaller) TreasuryContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "treasuryContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryContract is a free data retrieval call binding the contract method 0x18da0011.
//
// Solidity: function treasuryContract() view returns(address)
func (_UltraLightNode *UltraLightNodeSession) TreasuryContract() (common.Address, error) {
	return _UltraLightNode.Contract.TreasuryContract(&_UltraLightNode.CallOpts)
}

// TreasuryContract is a free data retrieval call binding the contract method 0x18da0011.
//
// Solidity: function treasuryContract() view returns(address)
func (_UltraLightNode *UltraLightNodeCallerSession) TreasuryContract() (common.Address, error) {
	return _UltraLightNode.Contract.TreasuryContract(&_UltraLightNode.CallOpts)
}

// TreasuryZROFees is a free data retrieval call binding the contract method 0xf47a5feb.
//
// Solidity: function treasuryZROFees() view returns(uint256)
func (_UltraLightNode *UltraLightNodeCaller) TreasuryZROFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "treasuryZROFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryZROFees is a free data retrieval call binding the contract method 0xf47a5feb.
//
// Solidity: function treasuryZROFees() view returns(uint256)
func (_UltraLightNode *UltraLightNodeSession) TreasuryZROFees() (*big.Int, error) {
	return _UltraLightNode.Contract.TreasuryZROFees(&_UltraLightNode.CallOpts)
}

// TreasuryZROFees is a free data retrieval call binding the contract method 0xf47a5feb.
//
// Solidity: function treasuryZROFees() view returns(uint256)
func (_UltraLightNode *UltraLightNodeCallerSession) TreasuryZROFees() (*big.Int, error) {
	return _UltraLightNode.Contract.TreasuryZROFees(&_UltraLightNode.CallOpts)
}

// UlnLookup is a free data retrieval call binding the contract method 0xea216c21.
//
// Solidity: function ulnLookup(uint16 ) view returns(bytes32)
func (_UltraLightNode *UltraLightNodeCaller) UlnLookup(opts *bind.CallOpts, arg0 uint16) ([32]byte, error) {
	var out []interface{}
	err := _UltraLightNode.contract.Call(opts, &out, "ulnLookup", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UlnLookup is a free data retrieval call binding the contract method 0xea216c21.
//
// Solidity: function ulnLookup(uint16 ) view returns(bytes32)
func (_UltraLightNode *UltraLightNodeSession) UlnLookup(arg0 uint16) ([32]byte, error) {
	return _UltraLightNode.Contract.UlnLookup(&_UltraLightNode.CallOpts, arg0)
}

// UlnLookup is a free data retrieval call binding the contract method 0xea216c21.
//
// Solidity: function ulnLookup(uint16 ) view returns(bytes32)
func (_UltraLightNode *UltraLightNodeCallerSession) UlnLookup(arg0 uint16) ([32]byte, error) {
	return _UltraLightNode.Contract.UlnLookup(&_UltraLightNode.CallOpts, arg0)
}

// AddInboundProofLibraryForChain is a paid mutator transaction binding the contract method 0x8207f79d.
//
// Solidity: function addInboundProofLibraryForChain(uint16 _chainId, address _library) returns()
func (_UltraLightNode *UltraLightNodeTransactor) AddInboundProofLibraryForChain(opts *bind.TransactOpts, _chainId uint16, _library common.Address) (*types.Transaction, error) {
	return _UltraLightNode.contract.Transact(opts, "addInboundProofLibraryForChain", _chainId, _library)
}

// AddInboundProofLibraryForChain is a paid mutator transaction binding the contract method 0x8207f79d.
//
// Solidity: function addInboundProofLibraryForChain(uint16 _chainId, address _library) returns()
func (_UltraLightNode *UltraLightNodeSession) AddInboundProofLibraryForChain(_chainId uint16, _library common.Address) (*types.Transaction, error) {
	return _UltraLightNode.Contract.AddInboundProofLibraryForChain(&_UltraLightNode.TransactOpts, _chainId, _library)
}

// AddInboundProofLibraryForChain is a paid mutator transaction binding the contract method 0x8207f79d.
//
// Solidity: function addInboundProofLibraryForChain(uint16 _chainId, address _library) returns()
func (_UltraLightNode *UltraLightNodeTransactorSession) AddInboundProofLibraryForChain(_chainId uint16, _library common.Address) (*types.Transaction, error) {
	return _UltraLightNode.Contract.AddInboundProofLibraryForChain(&_UltraLightNode.TransactOpts, _chainId, _library)
}

// EnableSupportedOutboundProof is a paid mutator transaction binding the contract method 0xeb0d4c31.
//
// Solidity: function enableSupportedOutboundProof(uint16 _chainId, uint16 _proofType) returns()
func (_UltraLightNode *UltraLightNodeTransactor) EnableSupportedOutboundProof(opts *bind.TransactOpts, _chainId uint16, _proofType uint16) (*types.Transaction, error) {
	return _UltraLightNode.contract.Transact(opts, "enableSupportedOutboundProof", _chainId, _proofType)
}

// EnableSupportedOutboundProof is a paid mutator transaction binding the contract method 0xeb0d4c31.
//
// Solidity: function enableSupportedOutboundProof(uint16 _chainId, uint16 _proofType) returns()
func (_UltraLightNode *UltraLightNodeSession) EnableSupportedOutboundProof(_chainId uint16, _proofType uint16) (*types.Transaction, error) {
	return _UltraLightNode.Contract.EnableSupportedOutboundProof(&_UltraLightNode.TransactOpts, _chainId, _proofType)
}

// EnableSupportedOutboundProof is a paid mutator transaction binding the contract method 0xeb0d4c31.
//
// Solidity: function enableSupportedOutboundProof(uint16 _chainId, uint16 _proofType) returns()
func (_UltraLightNode *UltraLightNodeTransactorSession) EnableSupportedOutboundProof(_chainId uint16, _proofType uint16) (*types.Transaction, error) {
	return _UltraLightNode.Contract.EnableSupportedOutboundProof(&_UltraLightNode.TransactOpts, _chainId, _proofType)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UltraLightNode *UltraLightNodeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UltraLightNode.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UltraLightNode *UltraLightNodeSession) RenounceOwnership() (*types.Transaction, error) {
	return _UltraLightNode.Contract.RenounceOwnership(&_UltraLightNode.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UltraLightNode *UltraLightNodeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UltraLightNode.Contract.RenounceOwnership(&_UltraLightNode.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0x4d3a0f7c.
//
// Solidity: function send(address _ua, uint64 , uint16 _dstChainId, bytes _path, bytes _payload, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_UltraLightNode *UltraLightNodeTransactor) Send(opts *bind.TransactOpts, _ua common.Address, arg1 uint64, _dstChainId uint16, _path []byte, _payload []byte, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _UltraLightNode.contract.Transact(opts, "send", _ua, arg1, _dstChainId, _path, _payload, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// Send is a paid mutator transaction binding the contract method 0x4d3a0f7c.
//
// Solidity: function send(address _ua, uint64 , uint16 _dstChainId, bytes _path, bytes _payload, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_UltraLightNode *UltraLightNodeSession) Send(_ua common.Address, arg1 uint64, _dstChainId uint16, _path []byte, _payload []byte, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _UltraLightNode.Contract.Send(&_UltraLightNode.TransactOpts, _ua, arg1, _dstChainId, _path, _payload, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// Send is a paid mutator transaction binding the contract method 0x4d3a0f7c.
//
// Solidity: function send(address _ua, uint64 , uint16 _dstChainId, bytes _path, bytes _payload, address _refundAddress, address _zroPaymentAddress, bytes _adapterParams) payable returns()
func (_UltraLightNode *UltraLightNodeTransactorSession) Send(_ua common.Address, arg1 uint64, _dstChainId uint16, _path []byte, _payload []byte, _refundAddress common.Address, _zroPaymentAddress common.Address, _adapterParams []byte) (*types.Transaction, error) {
	return _UltraLightNode.Contract.Send(&_UltraLightNode.TransactOpts, _ua, arg1, _dstChainId, _path, _payload, _refundAddress, _zroPaymentAddress, _adapterParams)
}

// SetChainAddressSize is a paid mutator transaction binding the contract method 0x87078f9f.
//
// Solidity: function setChainAddressSize(uint16 _chainId, uint256 _size) returns()
func (_UltraLightNode *UltraLightNodeTransactor) SetChainAddressSize(opts *bind.TransactOpts, _chainId uint16, _size *big.Int) (*types.Transaction, error) {
	return _UltraLightNode.contract.Transact(opts, "setChainAddressSize", _chainId, _size)
}

// SetChainAddressSize is a paid mutator transaction binding the contract method 0x87078f9f.
//
// Solidity: function setChainAddressSize(uint16 _chainId, uint256 _size) returns()
func (_UltraLightNode *UltraLightNodeSession) SetChainAddressSize(_chainId uint16, _size *big.Int) (*types.Transaction, error) {
	return _UltraLightNode.Contract.SetChainAddressSize(&_UltraLightNode.TransactOpts, _chainId, _size)
}

// SetChainAddressSize is a paid mutator transaction binding the contract method 0x87078f9f.
//
// Solidity: function setChainAddressSize(uint16 _chainId, uint256 _size) returns()
func (_UltraLightNode *UltraLightNodeTransactorSession) SetChainAddressSize(_chainId uint16, _size *big.Int) (*types.Transaction, error) {
	return _UltraLightNode.Contract.SetChainAddressSize(&_UltraLightNode.TransactOpts, _chainId, _size)
}

// SetConfig is a paid mutator transaction binding the contract method 0xf8e1734c.
//
// Solidity: function setConfig(uint16 _remoteChainId, address _ua, uint256 _configType, bytes _config) returns()
func (_UltraLightNode *UltraLightNodeTransactor) SetConfig(opts *bind.TransactOpts, _remoteChainId uint16, _ua common.Address, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _UltraLightNode.contract.Transact(opts, "setConfig", _remoteChainId, _ua, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xf8e1734c.
//
// Solidity: function setConfig(uint16 _remoteChainId, address _ua, uint256 _configType, bytes _config) returns()
func (_UltraLightNode *UltraLightNodeSession) SetConfig(_remoteChainId uint16, _ua common.Address, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _UltraLightNode.Contract.SetConfig(&_UltraLightNode.TransactOpts, _remoteChainId, _ua, _configType, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0xf8e1734c.
//
// Solidity: function setConfig(uint16 _remoteChainId, address _ua, uint256 _configType, bytes _config) returns()
func (_UltraLightNode *UltraLightNodeTransactorSession) SetConfig(_remoteChainId uint16, _ua common.Address, _configType *big.Int, _config []byte) (*types.Transaction, error) {
	return _UltraLightNode.Contract.SetConfig(&_UltraLightNode.TransactOpts, _remoteChainId, _ua, _configType, _config)
}

// SetDefaultAdapterParamsForChainId is a paid mutator transaction binding the contract method 0x8317814a.
//
// Solidity: function setDefaultAdapterParamsForChainId(uint16 _chainId, uint16 _proofType, bytes _adapterParams) returns()
func (_UltraLightNode *UltraLightNodeTransactor) SetDefaultAdapterParamsForChainId(opts *bind.TransactOpts, _chainId uint16, _proofType uint16, _adapterParams []byte) (*types.Transaction, error) {
	return _UltraLightNode.contract.Transact(opts, "setDefaultAdapterParamsForChainId", _chainId, _proofType, _adapterParams)
}

// SetDefaultAdapterParamsForChainId is a paid mutator transaction binding the contract method 0x8317814a.
//
// Solidity: function setDefaultAdapterParamsForChainId(uint16 _chainId, uint16 _proofType, bytes _adapterParams) returns()
func (_UltraLightNode *UltraLightNodeSession) SetDefaultAdapterParamsForChainId(_chainId uint16, _proofType uint16, _adapterParams []byte) (*types.Transaction, error) {
	return _UltraLightNode.Contract.SetDefaultAdapterParamsForChainId(&_UltraLightNode.TransactOpts, _chainId, _proofType, _adapterParams)
}

// SetDefaultAdapterParamsForChainId is a paid mutator transaction binding the contract method 0x8317814a.
//
// Solidity: function setDefaultAdapterParamsForChainId(uint16 _chainId, uint16 _proofType, bytes _adapterParams) returns()
func (_UltraLightNode *UltraLightNodeTransactorSession) SetDefaultAdapterParamsForChainId(_chainId uint16, _proofType uint16, _adapterParams []byte) (*types.Transaction, error) {
	return _UltraLightNode.Contract.SetDefaultAdapterParamsForChainId(&_UltraLightNode.TransactOpts, _chainId, _proofType, _adapterParams)
}

// SetDefaultConfigForChainId is a paid mutator transaction binding the contract method 0x6a14ac82.
//
// Solidity: function setDefaultConfigForChainId(uint16 _chainId, uint16 _inboundProofLibraryVersion, uint64 _inboundBlockConfirmations, address _relayer, uint16 _outboundProofType, uint64 _outboundBlockConfirmations, address _oracle) returns()
func (_UltraLightNode *UltraLightNodeTransactor) SetDefaultConfigForChainId(opts *bind.TransactOpts, _chainId uint16, _inboundProofLibraryVersion uint16, _inboundBlockConfirmations uint64, _relayer common.Address, _outboundProofType uint16, _outboundBlockConfirmations uint64, _oracle common.Address) (*types.Transaction, error) {
	return _UltraLightNode.contract.Transact(opts, "setDefaultConfigForChainId", _chainId, _inboundProofLibraryVersion, _inboundBlockConfirmations, _relayer, _outboundProofType, _outboundBlockConfirmations, _oracle)
}

// SetDefaultConfigForChainId is a paid mutator transaction binding the contract method 0x6a14ac82.
//
// Solidity: function setDefaultConfigForChainId(uint16 _chainId, uint16 _inboundProofLibraryVersion, uint64 _inboundBlockConfirmations, address _relayer, uint16 _outboundProofType, uint64 _outboundBlockConfirmations, address _oracle) returns()
func (_UltraLightNode *UltraLightNodeSession) SetDefaultConfigForChainId(_chainId uint16, _inboundProofLibraryVersion uint16, _inboundBlockConfirmations uint64, _relayer common.Address, _outboundProofType uint16, _outboundBlockConfirmations uint64, _oracle common.Address) (*types.Transaction, error) {
	return _UltraLightNode.Contract.SetDefaultConfigForChainId(&_UltraLightNode.TransactOpts, _chainId, _inboundProofLibraryVersion, _inboundBlockConfirmations, _relayer, _outboundProofType, _outboundBlockConfirmations, _oracle)
}

// SetDefaultConfigForChainId is a paid mutator transaction binding the contract method 0x6a14ac82.
//
// Solidity: function setDefaultConfigForChainId(uint16 _chainId, uint16 _inboundProofLibraryVersion, uint64 _inboundBlockConfirmations, address _relayer, uint16 _outboundProofType, uint64 _outboundBlockConfirmations, address _oracle) returns()
func (_UltraLightNode *UltraLightNodeTransactorSession) SetDefaultConfigForChainId(_chainId uint16, _inboundProofLibraryVersion uint16, _inboundBlockConfirmations uint64, _relayer common.Address, _outboundProofType uint16, _outboundBlockConfirmations uint64, _oracle common.Address) (*types.Transaction, error) {
	return _UltraLightNode.Contract.SetDefaultConfigForChainId(&_UltraLightNode.TransactOpts, _chainId, _inboundProofLibraryVersion, _inboundBlockConfirmations, _relayer, _outboundProofType, _outboundBlockConfirmations, _oracle)
}

// SetLayerZeroToken is a paid mutator transaction binding the contract method 0x52d3b500.
//
// Solidity: function setLayerZeroToken(address _layerZeroToken) returns()
func (_UltraLightNode *UltraLightNodeTransactor) SetLayerZeroToken(opts *bind.TransactOpts, _layerZeroToken common.Address) (*types.Transaction, error) {
	return _UltraLightNode.contract.Transact(opts, "setLayerZeroToken", _layerZeroToken)
}

// SetLayerZeroToken is a paid mutator transaction binding the contract method 0x52d3b500.
//
// Solidity: function setLayerZeroToken(address _layerZeroToken) returns()
func (_UltraLightNode *UltraLightNodeSession) SetLayerZeroToken(_layerZeroToken common.Address) (*types.Transaction, error) {
	return _UltraLightNode.Contract.SetLayerZeroToken(&_UltraLightNode.TransactOpts, _layerZeroToken)
}

// SetLayerZeroToken is a paid mutator transaction binding the contract method 0x52d3b500.
//
// Solidity: function setLayerZeroToken(address _layerZeroToken) returns()
func (_UltraLightNode *UltraLightNodeTransactorSession) SetLayerZeroToken(_layerZeroToken common.Address) (*types.Transaction, error) {
	return _UltraLightNode.Contract.SetLayerZeroToken(&_UltraLightNode.TransactOpts, _layerZeroToken)
}

// SetRemoteUln is a paid mutator transaction binding the contract method 0xed28580a.
//
// Solidity: function setRemoteUln(uint16 _remoteChainId, bytes32 _remoteUln) returns()
func (_UltraLightNode *UltraLightNodeTransactor) SetRemoteUln(opts *bind.TransactOpts, _remoteChainId uint16, _remoteUln [32]byte) (*types.Transaction, error) {
	return _UltraLightNode.contract.Transact(opts, "setRemoteUln", _remoteChainId, _remoteUln)
}

// SetRemoteUln is a paid mutator transaction binding the contract method 0xed28580a.
//
// Solidity: function setRemoteUln(uint16 _remoteChainId, bytes32 _remoteUln) returns()
func (_UltraLightNode *UltraLightNodeSession) SetRemoteUln(_remoteChainId uint16, _remoteUln [32]byte) (*types.Transaction, error) {
	return _UltraLightNode.Contract.SetRemoteUln(&_UltraLightNode.TransactOpts, _remoteChainId, _remoteUln)
}

// SetRemoteUln is a paid mutator transaction binding the contract method 0xed28580a.
//
// Solidity: function setRemoteUln(uint16 _remoteChainId, bytes32 _remoteUln) returns()
func (_UltraLightNode *UltraLightNodeTransactorSession) SetRemoteUln(_remoteChainId uint16, _remoteUln [32]byte) (*types.Transaction, error) {
	return _UltraLightNode.Contract.SetRemoteUln(&_UltraLightNode.TransactOpts, _remoteChainId, _remoteUln)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_UltraLightNode *UltraLightNodeTransactor) SetTreasury(opts *bind.TransactOpts, _treasury common.Address) (*types.Transaction, error) {
	return _UltraLightNode.contract.Transact(opts, "setTreasury", _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_UltraLightNode *UltraLightNodeSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _UltraLightNode.Contract.SetTreasury(&_UltraLightNode.TransactOpts, _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_UltraLightNode *UltraLightNodeTransactorSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _UltraLightNode.Contract.SetTreasury(&_UltraLightNode.TransactOpts, _treasury)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UltraLightNode *UltraLightNodeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UltraLightNode.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UltraLightNode *UltraLightNodeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UltraLightNode.Contract.TransferOwnership(&_UltraLightNode.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UltraLightNode *UltraLightNodeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UltraLightNode.Contract.TransferOwnership(&_UltraLightNode.TransactOpts, newOwner)
}

// UpdateHash is a paid mutator transaction binding the contract method 0x704316e5.
//
// Solidity: function updateHash(uint16 _srcChainId, bytes32 _lookupHash, uint256 _confirmations, bytes32 _blockData) returns()
func (_UltraLightNode *UltraLightNodeTransactor) UpdateHash(opts *bind.TransactOpts, _srcChainId uint16, _lookupHash [32]byte, _confirmations *big.Int, _blockData [32]byte) (*types.Transaction, error) {
	return _UltraLightNode.contract.Transact(opts, "updateHash", _srcChainId, _lookupHash, _confirmations, _blockData)
}

// UpdateHash is a paid mutator transaction binding the contract method 0x704316e5.
//
// Solidity: function updateHash(uint16 _srcChainId, bytes32 _lookupHash, uint256 _confirmations, bytes32 _blockData) returns()
func (_UltraLightNode *UltraLightNodeSession) UpdateHash(_srcChainId uint16, _lookupHash [32]byte, _confirmations *big.Int, _blockData [32]byte) (*types.Transaction, error) {
	return _UltraLightNode.Contract.UpdateHash(&_UltraLightNode.TransactOpts, _srcChainId, _lookupHash, _confirmations, _blockData)
}

// UpdateHash is a paid mutator transaction binding the contract method 0x704316e5.
//
// Solidity: function updateHash(uint16 _srcChainId, bytes32 _lookupHash, uint256 _confirmations, bytes32 _blockData) returns()
func (_UltraLightNode *UltraLightNodeTransactorSession) UpdateHash(_srcChainId uint16, _lookupHash [32]byte, _confirmations *big.Int, _blockData [32]byte) (*types.Transaction, error) {
	return _UltraLightNode.Contract.UpdateHash(&_UltraLightNode.TransactOpts, _srcChainId, _lookupHash, _confirmations, _blockData)
}

// ValidateTransactionProof is a paid mutator transaction binding the contract method 0x987fa2d5.
//
// Solidity: function validateTransactionProof(uint16 _srcChainId, address _dstAddress, uint256 _gasLimit, bytes32 _lookupHash, bytes32 _blockData, bytes _transactionProof) returns()
func (_UltraLightNode *UltraLightNodeTransactor) ValidateTransactionProof(opts *bind.TransactOpts, _srcChainId uint16, _dstAddress common.Address, _gasLimit *big.Int, _lookupHash [32]byte, _blockData [32]byte, _transactionProof []byte) (*types.Transaction, error) {
	return _UltraLightNode.contract.Transact(opts, "validateTransactionProof", _srcChainId, _dstAddress, _gasLimit, _lookupHash, _blockData, _transactionProof)
}

// ValidateTransactionProof is a paid mutator transaction binding the contract method 0x987fa2d5.
//
// Solidity: function validateTransactionProof(uint16 _srcChainId, address _dstAddress, uint256 _gasLimit, bytes32 _lookupHash, bytes32 _blockData, bytes _transactionProof) returns()
func (_UltraLightNode *UltraLightNodeSession) ValidateTransactionProof(_srcChainId uint16, _dstAddress common.Address, _gasLimit *big.Int, _lookupHash [32]byte, _blockData [32]byte, _transactionProof []byte) (*types.Transaction, error) {
	return _UltraLightNode.Contract.ValidateTransactionProof(&_UltraLightNode.TransactOpts, _srcChainId, _dstAddress, _gasLimit, _lookupHash, _blockData, _transactionProof)
}

// ValidateTransactionProof is a paid mutator transaction binding the contract method 0x987fa2d5.
//
// Solidity: function validateTransactionProof(uint16 _srcChainId, address _dstAddress, uint256 _gasLimit, bytes32 _lookupHash, bytes32 _blockData, bytes _transactionProof) returns()
func (_UltraLightNode *UltraLightNodeTransactorSession) ValidateTransactionProof(_srcChainId uint16, _dstAddress common.Address, _gasLimit *big.Int, _lookupHash [32]byte, _blockData [32]byte, _transactionProof []byte) (*types.Transaction, error) {
	return _UltraLightNode.Contract.ValidateTransactionProof(&_UltraLightNode.TransactOpts, _srcChainId, _dstAddress, _gasLimit, _lookupHash, _blockData, _transactionProof)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x07b18bde.
//
// Solidity: function withdrawNative(address _to, uint256 _amount) returns()
func (_UltraLightNode *UltraLightNodeTransactor) WithdrawNative(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UltraLightNode.contract.Transact(opts, "withdrawNative", _to, _amount)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x07b18bde.
//
// Solidity: function withdrawNative(address _to, uint256 _amount) returns()
func (_UltraLightNode *UltraLightNodeSession) WithdrawNative(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UltraLightNode.Contract.WithdrawNative(&_UltraLightNode.TransactOpts, _to, _amount)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x07b18bde.
//
// Solidity: function withdrawNative(address _to, uint256 _amount) returns()
func (_UltraLightNode *UltraLightNodeTransactorSession) WithdrawNative(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UltraLightNode.Contract.WithdrawNative(&_UltraLightNode.TransactOpts, _to, _amount)
}

// WithdrawZRO is a paid mutator transaction binding the contract method 0x8525b711.
//
// Solidity: function withdrawZRO(address _to, uint256 _amount) returns()
func (_UltraLightNode *UltraLightNodeTransactor) WithdrawZRO(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UltraLightNode.contract.Transact(opts, "withdrawZRO", _to, _amount)
}

// WithdrawZRO is a paid mutator transaction binding the contract method 0x8525b711.
//
// Solidity: function withdrawZRO(address _to, uint256 _amount) returns()
func (_UltraLightNode *UltraLightNodeSession) WithdrawZRO(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UltraLightNode.Contract.WithdrawZRO(&_UltraLightNode.TransactOpts, _to, _amount)
}

// WithdrawZRO is a paid mutator transaction binding the contract method 0x8525b711.
//
// Solidity: function withdrawZRO(address _to, uint256 _amount) returns()
func (_UltraLightNode *UltraLightNodeTransactorSession) WithdrawZRO(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _UltraLightNode.Contract.WithdrawZRO(&_UltraLightNode.TransactOpts, _to, _amount)
}

// UltraLightNodeAddInboundProofLibraryForChainIterator is returned from FilterAddInboundProofLibraryForChain and is used to iterate over the raw logs and unpacked data for AddInboundProofLibraryForChain events raised by the UltraLightNode contract.
type UltraLightNodeAddInboundProofLibraryForChainIterator struct {
	Event *UltraLightNodeAddInboundProofLibraryForChain // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UltraLightNodeAddInboundProofLibraryForChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UltraLightNodeAddInboundProofLibraryForChain)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UltraLightNodeAddInboundProofLibraryForChain)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UltraLightNodeAddInboundProofLibraryForChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UltraLightNodeAddInboundProofLibraryForChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UltraLightNodeAddInboundProofLibraryForChain represents a AddInboundProofLibraryForChain event raised by the UltraLightNode contract.
type UltraLightNodeAddInboundProofLibraryForChain struct {
	ChainId uint16
	Lib     common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddInboundProofLibraryForChain is a free log retrieval operation binding the contract event 0x802d55279d51813cb7a9a98e8fd2d7bec5346cb830901c11b85d1650cb857e9a.
//
// Solidity: event AddInboundProofLibraryForChain(uint16 indexed chainId, address lib)
func (_UltraLightNode *UltraLightNodeFilterer) FilterAddInboundProofLibraryForChain(opts *bind.FilterOpts, chainId []uint16) (*UltraLightNodeAddInboundProofLibraryForChainIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _UltraLightNode.contract.FilterLogs(opts, "AddInboundProofLibraryForChain", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &UltraLightNodeAddInboundProofLibraryForChainIterator{contract: _UltraLightNode.contract, event: "AddInboundProofLibraryForChain", logs: logs, sub: sub}, nil
}

// WatchAddInboundProofLibraryForChain is a free log subscription operation binding the contract event 0x802d55279d51813cb7a9a98e8fd2d7bec5346cb830901c11b85d1650cb857e9a.
//
// Solidity: event AddInboundProofLibraryForChain(uint16 indexed chainId, address lib)
func (_UltraLightNode *UltraLightNodeFilterer) WatchAddInboundProofLibraryForChain(opts *bind.WatchOpts, sink chan<- *UltraLightNodeAddInboundProofLibraryForChain, chainId []uint16) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _UltraLightNode.contract.WatchLogs(opts, "AddInboundProofLibraryForChain", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UltraLightNodeAddInboundProofLibraryForChain)
				if err := _UltraLightNode.contract.UnpackLog(event, "AddInboundProofLibraryForChain", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddInboundProofLibraryForChain is a log parse operation binding the contract event 0x802d55279d51813cb7a9a98e8fd2d7bec5346cb830901c11b85d1650cb857e9a.
//
// Solidity: event AddInboundProofLibraryForChain(uint16 indexed chainId, address lib)
func (_UltraLightNode *UltraLightNodeFilterer) ParseAddInboundProofLibraryForChain(log types.Log) (*UltraLightNodeAddInboundProofLibraryForChain, error) {
	event := new(UltraLightNodeAddInboundProofLibraryForChain)
	if err := _UltraLightNode.contract.UnpackLog(event, "AddInboundProofLibraryForChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UltraLightNodeAppConfigUpdatedIterator is returned from FilterAppConfigUpdated and is used to iterate over the raw logs and unpacked data for AppConfigUpdated events raised by the UltraLightNode contract.
type UltraLightNodeAppConfigUpdatedIterator struct {
	Event *UltraLightNodeAppConfigUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UltraLightNodeAppConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UltraLightNodeAppConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UltraLightNodeAppConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UltraLightNodeAppConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UltraLightNodeAppConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UltraLightNodeAppConfigUpdated represents a AppConfigUpdated event raised by the UltraLightNode contract.
type UltraLightNodeAppConfigUpdated struct {
	UserApplication common.Address
	ConfigType      *big.Int
	NewConfig       []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAppConfigUpdated is a free log retrieval operation binding the contract event 0xfc01bf86212a14151d51d1be5c2ac64d67d5ec823dfc6f53298d7ce3f3d3d252.
//
// Solidity: event AppConfigUpdated(address indexed userApplication, uint256 indexed configType, bytes newConfig)
func (_UltraLightNode *UltraLightNodeFilterer) FilterAppConfigUpdated(opts *bind.FilterOpts, userApplication []common.Address, configType []*big.Int) (*UltraLightNodeAppConfigUpdatedIterator, error) {

	var userApplicationRule []interface{}
	for _, userApplicationItem := range userApplication {
		userApplicationRule = append(userApplicationRule, userApplicationItem)
	}
	var configTypeRule []interface{}
	for _, configTypeItem := range configType {
		configTypeRule = append(configTypeRule, configTypeItem)
	}

	logs, sub, err := _UltraLightNode.contract.FilterLogs(opts, "AppConfigUpdated", userApplicationRule, configTypeRule)
	if err != nil {
		return nil, err
	}
	return &UltraLightNodeAppConfigUpdatedIterator{contract: _UltraLightNode.contract, event: "AppConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchAppConfigUpdated is a free log subscription operation binding the contract event 0xfc01bf86212a14151d51d1be5c2ac64d67d5ec823dfc6f53298d7ce3f3d3d252.
//
// Solidity: event AppConfigUpdated(address indexed userApplication, uint256 indexed configType, bytes newConfig)
func (_UltraLightNode *UltraLightNodeFilterer) WatchAppConfigUpdated(opts *bind.WatchOpts, sink chan<- *UltraLightNodeAppConfigUpdated, userApplication []common.Address, configType []*big.Int) (event.Subscription, error) {

	var userApplicationRule []interface{}
	for _, userApplicationItem := range userApplication {
		userApplicationRule = append(userApplicationRule, userApplicationItem)
	}
	var configTypeRule []interface{}
	for _, configTypeItem := range configType {
		configTypeRule = append(configTypeRule, configTypeItem)
	}

	logs, sub, err := _UltraLightNode.contract.WatchLogs(opts, "AppConfigUpdated", userApplicationRule, configTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UltraLightNodeAppConfigUpdated)
				if err := _UltraLightNode.contract.UnpackLog(event, "AppConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppConfigUpdated is a log parse operation binding the contract event 0xfc01bf86212a14151d51d1be5c2ac64d67d5ec823dfc6f53298d7ce3f3d3d252.
//
// Solidity: event AppConfigUpdated(address indexed userApplication, uint256 indexed configType, bytes newConfig)
func (_UltraLightNode *UltraLightNodeFilterer) ParseAppConfigUpdated(log types.Log) (*UltraLightNodeAppConfigUpdated, error) {
	event := new(UltraLightNodeAppConfigUpdated)
	if err := _UltraLightNode.contract.UnpackLog(event, "AppConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UltraLightNodeEnableSupportedOutboundProofIterator is returned from FilterEnableSupportedOutboundProof and is used to iterate over the raw logs and unpacked data for EnableSupportedOutboundProof events raised by the UltraLightNode contract.
type UltraLightNodeEnableSupportedOutboundProofIterator struct {
	Event *UltraLightNodeEnableSupportedOutboundProof // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UltraLightNodeEnableSupportedOutboundProofIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UltraLightNodeEnableSupportedOutboundProof)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UltraLightNodeEnableSupportedOutboundProof)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UltraLightNodeEnableSupportedOutboundProofIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UltraLightNodeEnableSupportedOutboundProofIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UltraLightNodeEnableSupportedOutboundProof represents a EnableSupportedOutboundProof event raised by the UltraLightNode contract.
type UltraLightNodeEnableSupportedOutboundProof struct {
	ChainId   uint16
	ProofType uint16
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEnableSupportedOutboundProof is a free log retrieval operation binding the contract event 0xec23bee6f88cfecebb09d6aaaed66f0ce110debc1f61117c8270a7116597df9a.
//
// Solidity: event EnableSupportedOutboundProof(uint16 indexed chainId, uint16 proofType)
func (_UltraLightNode *UltraLightNodeFilterer) FilterEnableSupportedOutboundProof(opts *bind.FilterOpts, chainId []uint16) (*UltraLightNodeEnableSupportedOutboundProofIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _UltraLightNode.contract.FilterLogs(opts, "EnableSupportedOutboundProof", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &UltraLightNodeEnableSupportedOutboundProofIterator{contract: _UltraLightNode.contract, event: "EnableSupportedOutboundProof", logs: logs, sub: sub}, nil
}

// WatchEnableSupportedOutboundProof is a free log subscription operation binding the contract event 0xec23bee6f88cfecebb09d6aaaed66f0ce110debc1f61117c8270a7116597df9a.
//
// Solidity: event EnableSupportedOutboundProof(uint16 indexed chainId, uint16 proofType)
func (_UltraLightNode *UltraLightNodeFilterer) WatchEnableSupportedOutboundProof(opts *bind.WatchOpts, sink chan<- *UltraLightNodeEnableSupportedOutboundProof, chainId []uint16) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _UltraLightNode.contract.WatchLogs(opts, "EnableSupportedOutboundProof", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UltraLightNodeEnableSupportedOutboundProof)
				if err := _UltraLightNode.contract.UnpackLog(event, "EnableSupportedOutboundProof", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEnableSupportedOutboundProof is a log parse operation binding the contract event 0xec23bee6f88cfecebb09d6aaaed66f0ce110debc1f61117c8270a7116597df9a.
//
// Solidity: event EnableSupportedOutboundProof(uint16 indexed chainId, uint16 proofType)
func (_UltraLightNode *UltraLightNodeFilterer) ParseEnableSupportedOutboundProof(log types.Log) (*UltraLightNodeEnableSupportedOutboundProof, error) {
	event := new(UltraLightNodeEnableSupportedOutboundProof)
	if err := _UltraLightNode.contract.UnpackLog(event, "EnableSupportedOutboundProof", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UltraLightNodeHashReceivedIterator is returned from FilterHashReceived and is used to iterate over the raw logs and unpacked data for HashReceived events raised by the UltraLightNode contract.
type UltraLightNodeHashReceivedIterator struct {
	Event *UltraLightNodeHashReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UltraLightNodeHashReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UltraLightNodeHashReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UltraLightNodeHashReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UltraLightNodeHashReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UltraLightNodeHashReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UltraLightNodeHashReceived represents a HashReceived event raised by the UltraLightNode contract.
type UltraLightNodeHashReceived struct {
	SrcChainId    uint16
	Oracle        common.Address
	LookupHash    [32]byte
	BlockData     [32]byte
	Confirmations *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterHashReceived is a free log retrieval operation binding the contract event 0x74bbc026808dcba59692d6a8bb20596849ca718e10e2432c6cdf48af865bc5d9.
//
// Solidity: event HashReceived(uint16 indexed srcChainId, address indexed oracle, bytes32 lookupHash, bytes32 blockData, uint256 confirmations)
func (_UltraLightNode *UltraLightNodeFilterer) FilterHashReceived(opts *bind.FilterOpts, srcChainId []uint16, oracle []common.Address) (*UltraLightNodeHashReceivedIterator, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _UltraLightNode.contract.FilterLogs(opts, "HashReceived", srcChainIdRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return &UltraLightNodeHashReceivedIterator{contract: _UltraLightNode.contract, event: "HashReceived", logs: logs, sub: sub}, nil
}

// WatchHashReceived is a free log subscription operation binding the contract event 0x74bbc026808dcba59692d6a8bb20596849ca718e10e2432c6cdf48af865bc5d9.
//
// Solidity: event HashReceived(uint16 indexed srcChainId, address indexed oracle, bytes32 lookupHash, bytes32 blockData, uint256 confirmations)
func (_UltraLightNode *UltraLightNodeFilterer) WatchHashReceived(opts *bind.WatchOpts, sink chan<- *UltraLightNodeHashReceived, srcChainId []uint16, oracle []common.Address) (event.Subscription, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _UltraLightNode.contract.WatchLogs(opts, "HashReceived", srcChainIdRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UltraLightNodeHashReceived)
				if err := _UltraLightNode.contract.UnpackLog(event, "HashReceived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseHashReceived is a log parse operation binding the contract event 0x74bbc026808dcba59692d6a8bb20596849ca718e10e2432c6cdf48af865bc5d9.
//
// Solidity: event HashReceived(uint16 indexed srcChainId, address indexed oracle, bytes32 lookupHash, bytes32 blockData, uint256 confirmations)
func (_UltraLightNode *UltraLightNodeFilterer) ParseHashReceived(log types.Log) (*UltraLightNodeHashReceived, error) {
	event := new(UltraLightNodeHashReceived)
	if err := _UltraLightNode.contract.UnpackLog(event, "HashReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UltraLightNodeInvalidDstIterator is returned from FilterInvalidDst and is used to iterate over the raw logs and unpacked data for InvalidDst events raised by the UltraLightNode contract.
type UltraLightNodeInvalidDstIterator struct {
	Event *UltraLightNodeInvalidDst // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UltraLightNodeInvalidDstIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UltraLightNodeInvalidDst)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UltraLightNodeInvalidDst)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UltraLightNodeInvalidDstIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UltraLightNodeInvalidDstIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UltraLightNodeInvalidDst represents a InvalidDst event raised by the UltraLightNode contract.
type UltraLightNodeInvalidDst struct {
	SrcChainId  uint16
	SrcAddress  []byte
	DstAddress  common.Address
	Nonce       uint64
	PayloadHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidDst is a free log retrieval operation binding the contract event 0xa2786598bd84ae4a299103996359e6cb4333404583256079dfc279386baf5832.
//
// Solidity: event InvalidDst(uint16 indexed srcChainId, bytes srcAddress, address indexed dstAddress, uint64 nonce, bytes32 payloadHash)
func (_UltraLightNode *UltraLightNodeFilterer) FilterInvalidDst(opts *bind.FilterOpts, srcChainId []uint16, dstAddress []common.Address) (*UltraLightNodeInvalidDstIterator, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}

	var dstAddressRule []interface{}
	for _, dstAddressItem := range dstAddress {
		dstAddressRule = append(dstAddressRule, dstAddressItem)
	}

	logs, sub, err := _UltraLightNode.contract.FilterLogs(opts, "InvalidDst", srcChainIdRule, dstAddressRule)
	if err != nil {
		return nil, err
	}
	return &UltraLightNodeInvalidDstIterator{contract: _UltraLightNode.contract, event: "InvalidDst", logs: logs, sub: sub}, nil
}

// WatchInvalidDst is a free log subscription operation binding the contract event 0xa2786598bd84ae4a299103996359e6cb4333404583256079dfc279386baf5832.
//
// Solidity: event InvalidDst(uint16 indexed srcChainId, bytes srcAddress, address indexed dstAddress, uint64 nonce, bytes32 payloadHash)
func (_UltraLightNode *UltraLightNodeFilterer) WatchInvalidDst(opts *bind.WatchOpts, sink chan<- *UltraLightNodeInvalidDst, srcChainId []uint16, dstAddress []common.Address) (event.Subscription, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}

	var dstAddressRule []interface{}
	for _, dstAddressItem := range dstAddress {
		dstAddressRule = append(dstAddressRule, dstAddressItem)
	}

	logs, sub, err := _UltraLightNode.contract.WatchLogs(opts, "InvalidDst", srcChainIdRule, dstAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UltraLightNodeInvalidDst)
				if err := _UltraLightNode.contract.UnpackLog(event, "InvalidDst", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvalidDst is a log parse operation binding the contract event 0xa2786598bd84ae4a299103996359e6cb4333404583256079dfc279386baf5832.
//
// Solidity: event InvalidDst(uint16 indexed srcChainId, bytes srcAddress, address indexed dstAddress, uint64 nonce, bytes32 payloadHash)
func (_UltraLightNode *UltraLightNodeFilterer) ParseInvalidDst(log types.Log) (*UltraLightNodeInvalidDst, error) {
	event := new(UltraLightNodeInvalidDst)
	if err := _UltraLightNode.contract.UnpackLog(event, "InvalidDst", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UltraLightNodeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UltraLightNode contract.
type UltraLightNodeOwnershipTransferredIterator struct {
	Event *UltraLightNodeOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UltraLightNodeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UltraLightNodeOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UltraLightNodeOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UltraLightNodeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UltraLightNodeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UltraLightNodeOwnershipTransferred represents a OwnershipTransferred event raised by the UltraLightNode contract.
type UltraLightNodeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UltraLightNode *UltraLightNodeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UltraLightNodeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UltraLightNode.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UltraLightNodeOwnershipTransferredIterator{contract: _UltraLightNode.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UltraLightNode *UltraLightNodeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UltraLightNodeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UltraLightNode.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UltraLightNodeOwnershipTransferred)
				if err := _UltraLightNode.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UltraLightNode *UltraLightNodeFilterer) ParseOwnershipTransferred(log types.Log) (*UltraLightNodeOwnershipTransferred, error) {
	event := new(UltraLightNodeOwnershipTransferred)
	if err := _UltraLightNode.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UltraLightNodePacketIterator is returned from FilterPacket and is used to iterate over the raw logs and unpacked data for Packet events raised by the UltraLightNode contract.
type UltraLightNodePacketIterator struct {
	Event *UltraLightNodePacket // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UltraLightNodePacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UltraLightNodePacket)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UltraLightNodePacket)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UltraLightNodePacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UltraLightNodePacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UltraLightNodePacket represents a Packet event raised by the UltraLightNode contract.
type UltraLightNodePacket struct {
	Payload []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPacket is a free log retrieval operation binding the contract event 0xe9bded5f24a4168e4f3bf44e00298c993b22376aad8c58c7dda9718a54cbea82.
//
// Solidity: event Packet(bytes payload)
func (_UltraLightNode *UltraLightNodeFilterer) FilterPacket(opts *bind.FilterOpts) (*UltraLightNodePacketIterator, error) {

	logs, sub, err := _UltraLightNode.contract.FilterLogs(opts, "Packet")
	if err != nil {
		return nil, err
	}
	return &UltraLightNodePacketIterator{contract: _UltraLightNode.contract, event: "Packet", logs: logs, sub: sub}, nil
}

// WatchPacket is a free log subscription operation binding the contract event 0xe9bded5f24a4168e4f3bf44e00298c993b22376aad8c58c7dda9718a54cbea82.
//
// Solidity: event Packet(bytes payload)
func (_UltraLightNode *UltraLightNodeFilterer) WatchPacket(opts *bind.WatchOpts, sink chan<- *UltraLightNodePacket) (event.Subscription, error) {

	logs, sub, err := _UltraLightNode.contract.WatchLogs(opts, "Packet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UltraLightNodePacket)
				if err := _UltraLightNode.contract.UnpackLog(event, "Packet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacket is a log parse operation binding the contract event 0xe9bded5f24a4168e4f3bf44e00298c993b22376aad8c58c7dda9718a54cbea82.
//
// Solidity: event Packet(bytes payload)
func (_UltraLightNode *UltraLightNodeFilterer) ParsePacket(log types.Log) (*UltraLightNodePacket, error) {
	event := new(UltraLightNodePacket)
	if err := _UltraLightNode.contract.UnpackLog(event, "Packet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UltraLightNodePacketReceivedIterator is returned from FilterPacketReceived and is used to iterate over the raw logs and unpacked data for PacketReceived events raised by the UltraLightNode contract.
type UltraLightNodePacketReceivedIterator struct {
	Event *UltraLightNodePacketReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UltraLightNodePacketReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UltraLightNodePacketReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UltraLightNodePacketReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UltraLightNodePacketReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UltraLightNodePacketReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UltraLightNodePacketReceived represents a PacketReceived event raised by the UltraLightNode contract.
type UltraLightNodePacketReceived struct {
	SrcChainId  uint16
	SrcAddress  []byte
	DstAddress  common.Address
	Nonce       uint64
	PayloadHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPacketReceived is a free log retrieval operation binding the contract event 0x2bd2d8a84b748439fd50d79a49502b4eb5faa25b864da6a9ab5c150704be9a4d.
//
// Solidity: event PacketReceived(uint16 indexed srcChainId, bytes srcAddress, address indexed dstAddress, uint64 nonce, bytes32 payloadHash)
func (_UltraLightNode *UltraLightNodeFilterer) FilterPacketReceived(opts *bind.FilterOpts, srcChainId []uint16, dstAddress []common.Address) (*UltraLightNodePacketReceivedIterator, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}

	var dstAddressRule []interface{}
	for _, dstAddressItem := range dstAddress {
		dstAddressRule = append(dstAddressRule, dstAddressItem)
	}

	logs, sub, err := _UltraLightNode.contract.FilterLogs(opts, "PacketReceived", srcChainIdRule, dstAddressRule)
	if err != nil {
		return nil, err
	}
	return &UltraLightNodePacketReceivedIterator{contract: _UltraLightNode.contract, event: "PacketReceived", logs: logs, sub: sub}, nil
}

// WatchPacketReceived is a free log subscription operation binding the contract event 0x2bd2d8a84b748439fd50d79a49502b4eb5faa25b864da6a9ab5c150704be9a4d.
//
// Solidity: event PacketReceived(uint16 indexed srcChainId, bytes srcAddress, address indexed dstAddress, uint64 nonce, bytes32 payloadHash)
func (_UltraLightNode *UltraLightNodeFilterer) WatchPacketReceived(opts *bind.WatchOpts, sink chan<- *UltraLightNodePacketReceived, srcChainId []uint16, dstAddress []common.Address) (event.Subscription, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}

	var dstAddressRule []interface{}
	for _, dstAddressItem := range dstAddress {
		dstAddressRule = append(dstAddressRule, dstAddressItem)
	}

	logs, sub, err := _UltraLightNode.contract.WatchLogs(opts, "PacketReceived", srcChainIdRule, dstAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UltraLightNodePacketReceived)
				if err := _UltraLightNode.contract.UnpackLog(event, "PacketReceived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacketReceived is a log parse operation binding the contract event 0x2bd2d8a84b748439fd50d79a49502b4eb5faa25b864da6a9ab5c150704be9a4d.
//
// Solidity: event PacketReceived(uint16 indexed srcChainId, bytes srcAddress, address indexed dstAddress, uint64 nonce, bytes32 payloadHash)
func (_UltraLightNode *UltraLightNodeFilterer) ParsePacketReceived(log types.Log) (*UltraLightNodePacketReceived, error) {
	event := new(UltraLightNodePacketReceived)
	if err := _UltraLightNode.contract.UnpackLog(event, "PacketReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UltraLightNodeRelayerParamsIterator is returned from FilterRelayerParams and is used to iterate over the raw logs and unpacked data for RelayerParams events raised by the UltraLightNode contract.
type UltraLightNodeRelayerParamsIterator struct {
	Event *UltraLightNodeRelayerParams // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UltraLightNodeRelayerParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UltraLightNodeRelayerParams)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UltraLightNodeRelayerParams)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UltraLightNodeRelayerParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UltraLightNodeRelayerParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UltraLightNodeRelayerParams represents a RelayerParams event raised by the UltraLightNode contract.
type UltraLightNodeRelayerParams struct {
	AdapterParams     []byte
	OutboundProofType uint16
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRelayerParams is a free log retrieval operation binding the contract event 0xb0c632f55f1e1b3b2c3d82f41ee4716bb4c00f0f5d84cdafc141581bb8757a4f.
//
// Solidity: event RelayerParams(bytes adapterParams, uint16 outboundProofType)
func (_UltraLightNode *UltraLightNodeFilterer) FilterRelayerParams(opts *bind.FilterOpts) (*UltraLightNodeRelayerParamsIterator, error) {

	logs, sub, err := _UltraLightNode.contract.FilterLogs(opts, "RelayerParams")
	if err != nil {
		return nil, err
	}
	return &UltraLightNodeRelayerParamsIterator{contract: _UltraLightNode.contract, event: "RelayerParams", logs: logs, sub: sub}, nil
}

// WatchRelayerParams is a free log subscription operation binding the contract event 0xb0c632f55f1e1b3b2c3d82f41ee4716bb4c00f0f5d84cdafc141581bb8757a4f.
//
// Solidity: event RelayerParams(bytes adapterParams, uint16 outboundProofType)
func (_UltraLightNode *UltraLightNodeFilterer) WatchRelayerParams(opts *bind.WatchOpts, sink chan<- *UltraLightNodeRelayerParams) (event.Subscription, error) {

	logs, sub, err := _UltraLightNode.contract.WatchLogs(opts, "RelayerParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UltraLightNodeRelayerParams)
				if err := _UltraLightNode.contract.UnpackLog(event, "RelayerParams", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayerParams is a log parse operation binding the contract event 0xb0c632f55f1e1b3b2c3d82f41ee4716bb4c00f0f5d84cdafc141581bb8757a4f.
//
// Solidity: event RelayerParams(bytes adapterParams, uint16 outboundProofType)
func (_UltraLightNode *UltraLightNodeFilterer) ParseRelayerParams(log types.Log) (*UltraLightNodeRelayerParams, error) {
	event := new(UltraLightNodeRelayerParams)
	if err := _UltraLightNode.contract.UnpackLog(event, "RelayerParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UltraLightNodeSetChainAddressSizeIterator is returned from FilterSetChainAddressSize and is used to iterate over the raw logs and unpacked data for SetChainAddressSize events raised by the UltraLightNode contract.
type UltraLightNodeSetChainAddressSizeIterator struct {
	Event *UltraLightNodeSetChainAddressSize // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UltraLightNodeSetChainAddressSizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UltraLightNodeSetChainAddressSize)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UltraLightNodeSetChainAddressSize)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UltraLightNodeSetChainAddressSizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UltraLightNodeSetChainAddressSizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UltraLightNodeSetChainAddressSize represents a SetChainAddressSize event raised by the UltraLightNode contract.
type UltraLightNodeSetChainAddressSize struct {
	ChainId uint16
	Size    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetChainAddressSize is a free log retrieval operation binding the contract event 0x0611bb2107e385b79ec826fff8ecc1c1248a7aae3c875c96668f8cfbf1734220.
//
// Solidity: event SetChainAddressSize(uint16 indexed chainId, uint256 size)
func (_UltraLightNode *UltraLightNodeFilterer) FilterSetChainAddressSize(opts *bind.FilterOpts, chainId []uint16) (*UltraLightNodeSetChainAddressSizeIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _UltraLightNode.contract.FilterLogs(opts, "SetChainAddressSize", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &UltraLightNodeSetChainAddressSizeIterator{contract: _UltraLightNode.contract, event: "SetChainAddressSize", logs: logs, sub: sub}, nil
}

// WatchSetChainAddressSize is a free log subscription operation binding the contract event 0x0611bb2107e385b79ec826fff8ecc1c1248a7aae3c875c96668f8cfbf1734220.
//
// Solidity: event SetChainAddressSize(uint16 indexed chainId, uint256 size)
func (_UltraLightNode *UltraLightNodeFilterer) WatchSetChainAddressSize(opts *bind.WatchOpts, sink chan<- *UltraLightNodeSetChainAddressSize, chainId []uint16) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _UltraLightNode.contract.WatchLogs(opts, "SetChainAddressSize", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UltraLightNodeSetChainAddressSize)
				if err := _UltraLightNode.contract.UnpackLog(event, "SetChainAddressSize", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetChainAddressSize is a log parse operation binding the contract event 0x0611bb2107e385b79ec826fff8ecc1c1248a7aae3c875c96668f8cfbf1734220.
//
// Solidity: event SetChainAddressSize(uint16 indexed chainId, uint256 size)
func (_UltraLightNode *UltraLightNodeFilterer) ParseSetChainAddressSize(log types.Log) (*UltraLightNodeSetChainAddressSize, error) {
	event := new(UltraLightNodeSetChainAddressSize)
	if err := _UltraLightNode.contract.UnpackLog(event, "SetChainAddressSize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UltraLightNodeSetDefaultAdapterParamsForChainIdIterator is returned from FilterSetDefaultAdapterParamsForChainId and is used to iterate over the raw logs and unpacked data for SetDefaultAdapterParamsForChainId events raised by the UltraLightNode contract.
type UltraLightNodeSetDefaultAdapterParamsForChainIdIterator struct {
	Event *UltraLightNodeSetDefaultAdapterParamsForChainId // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UltraLightNodeSetDefaultAdapterParamsForChainIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UltraLightNodeSetDefaultAdapterParamsForChainId)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UltraLightNodeSetDefaultAdapterParamsForChainId)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UltraLightNodeSetDefaultAdapterParamsForChainIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UltraLightNodeSetDefaultAdapterParamsForChainIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UltraLightNodeSetDefaultAdapterParamsForChainId represents a SetDefaultAdapterParamsForChainId event raised by the UltraLightNode contract.
type UltraLightNodeSetDefaultAdapterParamsForChainId struct {
	ChainId       uint16
	ProofType     uint16
	AdapterParams []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetDefaultAdapterParamsForChainId is a free log retrieval operation binding the contract event 0x4a5695eee2a74d548d5f5c485a3de99ace99e3b664c8e30a90f49be6ebb54932.
//
// Solidity: event SetDefaultAdapterParamsForChainId(uint16 indexed chainId, uint16 indexed proofType, bytes adapterParams)
func (_UltraLightNode *UltraLightNodeFilterer) FilterSetDefaultAdapterParamsForChainId(opts *bind.FilterOpts, chainId []uint16, proofType []uint16) (*UltraLightNodeSetDefaultAdapterParamsForChainIdIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var proofTypeRule []interface{}
	for _, proofTypeItem := range proofType {
		proofTypeRule = append(proofTypeRule, proofTypeItem)
	}

	logs, sub, err := _UltraLightNode.contract.FilterLogs(opts, "SetDefaultAdapterParamsForChainId", chainIdRule, proofTypeRule)
	if err != nil {
		return nil, err
	}
	return &UltraLightNodeSetDefaultAdapterParamsForChainIdIterator{contract: _UltraLightNode.contract, event: "SetDefaultAdapterParamsForChainId", logs: logs, sub: sub}, nil
}

// WatchSetDefaultAdapterParamsForChainId is a free log subscription operation binding the contract event 0x4a5695eee2a74d548d5f5c485a3de99ace99e3b664c8e30a90f49be6ebb54932.
//
// Solidity: event SetDefaultAdapterParamsForChainId(uint16 indexed chainId, uint16 indexed proofType, bytes adapterParams)
func (_UltraLightNode *UltraLightNodeFilterer) WatchSetDefaultAdapterParamsForChainId(opts *bind.WatchOpts, sink chan<- *UltraLightNodeSetDefaultAdapterParamsForChainId, chainId []uint16, proofType []uint16) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var proofTypeRule []interface{}
	for _, proofTypeItem := range proofType {
		proofTypeRule = append(proofTypeRule, proofTypeItem)
	}

	logs, sub, err := _UltraLightNode.contract.WatchLogs(opts, "SetDefaultAdapterParamsForChainId", chainIdRule, proofTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UltraLightNodeSetDefaultAdapterParamsForChainId)
				if err := _UltraLightNode.contract.UnpackLog(event, "SetDefaultAdapterParamsForChainId", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetDefaultAdapterParamsForChainId is a log parse operation binding the contract event 0x4a5695eee2a74d548d5f5c485a3de99ace99e3b664c8e30a90f49be6ebb54932.
//
// Solidity: event SetDefaultAdapterParamsForChainId(uint16 indexed chainId, uint16 indexed proofType, bytes adapterParams)
func (_UltraLightNode *UltraLightNodeFilterer) ParseSetDefaultAdapterParamsForChainId(log types.Log) (*UltraLightNodeSetDefaultAdapterParamsForChainId, error) {
	event := new(UltraLightNodeSetDefaultAdapterParamsForChainId)
	if err := _UltraLightNode.contract.UnpackLog(event, "SetDefaultAdapterParamsForChainId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UltraLightNodeSetDefaultConfigForChainIdIterator is returned from FilterSetDefaultConfigForChainId and is used to iterate over the raw logs and unpacked data for SetDefaultConfigForChainId events raised by the UltraLightNode contract.
type UltraLightNodeSetDefaultConfigForChainIdIterator struct {
	Event *UltraLightNodeSetDefaultConfigForChainId // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UltraLightNodeSetDefaultConfigForChainIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UltraLightNodeSetDefaultConfigForChainId)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UltraLightNodeSetDefaultConfigForChainId)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UltraLightNodeSetDefaultConfigForChainIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UltraLightNodeSetDefaultConfigForChainIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UltraLightNodeSetDefaultConfigForChainId represents a SetDefaultConfigForChainId event raised by the UltraLightNode contract.
type UltraLightNodeSetDefaultConfigForChainId struct {
	ChainId              uint16
	InboundProofLib      uint16
	InboundBlockConfirm  uint64
	Relayer              common.Address
	OutboundProofType    uint16
	OutboundBlockConfirm uint64
	Oracle               common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetDefaultConfigForChainId is a free log retrieval operation binding the contract event 0x5a76432853a0871c4e780def7f3ffc7912339b53f022ac31127fe5ff84a36fa1.
//
// Solidity: event SetDefaultConfigForChainId(uint16 indexed chainId, uint16 inboundProofLib, uint64 inboundBlockConfirm, address relayer, uint16 outboundProofType, uint64 outboundBlockConfirm, address oracle)
func (_UltraLightNode *UltraLightNodeFilterer) FilterSetDefaultConfigForChainId(opts *bind.FilterOpts, chainId []uint16) (*UltraLightNodeSetDefaultConfigForChainIdIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _UltraLightNode.contract.FilterLogs(opts, "SetDefaultConfigForChainId", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &UltraLightNodeSetDefaultConfigForChainIdIterator{contract: _UltraLightNode.contract, event: "SetDefaultConfigForChainId", logs: logs, sub: sub}, nil
}

// WatchSetDefaultConfigForChainId is a free log subscription operation binding the contract event 0x5a76432853a0871c4e780def7f3ffc7912339b53f022ac31127fe5ff84a36fa1.
//
// Solidity: event SetDefaultConfigForChainId(uint16 indexed chainId, uint16 inboundProofLib, uint64 inboundBlockConfirm, address relayer, uint16 outboundProofType, uint64 outboundBlockConfirm, address oracle)
func (_UltraLightNode *UltraLightNodeFilterer) WatchSetDefaultConfigForChainId(opts *bind.WatchOpts, sink chan<- *UltraLightNodeSetDefaultConfigForChainId, chainId []uint16) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _UltraLightNode.contract.WatchLogs(opts, "SetDefaultConfigForChainId", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UltraLightNodeSetDefaultConfigForChainId)
				if err := _UltraLightNode.contract.UnpackLog(event, "SetDefaultConfigForChainId", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetDefaultConfigForChainId is a log parse operation binding the contract event 0x5a76432853a0871c4e780def7f3ffc7912339b53f022ac31127fe5ff84a36fa1.
//
// Solidity: event SetDefaultConfigForChainId(uint16 indexed chainId, uint16 inboundProofLib, uint64 inboundBlockConfirm, address relayer, uint16 outboundProofType, uint64 outboundBlockConfirm, address oracle)
func (_UltraLightNode *UltraLightNodeFilterer) ParseSetDefaultConfigForChainId(log types.Log) (*UltraLightNodeSetDefaultConfigForChainId, error) {
	event := new(UltraLightNodeSetDefaultConfigForChainId)
	if err := _UltraLightNode.contract.UnpackLog(event, "SetDefaultConfigForChainId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UltraLightNodeSetLayerZeroTokenIterator is returned from FilterSetLayerZeroToken and is used to iterate over the raw logs and unpacked data for SetLayerZeroToken events raised by the UltraLightNode contract.
type UltraLightNodeSetLayerZeroTokenIterator struct {
	Event *UltraLightNodeSetLayerZeroToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UltraLightNodeSetLayerZeroTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UltraLightNodeSetLayerZeroToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UltraLightNodeSetLayerZeroToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UltraLightNodeSetLayerZeroTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UltraLightNodeSetLayerZeroTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UltraLightNodeSetLayerZeroToken represents a SetLayerZeroToken event raised by the UltraLightNode contract.
type UltraLightNodeSetLayerZeroToken struct {
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetLayerZeroToken is a free log retrieval operation binding the contract event 0x33d644987381deff4408951d55afa136f124e22a7810b163b2aaa3ebef770f64.
//
// Solidity: event SetLayerZeroToken(address indexed tokenAddress)
func (_UltraLightNode *UltraLightNodeFilterer) FilterSetLayerZeroToken(opts *bind.FilterOpts, tokenAddress []common.Address) (*UltraLightNodeSetLayerZeroTokenIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _UltraLightNode.contract.FilterLogs(opts, "SetLayerZeroToken", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &UltraLightNodeSetLayerZeroTokenIterator{contract: _UltraLightNode.contract, event: "SetLayerZeroToken", logs: logs, sub: sub}, nil
}

// WatchSetLayerZeroToken is a free log subscription operation binding the contract event 0x33d644987381deff4408951d55afa136f124e22a7810b163b2aaa3ebef770f64.
//
// Solidity: event SetLayerZeroToken(address indexed tokenAddress)
func (_UltraLightNode *UltraLightNodeFilterer) WatchSetLayerZeroToken(opts *bind.WatchOpts, sink chan<- *UltraLightNodeSetLayerZeroToken, tokenAddress []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _UltraLightNode.contract.WatchLogs(opts, "SetLayerZeroToken", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UltraLightNodeSetLayerZeroToken)
				if err := _UltraLightNode.contract.UnpackLog(event, "SetLayerZeroToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetLayerZeroToken is a log parse operation binding the contract event 0x33d644987381deff4408951d55afa136f124e22a7810b163b2aaa3ebef770f64.
//
// Solidity: event SetLayerZeroToken(address indexed tokenAddress)
func (_UltraLightNode *UltraLightNodeFilterer) ParseSetLayerZeroToken(log types.Log) (*UltraLightNodeSetLayerZeroToken, error) {
	event := new(UltraLightNodeSetLayerZeroToken)
	if err := _UltraLightNode.contract.UnpackLog(event, "SetLayerZeroToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UltraLightNodeSetRemoteUlnIterator is returned from FilterSetRemoteUln and is used to iterate over the raw logs and unpacked data for SetRemoteUln events raised by the UltraLightNode contract.
type UltraLightNodeSetRemoteUlnIterator struct {
	Event *UltraLightNodeSetRemoteUln // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UltraLightNodeSetRemoteUlnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UltraLightNodeSetRemoteUln)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UltraLightNodeSetRemoteUln)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UltraLightNodeSetRemoteUlnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UltraLightNodeSetRemoteUlnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UltraLightNodeSetRemoteUln represents a SetRemoteUln event raised by the UltraLightNode contract.
type UltraLightNodeSetRemoteUln struct {
	ChainId uint16
	Uln     [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetRemoteUln is a free log retrieval operation binding the contract event 0x0dad975e1d2fbe771c95cdcc7be9a1e61181de7173abe0a32b8f8f83140873e5.
//
// Solidity: event SetRemoteUln(uint16 indexed chainId, bytes32 uln)
func (_UltraLightNode *UltraLightNodeFilterer) FilterSetRemoteUln(opts *bind.FilterOpts, chainId []uint16) (*UltraLightNodeSetRemoteUlnIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _UltraLightNode.contract.FilterLogs(opts, "SetRemoteUln", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &UltraLightNodeSetRemoteUlnIterator{contract: _UltraLightNode.contract, event: "SetRemoteUln", logs: logs, sub: sub}, nil
}

// WatchSetRemoteUln is a free log subscription operation binding the contract event 0x0dad975e1d2fbe771c95cdcc7be9a1e61181de7173abe0a32b8f8f83140873e5.
//
// Solidity: event SetRemoteUln(uint16 indexed chainId, bytes32 uln)
func (_UltraLightNode *UltraLightNodeFilterer) WatchSetRemoteUln(opts *bind.WatchOpts, sink chan<- *UltraLightNodeSetRemoteUln, chainId []uint16) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _UltraLightNode.contract.WatchLogs(opts, "SetRemoteUln", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UltraLightNodeSetRemoteUln)
				if err := _UltraLightNode.contract.UnpackLog(event, "SetRemoteUln", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetRemoteUln is a log parse operation binding the contract event 0x0dad975e1d2fbe771c95cdcc7be9a1e61181de7173abe0a32b8f8f83140873e5.
//
// Solidity: event SetRemoteUln(uint16 indexed chainId, bytes32 uln)
func (_UltraLightNode *UltraLightNodeFilterer) ParseSetRemoteUln(log types.Log) (*UltraLightNodeSetRemoteUln, error) {
	event := new(UltraLightNodeSetRemoteUln)
	if err := _UltraLightNode.contract.UnpackLog(event, "SetRemoteUln", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UltraLightNodeSetTreasuryIterator is returned from FilterSetTreasury and is used to iterate over the raw logs and unpacked data for SetTreasury events raised by the UltraLightNode contract.
type UltraLightNodeSetTreasuryIterator struct {
	Event *UltraLightNodeSetTreasury // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UltraLightNodeSetTreasuryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UltraLightNodeSetTreasury)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UltraLightNodeSetTreasury)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UltraLightNodeSetTreasuryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UltraLightNodeSetTreasuryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UltraLightNodeSetTreasury represents a SetTreasury event raised by the UltraLightNode contract.
type UltraLightNodeSetTreasury struct {
	TreasuryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetTreasury is a free log retrieval operation binding the contract event 0xcb7ef3e545f5cdb893f5c568ba710fe08f336375a2d9fd66e161033f8fc09ef3.
//
// Solidity: event SetTreasury(address indexed treasuryAddress)
func (_UltraLightNode *UltraLightNodeFilterer) FilterSetTreasury(opts *bind.FilterOpts, treasuryAddress []common.Address) (*UltraLightNodeSetTreasuryIterator, error) {

	var treasuryAddressRule []interface{}
	for _, treasuryAddressItem := range treasuryAddress {
		treasuryAddressRule = append(treasuryAddressRule, treasuryAddressItem)
	}

	logs, sub, err := _UltraLightNode.contract.FilterLogs(opts, "SetTreasury", treasuryAddressRule)
	if err != nil {
		return nil, err
	}
	return &UltraLightNodeSetTreasuryIterator{contract: _UltraLightNode.contract, event: "SetTreasury", logs: logs, sub: sub}, nil
}

// WatchSetTreasury is a free log subscription operation binding the contract event 0xcb7ef3e545f5cdb893f5c568ba710fe08f336375a2d9fd66e161033f8fc09ef3.
//
// Solidity: event SetTreasury(address indexed treasuryAddress)
func (_UltraLightNode *UltraLightNodeFilterer) WatchSetTreasury(opts *bind.WatchOpts, sink chan<- *UltraLightNodeSetTreasury, treasuryAddress []common.Address) (event.Subscription, error) {

	var treasuryAddressRule []interface{}
	for _, treasuryAddressItem := range treasuryAddress {
		treasuryAddressRule = append(treasuryAddressRule, treasuryAddressItem)
	}

	logs, sub, err := _UltraLightNode.contract.WatchLogs(opts, "SetTreasury", treasuryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UltraLightNodeSetTreasury)
				if err := _UltraLightNode.contract.UnpackLog(event, "SetTreasury", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTreasury is a log parse operation binding the contract event 0xcb7ef3e545f5cdb893f5c568ba710fe08f336375a2d9fd66e161033f8fc09ef3.
//
// Solidity: event SetTreasury(address indexed treasuryAddress)
func (_UltraLightNode *UltraLightNodeFilterer) ParseSetTreasury(log types.Log) (*UltraLightNodeSetTreasury, error) {
	event := new(UltraLightNodeSetTreasury)
	if err := _UltraLightNode.contract.UnpackLog(event, "SetTreasury", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UltraLightNodeWithdrawNativeIterator is returned from FilterWithdrawNative and is used to iterate over the raw logs and unpacked data for WithdrawNative events raised by the UltraLightNode contract.
type UltraLightNodeWithdrawNativeIterator struct {
	Event *UltraLightNodeWithdrawNative // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UltraLightNodeWithdrawNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UltraLightNodeWithdrawNative)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UltraLightNodeWithdrawNative)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UltraLightNodeWithdrawNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UltraLightNodeWithdrawNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UltraLightNodeWithdrawNative represents a WithdrawNative event raised by the UltraLightNode contract.
type UltraLightNodeWithdrawNative struct {
	MsgSender common.Address
	To        common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawNative is a free log retrieval operation binding the contract event 0x3bfd26201736b5cb14a562ab3cfc2bef76901726e3a78483d6288af47131e1d9.
//
// Solidity: event WithdrawNative(address indexed msgSender, address indexed to, uint256 amount)
func (_UltraLightNode *UltraLightNodeFilterer) FilterWithdrawNative(opts *bind.FilterOpts, msgSender []common.Address, to []common.Address) (*UltraLightNodeWithdrawNativeIterator, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UltraLightNode.contract.FilterLogs(opts, "WithdrawNative", msgSenderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UltraLightNodeWithdrawNativeIterator{contract: _UltraLightNode.contract, event: "WithdrawNative", logs: logs, sub: sub}, nil
}

// WatchWithdrawNative is a free log subscription operation binding the contract event 0x3bfd26201736b5cb14a562ab3cfc2bef76901726e3a78483d6288af47131e1d9.
//
// Solidity: event WithdrawNative(address indexed msgSender, address indexed to, uint256 amount)
func (_UltraLightNode *UltraLightNodeFilterer) WatchWithdrawNative(opts *bind.WatchOpts, sink chan<- *UltraLightNodeWithdrawNative, msgSender []common.Address, to []common.Address) (event.Subscription, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UltraLightNode.contract.WatchLogs(opts, "WithdrawNative", msgSenderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UltraLightNodeWithdrawNative)
				if err := _UltraLightNode.contract.UnpackLog(event, "WithdrawNative", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawNative is a log parse operation binding the contract event 0x3bfd26201736b5cb14a562ab3cfc2bef76901726e3a78483d6288af47131e1d9.
//
// Solidity: event WithdrawNative(address indexed msgSender, address indexed to, uint256 amount)
func (_UltraLightNode *UltraLightNodeFilterer) ParseWithdrawNative(log types.Log) (*UltraLightNodeWithdrawNative, error) {
	event := new(UltraLightNodeWithdrawNative)
	if err := _UltraLightNode.contract.UnpackLog(event, "WithdrawNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UltraLightNodeWithdrawZROIterator is returned from FilterWithdrawZRO and is used to iterate over the raw logs and unpacked data for WithdrawZRO events raised by the UltraLightNode contract.
type UltraLightNodeWithdrawZROIterator struct {
	Event *UltraLightNodeWithdrawZRO // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UltraLightNodeWithdrawZROIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UltraLightNodeWithdrawZRO)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UltraLightNodeWithdrawZRO)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UltraLightNodeWithdrawZROIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UltraLightNodeWithdrawZROIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UltraLightNodeWithdrawZRO represents a WithdrawZRO event raised by the UltraLightNode contract.
type UltraLightNodeWithdrawZRO struct {
	MsgSender common.Address
	To        common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawZRO is a free log retrieval operation binding the contract event 0x3a20c8c3cd1848485ae8261a52398bb9b26f195b717306b3cf7f058e62c095d5.
//
// Solidity: event WithdrawZRO(address indexed msgSender, address indexed to, uint256 amount)
func (_UltraLightNode *UltraLightNodeFilterer) FilterWithdrawZRO(opts *bind.FilterOpts, msgSender []common.Address, to []common.Address) (*UltraLightNodeWithdrawZROIterator, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UltraLightNode.contract.FilterLogs(opts, "WithdrawZRO", msgSenderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UltraLightNodeWithdrawZROIterator{contract: _UltraLightNode.contract, event: "WithdrawZRO", logs: logs, sub: sub}, nil
}

// WatchWithdrawZRO is a free log subscription operation binding the contract event 0x3a20c8c3cd1848485ae8261a52398bb9b26f195b717306b3cf7f058e62c095d5.
//
// Solidity: event WithdrawZRO(address indexed msgSender, address indexed to, uint256 amount)
func (_UltraLightNode *UltraLightNodeFilterer) WatchWithdrawZRO(opts *bind.WatchOpts, sink chan<- *UltraLightNodeWithdrawZRO, msgSender []common.Address, to []common.Address) (event.Subscription, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UltraLightNode.contract.WatchLogs(opts, "WithdrawZRO", msgSenderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UltraLightNodeWithdrawZRO)
				if err := _UltraLightNode.contract.UnpackLog(event, "WithdrawZRO", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawZRO is a log parse operation binding the contract event 0x3a20c8c3cd1848485ae8261a52398bb9b26f195b717306b3cf7f058e62c095d5.
//
// Solidity: event WithdrawZRO(address indexed msgSender, address indexed to, uint256 amount)
func (_UltraLightNode *UltraLightNodeFilterer) ParseWithdrawZRO(log types.Log) (*UltraLightNodeWithdrawZRO, error) {
	event := new(UltraLightNodeWithdrawZRO)
	if err := _UltraLightNode.contract.UnpackLog(event, "WithdrawZRO", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
