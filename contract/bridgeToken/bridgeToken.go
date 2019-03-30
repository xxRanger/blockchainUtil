package bridgeToken

import (
	"blockchainUtil/contract"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"blockchainUtil/contract/GameToken"
)

const (
	ABI="[  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"spender\",     \"type\": \"address\"    },    {     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"approve\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"by\",     \"type\": \"address\"    },    {     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"consume\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"tokenId\",     \"type\": \"uint256\"    },    {     \"name\": \"user\",     \"type\": \"address\"    }   ],   \"name\": \"equipArmor\",   \"outputs\": [],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"tokenId\",     \"type\": \"uint256\"    },    {     \"name\": \"user\",     \"type\": \"address\"    }   ],   \"name\": \"equipWeapon\",   \"outputs\": [],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"user\",     \"type\": \"address\"    },    {     \"name\": \"amount\",     \"type\": \"uint256\"    }   ],   \"name\": \"exchange\",   \"outputs\": [],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"tokenID\",     \"type\": \"uint256\"    }   ],   \"name\": \"exchangeNFT\",   \"outputs\": [],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"to\",     \"type\": \"address\"    },    {     \"name\": \"tokenId\",     \"type\": \"uint256\"    }   ],   \"name\": \"mint\",   \"outputs\": [],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"user\",     \"type\": \"address\"    },    {     \"name\": \"amount\",     \"type\": \"uint256\"    }   ],   \"name\": \"pay\",   \"outputs\": [],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"tokenID\",     \"type\": \"uint256\"    },    {     \"name\": \"avatarOwner\",     \"type\": \"address\"    },    {     \"name\": \"gene\",     \"type\": \"uint256\"    },    {     \"name\": \"avatarLevel\",     \"type\": \"uint256\"    },    {     \"name\": \"weaponed\",     \"type\": \"bool\"    },    {     \"name\": \"armored\",     \"type\": \"bool\"    }   ],   \"name\": \"payNFT\",   \"outputs\": [],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"to\",     \"type\": \"address\"    },    {     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"reward\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"to\",     \"type\": \"address\"    },    {     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"transfer\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"from\",     \"type\": \"address\"    },    {     \"name\": \"to\",     \"type\": \"address\"    },    {     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"transferFrom\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"constant\": false,   \"inputs\": [    {     \"name\": \"tokenId\",     \"type\": \"uint256\"    }   ],   \"name\": \"upgrade\",   \"outputs\": [],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"function\"  },  {   \"inputs\": [    {     \"name\": \"totalSupply\",     \"type\": \"uint256\"    },    {     \"name\": \"tokenName\",     \"type\": \"string\"    },    {     \"name\": \"tokenSymbol\",     \"type\": \"string\"    },    {     \"name\": \"decimalUnits\",     \"type\": \"uint8\"    },    {     \"name\": \"_requiredSignatures\",     \"type\": \"uint256\"    }   ],   \"payable\": false,   \"stateMutability\": \"nonpayable\",   \"type\": \"constructor\"  },  {   \"anonymous\": false,   \"inputs\": [    {     \"indexed\": false,     \"name\": \"user\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"amount\",     \"type\": \"uint256\"    }   ],   \"name\": \"Exchange\",   \"type\": \"event\"  },  {   \"anonymous\": false,   \"inputs\": [    {     \"indexed\": false,     \"name\": \"user\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"amount\",     \"type\": \"uint256\"    }   ],   \"name\": \"Pay\",   \"type\": \"event\"  },  {   \"anonymous\": false,   \"inputs\": [    {     \"indexed\": false,     \"name\": \"tokenID\",     \"type\": \"uint256\"    },    {     \"indexed\": false,     \"name\": \"owner\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"gene\",     \"type\": \"uint256\"    },    {     \"indexed\": false,     \"name\": \"avatarLevel\",     \"type\": \"uint256\"    },    {     \"indexed\": false,     \"name\": \"weaponed\",     \"type\": \"bool\"    },    {     \"indexed\": false,     \"name\": \"armored\",     \"type\": \"bool\"    }   ],   \"name\": \"ExchangeNFT\",   \"type\": \"event\"  },  {   \"anonymous\": false,   \"inputs\": [    {     \"indexed\": false,     \"name\": \"tokenID\",     \"type\": \"uint256\"    },    {     \"indexed\": false,     \"name\": \"avatarOwner\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"gene\",     \"type\": \"uint256\"    },    {     \"indexed\": false,     \"name\": \"avatarLevel\",     \"type\": \"uint256\"    },    {     \"indexed\": false,     \"name\": \"weaponed\",     \"type\": \"bool\"    },    {     \"indexed\": false,     \"name\": \"armored\",     \"type\": \"bool\"    }   ],   \"name\": \"PayNFT\",   \"type\": \"event\"  },  {   \"anonymous\": false,   \"inputs\": [    {     \"indexed\": false,     \"name\": \"machine\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"player\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"Reward\",   \"type\": \"event\"  },  {   \"anonymous\": false,   \"inputs\": [    {     \"indexed\": false,     \"name\": \"machine\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"player\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"value\",     \"type\": \"uint256\"    }   ],   \"name\": \"Consume\",   \"type\": \"event\"  },  {   \"anonymous\": false,   \"inputs\": [    {     \"indexed\": true,     \"name\": \"_from\",     \"type\": \"address\"    },    {     \"indexed\": true,     \"name\": \"_to\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"_value\",     \"type\": \"uint256\"    }   ],   \"name\": \"Transfer\",   \"type\": \"event\"  },  {   \"anonymous\": false,   \"inputs\": [    {     \"indexed\": true,     \"name\": \"_owner\",     \"type\": \"address\"    },    {     \"indexed\": true,     \"name\": \"_spender\",     \"type\": \"address\"    },    {     \"indexed\": false,     \"name\": \"_value\",     \"type\": \"uint256\"    }   ],   \"name\": \"Approval\",   \"type\": \"event\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"\",     \"type\": \"address\"    }   ],   \"name\": \"_authorizdedMachines\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"owner\",     \"type\": \"address\"    },    {     \"name\": \"spender\",     \"type\": \"address\"    }   ],   \"name\": \"allowance\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"name\": \"avatar\",   \"outputs\": [    {     \"name\": \"gene\",     \"type\": \"uint256\"    },    {     \"name\": \"avatarLevel\",     \"type\": \"uint256\"    },    {     \"name\": \"weaponed\",     \"type\": \"bool\"    },    {     \"name\": \"armored\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"owner\",     \"type\": \"address\"    }   ],   \"name\": \"balanceOf\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [],   \"name\": \"decimals\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint8\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [],   \"name\": \"name\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"string\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"owner\",     \"type\": \"address\"    }   ],   \"name\": \"ownedAvatars\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"tokenId\",     \"type\": \"uint256\"    }   ],   \"name\": \"ownerOf\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"address\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [    {     \"name\": \"\",     \"type\": \"bytes32\"    }   ],   \"name\": \"payed\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"bool\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [],   \"name\": \"requiredSignatures\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [],   \"name\": \"symbol\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"string\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  },  {   \"constant\": true,   \"inputs\": [],   \"name\": \"totalSupply\",   \"outputs\": [    {     \"name\": \"\",     \"type\": \"uint256\"    }   ],   \"payable\": false,   \"stateMutability\": \"view\",   \"type\": \"function\"  } ]"
)

const (
	FuncPayToken string = "pay"
	FuncPayNFT string = "payNFT"
	FuncExchangeToken string = "exchange"
	FuncExchangeNFT string = "exchangeNFT"
)

const (
	EventPayToken string = "Pay"
	EventPayNFT string = "PayNFT"
	EventExchangeToken string = "Exchange"
	EventExchangeNFT string = "ExchangeNFT"
)

type BridgeTokenEventExchangeToken struct {
	User common.Address `json:"user"`
	Amount common.Address `json:"amount"`
}

//event ExchangeNFT(uint256 tokenID, address owner, uint256 gene, uint256 avatarLevel, bool weaponed, bool armored);
type BridgeTokenEventExchangeNFT struct {
	TokenID *big.Int `json:"tokenID"`
	Owner common.Address `json:"owner"`
	Gene *big.Int `json:"gene"`
	AvatarLevel *big.Int `json:"avatarLevel"`
	Weaponed bool `json:"weaponed"`
	Armored bool `json:"armored"`
}

type BridgeToken struct {
	*gameToken.GameToken
}

func NewBridgeToken(address common.Address) *BridgeToken {
	bridgeToken:=&BridgeToken{}
	bridgeToken.BaseContract = contract.NewBaseContract(address,ABI)
	return bridgeToken
}



