package simaland

type Author struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// MARK: - Category
// Модели для работы со списком категорий

type ListResponse[T any] struct {
	Items []T   `json:"items,omitempty"`
	Links Links `json:"_links,omitempty"`
	Meta  Meta  `json:"_meta,omitempty"`
}

type Category struct {
	ID                      int64   `json:"id,omitempty"`
	Sid                     int64   `json:"sid,omitempty"`
	Name                    string  `json:"name,omitempty"`
	Anchor                  string  `json:"anchor,omitempty"`
	Priority                int64   `json:"priority,omitempty"`
	PriorityHome            int64   `json:"priority_home,omitempty"`
	PriorityMenu            int64   `json:"priority_menu,omitempty"`
	IsHiddenInMenu          int64   `json:"is_hidden_in_menu,omitempty"`
	Path                    string  `json:"path,omitempty"`
	Level                   int64   `json:"level,omitempty"`
	Type                    int64   `json:"type,omitempty"`
	IsAdult                 int64   `json:"is_adult,omitempty"`
	HasLocoSlider           bool    `json:"has_loco_slider,omitempty"`
	HasDesign               bool    `json:"has_design,omitempty"`
	HasAsMainDesign         bool    `json:"has_as_main_design,omitempty"`
	IsItemDescriptionHidden bool    `json:"is_item_description_hidden,omitempty"`
	IsForMobileApp          bool    `json:"is_for_mobile_app,omitempty"`
	CategoryGroupID         *int64  `json:"category_group_id,omitempty"`
	HasDesktopDesign        bool    `json:"has_desktop_design,omitempty"`
	HasMobileDesign         bool    `json:"has_mobile_design,omitempty"`
	H1                      string  `json:"h1,omitempty"`
	IsComparable            bool    `json:"is_comparable,omitempty"`
	Photo                   *string `json:"photo,omitempty"`
	Icon                    string  `json:"icon,omitempty"`
	IsLeaf                  int64   `json:"is_leaf,omitempty"`
	FullSlug                string  `json:"full_slug,omitempty"`
	AsideCategoriesIDS      []int64 `json:"aside_categories_ids,omitempty"`
}

type Links struct {
	Self Link `json:"self,omitempty"`
	Next Link `json:"next,omitempty"`
	Last Link `json:"last,omitempty"`
}

type Link struct {
	Href string `json:"href,omitempty"`
}

type Meta struct {
	TotalCount  int64 `json:"totalCount,omitempty"`
	PageCount   int64 `json:"pageCount,omitempty"`
	CurrentPage int64 `json:"currentPage,omitempty"`
	PerPage     int64 `json:"perPage,omitempty"`
}

// MARK: - Item
// Товар на платформе

type Item struct {
	ID  int64 `json:"id,omitempty"`
	Sid int64 `json:"sid,omitempty"`

	Name       string `json:"name,omitempty"`
	CategoryID int64  `json:"category_id,omitempty"`

	Price           int64 `json:"price,omitempty"`
	PriceMax        int64 `json:"price_max,omitempty"`
	RetailPrice     int64 `json:"retail_price,omitempty"`
	DiscountPercent int64 `json:"discountPercent,omitempty"`

	Photos []struct {
		URLPart string `json:"url_part,omitempty"`
		Version int    `json:"version,omitempty"`
	} `json:"photos,omitempty"`

	PhotoVersions []struct {
		Number  string `json:"number,omitempty"`
		Version string `json:"version,omitempty"`
	} `json:"photoVersions,omitempty"`

	// ActionUrls struct {
	// 	Offer string `json:"offer,omitempty"`
	// } `json:"action_urls,omitempty"`

	// QtyRulesData struct {
	// 	On string `json:"on,omitempty"`
	// } `json:"qty_rules_data,omitempty"`

	// DateInfo struct {
	// 	MinDate string `json:"min_date,omitempty"`
	// 	MaxDate string `json:"max_date,omitempty"`
	// 	IsPaid  bool   `json:"is_paid,omitempty"`
	// } `json:"date_info,omitempty"`

	// Country struct {
	// 	ID       int64  `json:"id,omitempty"`
	// 	Name     string `json:"name,omitempty"`
	// 	FullName string `json:"full_name,omitempty"`
	// 	Alpha2   string `json:"alpha2,omitempty"`
	// } `json:"country,omitempty"`

	// IsDisabled                float64  `json:"is_disabled,omitempty"`
	// MinimumOrderQuantity      string   `json:"minimum_order_quantity,omitempty"`
	// PricePerSquareMeter       float64  `json:"price_per_square_meter,omitempty"`
	// PricePerLinearMeter       float64  `json:"price_per_linear_meter,omitempty"`
	// Currency                  string   `json:"currency,omitempty"`
	// BoxtypeID                 float64  `json:"boxtype_id,omitempty"`
	// BoxDepth                  float64  `json:"box_depth,omitempty"`
	// BoxHeight                 float64  `json:"box_height,omitempty"`
	// BoxWidth                  float64  `json:"box_width,omitempty"`
	// InBox                     float64  `json:"in_box,omitempty"`
	// InSet                     float64  `json:"in_set,omitempty"`
	// Depth                     float64  `json:"depth,omitempty"`
	// UnitID                    float64  `json:"unit_id,omitempty"`
	// Width                     float64  `json:"width,omitempty"`
	// Height                    float64  `json:"height,omitempty"`
	// CountryID                 float64  `json:"country_id,omitempty"`
	// CartMinDiff               string   `json:"cart_min_diff,omitempty"`
	// KeepPackage               float64  `json:"keep_package,omitempty"`
	// PerPackage                float64  `json:"per_package,omitempty"`
	// VideoFileURL              bool     `json:"video_file_url,omitempty"`
	// IsHit                     float64  `json:"is_hit,omitempty"`
	// IsPriceFixed              float64  `json:"is_price_fixed,omitempty"`
	// IsExclusive               float64  `json:"is_exclusive,omitempty"`
	// IsMotley                  float64  `json:"is_motley,omitempty"`
	// IsAdult                   float64  `json:"is_adult,omitempty"`
	// IsProtected               float64  `json:"is_protected,omitempty"`
	// CertificateTypeID         float64  `json:"certificate_type_id,omitempty"`
	// ParentItemID              float64  `json:"parent_item_id,omitempty"`
	// MaxQty                    float64  `json:"max_qty,omitempty"`
	// ItemMinQty                float64  `json:"min_qty,omitempty"`
	// QtyMultiplier             float64  `json:"qty_multiplier,omitempty"`
	// IsLoco                    float64  `json:"is_loco,omitempty"`
	// IsPaidDelivery            float64  `json:"is_paid_delivery,omitempty"`
	// PackageVolume             string   `json:"package_volume,omitempty"`
	// HasDiscount               float64  `json:"has_discount,omitempty"`
	// IsGift                    float64  `json:"is_gift,omitempty"`
	// IsBoxed                   float64  `json:"is_boxed,omitempty"`
	// ProductVolume             float64  `json:"product_volume,omitempty"`
	// BoxVolume                 float64  `json:"box_volume,omitempty"`
	// BoxCapacity               float64  `json:"box_capacity,omitempty"`
	// PackingVolumeFactor       float64  `json:"packing_volume_factor,omitempty"`
	// IsTireSpike               float64  `json:"is_tire_spike,omitempty"`
	// IsTireRunFlat             float64  `json:"is_tire_run_flat,omitempty"`
	// TireSeasonID              float64  `json:"tire_season_id,omitempty"`
	// TireDiameterID            float64  `json:"tire_diameter_id,omitempty"`
	// TireWidthID               float64  `json:"tire_width_id,omitempty"`
	// TireSectionHeightID       float64  `json:"tire_section_height_id,omitempty"`
	// TireLoadIndexID           float64  `json:"tire_load_index_id,omitempty"`
	// TireSpeedIndexID          float64  `json:"tire_speed_index_id,omitempty"`
	// WheelLzID                 float64  `json:"wheel_lz_id,omitempty"`
	// WheelWidthID              float64  `json:"wheel_width_id,omitempty"`
	// WheelDiameterID           float64  `json:"wheel_diameter_id,omitempty"`
	// WheelDiaID                float64  `json:"wheel_dia_id,omitempty"`
	// WheelPcdID                float64  `json:"wheel_pcd_id,omitempty"`
	// WheelEtID                 float64  `json:"wheel_et_id,omitempty"`
	// Isbn                      string   `json:"isbn,omitempty"`
	// ItemIsAddToCartMultiple   float64  `json:"is_add_to_cart_multiple,omitempty"`
	// SupplyPeriod              float64  `json:"supply_period,omitempty"`
	// HasAction                 float64  `json:"has_action,omitempty"`
	// HasActionDiscountSystem   float64  `json:"has_action_discount_system,omitempty"`
	// HasJewelryAction          float64  `json:"has_jewelry_action,omitempty"`
	// Has3_Pay2_Action          float64  `json:"has_3_pay_2_action,omitempty"`
	// HasBestFabric             float64  `json:"has_best_fabric,omitempty"`
	// HasBestTextile            float64  `json:"has_best_textile,omitempty"`
	// HasNumberOneMadeInRussia  float64  `json:"has_number_one_made_in_russia,omitempty"`
	// PhotoIndexes              []string `json:"photoIndexes,omitempty"`
	// PhotoURL                  string   `json:"photoUrl,omitempty"`
	// IsMarkdown                float64  `json:"is_markdown,omitempty"`
	// IsPrepayNeeded            float64  `json:"is_prepay_needed,omitempty"`
	// IsPaidDeliveryEkb         bool     `json:"is_paid_delivery_ekb,omitempty"`
	// MeanRating                float64  `json:"mean_rating,omitempty"`
	// CommentsCount             float64  `json:"comments_count,omitempty"`
	// MarkdownReason            string   `json:"markdown_reason,omitempty"`
	// IsWholesale               float64  `json:"is_wholesale,omitempty"`
	// IsWholesaleConservation   float64  `json:"is_wholesale_conservation,omitempty"`
	// Type                      float64  `json:"type,omitempty"`
	// IsShockPrice              bool     `json:"is_shock_price,omitempty"`
	// IsRecommended             bool     `json:"is_recommended,omitempty"`
	// Vat                       float64  `json:"vat,omitempty"`
	// IsExportToS3              bool     `json:"is_export_to_s3,omitempty"`
	// CurrencySign              string   `json:"currencySign,omitempty"`
	// IsEnough                  bool     `json:"isEnough,omitempty"`
	// IsAddToCartMultiple       bool     `json:"isAddToCartMultiple,omitempty"`
	// MinQty                    float64  `json:"minQty,omitempty"`
	// QtyRule                   string   `json:"qtyRule,omitempty"`
	// QtyRules                  string   `json:"qty_rules,omitempty"`
	// PluralNameFormat          string   `json:"pluralNameFormat,omitempty"`
	// InBoxPluralNameFormat     string   `json:"inBoxPluralNameFormat,omitempty"`
	// CanBuyByCredit            bool     `json:"can_buy_by_credit,omitempty"`
	// SupplierCode              string   `json:"supplier_code,omitempty"`
	// Weight                    float64  `json:"weight,omitempty"`
	// HasSpecialOffer           bool     `json:"has_special_offer,omitempty"`
	// HasDayDiscount            float64  `json:"has_day_discount,omitempty"`
	// HasErichKrause            float64  `json:"has_erich_krause,omitempty"`
	// HasTmGammaGifts           float64  `json:"has_tm_gamma_gifts,omitempty"`
	// HasSuperpriceOnLine       float64  `json:"has_superprice_on_line,omitempty"`
	// HasWeekDiscount           float64  `json:"has_week_discount,omitempty"`
	// Has3DaysDiscount          float64  `json:"has_3days_discount,omitempty"`
	// HasBestFabric2018         float64  `json:"has_best_fabric_2018,omitempty"`
	// HasPayLater               float64  `json:"has_pay_later,omitempty"`
	// HasNewRules               float64  `json:"has_new_rules,omitempty"`
	// HasItemMonth              float64  `json:"has_item_month,omitempty"`
	// HasBatteriesGift          float64  `json:"has_batteries_gift,omitempty"`
	// Has4_Pay2_Action          float64  `json:"has_4_pay_2_action,omitempty"`
	// HasTakeInstallmentsAction float64  `json:"has_take_installments_action,omitempty"`
	// WholesalePrice            float64  `json:"wholesale_price,omitempty"`
	// WholesalePriceText        string   `json:"wholesale_price_text,omitempty"`
	// IsPart                    bool     `json:"is_part,omitempty"`
	// IsRemoteStore             float64  `json:"is_remote_store,omitempty"`
	// IsSmallWholesaleAvailable bool     `json:"is_small_wholesale_available,omitempty"`
	// IsPlant                   bool     `json:"is_plant,omitempty"`
	// Color                     string   `json:"color,omitempty"`
	// ImageTitle                string   `json:"image_title,omitempty"`
	// ImageAlt                  string   `json:"image_alt,omitempty"`
	// ShortName                 string   `json:"short_name,omitempty"`
	// IsFreeDelivery            bool     `json:"is_free_delivery,omitempty"`
	// MinSumForFreeDelivery     float64  `json:"min_sum_for_free_delivery,omitempty"`
	// Img                       string   `json:"img,omitempty"`
	// IsEntranceTypeByWeight    bool     `json:"isEntranceTypeByWeight,omitempty"`
	// RealMinQty                float64  `json:"real_min_qty,omitempty"`
	// IsWeightedGoods           bool     `json:"is_weighted_goods,omitempty"`
	// HasGift                   bool     `json:"hasGift,omitempty"`
	// HasGiftAssignee           bool     `json:"hasGiftAssignee,omitempty"`
	// IsNovelty                 bool     `json:"isNovelty,omitempty"`
	// HasVolumeDiscount         bool     `json:"has_volume_discount,omitempty"`
	// Size                      string   `json:"size,omitempty"`
	// Stuff                     string   `json:"stuff,omitempty"`
	// EcommerceVariant          string   `json:"ecommerce_variant,omitempty"`
	// LoanCategoryID            float64  `json:"loan_category_id,omitempty"`
	// IsItemDescriptionHidden   bool     `json:"is_item_description_hidden,omitempty"`
	// IsFoundCheaperEnabled     bool     `json:"is_found_cheaper_enabled,omitempty"`

	// ReasonOfDisabling       interface{} `json:"reason_of_disabling,omitempty"`
	// NestedUnitID            interface{} `json:"nested_unit_id,omitempty"`
	// TrademarkID             interface{} `json:"trademark_id,omitempty"`
	// VideoFileName           interface{} `json:"video_file_name,omitempty"`
	// SeriesID                interface{} `json:"series_id,omitempty"`
	// IsLicensed              interface{} `json:"is_licensed,omitempty"`
	// OfferID                 interface{} `json:"offer_id,omitempty"`
	// HasUSB                  interface{} `json:"has_usb,omitempty"`
	// HasBattery              interface{} `json:"has_battery,omitempty"`
	// HasClockwork            interface{} `json:"has_clockwork,omitempty"`
	// HasSound                interface{} `json:"has_sound,omitempty"`
	// HasRadiocontrol         interface{} `json:"has_radiocontrol,omitempty"`
	// IsInertial              interface{} `json:"is_inertial,omitempty"`
	// IsOnACPower             interface{} `json:"is_on_ac_power,omitempty"`
	// HasRusVoice             interface{} `json:"has_rus_voice,omitempty"`
	// HasRusPack              interface{} `json:"has_rus_pack,omitempty"`
	// HasLight                interface{} `json:"has_light,omitempty"`
	// IsDayOffer              interface{} `json:"is_day_offer,omitempty"`
	// PageTitle               interface{} `json:"page_title,omitempty"`
	// PageKeywords            interface{} `json:"page_keywords,omitempty"`
	// PageDescription         interface{} `json:"page_description,omitempty"`
	// ModifierID              interface{} `json:"modifier_id,omitempty"`
	// ModifierValue           interface{} `json:"modifier_value,omitempty"`
	// GiftID                  interface{} `json:"gift_id,omitempty"`
	// SurfaceArea             interface{} `json:"surface_area,omitempty"`
	// LinearMeters            interface{} `json:"linear_meters,omitempty"`
	// NoveltedAt              interface{} `json:"novelted_at,omitempty"`
	// MinAge                  interface{} `json:"min_age,omitempty"`
	// Power                   interface{} `json:"power,omitempty"`
	// Volume                  interface{} `json:"volume,omitempty"`
	// TransportConditionID    interface{} `json:"transport_condition_id,omitempty"`
	// HasBodyDrawing          interface{} `json:"has_body_drawing,omitempty"`
	// HasCordCase             interface{} `json:"has_cord_case,omitempty"`
	// HasTeapot               interface{} `json:"has_teapot,omitempty"`
	// HasTermostat            interface{} `json:"has_termostat,omitempty"`
	// IsImprintable           interface{} `json:"is_imprintable,omitempty"`
	// PageCount               interface{} `json:"page_count,omitempty"`
	// AudioFilename           interface{} `json:"audio_filename,omitempty"`
	// Photo3DCount            interface{} `json:"photo_3d_count,omitempty"`
	// CustomQtyRulesData      interface{} `json:"custom_qty_rules_data,omitempty"`
	// BalancePluralNameFormat interface{} `json:"balancePluralNameFormat,omitempty"`
	// SpecialOfferID          interface{} `json:"special_offer_id,omitempty"`
	// MinSumOrder             interface{} `json:"min_sum_order,omitempty"`
	// NestedUnit              interface{} `json:"nestedUnit,omitempty"`
	// Offer                   interface{} `json:"offer,omitempty"`
	// PriceUnit               interface{} `json:"price_unit,omitempty"`
	// Modifier                interface{} `json:"modifier,omitempty"`
	// ModifiersCount          interface{} `json:"modifiers_count,omitempty"`
	// Trademark               interface{} `json:"trademark,omitempty"`
	// Series                  interface{} `json:"series,omitempty"`
	// TransitInSettlement     interface{} `json:"transit_in_settlement,omitempty"`
	// WholesalePriceUnit      interface{} `json:"wholesale_price_unit,omitempty"`
	// WholesaleText           interface{} `json:"wholesale_text,omitempty"`
	// ArrivalDate             interface{} `json:"arrivalDate,omitempty"`
}

type Settlement struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Region  string `json:"region"`
	Country string `json:"country"`
	Type    string `json:"type"`
	KladrID string `json:"kladr_id"`
}

type DeliveryPoint struct {
	ID           int64  `json:"id"`
	SettlementID int64  `json:"settlement_id"`
	Address      string `json:"address"`
}
