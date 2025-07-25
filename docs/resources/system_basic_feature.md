---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_system_basic_feature Resource - terraform-provider-sdwan"
subcategory: "Features - System"
description: |-
  This resource can manage a System Basic Feature.
  Minimum SD-WAN Manager version: 20.12.0
---

# sdwan_system_basic_feature (Resource)

This resource can manage a System Basic Feature.
  - Minimum SD-WAN Manager version: `20.12.0`

## Example Usage

```terraform
resource "sdwan_system_basic_feature" "example" {
  name                   = "Example"
  description            = "My Example"
  feature_profile_id     = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  timezone               = "UTC"
  config_description     = "example"
  location               = "example"
  gps_longitude          = -77
  gps_latitude           = 38
  gps_geo_fencing_enable = true
  gps_geo_fencing_range  = 100
  gps_sms_enable         = true
  gps_sms_mobile_numbers = [
    {
      number = "+11111233"
    }
  ]
  device_groups              = ["example"]
  controller_groups          = [1]
  overlay_id                 = 1
  port_offset                = 19
  port_hopping               = true
  control_session_pps        = 300
  track_transport            = true
  track_interface_tag        = 2
  console_baud_rate          = "9600"
  max_omp_sessions           = 24
  multi_tenant               = false
  track_default_gateway      = true
  admin_tech_on_failure      = true
  idle_timeout               = 10
  on_demand_enable           = true
  on_demand_idle_timeout     = 10
  transport_gateway          = false
  enhanced_app_aware_routing = "aggressive"
  site_types                 = ["type-1"]
  affinity_group_number      = 1
  affinity_group_preferences = [1]
  affinity_preference_auto   = false
  affinity_per_vrfs = [
    {
      affinity_group_number = 1
      vrf_range             = "123-456"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `feature_profile_id` (String) Feature Profile ID
- `name` (String) The name of the Feature

### Optional

- `admin_tech_on_failure` (Boolean) Collect admin-tech before reboot due to daemon failure
  - Default value: `true`
- `admin_tech_on_failure_variable` (String) Variable name
- `affinity_group_number` (Number) Affinity Group Number
  - Range: `1`-`63`
- `affinity_group_number_variable` (String) Variable name
- `affinity_group_preferences` (Set of Number) Affinity Group Preference
- `affinity_group_preferences_variable` (String) Variable name
- `affinity_per_vrfs` (Attributes List) Affinity Group Number for VRFs (see [below for nested schema](#nestedatt--affinity_per_vrfs))
- `affinity_preference_auto` (Boolean) Affinity Group Preference Auto
  - Default value: `false`
- `affinity_preference_auto_variable` (String) Variable name
- `config_description` (String) Set a text description of the device
- `config_description_variable` (String) Variable name
- `console_baud_rate` (String) Set the console baud rate
  - Choices: `1200`, `2400`, `4800`, `9600`, `19200`, `38400`, `57600`, `115200`
  - Default value: `9600`
- `console_baud_rate_variable` (String) Variable name
- `control_session_pps` (Number) Set the policer rate for control sessions
  - Range: `1`-`65535`
  - Default value: `300`
- `control_session_pps_variable` (String) Variable name
- `controller_groups` (Set of Number) Configure a list of comma-separated controller groups
- `controller_groups_variable` (String) Variable name
- `description` (String) The description of the Feature
- `device_groups` (Set of String) Device groups
- `device_groups_variable` (String) Variable name
- `enhanced_app_aware_routing` (String) Enable SLA Dampening and Enhanced App Routing.
  - Choices: `disabled`, `aggressive`, `moderate`, `conservative`
  - Default value: `disabled`
- `enhanced_app_aware_routing_variable` (String) Variable name
- `gps_geo_fencing_enable` (Boolean) Enable Geo fencing
  - Default value: `false`
- `gps_geo_fencing_range` (Number) Set the device’s geo fencing range
  - Range: `100`-`10000`
  - Default value: `100`
- `gps_geo_fencing_range_variable` (String) Variable name
- `gps_latitude` (Number) Set the device physical latitude
  - Range: `-90`-`90`
- `gps_latitude_variable` (String) Variable name
- `gps_longitude` (Number) Set the device physical longitude
  - Range: `-180`-`180`
- `gps_longitude_variable` (String) Variable name
- `gps_sms_enable` (Boolean) Enable device’s geo fencing SMS
  - Default value: `false`
- `gps_sms_mobile_numbers` (Attributes List) Set device’s geo fencing SMS phone number (see [below for nested schema](#nestedatt--gps_sms_mobile_numbers))
- `idle_timeout` (Number) Idle CLI timeout in minutes
  - Range: `0`-`300`
- `idle_timeout_variable` (String) Variable name
- `location` (String) Set the location of the device
- `location_variable` (String) Variable name
- `max_omp_sessions` (Number) Set the maximum number of OMP sessions <1..100> the device can have
  - Range: `1`-`100`
- `max_omp_sessions_variable` (String) Variable name
- `multi_tenant` (Boolean) Device is multi-tenant
  - Default value: `false`
- `multi_tenant_variable` (String) Variable name
- `on_demand_enable` (Boolean) Enable or disable On-demand Tunnel
  - Default value: `false`
- `on_demand_enable_variable` (String) Variable name
- `on_demand_idle_timeout` (Number) Set the idle timeout for on-demand tunnels
  - Range: `1`-`65535`
  - Default value: `10`
- `on_demand_idle_timeout_variable` (String) Variable name
- `overlay_id` (Number) Set the Overlay ID
  - Range: `1`-`4294967295`
  - Default value: `1`
- `overlay_id_variable` (String) Variable name
- `port_hopping` (Boolean) Enable port hopping
  - Default value: `true`
- `port_hopping_variable` (String) Variable name
- `port_offset` (Number) Set the TLOC port offset when multiple devices are behind a NAT
  - Range: `0`-`19`
  - Default value: `0`
- `port_offset_variable` (String) Variable name
- `site_types` (Set of String) Site Type
- `site_types_variable` (String) Variable name
- `timezone` (String) Set the timezone
  - Choices: `Europe/Andorra`, `Asia/Dubai`, `Asia/Kabul`, `America/Antigua`, `America/Anguilla`, `Europe/Tirane`, `Asia/Yerevan`, `Africa/Luanda`, `Antarctica/McMurdo`, `Antarctica/Rothera`, `Antarctica/Palmer`, `Antarctica/Mawson`, `Antarctica/Davis`, `Antarctica/Casey`, `Antarctica/Vostok`, `Antarctica/DumontDUrville`, `Antarctica/Syowa`, `America/Argentina/Buenos_Aires`, `America/Argentina/Cordoba`, `America/Argentina/Salta`, `America/Argentina/Jujuy`, `America/Argentina/Tucuman`, `America/Argentina/Catamarca`, `America/Argentina/La_Rioja`, `America/Argentina/San_Juan`, `America/Argentina/Mendoza`, `America/Argentina/San_Luis`, `America/Argentina/Rio_Gallegos`, `America/Argentina/Ushuaia`, `Pacific/Pago_Pago`, `Europe/Vienna`, `Australia/Lord_Howe`, `Antarctica/Macquarie`, `Australia/Hobart`, `Australia/Currie`, `Australia/Melbourne`, `Australia/Sydney`, `Australia/Broken_Hill`, `Australia/Brisbane`, `Australia/Lindeman`, `Australia/Adelaide`, `Australia/Darwin`, `Australia/Perth`, `Australia/Eucla`, `America/Aruba`, `Europe/Mariehamn`, `Asia/Baku`, `Europe/Sarajevo`, `America/Barbados`, `Asia/Dhaka`, `Europe/Brussels`, `Africa/Ouagadougou`, `Europe/Sofia`, `Asia/Bahrain`, `Africa/Bujumbura`, `Africa/Porto-Novo`, `America/St_Barthelemy`, `Atlantic/Bermuda`, `Asia/Brunei`, `America/La_Paz`, `America/Kralendijk`, `America/Noronha`, `America/Belem`, `America/Fortaleza`, `America/Recife`, `America/Araguaina`, `America/Maceio`, `America/Bahia`, `America/Sao_Paulo`, `America/Campo_Grande`, `America/Cuiaba`, `America/Santarem`, `America/Porto_Velho`, `America/Boa_Vista`, `America/Manaus`, `America/Eirunepe`, `America/Rio_Branco`, `America/Nassau`, `Asia/Thimphu`, `Africa/Gaborone`, `Europe/Minsk`, `America/Belize`, `America/St_Johns`, `America/Halifax`, `America/Glace_Bay`, `America/Moncton`, `America/Goose_Bay`, `America/Blanc-Sablon`, `America/Toronto`, `America/Nipigon`, `America/Thunder_Bay`, `America/Iqaluit`, `America/Pangnirtung`, `America/Resolute`, `America/Atikokan`, `America/Rankin_Inlet`, `America/Winnipeg`, `America/Rainy_River`, `America/Regina`, `America/Swift_Current`, `America/Edmonton`, `America/Cambridge_Bay`, `America/Yellowknife`, `America/Inuvik`, `America/Creston`, `America/Dawson_Creek`, `America/Vancouver`, `America/Whitehorse`, `America/Dawson`, `Indian/Cocos`, `Africa/Kinshasa`, `Africa/Lubumbashi`, `Africa/Bangui`, `Africa/Brazzaville`, `Europe/Zurich`, `Africa/Abidjan`, `Pacific/Rarotonga`, `America/Santiago`, `Pacific/Easter`, `Africa/Douala`, `Asia/Shanghai`, `Asia/Harbin`, `Asia/Chongqing`, `Asia/Urumqi`, `Asia/Kashgar`, `America/Bogota`, `America/Costa_Rica`, `America/Havana`, `Atlantic/Cape_Verde`, `America/Curacao`, `Indian/Christmas`, `Asia/Nicosia`, `Europe/Prague`, `Europe/Berlin`, `Europe/Busingen`, `Africa/Djibouti`, `Europe/Copenhagen`, `America/Dominica`, `America/Santo_Domingo`, `Africa/Algiers`, `America/Guayaquil`, `Pacific/Galapagos`, `Europe/Tallinn`, `Africa/Cairo`, `Africa/El_Aaiun`, `Africa/Asmara`, `Europe/Madrid`, `Africa/Ceuta`, `Atlantic/Canary`, `Africa/Addis_Ababa`, `Europe/Helsinki`, `Pacific/Fiji`, `Atlantic/Stanley`, `Pacific/Chuuk`, `Pacific/Pohnpei`, `Pacific/Kosrae`, `Atlantic/Faroe`, `Europe/Paris`, `Africa/Libreville`, `Europe/London`, `America/Grenada`, `Asia/Tbilisi`, `America/Cayenne`, `Europe/Guernsey`, `Africa/Accra`, `Europe/Gibraltar`, `America/Godthab`, `America/Danmarkshavn`, `America/Scoresbysund`, `America/Thule`, `Africa/Banjul`, `Africa/Conakry`, `America/Guadeloupe`, `Africa/Malabo`, `Europe/Athens`, `Atlantic/South_Georgia`, `America/Guatemala`, `Pacific/Guam`, `Africa/Bissau`, `America/Guyana`, `Asia/Hong_Kong`, `America/Tegucigalpa`, `Europe/Zagreb`, `America/Port-au-Prince`, `Europe/Budapest`, `Asia/Jakarta`, `Asia/Pontianak`, `Asia/Makassar`, `Asia/Jayapura`, `Europe/Dublin`, `Asia/Jerusalem`, `Europe/Isle_of_Man`, `Asia/Kolkata`, `Indian/Chagos`, `Asia/Baghdad`, `Asia/Tehran`, `Atlantic/Reykjavik`, `Europe/Rome`, `Europe/Jersey`, `America/Jamaica`, `Asia/Amman`, `Asia/Tokyo`, `Africa/Nairobi`, `Asia/Bishkek`, `Asia/Phnom_Penh`, `Pacific/Tarawa`, `Pacific/Enderbury`, `Pacific/Kiritimati`, `Indian/Comoro`, `America/St_Kitts`, `Asia/Pyongyang`, `Asia/Seoul`, `Asia/Kuwait`, `America/Cayman`, `Asia/Almaty`, `Asia/Qyzylorda`, `Asia/Aqtobe`, `Asia/Aqtau`, `Asia/Oral`, `Asia/Vientiane`, `Asia/Beirut`, `America/St_Lucia`, `Europe/Vaduz`, `Asia/Colombo`, `Africa/Monrovia`, `Africa/Maseru`, `Europe/Vilnius`, `Europe/Luxembourg`, `Europe/Riga`, `Africa/Tripoli`, `Africa/Casablanca`, `Europe/Monaco`, `Europe/Chisinau`, `Europe/Podgorica`, `America/Marigot`, `Indian/Antananarivo`, `Pacific/Majuro`, `Pacific/Kwajalein`, `Europe/Skopje`, `Africa/Bamako`, `Asia/Rangoon`, `Asia/Ulaanbaatar`, `Asia/Hovd`, `Asia/Choibalsan`, `Asia/Macau`, `Pacific/Saipan`, `America/Martinique`, `Africa/Nouakchott`, `America/Montserrat`, `Europe/Malta`, `Indian/Mauritius`, `Indian/Maldives`, `Africa/Blantyre`, `America/Mexico_City`, `America/Cancun`, `America/Merida`, `America/Monterrey`, `America/Matamoros`, `America/Mazatlan`, `America/Chihuahua`, `America/Ojinaga`, `America/Hermosillo`, `America/Tijuana`, `America/Santa_Isabel`, `America/Bahia_Banderas`, `Asia/Kuala_Lumpur`, `Asia/Kuching`, `Africa/Maputo`, `Africa/Windhoek`, `Pacific/Noumea`, `Africa/Niamey`, `Pacific/Norfolk`, `Africa/Lagos`, `America/Managua`, `Europe/Amsterdam`, `Europe/Oslo`, `Asia/Kathmandu`, `Pacific/Nauru`, `Pacific/Niue`, `Pacific/Auckland`, `Pacific/Chatham`, `Asia/Muscat`, `America/Panama`, `America/Lima`, `Pacific/Tahiti`, `Pacific/Marquesas`, `Pacific/Gambier`, `Pacific/Port_Moresby`, `Asia/Manila`, `Asia/Karachi`, `Europe/Warsaw`, `America/Miquelon`, `Pacific/Pitcairn`, `America/Puerto_Rico`, `Asia/Gaza`, `Asia/Hebron`, `Europe/Lisbon`, `Atlantic/Madeira`, `Atlantic/Azores`, `Pacific/Palau`, `America/Asuncion`, `Asia/Qatar`, `Indian/Reunion`, `Europe/Bucharest`, `Europe/Belgrade`, `Europe/Kaliningrad`, `Europe/Moscow`, `Europe/Volgograd`, `Europe/Samara`, `Asia/Yekaterinburg`, `Asia/Omsk`, `Asia/Novosibirsk`, `Asia/Novokuznetsk`, `Asia/Krasnoyarsk`, `Asia/Irkutsk`, `Asia/Yakutsk`, `Asia/Khandyga`, `Asia/Vladivostok`, `Asia/Sakhalin`, `Asia/Ust-Nera`, `Asia/Magadan`, `Asia/Kamchatka`, `Asia/Anadyr`, `Africa/Kigali`, `Asia/Riyadh`, `Pacific/Guadalcanal`, `Indian/Mahe`, `Africa/Khartoum`, `Europe/Stockholm`, `Asia/Singapore`, `Atlantic/St_Helena`, `Europe/Ljubljana`, `Arctic/Longyearbyen`, `Europe/Bratislava`, `Africa/Freetown`, `Europe/San_Marino`, `Africa/Dakar`, `Africa/Mogadishu`, `America/Paramaribo`, `Africa/Juba`, `Africa/Sao_Tome`, `America/El_Salvador`, `America/Lower_Princes`, `Asia/Damascus`, `Africa/Mbabane`, `America/Grand_Turk`, `Africa/Ndjamena`, `Indian/Kerguelen`, `Africa/Lome`, `Asia/Bangkok`, `Asia/Dushanbe`, `Pacific/Fakaofo`, `Asia/Dili`, `Asia/Ashgabat`, `Africa/Tunis`, `Pacific/Tongatapu`, `Europe/Istanbul`, `America/Port_of_Spain`, `Pacific/Funafuti`, `Asia/Taipei`, `Africa/Dar_es_Salaam`, `Europe/Kiev`, `Europe/Uzhgorod`, `Europe/Zaporozhye`, `Europe/Simferopol`, `Africa/Kampala`, `Pacific/Johnston`, `Pacific/Midway`, `Pacific/Wake`, `America/New_York`, `America/Detroit`, `America/Kentucky/Louisville`, `America/Kentucky/Monticello`, `America/Indiana/Indianapolis`, `America/Indiana/Vincennes`, `America/Indiana/Winamac`, `America/Indiana/Marengo`, `America/Indiana/Petersburg`, `America/Indiana/Vevay`, `America/Chicago`, `America/Indiana/Tell_City`, `America/Indiana/Knox`, `America/Menominee`, `America/North_Dakota/Center`, `America/North_Dakota/New_Salem`, `America/North_Dakota/Beulah`, `America/Denver`, `America/Boise`, `America/Phoenix`, `America/Los_Angeles`, `America/Anchorage`, `America/Juneau`, `America/Sitka`, `America/Yakutat`, `America/Nome`, `America/Adak`, `America/Metlakatla`, `Pacific/Honolulu`, `America/Montevideo`, `Asia/Samarkand`, `Asia/Tashkent`, `Europe/Vatican`, `America/St_Vincent`, `America/Caracas`, `America/Tortola`, `America/St_Thomas`, `Asia/Ho_Chi_Minh`, `Pacific/Efate`, `Pacific/Wallis`, `Pacific/Apia`, `Asia/Aden`, `Indian/Mayotte`, `Africa/Johannesburg`, `Africa/Lusaka`, `Africa/Harare`, `UTC`
  - Default value: `UTC`
- `timezone_variable` (String) Variable name
- `track_default_gateway` (Boolean) Enable or disable default gateway tracking
  - Default value: `true`
- `track_default_gateway_variable` (String) Variable name
- `track_interface_tag` (Number) OMP Tag attached to routes based on interface tracking
  - Range: `1`-`4294967295`
- `track_interface_tag_variable` (String) Variable name
- `track_transport` (Boolean) Configure tracking of transport
  - Default value: `true`
- `track_transport_variable` (String) Variable name
- `transport_gateway` (Boolean) Enable transport gateway
  - Default value: `false`
- `transport_gateway_variable` (String) Variable name

### Read-Only

- `id` (String) The id of the Feature
- `version` (Number) The version of the Feature

<a id="nestedatt--affinity_per_vrfs"></a>
### Nested Schema for `affinity_per_vrfs`

Optional:

- `affinity_group_number` (Number) Affinity Group Number
  - Range: `1`-`63`
- `affinity_group_number_variable` (String) Variable name
- `vrf_range` (String) Range of VRFs
- `vrf_range_variable` (String) Variable name


<a id="nestedatt--gps_sms_mobile_numbers"></a>
### Nested Schema for `gps_sms_mobile_numbers`

Optional:

- `number` (String) Mobile number, ex: 1231234414
- `number_variable` (String) Variable name

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
# Expected import identifier with the format: "system_basic_feature_id,feature_profile_id"
terraform import sdwan_system_basic_feature.example "f6b2c44c-693c-4763-b010-895aa3d236bd,f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
```
