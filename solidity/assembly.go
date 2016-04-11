package solidity
import "encoding/json"
var Abi_Assembly interface{}
func Init_assembly(){
    json.Unmarshal([]byte(`[{"constant":false,"inputs":[{"name":"ver","type":"uint256"}],"name":"getProposalHistory","outputs":[{"name":"","type":"bytes"},{"name":"","type":"bytes"},{"name":"","type":"address"}],"type":"function"},{"constant":false,"inputs":[],"name":"getProposal","outputs":[{"name":"","type":"bytes"},{"name":"","type":"bytes"},{"name":"","type":"address"}],"type":"function"},{"constant":false,"inputs":[{"name":"hop","type":"bytes"},{"name":"hod","type":"bytes"}],"name":"revisionProposal","outputs":[],"type":"function"},{"inputs":[{"name":"hop","type":"bytes"},{"name":"hod","type":"bytes"}],"type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"name":"adrs","type":"address"},{"indexed":false,"name":"version","type":"uint256"}],"name":"onRevisionedProposal","type":"event"}]`), &Abi_Assembly)
}
var Bin_Assembly="0x6060604052604051610dd8380380610dd8833981016040528080518201919060200180518201919060200150505b6100378282610048565b5b5050610a0b806103cd6000396000f35b60006000508054806001018281815481835581811511610194576003028160030283600052602060002091820191016101939190610081565b8082111561018f57600060008201600050805460018160011615610100020316600290046000825580601f106100b757506100f4565b601f0160209004906000526020600020908101906100f391906100d5565b808211156100ef57600081815060009055506001016100d5565b5090565b5b5060018201600050805460018160011615610100020316600290046000825580601f10610121575061015e565b601f01602090049060005260206000209081019061015d919061013f565b80821115610159576000818150600090555060010161013f565b5090565b5b506002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101610081565b5090565b5b5050509190906000526020600020906003020160005b6060604051908101604052808681526020018581526020013381526020015090919091506000820151816000016000509080519060200190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061022357805160ff1916838001178555610254565b82800160010185558215610254579182015b82811115610253578251826000505591602001919060010190610235565b5b50905061027f9190610261565b8082111561027b5760008181506000905550600101610261565b5090565b50506020820151816001016000509080519060200190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106102d657805160ff1916838001178555610307565b82800160010185558215610307579182015b828111156103065782518260005055916020019190600101906102e8565b5b5090506103329190610314565b8082111561032e5760008181506000905550600101610314565b5090565b505060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908302179055505050507fcd2c8ceb3b73cf18e882ced169bb77c1ec46bdea63000b8ec9b39b4d433fac8130600160006000508054905003604051808373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a15b50505660606040526000357c010000000000000000000000000000000000000000000000000000000090048063529440ae1461004f578063b9e2bea01461014f578063eb5804bb146102465761004d565b005b610065600480803590602001909190505061083e565b6040518080602001806020018473ffffffffffffffffffffffffffffffffffffffff1681526020018381038352868181518152602001915080519060200190808383829060006004602084601f0104600f02600301f150905090810190601f1680156100e55780820380516001836020036101000a031916815260200191505b508381038252858181518152602001915080519060200190808383829060006004602084601f0104600f02600301f150905090810190601f16801561013e5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b61015c6004805050610668565b6040518080602001806020018473ffffffffffffffffffffffffffffffffffffffff1681526020018381038352868181518152602001915080519060200190808383829060006004602084601f0104600f02600301f150905090810190601f1680156101dc5780820380516001836020036101000a031916815260200191505b508381038252858181518152602001915080519060200190808383829060006004602084601f0104600f02600301f150905090810190601f1680156102355780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b6102e16004808035906020019082018035906020019191908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050909091908035906020019082018035906020019191908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509090919050506102e3565b005b6000600050805480600101828181548183558181151161042f5760030281600302836000526020600020918201910161042e919061031c565b8082111561042a57600060008201600050805460018160011615610100020316600290046000825580601f10610352575061038f565b601f01602090049060005260206000209081019061038e9190610370565b8082111561038a5760008181506000905550600101610370565b5090565b5b5060018201600050805460018160011615610100020316600290046000825580601f106103bc57506103f9565b601f0160209004906000526020600020908101906103f891906103da565b808211156103f457600081815060009055506001016103da565b5090565b5b506002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690555060010161031c565b5090565b5b5050509190906000526020600020906003020160005b6060604051908101604052808681526020018581526020013381526020015090919091506000820151816000016000509080519060200190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106104be57805160ff19168380011785556104ef565b828001600101855582156104ef579182015b828111156104ee5782518260005055916020019190600101906104d0565b5b50905061051a91906104fc565b8082111561051657600081815060009055506001016104fc565b5090565b50506020820151816001016000509080519060200190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061057157805160ff19168380011785556105a2565b828001600101855582156105a2579182015b828111156105a1578251826000505591602001919060010190610583565b5b5090506105cd91906105af565b808211156105c957600081815060009055506001016105af565b5090565b505060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908302179055505050507fcd2c8ceb3b73cf18e882ced169bb77c1ec46bdea63000b8ec9b39b4d433fac8130600160006000508054905003604051808373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a15b5050565b60206040519081016040528060008152602001506020604051908101604052806000815260200150600060006000600050600160006000508054905003815481101561000257906000526020600020906003020160005b50905080600001600050816001016000508260020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16828054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561078a5780601f1061075f5761010080835404028352916020019161078a565b820191906000526020600020905b81548152906001019060200180831161076d57829003601f168201915b50505050509250818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108265780601f106107fb57610100808354040283529160200191610826565b820191906000526020600020905b81548152906001019060200180831161080957829003601f168201915b50505050509150935093509350610838565b50909192565b6020604051908101604052806000815260200150602060405190810160405280600081526020015060006000600060005085815481101561000257906000526020600020906003020160005b50905080600001600050816001016000508260020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16828054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109555780601f1061092a57610100808354040283529160200191610955565b820191906000526020600020905b81548152906001019060200180831161093857829003601f168201915b50505050509250818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109f15780601f106109c6576101008083540402835291602001916109f1565b820191906000526020600020905b8154815290600101906020018083116109d457829003601f168201915b50505050509150935093509350610a03565b50919390925056"