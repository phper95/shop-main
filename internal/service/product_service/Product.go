package product_service

import (
	"encoding/json"
	"errors"
	"gitee.com/phper95/pkg/httpclient"
	"gitee.com/phper95/pkg/sign"
	"gitee.com/phper95/pkg/strutil"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/unknwon/com"
	"net/http"
	"net/url"
	"shop/internal/models"
	"shop/internal/models/vo"
	"shop/internal/service/cate_service"
	"shop/internal/service/product_relation_service"
	"shop/internal/service/product_rule_service"
	productDto "shop/internal/service/product_service/dto"
	proVo "shop/internal/service/product_service/vo"
	"shop/pkg/constant"
	productEnum "shop/pkg/enums/product"
	"shop/pkg/global"
	"shop/pkg/logging"
	"shop/pkg/util"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Product struct {
	Id   int64
	Name string

	Enabled int

	PageNum  int
	PageSize int

	//M *models.StoreProductRule

	Ids []int64

	Dto productDto.StoreProduct

	SaleDto productDto.OnSale

	JsonObj map[string]interface{}

	Order int

	News       string
	PriceOrder string
	SalesOrder string
	Sid        string

	Uid int64

	Unique string

	Type string
}

//搜索结果响应结构
type searchResponse struct {
	Success bool                  `json:"success"`
	Code    int                   `json:"code"`
	Msg     string                `json:"msg"`
	Data    productSearchResponse `json:"data"`
}
type productSearchResponse struct {
	Total int64           `json:"total"`
	Hits  []*productIndex `json:"hits"`
}

type productIndex struct {
	Id int64 `json:"id"`
}

//get stock
func (d *Product) GetStock() int {
	var productAttrValue models.StoreProductAttrValue
	err := global.Db.Model(&models.StoreProductAttrValue{}).
		Where("`unique` = ?", d.Unique).
		Where("product_id = ?", d.Id).First(&productAttrValue).Error
	if err != nil {
		global.LOG.Error(err)
		return 0
	}
	return productAttrValue.Stock
}

func (d *Product) GetList() ([]proVo.Product, int, int) {
	maps := make(map[string]interface{})
	if d.Name != "" {
		maps["store_name"] = d.Name
	}
	if d.Enabled >= 0 {
		maps["is_show"] = d.Enabled
	}
	switch d.Order {
	case productEnum.STATUS_1:
		maps["is_best"] = 1
	case productEnum.STATUS_2:
		maps["is_new"] = 1
	case productEnum.STATUS_3:
		maps["is_benefit"] = 1
	case productEnum.STATUS_4:
		maps["is_hot"] = 1
	}

	if d.Sid != "" {
		maps["cate_id"] = com.StrTo(d.Sid).MustInt()
	}
	if d.News != "" {
		news := com.StrTo(d.News).MustInt()
		if news == 1 {
			maps["is_new"] = 1
		}
	}
	order := ""
	if d.SalesOrder != "" {
		if productEnum.ASC == d.SalesOrder {
			order = "sales asc"
		} else if productEnum.DESC == d.SalesOrder {
			order = "sales desc"
		}
	}
	if d.PriceOrder != "" {
		if productEnum.ASC == d.PriceOrder {
			order = "price asc"
		} else if productEnum.DESC == d.PriceOrder {
			order = "price desc"
		}
	}

	var productListVo []proVo.Product

	total, list := models.GetFrontAllProduct(d.PageNum, d.PageSize, maps, order)
	e := copier.Copy(&productListVo, list)
	if e != nil {
		global.LOG.Error(e)
	}
	totalNum := util.Int64ToInt(total)
	totalPage := util.GetTotalPage(totalNum, d.PageSize)
	return productListVo, totalNum, totalPage
}

// GetProductByIDs
func (d *Product) GetProductByIDs() []proVo.Product {
	var productListVo []proVo.Product
	if len(d.Ids) == 0 {
		return productListVo
	}
	list := models.GetProductByIDs(map[string]interface{}{"id": d.Ids})
	e := copier.Copy(&productListVo, list)
	if e != nil {
		global.LOG.Error(e)
	}
	return productListVo
}

// SearchGoods
func (d *Product) SearchGoods() ([]proVo.Product, int, int) {
	var productSearchList []proVo.Product
	//请求搜索接口
	params := url.Values{}
	params.Add("userid", strutil.Int64ToString(d.Uid))
	params.Add("keyword", d.Name)
	params.Add("page_num", strconv.Itoa(d.PageNum))
	params.Add("page_size", strconv.Itoa(d.PageSize))
	if len(d.News) > 0 {
		params.Add("news", d.News)
	}
	if len(d.PriceOrder) > 0 {
		params.Add("price_order", d.PriceOrder)
	}
	if len(d.SalesOrder) > 0 {
		params.Add("sales_order", d.SalesOrder)
	}

	global.LOG.Infof("SearchGoods params: %+v", d)

	apiCfg := global.CONFIG.Api
	productSearchHost := "http://localhost:9090"
	productSearchUri := "/api/v1/product-search"
	authorization, date, err := sign.New(apiCfg.SearchProductAK, apiCfg.SearchProductSK, constant.AuthorizationExpire).Generate(
		productSearchUri, http.MethodGet, params)
	if err != nil {
		global.LOG.Error(err, params)
		return nil, 0, 0
	}
	headerAuth := httpclient.WithHeader(constant.HeaderAuthField, authorization)
	headerAuthDate := httpclient.WithHeader(constant.HeaderAuthDateField, date)
	httpCode, body, err := httpclient.Get(productSearchHost+productSearchUri, params, httpclient.WithTTL(time.Second*5),
		headerAuth, headerAuthDate)
	if err != nil || httpCode != http.StatusOK {
		global.LOG.Error(" httpclient.Get error", err, httpCode, string(body))
		return nil, 0, 0
	}
	searchRes := &searchResponse{}
	err = json.Unmarshal(body, searchRes)
	if err != nil {
		global.LOG.Error("Unmarshal searchResponse error", err, string(body))
		return nil, 0, 0
	}
	if searchRes == nil {
		return productSearchList, 0, 0
	}
	if !searchRes.Success {
		global.LOG.Error("searchRes failed", string(body), searchRes)
		return nil, 0, 0
	}
	totalNum := int(searchRes.Data.Total)
	totalPage := util.GetTotalPage(totalNum, d.PageSize)
	productIDs := make([]int64, 0)
	for _, hit := range searchRes.Data.Hits {
		productIDs = append(productIDs, hit.Id)
	}
	global.LOG.Warnf("productIDs %v", productIDs)
	if len(productIDs) == 0 {
		return productSearchList, totalNum, totalPage
	}
	d.Ids = productIDs

	productSearchList = d.GetProductByIDs()
	global.LOG.Warnf("productSearchList", productSearchList)
	return productSearchList, totalNum, totalPage
}

func (d *Product) GetDetail() (*proVo.ProductDetail, error) {
	var (
		storeProduct models.StoreProduct
		productVo    proVo.Product
		err          error
	)
	err = global.Db.Model(&models.StoreProduct{}).
		Where("id = ?", d.Id).
		Where("is_show", 1).
		First(&storeProduct).Error
	if err != nil {
		global.LOG.Error(err)
		return nil, errors.New("获取商品失败")
	}
	//获取sku
	returnMap, err := getProductAttrDetail(d.Id)
	if err != nil {
		global.LOG.Error(err)
		return nil, errors.New("获取商品sku失败")
	}
	err = copier.Copy(&productVo, storeProduct)
	productVo.SliderImageArr = strings.Split(storeProduct.SliderImage, ",")
	if err != nil {
		global.LOG.Error(err)
		return nil, errors.New("商品转化失败")
	}
	//此处处理登录的用户
	//todo
	if d.Uid > 0 {
		isCollect := product_relation_service.IsRelation(d.Id, d.Uid)
		productVo.UserCollect = isCollect
	}

	//此处处理已经评论的数量-移动端需要，单个评价与好评旅，好评数量
	//todo
	//此处处理运费模板
	//todo
	detail := proVo.ProductDetail{
		StoreInfo:    productVo,
		ProductAttr:  returnMap["product_attr"].([]proVo.ProductAttr),
		ProductValue: returnMap["product_value"].(map[string]models.StoreProductAttrValue),
	}

	return &detail, nil
}

//获取商品sku
func getProductAttrDetail(productId int64) (map[string]interface{}, error) {
	var (
		storeProductAttrs    []models.StoreProductAttr
		productAttrValues    []models.StoreProductAttrValue
		mapp                 map[string]models.StoreProductAttrValue
		storeProductAttrList []proVo.ProductAttr
		err                  error
	)
	err = global.Db.Model(&models.StoreProductAttr{}).
		Where("product_id = ?", productId).
		Order("attr_values asc").Find(&storeProductAttrs).Error
	if err != nil {
		global.LOG.Error(err)
		return nil, err
	}
	err = global.Db.Model(&models.StoreProductAttrValue{}).
		Where("product_id = ?", productId).
		Find(&productAttrValues).Error
	if err != nil {
		global.LOG.Error(err)
		return nil, err
	}
	util.StructColumn(&mapp, productAttrValues, "", "Sku")
	//global.LOG.Info(mapp)

	for _, attr := range storeProductAttrs {
		stringList := strings.Split(attr.AttrValues, ",")
		var attrValues []productDto.AttrValue
		for _, str := range stringList {
			attrValue := productDto.AttrValue{
				Attr: str,
			}
			attrValues = append(attrValues, attrValue)
		}
		var attrVo proVo.ProductAttr
		err = copier.Copy(&attrVo, attr)
		if err != nil {
			global.LOG.Error(err)
			return nil, err
		}
		attrVo.AttrValue = attrValues
		attrVo.AttrValueArr = stringList
		storeProductAttrList = append(storeProductAttrList, attrVo)
	}

	returnMap := gin.H{
		"product_attr":  storeProductAttrList,
		"product_value": mapp,
	}
	return returnMap, nil
}

func (d *Product) OnSaleByProduct() error {
	return models.OnSaleByProduct(d.Id, d.SaleDto.Status)
}

func (d *Product) PublicFormatAttr() map[string]interface{} {
	return getFormatAttr(d.Id, d.JsonObj)
}

func (d *Product) AddOrSaveProduct() (model models.StoreProduct, err error) {
	var (
		productId int64
	)
	m := d.Dto
	copier.Copy(&model, &m)

	res := computeProduct(m.Attrs)
	model.Price = res["min_price"].(float64)
	model.OtPrice = res["min_ot_price"].(float64)
	model.Cost = res["min_cost"].(float64)
	model.Stock = res["stock"].(int)
	images := strings.Join(m.SliderImage, ",")
	model.SliderImage = images
	global.LOG.Warnf("model:%+v", model)
	if m.Id > 0 {
		err = models.UpdateByProduct(m.Id, &model)
		productId = m.Id
	} else {
		err = models.AddProduct(&model)
		if err != nil {
			global.LOG.Error("AddProduct error", err, model)
			return
		}
		productId = model.Id
	}

	//sku处理
	if m.SpecType == productEnum.SEPC_TYPE_0 {
		list1 := []string{"默认"}
		formatDetail := productDto.FormatDetail{
			Value:  "规格",
			Detail: list1,
		}
		productFormat := m.Attrs[0]
		productFormat.Value1 = "规格"
		productFormat.Detail = map[string]string{
			"规格": "默认",
		}
		err = insertProductSku([]productDto.FormatDetail{formatDetail}, []productDto.ProductFormat{productFormat}, productId)
	} else {
		err = insertProductSku(m.Items, m.Attrs, productId)
	}
	return
}

func (d *Product) GetProductInfo() map[string]interface{} {
	var (
		mapData      = make(map[string]interface{})
		StoreProduct models.StoreProduct
		productDto   productDto.StoreProductInfo
	)
	cateService := cate_service.Cate{}
	catList := cateService.GetProductCate()
	ruleService := product_rule_service.Rule{
		PageSize: 9999,
		PageNum:  1,
	}
	ruleList := ruleService.GetAll()
	mapData["cateList"] = catList
	mapData["ruleList"] = ruleList.Content
	if d.Id == 0 {
		return mapData
	}

	StoreProduct = models.GetProduct(d.Id)
	ee := copier.Copy(&productDto, StoreProduct)
	if ee != nil {
		logging.Error(ee)
	}
	productDto.SliderImage = strings.Split(StoreProduct.SliderImage, ",")
	res := models.GetProductAttrResult(d.Id)
	productDto.Attrs = res["value"]
	productDto.Items = res["attr"]

	mapData["productInfo"] = productDto

	return mapData
}

func (d *Product) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	if d.Name != "" {
		maps["name"] = d.Name
	}
	if d.Enabled >= 0 {
		maps["is_show"] = d.Enabled
	}

	total, list := models.GetAllProduct(d.PageNum, d.PageSize, maps)
	cateService := cate_service.Cate{}
	mapList := cateService.GetProductCate()
	return vo.ResultList{Content: list, TotalElements: total, ExtendData: mapList}
}

func (d *Product) Del() error {
	return models.DelByProduct(d.Ids)
}

func insertProductSku(items []productDto.FormatDetail, attrs []productDto.ProductFormat, productId int64) (err error) {
	err = models.DelByProductttr(productId)
	if err != nil {
		return err
	}
	err = models.DelByProductttrValue(productId)
	if err != nil {
		return err
	}
	err = models.AddProductAttr(items, productId)
	if err != nil {
		return err
	}

	err = models.AddProductttrValue(attrs, productId)
	if err != nil {
		return err
	}
	err = models.AddProductAttrResult(items, attrs, productId)
	if err != nil {
		return err
	}

	return

}

//计算获取属性结果最小值
func computeProduct(attrs []productDto.ProductFormat) map[string]interface{} {
	returnMap := make(map[string]interface{})

	var (
		minPrice   []float64
		minOtprice []float64
		minCost    []float64
		stock      []int
	)
	for _, dto := range attrs {
		price, _ := strconv.ParseFloat(dto.Price, 64)
		OtPrice, _ := strconv.ParseFloat(dto.Price, 64)
		cost, _ := strconv.ParseFloat(dto.Price, 64)
		s, _ := strconv.Atoi(dto.Stock)
		minPrice = append(minPrice, price)
		minOtprice = append(minOtprice, OtPrice)
		minCost = append(minCost, cost)
		stock = append(stock, s)
	}
	sort.Float64s(minPrice)
	sort.Float64s(minOtprice)
	sort.Float64s(minCost)
	returnMap["min_price"] = minPrice[0]
	returnMap["min_ot_price"] = minOtprice[0]
	returnMap["min_cost"] = minCost[0]
	returnMap["stock"] = util.GetSum(stock)
	return returnMap
}

//获取生成的商品sku
func getFormatAttr(id int64, jsonObj map[string]interface{}) map[string]interface{} {
	var (
		mapData          = make(map[string]interface{})
		formatDetailList []productDto.FormatDetail
		headerMapList    []map[string]interface{}
		valueMapList     []map[string]interface{}
		align            string = "center"
	)

	jsonByte, _ := json.Marshal(jsonObj["attrs"])
	json.Unmarshal(jsonByte, &formatDetailList)
	//logs.Info(formatDetailList)

	arr, ok := jsonObj["attrs"].([]interface{})
	if ok && len(arr) == 0 {
		return mapData
	}

	detail := attrFormat(formatDetailList)

	count := 0
	for _, mapData := range detail.Res {
		detailMap := mapData["detail"]
		valueMap := make(map[string]interface{})

		//组合表格头
		var i int = 0
		logging.Info(detailMap)
		if count == 0 {
			for kk, _ := range detailMap {
				headerMap := make(map[string]interface{})
				headerMap["title"] = kk
				headerMap["minWidth"] = 130
				headerMap["align"] = align
				myIntStr := strconv.Itoa(i + 1)
				headerMap["key"] = "value" + myIntStr
				headerMap["slot"] = "value" + myIntStr
				headerMapList = append(headerMapList, headerMap)
				i++
			}

			count++
		}

		//组合值
		j := 0
		skuArr := make([]string, 0, len(headerMapList))
		for _, kk := range headerMapList {
			key := "value" + strconv.Itoa(j+1)

			v := detailMap[kk["title"].(string)]
			valueMap[key] = detailMap[kk["title"].(string)]
			j++
			skuArr = append(skuArr, v)
		}
		sort.Strings(skuArr)
		sku := strings.Join(skuArr, ",")
		logging.Info("sku:" + sku)
		valueMap["detail"] = detailMap
		valueMap["pic"] = ""
		valueMap["price"] = "0"
		valueMap["cost"] = "0"
		valueMap["ot_price"] = "0"
		valueMap["stock"] = "0"
		valueMap["bar_code"] = ""
		valueMap["weight"] = "0"
		valueMap["volume"] = "0"
		//valueMap["brokerage"] = 0
		//valueMap["brokerage_two"] = 0
		//valueMap["pink_price"] = 0
		//valueMap["seckill_price"] = 0
		//valueMap["piink_stock"] = 0
		//valueMap["seckill_stock"] = 0
		if id > 0 {
			storeProductAttrValue := models.GetAttrValueByProductIdAndSku(id, sku)
			//
			valueMap["pic"] = storeProductAttrValue.Image
			valueMap["price"] = com.ToStr(storeProductAttrValue.Price)
			valueMap["cost"] = com.ToStr(storeProductAttrValue.Cost)
			valueMap["ot_price"] = com.ToStr(storeProductAttrValue.Price)
			valueMap["stock"] = com.ToStr(storeProductAttrValue.Stock)
			valueMap["bar_code"] = storeProductAttrValue.BarCode
			valueMap["weight"] = com.ToStr(storeProductAttrValue.Weight)
			valueMap["volume"] = com.ToStr(storeProductAttrValue.Volume)
			//valueMap["brokerage"] = storeProductAttrValue.Brokerage
			//valueMap["brokerage_two"] = storeProductAttrValue.BrokerageTwo
			//valueMap["pink_price"] = storeProductAttrValue.PinkPrice
			//valueMap["seckill_price"] = storeProductAttrValue.SeckillPrice
			//valueMap["piink_stock"] = storeProductAttrValue.PinkStock
			//valueMap["seckill_stock"] = storeProductAttrValue.SeckillStock
		}

		valueMapList = append(valueMapList, valueMap)
	}

	headerMapList = addMap(headerMapList, align)

	mapData["attr"] = formatDetailList
	mapData["value"] = valueMapList
	mapData["header"] = headerMapList

	return mapData
}

//组合map
func addMap(headerMapList []map[string]interface{}, align string) []map[string]interface{} {

	headMap := map[string]interface{}{
		"title":    "图片",
		"slot":     "pic",
		"align":    align,
		"minWidth": 80,
	}
	headerMapList = append(headerMapList, headMap)

	headMap = map[string]interface{}{
		"title":    "售价",
		"slot":     "price",
		"align":    align,
		"minWidth": 120,
	}
	headerMapList = append(headerMapList, headMap)

	headMap = map[string]interface{}{
		"title":    "成本价",
		"slot":     "cost",
		"align":    align,
		"minWidth": 140,
	}
	headerMapList = append(headerMapList, headMap)

	headMap = map[string]interface{}{
		"title":    "原价",
		"slot":     "ot_price",
		"align":    align,
		"minWidth": 140,
	}
	headerMapList = append(headerMapList, headMap)

	headMap = map[string]interface{}{
		"title":    "库存",
		"slot":     "stock",
		"align":    align,
		"minWidth": 140,
	}
	headerMapList = append(headerMapList, headMap)

	headMap = map[string]interface{}{
		"title":    "产品编号",
		"slot":     "bar_code",
		"align":    align,
		"minWidth": 140,
	}
	headerMapList = append(headerMapList, headMap)

	headMap = map[string]interface{}{
		"title":    "重量(kg)",
		"slot":     "weight",
		"align":    align,
		"minWidth": 140,
	}
	headerMapList = append(headerMapList, headMap)

	headMap = map[string]interface{}{
		"title":    "体积(m³)",
		"slot":     "volume",
		"align":    align,
		"minWidth": 140,
	}
	headerMapList = append(headerMapList, headMap)

	headMap = map[string]interface{}{
		"title":    "操作",
		"slot":     "action",
		"align":    align,
		"minWidth": 70,
	}
	headerMapList = append(headerMapList, headMap)

	return headerMapList

}

//组合sku规则算法
func attrFormat(formatDetailList []productDto.FormatDetail) productDto.Detail {
	var (
		data []string
		res  []map[string]map[string]string
	)

	if len(formatDetailList) > 1 { //当多个规则时候
		for i := 0; i < len(formatDetailList)-1; i++ {
			if i == 0 {
				data = formatDetailList[i].Detail
			}

			var tmp []string
			for _, v := range data {
				for _, g := range formatDetailList[i+1].Detail {
					rep2 := ""
					if i == 0 {
						rep2 = formatDetailList[i].Value + "_" + v + "-" + formatDetailList[i+1].Value + "_" + g
					} else {
						rep2 = v + "-" + formatDetailList[i+1].Value + "_" + g
					}

					tmp = append(tmp, rep2)

					if i == len(formatDetailList)-2 {
						var (
							rep4    = make(map[string]map[string]string)
							reptemp = make(map[string]string)
						)
						for _, h := range strings.Split(rep2, "-") {
							rep3 := strings.Split(h, "_")
							if len(rep3) > 1 {
								reptemp[rep3[0]] = rep3[1]
							} else {
								reptemp[rep3[0]] = ""
							}
						}

						rep4["detail"] = reptemp
						res = append(res, rep4)

					}

				}
			}

			if len(tmp) > 0 {
				data = tmp
			}

		}
	} else { //一个规则时候
		var dataArr []string
		for _, formatDetail := range formatDetailList {
			for _, str := range formatDetail.Detail {
				var map2 = make(map[string]map[string]string)
				dataArr = append(dataArr, formatDetail.Value+"_"+str)
				map1 := map[string]string{
					formatDetail.Value: str,
				}
				map2["detail"] = map1

				res = append(res, map2)
			}
		}

		s := strings.Join(dataArr, "-")
		data = append(data, s)
	}

	return productDto.Detail{
		Data: data,
		Res:  res,
	}
}
