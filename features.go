package gofiery

import (
	"net/http"
)

type Features struct {
	APPE                                  string `json:"APPE"`
	APPEImposition                        string `json:"APPE_IMPOSITION"`
	AtlanticFieryCodebase                 string `json:"ATLANTIC_FIERY_CODEBASE"`
	AutoTrapping                          string `json:"AUTOTRAPPING"`
	AdvancedSchedulingForImpose           string `json:"AdvancedSchedulingForImpose"`
	Archive                               string `json:"Archive"`
	BookletVersion                        string `json:"BOOKLETVERSION"`
	BackupRestoreFonts                    string `json:"BackupRestoreFonts"`
	Calibration                           string `json:"CALIBRATION"`
	ColorICCOutputIntent                  string `json:"COLOR_ICC_OUTPUTINTENT"`
	CommonLCD                             string `json:"COMMON_LCD"`
	Compose1FeatureSet                    string `json:"COMPOSE1_FEATURE_SET"`
	CorrectRasterInfo                     string `json:"CORRECT_RASTER_INFO"`
	ColorWise                             string `json:"ColorWise"`
	Compose                               string `json:"Compose"`
	ComposeLocalDongle                    string `json:"ComposeLocalDongle"`
	CustomImpose                          string `json:"CustomImpose"`
	DBPro                                 string `json:"DBPRO"`
	DBProVersion                          string `json:"DBPROVERSION"`
	DefaultServerPresetsVersion           string `json:"DEFAULT_SERVER_PRESETS_VERSION"`
	DirectMobilePrinting                  string `json:"DIRECT_MOBILE_PRINTING"`
	DiskSystem                            string `json:"DISK_SYSTEM"`
	DocumentConversion                    string `json:"DOCUMENT_CONVERSION"`
	DoFontArchive                         string `json:"DOFONTARCHIVE"`
	DualProcessor                         string `json:"DUAL_PROCESSOR"`
	DocumentPagesPrinted                  string `json:"DocumentPagesPrinted"`
	ECSSupport                            string `json:"ECS_SUPPORT"`
	EFISLP                                string `json:"EFI_SLP"`
	EOZ2Codebase                          string `json:"EOZ2_CODEBASE"`
	ExpectedNumRips                       string `json:"EXPECTED_NUM_RIPS"`
	EZRWSCapable                          string `json:"EZRWS_CAPABLE"`
	EnableCancelWaitingToPrint            string `json:"EnableCancelWaitingToPrint"`
	ExtraStatusInfo                       string `json:"ExtraStatusInfo"`
	FeatAuditCapable                      string `json:"FEAT_AUDIT_CAPABLE"`
	FeatConstraintCheck                   string `json:"FEAT_CONSTRAINT_CHECK"`
	FeatContainerReplicatefn              string `json:"FEAT_CONTAINER_REPLICATEFN"`
	FeatDynamicResmanResources            string `json:"FEAT_DYNAMIC_RESMAN_RESOURCES"`
	FeatFct                               string `json:"FEAT_FCT"`
	FeatFieryRipIsOffCapable              string `json:"FEAT_FIERY_RIP_IS_OFF_CAPABLE"`
	FeatFsh                               string `json:"FEAT_FSH"`
	FeatHealthmonitor                     string `json:"FEAT_HEALTHMONITOR"`
	FeatImposeEnhRegmarksBarcode          string `json:"FEAT_IMPOSE_ENH_REGMARKS_BARCODE"`
	FeatIntentionalAssert                 string `json:"FEAT_INTENTIONAL_ASSERT"`
	FeatIntentionalCrash                  string `json:"FEAT_INTENTIONAL_CRASH"`
	FeatJmiArchiveBugFixed                string `json:"FEAT_JMI_ARCHIVE_BUG_FIXED"`
	FeatJobexpertSupportedPdl             string `json:"FEAT_JOBEXPERT_SUPPORTED_PDL"`
	FeatJobflowInstalled                  string `json:"FEAT_JOBFLOW_INSTALLED"`
	FeatJobgroupsSupported                string `json:"FEAT_JOBGROUPS_SUPPORTED"`
	FeatJobExpertEnabled                  string `json:"FEAT_JOB_EXPERT_ENABLED"`
	FeatJobScheduler                      string `json:"FEAT_JOB_SCHEDULER"`
	FeatJpRuntimeSwitchCapable            string `json:"FEAT_JP_RUNTIME_SWITCH_CAPABLE"`
	FeatMultiPluginSupport                string `json:"FEAT_MULTI_PLUGIN_SUPPORT"`
	FeatPclFontCapable                    string `json:"FEAT_PCL_FONT_CAPABLE"`
	FeatPcAllowFamilyDuplicate            string `json:"FEAT_PC_ALLOW_FAMILY_DUPLICATE"`
	FeatPdfvtzipsupport                   string `json:"FEAT_PDFVTZIPSUPPORT"`
	FeatPdfPreflightPro                   string `json:"FEAT_PDF_PREFLIGHT_PRO"`
	FeatPreflightproImportExport          string `json:"FEAT_PREFLIGHTPRO_IMPORT_EXPORT"`
	FeatPreflightproSupportedPdl          string `json:"FEAT_PREFLIGHTPRO_SUPPORTED_PDL"`
	FeatPrintTimeEstimation               string `json:"FEAT_PRINT_TIME_ESTIMATION"`
	FeatRipdsInstalled                    string `json:"FEAT_RIPDS_INSTALLED"`
	FeatSpotColorGroupsToHide             string `json:"FEAT_SPOT_COLOR_GROUPS_TO_HIDE"`
	FeatTabFonts                          string `json:"FEAT_TAB_FONTS"`
	FeatUnifiedBr                         string `json:"FEAT_UNIFIED_BR"`
	FeatUserDefinedPagesize               string `json:"FEAT_USER_DEFINED_PAGESIZE"`
	FeatUseLpDefaultsConf                 string `json:"FEAT_USE_LP_DEFAULTS_CONF"`
	Fiery                                 string `json:"FIERY"`
	FieryPreflight                        string `json:"FIERY_PREFLIGHT"`
	FieryRasterPath                       string `json:"FIERY_RASTER_PATH"`
	FileSearchPath                        string `json:"FILE_SEARCH_PATH"`
	FinerTrayAlignAdjustment              string `json:"FINER_TRAY_ALIGN_ADJUSTMENT"`
	FjdfSupportFeatures                   string `json:"FJDF_SUPPORT_FEATURES"`
	FjdfSupportFeaturesVersion            string `json:"FJDF_SUPPORT_FEATURES_VERSION"`
	FontQueue                             string `json:"FONT_QUEUE"`
	ForceHttpsCapable                     string `json:"FORCE_HTTPS_CAPABLE"`
	ForcePrint                            string `json:"FORCE_PRINT"`
	Freeform                              string `json:"FREEFORM"`
	FierySetup                            string `json:"FierySetup"`
	FreeForm                              string `json:"FreeForm"`
	Ga1Halftone                           string `json:"GA1_HALFTONE"`
	Ga1Hotfolder                          string `json:"GA1_HOTFOLDER"`
	Ga1Iccprofiles                        string `json:"GA1_ICCPROFILES"`
	Ga1Papersim                           string `json:"GA1_PAPERSIM"`
	Ga1Softproof                          string `json:"GA1_SOFTPROOF"`
	Ga1Spoton                             string `json:"GA1_SPOTON"`
	Ga1Substitutecolor                    string `json:"GA1_SUBSTITUTECOLOR"`
	Ga1Trapping                           string `json:"GA1_TRAPPING"`
	Ga2Combineseps                        string `json:"GA2_COMBINESEPS"`
	Ga2Compoverprint                      string `json:"GA2_COMPOVERPRINT"`
	Ga2Compoverprintwspot                 string `json:"GA2_COMPOVERPRINTWSPOT"`
	Ga2Controlstrip                       string `json:"GA2_CONTROLSTRIP"`
	Ga2Dynamicscreen                      string `json:"GA2_DYNAMICSCREEN"`
	Ga2FeatureSet                         string `json:"GA2_FEATURE_SET"`
	Ga2Halftone                           string `json:"GA2_HALFTONE"`
	Ga2Hotfolders                         string `json:"GA2_HOTFOLDERS"`
	Ga2Ieve                               string `json:"GA2_IEVE"`
	Ga2Multicolormap                      string `json:"GA2_MULTICOLORMAP"`
	Ga2Overprint                          string `json:"GA2_OVERPRINT"`
	Ga2Papersimedit                       string `json:"GA2_PAPERSIMEDIT"`
	Ga2Postflight                         string `json:"GA2_POSTFLIGHT"`
	Ga2PostflightAppe                     string `json:"GA2_POSTFLIGHT_APPE"`
	Ga2Rasteredit                         string `json:"GA2_RASTEREDIT"`
	Ga2Softproof                          string `json:"GA2_SOFTPROOF"`
	Ga2Spoton2                            string `json:"GA2_SPOTON2"`
	Ga2Spotpro                            string `json:"GA2_SPOTPRO"`
	Ga2Substitutecolor                    string `json:"GA2_SUBSTITUTECOLOR"`
	Ga2Trapping                           string `json:"GA2_TRAPPING"`
	Ga2Usepdfxprofile                     string `json:"GA2_USEPDFXPROFILE"`
	GaFeatureSet                          string `json:"GA_FEATURE_SET"`
	GmTimeZone                            string `json:"GMTimeZone"`
	GranularBackupRestore                 string `json:"GRANULAR_BACKUP_RESTORE"`
	GreyImposeOverrides                   string `json:"GreyImposeOverrides"`
	HasDirectQueueForFontDownload         string `json:"HAS_DIRECT_QUEUE_FOR_FONT_DOWNLOAD"`
	HeadedSystem                          string `json:"HEADED_SYSTEM"`
	HfprintQueueAcceptsJobProperties      string `json:"HFPrintQueueAcceptsJobProperties"`
	HfEpsFilter                           string `json:"HF_EPS_FILTER"`
	HfJpegFilter                          string `json:"HF_JPEG_FILTER"`
	HfMsofficeFilter                      string `json:"HF_MSOFFICE_FILTER"`
	HfTiffFilter                          string `json:"HF_TIFF_FILTER"`
	Hotfolder                             string `json:"HOTFOLDER"`
	HotFolderPdf16Filter                  string `json:"HOT_FOLDER_PDF1_6_FILTER"`
	HowTo                                 string `json:"HOW_TO"`
	Ifax                                  string `json:"IFAX"`
	ImageviewerColorEdits                 string `json:"IMAGEVIEWER_COLOR_EDITS"`
	ImageviewerCurves                     string `json:"IMAGEVIEWER_CURVES"`
	ImageviewerObjectInspector            string `json:"IMAGEVIEWER_OBJECT_INSPECTOR"`
	Imagewise                             string `json:"IMAGEWISE"`
	Imposeversion                         string `json:"IMPOSEVERSION"`
	ImposeFeatureFinishermargin           string `json:"IMPOSE_FEATURE_FINISHERMARGIN"`
	ImposeHotFolderCustomImposition       string `json:"IMPOSE_HOT_FOLDER_CUSTOM_IMPOSITION"`
	ImposeImprovedPrescientDetection      string `json:"IMPOSE_IMPROVED_PRESCIENT_DETECTION"`
	Imposition                            string `json:"IMPOSITION"`
	IntentTicketing                       string `json:"INTENT_TICKETING"`
	Ipsec                                 string `json:"IPSEC"`
	Ipv6                                  string `json:"IPV6"`
	Imposition1                           string `json:"Imposition"`
	ImpositionWd                          string `json:"ImpositionWD"`
	JobflowSupport                        string `json:"JOBFLOW_SUPPORT"`
	JobErrorReportSupport                 string `json:"JOB_ERROR_REPORT_SUPPORT"`
	JobParallelRushrip                    string `json:"JOB_PARALLEL_RUSHRIP"`
	JobPreset                             string `json:"JOB_PRESET"`
	JobLogMmInfo                          string `json:"JobLogMMInfo"`
	JobOverrides                          string `json:"JobOverrides"`
	JobParallel                           string `json:"JobParallel"`
	JobReorder                            string `json:"JobReorder"`
	JobReorderOrderColumn                 string `json:"JobReorderOrderColumn"`
	LcdPages                              string `json:"LCDPages"`
	Maxcopies                             string `json:"MAXCOPIES"`
	MediaThumbdrive                       string `json:"MEDIA_THUMBDRIVE"`
	Multirip                              string `json:"MULTIRIP"`
	ManualDuplex                          string `json:"ManualDuplex"`
	MediaServer                           string `json:"MediaServer"`
	MemberPrinting                        string `json:"MemberPrinting"`
	NetworkFtpPassiveMode                 string `json:"NETWORK_FTP_PASSIVE_MODE"`
	NetworkMailEnhancedOptions            string `json:"NETWORK_MAIL_ENHANCED_OPTIONS"`
	NetworkPortfilterHttpssloff           string `json:"NETWORK_PORTFILTER_HTTPSSLOFF"`
	NetworkPortFiltering                  string `json:"NETWORK_PORT_FILTERING"`
	NetworkSnmpCommunitynameSettings      string `json:"NETWORK_SNMP_COMMUNITYNAME_SETTINGS"`
	NoOemlogo                             string `json:"NoOEMlogo"`
	NoVdpimposeuiForFf1                   string `json:"NoVDPImposeUI_ForFF1"`
	PatchQueue                            string `json:"PATCH_QUEUE"`
	PclFontsAgfa                          string `json:"PCL_FONTS_AGFA"`
	PcBackuprestoreUiEnable               string `json:"PC_BACKUPRESTORE_UI_ENABLE"`
	PcColorProfile                        string `json:"PC_ColorProfile"`
	PcFloatingPointDimension              string `json:"PC_FLOATING_POINT_DIMENSION"`
	PcImportExport                        string `json:"PC_ImportExport"`
	PcMmSelectiveDisable                  string `json:"PC_MM_SELECTIVE_DISABLE"`
	PcMediaAddDeleteDuplicate             string `json:"PC_MediaAddDeleteDuplicate"`
	PcPaperCatalogFeature                 string `json:"PC_PAPER_CATALOG_FEATURE"`
	PcSmartMedia                          string `json:"PC_SmartMedia"`
	PcTrayAssociation                     string `json:"PC_TrayAssociation"`
	PcAdd                                 string `json:"PC_add"`
	PcDelete                              string `json:"PC_delete"`
	PcDuplicate                           string `json:"PC_duplicate"`
	PcExport                              string `json:"PC_export"`
	PcImport                              string `json:"PC_import"`
	PcReset                               string `json:"PC_reset"`
	PdfVt                                 string `json:"PDF_VT"`
	Port9100                              string `json:"PORT_9100"`
	Postscript                            string `json:"POSTSCRIPT"`
	PpdNoConversion                       string `json:"PPD_NO_CONVERSION"`
	Ppml                                  string `json:"PPML"`
	PreflightHfAndVp                      string `json:"PREFLIGHT_HF_AND_VP"`
	PrinterControlLanguage                string `json:"PRINTER_CONTROL_LANGUAGE"`
	PrintAndDelete                        string `json:"PRINT_AND_DELETE"`
	Production1CancelOnMismatch           string `json:"PRODUCTION1_CANCEL_ON_MISMATCH"`
	Production1FeatureSet                 string `json:"PRODUCTION1_FEATURE_SET"`
	Production1Hotfolder                  string `json:"PRODUCTION1_HOTFOLDER"`
	Production1JobReorder                 string `json:"PRODUCTION1_JOB_REORDER"`
	Production1JobReorderPrintnext        string `json:"PRODUCTION1_JOB_REORDER_PRINTNEXT"`
	Production1JobReorderProcessnext      string `json:"PRODUCTION1_JOB_REORDER_PROCESSNEXT"`
	Production1MmPreviews                 string `json:"PRODUCTION1_MM_PREVIEWS"`
	Production1PaperCatalog               string `json:"PRODUCTION1_PAPER_CATALOG"`
	Production1RushJob                    string `json:"PRODUCTION1_RUSH_JOB"`
	Production1SamplePrint                string `json:"PRODUCTION1_SAMPLE_PRINT"`
	Production1Schedule                   string `json:"PRODUCTION1_SCHEDULE"`
	Production1SuspendOnMismatch          string `json:"PRODUCTION1_SUSPEND_ON_MISMATCH"`
	Production1Tabs                       string `json:"PRODUCTION1_TABS"`
	ProofPrintEnabled                     string `json:"PROOF_PRINT_ENABLED"`
	Preflight                             string `json:"Preflight"`
	PreflightBinaryOutput                 string `json:"PreflightBinaryOutput"`
	PrintNext                             string `json:"PrintNext"`
	PrintRush                             string `json:"PrintRush"`
	PrintSchedule                         string `json:"PrintSchedule"`
	ProcessNext                           string `json:"ProcessNext"`
	Qdb                                   string `json:"QDB"`
	RemoteDesktop                         string `json:"REMOTE_DESKTOP"`
	RipProgressUsePages                   string `json:"RIP_PROGRESS_USE_PAGES"`
	SecureDesktop                         string `json:"SECURE_DESKTOP"`
	SecureErase                           string `json:"SECURE_ERASE"`
	Security1FeatureSet                   string `json:"SECURITY1_FEATURE_SET"`
	Security1SafeErase                    string `json:"SECURITY1_SAFE_ERASE"`
	Security1SafeEraseV2                  string `json:"SECURITY1_SAFE_ERASE_V2"`
	Sendjoblog                            string `json:"SENDJOBLOG"`
	ShowLacRegistration                   string `json:"SHOW_LAC_REGISTRATION"`
	Sntp                                  string `json:"SNTP"`
	SoftwareLicensing                     string `json:"SOFTWARE_LICENSING"`
	SpdMapping                            string `json:"SPD_MAPPING"`
	SpoolPreviewSupported                 string `json:"SPOOL_PREVIEW_SUPPORTED"`
	SpotonJpWorkflow                      string `json:"SPOTON_JP_WORKFLOW"`
	SystemupdateSupport                   string `json:"SYSTEMUPDATE_SUPPORT"`
	ServerGeneratedThumbnails             string `json:"Server generated thumbnails"`
	SoftProof                             string `json:"SoftProof"`
	SuspendOnMismatch                     string `json:"SuspendOnMismatch"`
	SuspendOnMismatchV2                   string `json:"SuspendOnMismatchV2"`
	SuspendResume                         string `json:"SuspendResume"`
	TbiconStartStop                       string `json:"TBICON_StartStop"`
	TroyCodebase                          string `json:"TROY_CODEBASE"`
	TabsPrintingOn                        string `json:"TabsPrintingOn"`
	TrayAlignment                         string `json:"TrayAlignment"`
	TrayAlignmentTestPage                 string `json:"TrayAlignmentTestPage"`
	UserAuthentication                    string `json:"USER_AUTHENTICATION"`
	UserAuthenticationAdmin               string `json:"USER_AUTHENTICATION_ADMIN"`
	UserAuthenticationScanTimeUsers       string `json:"USER_AUTHENTICATION_SCAN_TIME_USERS"`
	UserDefinedPagesize                   string `json:"USER_DEFINED_PAGESIZE"`
	Utf8Version                           string `json:"UTF8_VERSION"`
	Vdpimpose                             string `json:"VDPImpose"`
	VdpImposition                         string `json:"VDP_IMPOSITION"`
	VirtualPrinters                       string `json:"VIRTUAL_PRINTERS"`
	Vps                                   string `json:"VPS"`
	VirtualPrinter                        string `json:"VirtualPrinter"`
	VirtualPrinterMaximumNumberOfPrinters string `json:"VirtualPrinterMaximumNumberOfPrinters"`
	WebtoolsVersion                       string `json:"WEBTOOLS_VERSION"`
	Xps                                   string `json:"XPS"`
	ClearCoatNColor                       string `json:"clear_coat/n_color"`
	CwsSetupDeviceType                    string `json:"cws setup device type"`
	CwsSetupKey                           string `json:"cws setup key"`
	DebugBuild                            string `json:"debug build"`
	HarmonyLibVersion                     string `json:"harmony lib version"`
	HeldJobReordering                     string `json:"held job reordering"`
	IpdsReady                             string `json:"ipds_ready"`
	JobReordering                         string `json:"job reordering"`
	NotActivated                          string `json:"not activated"`
	OperatingSystem                       string `json:"operating system"`
	PclSupport                            string `json:"pcl support"`
	PdfSupport                            string `json:"pdf support"`
	PostscriptSupport                     string `json:"postscript support"`
	PpmlSupport                           string `json:"ppml support"`
	ProductFamily                         string `json:"product family"`
	ProgressIndicator                     string `json:"progress indicator"`
	RacingSmartRip                        string `json:"racing smart rip"`
	RipAndHoldGeneratesNewJob             string `json:"rip and hold generates new job"`
}

// GetFeatures retrieves information about available features and capabilities on the Fiery server.
// Returns a pointer to Features containing the feature set supported by the server.
func (fc *FieryClient) GetFeatures() *Features {
	var features Features
	response := fc.Run(fc.Endpoint("features"), http.MethodGet)
	if response == nil {
		panic("Failed to retrieve features from the Fiery server.")
	}
	features = response.data.item.(Features)
	return &features
}

// HasFeature checks if a specific feature is enabled on the Fiery server.
// Returns true if the feature value is "yes", false otherwise or if features cannot be retrieved.
func (fc *FieryClient) HasFeature(field func(*Features) string) bool {
	features := fc.GetFeatures()
	if features == nil {
		return false
	}
	return field(features) == "yes"
}
