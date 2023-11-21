package odoo

import (
	"fmt"
)

// ServerConfig represents server.config model.
type ServerConfig struct {
	LastUpdate                         *Time       `xmlrpc:"__last_update,omptempty"`
	Config                             interface{} `xmlrpc:"config,omptempty"`
	CreateDate                         *Time       `xmlrpc:"create_date,omptempty"`
	CreateUid                          *Many2One   `xmlrpc:"create_uid,omptempty"`
	DisplayName                        *String     `xmlrpc:"display_name,omptempty"`
	Id                                 *Int        `xmlrpc:"id,omptempty"`
	OdooIAddonsPath                    *String     `xmlrpc:"odoo_I_addons_path,omptempty"`
	OdooIAdminPasswd                   *String     `xmlrpc:"odoo_I_admin_passwd,omptempty"`
	OdooIConfig                        *String     `xmlrpc:"odoo_I_config,omptempty"`
	OdooICsvInternalSep                *String     `xmlrpc:"odoo_I_csv_internal_sep,omptempty"`
	OdooIDataDir                       *String     `xmlrpc:"odoo_I_data_dir,omptempty"`
	OdooIDbHost                        *String     `xmlrpc:"odoo_I_db_host,omptempty"`
	OdooIDbMaxconn                     *String     `xmlrpc:"odoo_I_db_maxconn,omptempty"`
	OdooIDbName                        *String     `xmlrpc:"odoo_I_db_name,omptempty"`
	OdooIDbPassword                    *String     `xmlrpc:"odoo_I_db_password,omptempty"`
	OdooIDbPort                        *String     `xmlrpc:"odoo_I_db_port,omptempty"`
	OdooIDbSslmode                     *String     `xmlrpc:"odoo_I_db_sslmode,omptempty"`
	OdooIDbTemplate                    *String     `xmlrpc:"odoo_I_db_template,omptempty"`
	OdooIDbUser                        *String     `xmlrpc:"odoo_I_db_user,omptempty"`
	OdooIDbfilter                      *String     `xmlrpc:"odoo_I_dbfilter,omptempty"`
	OdooIDemo                          *String     `xmlrpc:"odoo_I_demo,omptempty"`
	OdooIDevMode                       *String     `xmlrpc:"odoo_I_dev_mode,omptempty"`
	OdooIEmailFrom                     *String     `xmlrpc:"odoo_I_email_from,omptempty"`
	OdooIFromFilter                    *String     `xmlrpc:"odoo_I_from_filter,omptempty"`
	OdooIGeoipDatabase                 *String     `xmlrpc:"odoo_I_geoip_database,omptempty"`
	OdooIGeventPort                    *String     `xmlrpc:"odoo_I_gevent_port,omptempty"`
	OdooIHttpEnable                    *String     `xmlrpc:"odoo_I_http_enable,omptempty"`
	OdooIHttpInterface                 *String     `xmlrpc:"odoo_I_http_interface,omptempty"`
	OdooIHttpPort                      *String     `xmlrpc:"odoo_I_http_port,omptempty"`
	OdooIImportPartial                 *String     `xmlrpc:"odoo_I_import_partial,omptempty"`
	OdooIInit                          *String     `xmlrpc:"odoo_I_init,omptempty"`
	OdooILanguage                      *String     `xmlrpc:"odoo_I_language,omptempty"`
	OdooILimitMemoryHard               *String     `xmlrpc:"odoo_I_limit_memory_hard,omptempty"`
	OdooILimitMemorySoft               *String     `xmlrpc:"odoo_I_limit_memory_soft,omptempty"`
	OdooILimitRequest                  *String     `xmlrpc:"odoo_I_limit_request,omptempty"`
	OdooILimitTimeCpu                  *String     `xmlrpc:"odoo_I_limit_time_cpu,omptempty"`
	OdooILimitTimeReal                 *String     `xmlrpc:"odoo_I_limit_time_real,omptempty"`
	OdooILimitTimeRealCron             *String     `xmlrpc:"odoo_I_limit_time_real_cron,omptempty"`
	OdooIListDb                        *String     `xmlrpc:"odoo_I_list_db,omptempty"`
	OdooILogDb                         *String     `xmlrpc:"odoo_I_log_db,omptempty"`
	OdooILogDbLevel                    *String     `xmlrpc:"odoo_I_log_db_level,omptempty"`
	OdooILogHandler                    *String     `xmlrpc:"odoo_I_log_handler,omptempty"`
	OdooILogLevel                      *String     `xmlrpc:"odoo_I_log_level,omptempty"`
	OdooILogfile                       *String     `xmlrpc:"odoo_I_logfile,omptempty"`
	OdooILongpollingPort               *String     `xmlrpc:"odoo_I_longpolling_port,omptempty"`
	OdooIMaxCronThreads                *String     `xmlrpc:"odoo_I_max_cron_threads,omptempty"`
	OdooIOsvMemoryAgeLimit             *String     `xmlrpc:"odoo_I_osv_memory_age_limit,omptempty"`
	OdooIOsvMemoryCountLimit           *String     `xmlrpc:"odoo_I_osv_memory_count_limit,omptempty"`
	OdooIOverwriteExistingTranslations *String     `xmlrpc:"odoo_I_overwrite_existing_translations,omptempty"`
	OdooIPgPath                        *String     `xmlrpc:"odoo_I_pg_path,omptempty"`
	OdooIPidfile                       *String     `xmlrpc:"odoo_I_pidfile,omptempty"`
	OdooIProxyMode                     *String     `xmlrpc:"odoo_I_proxy_mode,omptempty"`
	OdooIPublisherWarrantyUrl          *String     `xmlrpc:"odoo_I_publisher_warranty_url,omptempty"`
	OdooIReportgz                      *String     `xmlrpc:"odoo_I_reportgz,omptempty"`
	OdooIRootPath                      *String     `xmlrpc:"odoo_I_root_path,omptempty"`
	OdooIRunningEnv                    *String     `xmlrpc:"odoo_I_running_env,omptempty"`
	OdooISave                          *String     `xmlrpc:"odoo_I_save,omptempty"`
	OdooIScreencasts                   *String     `xmlrpc:"odoo_I_screencasts,omptempty"`
	OdooIScreenshots                   *String     `xmlrpc:"odoo_I_screenshots,omptempty"`
	OdooIServerWideModules             *String     `xmlrpc:"odoo_I_server_wide_modules,omptempty"`
	OdooIShellInterface                *String     `xmlrpc:"odoo_I_shell_interface,omptempty"`
	OdooISmtpPassword                  *String     `xmlrpc:"odoo_I_smtp_password,omptempty"`
	OdooISmtpPort                      *String     `xmlrpc:"odoo_I_smtp_port,omptempty"`
	OdooISmtpServer                    *String     `xmlrpc:"odoo_I_smtp_server,omptempty"`
	OdooISmtpSsl                       *String     `xmlrpc:"odoo_I_smtp_ssl,omptempty"`
	OdooISmtpSslCertificateFilename    *String     `xmlrpc:"odoo_I_smtp_ssl_certificate_filename,omptempty"`
	OdooISmtpSslPrivateKeyFilename     *String     `xmlrpc:"odoo_I_smtp_ssl_private_key_filename,omptempty"`
	OdooISmtpUser                      *String     `xmlrpc:"odoo_I_smtp_user,omptempty"`
	OdooIStopAfterInit                 *String     `xmlrpc:"odoo_I_stop_after_init,omptempty"`
	OdooISyslog                        *String     `xmlrpc:"odoo_I_syslog,omptempty"`
	OdooITestEnable                    *String     `xmlrpc:"odoo_I_test_enable,omptempty"`
	OdooITestFile                      *String     `xmlrpc:"odoo_I_test_file,omptempty"`
	OdooITestTags                      *String     `xmlrpc:"odoo_I_test_tags,omptempty"`
	OdooITransientAgeLimit             *String     `xmlrpc:"odoo_I_transient_age_limit,omptempty"`
	OdooITranslateIn                   *String     `xmlrpc:"odoo_I_translate_in,omptempty"`
	OdooITranslateModules              *String     `xmlrpc:"odoo_I_translate_modules,omptempty"`
	OdooITranslateOut                  *String     `xmlrpc:"odoo_I_translate_out,omptempty"`
	OdooIUnaccent                      *String     `xmlrpc:"odoo_I_unaccent,omptempty"`
	OdooIUpdate                        *String     `xmlrpc:"odoo_I_update,omptempty"`
	OdooIUpgradePath                   *String     `xmlrpc:"odoo_I_upgrade_path,omptempty"`
	OdooIWebsocketKeepAliveTimeout     *String     `xmlrpc:"odoo_I_websocket_keep_alive_timeout,omptempty"`
	OdooIWebsocketRateLimitBurst       *String     `xmlrpc:"odoo_I_websocket_rate_limit_burst,omptempty"`
	OdooIWebsocketRateLimitDelay       *String     `xmlrpc:"odoo_I_websocket_rate_limit_delay,omptempty"`
	OdooIWithoutDemo                   *String     `xmlrpc:"odoo_I_without_demo,omptempty"`
	OdooIWorkers                       *String     `xmlrpc:"odoo_I_workers,omptempty"`
	OdooIXSendfile                     *String     `xmlrpc:"odoo_I_x_sendfile,omptempty"`
	QueueJobIChannels                  *String     `xmlrpc:"queue_job_I_channels,omptempty"`
	SystemIArchitecture                *String     `xmlrpc:"system_I_architecture,omptempty"`
	SystemILocale                      *String     `xmlrpc:"system_I_locale,omptempty"`
	SystemILsbRelease                  *String     `xmlrpc:"system_I_lsb_release,omptempty"`
	SystemIOdoo                        *String     `xmlrpc:"system_I_odoo,omptempty"`
	SystemIOsName                      *String     `xmlrpc:"system_I_os_name,omptempty"`
	SystemIPlatform                    *String     `xmlrpc:"system_I_platform,omptempty"`
	SystemIPython                      *String     `xmlrpc:"system_I_python,omptempty"`
	SystemIRelease                     *String     `xmlrpc:"system_I_release,omptempty"`
	SystemIRevision                    *String     `xmlrpc:"system_I_revision,omptempty"`
	SystemIVersion                     *String     `xmlrpc:"system_I_version,omptempty"`
	WriteDate                          *Time       `xmlrpc:"write_date,omptempty"`
	WriteUid                           *Many2One   `xmlrpc:"write_uid,omptempty"`
}

// ServerConfigs represents array of server.config model.
type ServerConfigs []ServerConfig

// ServerConfigModel is the odoo model name.
const ServerConfigModel = "server.config"

// Many2One convert ServerConfig to *Many2One.
func (sc *ServerConfig) Many2One() *Many2One {
	return NewMany2One(sc.Id.Get(), "")
}

// CreateServerConfig creates a new server.config model and returns its id.
func (c *Client) CreateServerConfig(sc *ServerConfig) (int64, error) {
	ids, err := c.CreateServerConfigs([]*ServerConfig{sc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateServerConfig creates a new server.config model and returns its id.
func (c *Client) CreateServerConfigs(scs []*ServerConfig) ([]int64, error) {
	var vv []interface{}
	for _, v := range scs {
		vv = append(vv, v)
	}
	return c.Create(ServerConfigModel, vv)
}

// UpdateServerConfig updates an existing server.config record.
func (c *Client) UpdateServerConfig(sc *ServerConfig) error {
	return c.UpdateServerConfigs([]int64{sc.Id.Get()}, sc)
}

// UpdateServerConfigs updates existing server.config records.
// All records (represented by ids) will be updated by sc values.
func (c *Client) UpdateServerConfigs(ids []int64, sc *ServerConfig) error {
	return c.Update(ServerConfigModel, ids, sc)
}

// DeleteServerConfig deletes an existing server.config record.
func (c *Client) DeleteServerConfig(id int64) error {
	return c.DeleteServerConfigs([]int64{id})
}

// DeleteServerConfigs deletes existing server.config records.
func (c *Client) DeleteServerConfigs(ids []int64) error {
	return c.Delete(ServerConfigModel, ids)
}

// GetServerConfig gets server.config existing record.
func (c *Client) GetServerConfig(id int64) (*ServerConfig, error) {
	scs, err := c.GetServerConfigs([]int64{id})
	if err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of server.config not found", id)
}

// GetServerConfigs gets server.config existing records.
func (c *Client) GetServerConfigs(ids []int64) (*ServerConfigs, error) {
	scs := &ServerConfigs{}
	if err := c.Read(ServerConfigModel, ids, nil, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindServerConfig finds server.config record by querying it with criteria.
func (c *Client) FindServerConfig(criteria *Criteria) (*ServerConfig, error) {
	scs := &ServerConfigs{}
	if err := c.SearchRead(ServerConfigModel, criteria, NewOptions().Limit(1), scs); err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("server.config was not found with criteria %v", criteria)
}

// FindServerConfigs finds server.config records by querying it
// and filtering it with criteria and options.
func (c *Client) FindServerConfigs(criteria *Criteria, options *Options) (*ServerConfigs, error) {
	scs := &ServerConfigs{}
	if err := c.SearchRead(ServerConfigModel, criteria, options, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindServerConfigIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindServerConfigIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ServerConfigModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindServerConfigId finds record id by querying it with criteria.
func (c *Client) FindServerConfigId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ServerConfigModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("server.config was not found with criteria %v and options %v", criteria, options)
}
