/**
 * 电子账户付款查询
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package BsPaySdk

import "encoding/json"

type V2TradePaymentZxeAcctpyerqueryRequest struct {
    HuifuId string `json:"huifu_id" structs:"huifu_id"` // 商户号
    OrgReqDate string `json:"org_req_date" structs:"org_req_date"` // 原交易请求日期
    OrgHfSeqId string `json:"org_hf_seq_id" structs:"org_hf_seq_id"` // 原交易返回的全局流水号原交易请求流水号、原交易返回的全局流水号至少要送其中一项；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：00290TOP1GR210919004230P853ac13262200000&lt;/font&gt;
    OrgReqSeqId string `json:"org_req_seq_id" structs:"org_req_seq_id"` // 原交易请求流水号原交易请求流水号、原交易返回的全局流水号至少要送其中一项；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：2021091708126665001&lt;/font&gt;

	ExtendInfos map[string]interface{} `json:"extend_infos" structs:"extend_infos"` // 扩展字段
}

func (bp *BsPay) StrV2TradePaymentZxeAcctpyerqueryRequest(reqStr string) (map[string]interface{}, error) {
	reqParam := V2TradePaymentZxeAcctpyerqueryRequest{}
	json.Unmarshal([]byte(reqStr), &reqParam)
    json.Unmarshal([]byte(reqStr), &reqParam.ExtendInfos)
	return bp.V2TradePaymentZxeAcctpyerqueryRequest(reqParam)
}

func (bp *BsPay) V2TradePaymentZxeAcctpyerqueryRequest(reqParam V2TradePaymentZxeAcctpyerqueryRequest) (map[string]interface{}, error) {
	var url = BASE_API_TEST_URL_V2
	if bp.IsProdMode {
		url = BASE_API_URL_V2
	}
	reqUrl := url + V2_TRADE_PAYMENT_ZXE_ACCTPYERQUERY
    return PostRequest(reqUrl, ToMap(reqParam), bp.Msc)
}
