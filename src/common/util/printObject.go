package util

import (
	"bytes"
	"common/crypto"
	"protocol/core"
	"strconv"
	"time"
	"wallet_go/src/common/hexutil"
)

func printAccountResource(accountResource core.Account_AccountResource) string {
	var buffer bytes.Buffer
	//result += "energy_usage: ";
	//result += accountResource.getEnergyUsage();
	//result += "\n";
	//result += "frozen_balance_for_energy: ";
	//result += "{";
	//result += "\n";
	//result += "  amount: ";
	//result += accountResource.getFrozenBalanceForEnergy().getFrozenBalance();
	//result += "\n";
	//result += "  expire_time: ";
	//result += new Date(accountResource.getFrozenBalanceForEnergy().getExpireTime());
	//result += "\n";
	//result += "}";
	//result += "\n";
	//result += "latest_consume_time_for_energy: ";
	//result += accountResource.getLatestConsumeTimeForEnergy();
	//result += "\n";
	//result += "storage_limit: ";
	//result += accountResource.getStorageLimit();
	//result += "\n";
	//result += "storage_usage: ";
	//result += accountResource.getStorageUsage();
	//result += "\n";
	//result += "latest_exchange_storage_time: ";
	//result += accountResource.getLatestExchangeStorageTime();
	//result += "\n";
	return buffer.String()
}

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

	buffer.WriteString("free_net_usage: ")
	buffer.WriteString(strconv.FormatInt(ac.GetFreeNetUsage(), 10))
	buffer.WriteByte('\n')
	buffer.WriteString("net_usage: ")
	buffer.WriteString(strconv.FormatInt(ac.GetNetUsage(), 10))
	buffer.WriteByte('\n')

	if ac.GetCreateTime() != 0 {
		buffer.WriteString("create_time: ")
		buffer.WriteString(time.Unix(ac.GetCreateTime()/1000, 0).Format("2006-01-02 03:04:05 PM"))
		buffer.WriteByte('\n')
	}
	if len(ac.GetVotes()) > 0 {
		buffer.WriteString("votes: \n{\n")
		for _, vote := range ac.GetVotes() {
			buffer.WriteString("  vote_address: ")
			buffer.WriteString(crypto.B58checkencode(vote.VoteAddress))
			buffer.WriteByte('\n')
			buffer.WriteString("  vote_count: ")
			buffer.WriteString(strconv.FormatInt(vote.GetVoteCount(), 10))
			buffer.WriteByte('\n')
		}
		buffer.WriteString("}\n")
	}
	if len(ac.GetAsset()) > 0 {
		buffer.WriteString("asset: \n{\n")
		for name, assetBalance := range ac.GetAsset() {
			buffer.WriteString("  name: ")
			buffer.WriteString(name)
			buffer.WriteByte('\n')
			buffer.WriteString("  balance: ")
			buffer.WriteString(strconv.FormatInt(assetBalance, 10))
			buffer.WriteByte('\n')
			buffer.WriteString("  latest_asset_operation_time: ")
			buffer.WriteString(time.Unix(ac.GetLatestAssetOperationTime()[name]/1000, 0).Format("2006-01-02 03:04:05 PM"))
			buffer.WriteByte('\n')
			buffer.WriteString("  free_asset_net_usage: ")
			buffer.WriteString(strconv.FormatInt(ac.GetFreeAssetNetUsage()[name], 10))
			buffer.WriteByte('\n')
		}
		buffer.WriteString("}\n")
	}
	if len(ac.GetFrozenSupply()) > 0 {
		buffer.WriteString("frozen_supply: \n{\n")
		for _, frozen := range ac.GetFrozenSupply() {
			buffer.WriteString("  amount: ")
			buffer.WriteString(strconv.FormatInt(frozen.GetFrozenBalance(), 10))
			buffer.WriteByte('\n')
			buffer.WriteString("  expire_time: ")
			buffer.WriteString(time.Unix(frozen.GetExpireTime()/1000, 0).Format("2006-01-02 03:04:05 PM"))
			buffer.WriteByte('\n')
		}
	}
	if ac.GetLatestOprationTime() > 0 {
		buffer.WriteString("latest_opration_time: ")
		buffer.WriteString(time.Unix(ac.GetLatestOprationTime()/1000, 0).Format("2006-01-02 03:04:05 PM"))
		buffer.WriteByte('\n')
	}
	if ac.GetLatestConsumeTime() > 0 {
		buffer.WriteString("latest_consume_time: ")
		buffer.WriteString(time.Unix(ac.GetLatestConsumeTime()/1000, 0).Format("2006-01-02 03:04:05 PM"))
		buffer.WriteByte('\n')
	}
	if ac.GetLatestConsumeFreeTime() > 0 {
		buffer.WriteString("latest_consume_free_time: ")
		buffer.WriteString(time.Unix(ac.GetLatestConsumeFreeTime()/1000, 0).Format("2006-01-02 03:04:05 PM"))
		buffer.WriteByte('\n')
	}
	buffer.WriteString("allowance: ")
	buffer.WriteString(strconv.FormatInt(ac.GetAllowance(), 10))
	buffer.WriteByte('\n')
	if ac.GetLatestWithdrawTime() > 0 {
		buffer.WriteString("latest_withdraw_time: ")
		buffer.WriteString(time.Unix(ac.GetLatestWithdrawTime()/1000, 0).Format("2006-01-02 03:04:05 PM"))
		buffer.WriteByte('\n')
	}
	buffer.WriteString("code: ")
	buffer.WriteString(hexutil.Encode(ac.GetCode()))
	buffer.WriteByte('\n')
	buffer.WriteString("is_witness: ")
	buffer.WriteString(strconv.FormatBool(ac.GetIsWitness()))
	buffer.WriteByte('\n')
	buffer.WriteString("is_committee: ")
	buffer.WriteString(strconv.FormatBool(ac.GetIsCommittee()))
	buffer.WriteByte('\n')
	buffer.WriteString("asset_issued_name: ")
	buffer.WriteString(string(ac.GetAssetIssuedName()))
	buffer.WriteByte('\n')
	buffer.WriteString("accountResource: \n{\n")
	buffer.WriteString(printAccountResource(*ac.GetAccountResource()))
	buffer.WriteString("}\n")
	return buffer.String()
}
