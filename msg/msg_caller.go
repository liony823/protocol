package msg

import (
	"github.com/openimsdk/protocol/rpccall"
	"github.com/openimsdk/protocol/sdkws"
	"google.golang.org/grpc"
)

func InitMsg(conn *grpc.ClientConn) {
	GetMaxSeqCaller.SetConn(conn)
	GetMaxSeqsCaller.SetConn(conn)
	GetHasReadSeqsCaller.SetConn(conn)
	GetMsgByConversationIDsCaller.SetConn(conn)
	GetConversationMaxSeqCaller.SetConn(conn)
	PullMessageBySeqsCaller.SetConn(conn)
	GetSeqMessageCaller.SetConn(conn)
	SearchMessageCaller.SetConn(conn)
	SendMsgCaller.SetConn(conn)
	SetUserConversationsMinSeqCaller.SetConn(conn)
	ClearConversationsMsgCaller.SetConn(conn)
	UserClearAllMsgCaller.SetConn(conn)
	DeleteMsgsCaller.SetConn(conn)
	DeleteMsgPhysicalBySeqCaller.SetConn(conn)
	DeleteMsgPhysicalCaller.SetConn(conn)
	SetSendMsgStatusCaller.SetConn(conn)
	GetSendMsgStatusCaller.SetConn(conn)
	RevokeMsgCaller.SetConn(conn)
	MarkMsgsAsReadCaller.SetConn(conn)
	MarkConversationAsReadCaller.SetConn(conn)
	SetConversationHasReadSeqCaller.SetConn(conn)
	GetConversationsHasReadAndMaxSeqCaller.SetConn(conn)
	GetActiveUserCaller.SetConn(conn)
	GetActiveGroupCaller.SetConn(conn)
	GetServerTimeCaller.SetConn(conn)
	ClearMsgCaller.SetConn(conn)
	DestructMsgsCaller.SetConn(conn)
	GetActiveConversationCaller.SetConn(conn)
	AppendStreamMsgCaller.SetConn(conn)
	GetStreamMsgCaller.SetConn(conn)
	SetUserConversationMaxSeqCaller.SetConn(conn)
	SetUserConversationMinSeqCaller.SetConn(conn)
	GetLastMessageSeqByTime.SetConn(conn)
}

var (
	GetMaxSeqCaller                        = rpccall.NewRpcCaller[sdkws.GetMaxSeqReq, sdkws.GetMaxSeqResp](Msg_GetMaxSeq_FullMethodName)
	GetMaxSeqsCaller                       = rpccall.NewRpcCaller[GetMaxSeqsReq, SeqsInfoResp](Msg_GetMaxSeqs_FullMethodName)
	GetHasReadSeqsCaller                   = rpccall.NewRpcCaller[GetHasReadSeqsReq, SeqsInfoResp](Msg_GetHasReadSeqs_FullMethodName)
	GetMsgByConversationIDsCaller          = rpccall.NewRpcCaller[GetMsgByConversationIDsReq, GetMsgByConversationIDsResp](Msg_GetMsgByConversationIDs_FullMethodName)
	GetConversationMaxSeqCaller            = rpccall.NewRpcCaller[GetConversationMaxSeqReq, GetConversationMaxSeqResp](Msg_GetConversationMaxSeq_FullMethodName)
	PullMessageBySeqsCaller                = rpccall.NewRpcCaller[sdkws.PullMessageBySeqsReq, sdkws.PullMessageBySeqsResp](Msg_PullMessageBySeqs_FullMethodName)
	GetSeqMessageCaller                    = rpccall.NewRpcCaller[GetSeqMessageReq, GetSeqMessageResp](Msg_GetSeqMessage_FullMethodName)
	SearchMessageCaller                    = rpccall.NewRpcCaller[SearchMessageReq, SearchMessageResp](Msg_SearchMessage_FullMethodName)
	SendMsgCaller                          = rpccall.NewRpcCaller[SendMsgReq, SendMsgResp](Msg_SendMsg_FullMethodName)
	SetUserConversationsMinSeqCaller       = rpccall.NewRpcCaller[SetUserConversationsMinSeqReq, SetUserConversationsMinSeqResp](Msg_SetUserConversationsMinSeq_FullMethodName)
	ClearConversationsMsgCaller            = rpccall.NewRpcCaller[ClearConversationsMsgReq, ClearConversationsMsgResp](Msg_ClearConversationsMsg_FullMethodName)
	UserClearAllMsgCaller                  = rpccall.NewRpcCaller[UserClearAllMsgReq, UserClearAllMsgResp](Msg_UserClearAllMsg_FullMethodName)
	DeleteMsgsCaller                       = rpccall.NewRpcCaller[DeleteMsgsReq, DeleteMsgsResp](Msg_DeleteMsgs_FullMethodName)
	DeleteMsgPhysicalBySeqCaller           = rpccall.NewRpcCaller[DeleteMsgPhysicalBySeqReq, DeleteMsgPhysicalBySeqResp](Msg_DeleteMsgPhysicalBySeq_FullMethodName)
	DeleteMsgPhysicalCaller                = rpccall.NewRpcCaller[DeleteMsgPhysicalReq, DeleteMsgPhysicalResp](Msg_DeleteMsgPhysical_FullMethodName)
	SetSendMsgStatusCaller                 = rpccall.NewRpcCaller[SetSendMsgStatusReq, SetSendMsgStatusResp](Msg_SetSendMsgStatus_FullMethodName)
	GetSendMsgStatusCaller                 = rpccall.NewRpcCaller[GetSendMsgStatusReq, GetSendMsgStatusResp](Msg_GetSendMsgStatus_FullMethodName)
	RevokeMsgCaller                        = rpccall.NewRpcCaller[RevokeMsgReq, RevokeMsgResp](Msg_RevokeMsg_FullMethodName)
	MarkMsgsAsReadCaller                   = rpccall.NewRpcCaller[MarkMsgsAsReadReq, MarkMsgsAsReadResp](Msg_MarkMsgsAsRead_FullMethodName)
	MarkConversationAsReadCaller           = rpccall.NewRpcCaller[MarkConversationAsReadReq, MarkConversationAsReadResp](Msg_MarkConversationAsRead_FullMethodName)
	SetConversationHasReadSeqCaller        = rpccall.NewRpcCaller[SetConversationHasReadSeqReq, SetConversationHasReadSeqResp](Msg_SetConversationHasReadSeq_FullMethodName)
	GetConversationsHasReadAndMaxSeqCaller = rpccall.NewRpcCaller[GetConversationsHasReadAndMaxSeqReq, GetConversationsHasReadAndMaxSeqResp](Msg_GetConversationsHasReadAndMaxSeq_FullMethodName)
	GetActiveUserCaller                    = rpccall.NewRpcCaller[GetActiveUserReq, GetActiveUserResp](Msg_GetActiveUser_FullMethodName)
	GetActiveGroupCaller                   = rpccall.NewRpcCaller[GetActiveGroupReq, GetActiveGroupResp](Msg_GetActiveGroup_FullMethodName)
	GetServerTimeCaller                    = rpccall.NewRpcCaller[GetServerTimeReq, GetServerTimeResp](Msg_GetServerTime_FullMethodName)
	ClearMsgCaller                         = rpccall.NewRpcCaller[ClearMsgReq, ClearMsgResp](Msg_ClearMsg_FullMethodName)
	DestructMsgsCaller                     = rpccall.NewRpcCaller[DestructMsgsReq, DestructMsgsResp](Msg_DestructMsgs_FullMethodName)
	GetActiveConversationCaller            = rpccall.NewRpcCaller[GetActiveConversationReq, GetActiveConversationResp](Msg_GetActiveConversation_FullMethodName)
	AppendStreamMsgCaller                  = rpccall.NewRpcCaller[AppendStreamMsgReq, AppendStreamMsgResp](Msg_AppendStreamMsg_FullMethodName)
	GetStreamMsgCaller                     = rpccall.NewRpcCaller[GetStreamMsgReq, GetStreamMsgResp](Msg_GetStreamMsg_FullMethodName)
	SetUserConversationMaxSeqCaller        = rpccall.NewRpcCaller[SetUserConversationMaxSeqReq, SetUserConversationMaxSeqResp](Msg_SetUserConversationMaxSeq_FullMethodName)
	SetUserConversationMinSeqCaller        = rpccall.NewRpcCaller[SetUserConversationMinSeqReq, SetUserConversationMinSeqResp](Msg_SetUserConversationMinSeq_FullMethodName)
	GetLastMessageSeqByTime                = rpccall.NewRpcCaller[GetLastMessageSeqByTimeReq, GetLastMessageSeqByTimeResp](Msg_GetLastMessageSeqByTime_FullMethodName)
)
