package libXray

import (
	"github.com/nknorg/nkn-sdk-go"
)

//
//const NknTextType = nkn.TextType
//const NknSessionType = nkn.SessionType
//const NknSessionIDSize = nkn.SessionIDSize
//const NknMultiClientIdentifierRe = nkn.MultiClientIdentifierRe
//const NknMessageIDSize = nkn.MessageIDSize
//const NknDefaultSessionAllowAddr = nkn.DefaultSessionAllowAddr
//const NknBinaryType = nkn.BinaryType
//const NknAckType = nkn.AckType
//const NknAmountUnit = nkn.AmountUnit
//
//type NknAccount = nkn.Account
//type NknAmount = nkn.Amount
//type NknClientAddr = nkn.ClientAddr
//type NknClient = nkn.Client
//type NknOnConnect = nkn.OnConnect
//type NknOnConnectFunc = nkn.OnConnectFunc
//type NknOnError = nkn.OnError
//type NknOnErrorFunc = nkn.OnErrorFunc
//type NknOnMessage = nkn.OnMessage
//type NknOnMessageFunc = nkn.OnMessageFunc
//type NknSubscribers = nkn.Subscribers
//type NknClientConfig = nkn.ClientConfig
//type NknDialConfig = nkn.DialConfig
//type NknWalletConfig = nkn.WalletConfig
//type NknWallet = nkn.Wallet
//type NknTransactionConfig = nkn.TransactionConfig
//type NknSubscription = nkn.Subscription
//type NknScryptConfig = nkn.ScryptConfig
//type NknRPCConfigInterface = nkn.RPCConfigInterface
//type NknRPCConfig = nkn.RPCConfig
//type NknResolver = nkn.Resolver
//type NknRegistrant = nkn.Registrant
//type NknNodeState = nkn.NodeState
//type NknNode = nkn.Node
//type NknNanoPayClaimer = nkn.NanoPayClaimer
//type NknNanoPay = nkn.NanoPay
//type NknMultiClient = nkn.MultiClient
//type NknMessageConfig = nkn.MessageConfig
//type NknMessage = nkn.Message
//type NknErrorWithCode = nkn.ErrorWithCode
//
//// vars
//var NknNewStringMap = nkn.NewStringMap
//var NknNewStringArray = nkn.NewStringArray
//var NknErrWrongRecipient = nkn.ErrWrongRecipient
//var NknErrWrongPassword = nkn.ErrWrongPassword
//var NknErrUnencryptedMessage = nkn.ErrUnencryptedMessage
//var NknErrResolveLimit = nkn.ErrResolveLimit
//var NknErrNotNanoPay = nkn.ErrNotNanoPay
//var NknErrNoDestination = nkn.ErrNoDestination
//var NknErrNilWebsocketConn = nkn.ErrNilWebsocketConn
//var NknErrNilClient = nkn.ErrNilClient
//var NknErrNanoPayClosed = nkn.ErrNanoPayClosed
//var NknErrMessageOversize = nkn.ErrMessageOversize
//var NknErrKeyNotInMap = nkn.ErrKeyNotInMap
//var NknErrInvalidWalletVersion = nkn.ErrInvalidWalletVersion
//var NknErrInvalidResolver = nkn.ErrInvalidResolver
//var NknErrInvalidPubkeySize = nkn.ErrInvalidPubkeySize
//var NknErrInvalidPubkeyOrName = nkn.ErrInvalidPubkeyOrName
//var NknErrInvalidPayloadType = nkn.ErrInvalidPayloadType
//var NknErrInvalidDestination = nkn.ErrInvalidDestination
//var NknErrInvalidAmount = nkn.ErrInvalidAmount
//var NknErrInsufficientBalance = nkn.ErrInsufficientBalance
//var NknErrExpiredNanoPayTxn = nkn.ErrExpiredNanoPayTxn
//var NknErrDecryptFailed = nkn.ErrDecryptFailed
//var NknErrCreateClientFailed = nkn.ErrCreateClientFailed
//var NknErrConnectFailed = nkn.ErrConnectFailed
//var NknErrClosed = nkn.ErrClosed
//var NknErrAddrNotAllowed = nkn.ErrAddrNotAllowed
//var NknDefaultWalletConfig = nkn.DefaultWalletConfig
//var NknDefaultTransactionConfig = nkn.DefaultTransactionConfig
//var NknDefaultStunServerAddr = nkn.DefaultStunServerAddr
//var NknDefaultSessionConfig = nkn.DefaultSessionConfig
//var NknDefaultSeedRPCServerAddr = nkn.DefaultSeedRPCServerAddr
//var NknDefaultRPCConfig = nkn.DefaultRPCConfig
//var NknDefaultMessageConfig = nkn.DefaultMessageConfig
//var NknDefaultDialConfig = nkn.DefaultDialConfig
//var NknDefaultClientConfig = nkn.DefaultClientConfig
//var NknNewStringMapWithSize = nkn.NewStringMapWithSize
//var NknNewStringArrayFromString = nkn.NewStringArrayFromString
//
//// Functions
//var NknGetBalance = nkn.GetBalance
//var NknGetDefaultWalletConfig = nkn.GetDefaultWalletConfig
//var NknWalletFromJSON = nkn.WalletFromJSON
//var NknUnsubscribeContext = nkn.UnsubscribeContext
//var NknUnsubscribe = nkn.Unsubscribe
//var NknTransferNameContext = nkn.TransferNameContext
//var NknTransferName = nkn.TransferName
//var NknTransferContext = nkn.TransferContext
//var NknTransfer = nkn.Transfer
//var NknSubscribeContext = nkn.SubscribeContext
//var NknSendRawTransactionContext = nkn.SendRawTransactionContext
//var NknSendRawTransaction = nkn.SendRawTransaction
//var NknRPCCall = nkn.RPCCall
//var NknResolveDests = nkn.ResolveDests
//var NknResolveDestN = nkn.ResolveDestN
//var NknResolveDest = nkn.ResolveDest
//var NknRegisterNameContext = nkn.RegisterNameContext
//var NknRegisterName = nkn.RegisterName
//var NknNewWallet = nkn.NewWallet
//var NknNewReplyPayload = nkn.NewReplyPayload
//var NknNewNanoPayClaimer = nkn.NewNanoPayClaimer
//var NknNewNanoPay = nkn.NewNanoPay
//var NknNewMultiClient = nkn.NewMultiClient
//var NknMergeWalletConfig = nkn.MergeWalletConfig
//var NknMergeTransactionConfig = nkn.MergeTransactionConfig
//var NknMergeMessageConfig = nkn.MergeMessageConfig
//var NknMergeDialConfig = nkn.MergeDialConfig
//var NknMergeClientConfig = nkn.MergeClientConfig
//var NknGetWssAddrContext = nkn.GetWssAddrContext
//var NknGetWssAddr = nkn.GetWssAddr
//var NknGetWsAddrContext = nkn.GetWsAddrContext
//var NknGetWsAddr = nkn.GetWsAddr
//var NknGetSubscriptionContext = nkn.GetSubscriptionContext
//var NknGetSubscription = nkn.GetSubscription
//var NknGetSubscribersCountContext = nkn.GetSubscribersCountContext
//var NknGetSubscribersCount = nkn.GetSubscribersCount
//var NknGetSubscribersContext = nkn.GetSubscribersContext
//var NknGetSubscribers = nkn.GetSubscribers
//var NknGetRegistrantContext = nkn.GetRegistrantContext
//var NknGetRegistrant = nkn.GetRegistrant
//var NknGetPeerAddrContext = nkn.GetPeerAddrContext
//var NknGetPeerAddr = nkn.GetPeerAddr
//var NknGetNonceContext = nkn.GetNonceContext
//var NknGetNonce = nkn.GetNonce
//var NknGetNodeStateContext = nkn.GetNodeStateContext
//var NknGetNodeState = nkn.GetNodeState
//var NknGetHeightContext = nkn.GetHeightContext
//var NknGetHeight = nkn.GetHeight
//var NknGetDefaultTransactionConfig = nkn.GetDefaultTransactionConfig
//var NknGetDefaultSessionConfig = nkn.GetDefaultSessionConfig
//var NknGetDefaultRPCConfig = nkn.GetDefaultRPCConfig
//var NknGetDefaultMessageConfig = nkn.GetDefaultMessageConfig
//var NknGetDefaultDialConfig = nkn.GetDefaultDialConfig
//var NknGetDefaultClientConfig = nkn.GetDefaultClientConfig
//var NknGetBalanceContext = nkn.GetBalanceContext
//var NknDeleteNameContext = nkn.DeleteNameContext
//var NknDeleteName = nkn.DeleteName
//var NknVerifyWalletAddress = nkn.VerifyWalletAddress
//var NknRandomBytes = nkn.RandomBytes
//var NknPubKeyToWalletAddr = nkn.PubKeyToWalletAddr
//var NknNewOnMessage = nkn.NewOnMessage
//var NknNewOnError = nkn.NewOnError
//var NknNewOnConnect = nkn.NewOnConnect
//var NknNewClient = nkn.NewClient
//var NknNewClientAddr = nkn.NewClientAddr
//var NknNewAmount = nkn.NewAmount
//var NknNewAccount = nkn.NewAccount
//var NknMeasureSeedRPCServerContext = nkn.MeasureSeedRPCServerContext
//var NknMeasureSeedRPCServer = nkn.MeasureSeedRPCServer
//var NknClientAddrToWalletAddr = nkn.ClientAddrToWalletAddr
//var NknClientAddrToPubKey = nkn.ClientAddrToPubKey
//var NknSubscribe = nkn.Subscribe

func NknNewAccount(seed []byte) (*nkn.Account, error) {
	return nkn.NewAccount(seed)
}

// Nkn test
func NknTest() string {
	return "NKN TEST"
}
