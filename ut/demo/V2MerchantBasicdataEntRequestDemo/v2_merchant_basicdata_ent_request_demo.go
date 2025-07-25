/**
 * 企业商户进件 - 示例
 *
 * @Author sdk-generator
 * @Description 汇付天下
 */
package V2MerchantBasicdataEntRequestDemo

import (
    "encoding/json"
	"fmt"
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
)

func V2MerchantBasicdataEntRequestDemo() {
    // 1. 数据初始化
	dgSDK, _ := BsPaySdk.NewBsPay(true, "./config/config.json")

    // 2.组装请求参数
    dgReq := BsPaySdk.V2MerchantBasicdataEntRequest{
        // 请求流水号
        ReqSeqId:tool.GetReqSeqId(),
        // 请求日期
        ReqDate:tool.GetCurrentDate(),
        // 渠道商号
        UpperHuifuId:"6666000003080000",
        // 商户名称
        RegName:"集成企业商户8664",
        // 商户简称
        ShortName:"企业商户3471",
        // 小票名称
        ReceiptName:"盈盈超市",
        // 公司类型
        EntType:"1",
        // 所属行业参考[汇付MCC编码](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_hfmccbm) ；当use_head_info_flag&#x3D;Y时不填&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：5311&lt;/font&gt;
        Mcc:"5411",
        // 经营类型1：实体，2：虚拟 ；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1&lt;/font&gt; ；当use_head_info_flag&#x3D;Y时不填
        BusiType:"1",
        // 场景类型
        // SceneType:"test",
        // 证照图片
        // LicensePic:"test",
        // 证照编号
        LicenseCode:"20220422267883697",
        // 证照有效期类型
        LicenseValidityType:"1",
        // 证照有效期开始日期
        LicenseBeginDate:"20200401",
        // 证照有效期截止日期格式：yyyyMMdd。&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20220125&lt;/font&gt;&lt;br/&gt;  当license_validity_type&#x3D;0时必填；当license_validity_type&#x3D;1时为空；当use_head_info_flag&#x3D;Y时不填
        LicenseEndDate:"",
        // 成立时间
        // FoundDate:"test",
        // 注册资本保留两位小数；条件选填，国营企业、私营企业、外资企业、事业单位、其他、集体经济必填，政府机构、个体工商户可为空；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：100.00&lt;/font&gt;
        // RegCapital:"test",
        // 注册区
        RegDistrictId:"350203",
        // 注册详细地址
        RegDetail:"吉林省长春市思明区解放2路59096852",
        // 经营区
        DistrictId:"310104",
        // 经营详细地址scene_type&#x3D;OFFLINE/ALL时必填；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：上海市徐汇区XX路XX号&lt;/font&gt;
        DetailAddr:"吉林省长春市思明区解放1路49227677",
        // 法人姓名
        LegalName:"陈立一",
        // 法人证件类型
        LegalCertType:"00",
        // 法人证件号码
        LegalCertNo:"321084198912060000",
        // 法人证件有效期类型
        LegalCertValidityType:"1",
        // 法人证件有效期开始日期
        LegalCertBeginDate:"20121201",
        // 法人证件有效期截止日期日期格式：yyyyMMdd， &lt;font color&#x3D;&quot;green&quot;&gt;示例值：20220125&lt;/font&gt;&lt;br/&gt;当legal_cert_validity_type&#x3D;0时必填；&lt;br/&gt;当legal_cert_validity_type&#x3D;1时为空；&lt;br/&gt;当use_head_info_flag&#x3D;Y时不填
        LegalCertEndDate:"20301201",
        // 法人证件地址
        // LegalAddr:"test",
        // 法人身份证国徽面
        // LegalCertBackPic:"test",
        // 法人身份证人像面
        // LegalCertFrontPic:"test",
        // 联系人手机号
        ContactMobileNo:"13111112222",
        // 联系人电子邮箱
        ContactEmail:"jeff.peng@huifu.com",
        // 管理员账号
        LoginName:"LG20220422267883697",
        // 开户许可证企业商户需要，结算账号为对公账户必填；通过[图片上传接口](https://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc)上传材料；文件类型：F08；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // RegAcctPic:"test",
        // 基本存款账户编号或核准号条件选填；当use_head_info_flag&#x3D;Y时不填 ；&lt;br/&gt;基本存款账户信息请填写基本存款账户编号；开户许可证请填写核准号。&lt;br/&gt;当注册地址或经营地址为如下地区时必填：江苏省、浙江省、湖南省、湖北省、云南省、贵州省、陕西省、河南省、吉林省、黑龙江省、福建省、海南省、重庆市、青海省、宁夏回族自治区；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：J2900123456789&lt;/font&gt;
        OpenLicenceNo:"",
        // 银行卡信息配置
        CardInfo:getB2327a4a46a74af1B6da314607ca4387(),
        // 卡正面**对私必填**。通过[图片上传接口](https://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc)上传材料；文件类型：F13；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // SettleCardFrontPic:"test",
        // 持卡人身份证国徽面**对私必填**。通过[图片上传接口](https://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc)上传材料；文件类型：F56；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // SettleCertBackPic:"test",
        // 持卡人身份证人像面**对私必填**。通过[图片上传接口](https://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc)上传材料；文件类型：F55；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // SettleCertFrontPic:"test",
        // 授权委托书**对私非法人、对公非同名结算必填**；通过[图片上传接口](https://paas.huifu.com/open/doc/api/#/shgl/shjj/api_shjj_shtpsc)上传材料；文件类型：F15；开通银行电子账户（中信E管家）需提供F520；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：57cc7f00-600a-33ab-b614-6221bbf2e530&lt;/font&gt;
        // AuthEnturstPic:"test",
        // 上级汇付Id如果head_office_flag&#x3D;0，则字段必填，如果head_office_flag&#x3D;1，上级汇付Id不可传&lt;br/&gt;如果headOfficeFlag&#x3D;0，useHeadInfoFlag&#x3D;Y,且head_huifu_id不为空则基本信息部分复用上级的基本信息。&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123123123&lt;/font&gt;
        // HeadHuifuId:"test",
        // 商户ICP备案编号商户ICP备案编号或网站许可证号；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：沪ICP备06046402号-28 &lt;/font&gt;&lt;br/&gt;类型为PC网站时，且为企业商户，且开通快捷或网银，或大额转账，或余额支付或分账业务（20%（不含）-100%），或为个人商户开通分账业务（10%（不含）-100%），必填
        MerIcp:"",
        // 店铺门头照
        // StoreHeaderPic:"test",
        // 店铺内景/工作区域照
        // StoreIndoorPic:"test",
        // 店铺收银台/公司前台照
        // StoreCashierDeskPic:"test",
    }
    // 设置非必填字段
	dgReq.ExtendInfos = getExtendInfos()

    fmt.Println("请求时间:", tool.GetCurrentTime())
	respStr, _ := json.Marshal(dgReq)
    fmt.Println("请求数据:", string(respStr))

    // 3. 发起API调用
    resp, err := dgSDK.V2MerchantBasicdataEntRequest(dgReq)
  	if err != nil {
		fmt.Println("请求异常:", err)
	}

	fmt.Println("返回数据:", resp)
}

/**
 * 非必填字段
 */
func getExtendInfos() map[string]interface{} {
    // 设置非必填字段
    extendInfoMap := make(map[string]interface{})
    // 商户英文名称
    extendInfoMap["mer_en_name"] = ""
    // 证照类型
    extendInfoMap["license_type"] = ""
    // 经营范围
    // extendInfoMap["business_scope"] = ""
    // 注册省
    extendInfoMap["reg_prov_id"] = "350000"
    // 注册市
    extendInfoMap["reg_area_id"] = "350200"
    // 经营省
    extendInfoMap["prov_id"] = "310000"
    // 经营市
    extendInfoMap["area_id"] = "310100"
    // 法人手机号
    // extendInfoMap["legal_mobile_no"] = ""
    // 受益人列表
    // extendInfoMap["beneficiary_info"] = getF13a0b43F23d4aab91a0E326ec89a351()
    // 联系人姓名
    extendInfoMap["contact_name"] = "联系人"
    // 商户通知标识
    extendInfoMap["sms_send_flag"] = "Y"
    // 客服电话
    extendInfoMap["service_phone"] = "021-121111221"
    // 结算业务配置
    extendInfoMap["settle_config"] = getA908da42065845b8852550b11a12c8c0()
    // 取现业务配置
    extendInfoMap["cash_config"] = get33768d6aA9394ce5A369A88b918d2ea8()
    // 商户身份
    // extendInfoMap["head_office_flag"] = ""
    // 使用上级资料信息
    // extendInfoMap["use_head_info_flag"] = ""
    // 商户主页URL
    extendInfoMap["mer_url"] = ""
    // 外部商户号
    // extendInfoMap["ext_mer_id"] = ""
    // 备注
    // extendInfoMap["remarks"] = ""
    // 异步请求地址
    extendInfoMap["async_return_url"] = "virgo://http://192.168.85.157:30031/sspm/testVirgo"
    // 斗拱e账户功能配置
    // extendInfoMap["elec_acct_config"] = get0882475c315d4249A27dFd4dc2d278a6()
    // 股东信息
    // extendInfoMap["share_holder_info_list"] = getC12a6d98Fae14c13B3b1Eb20462fa0af()
    // 扩展资料包
    // extendInfoMap["extended_material_list"] = get008cc24b42ff4cd88c31663c087a4836()
    return extendInfoMap
}

func getF13a0b43F23d4aab91a0E326ec89a351() string {
    dto := make(map[string]interface{})
    // 受益人名称
    // dto["bo_name"] = "test"
    // 受益人证件类型
    // dto["bo_type"] = "test"
    // 受益人证件号
    // dto["bo_no"] = "test"
    // 受益人证件有效期开始时间
    // dto["bo_date_start"] = "test"
    // 受益人证件有效期结束时间
    // dto["bo_dead_line"] = "test"
    // 受益人证件地址
    // dto["bo_address"] = "test"
    // 受益人手机号开通全域资金管理业务时必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：13911111111&lt;/font&gt;
    // dto["bo_mobile_no"] = "test"
    // 最终受益人受益方式A01：直接或间接控股25%（含）以上 &lt;br/&gt;A02：通过人事、财务等其他方式对公司进行控制 &lt;br/&gt;A03：高级管理人员 &lt;br/&gt;A04：法人或公司负责人 &lt;br/&gt;A05：其他&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：A01&lt;/font&gt;&lt;br/&gt;开通全域资金管理业务时必填
    // dto["final_beneficiary_mode"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func getB2327a4a46a74af1B6da314607ca4387() string {
    dto := make(map[string]interface{})
    // 银行账户类型
    dto["card_type"] = "0"
    // 银行账户名
    dto["card_name"] = "上海以道数据服务中心"
    // 银行账号
    dto["card_no"] = "698043508"
    // 银行所在市
    // dto["area_id"] = "test"
    // 银行编码当card_type&#x3D;0时必填；参考： [银行编码](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_yhbm)；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：01020000&lt;/font&gt;；
    // dto["bank_code"] = "test"
    // 联行号当card_type&#x3D;0时必填，参考：[银行支行编码](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_yhzhbm) 当card_type&#x3D;0时必填， 当card_type&#x3D;1或2时非必填 &lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：102290026507&lt;/font&gt;
    dto["branch_code"] = "305290002096"
    // 持卡人证件类型持卡人证件类型，参见《[自然人证件类型](https://paas.huifu.com/open/doc/api/#/api_ggcsbm?id&#x3D;%e8%87%aa%e7%84%b6%e4%ba%ba%e8%af%81%e4%bb%b6%e7%b1%bb%e5%9e%8b)》。&lt;br/&gt; 当card_type&#x3D;0时为空， 当card_type&#x3D;1或2时必填； &lt;font color&#x3D;&quot;green&quot;&gt;示例值：00&lt;/font&gt;&lt;br/&gt;持卡人证件类型04、11需要补充 F101【结算人】港澳台居民来往内地通行证&lt;br/&gt;持卡人证件类型13需要补充 F513【结算人】外国人居留证&lt;br/&gt;持卡人证件类型14、15需要补充 F514【结算人】港澳台居住证&lt;br/&gt;其它持卡人证件类型补充F102【结算人】其它证件材料&lt;br/&gt;补充材料填写在extended_material_list 扩展资料包中
    // dto["cert_type"] = "test"
    // 持卡人证件号码对私必填；年龄不小于18岁且不能大于80岁；如持卡人证件类型为00：身份证，则填写身份证号码 ；&lt;br/&gt;&lt;font color&#x3D;&quot;green&quot;&gt;示例值：310112200001018888&lt;/font&gt;
    // dto["cert_no"] = "test"
    // 持卡人证件有效期截止日期日期格式：yyyyMMdd，&lt;font color&#x3D;&quot;green&quot;&gt;示例值：20220125&lt;/font&gt;&lt;br/&gt;  当cert_validity_type&#x3D;0时必填；当cert_validity_type&#x3D;1时为空
    // dto["cert_end_date"] = "test"
    // 银行所在省
    // dto["prov_id"] = ""
    // 持卡人证件有效期类型
    // dto["cert_validity_type"] = ""
    // 持卡人证件有效期开始日期
    // dto["cert_begin_date"] = ""
    // 银行卡绑定手机号
    // dto["mp"] = ""
    // 默认结算卡标志
    // dto["is_settle_default"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getA908da42065845b8852550b11a12c8c0() string {
    dto := make(map[string]interface{})
    // 结算周期
    dto["settle_cycle"] = "D1"
    // 结算手续费外扣商户号填写承担手续费的汇付商户号；当out_settle_flag&#x3D;1时必填，否则非必填；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123123123&lt;/font&gt;
    dto["out_settle_huifuid"] = ""
    // 结算批次号settle_cycle为TS时不填。当settle_pattern&#x3D;P0批次结算时必填，即将按指定结算批次号进行资金结算；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：100&lt;/font&gt;；[参见结算批次说明](https://paas.huifu.com/open/doc/api/#/csfl/api_csfl_jspc)
    dto["settle_batch_no"] = ""
    // 自定义结算处理时间settle_cycle为TS时不填。当settle_pattern&#x3D;P1/P2自定义时间结算时必填；注意：00:00到00:30不能指定&lt;br/&gt;到账时间，格式：HHmmss；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：103000&lt;/font&gt;
    dto["settle_time"] = ""
    // 节假日结算手续费率settle_cycle为D1、TS时必填。单位%，需保留小数点后两位。取值范围[0.00，100.00]，不收费请填写0.00；settle_cycle&#x3D;T1时，不生效 ；settle_cycle为D1时，遇节假日按此费率结算 ；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.05&lt;/font&gt;
    dto["fixed_ratio"] = "5.00"
    // 节假日结算手续费固定金额settle_cycle为D1、TS时必填。单位元，需保留小数点后两位。不收费请填写0.00；settle_cycle结算周期为D1时，遇节假日按此费率结算 ；&lt;br/&gt; &lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt;
    // dto["constant_amt"] = "test"
    // 起结金额
    dto["min_amt"] = "1.00"
    // 留存金额
    dto["remained_amt"] = "2.00"
    // 结算摘要
    dto["settle_abstract"] = "abstract"
    // 手续费外扣标记
    dto["out_settle_flag"] = "2"
    // 结算手续费外扣账户类型
    dto["out_settle_acct_type"] = ""
    // 结算方式
    dto["settle_pattern"] = ""
    // 是否优先到账
    dto["is_priority_receipt"] = ""
    // 工作日结算手续费率
    // dto["workday_fixed_ratio"] = ""
    // 工作日结算手续费固定金额
    // dto["workday_constant_amt"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func get33768d6aA9394ce5A369A88b918d2ea8() string {
    dto := make(map[string]interface{})
    // 取现手续费（固定/元）fix_amt与fee_rate至少填写一项， 需保留小数点后两位，不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt;注：当cash_type&#x3D;D1时为节假日取现手续费；当cash_type&#x3D;T1时为工作日取现手续费
    dto["fix_amt"] = "1.00"
    // 取现手续费率（%）fix_amt与fee_rate至少填写一项，需保留小数点后两位，取值范围[0.00,100.00]，不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.05&lt;/font&gt;注：1、如果fix_amt与fee_rate都填写了则手续费&#x3D;fix_amt+支付金额\*fee_rate2、当cash_type&#x3D;D1时为节假日取现手续费；当cash_type&#x3D;T1时为工作日取现手续费
    dto["fee_rate"] = ""
    // D1工作日取现手续费固定金额单位元，需保留小数点后两位。不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：1.00&lt;/font&gt;cash_type&#x3D;T1时，不生效 ；cash_type取现类型为D1时，遇工作日按此费率结算，若未配置则默认按照节假日手续费计算
    // dto["weekday_fix_amt"] = "test"
    // D1工作日取现手续费率单位%，需保留小数点后两位。取值范围[0.00，100.00]，不收费请填写0.00；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：0.05&lt;/font&gt;cash_type&#x3D;T1时，不生效 ；cash_type取现类型为D1时，遇工作日按此费率结算 ，若未配置则默认按照节假日手续费计算
    // dto["weekday_fee_rate"] = "test"
    // 手续费承担方手续费外扣时必需指定手续费承担方ID ；&lt;font color&#x3D;&quot;green&quot;&gt;示例值：6666000123123123&lt;/font&gt;
    // dto["out_fee_huifu_id"] = "test"
    // 取现类型
    dto["cash_type"] = "D0"
    // 是否取现手续费外扣
    // dto["out_fee_flag"] = ""
    // 手续费外扣的账户类型
    // dto["out_fee_acct_type"] = ""
    // 是否优先到账
    // dto["is_priority_receipt"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get622ded0e267647a8Ba6b9a7009c3ef51() interface{} {
    dto := make(map[string]interface{})
    // 银行编码
    // dto["bank_code"] = "test"
    // 支行联行号
    // dto["branch_code"] = "test"
    // 支行名称
    // dto["branch_name"] = "test"
    // 银行账户名
    // dto["card_name"] = "test"
    // 银行卡号
    // dto["card_no"] = "test"
    // 银行账户类型
    // dto["card_type"] = "test"
    // 银行所在省
    // dto["prov_id"] = ""
    // 银行所在市
    // dto["area_id"] = ""
    // 默认卡标识
    // dto["default_cash_flag"] = ""

    dtoList := [1]interface{}{dto}
    return dtoList
}

func get0882475c315d4249A27dFd4dc2d278a6() string {
    dto := make(map[string]interface{})
    // 电子账户开关
    // dto["switch_state"] = "test"
    // 账户类型
    // dto["acct_type"] = "test"
    // 电子账户取现手续费承担方
    // dto["cash_fee_party"] = "test"
    // 场景类型
    // dto["scene"] = "test"
    // 角色类型
    // dto["role_type"] = "test"
    // 签约成功标志
    // dto["sign_success_flag"] = "test"
    // 银行卡信息
    // dto["elec_card_list"] = get622ded0e267647a8Ba6b9a7009c3ef51()
    // 中信签约短信流水号
    // dto["elec_acct_sign_seq_id"] = ""

    dtoByte, _ := json.Marshal(dto)
    return string(dtoByte)
}

func getC12a6d98Fae14c13B3b1Eb20462fa0af() string {
    dto := make(map[string]interface{})
    // 股东姓名
    // dto["name"] = "test"
    // 股东证件类型
    // dto["cert_type"] = "test"
    // 股东证件号码
    // dto["cert_no"] = "test"
    // 股东证件有效期类型
    // dto["cert_validity_type"] = "test"
    // 股东证件有效期起始日
    // dto["cert_begin_date"] = "test"
    // 股东证件有效期到期日
    // dto["cert_end_date"] = ""

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

func get008cc24b42ff4cd88c31663c087a4836() string {
    dto := make(map[string]interface{})
    // 文件id
    // dto["file_id"] = "test"
    // 文件类型
    // dto["file_type"] = "test"

    dtoList := [1]interface{}{dto}
    dtoByte, _ := json.Marshal(dtoList)
    return string(dtoByte)
}

