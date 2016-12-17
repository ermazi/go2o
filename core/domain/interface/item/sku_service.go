package item

type (
	ISkuService interface {
		// 将SKU字符串转为字典,如: 1:2;2:3
		SpecDataToMap(specData string) map[int]int
		// 获取规格和项的数组
		GetSpecItemArray(sku []*Sku) ([]int, []int)
	}

	// 商品SKU
	Sku struct {
		// 编号
		Id int32 `db:"id" pk:"yes" auto:"yes"`
		// 产品编号
		ProductId int32 `db:"product_id"`
		// 商品编号
		ItemId int32 `db:"item_id"`
		// 标题
		Title string `db:"title"`
		// 图片
		Image string `db:"image"`
		// 规格数据
		SpecData string `db:"spec_data"`
		// 规格字符
		SpecWord string `db:"spec_word"`
		// 产品编码
		Code string `db:"code"`
		// 参考价
		RetailPrice float32 `db:"-"`
		// 价格（分)
		Price float32 `db:"price"`
		// 成本（分)
		Cost float32 `db:"cost"`
		// 重量(克)
		Weight int32 `db:"weight"`
		// 体积（毫升)
		Bulk int32 `db:"bulk"`
		// 库存
		Stock int32 `db:"stock"`
		// 已销售数量
		SaleNum int32 `db:"sale_num"`
	}
)
