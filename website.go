package odoo

import (
	"fmt"
)

// Website represents website model.
type Website struct {
	LastUpdate                      *Time      `xmlrpc:"__last_update,omptempty"`
	AccountOnCheckout               *Selection `xmlrpc:"account_on_checkout,omptempty"`
	AddToCartAction                 *Selection `xmlrpc:"add_to_cart_action,omptempty"`
	AllPricelistIds                 *Relation  `xmlrpc:"all_pricelist_ids,omptempty"`
	AuthSignupUninvited             *Selection `xmlrpc:"auth_signup_uninvited,omptempty"`
	AutoRedirectLang                *Bool      `xmlrpc:"auto_redirect_lang,omptempty"`
	CartAbandonedDelay              *Float     `xmlrpc:"cart_abandoned_delay,omptempty"`
	CartRecoveryMailTemplateId      *Many2One  `xmlrpc:"cart_recovery_mail_template_id,omptempty"`
	CdnActivated                    *Bool      `xmlrpc:"cdn_activated,omptempty"`
	CdnFilters                      *String    `xmlrpc:"cdn_filters,omptempty"`
	CdnUrl                          *String    `xmlrpc:"cdn_url,omptempty"`
	ChannelId                       *Many2One  `xmlrpc:"channel_id,omptempty"`
	CompanyId                       *Many2One  `xmlrpc:"company_id,omptempty"`
	ConfiguratorDone                *Bool      `xmlrpc:"configurator_done,omptempty"`
	ContactUsButtonUrl              *String    `xmlrpc:"contact_us_button_url,omptempty"`
	CookiesBar                      *Bool      `xmlrpc:"cookies_bar,omptempty"`
	CreateDate                      *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                       *Many2One  `xmlrpc:"create_uid,omptempty"`
	CrmDefaultTeamId                *Many2One  `xmlrpc:"crm_default_team_id,omptempty"`
	CrmDefaultUserId                *Many2One  `xmlrpc:"crm_default_user_id,omptempty"`
	CurrencyId                      *Many2One  `xmlrpc:"currency_id,omptempty"`
	CustomCodeFooter                *String    `xmlrpc:"custom_code_footer,omptempty"`
	CustomCodeHead                  *String    `xmlrpc:"custom_code_head,omptempty"`
	DefaultLangId                   *Many2One  `xmlrpc:"default_lang_id,omptempty"`
	DisplayName                     *String    `xmlrpc:"display_name,omptempty"`
	Domain                          *String    `xmlrpc:"domain,omptempty"`
	EnabledPortalReorderButton      *Bool      `xmlrpc:"enabled_portal_reorder_button,omptempty"`
	Favicon                         *String    `xmlrpc:"favicon,omptempty"`
	FirebaseAdminKeyFile            *String    `xmlrpc:"firebase_admin_key_file,omptempty"`
	FirebaseEnablePushNotifications *Bool      `xmlrpc:"firebase_enable_push_notifications,omptempty"`
	FirebaseProjectId               *String    `xmlrpc:"firebase_project_id,omptempty"`
	FirebasePushCertificateKey      *String    `xmlrpc:"firebase_push_certificate_key,omptempty"`
	FirebaseSenderId                *String    `xmlrpc:"firebase_sender_id,omptempty"`
	FirebaseUseOwnAccount           *Bool      `xmlrpc:"firebase_use_own_account,omptempty"`
	FirebaseWebApiKey               *String    `xmlrpc:"firebase_web_api_key,omptempty"`
	GoogleAnalyticsKey              *String    `xmlrpc:"google_analytics_key,omptempty"`
	GoogleMapsApiKey                *String    `xmlrpc:"google_maps_api_key,omptempty"`
	GoogleSearchConsole             *String    `xmlrpc:"google_search_console,omptempty"`
	HasSocialDefaultImage           *Bool      `xmlrpc:"has_social_default_image,omptempty"`
	HomepageUrl                     *String    `xmlrpc:"homepage_url,omptempty"`
	Id                              *Int       `xmlrpc:"id,omptempty"`
	KarmaProfileMin                 *Int       `xmlrpc:"karma_profile_min,omptempty"`
	LanguageCount                   *Int       `xmlrpc:"language_count,omptempty"`
	LanguageIds                     *Relation  `xmlrpc:"language_ids,omptempty"`
	Logo                            *String    `xmlrpc:"logo,omptempty"`
	MenuId                          *Many2One  `xmlrpc:"menu_id,omptempty"`
	Name                            *String    `xmlrpc:"name,omptempty"`
	NotificationRequestBody         *String    `xmlrpc:"notification_request_body,omptempty"`
	NotificationRequestDelay        *Int       `xmlrpc:"notification_request_delay,omptempty"`
	NotificationRequestIcon         *String    `xmlrpc:"notification_request_icon,omptempty"`
	NotificationRequestTitle        *String    `xmlrpc:"notification_request_title,omptempty"`
	PartnerId                       *Many2One  `xmlrpc:"partner_id,omptempty"`
	PlausibleSharedKey              *String    `xmlrpc:"plausible_shared_key,omptempty"`
	PlausibleSite                   *String    `xmlrpc:"plausible_site,omptempty"`
	PreventZeroPriceSale            *Bool      `xmlrpc:"prevent_zero_price_sale,omptempty"`
	PreventZeroPriceSaleText        *String    `xmlrpc:"prevent_zero_price_sale_text,omptempty"`
	PricelistId                     *Many2One  `xmlrpc:"pricelist_id,omptempty"`
	PricelistIds                    *Relation  `xmlrpc:"pricelist_ids,omptempty"`
	ProductPageGridColumns          *Int       `xmlrpc:"product_page_grid_columns,omptempty"`
	ProductPageImageLayout          *Selection `xmlrpc:"product_page_image_layout,omptempty"`
	ProductPageImageSpacing         *Selection `xmlrpc:"product_page_image_spacing,omptempty"`
	ProductPageImageWidth           *Selection `xmlrpc:"product_page_image_width,omptempty"`
	RobotsTxt                       *String    `xmlrpc:"robots_txt,omptempty"`
	SalespersonId                   *Many2One  `xmlrpc:"salesperson_id,omptempty"`
	SalesteamId                     *Many2One  `xmlrpc:"salesteam_id,omptempty"`
	SendAbandonedCartEmail          *Bool      `xmlrpc:"send_abandoned_cart_email,omptempty"`
	Sequence                        *Int       `xmlrpc:"sequence,omptempty"`
	ShopDefaultSort                 *Selection `xmlrpc:"shop_default_sort,omptempty"`
	ShopExtraFieldIds               *Relation  `xmlrpc:"shop_extra_field_ids,omptempty"`
	ShopPpg                         *Int       `xmlrpc:"shop_ppg,omptempty"`
	ShopPpr                         *Int       `xmlrpc:"shop_ppr,omptempty"`
	SocialDefaultImage              *String    `xmlrpc:"social_default_image,omptempty"`
	SocialFacebook                  *String    `xmlrpc:"social_facebook,omptempty"`
	SocialGithub                    *String    `xmlrpc:"social_github,omptempty"`
	SocialInstagram                 *String    `xmlrpc:"social_instagram,omptempty"`
	SocialLinkedin                  *String    `xmlrpc:"social_linkedin,omptempty"`
	SocialTwitter                   *String    `xmlrpc:"social_twitter,omptempty"`
	SocialYoutube                   *String    `xmlrpc:"social_youtube,omptempty"`
	SpecificUserAccount             *Bool      `xmlrpc:"specific_user_account,omptempty"`
	ThemeId                         *Many2One  `xmlrpc:"theme_id,omptempty"`
	UserId                          *Many2One  `xmlrpc:"user_id,omptempty"`
	WarehouseId                     *Many2One  `xmlrpc:"warehouse_id,omptempty"`
	WebsiteSlideGoogleAppKey        *String    `xmlrpc:"website_slide_google_app_key,omptempty"`
	WriteDate                       *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                        *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// Websites represents array of website model.
type Websites []Website

// WebsiteModel is the odoo model name.
const WebsiteModel = "website"

// Many2One convert Website to *Many2One.
func (w *Website) Many2One() *Many2One {
	return NewMany2One(w.Id.Get(), "")
}

// CreateWebsite creates a new website model and returns its id.
func (c *Client) CreateWebsite(w *Website) (int64, error) {
	ids, err := c.CreateWebsites([]*Website{w})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsite creates a new website model and returns its id.
func (c *Client) CreateWebsites(ws []*Website) ([]int64, error) {
	var vv []interface{}
	for _, v := range ws {
		vv = append(vv, v)
	}
	return c.Create(WebsiteModel, vv)
}

// UpdateWebsite updates an existing website record.
func (c *Client) UpdateWebsite(w *Website) error {
	return c.UpdateWebsites([]int64{w.Id.Get()}, w)
}

// UpdateWebsites updates existing website records.
// All records (represented by ids) will be updated by w values.
func (c *Client) UpdateWebsites(ids []int64, w *Website) error {
	return c.Update(WebsiteModel, ids, w)
}

// DeleteWebsite deletes an existing website record.
func (c *Client) DeleteWebsite(id int64) error {
	return c.DeleteWebsites([]int64{id})
}

// DeleteWebsites deletes existing website records.
func (c *Client) DeleteWebsites(ids []int64) error {
	return c.Delete(WebsiteModel, ids)
}

// GetWebsite gets website existing record.
func (c *Client) GetWebsite(id int64) (*Website, error) {
	ws, err := c.GetWebsites([]int64{id})
	if err != nil {
		return nil, err
	}
	if ws != nil && len(*ws) > 0 {
		return &((*ws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of website not found", id)
}

// GetWebsites gets website existing records.
func (c *Client) GetWebsites(ids []int64) (*Websites, error) {
	ws := &Websites{}
	if err := c.Read(WebsiteModel, ids, nil, ws); err != nil {
		return nil, err
	}
	return ws, nil
}

// FindWebsite finds website record by querying it with criteria.
func (c *Client) FindWebsite(criteria *Criteria) (*Website, error) {
	ws := &Websites{}
	if err := c.SearchRead(WebsiteModel, criteria, NewOptions().Limit(1), ws); err != nil {
		return nil, err
	}
	if ws != nil && len(*ws) > 0 {
		return &((*ws)[0]), nil
	}
	return nil, fmt.Errorf("website was not found with criteria %v", criteria)
}

// FindWebsites finds website records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsites(criteria *Criteria, options *Options) (*Websites, error) {
	ws := &Websites{}
	if err := c.SearchRead(WebsiteModel, criteria, options, ws); err != nil {
		return nil, err
	}
	return ws, nil
}

// FindWebsiteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebsiteModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebsiteId finds record id by querying it with criteria.
func (c *Client) FindWebsiteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("website was not found with criteria %v and options %v", criteria, options)
}
