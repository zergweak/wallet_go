package asset

var assetKey map[string]string
var assetAddress map[string]string
var assetName []string

func InitAsset() {
	assetKey = make(map[string]string)
	assetAddress = make(map[string]string)
	assetKey["ALLEExchange"] = "1a6695ac397dcc2b7a452c1a3eea546c9c00428062973b982bc82e5b5c7670f2"
	assetAddress["ALLEExchange"] = "TTbioAsbefqWxtoGk3MsKkUw7jgtFKPH9E"
	assetKey["TronWalletMe"] = "ec2575898d6a23b9e9ecb65507f8a043bfb4730c7d3c76d61a83867936fd74ad"
	assetAddress["TronWalletMe"] = "TCTb1zM4zQ6Ai5XmU9Lb19NEmL6HToxs86"
	assetKey["TronsTronics"] = "392fb4acd656d340a0244c36b3815f373e25552aaeb3b13592fa1eca44fe838d"
	assetAddress["TronsTronics"] = "TWvzaTM8UrpSeXufseXJqtkfc5GcRieprF"
	assetKey["RingCoin"] = "cdb66604cb9b2c78767911b5fdefd4d860a839179071b30c31a82029528172d1"
	assetAddress["RingCoin"] = "TRonANwbu14QM4Pnu3PdUPsEMpUw2bcssg"
	assetKey["BitGuild"] = "c44108cf141e16127a11b5f499a576f44eeeb7c6509297e0de7af3ce0e3e2552"
	assetAddress["BitGuild"] = "TV9Qi1uauZAt8KZWopHV1LbQd5RrLBrS3V"
	assetKey["FileCoin"] = "2d390ee5d3920dc0376dcec977ad7e91514927c10b1c1e421857fa204e756f91"
	assetAddress["FileCoin"] = "TKkJnDF8mRH24YZ7VKBfxhBUQrp1xn8c2S"

	assetName = append(assetName, "ALLEExchange")
	assetName = append(assetName, "TronWalletMe")
	assetName = append(assetName, "TronsTronics")
	assetName = append(assetName, "RingCoin")
	assetName = append(assetName, "BitGuild")
	assetName = append(assetName, "FileCoin")
}

func GetAsset(name string) (key string, address string) {
	return assetKey[name], assetAddress[name]
}

func GetAddressList() []string {
	var address []string
	for i := 0; i < len(assetName); i++ {
		address = append(address, assetAddress[assetName[i]])
	}
	address = append(address, "TFZuTwsDvpJQqbAZv9czKP4wbU975YkWjL")
	return address
}

func GetAssetList() [] string {
	return assetName
}
