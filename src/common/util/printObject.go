package util

import (
	"bytes"
	"common/crypto"
	"protocol/core"
	"strconv"
	"time"
)

func PrintAccount(ac core.Account) string {
	var buffer bytes.Buffer
	buffer.WriteByte('\n')
	buffer.WriteString("address: ")
	buffer.WriteString(crypto.B58checkencode(ac.Address))
	buffer.WriteByte('\n')
	if len(ac.AccountId) != 0 {
		buffer.WriteString("account_id: ")
		buffer.Write(ac.AccountId)
		buffer.WriteByte('\n')
	}

	if len(ac.AccountName) != 0 {
		buffer.WriteString("account_name: ")
		buffer.Write(ac.AccountName)
		buffer.WriteByte('\n')
	}
	buffer.WriteString("type: ")
	buffer.WriteString(ac.GetType().String())
	buffer.WriteByte('\n')
	buffer.WriteString("balance: ")
	buffer.WriteString(strconv.FormatInt(ac.GetBalance(), 10))
	buffer.WriteByte('\n')

	if len(ac.GetFrozen()) > 0 {
		buffer.WriteString("frozen: \n{\n")
		for _, frozen := range ac.GetFrozen() {
			buffer.WriteString("  frozen_balance: ")
			buffer.WriteString(strconv.FormatInt(frozen.GetFrozenBalance(), 10))
			buffer.WriteByte('\n')
			buffer.WriteString("  expire_time: ")
			buffer.WriteString(time.Unix(frozen.GetExpireTime()/1000, 0).Format("2006-01-02 03:04:05 PM"))
			buffer.WriteByte('\n')
		}
		buffer.WriteString("}\n")
	}

	//result += "free_net_usage: ";
	//result += account.getFreeNetUsage();
	//result += "\n";
	//result += "net_usage: ";
	//result += account.getNetUsage();
	//result += "\n";
	//if (account.getCreateTime() != 0) {
	//	result += "create_time: ";
	//	result += new
	//	Date(account.getCreateTime());
	//	result += "\n";
	//}
	//if (account.getVotesCount() > 0) {
	//	for
	//	(Vote
	//vote:
	//	account.getVotesList()) {
	//		result += "votes";
	//		result += "\n";
	//		result += "{";
	//		result += "\n";
	//		result += "  vote_address: ";
	//		result += WalletApi.encode58Check(vote.getVoteAddress().toByteArray());
	//		result += "\n";
	//		result += "  vote_count: ";
	//		result += vote.getVoteCount();
	//		result += "\n";
	//		result += "}";
	//		result += "\n";
	//	}
	//}
	//if (account.getAssetCount() > 0) {
	//	for
	//	(String
	//name:
	//	account.getAssetMap().keySet()) {
	//		result += "asset";
	//		result += "\n";
	//		result += "{";
	//		result += "\n";
	//		result += "  name: ";
	//		result += name;
	//		result += "\n";
	//		result += "  balance: ";
	//		result += account.getAssetMap().get(name);
	//		result += "\n";
	//		result += "  latest_asset_operation_time: ";
	//		result += account.getLatestAssetOperationTimeMap().get(name);
	//		result += "\n";
	//		result += "  free_asset_net_usage: ";
	//		result += account.getFreeAssetNetUsageMap().get(name);
	//		result += "\n";
	//		result += "}";
	//		result += "\n";
	//	}
	//}
	//if (account.getFrozenSupplyCount() > 0) {
	//	for
	//	(Frozen
	//frozen:
	//	account.getFrozenSupplyList()) {
	//		result += "frozen_supply";
	//		result += "\n";
	//		result += "{";
	//		result += "\n";
	//		result += "  amount: ";
	//		result += frozen.getFrozenBalance();
	//		result += "\n";
	//		result += "  expire_time: ";
	//		result += new
	//		Date(frozen.getExpireTime());
	//		result += "\n";
	//		result += "}";
	//		result += "\n";
	//	}
	//}
	//result += "latest_opration_time: ";
	//result += new
	//Date(account.getLatestOprationTime());
	//result += "\n";
	//result += "latest_consume_time: ";
	//result += account.getLatestConsumeTime();
	//result += "\n";
	//result += "latest_consume_free_time: ";
	//result += account.getLatestConsumeFreeTime();
	//result += "\n";
	//result += "allowance: ";
	//result += account.getAllowance();
	//result += "\n";
	//result += "latest_withdraw_time: ";
	//result += new
	//Date(account.getLatestWithdrawTime());
	//result += "\n";
	//
	////    result += "code: ";
	////    result += account.getCode();
	////    result += "\n";
	////
	//result += "is_witness: ";
	//result += account.getIsWitness();
	//result += "\n";
	////
	////    result += "is_committee: ";
	////    result += account.getIsCommittee();
	////    result += "\n";
	//
	//result += "asset_issued_name: ";
	//result += account.getAssetIssuedName().toStringUtf8();
	//result += "\n";
	//result += "accountResource: {\n";
	//result += printAccountResource(account.getAccountResource());
	//result += "}\n";
	//return result;
	return buffer.String()
}
